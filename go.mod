module github.com/briggs/test

go 1.16

//replace github.com/briggs/myfunc => ../myfunc

//require github.com/briggs/myfunc v0.0.0-00010101000000-000000000000

require github.com/go-sql-driver/mysql v1.5.0 // indirect
