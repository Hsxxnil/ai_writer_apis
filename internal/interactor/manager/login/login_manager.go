package login

import (
	"errors"

	"github.com/bytedance/sonic"

	"ai_writer/config"

	jwxModel "ai_writer/internal/interactor/models/jwx"
	loginsModel "ai_writer/internal/interactor/models/logins"
	usersModel "ai_writer/internal/interactor/models/users"
	"ai_writer/internal/interactor/pkg/jwx"
	"ai_writer/internal/interactor/pkg/util"
	"ai_writer/internal/interactor/pkg/util/code"
	"ai_writer/internal/interactor/pkg/util/log"
	jwxService "ai_writer/internal/interactor/service/jwx"
	userService "ai_writer/internal/interactor/service/user"

	"gorm.io/gorm"
)

type Manager interface {
	Login(input *loginsModel.Login) (int, any)
	Refresh(input *jwxModel.Refresh) (int, any)
}

type manager struct {
	UserService userService.Service
	JwxService  jwxService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		UserService: userService.Init(db),
		JwxService:  jwxService.Init(),
	}
}

func (m *manager) Login(input *loginsModel.Login) (int, any) {

	// 驗證帳密
	acknowledge, userBase, err := m.UserService.AcknowledgeUser(&usersModel.Field{
		UserName: util.PointerString(input.UserName),
		Password: util.PointerString(input.Password),
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if acknowledge == false {
		return code.PermissionDenied, code.GetCodeMessage(code.PermissionDenied, "Incorrect user_name or password.")
	}

	// 產生accessToken
	output := &jwxModel.Token{}
	accessToken, err := m.JwxService.CreateAccessToken(&jwxModel.JWX{
		UserID: userBase.ID,
		Name:   userBase.Name,
	})

	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	accessTokenByte, _ := sonic.Marshal(accessToken)
	err = sonic.Unmarshal(accessTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 產生refreshToken
	refreshToken, err := m.JwxService.CreateRefreshToken(&jwxModel.JWX{
		UserID: userBase.ID,
		Name:   userBase.Name,
	})

	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	refreshTokenByte, _ := sonic.Marshal(refreshToken)
	err = sonic.Unmarshal(refreshTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Refresh(input *jwxModel.Refresh) (int, any) {
	// 驗證refreshToken
	j := &jwx.JWT{
		PublicKey: config.RefreshPublicKey,
		Token:     input.RefreshToken,
	}

	if len(input.RefreshToken) == 0 {
		return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is null.")
	}

	j, err := j.Verify()
	if err != nil {
		log.Error(err)
		return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is error.")
	}

	field, err := m.UserService.GetBySingle(&usersModel.Field{
		ID: j.Other["user_id"].(string),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is error.")
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is error.")
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 產生accessToken
	token, err := m.JwxService.CreateAccessToken(&jwxModel.JWX{
		UserID: field.ID,
		Name:   field.Name,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	token.RefreshToken = input.RefreshToken
	return code.Successful, code.GetCodeMessage(code.Successful, token)
}
