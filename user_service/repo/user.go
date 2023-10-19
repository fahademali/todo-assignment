package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"user_service/models"
)

type IUserRepo interface {
	GetUserByEmail(email string) (models.User, error)
	InsertUser(u models.SignupRequest) error
	VerifyUser(email string) error
	SetAdminRole(email string) error
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
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.IsVerified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("email %s is not linked to any user", email)
		}
		return user, fmt.Errorf("GetUserByEmail %s: %v", email, err)
	}

	return user, nil
}

func (ur *UserRepo) InsertUser(u models.SignupRequest) error {
	user, err := ur.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *", u.UserName, u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("InsertUser: %v", err)
	}
	fmt.Println("INSERTED USER IMMEDIATELY BACK.................")
	fmt.Println(user)
	return nil
}

func (ur *UserRepo) VerifyUser(email string) error {
	_, err := ur.db.Exec("UPDATE users SET is_verified = true where email = $1", email)
	if err != nil {
		return fmt.Errorf("VerifyUser: %v", err)
	}
	return nil
}

func (ur *UserRepo) SetAdminRole(email string) error {

	_, err := ur.db.Exec("UPDATE users SET role = 'admin' where email =$1", email)
	if err != nil {
		return fmt.Errorf("SetAdminRole: %v", err)
	}
	return nil
}
