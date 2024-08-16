package postgresql

import (
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (repo AuthRepo) GetById(authId uuid.UUID) (*entity.UserAuth, error) {
	const op = "postgresql.AuthRepo.GetById"

	fUserAuth := entity.UserAuth{
		Token: entity.AuthToken{},
	}

	row := repo.db.QueryRow(`SELECT * FROM user_auth WHERE id = $1`, authId)
	err := row.Scan(
		&fUserAuth.Id,
		&fUserAuth.UserId,
		&fUserAuth.Token.Refresh,
		&fUserAuth.Token.Access,
		&fUserAuth.DeviceName,
		&fUserAuth.Token.AccessExpiresAt,
		&fUserAuth.Token.RefreshExpiresAt,
		&fUserAuth.CreatedAt,
		&fUserAuth.UpdatedAt,
	)
	if err != nil {
		return &fUserAuth, ers.ThrowMessage(op, fmt.Errorf("auth row not found"))
	}

	return &fUserAuth, nil
}

func (repo AuthRepo) Update(uAuth *entity.UserAuth) error {
	const op = "postgresql.AuthRepo.Update"

	newUserStmt, err := repo.db.Prepare(
		`UPDATE "user_auth" SET user_id = $2, refresh_token = $3, access_token = $4, device_name = $5, access_expires_at = $6, refresh_expires_at = $7, created_at = $8, updated_at = $9 WHERE id = $1`,
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

func (repo AuthRepo) DeleteItem(authId uuid.UUID) error {
	const op = "postgresql.AuthRepo.DeleteItem"

	newUserStmt, err := repo.db.Prepare(
		`DELETE FROM user_auth WHERE id = $1`,
	)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	_, err = newUserStmt.Exec(authId)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}

func (repo AuthRepo) RemoveExistsForDevice(userId uuid.UUID, deviceName string) error {
	const op = "postgresql.AuthRepo.RemoveExistsForDevice"

	newUserStmt, err := repo.db.Prepare(
		`DELETE FROM user_auth WHERE user_id = $1 AND device_name = $2`,
	)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	_, err = newUserStmt.Exec(userId, deviceName)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
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
