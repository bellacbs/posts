package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/bellacbs/posts/proto-buffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPostServiceClient(conn)

	unary(client)
	serverSide(client)
}

func unary(client pb.PostServiceClient) {
	createPostResponse, err := client.CreatePost(context.Background(), &pb.Post{
		Title:   "Post Title",
		Content: "Content Post",
	})
	if err != nil {
		log.Fatalf("Error creating post: %v", err)
	}
	log.Printf("CreatePost response: %v", createPostResponse.Success)
}

func serverSide(client pb.PostServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	getPostsStream, err := client.GetPosts(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error getting posts: %v", err)
	}

	for {
		post, err := getPostsStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving post: %v", err)
		}

		log.Printf("Received post: %v", post)
	}
}
