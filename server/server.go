package server

import (
	"context"
	"fmt"
	"log"
	"net"

	Post "github.com/bellacbs/posts/post"
	pd "github.com/bellacbs/posts/proto-buffer"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Server struct {
	PostList []*Post.Post
	clients  map[pd.PostService_GetPostsServer]struct{}
	pd.UnimplementedPostServiceServer
}

func (s *Server) CreatePost(ctx context.Context, req *pd.Post) (*pd.Success, error) {
	fmt.Println("CreatePost method called")
	id := uuid.New()
	post := &Post.Post{
		ID:      id.String(),
		Title:   req.Title,
		Content: req.Content,
	}
	s.PostList = append(s.PostList, post)

	updatePosts := s.getUpdatePost()

	for client := range s.clients {
		err := client.Send(updatePosts)
		if err != nil {
			delete(s.clients, client)
		}
	}
	return &pd.Success{Success: true}, nil
}

func (s *Server) GetPosts(req *pd.Empty, stream pd.PostService_GetPostsServer) error {
	fmt.Println("GetPost method called")
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

	updatePosts := s.getUpdatePost()

	for clientStream := range s.clients {
		if err := clientStream.Send(updatePosts); err != nil {
			return err
		}
	}
	<-stream.Context().Done()
	return nil
}

func (s *Server) getUpdatePost() *pd.Posts {
	updatePosts := &pd.Posts{}
	for _, p := range s.PostList {
		updatePosts.Posts = append(updatePosts.Posts, &pd.Post{
			Id:      p.ID,
			Title:   p.Title,
			Content: p.Content,
		})
	}
	return updatePosts
}

func Init() {
	listen, err := net.Listen("tcp", port)
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
