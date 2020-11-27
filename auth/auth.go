package auth

import (
	"context"
)

// AuthArgs 请求参数
type AuthArgs struct {
}

// AuthReply 响应数据
type AuthReply struct {
	Token string
}

type Auth struct {
}

// GetToken 获取token
func (t *Auth) GetToken(ctx context.Context, args *AuthArgs, reply *AuthReply) error {
	token, err := generateToken()
	if err != nil {
		return err
	}
	reply.Token = token
	return nil
}
