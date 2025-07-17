package grpc

// cmd/server/main.go
import (
	"context"
	"log"
	"net"
	"net/http"

	"cex-core-api/app/internal/service/order"
	order_v1 "cex-core-api/gen/order/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run() {
	grpcAddr := ":9090"
	httpAddr := ":8081"

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	order_v1.RegisterOrderServiceServer(s, &order.OrderService{})

	go func() {
		log.Printf("gRPC server listening at %s", grpcAddr)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := order_v1.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	log.Printf("HTTP Gateway listening at %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
