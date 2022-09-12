package user_test

import (
	"dealljobs/domain/request"
	"dealljobs/domain/user"
	"dealljobs/domain/user/impl"
	"dealljobs/domain/user/mocks"
	"testing"
)

type itemServiceTest struct {
	MockRepo    *mocks.Repository
	MockService *mocks.Service
	Service     user.Service
}

var svcTest itemServiceTest

func init() {
	mockRepo := new(mocks.Repository)
	mockService := new(mocks.Service)

	svcTest = itemServiceTest{
		MockRepo:    mockRepo,
		MockService: mockService,
		Service: impl.NewService(
			mockRepo,
		),
	}
}

func TestCreateItemIfNotAny(t *testing.T) {
	req := request.CreateUserRequest{}

	it := user.User{
		Name: "ITEM-123",
	}

	t.Run("positive result", func(t *testing.T) {
		svcTest.MockRepo.On("GetUserByUsername", it).Return(true, nil).Once()
		svcTest.MockRepo.On("Persist", it).Once()
		svcTest.Service.CreateUserIfNotAny(req)
	})

}
