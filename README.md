*IP流量控管練習*
====================
&emsp;

- [專案功能](#專案功能) 
- [流量管制目的](#流量管制目的)
- [開發工具](#開發工具)
- [執行](#執行)
- [Api文件](#api文件)
- [環境變數設定檔](#環境變數設定檔)
- [目錄結構說明](#目錄結構說明)
- [Golang 套件選擇](#golang-套件選擇)
- [流量控管設計概念](#流量控管設計概念)

&emsp;

## 專案功能與概念
除了*流量控管*外，還有簡易的*使用者系統*、*驗證系統*、*發文系統*

&emsp;

- 專案運用了**RESTful API**的概念設計
- 流量管制為**全域的middware**，只要有請求發生都會列入計算
- 除了**登入、與創建使用者**，其他api都要經過驗證jwt token 的 middware

&emsp;

## 流量管制目的

- 限制每小時來自同一個 IP 的請求數量**不得超過 1000**
- 在 response headers 中加入剩餘的請求數量 (X-RateLimit-Remaining) 以及 rate limit 歸零的時間 (X-RateLimit-Reset)
- 如果超過限制的話就回傳 **429 (Too Many Requests)**

&emsp;

## 開發工具

- 開發語言 ： Golang
- 資料庫 ： MySQL
- 快取資料庫 ： Redis
- Api 文件 ： Swagger

&emsp;
 
## 執行
```
go run main.go
```

&emsp;

## Api文件

- SwaggerUI : `http://localhost:8080/swagger/index.html`

&emsp;

## 環境變數設定檔
```
config/config.yml
```

## 目錄結構說明

- api  
  - 有關api相關的設定 ex:`gin.group` `api URL`
- config
  - auth
    - 有關 JWT Token 的發放與驗證
  - db
    - 連結 database 的設定
  - rdb
    - 連結 redis 的設定
  - config
    - 讀取外部環境變數 ex: `db` `redis` `jwt secret`
- docs
  - swagger套件相關
- init
  - migrations
    - GORM 的 AutoMigrate
- pkg
  - components
    - Handler function
  - middware
    - 中介程式 ex: `flow control` `auth`
  - models
    - models

&emsp;

## Golang 套件選擇

- ORM ： [GORM](https://gorm.io)
- Web framework : [Gin](https://github.com/gin-gonic/gin)
- Redis : [go-redis](https://github.com/go-redis/redis)
- Swagger : [swag](https://github.com/swaggo/swag)
- JWT : [jwt-go](https://github.com/dgrijalva/jwt-go)
- 讀取外部環境變數 : [configor](https://github.com/jinzhu/configor)
- Testing : [testify](https://github.com/stretchr/testify)
- 密碼雜湊 : [bcrypt](https://golang.org/x/crypto/bcrypt)

&emsp;

## 流量控管設計概念

- 為什麼選擇`redis`
  - 因為redis速度非常快，而流量控管只需要紀錄 IP,請求次數,重置時間，不需要使用到mysql之類的database 影響效能

&emsp;

- redis設計
  - 使用`Hash`型別
  - *Key* : IP位址
    - value count : 紀錄次數
    - value reset : 紀錄重置時間

&emsp;

- middware邏輯設計
  - 首先檢查**使用者IP**的 Key是否存在 (不存在表示此ip第一次造訪網站)
    - 不存在：**建立新Key值**，設定reset為一小時後，次數為1
  - 如果存在：先檢查現在時間是否已經大於重置時間
    - 如果大於重置時間：直接更新reset為一小時後，次數為1
  - 如果還沒到重置時間：檢查次數是否超過上限
    - 超過上限，回傳429
  - 沒超過上限，次數加1，reset不變

&emsp;