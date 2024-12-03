package student

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	db, err := sql.Open("mysql", "root:4542@tcp(localhost)/College_Internal_System")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic (no i dont think i will)
	}
	defer db.Close()
}
