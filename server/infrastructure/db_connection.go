package infrastructure

import "database/sql"

func GetSqliteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}

func CloseSQLDB(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
}
