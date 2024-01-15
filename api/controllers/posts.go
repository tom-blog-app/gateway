// post_controller.go
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	postProto "github.com/tom-blog-app/blog-proto/post"
	"github.com/tom-blog-app/gataway/api/services"
	"log"
	"net/http"
)

//type Post struct {
//	Id    string `json:"id"`
//	Title string `json:"title"`
//	Body  string `json:"body"`
//}

func SetupPostController(e *echo.Echo) {
	controller := &PostController{
		e,
		services.NewPostService(),
	}
	fmt.Println("Setup Post controller...")
	postGroup := e.Group("/posts")
	postGroup.POST("", controller.createPost)
	postGroup.GET("/:id", controller.getPostById)
	postGroup.PUT("/:id", controller.updatePostById)
	postGroup.DELETE("/:id", controller.deletePostById)
	postGroup.GET("", controller.getAllPosts)
	postGroup.GET("/author/:authorId", controller.getAllPostsByAuthor)
}

type PostController struct {
	*echo.Echo
	postService services.PostService
}

func (pc *PostController) createPost(c echo.Context) error {
	post := new(postProto.PostRequest)
	if err := json.NewDecoder(c.Request().Body).Decode(post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	log.Println("CreatePost request received", post)
	newPost, err := pc.postService.CreatePost(post)
	if err != nil {
		log.Println("Error1: ", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newPost)
}

func (pc *PostController) getPostById(c echo.Context) error {
	id := c.Param("id")
	post, err := pc.postService.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (pc *PostController) updatePostById(c echo.Context) error {
	id := c.Param("id")

	var post postProto.UpdatePostRequest

	if err := c.Bind(&post); err != nil {
		return err
	}

	post.Id = id
	updatedPost, err := pc.postService.UpdatePostById(&post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedPost)
}

func (pc *PostController) deletePostById(c echo.Context) error {
	id := c.Param("id")
	success, err := pc.postService.DeletePost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]bool{"success": success})
}

func (pc *PostController) getAllPosts(c echo.Context) error {
	posts, err := pc.postService.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func (pc *PostController) getAllPostsByAuthor(c echo.Context) error {
	authorId := c.Param("authorId")
	posts, err := pc.postService.GetAllPostsByAuthorId(authorId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}
