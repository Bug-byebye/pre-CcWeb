package main

import (
	"CC-web/grpc/auth_rpc"
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
)

type server struct {
	auth_rpc.UnimplementedAuthServiceServer
}

func (s *server) Login(ctx context.Context, req *auth_rpc.LoginRequest) (*auth_rpc.LoginResponse, error) {
	// 这里可以加入验证用户名密码的逻辑
	// 假设我们验证成功并生成一个 token
	return &auth_rpc.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "Login successful",
		UserId:     123,
		Token:      "some-jwt-token",
	}, nil
}

func (s *server) Register(ctx context.Context, req *auth_rpc.RegisterRequest) (*auth_rpc.RegisterResponse, error) {
	// 这里可以加入注册逻辑，例如检查用户名是否存在等
	// 假设注册成功并生成一个 token
	return &auth_rpc.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "Registration successful",
		UserId:     123,
		Token:      "some-jwt-token",
	}, nil
}

func (s *server) Authenticate(ctx context.Context, req *auth_rpc.AuthenticateRequest) (*auth_rpc.AuthenticateResponse, error) {
	// 假设我们简单验证 token 是否有效
	if req.Token == "some-jwt-token" {
		return &auth_rpc.AuthenticateResponse{
			StatusCode: 0,
			StatusMsg:  "Authentication successful",
			UserId:     123,
		}, nil
	}

	return &auth_rpc.AuthenticateResponse{
		StatusCode: 1,
		StatusMsg:  "Invalid token",
	}, nil
}

func main() {
	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth_rpc.RegisterAuthServiceServer(s, &server{})

	fmt.Println("Auth service is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
