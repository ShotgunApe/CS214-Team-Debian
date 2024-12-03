package student

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	StudentID       string `json:StudentID`
	FirstName       string `json:FirstName`
	LastName        string `json:LastName`
	Email           string `json:Email`
	Major           string `json:Major`
	ClassYear       string `json:ClassYear`
	DateOfBirth     string `json:DateOfBirth"`
	Gender          string `json:Gender`
	PhoneNumber     string `json:PhoneNumber`
	Address         string `json:Address`
	CalculatedGrade string `json:CalculatedGrade`
	CumulativeGPA   string `json:CumulativeGPA`
	SemesterGPA     string `json:SemesterGPA`
}

type instructor struct {
	InstructorID string `json:InstructorID`
	FirstName    string `json:FirstName`
	LastName     string `json:LastName`
	Email        string `json:Email`
	Department   string `json:Department`
}

type connection struct {
	Clients    []*student    `json:"Clients"`
	ClientsTwo []*instructor `json:"ClientsTwo"`
}

func Init() {
	db, err := sql.Open("mysql", "root:4542@tcp(localhost)/College_Internal_System") //needs to be changed for your database config
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic (no i dont think i will)
	}
	defer db.Close()
}

func GetStudents() []byte {
	db, err := sql.Open("mysql", "root:4542@tcp(localhost)/College_Internal_System") //needs to be changed for your database config
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic (no i dont think i will)
	}
	defer db.Close()
	// Execute the query
	rows, err := db.Query("SELECT * FROM Students")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// init json
	var students []*student

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.

		// should probably marshall data as json to be sent when called for and returned!
		students = append(students, &student{
			StudentID:       string(values[0]),
			FirstName:       string(values[1]),
			LastName:        string(values[2]),
			Email:           string(values[3]),
			Major:           string(values[4]),
			ClassYear:       string(values[5]),
			DateOfBirth:     string(values[6]),
			Gender:          string(values[7]),
			PhoneNumber:     string(values[8]),
			Address:         string(values[9]),
			CalculatedGrade: string(values[10]),
			CumulativeGPA:   string(values[11]),
			SemesterGPA:     string(values[12]),
		})
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app - nah
	}

	// now professors!
	var instructors []*instructor

	rows, err = db.Query("SELECT * FROM Instructors")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Get column names
	columns, err = rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values = make([]sql.RawBytes, len(columns))

	scanArgs = make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.

		// should probably marshall data as json to be sent when called for and returned!
		instructors = append(instructors, &instructor{
			InstructorID: string(values[0]),
			FirstName:    string(values[1]),
			LastName:     string(values[2]),
			Email:        string(values[3]),
			Department:   string(values[4]),
		})
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app - nah
	}

	p, _ := json.Marshal(connection{Clients: students, ClientsTwo: instructors})
	fmt.Printf("%s\n", p)
	return p
}
