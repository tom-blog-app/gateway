// post_controller.go
package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
type PostCRUD interface {
	CreatePost(post Post) (Post, error)
	DeletePost(id string) (bool, error)
	GetAllPosts() ([]Post, error)
	UpdatePostById(id string, post Post) (Post, error)
	//DeletePostById(id string) (bool, error)
}

type PostController struct {
	*echo.Echo
}

func (pc *PostController) createPost(c echo.Context) error {
	fmt.Println("createPost")
	//var post Post
	//if err := c.Bind(&post); err != nil {
	//	return err
	//}
	//newPost, err := pc.Service.CreatePost(post)
	//if err != nil {
	//	return err
	//}
	return c.JSON(http.StatusCreated, &Post{
		Id:    "1",
		Title: "test1",
		Body:  "test1",
	})
}

func (pc *PostController) getPostById(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("GetPostById id", id)
	//post, err := pc.Service.GetPostById(id)
	//if err != nil {
	//	return err
	//}
	return c.JSON(http.StatusOK, &Post{
		Id:    id,
		Title: "test1",
		Body:  "test1",
	})
}

func (pc *PostController) updatePostById(c echo.Context) error {
	id := c.Param("id")
	//var post Post
	//if err := c.Bind(&post); err != nil {
	//	return err
	//}
	//updatedPost, err := pc.Service.UpdatePostById(id, post)
	//if err != nil {
	//	return err
	//}
	return c.JSON(http.StatusOK, &Post{
		Id:    id,
		Title: "test1",
		Body:  "test1",
	})
}

func (pc *PostController) deletePostById(c echo.Context) error {
	id := c.Param("id")
	//_, err := pc.Service.DeletePostById(id)
	//if err != nil {
	//	return err
	//}
	return c.JSON(http.StatusOK, id)
}

func (pc *PostController) getAllPosts(c echo.Context) error {
	//posts, err := pc.Service.getAllPosts()
	//if err != nil {
	//	return err
	//}

	posts := []Post{{
		Id:    "1",
		Title: "test1",
		Body:  "test1",
	}}
	return c.JSON(http.StatusOK, posts)
}

func SetupPostController(e *echo.Echo) {
	controller := &PostController{e}
	fmt.Println("setup post controller")
	e.GET("/help", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!2")
	})
	postGroup := e.Group("/posts")
	postGroup.POST("/", controller.createPost)
	postGroup.GET("/:id", controller.getPostById)
	postGroup.PUT("/:id", controller.updatePostById)
	postGroup.DELETE("/:id", controller.deletePostById)
	postGroup.GET("", controller.getAllPosts)
}
