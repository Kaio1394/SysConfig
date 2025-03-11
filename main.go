package main

import (
	"SysConfig/database"
	"SysConfig/internal/routes"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	dd, err := database.ConnectDatabase()
	if err != nil {
		return
	}
	server := gin.Default()

	routes.RegisterRouteAgent(server, dd)
	routes.RegisterRouteLog(server, dd)

	_ = server.Run(":8080")
}
