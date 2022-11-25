package deletecookie

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

func TestDeleteCookie(t *testing.T) {
	cookieGatewayMock := &cookieGatewayMock{}
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	cookieGatewayMock.On("Delete").Return(nil)
	getCookie := NewDeleteCookie(cookieGatewayMock)
	output, err := getCookie.Execute(&DeleteCookieInputDTO{
		ID: cookie.ID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, true, output.Success)
	cookieGatewayMock.AssertNumberOfCalls(t, "Delete", 1)
}

func (m *cookieGatewayMock) List(limit int) ([]*entity.Cookie, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Cookie), args.Error(1)
}

func TestDeleteCookieDontExists(t *testing.T) {
	cookieGatewayMock := &cookieGatewayMock{}
	cookieGatewayMock.On("Delete").Return(errors.New("Cookie not found"))
	getCookie := NewDeleteCookie(cookieGatewayMock)
	output, err := getCookie.Execute(&DeleteCookieInputDTO{
		ID: "123",
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Cookie not found")
	assert.Nil(t, output)
}
