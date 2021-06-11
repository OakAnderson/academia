package repository

import (
	"database/sql"
	"fmt"
	"github.com/OakAnderson/academia/data/model"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
 	err error
	dbhost string = os.Getenv("ACADEMIA_HOST_DB")
)

func FindById(id int) model.Cliente {
	var row = execQueryRow(fmt.Sprintf("SELECT * FROM clientes WHERE id = %d", id ))

	var cliente = model.Cliente{}

	err = cliente.Scan(row)
	if err != nil { fmt.Println(err.Error()) }

	return cliente
}

func FindAll() []model.Cliente {
	var rows, err = execQuery("SELECT * FROM clientes")
	if err != nil { fmt.Println(err.Error())}

	var clientes []model.Cliente

	if rows != nil {
		for rows.Next() {
			var c = model.Cliente{}

			err = c.Scan(rows)
			if err != nil {
				fmt.Println(err.Error())
			}

			clientes = append(clientes, c)
		}
	}

	return clientes
}

func execQueryRow(query string) *sql.Row {
	db, err = sql.Open("mysql", dbhost)
	defer db.Close()
	if err != nil { fmt.Println(err.Error()) }

	return db.QueryRow(query)
}

func execQuery(query string) (*sql.Rows, error) {
	db, err = sql.Open("mysql", dbhost)
	defer db.Close()
	if err != nil { fmt.Println(err.Error()) }

	return db.Query(query)
}
