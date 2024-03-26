package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Saakhr/go-affichage-de-notes/templates"
	"github.com/Saakhr/go-affichage-de-notes/templates/pages"
	"github.com/Saakhr/go-affichage-de-notes/templates/pages/comps"
	"github.com/a-h/templ"
	"github.com/google/uuid"

	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// indexViewHandler handles a view for the index page.

func (api *ApiConf) indexViewHandler(c echo.Context, user User) error {

	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"Affichage de Notes", // define meta description
		"Affichage de Notes", // define meta description
	)

	// Define template body content.

	// Define template layout for index page.
	var indexTemplate templ.Component
	indexTemplate = templates.Layout(
		"Affichage de Notes",       // define title text
		metaTags,                   // define meta tags
		pages.Dashboard(user.role), // define body content
	)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)

}

// showContentAPIHandler handles an API endpoint to show content.
func (api *ApiConf) add_etu_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	specs := getGroups(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_etu(specs))
}
func (api *ApiConf) mng_etu_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	specs := getGroups(api.db)
	users := getAllStudentswithspec(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_etu(users, specs))
}
func (api *ApiConf) add_etu(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	specs := getGroups(api.db)
	etucname := c.FormValue("user")
	password := c.FormValue("password")
	specname := c.FormValue("spec")
	var idspec int
	for _, v := range specs {
		if v.Name == specname {
			idspec = v.ID
		}
	}
	addStudant(api.db, [2]string{etucname, password}, idspec)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_etu(specs))
}
func (api *ApiConf) edit_etu(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	specs := getGroups(api.db)
	etucname := c.FormValue("user")
	password := c.FormValue("password")
	groupid, err := strconv.Atoi(c.FormValue("spec"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	index, err := strconv.Atoi(c.QueryParam("index"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	updateStudant(api.db, [2]string{etucname, password}, groupid, userid)
	student := getAStudentwithspec(api.db, userid)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, comps.Item_etu(index, student, &specs))
}
func (api *ApiConf) delete_etu(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteStudant(api.db, userid)
	return c.NoContent(200)

}

func add_spec_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_spec())
}
func (api *ApiConf) add_spec(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	specname := c.FormValue("name")
	addSpec(api.db, specname)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_spec())
}
func (api *ApiConf) mng_spec_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	specs := getSpecialities(api.db)
	mods := getMods(api.db)
	users := getAllSpecsWithModules(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_spec(users, mods, specs))
}

