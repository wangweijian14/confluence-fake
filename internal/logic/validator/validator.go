package validatorI

import (
	"confluence_fake/internal/service"
	"fmt"
)

type sValidatorI struct {
}

func init() {
	service.RegisterValidatorI(New())
}
func New() *sValidatorI {
	return &sValidatorI{}
}

func (s *sValidatorI) Ok() {
	fmt.Println("ok")
}
