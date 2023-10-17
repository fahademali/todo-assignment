package repo

import (
	"database/sql"
	"fmt"

	"user_service/models"
)

type IUserRepo interface {
	GetUserByEmail(email string) (models.User, error)
	InsertUser(u models.SignupRequestBody) (bool, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) GetUserByEmail(email string) (models.User, error) {

	var user models.User
	row := ur.db.QueryRow("Select * From users WHERE email = $1", email)

	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role, &user.IsVerified); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("email %v is not linked to any user", email)
		}
		return user, fmt.Errorf("GetUserByEmail %v: %v", email, err)
	}
	return user, nil
}

func (ur *UserRepo) InsertUser(u models.SignupRequestBody) (bool, error) {
	_, err := ur.db.Exec("INSERT INTO users (username, email, password, role, is_verified) VALUES ($1, $2, $3, $4, $5)", u.UserName, u.Email, u.Password, u.Role, u.IsVerified)
	if err != nil {
		return false, fmt.Errorf("InsertUser: %v", err)
	}
	return true, nil
}
