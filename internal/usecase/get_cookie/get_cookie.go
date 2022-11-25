package getcookie

import (
	"github.com/daviamorim29/cookieapi/internal/entity"
	"github.com/daviamorim29/cookieapi/internal/gateway"
)

type GetCookieInputDTO struct {
	ID string
}

// type GetCookieOutputDTO struct {
// 	ID        string
// 	Abck      string
// 	UserAgent string
// 	Proxy     string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type GetCookie struct {
	cookieGateway gateway.CookieGateway
}

func NewGetCookie(cookieGateway gateway.CookieGateway) *GetCookie {
	return &GetCookie{
		cookieGateway: cookieGateway,
	}
}

func (g *GetCookie) Execute(input *GetCookieInputDTO) (*entity.Cookie, error) {
	cookie, err := g.cookieGateway.Get(input.ID)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}
