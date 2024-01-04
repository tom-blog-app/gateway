package controllers

//
//import (
//	"github.com/labstack/echo/v4"
//)
//
//type User struct {
//	Id    string `json:"id"`
//	Email string `json:"email"`
//}
//
//type UserCRUD interface {
//	CreateUser(user User) (User, error)
//	DeleteUser(id string) (User, error)
//}
//
//type UserController struct {
//}
//
//func CreateUser(ctx echo.Context) (User, error) {
//	return User{
//		Id:    "1",
//		Email: "test",
//	}, nil
//}
//
//func DeleteUser(ctx echo.Context) (bool, error) {
//	return true, nil
//}
//
//func (uc *UserController) SetupUserController(e *echo.Echo) {
//	controller := &UserController{
//		e,
//	}
//
//	e.POST("/users", controller.CreateUser)
//	e.GET("/users/:id", controller.GetUserById)
//	e.PUT("/users/:id", controller.UpdateUserById)
//	e.DELETE("/users/:id", controller.DeleteUserById)
//	e.GET("/users", controller.GetAllUsers)
//}
