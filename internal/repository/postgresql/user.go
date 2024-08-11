package postgresql

import (
	"app/pkg/domain/entity"
	"database/sql"
	"fmt"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (uR UserRepo) CreatePersonal(newUser *entity.User) error {
	const op = "postgresql.UserRepo.CreatePersonal"

	tx, err := uR.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	newUserStmt, err := tx.Prepare(
		`INSERT INTO "user"(id, encrypted_password, salt, created_at, updated_at, account_type, role, email, name, description, avatar_url, cover_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = newUserStmt.Exec(
		newUser.Id,
		newUser.EncryptedPassword,
		newUser.Salt,
		newUser.CreatedAt,
		newUser.UpdatedAt,
		newUser.AccountType,
		newUser.Role,
		newUser.Email,
		newUser.Name,
		newUser.Description,
		newUser.AvatarUrl,
		newUser.CoverUrl,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	userSettingsStmt, err := tx.Prepare(
		`INSERT INTO "user_settings"(user_id, news_line_default, news_line_sort) VALUES ($1, $2, $3)`,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = userSettingsStmt.Exec(
		newUser.Id,
		newUser.Settings.NewsLineDefault,
		newUser.Settings.NewsLineSort,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
