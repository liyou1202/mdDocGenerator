package main

import (
	"fmt"
	"os"

	"github.com/mdDocGenerator/config"
	"github.com/mdDocGenerator/model"
	"github.com/mdDocGenerator/pkg"
)

func main() {
	// 從配置中獲取 API 文檔模板
	template := config.APITemplates["GET"]

	// 創建新的 markdown 文檔
	doc := pkg.NewDoc()

	// 構建 API 文檔結構
	apiDoc := model.APIDoc{
		Title:       template.Title,
		Description: template.Description,
		Path:        template.Path,
		Protocol:    template.Protocol,
		Method:      template.Method,
		Headers:     template.Headers,
		Request: model.DocField{
			Name: "request",
			Type: "object",
			Fields: []model.DocField{
				{
					Name:        "uid",
					Required:    false,
					Type:        "string",
					Description: "使用者ID",
					Note:        "",
				},
				{
					Name:        "deliveryId",
					Required:    false,
					Type:        "array",
					Description: "配送ID列表",
					Note:        "",
				},
				{
					Name:        "comment",
					Required:    false,
					Type:        "string",
					Description: "重設原因",
					Note:        "",
				},
				{
					Name: "base64",
					Type: "object",
					Fields: []model.DocField{
						{
							Name:        "file",
							Required:    true,
							Type:        "string",
							Description: "base64 編碼字串",
							Note:        "",
						},
						{
							Name:        "filename",
							Required:    true,
							Type:        "string",
							Description: "檔案名稱",
							Note:        "",
						},
					},
				},
			},
		},
		RequestExample:  config.RequestExamples["resetSchedule"],
		Response: model.DocField{
			Name: "response",
			Type: "object",
			Fields: []model.DocField{
				{
					Name:        "success",
					Required:    false,
					Type:        "boolean",
					Description: "是否成功",
					Note:        "",
				},
				{
					Name:        "message",
					Required:    false,
					Type:        "string",
					Description: "回應訊息",
					Note:        "",
				},
			},
		},
		ResponseExample: config.ResponseExamples["resetSchedule"],
	}

	// 生成 API 文檔
	doc.GenerateAPIDoc(apiDoc)

	// 導出文檔
	if err := doc.Export("output/api_doc.md"); err != nil {
		fmt.Printf("導出文檔失敗: %v\n", err)
		os.Exit(1)
	}
}
