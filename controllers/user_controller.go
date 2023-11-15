package controllers

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "Back-Flutter/models"
    "strconv"
)

// Obtener todos los usuarios
func GetAllUsers(c echo.Context) error {
    users := models.GetAllUsers()
    return c.JSON(http.StatusOK, users)
}

// Obtener un usuario por ID
func GetUserByID(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.String(http.StatusBadRequest, "ID inválido")
    }

    user := models.GetUserByID(id)
    if user == nil {
        return c.String(http.StatusNotFound, "Usuario no encontrado")
    }

    return c.JSON(http.StatusOK, user)
}

// Crear un nuevo usuario
func CreateUser(c echo.Context) error {
    var newUser models.User
    if err := c.Bind(&newUser); err != nil {
        return c.String(http.StatusBadRequest, "Datos de usuario no válidos")
    }

    models.CreateUser(newUser)
    return c.String(http.StatusCreated, "Usuario creado exitosamente")
}

// Actualizar un usuario existente
func UpdateUser(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.String(http.StatusBadRequest, "ID inválido")
    }

    var updatedUser models.User
    if err := c.Bind(&updatedUser); err != nil {
        return c.String(http.StatusBadRequest, "Datos de usuario no válidos")
    }

    success := models.UpdateUser(id, updatedUser)
    if success {
        return c.String(http.StatusOK, "Usuario actualizado exitosamente")
    }

    return c.String(http.StatusNotFound, "Usuario no encontrado")
}

// Eliminar un usuario por ID
func DeleteUser(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.String(http.StatusBadRequest, "ID inválido")
    }

    success := models.DeleteUser(id)
    if success {
        return c.String(http.StatusOK, "Usuario eliminado exitosamente")
    }

    return c.String(http.StatusNotFound, "Usuario no encontrado")
}
