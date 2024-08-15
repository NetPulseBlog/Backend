package postgresql

import (
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"database/sql"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (repo AuthRepo) Create(uAuth entity.UserAuth) error {
	const op = "postgresql.AuthRepo.Create"

	newUserStmt, err := repo.db.Prepare(
		`INSERT INTO "user_auth" (id, user_id, refresh_token, access_token, device_name, access_expires_at, refresh_expires_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
	)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	_, err = newUserStmt.Exec(
		uAuth.Id,
		uAuth.UserId,
		uAuth.Token.Refresh,
		uAuth.Token.Access,
		uAuth.DeviceName,
		uAuth.Token.AccessExpiresAt.UTC(),
		uAuth.Token.RefreshExpiresAt.UTC(),
		uAuth.CreatedAt.UTC(),
		uAuth.UpdatedAt.UTC(),
	)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}
