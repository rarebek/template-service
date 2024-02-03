package repo

import (
	pb "template-service3/genproto/user_service"
)

//rpc Create(User) returns (User);
//rpc GetUserById(GetUserId) returns (UserWithPostsAndComments);
//rpc UpdateUser(User) returns (User);
//rpc DeleteUser(GetUserId) returns (User);
//rpc GetAllUsers(GetAllUsersRequest) returns (AllUsers);

// UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	GetUserById(*pb.GetUserId) (*pb.UserWithPostsAndComments, error)
	UpdateUser(*pb.User) (*pb.User, error)
	DeleteUser(*pb.GetUserId) (*pb.User, error)
	GetAllUsers(*pb.GetAllUsersRequest) (*pb.AllUsers, error)
}
