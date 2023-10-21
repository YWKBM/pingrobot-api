package pkg

import (
    "database/sql"
    "fmt"
  
    _ "github.com/lib/pq"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresConnection(info ConnectionInfo) (*sql.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=%s", info.Host, info.Port, info.Username, info.Password, info.DBName, info.SSLMode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil{
		return nil, err
	}

	return db, nil
}
