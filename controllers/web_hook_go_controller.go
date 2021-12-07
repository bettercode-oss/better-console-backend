package controllers

import (
	"better-console-backend/adapters"
	"better-console-backend/middlewares"
	"github.com/labstack/echo"
	"net/http"
)

type WebHookGoController struct {
}

func (controller WebHookGoController) Init(g *echo.Group) {
	g.GET("/ignores", controller.GetIgnoredItems, middlewares.CheckPermission([]string{"*"}))
	g.POST("/ignores", controller.AddIgnoreItem, middlewares.CheckPermission([]string{"*"}))
	g.GET("/template", controller.GetTemplate, middlewares.CheckPermission([]string{"*"}))
	g.PUT("/template", controller.UpdateTemplate, middlewares.CheckPermission([]string{"*"}))
	g.POST("/template/check", controller.CheckTemplateSyntax, middlewares.CheckPermission([]string{"*"}))
	g.POST("/template/reload", controller.ReloadTemplate, middlewares.CheckPermission([]string{"*"}))
}

func (WebHookGoController) GetIgnoredItems(ctx echo.Context) error {
	items, err := adapters.WebHookGoRestAdapter{}.GetIgnoredItems()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, items)
}

func (WebHookGoController) AddIgnoreItem(ctx echo.Context) error {
	var newIgnoredItem map[string]interface{}

	if err := ctx.Bind(&newIgnoredItem); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := adapters.WebHookGoRestAdapter{}.AddIgnoreItem(newIgnoredItem)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (WebHookGoController) GetTemplate(ctx echo.Context) error {
	template, err := adapters.WebHookGoRestAdapter{}.GetTemplate()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, template)
}

func (WebHookGoController) UpdateTemplate(ctx echo.Context) error {
	var template map[string]interface{}

	if err := ctx.Bind(&template); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := adapters.WebHookGoRestAdapter{}.UpdateTemplate(template)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (WebHookGoController) CheckTemplateSyntax(ctx echo.Context) error {
	var template map[string]interface{}

	if err := ctx.Bind(&template); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := adapters.WebHookGoRestAdapter{}.CheckTemplateSyntax(template)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}

func (WebHookGoController) ReloadTemplate(ctx echo.Context) error {
	err := adapters.WebHookGoRestAdapter{}.ReloadTemplate()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
