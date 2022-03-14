package endpoint

import (
	"context"

	"github.com/Fingann/Nibble/models"
	"github.com/Fingann/Nibble/services"
	trans "github.com/Fingann/Nibble/transport"

	"gorm.io/gorm"
)

func MakeRegistrationEndpoint(userService services.UserService) Endpoint[trans.RegistrationRequest, trans.RegistrationResponse] {
	return func(ctx context.Context, request trans.RegistrationRequest) (trans.RegistrationResponse, error) {

		id, err := userService.Register(request.Email, request.Username, request.Password)
		if err != nil {
			return trans.RegistrationResponse{}, err
		}

		return trans.RegistrationResponse{ID: id}, nil
	}
}

func MakeLoginEndpoint(userService services.UserService, jwtService services.JWTService) Endpoint[trans.LoginRequest, trans.LoginResponse] {
	return func(ctx context.Context, request trans.LoginRequest) (trans.LoginResponse, error) {
		exists, err := userService.Login(request.Username, request.Password)
		if err != nil {
			return trans.LoginResponse{}, err
		}
		if exists {
			token, err := jwtService.GenerateToken(request.Username, true)
			if err != nil {
				return trans.LoginResponse{}, err
			}
			return trans.LoginResponse{Token: token}, nil
		}
		return trans.LoginResponse{Error: "WRONG CREDENTIALS"}, nil
	}
}

func MakeUserRetrieveEndpoint(userService services.UserService) Endpoint[trans.UserRetrieveRequest, trans.UserRetrieveResponse] {
	return func(ctx context.Context, request trans.UserRetrieveRequest) (trans.UserRetrieveResponse, error) {
		user := models.User{Model: gorm.Model{ID: request.Id}}
		result, err := userService.First(user)
		if err != nil {
			return trans.UserRetrieveResponse{}, err
		}
		return trans.UserRetrieveResponse{
			ID:       result.ID,
			Username: result.Username,
			Email:    result.Email,
		}, nil
	}
}

func MakeUserUpdateEndpoint(userService services.UserService) Endpoint[trans.UserUpdateRequest, trans.UserUpdateResponse] {
	return func(c context.Context, request trans.UserUpdateRequest) (trans.UserUpdateResponse, error) {
		user := models.User{
			Model:    gorm.Model{ID: request.Id},
			Username: request.Username,
			Email:    request.Email,
		}
		result, err := userService.Update(user)
		if err != nil {
			return trans.UserUpdateResponse{}, err
		}
		return trans.UserUpdateResponse{ID: result.ID,
			Username: result.Email,
			Email:    result.Email,
		}, nil
	}
}

func MakeUserDeleteEndpoint(userService services.UserService) Endpoint[trans.UserDeleteRequest, trans.UserDeleteResponse] {
	return func(c context.Context, request trans.UserDeleteRequest) (trans.UserDeleteResponse, error) {
		err := userService.Delete(request.ID)
		if err != nil {
			return trans.UserDeleteResponse{}, err
		}
		return trans.UserDeleteResponse{ID: request.ID}, nil
	}
}
