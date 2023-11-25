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
	clients  map[pd.PostService_GetPostsServer]struct{}
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

	updatePosts := &pd.Posts{}
	for _, p := range s.PostList {
		updatePosts.Posts = append(updatePosts.Posts, &pd.Post{
			Id:      p.ID,
			Title:   p.Title,
			Content: p.Content,
		})
	}

	for client := range s.clients {
		err := client.Send(updatePosts)
		if err != nil {
			delete(s.clients, client)
		}
	}
	return &pd.Success{Success: true}, nil
}

func (s *Server) GetPosts(req *pd.Empty, stream pd.PostService_GetPostsServer) error {
	if s.clients == nil {
		s.clients = make(map[pd.PostService_GetPostsServer]struct{})
	}
	updateChanel := make(chan *pd.Posts)
	s.clients[stream] = struct{}{}
	go func() {
		defer close(updateChanel)
		for {
			select {
			case <-stream.Context().Done():
				delete(s.clients, stream)
				return
			}
		}
	}()

	for update := range updateChanel {
		if err := stream.Send(update); err != nil {
			return nil
		}
	}
	return nil
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
