package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Logininfo struct {
	gorm.Model
	Username string `form : "username" binding : "required" gorm : "unique;not null"`
	Password string `form : "password" binding : "required"`
}

//DB初期化
func DbInit() {
	db, err := gorm.Open("sqlite3", "./database.sqlite3")
	checkerr(err)
	db.AutoMigrate(&Logininfo{})
	defer db.Close()
}

func DbGetOne(username string) Logininfo{
	db, err := gorm.Open("sqlite3", "./database.sqlite3")
	checkerr(err)
	var user Logininfo
	db.Where(&user, "User = ?", username)
	db.Close()
	return user
}

//エラーチェック
func checkerr(err error) {
	if err != nil {
		
		panic(err)
	}
}