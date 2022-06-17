package membership

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		app.Create(req)

		alreadyExistedNameReq := CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		}

		_, err := app.Create(alreadyExistedNameReq)
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("already existed name"), err)
		}
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{
			UserName:       "",
			MembershipType: "naver",
		}

		_, err := app.Create(req)
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("there is no user name"), err)
		}
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{
			UserName:       "jenny",
			MembershipType: "",
		}

		_, err := app.Create(req)
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("there is no membership type"), err)
		}
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{
			UserName:       "jenny",
			MembershipType: "google",
		}

		_, err := app.Create(req)
		if assert.Error(t, err) {
			assert.Equal(t, errors.New("not supported membership"), err)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		res, err := app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "payco",
		})

		req := UpdateRequest{
			ID:             res.ID,
			UserName:       "jenny",
			MembershipType: "naver",
		}

		_, err = app.Update(req)
		membershipFromData, _ := app.repository.data[res.ID]

		assert.Equal(t, "naver", membershipFromData.MembershipType)
		assert.Nil(t, err)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "payco",
		})

		_, err := app.Update(UpdateRequest{
			ID:             "uuid",
			UserName:       "jenny",
			MembershipType: "payco",
		})

		if assert.Error(t, err) {
			assert.Equal(t, errors.New("already existed name"), err)
		}
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "",
			UserName:       "jenny",
			MembershipType: "payco",
		})

		if assert.Error(t, err) {
			assert.Equal(t, errors.New("there is no id"), err)
		}
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "uuid",
			UserName:       "",
			MembershipType: "payco",
		})

		if assert.Error(t, err) {
			assert.Equal(t, errors.New("there is no user name"), err)
		}
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "uuid",
			UserName:       "jenny",
			MembershipType: "",
		})

		if assert.Error(t, err) {
			assert.Equal(t, errors.New("there is no membership type"), err)
		}
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		_, err := app.Update(UpdateRequest{
			ID:             "uuid",
			UserName:       "jenny",
			MembershipType: "google",
		})

		if assert.Error(t, err) {
			assert.Equal(t, errors.New("not supported membership"), err)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		res, err := app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		})

		err = app.Delete(res.ID)
		assert.Nil(t, err)
		assert.Equal(t, Membership{}, app.repository.data[res.ID])
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		err := app.Delete("")

		if assert.Error(t, err) {
			assert.Equal(t, errors.New("there is no id"), err)
		}
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		})

		err := app.Delete("uuid")

		if assert.Error(t, err) {
			assert.Equal(t, errors.New("there is no id"), err)
		}
	})
}

func TestRead(t *testing.T) {
	t.Run("멤버십 id로 멤버십을 가져온다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		response, err := app.Create(CreateRequest{
			UserName:       "jenny",
			MembershipType: "naver",
		})

		readResponse, err := app.Read(response.ID)

		assert.Nil(t, err)
		assert.NotEmpty(t, readResponse)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {

	})
}
