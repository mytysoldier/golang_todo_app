package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang_todo_app/func/database"
)

// パスルーティング定義
func Router() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	database.DbInit()

	// 初期表示
	router.GET("/", func(ctx *gin.Context) {
		todos := database.DbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// TODO作成
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		database.DbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	// TODO詳細
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := database.DbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	// TODO更新
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		database.DbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})

	// TODO削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := database.DbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	// TODO削除
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		database.DbDelete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
