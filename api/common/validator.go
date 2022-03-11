package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator[T any] struct {
}

func (v *Validator[T]) Validate(c *gin.Context, item *T) error {
	if err := c.ShouldBindJSON(item); err != nil {
		return err
	}
	err := validator.New().StructCtx(c.Request.Context(), item)
	if err != nil {
		return err
	}
	return nil
}
