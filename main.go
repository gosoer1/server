package main

import (
	//"context"
	"app/server"
	"fmt"
	//"log"
	//"net"
	//"google.golang.org/grpc"
)

type server_changer interface {
	changer()
}

type grpc struct{}
type http struct{}

func (a *grpc) changer() {

	done := make(chan bool)
	go server.ServerGrps(done)
	<-done
	server.Togrps()

	select {}
}

func (a *http) changer() {
	var X int
	fmt.Println("Сервер http запущен")
	fmt.Println("Введите число Х")
	fmt.Scan(&X)
	server.Get(X)
	fmt.Println("Ваш сервер по адресу: localhost:8080")
	go server.ServerHttp()

	select {}
}

func main() {
	var chang_name string
	fmt.Println("Выберите сервер http or grpc")
	fmt.Scan(&chang_name)
	var s server_changer
	if chang_name == "grpc" {
		s = &grpc{}
	} else if chang_name == "http" {
		s = &http{}
	} else {
		fmt.Println("Сервера такого нет ты криворукое чучело")
	}

	s.changer()

}
