package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserView(ctx *fiber.Ctx) error {
	//cookie_jar := ctx.Locals("cookie_jar").(*session.Store)

	database, err := sql.Open("mysql", DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	username := ctx.Params("username")
	stmt, err := database.Prepare("SELECT id FROM users WHERE username = ?;")
	if err != nil {
		return send_internal_error(ctx)
	}

	defer stmt.Close()
	row := stmt.QueryRow(username)

	user_model := User{}

	err = row.Scan(
		&user_model.ID,
	)

	if err != nil {
		return send_error(ctx, UserNotFound)
	}

	return ctx.Render("profile", fiber.Map{
		"Username": username,
	})
}

func LoginView(ctx *fiber.Ctx) error {
	cookie_jar := ctx.Locals("cookie_jar").(*session.Store)

	session, err := cookie_jar.Get(ctx)
	if err != nil {
		return send_redirect(ctx, "/")
	}

	if !session.Fresh() {
		return send_redirect(ctx, "/")
	}

	return ctx.SendFile("./views/login.html")
}

func RegisterView(ctx *fiber.Ctx) error {
	cookie_jar := ctx.Locals("cookie_jar").(*session.Store)

	session, err := cookie_jar.Get(ctx)
	if err != nil {
		return send_redirect(ctx, "/")
	}

	if !session.Fresh() {
		return send_redirect(ctx, "/")
	}

	return ctx.SendFile("./views/register.html")
}

func DashboardView(ctx *fiber.Ctx) error {
	cookie_jar := ctx.Locals("cookie_jar").(*session.Store)

	database, err := sql.Open("mysql", DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	session, err := cookie_jar.Get(ctx)
	if err != nil {
		return send_redirect(ctx, "/login")
	}

	if session == nil || session.Fresh() || session.Get("id") == nil {
		session.Destroy()
		return send_redirect(ctx, "/login")
	}

	id := session.Get("id").(int64)
	stmt, err := database.Prepare("SELECT username FROM users WHERE id = ?;")
	if err != nil {
		return send_internal_error(ctx)
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	var username string

	if err := row.Scan(&username); err != nil {
		return send_internal_error(ctx)
	}

	return ctx.Render("dashboard", fiber.Map{
		"Username": username,
	})
}
