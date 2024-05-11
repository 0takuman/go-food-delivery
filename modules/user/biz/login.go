package userbiz

import (
	"context"
	appctx "food-delivery/components/appcontext"
	"food-delivery/components/tokenprovider"
	usermodel "food-delivery/modules/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBussiness struct {
	appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(loginStorage LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBussiness {
	return &loginBussiness{
		storeUser:     loginStorage,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBussiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	return accessToken, nil
}
