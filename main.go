package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "Back-Flutter/controllers"
)

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Rutas CRUD para usuarios
    e.GET("/users", controllers.GetAllUsers)
    e.GET("/users/:id", controllers.GetUserByID)
    e.POST("/users", controllers.CreateUser)
    e.PUT("/users/:id", controllers.UpdateUser)
    e.DELETE("/users/:id", controllers.DeleteUser)

    // Iniciar el servidor web
    e.Start(":8080")
}
