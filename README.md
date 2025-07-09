# AI Writer APIs

一套以 **Golang** + **PostgreSQL** 為後端、**Angular** 為前端框架開發的 **AI 產文平台**，結合 ChatGPT 語言模型，提供使用者高效產出內容的解決方案。
平台特色如下：
* 自訂參數輸入：使用者可依需求輸入主題、語氣、長度、關鍵字等條件，精準控制生成內容風格與方向
* AI 智能生成：透過 ChatGPT 快速生成高品質的文章草稿、文案或腳本
* 即時預覽與編輯：前端提供直覺式操作介面，產文結果可即時預覽與編輯，方便進一步調整與應用
* 內容歷程與儲存：可記錄使用者產出的內容歷史，方便追蹤與重複使用

系統採用前後端分離設計，具備良好的可擴充性與維護彈性，適用於行銷文案、部落格創作、自媒體經營等多種應用場景，有效提升內容產出效率與品質。

#Golang #Gin #PostgreSQL #Angular #PrimeNG #Swagger #ＯpenAI

## 專案連結

* 前端畫面：[點我查看](https://hsxxnil.notion.site/AI-11c5b51f95f581e3944bd0cc85581f2c)
* Swagger API 文件：[點我查看](https://hsxxnil.github.io/swagger-ui/?urls.primaryName=AI)

## 安裝
1. 下載專案

```bash
git clone https://github.com/Hsxxnil/ai_writer_apis.git
cd ai_writer_apis
```

2. 建立 Makefile

> 請根據您的作業系統選擇對應的範本進行複製：
* Linux / macOS
```bash
cp Makefile.example.linux Makefile
```

* Windows
```bash
copy Makefile.example.windows Makefile
```

3. 初始化

> 如為初次建立開發環境，請先根據您的作業系統安裝必要套件：
* Linux / macOS
```bash
brew install golang-migrate golangci-lint protobuf
```

* Windows（建議使用 Scoop，或手動安裝以下套件）：
```bash
scoop install golang-migrate golangci-lint protobuf
```

> 執行以下指令將自動安裝依賴套件並建立必要的目錄結構：
```bash
make setup
```

4. 設定環境參數

> 開啟並編輯以下檔案，填入資料庫連線資訊、JWT 金鑰等必要參數：
```file
config/debug_config.go
```

5. 更新套件

>執行以下指令升級相關套件
```bash
make update_lib
```

## 資料庫遷移

> 執行以下指令使用[golang-migrate](https://github.com/golang-migrate/migrate)做資料庫遷移及做資料表版控：
```bash
make migration
```

## 執行
> 執行以下指令在本地端啟動伺服器並自動重載：
```bash
make air
```

## License

本專案使用的 [Vodka](https://github.com/dylanlyu/vodka) 採用 [MIT License](https://opensource.org/licenses/MIT) 授權。
