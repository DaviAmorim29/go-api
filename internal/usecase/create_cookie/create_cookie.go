package createcookie

import (
	"github.com/daviamorim29/cookieapi/internal/entity"
	"github.com/daviamorim29/cookieapi/internal/gateway"
)

type CreateCookieInputDTO struct {
	Abck      string `json:"abck"`
	UserAgent string `json:"user_agent"`
	Proxy     string `json:"proxy"`
}

type CreateCookieOutputDTO struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CreateCookie struct {
	cookieGateway gateway.CookieGateway
}

func NewCreateCookie(cookieGateway gateway.CookieGateway) *CreateCookie {
	return &CreateCookie{
		cookieGateway: cookieGateway,
	}
}

func (c *CreateCookie) Execute(input *CreateCookieInputDTO) (*CreateCookieOutputDTO, error) {
	cookie, err := entity.NewCookie(input.Abck, input.UserAgent, input.Proxy)
	if err != nil {
		return nil, err
	}
	err = c.cookieGateway.Save(cookie)
	if err != nil {
		return nil, err
	}
	return &CreateCookieOutputDTO{
		Success: true,
		Message: "Cookie created successfully",
	}, nil
}
