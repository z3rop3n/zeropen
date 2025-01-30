package postgresdriver

import (
	"fmt"
	// "go/types"

	sdb "github.com/zeropen/app/sazs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (p *PostgresDB) Connect() (*sdb.Config, error) {
	if p.DB != nil {
		return nil, nil
	}
	// Connect to Postgres
	dsn := "host=127.0.0.1 user=user password=password dbname=surakshadb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	p.DB = db
	if err != nil {
		fmt.Println("Failed to connect to Postgres")
		return nil, err
	}
	fmt.Println("Connected to Postgres")
	config := sdb.Config{
		UserQuery: NewPostgresUserQuery(db),
		OTPQuery:  NewPostgresOTPQuery(db),
	}
	return &config, nil
}
