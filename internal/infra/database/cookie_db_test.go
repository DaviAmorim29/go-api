package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/daviamorim29/cookieapi/internal/entity"
	"github.com/stretchr/testify/suite"
)

type CookieDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	cookieDB *CookieDB
}

func (s *CookieDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	// Create table
	_, err = db.Exec("CREATE TABLE cookies (id VARCHAR(255), abck TEXT, user_agent TEXT, proxy TEXT, created_at date, updated_at date)")
	s.Nil(err)
	s.cookieDB = NewCookieDB(s.db)
}

func (s *CookieDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	// drop table
	_, err := s.db.Exec("DROP TABLE cookies")
	s.Nil(err)
}

func TestCookieDBTestSuite(t *testing.T) {
	suite.Run(t, new(CookieDBTestSuite))
}

func (s *CookieDBTestSuite) TestSave() {
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	err := s.cookieDB.Save(cookie)
	s.Nil(err)
}

func (s *CookieDBTestSuite) TestGetRandom() {
	// save cookie first
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	s.cookieDB.Save(cookie)
	// get random cookie
	cookie, err := s.cookieDB.GetRandom()
	s.Nil(err)
	s.NotNil(cookie)
	s.Equal("abck", cookie.Abck)
	s.Equal("daviamorim29", cookie.UserAgent)
	s.Equal("proxyUrl", cookie.Proxy)
}

func (s *CookieDBTestSuite) TestGet() {
	// save cookie first
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	s.cookieDB.Save(cookie)
	// get cookie
	cookie, err := s.cookieDB.Get(cookie.ID)
	s.Nil(err)
	s.NotNil(cookie)
	s.Equal("abck", cookie.Abck)
	s.Equal("daviamorim29", cookie.UserAgent)
	s.Equal("proxyUrl", cookie.Proxy)
}

func (s *CookieDBTestSuite) TestDeleteCookie() {
	// save cookie first
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	s.cookieDB.Save(cookie)
	// delete cookie
	err := s.cookieDB.Delete(cookie.ID)
	s.Nil(err)
	// get cookie
	cookie, err = s.cookieDB.Get(cookie.ID)
	s.NotNil(err)
	s.Nil(cookie)
}

func (s *CookieDBTestSuite) TestDeleteCookieDontExists() {
	// delete cookie
	err := s.cookieDB.Delete("id")
	s.NotNil(err)
	s.Equal("cookie not found", err.Error())
}

func (s *CookieDBTestSuite) TestList() {
	// save cookie first
	cookie, _ := entity.NewCookie("abck", "daviamorim29", "proxyUrl")
	s.cookieDB.Save(cookie)
	// list cookies
	cookies, err := s.cookieDB.List(1)
	s.Nil(err)
	s.NotNil(cookies)
	s.Equal(1, len(cookies))
	s.Equal("abck", cookies[0].Abck)
	s.Equal("daviamorim29", cookies[0].UserAgent)
	s.Equal("proxyUrl", cookies[0].Proxy)
}
