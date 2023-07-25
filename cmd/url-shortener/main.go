package main

import (
	"RestAPITest/internal/config"
	"fmt"
)

func main() {

	sfg := config.MustLoad()
	fmt.Println(sfg)

	// TODO: init config: cleanenv

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run  server:

}
