package repo

import (
	"database/sql"
	"fmt"

	"user_service/models"
)

type IUserRepo interface {
	GetUserByEmail(email string) (models.User, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) GetUserByEmail(email string) (models.User, error) {

	var user models.User
	row := ur.db.QueryRow("Select id, username, password From users WHERE email = $1", email)

	if err := row.Scan(&user.Email, &user.Password, &user.Id); err != nil {
		fmt.Println("Error Occured")
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("email %v is not linked to any user", email)
		}
		return user, fmt.Errorf("userByEmail %v: %v", email, err)
	}
	return user, nil

}
