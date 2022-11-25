package database

import (
	"database/sql"
	"errors"

	"github.com/daviamorim29/cookieapi/internal/entity"
)

type CookieDB struct {
	DB *sql.DB
}

func NewCookieDB(db *sql.DB) *CookieDB {
	return &CookieDB{DB: db}
}

func (c *CookieDB) Save(cookie *entity.Cookie) error {
	smtp, err := c.DB.Prepare("INSERT INTO cookies (id, abck, user_agent, proxy, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = smtp.Exec(cookie.ID, cookie.Abck, cookie.UserAgent, cookie.Proxy, cookie.CreatedAt, cookie.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (c *CookieDB) GetRandom() (*entity.Cookie, error) {
	cookie := &entity.Cookie{}
	stmt, err := c.DB.Prepare("SELECT id, abck, user_agent, proxy, created_at, updated_at FROM cookies LIMIT 1")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow().Scan(&cookie.ID, &cookie.Abck, &cookie.UserAgent, &cookie.Proxy, &cookie.CreatedAt, &cookie.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (c *CookieDB) Get(id string) (*entity.Cookie, error) {
	cookie := &entity.Cookie{}
	stmt, err := c.DB.Prepare("SELECT id, abck, user_agent, proxy, created_at, updated_at FROM cookies WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&cookie.ID, &cookie.Abck, &cookie.UserAgent, &cookie.Proxy, &cookie.CreatedAt, &cookie.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (c *CookieDB) Delete(id string) error {
	stmt, err := c.DB.Prepare("DELETE FROM cookies WHERE id = ?")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	value, err := result.RowsAffected()
	if err != nil {
		return errors.New("error to delete cookie")
	}
	if value == 0 {
		return errors.New("cookie not found")
	}

	return nil
}

func (c *CookieDB) List(limit int) ([]*entity.Cookie, error) {
	cookies := []*entity.Cookie{}
	stmt, err := c.DB.Prepare("SELECT id, abck, user_agent, proxy, created_at, updated_at FROM cookies LIMIT ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cookie := &entity.Cookie{}
		err = rows.Scan(&cookie.ID, &cookie.Abck, &cookie.UserAgent, &cookie.Proxy, &cookie.CreatedAt, &cookie.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cookies = append(cookies, cookie)
	}
	return cookies, nil
}
