# GoChat
使用Golang和JavaScript架設一個即時聊天的服務

## Features
- Register 註冊會員
- Login 使用帳號密碼登入
- Find Account 查詢帳號
- Add Friend 加入帳號為好友
- Show Friends 查詢已加好友清單
- Send Message 傳送訊息給好友
- Profile Config 設定帳號資訊
- Passowrd Change 修改密碼

## Dependencies
**Web Framework**
- go get -u github.com/gin-gonic/gin

**Configuration**
- go get -u github.com/spf13/viper

**Database**
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/mysql

**Logger**
- go get -u github.com/sirupsen/logrus

**UnitTest**
- go get -u github.com/stretchr/testify/assert 

**UUID**
- go get -u github.com/google/uuid

