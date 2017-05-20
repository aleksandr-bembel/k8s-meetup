package main

import (
	"fmt"
	"log"
	"os"

	"github.com/takama/router"
)

// Run server: go run main.go
// Try few requests:
// - curl http://localhost:8000
// - curl http://localhost:8000/test-some-path
func main() {
	logger := log.New(os.Stdout, "[step-by-step]", log.LstdFlags)
	logger.Print("Приложение запускается...")

	r := router.New()

	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		logger.Fatal("Порт не задан")
	}

	r.GET("/", home)

	logger.Print("Готовимся слушать порт " + port)
	r.Listen(":" + port)
}

func home(c *router.Control) {
	log.Print("Хэндрел home поймал запрос")
	fmt.Fprintf(
		c.Writer, "URL.Path = %q\n", c.Request.URL.Path,
	)

}
