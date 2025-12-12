package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "grpc-crud-microservice/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func (s *userServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var id int32
	err := s.db.QueryRowContext(
		ctx,
		"INSERT INTO users(name, email) VALUES($1, $2) RETURNING id",
		req.Name, req.Email,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	var id int32
	var name, email string

	err := s.db.QueryRowContext(
		ctx,
		"SELECT id, name, email FROM users WHERE id=$1",
		req.Id,
	).Scan(&id, &name, &email)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    id,
		Name:  name,
		Email: email,
	}, nil
}

func main() {
	connStr := "host=localhost port=5432 user=admin password=password dbname=usersdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("db open error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("db ping error:", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServer{db: db})

	fmt.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("serve error:", err)
	}
}
