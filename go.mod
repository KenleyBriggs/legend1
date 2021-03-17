module github.com/briggs/test

go 1.16

//replace github.com/briggs/myfunc => ../myfunc

//require github.com/briggs/myfunc v0.0.0-00010101000000-000000000000

require (
	github.com/cloudevents/sdk-go/v2 v2.3.1 // indirect
	github.com/denisenkom/go-mssqldb v0.9.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670 // indirect
	knative.dev/pkg v0.0.0-20210315160101-6a33a1ab29ac // indirect
)
