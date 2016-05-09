package main

import (
	"github.com/micro/go-micro"
	proto "github.com/Rakanixu/factorial-micro-service/server/proto"
	"golang.org/x/net/context"
	"fmt"
)

func main() {
	service := micro.NewService(micro.Name("factorial.client"))

	factorial := proto.NewFactorialClient("factorial", service.Client())

	rsp, err := factorial.CalcFactorial(
		context.TODO(),
		&proto.FactorialRequest{
			Number: 7,
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Result)
}
