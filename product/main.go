package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/harikesh-yadav/go-grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var serverAdd = "0.0.0.0:8001"

func main() {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(serverAdd, opts...)

	if err != nil {
		log.Fatalf("Grpc server failed to connect with error %v", err)
	}
	defer conn.Close()
	client := pb.NewInventoryServiceClient(conn)

	route := gin.Default()

	route.GET("/item/:name", getHandler(client))

	route.Run("localhost:8000")
}

func getHandler(client pb.InventoryServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		items, _ := client.GetItemByName(context.Background(), wrapperspb.String(name))

		c.JSON(http.StatusOK, gin.H{
			"msg":  "OK",
			"data": *items,
		})
	}
}
