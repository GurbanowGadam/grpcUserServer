package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/GurbanowGadam/grpcUserServer"
	"google.golang.org/grpc"
)

func main() {
	host:="127.0.0.1"
	port:="7070"
	address := host+":"+port	
	
	// creds, err := credentials.NewClientTLSFromFile("/path/server.crt", "")
	// if err != nil {
	// 	log.Fatalf("Failed to load server certificate: %v", err)
	// }

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	// conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("grpc.Dial => ", err)
	}
	defer conn.Close()

	client := grpcUserServer.NewUserServerServiceClient(conn)

	for i := 0; i < 5; i++ {
		fmt.Println("I => ", i)
		err := GetProfileData(client, i)
		if err != nil {
			fmt.Println("ERROR => ", i, err.Error())
		}

		err = GetBlockedLists(client, i)
		if err != nil {
			fmt.Println("ERROR => ", i, err.Error())
		}
	}
	fmt.Println("-----------------------------------------------------")
	err = GetUsers(client, 10)
	if err != nil {
		fmt.Println("ERROR => ", 10, err.Error())
	}
	
}

func GetProfileData(client grpcUserServer.UserServerServiceClient, i int) error {
	ctx := context.Background()

	resp, err := client.GetProfileData(ctx, &grpcUserServer.GetProfileDataRequest{UserID: strconv.Itoa(i)})
	if err != nil {
		// fmt.Println("client.GetProfileData => ", err)
		return err
	}
	fmt.Println("resp => ", resp)	
	return nil
}


func GetBlockedLists(client grpcUserServer.UserServerServiceClient, i int) error {
	ctx := context.Background()

	resp, err := client.GetBlockedLists(ctx, &grpcUserServer.GetBlockedListsRequest{UserID: strconv.Itoa(i)})
	if err != nil {
		// fmt.Println("client.GetProfileData => ", err)
		return err
	}
	fmt.Println("resp => ", resp)	
	return nil
}

func GetUsers(client grpcUserServer.UserServerServiceClient, i int) error {
	ctx := context.Background()

	resp, err := client.GetUsers(ctx, &grpcUserServer.GetUsersRequest{LastUser: strconv.Itoa(i)})
	if err != nil {
		// fmt.Println("client.GetProfileData => ", err)
		return err
	}
	fmt.Println("resp => ", resp)	
	return nil
}
