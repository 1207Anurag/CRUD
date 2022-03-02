package main

import (
        "database/sql"
        "encoding/json"
        "fmt"
       // "log"
        "net/http"
        "strconv"
        "github.com/gorilla/mux"
      CRUD "github.com/1207Anurag/CRUD"
      
       
)

type emp struct {
        Id    string
        Name  string
        Email string
        Role  string
}

// type Handler struct {
//         DB *sql.DB
// }

func HomeLink(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome Home!")
}

//post request

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        var employee emp
      _ = json.NewDecoder(r.Body).Decode(&employee)

        CRUD.InsertData(employee, DB, employee.Name, employee.Email, employee.Role)
        json.NewEncoder(w).Encode(employee)

}

//read from employees

func GetEmployee(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        empID := vars["id"]
        empIDint, _ := strconv.Atoi(empID)
        employee, err := CRUD.GetById(DB, empIDint)
        if err!=nil{
                fmt.Println("Error")
        }
        json.NewEncoder(w).Encode(employee)
}



func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
        empID := mux.Vars(r)["id"]
        empIDint, _ := strconv.Atoi(empID)
        CRUD.RemoveById(DB, empIDint)
        fmt.Fprintf(w, "The employee with ID %v has been deleted successfully", empID)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
        empID := mux.Vars(r)["id"]
        empIDint, _ := strconv.Atoi(empID)
        var updatedEmployee emp

        _ = json.NewDecoder(r.Body).Decode(&updatedEmployee)
        CRUD.UpdateById(DB, empIDint, updatedEmployee.Name, updatedEmployee.Email, updatedEmployee.Role)
        fmt.Fprintf(w, "The employee with ID %v has been updated successfully", empID)

}

var DB *sql.DB

func main() {

        DB = CRUD.DbConn("Employee_Db")
        router := mux.NewRouter()
        router.HandleFunc("/", HomeLink)
        router.HandleFunc("/employee", CreateEmployee).Methods("POST")
        router.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
        router.HandleFunc("/employees/{id}", UpdateEmployee).Methods("PATCH")
        router.HandleFunc("/employees/{id}", DeleteEmployee).Methods("DELETE")

        http.ListenAndServe(":8080", router)

}