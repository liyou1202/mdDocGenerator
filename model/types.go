package model

// APIDocTemplate 定義 API 文檔的基本結構
type APIDocTemplate struct {
	Title       string
	Description string
	Path        string
	Protocol    string
	Method      string
	Headers     []Header
}

// APIDoc 定義 API 文檔的結構
type APIDoc struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Path            string   `json:"path"`
	Protocol        string   `json:"protocol"`
	Method          string   `json:"method"`
	Headers         []Header `json:"headers"`
	Request         DocField `json:"request"`
	RequestExample  string   `json:"requestExample"`
	Response        DocField `json:"response"`
	ResponseExample string   `json:"responseExample"`
}

// Header 定義 HTTP 標頭的結構
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// DocField 定義文檔欄位的結構
type DocField struct {
	Name        string     `json:"name"`
	Required    bool       `json:"required"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
	Note        string     `json:"note"`
	Fields      []DocField `json:"fields,omitempty"`
}

// DocTemplate 定義文檔模板的結構
type DocTemplate struct {
	Title        string
	Fields       []FieldInfo
	SubTemplates []SubTemplate
}

// SubTemplate 定義子模板的結構
type SubTemplate struct {
	Name   string
	Fields []FieldInfo
}

// FieldInfo 定義欄位資訊的結構
type FieldInfo struct {
	Name        string
	DataType    string
	Description string
	Note        string
}
