package postgres

import (
	"template-service3/config"
	pb "template-service3/genproto/user_service"
	"template-service3/pkg/db"
	"template-service3/storage/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.UserStorageI
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	pgPool, cleanUp, _ := db.ConnectToDB(config.Load())
	s.Repository = NewUserRepo(pgPool)
	s.CleanupFunc = cleanUp
}

func (s *UserRepositoryTestSuite) TestUserCRUD() {
	user := &pb.User{
		FirstName: "Test User Name 1",
		LastName:  "Test User Last Name 1",
		Age:       20,
		Gender:    0,
	}

	createdUser, err := s.Repository.Create(user)
	s.Suite.NotNil(createdUser)
	s.Suite.NoError(err)
	s.Suite.Equal(user.FirstName, createdUser.FirstName)
	s.Suite.Equal(user.LastName, createdUser.LastName)
	s.Suite.Equal(user.Age, createdUser.Age)
	s.Suite.Equal(user.Gender, createdUser.Gender)

	userId := pb.GetUserId{UserId: createdUser.Id}
	getUser, err := s.Repository.GetUserById(&userId)
	s.Suite.NotNil(getUser)
	s.Suite.NoError(err)
	s.Suite.Equal(getUser.FirstName, createdUser.FirstName)
	s.Suite.Equal(getUser.LastName, createdUser.LastName)
	s.Suite.Equal(getUser.Age, createdUser.Age)
	s.Suite.Equal(getUser.Gender, createdUser.Gender)

	createdUser.FirstName = "Updated User Name 1"
	createdUser.LastName = "Updated User Last Name 1"

	updateUser, err := s.Repository.UpdateUser(createdUser)
	s.Suite.NotNil(updateUser)
	s.Suite.NoError(err)

	updatedUser, err := s.Repository.GetUserById(&userId)
	s.Suite.NotNil(getUser)
	s.Suite.NoError(err)
	s.Suite.Equal(updatedUser.FirstName, createdUser.FirstName)
	s.Suite.Equal(updatedUser.LastName, createdUser.LastName)
	s.Suite.Equal(updatedUser.Age, createdUser.Age)
	s.Suite.Equal(updatedUser.Gender, createdUser.Gender)

	users, err := s.Repository.GetAllUsers(&pb.GetAllUsersRequest{})
	s.Suite.NotNil(users)
	s.Suite.NoError(err)

	deletedUser, err := s.Repository.DeleteUser(&userId)
	s.Suite.NotNil(deletedUser)
	s.Suite.NoError(err)
	s.Suite.Equal(deletedUser.FirstName, createdUser.FirstName)
	s.Suite.Equal(deletedUser.LastName, createdUser.LastName)
	s.Suite.Equal(deletedUser.Age, createdUser.Age)
	s.Suite.Equal(deletedUser.Gender, createdUser.Gender)

}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	s.CleanupFunc()
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
