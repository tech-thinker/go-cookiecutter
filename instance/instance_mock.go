package instance

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
)

type InstanceMock struct {
	mock.Mock
}

func (instance *InstanceMock) Destroy() error {
	args := instance.Called()
	return args.Error(1)
}

func (instance *InstanceMock) DB() orm.Ormer {
	args := instance.Called()
	return args.Get(0).(orm.Ormer)
}

func (instance *InstanceMock) Validator() *validator.Validate {
	args := instance.Called()
	return args.Get(0).(*validator.Validate)
}

func NewInstanceMock() InstanceMock {
	return InstanceMock{}
}
