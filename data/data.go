package data

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq" // postgresql driver
)

func dbFunc(db *sql.DB) string {

	rdo := "se conecto satisfactoriamente"

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
		rdo += fmt.Sprintf("Error creating database table: %q", err)
	}
	rdo += " \n tabla creada"

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		rdo += fmt.Sprintf("Error incrementing tick: %q", err)
	}

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		rdo += fmt.Sprintf("Error incrementing tick: %q", err)
	}

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		rdo += fmt.Sprintf("Error incrementing tick: %q", err)
	}

	rdo += " \n registros insertado"

	rows, err := db.Query("SELECT tick FROM ticks")
	if err != nil {

		rdo += fmt.Sprintf("Error reading ticks: %q", err)
	}

	defer rows.Close()
	for rows.Next() {
		var tick time.Time
		if err := rows.Scan(&tick); err != nil {

			rdo += fmt.Sprintf("Error scanning ticks: %q", err)

		}

		rdo += fmt.Sprintf("Read from DB: %s\n", tick.String())
	}

	return rdo

}

// Abrir  se utiliza para inicializar la base de datos
func Abrir() string {

	// process.env.DATABASE_URL
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	psqlInfo := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return dbFunc(db) + os.Getenv("PORT") + os.Getenv("DATABASE_URL")

}
