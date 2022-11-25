package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Cookie struct {
	ID        string    `json:"id"`
	Abck      string    `json:"abck"`
	UserAgent string    `json:"user_agent"`
	Proxy     string    `json:"proxy"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCookie(abck string, userAgent string, proxy string) (*Cookie, error) {
	cookie := &Cookie{
		ID:        uuid.New().String(),
		Abck:      abck,
		UserAgent: userAgent,
		Proxy:     proxy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := cookie.Validate()
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (c *Cookie) Validate() error {
	if (c.Abck == "") || (c.UserAgent == "") || (c.Proxy == "") || (c.ID == "") {
		return errors.New("Invalid cookie")
	}
	return nil
}
