package main

import (
	"book_service/app"
)

func main() {
	app := app.NewApplication()
	app.Init()
	app.Run()
	app.Wait()
	app.Cleanup()
}
