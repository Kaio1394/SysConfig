package main

import (
	"SysConfig/config"
	"SysConfig/database"
	"SysConfig/internal/middleware"
	"SysConfig/internal/routes"
	"github.com/gin-gonic/gin"
	"strconv"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	dd, err := database.ConnectDatabase()
	if err != nil {
		return
	}
	server := gin.Default()
	server.Use(middleware.AuthMiddleware())
	routes.RegisterRouteAgent(server, dd)
	routes.RegisterRouteLog(server, dd)
	_ = server.Run(":" + strconv.Itoa(config.ConfigViper.Server.Port))
}
