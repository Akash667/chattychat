package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(ctx context.Context, user *UserCreateReq) (*UserCreateRes, error)
	LoginUser(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
}

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email"  db:"email"`
}

type UserCreateReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

type UserCreateRes struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type LoginUserReq struct {
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

type LoginUserRes struct {
	accessToken string
	Username    string `json:"username" db:"username"`
	ID          int64  `json:"id" db:"id"`
}
