package membership_test

import (
	"github.com/stretchr/testify/assert"
	"golang_web_programming/app/membership"
	"testing"
)

func TestService_ReadAllMembership(t *testing.T) {
	data := map[string]membership.Membership{}
	repo := membership.NewRepository(data)
	service := membership.NewService(*repo)
	t.Run("limit이 없는 경우", func(t *testing.T) {
		req := membership.ReadAllRequest{
			Limit:  "",
			Offset: "2",
		}
		res := service.ReadAllMembership(req)

		assert.Empty(t, res)
	})
}
