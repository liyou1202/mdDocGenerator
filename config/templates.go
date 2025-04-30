package config

import "github.com/mdDocGenerator/model"

// APITemplates 定義了所有 API 文檔的模板
var APITemplates = map[string]model.APIDocTemplate{
	"GET": {
		Title:       "取得詳情",
		Description: "描述",
		Path:        "/api/admin/getSomthing",
		Protocol:    "HTTP",
		Method:      "GET",
		Headers: []model.Header{
			{Key: "Accept", Value: "application/json"},
			{Key: "Token", Value: "JWT"},
		},
	},
	"POST": {
		Title:       "建立詳情",
		Description: "描述",
		Path:        "/api/admin/postSomething",
		Protocol:    "HTTP",
		Method:      "POST",
		Headers: []model.Header{
			{Key: "Accept", Value: "application/json"},
			{Key: "Token", Value: "JWT"},
		},
	},
	"DELETE": {
		Title:       "刪除詳情",
		Description: "描述",
		Path:        "/api/admin/deleteSomething",
		Protocol:    "HTTP",
		Method:      "DELETE",
		Headers: []model.Header{
			{Key: "Accept", Value: "application/json"},
			{Key: "Token", Value: "JWT"},
		},
	},
}

// RequestExamples 定義了所有 API 請求的範例
var RequestExamples = map[string]string{
	"resetSchedule": `{
	"uid": "",
	"deliveryId": ["9a23da1b-97eb-48d9-afd3-2f6116f10cf9"],
	"comment": "重設原因",
	"base64": {
		"file": "base 64 encode string",
		"filename": "name.jpg"
	}
}`,
}

// ResponseExamples 定義了所有 API 回應的範例
var ResponseExamples = map[string]string{
	"resetSchedule": `{
    "success": true,
    "message": "行程重設成功"
}`,
}
