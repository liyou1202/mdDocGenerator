package pkg

import (
	"fmt"
	"os"
	"strings"
)

// MarkdownDoc 定義 Markdown 文檔的結構
type MarkdownDoc struct {
	content *strings.Builder
}

// NewDoc 創建新的 Markdown 文檔
func NewDoc() *MarkdownDoc {
	return &MarkdownDoc{
		content: &strings.Builder{},
	}
}

// WriteLine 寫入一行內容
func (doc *MarkdownDoc) WriteLine(content string) {
	doc.content.WriteString(content + "\n")
}

// AddTitle 添加標題
func (doc *MarkdownDoc) AddTitle(t string, lv int) *MarkdownDoc {
	if lv > 6 || lv < 1 {
		fmt.Printf("failed to add Title %s in level: %d\n", t, lv)
		return doc
	}
	mdSyntax := strings.Repeat("#", lv) + " " + t
	doc.WriteLine(mdSyntax)
	return doc
}

// AddInterval 添加分隔線
func (doc *MarkdownDoc) AddInterval() *MarkdownDoc {
	mdSyntax := strings.Repeat("-", 3)
	doc.WriteLine(mdSyntax)
	return doc
}

// AddImage 添加圖片
func (doc *MarkdownDoc) AddImage(placeholder, path, title string) *MarkdownDoc {
	mdSyntax := fmt.Sprintf("![%s](%s) %s", placeholder, path, title)
	doc.WriteLine(mdSyntax)
	return doc
}

// AddBlankLines 添加空行
func (doc *MarkdownDoc) AddBlankLines(lv int) *MarkdownDoc {
	for i := 0; i < lv; i++ {
		doc.WriteLine("")
	}
	return doc
}

// AddCodeBlock 添加代碼塊
func (doc *MarkdownDoc) AddCodeBlock(code, language string) *MarkdownDoc {
	mdSyntax := fmt.Sprintf("```%s\n%s\n```", language, code)
	doc.WriteLine(mdSyntax)
	return doc
}

// AddLink 添加鏈接
func (doc *MarkdownDoc) AddLink(text, path string) *MarkdownDoc {
	mdSyntax := fmt.Sprintf("[%s](%s)", text, path)
	doc.content.WriteString(mdSyntax)
	return doc
}

// AddContent 添加內容
func (doc *MarkdownDoc) AddContent(text string) *MarkdownDoc {
	doc.content.WriteString(text)
	return doc
}

// Export 導出文檔
func (doc *MarkdownDoc) Export(filename string) error {
	return os.WriteFile(filename, []byte(doc.content.String()), os.ModePerm)
}