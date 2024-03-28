package main

import (
	"github.com/Saakhr/go-affichage-de-notes/templates"
	"github.com/Saakhr/go-affichage-de-notes/templates/pages"
	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type User struct {
	id       int
	username string
	session  string
	role     string
}
type authHandler func(echo.Context, User) error

func (api *ApiConf) middlewareAuth(handler authHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		userid := sess.Values["userid"]
		query := `
		SELECT UserID, Username, Session, Role	FROM Users WHERE UserID = ?;
	`
		var user User
		_ = api.db.QueryRow(query, userid).Scan(&user.id, &user.username, &user.session, &user.role)
		if sess.Values["session"] != user.session {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

			metaTags := pages.MetaTags(
				"Affichage de Notes", // define meta description
				"Affichage de Notes", // define meta description
			)
			var indexTemplate templ.Component
			indexTemplate = templates.Layout(
				"Affichage de Notes",     // define title text
				metaTags,                 // define meta tags
				pages.BodyContent(false), // define body content
			)
			return htmx.NewResponse().Retarget("body").RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)

		}
		return handler(c, user)

	}

}

func (api *ApiConf) middlewareAdminPer(handler authHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		userid := sess.Values["userid"]
		query := `
		SELECT UserID, Username, Session, Role	FROM Users WHERE UserID = ?;
	`
		var user User
		_ = api.db.QueryRow(query, userid).Scan(&user.id, &user.username, &user.session, &user.role)
		if sess.Values["session"] != user.session {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

			metaTags := pages.MetaTags(
				"Affichage de Notes", // define meta description
				"Affichage de Notes", // define meta description
			)
			var indexTemplate templ.Component
			indexTemplate = templates.Layout(
				"Affichage de Notes",     // define title text
				metaTags,                 // define meta tags
				pages.BodyContent(false), // define body content
			)
			return htmx.NewResponse().Retarget("body").RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)

		}

		if user.role != "admin" {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

			metaTags := pages.MetaTags(
				"Affichage de Notes", // define meta description
				"Affichage de Notes", // define meta description
			)
			var indexTemplate templ.Component
			indexTemplate = templates.Layout(
				"Affichage de Notes",       // define title text
				metaTags,                   // define meta tags
				pages.Dashboard(user.role), // define body content
			)
			return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)

		}
		return handler(c, user)
	}
}
func (api *ApiConf) middlewareProfPer(handler authHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		userid := sess.Values["userid"]
		query := `
		SELECT UserID, Username, Session, Role	FROM Users WHERE UserID = ?;
	`
		var user User
		_ = api.db.QueryRow(query, userid).Scan(&user.id, &user.username, &user.session, &user.role)
		if sess.Values["session"] != user.session {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

			metaTags := pages.MetaTags(
				"Affichage de Notes", // define meta description
				"Affichage de Notes", // define meta description
			)
			var indexTemplate templ.Component
			indexTemplate = templates.Layout(
				"Affichage de Notes",     // define title text
				metaTags,                 // define meta tags
				pages.BodyContent(false), // define body content
			)
			return htmx.NewResponse().Retarget("body").RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)

		}

		if user.role != "prof" {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

			metaTags := pages.MetaTags(
				"Affichage de Notes", // define meta description
				"Affichage de Notes", // define meta description
			)
			var indexTemplate templ.Component
			indexTemplate = templates.Layout(
				"Affichage de Notes",       // define title text
				metaTags,                   // define meta tags
				pages.Dashboard(user.role), // define body content
			)
			return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)

		}
		return handler(c, user)
	}
}
