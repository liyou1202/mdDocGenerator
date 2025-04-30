# MD Doc Generator

這是一個用於生成 API 文檔的工具，可以根據預定義的模板自動生成 Markdown 格式的 API 文檔。

## 功能特點

- 支援自定義 API 文檔模板
- 自動生成標準化的 Markdown 文檔
- 支援請求和回應範例的自定義
- 支援 HTTP 標頭設定
- 支援多種資料型別和巢狀結構

## 安裝

```bash
# 克隆專案
git clone [your-repository-url]

# 進入專案目錄
cd mdDocGenerator

# 安裝依賴
go mod download
```

## 使用方法

### 1. 定義 API 文檔模板

在 `config/templates.go` 中定義 API 文檔模板：

```go
var APITemplates = map[string]model.APIDocTemplate{
    "yourAPI": {
        Title:       "API 標題",
        Description: "API 描述",
        Path:        "/api/your/path",
        Protocol:    "HTTP",
        Method:      "POST",
        Headers: []model.Header{
            {Key: "Accept", Value: "application/json"},
            {Key: "Token", Value: "JWT"},
        },
    },
}
```

### 2. 定義請求和回應範例

同樣在 `config/templates.go` 中定義：

```go
var RequestExamples = map[string]string{
    "yourAPI": `{
        "field1": "value1",
        "field2": "value2"
    }`,
}

var ResponseExamples = map[string]string{
    "yourAPI": `{
        "success": true,
        "message": "操作成功"
    }`,
}
```

### 3. 生成文檔

```go
func main() {
    // 從配置中獲取 API 文檔模板
    template := config.APITemplates["yourAPI"]

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
        Request:     yourRequestFields,
        Response:    yourResponseFields,
    }

    // 生成 API 文檔
    doc.GenerateAPIDoc(apiDoc)

    // 導出文檔
    doc.Export("output/api_doc.md")
}
```

## 輸出範例

生成的文檔將包含以下部分：

- API 基本資訊（標題、描述、路徑等）
- 請求標頭
- 請求參數說明
- 請求範例
- 回應參數說明
- 回應範例

## 注意事項

- 確保所有必要的欄位都已正確定義
- 檢查生成的文檔格式是否符合預期
- 定期更新 API 文檔以保持同步

## 貢獻

歡迎提交 Issue 和 Pull Request 來改進這個工具。

## 授權

[授權類型] - 查看 [LICENSE](LICENSE) 文件了解更多資訊。