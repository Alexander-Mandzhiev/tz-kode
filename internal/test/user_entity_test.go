package test

import (
	"testing"
	"tz-kode/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *entity.User
		IsValid bool
	}{
		{
			name: "valid",
			user: func() *entity.User {
				return TestUser(t)
			},
			IsValid: true,
		},
		{
			name: "empty email",
			user: func() *entity.User {
				u := TestUser(t)
				u.Email = ""
				return u
			},
			IsValid: false,
		},
		{
			name: "invalid email",
			user: func() *entity.User {
				u := TestUser(t)
				u.Email = "213fs"
				return u
			},
			IsValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.IsValid {
				assert.NoError(t, tc.user().Valiedate())
			} else {
				assert.Error(t, tc.user().Valiedate())
			}
		})
	}
}
