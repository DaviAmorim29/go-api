package getcookie

import (
	"errors"
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

func TestGetCookieExists(t *testing.T) {
	cookieGatewayMock := &cookieGatewayMock{}
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	cookieGatewayMock.On("Get").Return(cookie, nil)
	getCookie := NewGetCookie(cookieGatewayMock)
	output, err := getCookie.Execute(&GetCookieInputDTO{
		ID: cookie.ID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "abck", output.Abck)
	assert.Equal(t, "daviamorim29", output.UserAgent)
	assert.Equal(t, "proxyUrl", output.Proxy)
	cookieGatewayMock.AssertNumberOfCalls(t, "Get", 1)
}

func (m *cookieGatewayMock) List(limit int) ([]*entity.Cookie, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Cookie), args.Error(1)
}

func TestGetCookieDontExists(t *testing.T) {
	cookieGatewayMock := &cookieGatewayMock{}
	cookieGatewayMock.On("Get").Return(&entity.Cookie{}, errors.New("Cookie not found"))
	getCookie := NewGetCookie(cookieGatewayMock)
	output, err := getCookie.Execute(&GetCookieInputDTO{
		ID: "123",
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Cookie not found")
	assert.Nil(t, output)
	cookieGatewayMock.AssertNumberOfCalls(t, "Get", 1)
}
