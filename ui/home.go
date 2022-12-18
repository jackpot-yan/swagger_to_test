package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/http"
	"swagger_to_test/models"
)

func Home() {
	page := app.New()
	var awType models.AnalysisArgs
	window := page.NewWindow("swagger接口方法自动封装")
	urlInput := widget.NewEntry()
	urlInput.SetPlaceHolder("请输入swagger_json url")
	content := container.NewVBox(urlInput, widget.NewButton("测试swagger链接", func() {
		resp, err := http.Get(urlInput.Text)
		awType.Url = urlInput.Text
		if err != nil {
			info := DialogInfo(&window, "测试链接失败", err.Error())
			info.Resize(fyne.NewSize(100, 200))
			info.Show()
		} else {
			if resp.StatusCode != 200 {
				info := DialogInfo(&window, "状态码不为200,请检查", fmt.Sprintf("状态码为%d", resp.StatusCode))
				info.Resize(fyne.NewSize(100, 200))
				info.Show()
			} else {
				info := DialogInfo(&window, "测试成功！", "您现在可以生成对应请求方法了！")
				info.Resize(fyne.NewSize(100, 200))
				info.Show()
			}
		}
	}))
	options := []string{"Python", "JavaScript", "Golang"}
	provinceSelect := widget.NewSelect(options, func(s string) {
		awType.AwType = s
	})
	content.Add(provinceSelect)
	window.SetContent(content)
	window.Resize(fyne.NewSize(450, 500))
	window.ShowAndRun()
}
