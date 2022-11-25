package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCookie(t *testing.T) {
	cookie, err := NewCookie("abck", "daviamorim29", "proxyUrl")
	assert.NotNil(t, cookie)
	assert.Nil(t, err)
	assert.Equal(t, "abck", cookie.Abck)
	assert.Equal(t, "daviamorim29", cookie.UserAgent)
	assert.Equal(t, "proxyUrl", cookie.Proxy)
}

func TestCookieWithInvalids(t *testing.T) {
	cookie, err := NewCookie("", "", "")
	assert.Nil(t, cookie)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid cookie", err.Error())
}
