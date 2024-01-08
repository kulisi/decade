package decade

import (
	"github.com/kardianos/service"
	"github.com/urfave/cli/v2"
)

// 定义一个卡片的通用接口

type ICard interface {
	Name() string
	Description() string
	Version() string

	StartAction() cli.ActionFunc
	GenericAction() cli.ActionFunc

	CardIn()
	CardOut()

	Start(s service.Service) error
	Stop(s service.Service) error
}
