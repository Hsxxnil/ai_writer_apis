package article

import (
	store "ai_writer/internal/entity/postgresql/article"
	db "ai_writer/internal/entity/postgresql/db/articles"
	model "ai_writer/internal/interactor/models/articles"

	"github.com/bytedance/sonic"

	"ai_writer/internal/interactor/pkg/util"
	"ai_writer/internal/interactor/pkg/util/log"
	"ai_writer/internal/interactor/pkg/util/uuid"

	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Create(input *model.Create) (output *db.Base, err error)
	GetByList(input *model.Fields) (quantity int64, output []*db.Base, err error)
	GetBySingle(input *model.Field) (output *db.Base, err error)
	GetByQuantity(input *model.Field) (quantity int64, err error)
	Update(input *model.Update) (err error)
	Delete(input *model.Update) (err error)
}

type service struct {
	Repository store.Entity
}

func Init(db *gorm.DB) Service {
	return &service{
		Repository: store.Init(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {
	return &service{
		Repository: s.Repository.WithTrx(tx),
	}
}

func (s *service) Create(input *model.Create) (output *db.Base, err error) {
	base := &db.Base{}
	marshal, err := sonic.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = sonic.Unmarshal(marshal, &base)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	base.ID = util.PointerString(uuid.CreatedUUIDString())
	base.CreatedAt = util.PointerTime(util.NowToUTC())
	base.UpdatedAt = util.PointerTime(util.NowToUTC())

	err = s.Repository.Create(base)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err = sonic.Marshal(base)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = sonic.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByList(input *model.Fields) (quantity int64, output []*db.Base, err error) {
	field := &db.Base{}
	marshal, err := sonic.Marshal(input)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	err = sonic.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	quantity, fields, err := s.Repository.GetByList(field)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	marshal, err = sonic.Marshal(fields)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	err = sonic.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *service) GetBySingle(input *model.Field) (output *db.Base, err error) {
	field := &db.Base{}
	marshal, err := sonic.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = sonic.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	single, err := s.Repository.GetBySingle(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err = sonic.Marshal(single)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = sonic.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) Delete(input *model.Update) (err error) {
	field := &db.Base{}
	marshal, err := sonic.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	err = sonic.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Repository.Delete(field)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *service) Update(input *model.Update) (err error) {
	field := &db.Base{}
	marshal, err := sonic.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	err = sonic.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Repository.Update(field)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *service) GetByQuantity(input *model.Field) (quantity int64, err error) {
	field := &db.Base{}
	marshal, err := sonic.Marshal(input)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	err = sonic.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	quantity, err = s.Repository.GetByQuantity(field)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return quantity, nil
}
