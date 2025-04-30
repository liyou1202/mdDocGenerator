package pkg

import (
	"fmt"
	"github.com/mdDocGenerator/model"
	"strings"
)

// generateFieldsDoc 遞迴生成欄位文檔
func (doc *MarkdownDoc) generateFieldsDoc(field model.DocField, level int) string {
	indent := strings.Repeat("  ", level)
	content := fmt.Sprintf("%s- %s (%s): %s\n", indent, field.Name, field.Type, field.Description)

	for _, subField := range field.Fields {
		content += doc.generateFieldsDoc(subField, level+1)
	}

	return content
}

// GenerateAPIDoc 生成 API 文檔
func (doc *MarkdownDoc) GenerateAPIDoc(apiDoc model.APIDoc) *MarkdownDoc {
	// 標題
	doc.AddTitle(apiDoc.Title, 3).
		AddBlankLines(1)

	// 基本資訊
	doc.AddContent("- PATH: `" + apiDoc.Path + "`").
		AddBlankLines(1)
	doc.AddContent("- Protocol: " + apiDoc.Protocol).
		AddBlankLines(1)
	doc.AddContent("- Method: " + apiDoc.Method).
		AddBlankLines(1)

	// Headers
	doc.AddContent("- Header:").
		AddBlankLines(1)
	for _, header := range apiDoc.Headers {
		doc.AddContent(fmt.Sprintf("  - `%s: %s`", header.Key, header.Value)).
			AddBlankLines(1)
	}
	doc.AddBlankLines(2)

	// 請求體
	if len(apiDoc.Request.Fields) > 0 {
		doc.AddTitle("Request", 4).
			AddBlankLines(1)

		// 請求體表格
		doc.AddContent("| name            | required | data type  | description                      | note |").
			AddBlankLines(1)
		doc.AddContent("| --------------- | -------- | ---------- | -------------------------------- | ---------- |").
			AddBlankLines(1)

		// 生成請求體欄位表格
		for _, field := range apiDoc.Request.Fields {
			doc.generateFieldRow(field, 0)
		}
		doc.AddBlankLines(1)

		// 如果有請求體範例
		if apiDoc.RequestExample != "" {
			doc.AddContent("```json").
				AddBlankLines(1)
			doc.AddContent(apiDoc.RequestExample).
				AddBlankLines(1)
			doc.AddContent("```").
				AddBlankLines(3)
		}
	}

	// 回應內容
	if len(apiDoc.Response.Fields) > 0 || apiDoc.Response.Type != "" {
		doc.AddTitle("Response Body", 4).
			AddBlankLines(1)

		// 回應內容表格
		doc.AddContent("| name            | required | data type  | description                      | note |").
			AddBlankLines(1)
		doc.AddContent("| --------------- | -------- | ---------- | -------------------------------- | ---------- |").
			AddBlankLines(1)
			
		// 生成回應內容欄位表格
		doc.generateFieldRow(apiDoc.Response, 0)
		doc.AddBlankLines(1)

		// 如果有回應範例
		if apiDoc.ResponseExample != "" {
			doc.AddContent("```json").
				AddBlankLines(1)
			doc.AddContent(apiDoc.ResponseExample).
				AddBlankLines(1)
			doc.AddContent("```").
				AddBlankLines(3)
		}
	}

	return doc
}

// generateFieldRow 生成欄位行
func (doc *MarkdownDoc) generateFieldRow(field model.DocField, depth int) {
    required := ""
    if field.Required {
        required = "Y"
    }

    // 如果是物件類型且有子欄位，創建新的表格
    if field.Type == "object" && len(field.Fields) > 0 {
        // 先添加當前欄位到主表格
        doc.AddContent(fmt.Sprintf("| %-15s | %-8s | %-10s | %-20s | %-10s |",
            field.Name,
            required,
            field.Type,
            field.Description,
            field.Note)).AddBlankLines(3)

        // 為子欄位創建新的表格
        doc.AddTitle(fmt.Sprintf("%s", field.Name), 5).
		AddBlankLines(1)

        // 添加子表格標題
        doc.AddContent("| name            | required | data type  | description                      | note |").
            AddBlankLines(1)
        doc.AddContent("| --------------- | -------- | ---------- | -------------------------------- | ---------- |").
            AddBlankLines(1)

        // 遞迴處理子欄位
        for _, subField := range field.Fields {
            doc.generateFieldRow(subField, 0)
        }
        doc.AddBlankLines(1)
    } else {
        // 一般欄位直接添加到當前表格
        doc.AddContent(fmt.Sprintf("| %-15s | %-8s | %-10s | %-20s | %-10s |",
            field.Name,
            required,
            field.Type,
            field.Description,
            field.Note)).AddBlankLines(1)
    }
}
