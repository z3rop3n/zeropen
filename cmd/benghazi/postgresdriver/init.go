package postgresdriver

import (
	"fmt"
	// "go/types"

	sdb "github.com/zeropen/app/sazs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresDB struct {
	DB       *gorm.DB
	Host     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	User     string
}

func NewPostgresDB(Host string, User string, Password string, DBName string, Port string, SSLMode string,
) *PostgresDB {
	return &PostgresDB{
		Host:     Host,
		Password: Password,
		DBName:   DBName,
		Port:     Port,
		User:     User,
		SSLMode:  SSLMode,
	}
}

func (p *PostgresDB) Connect() (*sdb.Config, error) {
	if p.DB != nil {
		return nil, nil
	}
	// Connect to Postgres
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", p.Host, p.User, p.Password, p.DBName, p.Port, p.SSLMode)
	// dsn := "host=127.0.0.1 user=user password=password dbname=surakshadb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	p.DB = db
	if err != nil {
		fmt.Println("Failed to connect to Postgres")
		return nil, err
	}

	fmt.Println("Connected to Postgres")
	config := sdb.Config{
		UserQuery:         NewPostgresUserQuery(db),
		OTPQuery:          NewPostgresOTPQuery(db),
		RefreshTokenQuery: NewPostgresRefreshTokenQuery(db),
	}
	return &config, nil
}
