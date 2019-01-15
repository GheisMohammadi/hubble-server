package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Postgre adapter
type Postgre struct {
	Host     string  //"localhost"
	Port     int     //5432
	User     string  //"postgres"
	Password string  //"your-password"
	DBname   string  //"calhounio_demo"
	ObjDB    *sql.DB //Opened DB
}

//Connect to database
func (obe *Postgre) Connect() error {
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
func (obe *Postgre) Disconnect() {
	obe.ObjDB.Close()
}

//InsertAccount add new Account to accounts table
func (obe *Postgre) InsertAccount(acc *Account) error {

	sqlStatement := `INSERT INTO accounts (address, public_key, balance, permission,sequence,code)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`
	id := 0
	err := obe.ObjDB.QueryRow(sqlStatement, acc.Address, acc.PublicKey, acc.Balance, acc.Permission, acc.Sequence, acc.Code).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println("New record ID is:", id)
	return nil
}

//UpdateAccount modifies all fields for selected account
func (obe *Postgre) UpdateAccount(id int, acc *Account) error {
	sqlStatement := `UPDATE accounts
				SET address = $2, public_key = $3, balance = $4, permission = $5, sequence = $6, code = $7
				WHERE id = $1
				RETURNING id, address;`
	var retAddress string
	var retID int
	err := obe.ObjDB.QueryRow(sqlStatement, id, acc.Address, acc.PublicKey, acc.Balance, acc.Permission, acc.Sequence, acc.Code).Scan(&retID, &retAddress)

	if err != nil {
		return err
	}

	return nil
}

//GetAccount finds account in db and returns its data
func (obe *Postgre) GetAccount(id int) (*Account, error) {
	sqlStatement := `SELECT * FROM accounts 
					 WHERE id=$1;`
	acc := &Account{Address: "", PublicKey: "", Balance: 0.0, Permission: "", Sequence: 0, Code: ""}
	row := obe.ObjDB.QueryRow(sqlStatement, id)
	err := row.Scan(&acc.Address, &acc.PublicKey, &acc.Balance, &acc.Permission, &acc.Sequence, &acc.Code, &acc.ID)
	switch err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return acc, nil
	default:
		return nil, err
	}
}

//GetBlocksTableLastID returns last block number
func (obe *Postgre) GetBlocksTableLastID() (uint64, error) {
	sqlStatement := `SELECT MAX(id) FROM blocks`
	row := obe.ObjDB.QueryRow(sqlStatement)
	var LastID uint64
	err := row.Scan(LastID)
	switch err {
	case sql.ErrNoRows:
		return 0, err
	case nil:
		return LastID, nil
	default:
		return 0, err
	}
}
