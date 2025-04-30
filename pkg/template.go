package pkg

import (
	"fmt"

	"github.com/mdDocGenerator/model"
)

// GenerateFromTemplate 從模板生成文檔
func (doc *MarkdownDoc) GenerateFromTemplate(template model.DocTemplate) *MarkdownDoc {
	doc.AddTitle(template.Title, 4).
		AddBlankLines(1)

	doc.AddContent("| 欄位名稱 | 資料型態 | 說明 | 備註 |").
		AddBlankLines(0)
	doc.AddContent("| -------- | -------- | ---- | ---- |").
		AddBlankLines(0)

	for _, field := range template.Fields {
		doc.AddContent(fmt.Sprintf("| %s | %s | %s | %s |",
			field.Name, field.DataType, field.Description, field.Note)).
			AddBlankLines(0)
	}

	doc.AddBlankLines(2)

	for _, subTemplate := range template.SubTemplates {
		doc.AddTitle(subTemplate.Name+" 結構:", 4).
			AddBlankLines(1)

		doc.AddContent("| 欄位名稱 | 資料型態 | 說明 | 備註 |").
			AddBlankLines(0)
		doc.AddContent("| -------- | -------- | ---- | ---- |").
			AddBlankLines(0)

		for _, field := range subTemplate.Fields {
			doc.AddContent(fmt.Sprintf("| %s | %s | %s | %s |",
				field.Name, field.DataType, field.Description, field.Note)).
				AddBlankLines(0)
		}

		doc.AddBlankLines(2)
	}

	return doc
}