package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/prometheus/common/log"
)

const (
	envUsername = "SQL_USERNAME"
	envPassword = "SQL_PASSWORD"
)

type Adapter struct {
	Three int
	Six   int
	Ten   int

	Username string
	Password string
	Server   string
	Port     int
	User     string

	Database string

	db *sql.DB
}

func init() {

}

type PricingData struct {
	Three int `json:"three"`
	Six   int `json:"six"`
	Ten   int `json:"ten"`
}

func main() {

	a := &Adapter{}

	a.Username = os.Getenv(envUsername)
	a.Password = os.Getenv(envPassword)
	a.Server = os.Getenv("SQL_SERVER")
	a.Port = 1433 //os.Getenv("SQL_PORT")
	a.Database = os.Getenv("SQL_DATABASE")

	// build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", a.Server, a.User, a.Password, a.Port, a.Database)

	fmt.Println(connString)

	var err error
	// create connection pool
	a.db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(a.Username)

	log.Info("Starting the HTTP server now..")
	http.HandleFunc("/pricing", a.getPricing)
	http.HandleFunc("/insertrow", a.insertNewRow)
	http.ListenAndServe(":8080", nil)

	// DataBase()
}

func (a *Adapter) getPricing(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error occured reading body: %v", err)
		return
	}

	i := &PricingData{}
	if err := json.Unmarshal(body, i); err != nil {
		log.Errorf("Error occured unmarsaling data: %v", err)
		return
	}

	// call sql for updated values

	a.Three = 12
	a.Ten = 122
	a.Six = 322

	a.MetalLegendPricing(i.Three, i.Six, i.Ten)

	fmt.Fprintf(w, "got pricing!\n")
}

func (a *Adapter) MetalLegendPricing(three int, six int, ten int) {

	metal28GaTuffRib := ((three * a.Three) + (six * a.Six) + (ten * a.Ten))
	metal26GaTuffRib := ((three * 975) + (six * 1005) + (ten * 1035))
	metal26Ga100NS := ((three * 1050) + (six * 1580) + (ten * 1610))
	metal24Ga100NS := ((three * 1650) + (six * 1680) + (ten * 1710))

	fmt.Print(metal28GaTuffRib)
	metal28GaTuffRibString := FormatCurrency(metal28GaTuffRib)
	metal26GaTuffRibString := FormatCurrency(metal26GaTuffRib)
	metal26Ga100NSString := FormatCurrency(metal26Ga100NS)
	metal24Ga100NSString := FormatCurrency(metal24Ga100NS)

	fmt.Println("Legend Metal Pricing:")
	fmt.Println("28 gauge Tuff Rib:", metal28GaTuffRibString)
	fmt.Println("26 gauage Tuff Rib:", metal26GaTuffRibString)
	fmt.Println("26 guage 1 inch Nail Strip Standing Seam:", metal26Ga100NSString)
	fmt.Println("24 gauge 1 inch Nail Strip Standing Seam:", metal24Ga100NSString)

}

func (a *Adapter) DataBase() {

	// fmt.Println("Connected!")
	// fmt.Println()
	// fmt.Println("What would you like to do?")
	// fmt.Println()
	// fmt.Println("Available options are:")
	// fmt.Println("Insert")
	// fmt.Println("Update")
	// fmt.Println("Delete")
	// fmt.Println("Exit")
	// fmt.Println()
	// var whatToDo string
	// fmt.Scanln(&whatToDo)

	// if whatToDo == "Insert" {
	// 	// create new product
	// 	var prodName string
	// 	var three int
	// 	var six int
	// 	var ten int

	// 	fmt.Println()
	// 	fmt.Println("Enter your product/system name:")
	// 	fmt.Scanln(&prodName)
	// 	fmt.Println()
	// 	fmt.Println("Enter your pricing as a whole number.")
	// 	fmt.Println("Example:")
	// 	fmt.Println("$8.87/lf is entered as 887")
	// 	fmt.Println("Enter your three twelve pricing:")
	// 	fmt.Scanln(&three)
	// 	fmt.Println("Enter your six twelve pricing:")
	// 	fmt.Scanln(&six)
	// 	fmt.Println("Enter your ten twelve pricing:")
	// 	fmt.Scanln(&ten)

	// 	createProd, err := CreateProduct(prodName, three, six, ten)
	// 	if err != nil {
	// 		log.Fatal("Error creating product: ", err.Error())
	// 	}

	// 	fmt.Printf("Inserted UID: %d successfully.\n", createProd)
	// 	fmt.Println("Your product UIDs are:")

	// 	ReadProduct()
	// 	if err != nil {
	// 		log.Fatal("Error reading products: ", err.Error())
	// 	}

	// 	fmt.Println("Don't forget them!")
	// }

	// if whatToDo == "Update" {
	// 	// create new product
	// 	var prodName string
	// 	var three int
	// 	var six int
	// 	var ten int
	// 	var UID int

	// 	fmt.Println()
	// 	fmt.Println("Enter the UID of the product you would like to update.")
	// 	fmt.Scanln(&UID)
	// 	fmt.Println()
	// 	fmt.Println("Sorry I'm not flexible, please enter the entire set of information")
	// 	fmt.Println("Enter your product/system name:")
	// 	fmt.Scanln(&prodName)
	// 	fmt.Println("Enter your pricing as a whole number.")
	// 	fmt.Println("Example:")
	// 	fmt.Println("$8.87/lf is entered as 887")
	// 	fmt.Println("Enter your three twelve pricing:")
	// 	fmt.Scanln(&three)
	// 	fmt.Println("Enter your six twelve pricing:")
	// 	fmt.Scanln(&six)
	// 	fmt.Println("Enter your ten twelve pricing:")
	// 	fmt.Scanln(&ten)

	// 	updatedRows, err := UpdateProduct(prodName, three, six, ten, UID)
	// 	if err != nil {
	// 		log.Fatal("Error updating product: ", err.Error())
	// 	}
	// 	fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

	// }
	// if whatToDo == "Delete" {

	// 	var UID int

	// 	fmt.Println("What product UID would you like to have deleted. REMEMBER this operation cannot be undone.")
	// 	fmt.Scanln(&UID)
	// 	deletedRows, err := DeleteProduct(UID)
	// 	if err != nil {
	// 		log.Fatal("Error deleting product: ", err.Error())
	// 	}
	// 	fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)

	// }

	// if whatToDo == "Exit" {
	// 	rows, err := db.Query("SELECT * FROM metal", 1)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer rows.Close()
	// } else {
	// 	fmt.Println("Invalid Entry Try Again")
	// 	DataBase()
	// }

}

type RowDataPayload struct {
	Pd             PricingData
	ProductionName string `json:"prodName"`
}

func (a *Adapter) insertNewRow(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error occured reading body: %v", err)
		return
	}

	i := &RowDataPayload{}
	if err := json.Unmarshal(body, i); err != nil {
		log.Errorf("Error occured unmarsaling data: %v", err)
		return
	}

	createProd, err := CreateProduct(i.ProductionName, i.Pd.Three, i.Pd.Six, i.Pd.Ten)
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
