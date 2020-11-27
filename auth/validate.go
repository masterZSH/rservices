package auth

import (
	"context"
	"errors"
)

// ValidateArgs 请求参数
type ValidateArgs struct {
	Token string
}

// ValidateReply 响应数据
type ValidateReply struct {
}

// Validate validate token
type Validate int

var (
	// ErrInvalidToken invalid token error
	ErrInvalidToken = errors.New("invalid token")
)

// ValidateToken validate token
func (v *Validate) ValidateToken(ctx context.Context, args *ValidateArgs, reply *ValidateReply) error {
	err := validateToken(args.Token)
	if err != nil {
		return err
	}
	return nil
}
