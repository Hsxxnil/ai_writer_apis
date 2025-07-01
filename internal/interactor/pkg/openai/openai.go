package openai

import (
	"ai_writer/internal/interactor/pkg/util/log"
	"context"
	"math"

	"github.com/bytedance/sonic"
	"github.com/sashabaranov/go-openai"
)

// config your openai api key and organization id
const (
	openApiKey     = "REMOVED"
	organizationID = "REMOVED"
)

// UploadFileRequest is a used to upload file to openai
type UploadFileRequest struct {
	// the name of the uploaded file in OpenAI
	FileName string
	// the bytes of the file
	File []byte
	// the purpose of the file
	Purpose openai.PurposeType
}

// FileResults is a used to store the file results
type FileResults struct {
	// the name of the uploaded file in OpenAI
	FileName string `json:"filename"`
	// the id of the file
	FileID string `json:"id"`
	// Status of the file
	Status string `json:"status"`
	// the purpose of the file
	Purpose string `json:"purpose"`
}

// FineTuningRequest is a used to fine-tuning the GPT model
type FineTuningRequest struct {
	// the name of the training file
	TrainingFileName *string
	// the bytes of the training file
	TrainingFile []byte
	// the name of the validation file
	ValidationFileName *string
	// the bytes of the validation file
	ValidationFile []byte
	// the model to fine-tune
	Model *string
}

// Chat is a used to chat with GPT
func Chat(input string, limit int) (output string, err error) {
	// calculate the token
	var fli float64 = float64(limit)
	token := int(math.Floor((fli * 1.5) + 0.5))

	// create the client
	client := openai.NewClient(openApiKey)

	// configure the request
	req := openai.ChatCompletionRequest{
		MaxTokens: token,
		Model:     openai.GPT3Dot5Turbo1106,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "你是一個產品行銷文章專家",
			},
		},
	}
	log.Info("Thinking start...")

	// add the user input to the request
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: input,
	})

	// send the request
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Error(err)
		return "", err
	}
	req.Messages = append(req.Messages, resp.Choices[0].Message)

	return req.Messages[len(req.Messages)-1].Content, nil
}

// FineTuning is a used to fine-tuning the GPT model
func FineTuning(input FineTuningRequest) error {
	// create the client
	client := openai.NewClient(openApiKey)

	// upload files
	var trainingFile, validationFile string
	if input.TrainingFileName != nil && input.TrainingFile != nil {
		// config the file request
		file := UploadFileRequest{
			FileName: *input.TrainingFileName,
			File:     input.TrainingFile,
			Purpose:  openai.PurposeFineTune,
		}

		// upload the file
		result, err := UploadFileToOpenAi(file)
		if err != nil {
			log.Error(err)
			return err
		}

		trainingFile = result.FileID
	}

	if input.ValidationFileName != nil && input.ValidationFile != nil {
		// config the file request
		file := UploadFileRequest{
			FileName: *input.ValidationFileName,
			File:     input.ValidationFile,
			Purpose:  openai.PurposeFineTune,
		}

		// upload the file
		result, err := UploadFileToOpenAi(file)
		if err != nil {
			log.Error(err)
			return err
		}

		validationFile = result.FileID
	}

	// configure the request for fine-tuning
	req := openai.FineTuningJobRequest{
		Model:          *input.Model,
		TrainingFile:   trainingFile,
		ValidationFile: validationFile,
	}

	// send the request
	resp, err := client.CreateFineTuningJob(context.Background(), req)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug(resp)

	return nil
}

// UploadFileToOpenAi is a used to upload file to openai
func UploadFileToOpenAi(input UploadFileRequest) (*FileResults, error) {
	output := &FileResults{}
	// create the client
	client := openai.NewClient(openApiKey)

	// configure the request
	req := openai.FileBytesRequest{
		Name:    input.FileName,
		Bytes:   input.File,
		Purpose: input.Purpose,
	}

	// send the request
	result, err := client.CreateFileBytes(context.Background(), req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultBytes, err := sonic.Marshal(result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = sonic.Unmarshal(resultBytes, output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}
