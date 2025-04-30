package main

import (
	"fmt"
	"os"

	"github.com/mdDocGenerator/model"
	"github.com/mdDocGenerator/pkg"
)

func main() {
	template := model.APIDocTemplate{
		Title:       "重設行程",
		Description: "",
		Path:        "/api/fleet/resetSchedule",
		Protocol:    "HTTP",
		Method:      "POST",
		Headers: []model.Header{
			{Key: "Accept", Value: "application/json"},
			{Key: "Token", Value: "JWT"},
		},

	}

	// Create a new markdown document
	doc := pkg.NewDoc()

	// Convert APIDocTemplate to APIDoc
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

		RequestExample: `{
	"uid": "",
	"deliveryId": ["9a23da1b-97eb-48d9-afd3-2f6116f10cf9"],
	"comment": "重設原因",
	"base64": {
		"file": "base 64 encode string",
		"filename": "name.jpg"
	}
}`,
		Response: model.DocField{
            Name: "response",
            Type: "object",
            Fields: []model.DocField{
                {
                    Name:        "success",
                    Required:    false,
                    Type:        "boolean",
                    Description: "是否成功",
                    Note: "",
                },
                {
                    Name:        "message",
                    Required:    false,
                    Type:        "string",
                    Description: "回應訊息",
                    Note: "",
                },
            },
        },
        ResponseExample: `{
    "success": true,
    "message": "行程重設成功"
}`,
	}

	// Generate API documentation
	doc.GenerateAPIDoc(apiDoc)

	// Export the document
	err := doc.Export("output/api_doc.md")
	if err != nil {
		fmt.Printf("Failed to export document: %v\n", err)
		os.Exit(1)
	}
}
