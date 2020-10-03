module github.com/golang_todo_app

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-sqlite3 v1.14.3
)

// replace local.packages/database => ./func/database/database
