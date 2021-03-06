package usecase_auth

import (
	"fmt"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	domain_service_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/users"
)

type IUserAuthUsecase interface {
	SignUp(*UserSignUpInputCommand) (*model.User, error)
}

type UserAuthUsecase struct {
	userRepository repository.UserRepository
	userService    domain_service_users.UserService
}

// NewAuthUsecase is create UserAuthUsecase Instance
func NewAuthUsecase(userRepository repository.UserRepository) IUserAuthUsecase {
	userService := domain_service_users.NewUserService(userRepository)
	userAuthUsecase := UserAuthUsecase{
		userRepository: userRepository,
		userService:    *userService,
	}

	return &userAuthUsecase
}

// SignUp is create user
func (usecase *UserAuthUsecase) SignUp(command *UserSignUpInputCommand) (*model.User, error) {
	// create object
	email, err := user.NewEmail(command.Email)
	uid := command.Uid
	firstName, err := user.NewFirstName("new user")
	lastName, err := user.NewLastName("new user")
	displayName, err := user.NewDisplayName("new user displayname")

	if err != nil {
		return nil, err
	}

	user := model.NewUser(email, uid, firstName, lastName, displayName)
	// Find user
	isExistUser := usecase.userService.EmailExists(user.Email)

	if isExistUser {
		return nil, fmt.Errorf("%s is already exists", command.Email)
	}

	usecase.userRepository.Save(user)

	return user, nil
}
