package user

import (
	"errors"

	"github.com/bytedance/sonic"

	"ai_writer/internal/interactor/pkg/util"

	userModel "ai_writer/internal/interactor/models/users"
	userService "ai_writer/internal/interactor/service/user"

	// userRoleService "ai_writer/internal/interactor/service/user_role"

	"gorm.io/gorm"

	"ai_writer/internal/interactor/pkg/util/code"
	"ai_writer/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *userModel.Create) (int, any)
	GetByList(input *userModel.Fields) (int, any)
	GetBySingle(input *userModel.Field) (int, any)
	Delete(trx *gorm.DB, input *userModel.Field) (int, any)
	Update(trx *gorm.DB, input *userModel.Update) (int, any)
}

type manager struct {
	UserService userService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		UserService: userService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *userModel.Create) (int, any) {
	defer trx.Rollback()

	// 判斷使用者ID是否重複
	quantity, _ := m.UserService.GetByQuantity(&userModel.Field{
		UserName: util.PointerString(input.UserName),
	})

	if quantity > 0 {
		log.Info("UserName already exists. UserName: ", input.UserName)
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "User already exists.")
	}

	userBase, err := m.UserService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, userBase.ID)
}

func (m *manager) GetByList(input *userModel.Fields) (int, any) {
	output := &userModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, userBase, err := m.UserService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	userByte, err := sonic.Marshal(userBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = sonic.Unmarshal(userByte, &output.Users)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *userModel.Field) (int, any) {
	userBase, err := m.UserService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &userModel.Single{}
	userByte, _ := sonic.Marshal(userBase)
	err = sonic.Unmarshal(userByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(trx *gorm.DB, input *userModel.Field) (int, any) {
	defer trx.Rollback()

	_, err := m.UserService.GetBySingle(&userModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.UserService.WithTrx(trx).Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *userModel.Update) (int, any) {
	defer trx.Rollback()

	userBase, err := m.UserService.GetBySingle(&userModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.UserService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, userBase.ID)
}
