package getcookierandom

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

func TestGetCookieRandom(t *testing.T) {
	cookieGatewayMock := &cookieGatewayMock{}
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	cookieGatewayMock.On("GetRandom").Return(cookie)
	getCookieRandom := NewGetRandomCookie(cookieGatewayMock)
	output, err := getCookieRandom.Execute()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	cookieGatewayMock.AssertNumberOfCalls(t, "GetRandom", 1)
	assert.NotNil(t, output)
	assert.Equal(t, "abck", output.Abck)
	assert.Equal(t, "daviamorim29", output.UserAgent)
	assert.Equal(t, "proxyUrl", output.Proxy)
	assert.Nil(t, err)
}
