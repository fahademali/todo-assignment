package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"user_service/log"
	"user_service/models"
)

type IUserRepo interface {
	SetAdminRole(email string) error
	GetByEmail(email string) (models.User, error)
	Insert(u models.SignupRequest) error
	InsertTx(ctx context.Context, tx *sql.Tx, u models.SignupRequest) error
	UpdateByEmail(email string, user models.User) error
	ExecTx(ctx context.Context) (*sql.Tx, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) GetByEmail(email string) (models.User, error) {

	var user models.User
	row := ur.db.QueryRow("Select * From users WHERE email = $1", email)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.IsVerified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("email %s is not linked to any user", email)
		}
		return user, fmt.Errorf("GetByEmail %s: %v", email, err)
	}

	return user, nil
}

func (ur *UserRepo) Insert(u models.SignupRequest) error {
	_, err := ur.db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", u.UserName, u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("Insert: %v", err)
	}
	fmt.Println("INSERTED USER IMMEDIATELY BACK.................")
	return nil
}

func (ur *UserRepo) InsertTx(ctx context.Context, tx *sql.Tx, u models.SignupRequest) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", u.UserName, u.Email, u.Password)
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

func (ur *UserRepo) UpdateByEmail(email string, user models.User) error {
	log.GetLog().Info(user.Username)
	log.GetLog().Info(user.Role)
	log.GetLog().Info(user.IsVerified)
	log.GetLog().Info(user.Email)
	_, err := ur.db.Exec("UPDATE users SET username = $1, role = $2, is_verified = $3 where email = $4", user.Username, user.Role, user.IsVerified, user.Email)
	if err != nil {
		return fmt.Errorf("UpdateByEmail: %v", err)
	}
	return nil
}

func (ur *UserRepo) ExecTx(ctx context.Context) (*sql.Tx, error) {
	return ur.db.BeginTx(ctx, nil)
}
