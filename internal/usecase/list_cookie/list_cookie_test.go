package list_cookie

import (
	"testing"

	"github.com/daviamorim29/cookieapi/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type cookieGatewayMock struct {
	mock.Mock
}

func (m *cookieGatewayMock) Save(cookie *entity.Cookie) error {
	args := m.Called(cookie)
	return args.Error(0)
}

func (m *cookieGatewayMock) GetRandom() (*entity.Cookie, error) {
	args := m.Called()
	return args.Get(0).(*entity.Cookie), args.Error(1)
}

func (m *cookieGatewayMock) Get(id string) (*entity.Cookie, error) {
	args := m.Called()
	return args.Get(0).(*entity.Cookie), args.Error(1)
}

func (m *cookieGatewayMock) Delete(id string) error {
	args := m.Called()
	return args.Error(0)
}

func (m *cookieGatewayMock) List(limit int) ([]*entity.Cookie, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Cookie), args.Error(1)
}

func TestGetListCookie(t *testing.T) {
	cookieGatewayMock := &cookieGatewayMock{}
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	cookieGatewayMock.On("List").Return([]*entity.Cookie{cookie}, nil)
	listCookie := NewListCookie(cookieGatewayMock)
	output, err := listCookie.Execute(&ListCookieInputDTO{
		Limit: 5,
	})
	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, "abck", output[0].Abck)
	assert.Equal(t, "daviamorim29", output[0].UserAgent)
	cookieGatewayMock.AssertNumberOfCalls(t, "List", 1)
}
