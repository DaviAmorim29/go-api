package gateway

import "github.com/daviamorim29/cookieapi/internal/entity"

type CookieGateway interface {
	Save(cookie *entity.Cookie) error
	GetRandom() (*entity.Cookie, error)
	Get(id string) (*entity.Cookie, error)
	Delete(id string) error
	List(limit int) ([]*entity.Cookie, error)
}