func (api *ApiConf) edit_spec(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	etucname := c.FormValue("user")
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteSpecCourses(api.db, userid)
	updateSpec(api.db, etucname, userid)
	modules := c.Request().Form["modules"]
	modulescoe := c.Request().Form["modules-coe"]
	modulescoe = removeEmptyStrings(modulescoe)
	if len(modules) != len(modulescoe) {
		return c.NoContent(http.StatusBadRequest)
	}
	modulid, err := stringsToInts(modules)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	coes, err := stringsToInts(modulescoe)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	for i, v := range modulid {
		addSpecCourse(api.db, [3]int{v, userid, coes[i]})
	}

	specs := getSpecialities(api.db)
	mods := getMods(api.db)
	users := getAllSpecsWithModules(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_spec(users, mods, specs))

}
func (api *ApiConf) delete_spec(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteSpec(api.db, userid)
	return c.NoContent(200)

}
func add_prof_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_prof())
}
func (api *ApiConf) add_prof(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	etucname := c.FormValue("user")
	password := c.FormValue("password")
	addProf(api.db, [2]string{etucname, password})
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_prof())
}
func (api *ApiConf) mng_prof_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	profwith := getAllProfsWithModulesAndGroups(api.db)
	profs := getAllProfs(api.db)
	gmodules := getAllGroupsAndCourses(api.db)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_prof(profwith, gmodules, profs))
}
func (api *ApiConf) edit_prof(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	etucname := c.FormValue("user")
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteProfModules(api.db, userid)
	updateProfUsername(api.db, userid, etucname)
	groups := c.Request().Form["group"]
	if len(groups) == 0 {
		return c.NoContent(200)
	}
	groupint, err := stringsToInts(groups)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	for i, v := range groups {
		modug := c.Request().Form["group-"+v]
		if len(modug) == 0 {
			return c.NoContent(http.StatusBadRequest)
		}
		modulid, err := stringsToInts(modug)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		for _, v := range modulid {
			addTeacherModule(api.db, userid, v, groupint[i])
		}

	}
	profwith := getAllProfsWithModulesAndGroups(api.db)
	profs := getAllProfs(api.db)
	gmodules := getAllGroupsAndCourses(api.db)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_prof(profwith, gmodules, profs))

}
func (api *ApiConf) delete_prof(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteProf(api.db, userid)
	return c.NoContent(200)

}
func (api *ApiConf) add_mod_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_mod())
}
func (api *ApiConf) add_mod(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	name := c.FormValue("name")
	if name == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	addMod(api.db, name)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_mod())
}
func (api *ApiConf) mng_mod_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	specs := getMods(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_mod(specs))
}
func (api *ApiConf) edit_mod(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	etucname := c.FormValue("user")
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	index, err := strconv.Atoi(c.QueryParam("index"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	updateMod(api.db, etucname, userid)
	student := getAMod(api.db, userid)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, comps.Item_mod(index, student))
}
func (api *ApiConf) delete_mod(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteMod(api.db, userid)
	return c.NoContent(200)

}
func (api *ApiConf) add_grp_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	specs := getSpecialities(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_group(specs))
}
func (api *ApiConf) add_grp(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	name := c.FormValue("name")
	if name == "" {
		fmt.Println("dd")
		return c.NoContent(http.StatusBadRequest)
	}
	spec := c.FormValue("spec")
	specid, err := strconv.Atoi(spec)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	addGrp(api.db, name, specid)

	specs := getSpecialities(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Add_group(specs))
}
func (api *ApiConf) mng_grp_page(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	specs := getSpecialities(api.db)
	users := getAllGroupswithspecNAME(api.db)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.Mng_grp(users, specs))
}
func (api *ApiConf) edit_grp(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	specs := getGroups(api.db)
	etucname := c.FormValue("user")
	groupid, err := strconv.Atoi(c.FormValue("spec"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	index, err := strconv.Atoi(c.QueryParam("index"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	updateGrp(api.db, etucname, userid, groupid)
	student := getAGroupwithspec(api.db, userid)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, comps.Item_grp(index, student, &specs))
}
func (api *ApiConf) delete_grp(c echo.Context, user User) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	userid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	deleteGrp(api.db, userid)
	return c.NoContent(200)

}

// login

func (api *ApiConf) LoginApiHandler(c echo.Context) error {
	if !htmx.IsHTMX(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	user := c.FormValue("user")
	pass := c.FormValue("password")
	query := "SELECT UserID, Password FROM Users WHERE Username = ?"
	row := api.db.QueryRow(query, user)
	var userID int
	var password string
	err := row.Scan(&userID, &password)
	if err != nil {
		fmt.Println("User not found", err)
		return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.BodyContent(true))
	}
	if password == pass {
		fmt.Println("Login successful! UserID:", userID)
		id := uuid.NewString()
		updateSession(api.db, userID, id)
		sess, _ := session.Get("session", c)
		sess.Values["session"] = id
		sess.Values["userid"] = userID
		//store session
		err := sess.Save(c.Request(), c.Response())
		if err != nil {
			log.Println("Session error: ", err)
		}
		query := `
		SELECT UserID, Username, Session, Role	FROM Users WHERE UserID = ?;
	`
		var user User
		_ = api.db.QueryRow(query, userID).Scan(&user.id, &user.username, &user.session, &user.role)
		return api.indexViewHandler(c, user)

	} else {
		fmt.Println("Incorrect password")
		return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, pages.BodyContent(true))
	}
}
