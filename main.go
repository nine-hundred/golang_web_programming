package main

import "golang_web_programming/app"

func main() {
	config := app.DefaultConfig()
	e := app.NewEcho(*config)

	e.Logger.Fatal(e.Start(":8080"))
}
