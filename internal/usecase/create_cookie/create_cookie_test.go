package createcookie

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

func TestCreateCookie_Execute(t *testing.T) {
	cookieGatewayMock := &cookieGatewayMock{}
	cookieGatewayMock.On("Save", mock.Anything).Return(nil)
	createCookie := NewCreateCookie(cookieGatewayMock)
	input := &CreateCookieInputDTO{
		Abck:      "abck",
		UserAgent: "daviamorim29",
		Proxy:     "proxyUrl",
	}
	output, err := createCookie.Execute(input)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	cookieGatewayMock.AssertExpectations(t)
	cookieGatewayMock.AssertNumberOfCalls(t, "Save", 1)
	assert.Equal(t, true, output.Success)
	assert.Equal(t, "Cookie created successfully", output.Message)
}
