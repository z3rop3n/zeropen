package postgresdriver

// type PostgresTokenQuery struct {
// 	db *gorm.DB
// }

// func NewPostgresTokenQuery(db *gorm.DB) *PostgresTokenQuery {
// 	return &PostgresTokenQuery{db: db}
// }

// func (p PostgresTokenQuery) GetToken(_ string) (types.Token, error) {
// 	var token types.Token
// 	if err := p.db.First(&token).Error; err != nil {
// 		return token, err
// 	} else {
// 		return types.Token{}, err
// 	}
// }
