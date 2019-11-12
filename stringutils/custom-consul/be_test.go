package custom_consul

import (
	"context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"testing"
	"time"
)
const (
	target      = "consul://trungtvq.ddns.net:8500/helloworld"
	defaultName = "world"
)

func BenchmarkF(b *testing.B) {
	// run the Fib function b.N times

	Init()
	// Set up a connection to the server.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, _ := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName


	for n := 0; n < b.N; n++ {

		ctx, _ = context.WithTimeout(context.Background(), time.Second)
		c.SayHello(ctx, &pb.HelloRequest{Name: name})
	}
}