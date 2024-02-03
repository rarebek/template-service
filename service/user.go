package service

import (
	"context"
	"fmt"
	"log"
	pbc "template-service3/genproto/comment_service"
	pbp "template-service3/genproto/post_service"
	pb "template-service3/genproto/user_service"
	l "template-service3/pkg/logger"
	grpcClient "template-service3/service/grpc_client"
	"template-service3/storage"

	"github.com/jmoiron/sqlx"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
	pb.UnimplementedUserServiceServer
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().Create(req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().UpdateUser(req)
}

func (s *UserService) DeleteUser(ctx context.Context, reqId *pb.GetUserId) (*pb.User, error) {
	return s.storage.User().DeleteUser(reqId)
}

func (s *UserService) GetUserById(ctx context.Context, reqId *pb.GetUserId) (*pb.UserWithPostsAndComments, error) {
	//posts -> post -> comments -> comment -> owner
	user, err := s.storage.User().GetUserById(reqId)
	if err != nil {
		log.Fatal("cannot get user by id", err.Error())
	}
	var posts []*pb.Post
	allPosts, err := s.client.PostService().GetPostsByOwnerId(ctx, &pbp.GetOwnerId{
		OwnerId: user.Id,
	})
	if err != nil {
		log.Fatal("cannot get posts of user", err.Error())
	}

	for _, post := range allPosts.Posts {
		comments, err := s.client.CommentService().GetAllCommentsByPostId(ctx, &pbc.GetPostID{
			PostId: post.Id,
		})
		if err != nil {
			log.Fatal("cannot get comments of a post", err.Error())
		}

		p := pb.Post{
			Id:       post.Id,
			Title:    post.Title,
			ImageUrl: post.ImageUrl,
			OwnerId:  post.OwnerId,
		}

		for _, comment := range comments.Comments {
			c := pb.Comment{
				Id:        comment.Id,
				Content:   comment.Content,
				PostId:    comment.PostId,
				OwnerId:   comment.OwnerId,
				CreatedAt: comment.CreatedAt,
				UpdatedAt: comment.UpdatedAt,
				DeletedAt: comment.DeletedAt,
			}
			user, err := s.storage.User().GetUserById(&pb.GetUserId{
				UserId: c.OwnerId,
			})

			if c.Owner == nil {
				c.Owner = &pb.User{}
			}

			if err != nil {
				log.Fatal("cannot get owner of the comment", err.Error())
			}
			c.Owner.Id = user.Id
			c.Owner.FirstName = user.FirstName
			c.Owner.LastName = user.LastName
			c.Owner.Age = user.Age
			c.Owner.Gender = user.Gender
			p.Comments = append(p.Comments, &c)
		}
		posts = append(posts, &p)
	}
	user.Posts = posts

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.AllUsers, error) {
	users, err := s.storage.User().GetAllUsers(req)
	if err != nil {
		log.Fatal("cannot get all users", err.Error())
	}
	for _, user := range users.Users {
		var posts []*pb.Post
		allPosts, err := s.client.PostService().GetPostsByOwnerId(ctx, &pbp.GetOwnerId{
			OwnerId: user.Id,
		})
		if err != nil {
			log.Fatal("cannot get posts of a user", err.Error())
		}

		for _, post := range allPosts.Posts {
			p := pb.Post{
				Id:       post.Id,
				Title:    post.Title,
				ImageUrl: post.ImageUrl,
				OwnerId:  post.OwnerId,
				//Comments: nil,
			}

			comments, err := s.client.CommentService().GetAllCommentsByPostId(ctx, &pbc.GetPostID{
				PostId: post.Id,
			})
			if err != nil {
				log.Fatal("cannot get comment of the post", err.Error())
			}
			var postComments []*pb.Comment
			for _, comment := range comments.Comments {
				c := pb.Comment{
					Id:        comment.Id,
					Content:   comment.Content,
					PostId:    comment.PostId,
					OwnerId:   comment.OwnerId,
					CreatedAt: comment.CreatedAt,
					UpdatedAt: comment.UpdatedAt,
					DeletedAt: comment.DeletedAt,
					//Owner:     nil,
				}
				owner, err := s.storage.User().GetUserById(&pb.GetUserId{
					UserId: c.OwnerId,
				})
				fmt.Println(owner, "fwefwe")
				if err != nil {
					log.Fatal("cannot get owner of the comment", err.Error())
				}

				if c.Owner == nil {
					c.Owner = &pb.User{}
				}

				c.Owner.Id = owner.Id
				c.Owner.FirstName = owner.FirstName
				c.Owner.LastName = owner.LastName
				c.Owner.Age = owner.Age
				c.Owner.Gender = owner.Gender

				postComments = append(postComments, &c)
			}
			p.Comments = postComments

			posts = append(posts, &p)
		}
		user.Posts = posts
	}

	return users, nil
}

func (s *UserService) mustEmbedUnimplementedUserServiceServer() {}
