package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() *sql.DB {
	var db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	return db
}

func sqlQuery() {
	var db = connect()
	defer db.Close()

	var age = 27
	var rows, err = db.Query("select id, name, grade from tb_student where age = ?", age)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func sqlQueryRow() {
	var db = connect()
	defer db.Close()

	var result = student{}
	var id = "E001"
	var err = db.QueryRow("select name, grade from tb_student where id = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("name: %sngrade: %d\n", result.name, result.grade)
}

/*
Eksekusi Query Menggunakan Prepare()
Teknik prepared statement adalah teknik penulisan query di awal dengan kelebihan bisa di
re-use atau digunakan banyak kali untuk eksekusi yang berbeda-beda.
Metode ini bisa digabung dengan Query() maupun QueryRow().
Berikut merupakan contoh penerapannya
*/

func sqlPrepare() {
	var db = connect()
	defer db.Close()

	var statement, err = db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	var result1 = student{}
	statement.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %sgrade: %d\n", result1.name, result1.grade)

	var result2 = student{}
	statement.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %sgrade: %d\n", result2.name, result2.grade)

	var result3 = student{}
	statement.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %sgrade: %d\n", result3.name, result3.grade)
}

// insert, update & delete menggunakan Exec()

func sqlExec() {
	var db = connect()
	defer db.Close()

	var err error

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G002", "Lamborgini", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G002")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec("delete from tb_student where id = ?", "G002")
	if err != nil {
		fmt.Println(err.Error())
	}

	/*
	// menggunakan metode prepared statement
	var stmt, _ = db.Prepare("insert into tb_student values (?, ?, ?, ?)")stmt.Exec("G001", "Galahad", 29, 2)
	// menggunakan metode biasa
	_, _ = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	*/
}

func main() {
	fmt.Println("SqlQuery->"); sqlQuery()
	fmt.Println("SqlQueryRow->"); sqlQueryRow()
	fmt.Println("SqlPrepare->"); sqlPrepare()
	fmt.Println("SqlExec CRUD->"); sqlExec()
}
