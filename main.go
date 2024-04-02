package main

import (
	"github.com/russelldias98/go-api/initializers"
	"github.com/russelldias98/go-api/routes"
)

func main() {
	initializers.SetupEnv()
	initializers.ConnectToDB()
	routes.SetupRouter()
}
