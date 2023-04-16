package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	desc "github.com/Constantine27K/crnt-auth-service/pkg/api/auth"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	authType = "Bearer"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	request := req.GetSecret()
	if request == nil {
		log.Error("secret is empty")
		return nil, status.Error(codes.InvalidArgument, "request is empty")
	}

	secret, err := i.secretStorage.GetByLogin(request.GetLogin())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("no such user",
				zap.String("login", request.GetLogin()),
			)
			return nil, status.Error(codes.InvalidArgument, "no such user")
		}
		log.Error("cannot get secret by login",
			zap.Error(err),
		)
		return nil, err
	}

	user, err := i.userService.GetUserByLogin(ctx, req.GetSecret().GetLogin())
	if err != nil {
		log.Error("failed to get user by login",
			zap.Error(err),
		)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(secret.GetPassword()), []byte(request.GetPassword()))
	if err != nil {
		log.Error("password is incorrect")
		return nil, status.Error(codes.InvalidArgument, "password is incorrect")
	}

	duration, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_DURATION_HOURS"), 10, 64)
	if err != nil {
		log.Error("failed to parse duration",
			zap.Error(err),
		)
		return nil, err
	}

	tokenValue, err := i.tokenMaker.CreateToken(secret.GetLogin(), secret.GetRole(), user.GetTeam(), time.Duration(duration)*time.Hour)
	if err != nil {
		log.Error("failed to generate token",
			zap.Error(err),
		)
		return nil, err
	}

	accessToken := fmt.Sprintf("%s %s", authType, tokenValue)

	return &desc.LoginResponse{AccessToken: accessToken}, nil
}
