package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5433
	user = "postgres"
	password = "root"
	dbname = "test_go"
)

var (
	db *sql.DB

	err error
)

type Employee struct {
	ID int
	Full_name string
	Email string
	Age int
	Division string
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *`

	err = db.QueryRow(sqlStatement, "Fahril Rizal", "admin@gmail.com", 20, "IT").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Data pegawai baru : %v\n", employee)
}

func GetEmployees() {
	var result = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}

	fmt.Println("Data Pegawai:", result)
}

func UpdateEmployee() {
	sqlStatement := `UPDATE employees 
	SET full_name = $2, email = $3, age = $4, division = $5
	WHERE id = $1`

	res, err := db.Exec(sqlStatement, 1 , "Mochammad Fahril Rizal", "mfahrilrizal@gmail.com", 25, "Management")
	
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Update data:", count)
}

func DeleteEmployee() {
	sqlStatement := `DELETE FROM employees WHERE id = $1`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hapus data:", count)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB berhasil terhubung")

	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()
	// DeleteEmployee()
}