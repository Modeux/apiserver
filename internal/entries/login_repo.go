package entries

import (
	"database/sql"
	"github.com/pkg/errors"
	"pornrangers/entities"
	"pornrangers/pkg/databases"
)

type LoginRepo struct {
	db databases.DBInterface
}

func NewLoginRepo(db databases.DBInterface) *LoginRepo {
	return &LoginRepo{db: db}
}

func (l *LoginRepo) GetUserByEmail(email string) (entities.UserLogin, error) {
	var user entities.UserLogin
	query := "SELECT `id`, `name`, `email`, `password` FROM `users` WHERE `email` = ? ORDER BY `id` DESC LIMIT 1"
	if err := l.db.GetConn(databases.Mysql).Get(&user, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.UserLogin{}, nil
		}
		return user, errors.Wrap(err, "Login query")
	}
	return user, nil
}
