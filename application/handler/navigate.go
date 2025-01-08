package handler

import (
	"github.com/coscms/webcore/library/navigate"
	"github.com/webx-top/echo"
)

var LeftNavigate = &navigate.Item{
	Display: true,
	Name:    echo.T(`离线下载`),
	Action:  `download`,
	Icon:    `download`,
	Children: &navigate.List{
		{
			Display: true,
			Name:    echo.T(`下载管理`),
			Action:  `index.html`,
		},
		{
			Display: true,
			Name:    echo.T(`文件管理`),
			Action:  `file`,
		},
		{
			Display: false,
			Name:    echo.T(`总进度信息`),
			Action:  `progress.json`,
		},
		{
			Display: false,
			Name:    echo.T(`添加任务`),
			Action:  `add_task`,
		},
		{
			Display: false,
			Name:    echo.T(`删除任务`),
			Action:  `remove_task`,
		},
		{
			Display: false,
			Name:    echo.T(`启动任务`),
			Action:  `start_task`,
		},
		{
			Display: false,
			Name:    echo.T(`停止任务`),
			Action:  `stop_task`,
		},
		{
			Display: false,
			Name:    echo.T(`重启任务`),
			Action:  `restart_task`,
		},
		{
			Display: false,
			Name:    echo.T(`启动所有任务`),
			Action:  `start_all_task`,
		},
		{
			Display: false,
			Name:    echo.T(`停止所有任务`),
			Action:  `stop_all_task`,
		},
		{
			Display: false,
			Name:    echo.T(`单个文件进度信息`),
			Action:  `progress/*`,
		},
	},
}
