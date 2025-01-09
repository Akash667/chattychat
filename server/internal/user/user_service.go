package user

import (
	"context"
	"strconv"

	"github.com/akrawat667/baseChat/server/utils"
	"github.com/golang-jwt/jwt/v5"
)

const (
	secretkey = "bruh"
)

type service struct {
	Repository
}

func NewService(r Repository) Service {

	return &service{r}

}

func (s *service) CreateUser(ctx context.Context, user *UserCreateReq) (*UserCreateRes, error) {

	var u User

	u.Email = user.Email
	u.Username = user.Username
	u.Password = utils.HashPassword(user.Password)
	createdUser, err := s.Repository.CreateUser(ctx, &u)
	if err != nil {
		return &UserCreateRes{}, err
	}

	u.ID = createdUser.ID

	return &UserCreateRes{ID: u.ID, Username: u.Username, Email: u.Email}, nil

}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type WrongPassword struct{}

func (w WrongPassword) Error() string {
	return "incorrect Password"
}

func (s *service) LoginUser(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {

	user, err := s.Repository.GetUser(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return &LoginUserRes{}, &WrongPassword{}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(user.ID)),
		},
	})

	accessToken, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: accessToken, Username: user.Username, ID: user.ID}, nil

}
