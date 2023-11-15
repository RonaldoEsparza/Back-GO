package models

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

// Mock data para simular una base de datos
var users = []User{
    {ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

// Funciones para operaciones CRUD básicas

func GetAllUsers() []User {
    return users
}

func GetUserByID(id int) *User {
    for _, user := range users {
        if user.ID == id {
            return &user
        }
    }
    return nil
}

func CreateUser(user User) {
    user.ID = len(users) + 1
    users = append(users, user)
}

func UpdateUser(id int, updatedUser User) bool {
    for i, user := range users {
        if user.ID == id {
            users[i] = updatedUser
            return true
        }
    }
    return false
}

func DeleteUser(id int) bool {
    for i, user := range users {
        if user.ID == id {
            // Eliminar el elemento de users
            users = append(users[:i], users[i+1:]...)
            return true
        }
    }
    return false
}
