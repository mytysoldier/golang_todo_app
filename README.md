# golang_todo_app
GO言語を用いたTODOサンプルアプリケーション

## 使用技術
___
* 言語  
 GO  
 HTML  
 CSS
* フレームワーク  
 Gin  
 GORM
* データベース  
Sqlite3
* エディター  
Visual Studio Code  
※任意のものでOKです。

## 動かすための準備
___
* GOのインストール  
※MacのHomebrewでのインストール手順です。  
brew install go
* Ginのインストール  
go get github.com/gin-gonic/gin
* GORMのインストール  
go get github.com/jinzhu/gorm
* Sqlite3のインストール  
go get github.com/mattn/go-sqlite3
* VSCodeにプロジェクトをインポート  
※この際の注意点として、プロジェクトは$GOPATH配下に配置する必要がある。  
※例えば、Githubからクローンしてきたプロジェクトであれば、  
$GOPATH/src/github.com/{gitのusername}/プロジェクト名  
main.goを開きF5キーで実行  

## 今後のTODO
___
* フロントをVue.jsに変更
* DBをNoSQLに変更
* Docker導入
