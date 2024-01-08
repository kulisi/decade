package decade

import (
	"log"

	"github.com/urfave/cli/v2"
)

// 封装 urfave/cli 为选项模式的 Decade
type Decade struct {
	cli *cli.App
}

// 调用第三方库 urfave/cli 的 Run 函数
func (d *Decade) Run(arguments []string) {
	err := d.cli.Run(arguments)
	if err != nil {
		log.Fatalln("Cli Run Error: " + err.Error())
	}
}

// 克隆 Decade
func (d *Decade) clone() *Decade {
	clone := *d
	return &clone
}

// 给 Decade 应用选项
func (d *Decade) WithOptions(opts ...Option) *Decade {
	decade := d.clone()
	for _, opt := range opts {
		opt.apply(decade)
	}
	return decade
}

// 构造函数
func New(card ICard, opts ...Option) *Decade {
	cmd := &cli.App{
		Name:    card.Name(),
		Usage:   card.Description(),
		Version: card.Version(),
		Commands: []*cli.Command{
			{Name: "install", Action: card.GenericAction(), Usage: "安装服务"},
			{Name: "uninstall", Action: card.GenericAction(), Usage: "卸载服务"},
			{Name: "start", Action: card.GenericAction(), Usage: "启动服务"},
			{Name: "stop", Action: card.GenericAction(), Usage: "停止服务"},
			{Name: "restart", Action: card.GenericAction(), Usage: "重启服务"},
		},
		Action: card.StartAction(),
	}

	decade := &Decade{cli: cmd}

	for _, opt := range opts {
		opt.apply(decade)
	}

	return decade
}
