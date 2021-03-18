package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "briggs-roofing-pricing.database.windows.net"
var port = 1433
var database = "Legend"

func CreateProduct(prod string, three, six, ten int) (int64, error) {

	ctx := context.Background()
	var err error

	if db == nil {
		err = errors.New("CreateProduct: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
    INSERT INTO dbo.metal (PROD_NAME, THREE_TWELVE, SIX_TWELVE, TEN_TWELVE) VALUES (@PROD_NAME, @THREE_TWELVE, @SIX_TWELVE, @TEN_TWELVE);
    select isNull(SCOPE_IDENTITY(), -1);
    `

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("PROD_NAME", prod),
		sql.Named("THREE_TWELVE", three),
		sql.Named("SIX_TWELVE", six),
		sql.Named("TEN_TWELVE", ten))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func ReadProduct() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT UID FROM dbo.metal;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var prodName string
		var id int

		// Get values from row.
		err := rows.Scan(&id)
		if err != nil {
			return -1, err
		}

		fmt.Printf("UID: %d, Product name: %s\n", id, prodName)
		count++
	}

	return count, nil
}

func UpdateProduct(prod string, three int, six int, ten int, UID int) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE dbo.metal SET PROD_NAME = @PROD_NAME, THREE_TWELVE = @THREE_TWELVE, SIX_TWELVE = @SIX_TWELVE, TEN_TWELVE = @TEN_TWELVE WHERE UID = @UID")

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("PROD_NAME", prod),
		sql.Named("THREE_TWELVE", three),
		sql.Named("SIX_TWELVE", six),
		sql.Named("UID", UID),
		sql.Named("TEN_TWELVE", ten))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func DeleteProduct(UID int) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM dbo.metal WHERE UID = @UID;")

	// Execute non-query with named parameters
	result, err := db.ExecContext(ctx, tsql, sql.Named("UID", UID))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
