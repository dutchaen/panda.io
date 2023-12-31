package main

import (
	"bamboo/api"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	cookie_jar := session.New(session.Config{
		Expiration: time.Until(time.Now().Add(time.Hour * 24 * 30 * 12)),
	})

	db, err := sql.Open("mysql", api.DATABASE_AUTHENTICATION)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to the Database.")

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./views/static")

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("cookie_jar", cookie_jar)

		return ctx.Next()
	})

	app.Get("/", api.DashboardView)
	app.Get("/login", api.LoginView)
	app.Get("/register", api.RegisterView)
	app.Get("/@:username", api.UserView)

	app.Post("/api/login", api.LoginRest)
	app.Post("/api/register", api.RegisterRest)
	app.Get("/api/user/@:username", api.UserLookupRest)

	app.Get("/api/account/@me", api.GetMeRest)
	app.Post("/api/account/@me", api.WriteMeRest)

	fmt.Println("opening bamboo server")

	log.Fatal(app.Listen(":3000"))
}
