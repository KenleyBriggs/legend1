package azuresqltarget

import (
	"context"
	"errors"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"knative.dev/pkg/logging"

	libadapter "github.com/triggermesh/knative-targets/pkg/adapter"
	targetce "github.com/triggermesh/knative-targets/pkg/adapter/cloudevents"
	"github.com/triggermesh/knative-targets/pkg/apis/targets/v1alpha1"
)


type TsqlEvent struct {
	Query string `json:"query"`
}

// NewTarget creates a Azure Sql target adapter
func NewTarget(ctx context.Context, envAcc libadapter.EnvConfigAccessor, ceClient cloudevents.Client) libadapter.Adapter {
	logger := logging.FromContext(ctx)
	env := envAcc.(*envAccessor)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		env.ServerURL, env.ServerUser, env.ServerPassword, env.ServerPort, env.ServerDatabase)
	var err error
	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		logger.Errorf("Error creating connection pool: ", err.Error())
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	logger.Info("Connected!")

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStaticResponseType(v1alpha1.EventTypeSalesforceAPICallResponse))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	return &azureSQLAdapter{
		sqlClient: db,

		replier:  replier,
		ceClient: ceClient,
		logger:   logger,
	}
}

var _ libadapter.Adapter = (*azureSQLAdapter)(nil)

type azureSQLAdapter struct {
	sqlClient *sql.DB

	replier  *targetce.Replier
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
}

func (a *azureSQLAdapter) Start(ctx context.Context) error {
	a.logger.Info("Starting Azure SQL adapter")

	if err := a.ceClient.StartReceiver(ctx, a.dispatch); err != nil {
		return err
	}
	return nil
}

func (a *azureSQLAdapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {

	switch event.Type() {
	case "io.triggermesh.azuresql.tsql":
		id, err := a.tsqlQuery(event)
		if err != nil {
			return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
		}
		return a.replier.Ok(&event, id)
	default:
		a.logger.Errorf("Event type %q is not supported", event.Type())
		return nil, cloudevents.ResultNACK
	}
}

func (a *azureSQLAdapter) tsqlQuery(event cloudevents.Event) (int64, error) {
	a.logger.Info("tsqlQeury called")
	tsql := &TsqlEvent{}
	if err := event.DataAs(tsql); err != nil {
		a.logger.Errorw("Error processing incoming event data ", zap.Error(err))
		return -1, err
	}

	ctx := context.Background()
	var err error

	if a.sqlClient == nil {
		err = errors.New("db is null")
		return -1, err
	}

	// Check if database is alive.
	err = a.sqlClient.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	stmt, err := a.sqlClient.Prepare(tsql.Query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	row := stmt.QueryRowContext(ctx)
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}
