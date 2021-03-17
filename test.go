package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"golang.org/x/crypto/ssh/terminal"
	"log"
)

var db *sql.DB

var server = "briggs-roofing-pricing.database.windows.net"
var port = 1433
var user string
var password string
var database = "Legend"

func DataBase() {

	var err error

	// sets username and password
	fmt.Println("Please enter your username:")
	fmt.Scanln(&user)

	fmt.Println("Please enter your password:")
	password, err := terminal.ReadPassword(0)
	if err == nil {
	}

	// build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	// create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected!")
	fmt.Println()
	fmt.Println("What would you like to do?")
	fmt.Println()
	fmt.Println("Available options are:")
	fmt.Println("Insert")
	fmt.Println("Update")
	fmt.Println("Delete")
	fmt.Println("Exit")
	fmt.Println()
	var whatToDo string
	fmt.Scanln(&whatToDo)

	if whatToDo == "Insert" {
		// create new product
		var prodName string
		var three int
		var six int
		var ten int

		fmt.Println()
		fmt.Println("Enter your product/system name:")
		fmt.Scanln(&prodName)
		fmt.Println()
		fmt.Println("Enter your pricing as a whole number.")
		fmt.Println("Example:")
		fmt.Println("$8.87/lf is entered as 887")
		fmt.Println("Enter your three twelve pricing:")
		fmt.Scanln(&three)
		fmt.Println("Enter your six twelve pricing:")
		fmt.Scanln(&six)
		fmt.Println("Enter your ten twelve pricing:")
		fmt.Scanln(&ten)

		createProd, err := CreateProduct(prodName, three, six, ten)
		if err != nil {
			log.Fatal("Error creating product: ", err.Error())
		}

		fmt.Printf("Inserted UID: %d successfully.\n", createProd)
		fmt.Println("Your product UIDs are:")

		ReadProduct()
		if err != nil {
			log.Fatal("Error reading products: ", err.Error())
		}

		fmt.Println("Don't forget them!")
	}

	if whatToDo == "Update" {
		// create new product
		var prodName string
		var three int
		var six int
		var ten int
		var UID int

		fmt.Println()
		fmt.Println("Enter the UID of the product you would like to update.")
		fmt.Scanln(&UID)
		fmt.Println()
		fmt.Println("Sorry I'm not flexible, please enter the entire set of information")
		fmt.Println("Enter your product/system name:")
		fmt.Scanln(&prodName)
		fmt.Println("Enter your pricing as a whole number.")
		fmt.Println("Example:")
		fmt.Println("$8.87/lf is entered as 887")
		fmt.Println("Enter your three twelve pricing:")
		fmt.Scanln(&three)
		fmt.Println("Enter your six twelve pricing:")
		fmt.Scanln(&six)
		fmt.Println("Enter your ten twelve pricing:")
		fmt.Scanln(&ten)

		updatedRows, err := UpdateProduct(prodName, three, six, ten, UID)
		if err != nil {
			log.Fatal("Error updating product: ", err.Error())
		}
		fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

	}
	if whatToDo == "Delete" {

		var UID int

        fmt.Println("What product UID would you like to have deleted. REMEMBER this operation cannot be undone.")
        fmt.Scanln(&UID)
		deletedRows, err := DeleteProduct(UID)
		if err != nil {
			log.Fatal("Error deleting product: ", err.Error())
		}
		fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)

	}

    if whatToDo == "Exit" {
        rows, err := db.Query("SELECT * FROM metal", 1)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()
    }
}

func CreateProduct(prod string, three int, six int, ten int) (int64, error) {

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
