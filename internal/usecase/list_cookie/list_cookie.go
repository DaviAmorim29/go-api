package list_cookie

import (
	"github.com/daviamorim29/cookieapi/internal/entity"
	"github.com/daviamorim29/cookieapi/internal/gateway"
)

type ListCookieInputDTO struct {
	Limit int
}

type ListCookie struct {
	cookieGateway gateway.CookieGateway
}

func NewListCookie(cookieGateway gateway.CookieGateway) *ListCookie {
	return &ListCookie{
		cookieGateway: cookieGateway,
	}
}

func (l *ListCookie) Execute(input *ListCookieInputDTO) ([]*entity.Cookie, error) {
	cookies, err := l.cookieGateway.List(input.Limit)
	if err != nil {
		return nil, err
	}
	return cookies, nil
}
