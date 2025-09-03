package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define route for serving HTML at root path
	e.GET("/", func(c echo.Context) error {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go Echo Server</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; }
				h1 { color: #333; }
			</style>
		</head>
		<body>
			<h1>Hello from Go Echo Framework!</h1>
			<p>Server is running on port 8080</p>
		</body>
		</html>
		`
		return c.HTML(http.StatusOK, html)
	})

	// Start server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
