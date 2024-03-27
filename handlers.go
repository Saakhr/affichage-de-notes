package main

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/Saakhr/go-affichage-de-notes/templates"
	"github.com/Saakhr/go-affichage-de-notes/templates/pages"
	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (api *ApiConf) add_note_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	students := getStudentsByTeacherID(api.db, user.id)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_note(students))
}
func (api *ApiConf) add_note(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	tds := c.FormValue("TD")
	var TD, TP, EXAM, PROJECT float64
	var err error
	if tds == "" {
		TD = -1
	} else {
		TD, err = strconv.ParseFloat(tds, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	tps := c.FormValue("TP")
	if tps == "" {
		TP = -1
	} else {
		TP, err = strconv.ParseFloat(tps, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	exams := c.FormValue("EXAM")
	if exams == "" {
		EXAM = -1
	} else {
		EXAM, err = strconv.ParseFloat(exams, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	Projects := c.FormValue("PROJECT")
	if Projects == "" {
		PROJECT = -1
	} else {
		PROJECT, err = strconv.ParseFloat(Projects, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	options := strings.Split(c.FormValue("stud"), "-")
	if len(options) != 2 {
		return c.NoContent(http.StatusBadRequest)
	}
	userid, err := strconv.Atoi(options[0])
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	courseid, err := strconv.Atoi(options[1])
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	addGrade(api.db, user.id, courseid, userid, [4]float64{EXAM, TD, TP, PROJECT})

	students := getStudentsByTeacherID(api.db, user.id)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_note(students))
}

func (api *ApiConf) mng_note_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	grads := getGradesByTeacherID(api.db, user.id)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_notes(grads))
}
func (api *ApiConf) mng_note(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	gradeid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tds := c.FormValue("TD")
	var TD, TP, EXAM, PROJECT float64
	if tds == "" {
		TD = -1
	} else {
		TD, err = strconv.ParseFloat(tds, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	tps := c.FormValue("TP")
	if tps == "" {
		TP = -1
	} else {
		TP, err = strconv.ParseFloat(tps, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	exams := c.FormValue("EXAM")
	if exams == "" {
		EXAM = -1
	} else {
		EXAM, err = strconv.ParseFloat(exams, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	Projects := c.FormValue("PROJECT")
	if Projects == "" {
		PROJECT = -1
	} else {
		PROJECT, err = strconv.ParseFloat(Projects, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	updateGrade(api.db, gradeid, [4]float64{EXAM, TD, TP, PROJECT})

	grads := getGradesByTeacherID(api.db, user.id)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_notes(grads))

}
func (api *ApiConf) delete_note(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	gradeid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteGrade(api.db, gradeid)
	return c.NoContent(200)
}

func (api *ApiConf) SeeNotes(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	grads := getGradesWithSpecializationAndGroup(api.db, user.id)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.See_notes(grads))
}
func (api *ApiConf) Logout(c echo.Context, user User) error {
	sess, _ := session.Get("session", c)
	sess.Values["userid"] = ""
	sess.Values["session"] = ""
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
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)

}
