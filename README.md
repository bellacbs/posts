<h1 align="center" id="top">== Posts - Golang ==</h1>

<p align="center">
  <a href="#sobre">About</a> &#xa0; | &#xa0; 
  <a href="#funciona">What works</a> &#xa0; | &#xa0;
  <a href="#pendente">In development</a> &#xa0; | &#xa0;
  <a href="#requirements">Requirements</a> &#xa0; | &#xa0;
</p>

<h2 id="sobre">:notebook: About </h2>

<p align="center">:rocket: Project developed to demonstrate two type of gRPC (Unary and Server Streaming)</p>

<h2 id="tecnologias"> 🛠 Technologies and programming languages </h2>

The following libraries and languages were used in the project's construction:

* Go
* gRPC
* Makefile
* Gin

<h2 id="funciona">:heavy_check_mark: What works</h2>

* Send Posts using unary call;</br>
* Receiving new posts in real time using Server Streaming;</br>

 
<h2 id="pendente">:construction: In development</h2>

- [x] Client streaming;
- [x] Bidirectional streaming;
- [x] JS client;
- [x] Clean Architecture;

<h2 id="requirements">:leftwards_arrow_with_hook: Prerequisites</h2>

Before you start, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Go](https://go.dev/doc/install), [gRPC and Protobuffer package dependencies](https://grpc.io/docs/languages/go/quickstart/) and if you want the development mode, install [gin](https://github.com/codegangsta/gin) to live reload. 
Additionally, it's good to have a code editor to work with, such as [VSCode](https://code.visualstudio.com/)

## Installing protobuffer

### Linux

```
sudo apt install -y protobuf-compiler
```

### MacOS

```
brew install protobuff
```

### gRPC and Protobuffer package dependencies
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

NOTE: You should add the `protoc-gen-go-grpc` to your PATH

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Install gin

```
go get github.com/codegangsta/gin
go install github.com/codegangsta/gin
```

<h4>:checkered_flag: Running the project </h4>

```bash
# Clone this repository

# To start the server as dev
$ make run dev

# To build and start server
$ make run build
$ make run run

# The server will start on port 50051 - access it on <grpc://localhost:50051>

# To run client, in the source of this directory, open a new terminal, and execute 
$ go run client/golang/main.go

# you can also use clients like Postman and Insomnia by importing the proto file at /proto/post.proto
```


<a href="#top">Go back to the top</a>
