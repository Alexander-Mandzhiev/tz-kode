package test

import (
	"testing"
	"time"
	"tz-kode/internal/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryUser_Create(t *testing.T) {
	repo, teardown := TestRepository(t, databaseURL)
	defer teardown("users")

	id := uuid.New().String()
	err := repo.Create(&entity.User{
		ID:        id,
		Username:  "",
		Email:     "user1",
		Password:  "124",
		CreatedAt: time.Now(),
	})

	assert.Error(t, err)

	u := TestUser(t)
	assert.NoError(t, repo.User.Create(u))
	assert.NotNil(t, u)

}

func TestRepositoryUser_FingByEmain(t *testing.T) {
	repo, teardown := TestRepository(t, databaseURL)
	defer teardown("users")

	email := "user@mail.com"

	_, err := repo.FindByEmain(email)
	assert.Error(t, err)

	u := TestUser(t)
	repo.User.Create(u)

	tu, err := repo.FindByEmain(u.Email)

	assert.NoError(t, err)
	assert.NotNil(t, tu)
}

func TestRepositoryUser_FingById(t *testing.T) {
	repo, teardown := TestRepository(t, databaseURL)
	defer teardown("users")

	id := uuid.New().String()
	_, err := repo.FindById(id)
	assert.Error(t, err)

	u := TestUser(t)
	repo.User.Create(u)

	tu, err := repo.FindById(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, tu)
}
