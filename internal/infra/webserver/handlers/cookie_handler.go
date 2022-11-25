package handlers

import (
	"strconv"

	"github.com/daviamorim29/cookieapi/internal/infra/database"
	createcookie "github.com/daviamorim29/cookieapi/internal/usecase/create_cookie"
	deletecookie "github.com/daviamorim29/cookieapi/internal/usecase/delete_cookie"
	get_cookie "github.com/daviamorim29/cookieapi/internal/usecase/get_cookie"
	get_cookie_random "github.com/daviamorim29/cookieapi/internal/usecase/get_cookie_random"
	"github.com/daviamorim29/cookieapi/internal/usecase/list_cookie"
	"github.com/gin-gonic/gin"
)

type CookieHandler struct {
	CookieDb *database.CookieDB
}

type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewCookieHandler(router *gin.Engine, cookieDb *database.CookieDB) {
	cookieHandler := &CookieHandler{
		CookieDb: cookieDb,
	}
	router.POST("/cookie", cookieHandler.CreateCookie)
	router.GET("/cookie/:id", cookieHandler.GetCookie)
	router.GET("/cookie", cookieHandler.GetRandomCookie)
	router.DELETE("/cookie/:id", cookieHandler.DeleteCookie)
	router.GET("/cookies", cookieHandler.List)
}

func (c *CookieHandler) CreateCookie(ctx *gin.Context) {
	var cookie createcookie.CreateCookieInputDTO
	if err := ctx.BindJSON(&cookie); err != nil {
		ctx.Writer.WriteHeader(400)
		ctx.JSON(400, Error{
			Success: false,
			Message: "Invalid JSON",
		})
		return
	}

	// save cookie
	createCookie := createcookie.NewCreateCookie(c.CookieDb)
	output, err := createCookie.Execute(&cookie)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		ctx.JSON(500, Error{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, output)
}

func (c *CookieHandler) DeleteCookie(ctx *gin.Context) {
	// delete cookie
	id := ctx.Params.ByName("id")
	deleteCookie := deletecookie.NewDeleteCookie(c.CookieDb)
	output, err := deleteCookie.Execute(&deletecookie.DeleteCookieInputDTO{
		ID: id,
	})
	if err != nil {
		ctx.Writer.WriteHeader(500)
		ctx.JSON(500, Error{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, output)
}

func (c *CookieHandler) GetCookie(ctx *gin.Context) {
	// get cookie
	id := ctx.Params.ByName("id")
	cookie := get_cookie.NewGetCookie(c.CookieDb)
	output, err := cookie.Execute(&get_cookie.GetCookieInputDTO{
		ID: id,
	})
	if err != nil {
		ctx.JSON(404, Error{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, output)
}

func (c *CookieHandler) GetRandomCookie(ctx *gin.Context) {
	// get random cookie
	cookie := get_cookie_random.NewGetRandomCookie(c.CookieDb)
	output, err := cookie.Execute()
	if err != nil {
		ctx.JSON(404, Error{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, output)
}

func (c *CookieHandler) List(ctx *gin.Context) {
	// list cookies
	// transform to int
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 1
	}
	cookie := list_cookie.NewListCookie(c.CookieDb)
	output, err := cookie.Execute(&list_cookie.ListCookieInputDTO{
		Limit: limit,
	})
	if err != nil {
		ctx.JSON(400, Error{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, output)
}
