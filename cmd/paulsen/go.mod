module github.com/zeropen/app/paulsen

go 1.22.1

replace github.com/zeropen/app/spector => ../spector

replace github.com/zeropen/app/sazs => ../sazs

replace github.com/zeropen/app/benghazi/postgresdriver => ../benghazi/postgresdriver

replace github.com/zeropen => ../../

require (
	github.com/joho/godotenv v1.5.1
	github.com/zeropen/app/benghazi/postgresdriver v0.0.0-00010101000000-000000000000
	github.com/zeropen/app/spector v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/zeropen v0.0.0-00010101000000-000000000000 // indirect
	github.com/zeropen/app/sazs v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gorm.io/driver/postgres v1.5.11 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
