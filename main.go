package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template struct implements echo.Renderer interface
type Template struct {
	templates *template.Template
}

// Render method for Template struct
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Create a new Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	// Helper function to create template data with common fields
	templateData := func(title string) map[string]interface{} {
		return map[string]interface{}{
			"Port":        port,
			"Title":       title,
			"CurrentTime": time.Now().UTC().Format("2006-01-02 15:04:05"),
		}
	}

	// Define route for serving HTML at root path
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", templateData("Home"))
	})

	// Define route for about page
	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", templateData("About"))
	})

	// Start server on the configured port
	e.Logger.Fatal(e.Start(":" + port))
}
