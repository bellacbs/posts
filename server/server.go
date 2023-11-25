package server

import (
	"context"
	"flag"
	"log"
	"net"

	pd "github.com/bellacbs/posts/proto-buffer"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type grpcServer struct {
	pd.UnimplementedPostServiceServer
}

func (s *grpcServer) CreatePost(ctx context.Context, req *pd.Post) (*pd.Success, error) {
	log.Printf(req.GetContent())
	log.Printf(req.GetTitle())
	return &pd.Success{Success: true}, nil
}

func (s *grpcServer) GetPosts(ctx context.Context, req *pd.Empty) (*pd.Posts, error) {
	post := &pd.Post{
		Id:      "123456",
		Title:   "Test",
		Content: "Content Text",
	}
	posts := &pd.Posts{Posts: []*pd.Post{post}}
	return posts, nil
}

func Init() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed too listen: %v", err)
	}
	server := grpc.NewServer()
	pd.RegisterPostServiceServer(server, &grpcServer{})
	if err := server.Serve(listen); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
