package services

import (
	"context"
	"fmt"
	"log"
	"time"

	//"github.com/tom-blog-app/gataway/models"
	postProto "github.com/tom-blog-app/blog-proto/post"
	"google.golang.org/grpc"
)

type PostService interface {
	CreatePost(post *postProto.PostRequest) (*postProto.Post, error)
	GetPost(id string) (*postProto.Post, error)
	DeletePost(id string) (bool, error)
	UpdatePostById(post *postProto.UpdatePostRequest) (*postProto.Post, error)
	GetAllPosts() ([]*postProto.Post, error)
	GetAllPostsByAuthorId(authorId string) ([]*postProto.Post, error)
}

type postService struct {
	client postProto.PostServiceClient
}

func NewPostService() PostService {
	conn, err := grpc.Dial("post-service:50002", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Error connecting to post service:", err)
		return nil
	}
	client := postProto.NewPostServiceClient(conn)
	return &postService{client: client}
}

func (s *postService) CreatePost(post *postProto.PostRequest) (*postProto.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//currentTime := timestamppb.New(time.Now())
	//post.CreatedAt = currentTime
	//post.UpdatedAt = currentTime
	log.Println("Try to send request to post service...")
	res, err := s.client.CreatePost(ctx, &postProto.PostRequest{
		Title:    post.Title,
		Content:  post.Content,
		AuthorId: post.AuthorId,
	})
	log.Println("service CreatePost request received")
	if err != nil {
		return nil, fmt.Errorf("could not create post: %v", err)
	}

	// Return the created post
	return res.GetPost(), nil
}

func (s *postService) GetPost(id string) (*postProto.Post, error) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetPost RPC method
	res, err := s.client.GetPost(ctx, &postProto.GetPostRequest{
		Id: id,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get post: %v", err)
	}

	// Return the fetched post
	return res.GetPost(), nil
}

func (s *postService) DeletePost(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the DeletePost RPC method
	_, err := s.client.DeletePost(ctx, &postProto.GetPostRequest{
		Id: id,
	})
	if err != nil {
		return false, fmt.Errorf("could not delete post: %v", err)
	}

	// Return true if the post was successfully deleted
	return true, nil
}

func (s *postService) UpdatePostById(post *postProto.UpdatePostRequest) (*postProto.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the UpdatePost RPC method
	res, err := s.client.UpdatePost(ctx, &postProto.UpdatePostRequest{
		Id:       post.Id,
		Title:    post.Title,
		Content:  post.Content,
		AuthorId: post.AuthorId,
	})
	if err != nil {
		return nil, fmt.Errorf("could not update post: %v", err)
	}

	// Return the updated post
	return res.GetPost(), nil
}

func (s *postService) GetAllPosts() ([]*postProto.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the ListPosts RPC method
	res, err := s.client.ListPosts(ctx, &postProto.GetPostListRequest{})
	if err != nil {
		return nil, fmt.Errorf("could not list posts: %v", err)
	}
	// Return the list of posts
	return res.GetPosts(), nil
}

func (s *postService) GetAllPostsByAuthorId(authorId string) ([]*postProto.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the ListPosts RPC method
	res, err := s.client.ListPostsByAuthor(ctx, &postProto.GetPostListByAuthorRequest{
		AuthorId: authorId,
	})
	if err != nil {
		return nil, fmt.Errorf("could not list posts: %v", err)
	}

	// Return the list of posts
	return res.GetPosts(), nil
}
