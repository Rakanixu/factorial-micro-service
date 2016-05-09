package handler

import (
	proto "github.com/Rakanixu/factorial-micro-service/server/proto"
	"golang.org/x/net/context"
	"github.com/Rakanixu/factorial"
)

type Factorial struct{}

func (g *Factorial) CalcFactorial(ctx context.Context, req *proto.FactorialRequest, rsp *proto.FactorialResponse) error {
	rsp.Result = factorial.CalculateFactorial(req.Number)
	return nil
}
