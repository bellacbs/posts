package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	Post "github.com/bellacbs/posts/post"
	pd "github.com/bellacbs/posts/proto-buffer"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server struct {
	PostList []*Post.Post
	pd.UnimplementedPostServiceServer
}

func (s *Server) CreatePost(ctx context.Context, req *pd.Post) (*pd.Success, error) {
	id := uuid.New()
	post := &Post.Post{
		ID:      id.String(),
		Title:   req.Title,
		Content: req.Content,
	}
	s.PostList = append(s.PostList, post)
	return &pd.Success{Success: true}, nil
}

func (s *Server) GetPosts(ctx context.Context, req *pd.Empty) (*pd.Posts, error) {
	posts := &pd.Posts{}
	for _, p := range s.PostList {
		posts.Posts = append(posts.Posts, &pd.Post{
			Id:      p.ID,
			Title:   p.Title,
			Content: p.Content,
		})
	}
	return posts, nil
}

func Init() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed too listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server := &Server{}
	pd.RegisterPostServiceServer(grpcServer, server)
	fmt.Printf("Server listening on port 5001...\n")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
