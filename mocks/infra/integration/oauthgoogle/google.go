package mock_oauthgoogle

import (
	"context"
	dto "ex_service/src/app/dto/user"

	"github.com/stretchr/testify/mock"
)

type MockOauthGoogleService struct {
	mock.Mock
}

func NewMockOauthGoogleService() *MockOauthGoogleService {
	return &MockOauthGoogleService{}
}

func (m *MockOauthGoogleService) GetUrl(state string) string {
	panic("TODO: Implement")
}

func (m *MockOauthGoogleService) ExchangeCode(ctx context.Context, code string) (token string, err error) {
	panic("TODO: Implement")
}

func (m *MockOauthGoogleService) GetUser(ctx context.Context, token string) (user *dto.UserInfoGoogleDTO, err error) {
	panic("TODO: Implement")
}
