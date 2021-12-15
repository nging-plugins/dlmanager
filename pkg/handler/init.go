package download

import (
	"github.com/admpub/godownloader/service"
	"github.com/webx-top/echo"
	mw "github.com/webx-top/echo/middleware"

	"github.com/admpub/nging/v4/application/handler"
	"github.com/admpub/nging/v4/application/library/config"

	"github.com/nging-plugins/dbmanager/pkg/library/dbmanager/driver/mysql"
)

var downloadDir = func() string {
	if len(config.DefaultConfig.Download.SavePath) == 0 {
		return service.GetDownloadPath()
	}
	return config.DefaultConfig.Download.SavePath
}

func init() {
	server := &service.DServ{}
	server.SetTmpl(`download/index`)
	server.SetSavePath(downloadDir)
	mysql.SQLTempDir = downloadDir //将SQL文件缓存到下载目录里面方便管理
	handler.RegisterToGroup(`/download`, func(g echo.RouteRegister) {
		server.Register(g, true)
		g.Route(`GET,POST`, `/file`, File, mw.CORS())
	})
}
