# decade

#### 例子

``` go 
package driver

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/urfave/cli/v2"
)

// 实现接口 IDecadeApplication
type ExampleCard struct {
	version string

	cfg service.Config

	start   cli.ActionFunc
	generic cli.ActionFunc
	// web 服务
	http http.Server
}

// 实现接口 IDecadeApplication 的 StartAction 函数
func (card *ExampleCard) StartAction() cli.ActionFunc {
	return card.start
}

// 实现接口 IDecadeApplication 的 GenericAction 函数
func (card *ExampleCard) GenericAction() cli.ActionFunc {
	return card.generic
}

// 实现接口 IDecadeApplication 的 Name 函数
func (card *ExampleCard) Name() string {
	return card.cfg.Name
}

// 实现接口 IDecadeApplication 的 Description 函数
func (card *ExampleCard) Description() string {
	return card.cfg.Description
}

// 实现接口 IDecadeApplication 的 Version 函数
func (card *ExampleCard) Version() string {
	return card.version
}

// 实现接口 IDecadeApplication 的 CardIn 函数
func (card *ExampleCard) CardIn() {
	if err := card.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server Start Error:", err)
	}
}

// 实现接口 IDecadeApplication 的 CardOut 函数
func (card *ExampleCard) CardOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := card.http.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
}

// 实现接口 IDecadeApplication 的 Start 函数
func (card *ExampleCard) Start(s service.Service) error {
	go card.CardIn()
	return nil
}

// 实现接口 IDecadeApplication 的 Stop 函数
func (card *ExampleCard) Stop(s service.Service) error {
	card.CardOut()
	return nil
}

func NewExampleCard() ICard {

	web := &ExampleCard{}
	web.cfg = service.Config{
		Name:        "Decade",
		DisplayName: "假面骑士Decade",
		Description: "是一个路过的打工仔",
	}

	web.version = "0.0.1"
	web.start = func(ctx *cli.Context) error {
		svc, _ := service.New(web, &web.cfg)
		svc.Run()
		return nil
	}
	web.generic = func(ctx *cli.Context) error {
		svc, _ := service.New(web, &web.cfg)
		service.Control(svc, ctx.Command.Name)
		return nil
	}

	gin.SetMode(gin.ReleaseMode)

	routes := gin.Default()

	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 50000,
			"time": time.Now().String(),
			"seed": time.Now().Unix(),
		})
	})

	web.http.Handler = routes

	web.http.Addr = ":8080"

	return web
}

```

#### 安装服务
```cmd
./decade.exe install
```

#### 卸载服务
```cmd
./decade.exe uninstall
```

#### 启动服务
```cmd
./decade.exe start
```

#### 停止服务
```cmd
./decade.exe stop
```

#### 重启服务
```cmd
./decade.exe restart
```