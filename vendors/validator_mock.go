package vendors

import (
	"github.com/stretchr/testify/mock"
)

type ModelValidatorMock struct {
	mock.Mock
}

func (v *ModelValidatorMock) Struct(doc interface{}) []error {
	args := v.Called(doc)
	err := args.Error(0)
	if err != nil {
		return []error{err}
	}
	return nil
}

func NewModelValidatorMock() ModelValidatorMock {
	return ModelValidatorMock{}
}
