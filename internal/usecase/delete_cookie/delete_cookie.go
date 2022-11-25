package deletecookie

import (
	"github.com/daviamorim29/cookieapi/internal/gateway"
)

type DeleteCookieInputDTO struct {
	ID string
}

type DeleteCookieOutputDTO struct {
	Success bool
	Message string
}

type DeleteCookie struct {
	cookieGateway gateway.CookieGateway
}

func NewDeleteCookie(cookieGateway gateway.CookieGateway) *DeleteCookie {
	return &DeleteCookie{
		cookieGateway: cookieGateway,
	}
}

func (g *DeleteCookie) Execute(input *DeleteCookieInputDTO) (*DeleteCookieOutputDTO, error) {
	err := g.cookieGateway.Delete(input.ID)
	if err != nil {
		return nil, err
	}
	return &DeleteCookieOutputDTO{
		Success: true,
		Message: "Cookie deleted successfully",
	}, nil
}
