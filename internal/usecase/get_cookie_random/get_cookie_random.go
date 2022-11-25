package getcookierandom

import (
	"time"

	"github.com/daviamorim29/cookieapi/internal/entity"
	"github.com/daviamorim29/cookieapi/internal/gateway"
)

type GetCookieRandomOutputDTO struct {
	ID        string
	Abck      string
	UserAgent string
	Proxy     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetRandomCookie struct {
	cookieGateway gateway.CookieGateway
}

func NewGetRandomCookie(cookieGateway gateway.CookieGateway) *GetRandomCookie {
	return &GetRandomCookie{
		cookieGateway: cookieGateway,
	}
}

func (g *GetRandomCookie) Execute() (*entity.Cookie, error) {
	cookie, err := g.cookieGateway.GetRandom()
	if err != nil {
		return nil, err
	}
	return cookie, nil
}
