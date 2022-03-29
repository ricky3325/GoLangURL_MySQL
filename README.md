# URL Shortener
###啟動服務
```
docker-compose up -d
```
###第一次使用
連接資料庫:<http://localhost/myproject?username=admin&password=0>
建立資料表：<http://localhost/myproject?username=admin&password=1>

###每次開啟
連接資料庫:<http://localhost/myproject?username=admin&password=0>
**RUN Unit Tests**
```
docker exec -it myproject /bin/sh
go test -v
```

###API 操作

----
#####POST http://localhost/login1

**Request methods**

| Request methods/headers | Value |
| ------------- | ------------------------------ |
| Method      | POST       |
| Content-Type   | application/json     |

**Request parameters**

| Parameter name | Required/optional | Type | Description |
| --------- | ------------ |------ |------------ |
| Url      | Required    |	string    |原始網址    |
| ExpireAt   | Required  |	string    | 到期時間格式：2022-03-29 10:56:00    |

**Response**

| Response header | Value |
| ------------- | ------------------------------ |
| Status  | 200: Success       |
| Content-Type   | application/json     |

**Response body**
The response body is a JSON object type.

| Name | Type | Description |
| --------- |------ |------------ |
| id      |	string    |短網址的ID    |
| shortUrl  |	string |完整短網址|

----

#####GET http://localhost/login1/{id}

**Request methods**

| Request methods/headers | Value |
| ------------- | ------------------------------ |
| Method      | GET       |


**Response**

| Response header | Value |
| ------------- | ------------------------------ |
| Status  | 200: Success<br>404: Not Found|


**Response body**
The response body is a JSON object type.

| 狀態 |動作|
| ------------- | ------------------------------ |
| 200  | 跳轉到網址|
| 404  | 跳轉到404頁面|
