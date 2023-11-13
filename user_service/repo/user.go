package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"user_service/models"
)

type IUserRepo interface {
	SetAdminRole(email string) error
	GetByEmail(email string) (models.User, error)
	GetByIDs(userIDs []int64) ([]models.User, error)
	Insert(ctx context.Context, tx *sql.Tx, newUser models.User) (models.User, error)
	Update(user models.User) error
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
	row := ur.db.QueryRow("Select * from users WHERE email = $1", email)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.IsVerified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("email %s is not linked to any user", email)
		}
		return user, fmt.Errorf("GetByEmail %s: %v", email, err)
	}

	return user, nil
}

func (ur *UserRepo) GetByIDs(userIDs []int64) ([]models.User, error) {
	var users []models.User

	placeholders := make([]string, len(userIDs))
	values := make([]interface{}, len(userIDs))
	for i, id := range userIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		values[i] = id
	}
	placeholderString := strings.Join(placeholders, ", ")

	query := fmt.Sprintf("SELECT id, email FROM users WHERE id IN (%s)", placeholderString)

	rows, err := ur.db.Query(query, values...)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (ur *UserRepo) Insert(ctx context.Context, tx *sql.Tx, newUser models.User) (models.User, error) {
	var user models.User
	err := tx.QueryRowContext(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *", newUser.Username, newUser.Email, newUser.Password).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Password, &user.IsVerified)
	if err != nil {
		return user, fmt.Errorf("Insert: %v", err)
	}
	return user, nil
}

func (ur *UserRepo) SetAdminRole(email string) error {

	_, err := ur.db.Exec("UPDATE users SET role = 'admin' where email =$1", email)
	if err != nil {
		return fmt.Errorf("SetAdminRole: %v", err)
	}
	return nil
}

func (ur *UserRepo) Update(user models.User) error {
	_, err := ur.db.Exec("UPDATE users SET username = $1, is_verified = $2 where email = $3", user.Username, user.IsVerified, user.Email)
	if err != nil {
		return fmt.Errorf("UpdateByEmail: %v", err)
	}
	return nil
}

func (ur *UserRepo) ExecTx(ctx context.Context) (*sql.Tx, error) {
	return ur.db.BeginTx(ctx, nil)
}
