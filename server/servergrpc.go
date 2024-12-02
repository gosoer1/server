package server

import (
	grpc_server "app/grpc/grpc"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	grpc_server.UnimplementedNumberServiceServer
}

func (s *server) SendNumber(ctx context.Context, req *grpc_server.NumberRequest) (*grpc_server.NumberResponse, error) {
	number := req.GetNumber()
	result := number * number

	fmt.Printf("Степень: %d\n", result)
	return &grpc_server.NumberResponse{
		Result:  int32(result),
		Message: fmt.Sprintf("Степень %d !", result)}, nil
}

func ServerGrps(done chan bool) {
	grpcServer := grpc.NewServer()
	grpc_server.RegisterNumberServiceServer(grpcServer, &server{})
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	fmt.Println("Сервер запущен на порту :50051")
	done <- true
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка при запуске gRPC сервера: %v", err)
	}
}

func Togrps() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close()

	client := grpc_server.NewNumberServiceClient(conn)

	var num int
	fmt.Print("Введите число: ")
	_, err = fmt.Scan(&num)
	if err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}

	req := &grpc_server.NumberRequest{Number: int32(num)}
	resp, err := client.SendNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове метода: %v", err)
	}

	fmt.Println("Ответ от сервера:", resp.GetMessage())
}
