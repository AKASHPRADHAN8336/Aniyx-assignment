package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AKASHPRADHAN8336/aniyxProject/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, name, dob string) (int32, error)
	GetByID(ctx context.Context, id int32) (sqlc.User, error)
	List(ctx context.Context) ([]sqlc.User, error)
	Update(ctx context.Context, id int32, name, dob string) error
	Delete(ctx context.Context, id int32) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, name, dob string) (int32, error) {
	fmt.Printf("DEBUG: Creating user: %s, %s\n", name, dob)

	result, err := r.db.ExecContext(ctx,
		"INSERT INTO users (name, dob) VALUES (?, ?)",
		name, dob,
	)
	if err != nil {
		fmt.Printf("DEBUG: Create error: %v\n", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("DEBUG: LastInsertId error: %v\n", err)
		return 0, err
	}

	fmt.Printf("DEBUG: Created user ID: %d\n", id)
	return int32(id), nil
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (sqlc.User, error) {
	fmt.Printf("DEBUG: Getting user by ID: %d\n", id)

	row := r.db.QueryRowContext(ctx, "SELECT id, name, dob FROM users WHERE id = ?", id)

	var user sqlc.User
	err := row.Scan(&user.ID, &user.Name, &user.Dob)
	if err != nil {
		fmt.Printf("DEBUG: GetByID error: %v\n", err)
		return user, err
	}

	fmt.Printf("DEBUG: Found user: %+v\n", user)
	return user, nil
}

func (r *userRepository) List(ctx context.Context) ([]sqlc.User, error) {
	fmt.Println("DEBUG: Listing all users")

	rows, err := r.db.QueryContext(ctx, "SELECT id, name, dob FROM users ORDER BY id")
	if err != nil {
		fmt.Printf("DEBUG: List query error: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var users []sqlc.User
	for rows.Next() {
		var user sqlc.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Dob); err != nil {
			fmt.Printf("DEBUG: Row scan error: %v\n", err)
			return nil, err
		}
		users = append(users, user)
		fmt.Printf("DEBUG: Found user: ID=%d, Name=%s, Dob=%s\n", user.ID, user.Name, user.Dob)
	}

	fmt.Printf("DEBUG: Total users found: %d\n", len(users))
	return users, nil
}

func (r *userRepository) Update(ctx context.Context, id int32, name, dob string) error {
	fmt.Printf("DEBUG: Updating user %d: %s, %s\n", id, name, dob)

	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET name = ?, dob = ? WHERE id = ?",
		name, dob, id,
	)
	if err != nil {
		fmt.Printf("DEBUG: Update error: %v\n", err)
	}

	return err
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	fmt.Printf("DEBUG: Deleting user: %d\n", id)

	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	if err != nil {
		fmt.Printf("DEBUG: Delete error: %v\n", err)
	}

	return err
}
