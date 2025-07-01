package article

import (
	articleModel "ai_writer/internal/interactor/models/articles"
	"ai_writer/internal/interactor/pkg/openai"
	"ai_writer/internal/interactor/pkg/util"
	"ai_writer/internal/interactor/pkg/util/code"
	"ai_writer/internal/interactor/pkg/util/log"
	articleService "ai_writer/internal/interactor/service/article"
	"errors"
	"fmt"
	"reflect"
	"slices"

	"github.com/bytedance/sonic"

	"gorm.io/gorm"
)

type Manager interface {
	Create(trx *gorm.DB, input *articleModel.Create) (int, any)
	GetByList(input *articleModel.Fields) (int, any)
	GetBySingle(input *articleModel.Field) (int, any)
	Delete(input *articleModel.Update) (int, any)
	Update(input *articleModel.Update) (int, any)
}

type manager struct {
	ArticleService articleService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ArticleService: articleService.Init(db),
	}
}

// checkEmptyFields is a helper function to check empty fields
func checkEmptyFields(fieldsNameMap map[string]string, input any) (result string) {

	// get the value of the input
	inputValue := reflect.ValueOf(input).Elem()

	// loop through the fields
	for i := 0; i < inputValue.NumField(); i++ {
		field := inputValue.Field(i)
		fieldName := inputValue.Type().Field(i).Name
		fieldDisplayName := fieldsNameMap[fieldName]

		isEmpty := func() bool {
			switch field.Kind() {
			case reflect.String:
				return field.String() == ""
			case reflect.Ptr, reflect.Interface:
				return field.IsNil()
			default:
				return false
			}
		}()

		if !isEmpty {
			if fieldDisplayName != "" {
				result += fmt.Sprintf("%s: %v\n", fieldDisplayName, field)
			}
		}
	}

	return result
}

func (m *manager) Create(trx *gorm.DB, input *articleModel.Create) (int, any) {
	defer trx.Rollback()

	// map the fields name
	fieldsNameMap := map[string]string{
		"KeyMessage":        "關鍵訊息",
		"Story":             "口碑切角",
		"ProductName":       "產品名稱",
		"ProductFeature":    "產品特性",
		"ProductHighlights": "產品亮點",
	}

	// transform the age
	slices.Sort(input.Age)
	input.Ages = fmt.Sprintf("%d~%d歲", input.Age[0], input.Age[1])
	input.FromAge = input.Age[0]
	input.ToAge = input.Age[1]

	// get the non-empty fields
	fields := checkEmptyFields(fieldsNameMap, input)

	// send request to openai
	result, err := openai.Chat(fmt.Sprintf("模擬%s%s%s版,依據性別%s,%s,%s人物設定,其他設定%s寫出%s類型,%s風格,在%d字內正體中文文章含標點符號", input.Forum, input.ContentType, input.Board, input.Gender, input.CharacterTrait, input.CharacterRemarks, fields, input.Type, input.Style, input.WordLimit), input.WordLimit)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err)
	}

	input.AIArticle = result

	// store the data
	articleBase, err := m.ArticleService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// transform the data to output
	output := articleModel.Single{}
	articleByte, err := sonic.Marshal(articleBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = sonic.Unmarshal(articleByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByList(input *articleModel.Fields) (int, any) {
	output := &articleModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, articleBase, err := m.ArticleService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	articleByte, err := sonic.Marshal(articleBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = sonic.Unmarshal(articleByte, &output.Articles)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *articleModel.Field) (int, any) {
	articleBase, err := m.ArticleService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &articleModel.Single{}
	articleByte, _ := sonic.Marshal(articleBase)
	err = sonic.Unmarshal(articleByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *articleModel.Update) (int, any) {
	_, err := m.ArticleService.GetBySingle(&articleModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ArticleService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *articleModel.Update) (int, any) {
	articleBase, err := m.ArticleService.GetBySingle(&articleModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ArticleService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, articleBase.ID)
}
