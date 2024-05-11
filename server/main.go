package main

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/GurbanowGadam/grpcUserServer"
	"google.golang.org/grpc"
)

func main() {
	host:="127.0.0.1"
	port:="7070"
	address := host+":"+port

	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
	}

	// creds, err := credentials.NewServerTLSFromFile("/home/gadamgurban/learn/grpcUserServer/ssl/hazarlogistik.crt", "/home/gadamgurban/learn/grpcUserServer/ssl/private.key")
	// if err != nil {
	// 	log.Fatalf("Failed to load server certificate: %v", err)
	// }

	// srv := grpc.NewServer(grpc.Creds(creds))
	srv := grpc.NewServer()
	grpcUserServer.RegisterUserServerServiceServer(srv, &myUserServerService{})

	fmt.Println("Starting server...", address)
	panic(srv.Serve(lis))
}

type myUserServerService struct {
	grpcUserServer.UnimplementedUserServerServiceServer
}

func (m *myUserServerService) GetProfileData(ctx context.Context, req *grpcUserServer.GetProfileDataRequest) (*grpcUserServer.GetProfileDataResponse, error) {
	var profileData grpcUserServer.GetProfileDataResponse

	userID := req.UserID
	fmt.Println(userID)
	profileData.UserID = userID
	profileData.FullName = "GadamGurban"
	profileData.AvatarName = "/path/avatar.jpg"
	profileData.Type = "men"

	return &profileData, nil
}


func (m *myUserServerService) GetBlockedLists(ctx context.Context, req *grpcUserServer.GetBlockedListsRequest) (*grpcUserServer.GetBlockedListsResponse, error) {
	var blockedList grpcUserServer.GetBlockedListsResponse

	userID := req.UserID
	fmt.Println(userID)
	blockedList.UserID = userID
	blockedList.FullName = "ElonMusk"

	return &blockedList, nil
}

func (m *myUserServerService) GetUsers(ctx context.Context, req *grpcUserServer.GetUsersRequest) (*grpcUserServer.GetUsersResponse, error) {
	var (
		usersList grpcUserServer.GetUsersResponse
		user grpcUserServer.GetUserResponse 
	)

	lastUserID := req.LastUser
	fmt.Println("lastUserID => ", lastUserID)
	for i := 0; i < 3; i++ {
		user.FullName = "Aman"
		user.UserID = strconv.Itoa(i)
		usersList.Users = append(usersList.Users, &user)
	}

	return &usersList, nil
}
