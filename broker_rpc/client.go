package broker_rpc

import (
	fmt "fmt"
	"log"
	"os"

	"google.golang.org/grpc"
)

var serverAddr = "localhost:27492"

// AdminService a sparkswapd grpc client
func HealthCheck() {
	conn, err := grpc.Dial(serverAddr)

	if err != nil {
		fmt.Printf("error happened so whatever")
		log.Fatal(err)
		os.Exit(1)
	}

	client := NewAdminServiceClient(conn)
	res, err := client.HealthCheck()

	if err != nil {
		fmt.Errorf("Healthcheck failed for some reason: %v", err)
		log.Fatal(err)
		os.Exit(1)
	}

	defer conn.Close()
}
