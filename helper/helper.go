package helper

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	StatusUnprocessableEntity = 422
	StatusBadRequest          = 400
	StatusInternalServerError = 500
	StatusCreated             = 201
	StatusOK                  = 200
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, v := range validationErr {
			errors = append(errors, v.Error())
		}
	} else if syntaxErr, ok := err.(*json.SyntaxError); ok {
		errors = append(errors, syntaxErr.Error())
	} else {
		log.Printf("Unexpected error: %v", err)
		errors = append(errors, "Internal server error")
	}

	return errors
}

func GenerateUniqueID(length int) string {
	if length == 0 {
		length = 4
	}
	uniqCode := time.Now().Format("20060102150405")
	uniqCode = uniqCode[len(uniqCode)-length:]
	return uniqCode
}
