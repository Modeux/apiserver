package entries

import (
	"github.com/pkg/errors"
	"pornrangers/entities"
	"pornrangers/pkg/databases"
)

type RegisterRepo struct {
	db databases.DBInterface
}

func NewRegisterRepo(db databases.DBInterface) *RegisterRepo {
	return &RegisterRepo{db: db}
}

func (e *RegisterRepo) CheckEmail(email string) (bool, error) {
	var exist bool
	if err := e.db.GetConn(databases.Mysql).Get(&exist, "SELECT EXISTS(SELECT 1 FROM `users` WHERE `email` = ? ORDER BY `id` LIMIT 1)", email); err != nil {
		return false, errors.Wrap(err, "check email duplicate")
	}
	return exist, nil
}

func (e *RegisterRepo) InsertUser(data entities.SignUpData) error {
	userQuery := "INSERT INTO `users` (`name`, `email`, `password`, `created_at`, `updated_at`) VALUES (:name, :email, :password, :created_at, :updated_at)"
	_, err := e.db.GetConn(databases.Mysql).NamedExec(userQuery, &data)
	if err != nil {
		return errors.Wrap(err, "insert user")
	}
	return nil
}
