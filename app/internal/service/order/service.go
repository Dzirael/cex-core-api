package order

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "cex-core-api/gen/order/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func (s *OrderService) GetSupportedPairs(ctx context.Context, req *pb.GetSupportedPairsRequest) (*pb.GetSupportedPairsResponse, error) {
	return &pb.GetSupportedPairsResponse{
		Pairs: []*pb.MarketPair{
			{
				Pair: "BTC-USDT",
			},
		},
	}, nil
}

func (s *OrderService) GetOrderPair(ctx context.Context, req *pb.GetOrderPairRequest) (*pb.GetOrderPairResponse, error) {
	return &pb.GetOrderPairResponse{}, nil
}

func (s *OrderService) GetOrderPairChart(ctx context.Context, req *pb.GetOrderPairChartRequest) (*pb.GetOrderPairChartResponse, error) {
	return &pb.GetOrderPairChartResponse{}, nil
}

func (s *OrderService) GetOrderHistory(ctx context.Context, req *pb.GetOrderHistoryRequest) (*pb.GetOrderHistoryResponse, error) {
	return &pb.GetOrderHistoryResponse{}, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return &pb.CreateOrderResponse{}, nil
}

func (s *OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	return &pb.CancelOrderResponse{}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	return &pb.UpdateOrderResponse{}, nil
}

func (s *OrderService) StreamOrderUpdates(req *pb.StreamOrderUpdatesRequest, stream pb.OrderService_StreamOrderUpdatesServer) error {
	return nil
}

func (s *OrderService) StreamOrderBook(req *pb.StreamOrderBookRequest, stream pb.OrderService_StreamOrderBookServer) error {
	pair := req.Pair
	depth := req.Depth
	if depth == 0 {
		depth = 5
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			response := &pb.StreamOrderBookResponse{
				Pair: pair,
				Asks: []*pb.OrderBookEntry{
					{Price: 31000.5, Quantity: 0.1},
					{Price: 31010.0, Quantity: 0.2},
				},
				Bids: []*pb.OrderBookEntry{
					{Price: 30990.0, Quantity: 0.15},
					{Price: 30985.5, Quantity: 0.3},
				},
			}

			if err := stream.Send(response); err != nil {
				return fmt.Errorf("failed to stream order book: %w", err)
			}

		case <-stream.Context().Done():
			log.Printf("stream closed for pair %s", pair)
			return nil
		}
	}
}
