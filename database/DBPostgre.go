package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBPostgre struct {
	Host     string  //"localhost"
	Port     int     //5432
	User     string  //"postgres"
	Password string  //"your-password"
	DBname   string  //"calhounio_demo"
	ObjDB    *sql.DB //Opened DB
}

//Connect to database
func (obe *DBPostgre) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", obe.Host, obe.Port, obe.User, obe.Password, obe.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		db.Close()
		return err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return err
	}
	obe.ObjDB = db
	return nil
}

//Disconnect close connection to database
func (obe *DBPostgre) Disconnect() {
	obe.ObjDB.Close()
}

//InsertAccount add new Account to accounts table
func (obe *DBPostgre) InsertAccount(address string, publicKey string, balance float64, perm string,
	seq string, code string) error {

	sqlStatement := `INSERT INTO accounts (address, public_key, balance, permission,sequence,code)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`
	id := 0
	err := obe.ObjDB.QueryRow(sqlStatement, "Addr123", "ABC", 1000, "Perm456", "Seq123", "CodeF1F2").Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println("New record ID is:", id)
	return nil
}
