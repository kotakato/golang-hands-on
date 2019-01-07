package web

import (
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

// NewEchoCustomValidator はEchoCustomValidatorを作成してポインターを取得する。
func NewEchoCustomValidator() echo.Validator {
	return &EchoCustomValidator{validator: validator.New()}
}

// EchoCustomValidator はecho.Validatorのgopkg.in/go-playground/validatorによる実装。
type EchoCustomValidator struct {
	validator *validator.Validate
}

// Validate は引数を構造体のフィールドタグを使用してバリデーションする。
func (cv *EchoCustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
