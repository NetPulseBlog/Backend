package postgresql

import (
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
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

func (repo UserRepo) FindByEmail(email string) (entity.User, error) {
	fUserSettings := entity.UserSettings{}
	fUser := entity.User{}

	row := repo.db.QueryRow(`
	   SELECT 
			u.*,
			(SELECT COUNT(*) FROM user_subscription WHERE subscriber_id = u.id) AS subscriptions_count,
			(SELECT COUNT(*) FROM user_subscription WHERE subscribed_user_id = u.id) AS subscribers_count,
			us.news_line_default, us.news_line_sort
		FROM 
			"user" u
		INNER JOIN 
			"user_settings" us ON u.id = us.user_id
		WHERE 
			u.email = $1
    `, email)
	if row == nil {
		return fUser, fmt.Errorf("user with email %s not found", email, entity.ErrUserNotFound)
	}

	err := row.Scan(
		&fUser.Id,
		&fUser.EncryptedPassword,
		&fUser.Salt,
		&fUser.CreatedAt,
		&fUser.UpdatedAt,
		&fUser.AccountType,
		&fUser.Role,
		&fUser.Email,
		&fUser.Name,
		&fUser.Description,
		&fUser.AvatarUrl,
		&fUser.CoverUrl,
		&fUser.SubscriptionsCount,
		&fUser.SubscribersCount,
		&fUserSettings.NewsLineDefault,
		&fUserSettings.NewsLineSort,
	)
	if err != nil {
		return fUser, fmt.Errorf("user with email %s not found", email, entity.ErrUserNotFound)
	}

	fUserSettings.UserId = fUser.Id
	fUser.Settings = fUserSettings

	return fUser, nil
}

func (repo UserRepo) CreatePersonal(newUser *entity.User) error {
	const op = "postgresql.UserRepo.CreatePersonal"

	tx, err := repo.db.Begin()
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	newUserStmt, err := tx.Prepare(
		`INSERT INTO "user"(id, encrypted_password, salt, created_at, updated_at, account_type, role, email, name, description, avatar_url, cover_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
	)
	if err != nil {
		return ers.ThrowMessage(op, err)
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
		return ers.ThrowMessage(op, err)
	}

	userSettingsStmt, err := tx.Prepare(
		`INSERT INTO "user_settings"(user_id, news_line_default, news_line_sort) VALUES ($1, $2, $3)`,
	)
	if err != nil {
		tx.Rollback()
		return ers.ThrowMessage(op, err)
	}

	_, err = userSettingsStmt.Exec(
		newUser.Id,
		newUser.Settings.NewsLineDefault,
		newUser.Settings.NewsLineSort,
	)
	if err != nil {
		tx.Rollback()
		return ers.ThrowMessage(op, err)
	}

	err = tx.Commit()
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}
