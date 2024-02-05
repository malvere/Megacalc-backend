package sqlc__test

import (
	"context"
	"testing"
	"tg-backend/internal/db/sqlc"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *sqlc.User {
	c := createRandomCode(t)
	require.NotEmpty(t, c)

	arg := sqlc.CreateUserParams{
		UserID:       uuid.New(),
		TelegramID:   generateRandomString(8),
		InviteCodeID: c.CodeID,
	}
	u, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	return &u
}

func createRandomCode(t *testing.T) *sqlc.InviteCode {
	arg := sqlc.CreateCodeParams{
		CodeID: uuid.New(),
		Code:   generateRandomString(4),
		Active: true,
	}
	c, err := testQueries.CreateCode(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	return &c
}
func TestCreateUser(t *testing.T) {
	u1 := createRandomUser(t)
	require.NotEmpty(t, u1)
}

func TestAll(t *testing.T) {
	ctx := context.Background()
	u1 := createRandomUser(t)
	require.NotEmpty(t, u1)

	uf1, err := testQueries.FindByTGID(ctx, u1.TelegramID)
	require.NoError(t, err)
	require.NotEmpty(t, uf1)

	c1 := createRandomCode(t)
	argc := sqlc.UpdateInviteByTGIDParams{
		TelegramID:   uf1.TelegramID,
		InviteCodeID: c1.CodeID,
	}
	err = testQueries.UpdateInviteByTGID(ctx, argc)
	require.NoError(t, err)

	err = testQueries.UpdateCode(ctx, sqlc.UpdateCodeParams{
		Code:   c1.Code,
		Active: false,
	})
	require.NoError(t, err)

	err = testQueries.DeleteByTGID(ctx, uf1.TelegramID)
	require.NoError(t, err)

	err = testQueries.DeleteCode(ctx, c1.Code)
	require.NoError(t, err)

}
