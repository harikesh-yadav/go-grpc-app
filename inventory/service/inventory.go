package service

import (
	"context"
	"log"

	pb "github.com/harikesh-yadav/go-grpc-app/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
}

func (s *InventoryServer) GetItemByName(ctx context.Context, wr *wrapperspb.StringValue) (*pb.Items, error) {
	log.Printf("GetItemByName called for value %s \n", wr.GetValue())
	items := pb.Items{
		Items: []*pb.Item{
			{Id: 1001,
				Name:        "Harikesh Yadav",
				Description: "test 1",
			},
			{Id: 1002,
				Name:        "Harikesh Kumar",
				Description: "test 2",
			},
		},
	}
	return &items, nil
}
