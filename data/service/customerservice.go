package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"restapi/app/common"
	"restapi/data/model"

	"github.com/gorilla/mux"
)

//GetAllCustomer is...
func GetAllCustomer() []model.Customer {
	db := common.GetDatabase()
	defer db.Close()
	var customers []model.Customer
	rows, err := db.Query("SELECT *FROM customer")
	common.CheckError(err)
	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(&customer.Customerid, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Dateofbirth, &customer.Mobilenumber)
		common.CheckError(err)
		customers = append(customers, customer)
	}
	return customers
}

//SaveOneCustomer is...
func SaveOneCustomer(r *http.Request) {
	db := common.GetDatabase()
	defer db.Close()
	fmt.Println("save customer")
	query := "INSERT INTO customer (firstname,lastname,email,dateofbirth,mobilenumber) VALUES($1,$2,$3,$4,$5)"
	customer := model.Customer{
		FirstName:    r.FormValue("firstname"),
		LastName:     r.FormValue("lastname"),
		Email:        r.FormValue("email"),
		Dateofbirth:  r.FormValue("dateofbirth"),
		Mobilenumber: r.FormValue("mobilenumber"),
	}
	_, err := db.Exec(query, customer.FirstName, customer.LastName, customer.Email, customer.Dateofbirth, customer.Mobilenumber)
	common.CheckError(err)
}

//GetOneCustomer is...
func GetOneCustomer(r *http.Request) (model.Customer, error) {
	var customer model.Customer
	db := common.GetDatabase()
	defer db.Close()
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	query := "SELECT *FROM customer WHERE id=" + id
	row := db.QueryRow(query)
	err := row.Scan(&customer.Customerid, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Dateofbirth, &customer.Mobilenumber)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

//DeleteOneCustomer is...
func DeleteOneCustomer(r *http.Request) {
	fmt.Println("delete one customer")
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	db := common.GetDatabase()
	defer db.Close()
	query := "DELETE FROM CUSTOMER WHERE id=$1"
	rows, err := db.Exec(query, id)
	common.CheckError(err)
	_, err = rows.RowsAffected()
	common.CheckError(err)
}

//UpdateOneCustomer is...
func UpdateOneCustomer(r *http.Request) {
	id := mux.Vars(r)["id"]
	var customer model.Customer
	body, err := ioutil.ReadAll(r.Body)
	common.CheckError(err)
	err = json.Unmarshal(body, &customer)
	common.CheckError(err)
	customer.Customerid = id
	query := "UPDATE customer SET firstname=$1 , lastname=$2 , email=$3 , dateofbirth=$4 , mobilenumber=$5 WHERE id=$6"
	db := common.GetDatabase()
	defer db.Close()
	rows, err := db.Exec(query, customer.FirstName, customer.LastName, customer.Email, customer.Dateofbirth, customer.Mobilenumber, customer.Customerid)
	common.CheckError(err)
	_, err = rows.RowsAffected()
	common.CheckError(err)
}
