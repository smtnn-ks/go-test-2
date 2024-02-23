package router

import (
	"github.com/gofiber/fiber/v2"
)

type registerDto_t struct {
	Name     string `json:"name" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}

type newsDto_t struct {
	Title      string  `json:"Title" validate:"max=255"`
	Content    string  `json:"Content" validate:"max=1000"`
	Categories []int64 `json:"Categories"`
}

type list_t struct {
	Id int64 `json:"Id"`
	newsDto_t
}

type listOut_t struct {
	Success bool     `json:"Success"`
	News    []list_t `json:"News"`
}

func Init() (app *fiber.App) {
	authApp := fiber.New()
	authApp.Post("/", registerHandler)

	guarded := fiber.New()
	guarded.Use(authMiddleware)
	guarded.Post("/create", createHandler)
	guarded.Post("/edit/:id", editHandler)
	guarded.Get("/list", listHandler)

	app = fiber.New()
	app.Mount("/register", authApp)
	app.Mount("/", guarded)

	return
}
