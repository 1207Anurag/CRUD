package main

import (
	"database/sql"
	//"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//check for error
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//emp struct

type emp struct {
	id                int
	name, email, role string
}

//open connection
func dbCon() *sql.DB {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/go_db")
	CheckErr(err)
	return db
}

func main() {
	db := dbCon()
	res, _ := GetById(db, 1)
	fmt.Println(res)
	
	err := UpdateById(db, "Tarun", 1)
	if err != nil {
		log.Fatal(err)
	}
	err:=InsertData(db,"Anurag Chaubey","anurag@zopsmart.com","intern")
	if err!=nil{
		log.Fatal(err)
	}
	

}

func InsertData(db *sql.DB, name string, email string, role string) error {
	query := "INSERT INTO employee (Name,Email,role) values (?,?,?);"

	res, err := db.Prepare(query)

	if err != nil {
		return err
	}
	_, err = res.Exec(name, email, role)

	return err
}

// GetByID
// func GetById(id int) *sql.Rows{
// db:=dbCon()
// q:="SELECT * FROM employee WHERE id=?"
// res,e:=db.Query(q,id)
// CheckErr(e)
// fmt.Println("A row deleted")
// defer db.Close()
// return res

// }

func GetById(db *sql.DB, id int) (*emp, error) {

	var u emp

	q := "SELECT * FROM employee WHERE id=?"
	res, e := db.Query(q, id)
	if e != nil {
		return nil, e
	}
	defer res.Close()
	for res.Next() {

		e = res.Scan(&u.id, &u.name, &u.email, &u.role)
		if e != nil {
			return nil, e
		}

		// fmt.Println(u.id,u.name,u.email,u.role)

	}
	return &u, nil

}

//update
func UpdateById(db *sql.DB, name string, id int) error {
	q := "UPDATE employee SET name=? WHERE id=?"

	_, e := db.Exec(q, name, id)
	//CheckErr(e)
	if e != nil {
		return e
	}
	fmt.Println("A row is updated")

	return nil
}

//delete

func RemoveById(db *sql.DB, id int) error {

	q := "DELETE FROM employee WHERE id=?"

	_, e := db.Exec(q, id)
	if e != nil {
		return e
	}
	//defer del.Close()
	return nil
}
