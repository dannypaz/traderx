package broker_rpc

import (
	"context"
	fmt "fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

var serverAddr = "localhost:27492"

// HealthCheck a sparkswapd grpc client
func HealthCheck() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	client := NewAdminServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.HealthCheck(ctx, &GoogleProtobuf_Empty{})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("received response: %v", res)

	defer conn.Close()
}
