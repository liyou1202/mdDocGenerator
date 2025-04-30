### 取得詳情

- PATH: `/api/admin/getSomthing`
- Protocol: HTTP
- Method: GET
- Header:
  - `Accept: application/json`
  - `Token: JWT`


#### Request

| name            | required | data type  | description                      | note |
| --------------- | -------- | ---------- | -------------------------------- | ---------- |
| uid             |          | string     | 使用者ID                |            |
| deliveryId      |          | array      | 配送ID列表               |            |
| comment         |          | string     | 重設原因                 |            |
| base64          |          | object     |                      |            |


##### base64

| name            | required | data type  | description                      | note |
| --------------- | -------- | ---------- | -------------------------------- | ---------- |
| file            | Y        | string     | base64 編碼字串          |            |
| filename        | Y        | string     | 檔案名稱                 |            |


```json
{
	"uid": "",
	"deliveryId": ["9a23da1b-97eb-48d9-afd3-2f6116f10cf9"],
	"comment": "重設原因",
	"base64": {
		"file": "base 64 encode string",
		"filename": "name.jpg"
	}
}
```


#### Response Body

| name            | required | data type  | description                      | note |
| --------------- | -------- | ---------- | -------------------------------- | ---------- |
| response        |          | object     |                      |            |


##### response

| name            | required | data type  | description                      | note |
| --------------- | -------- | ---------- | -------------------------------- | ---------- |
| success         |          | boolean    | 是否成功                 |            |
| message         |          | string     | 回應訊息                 |            |


```json
{
    "success": true,
    "message": "行程重設成功"
}
```


