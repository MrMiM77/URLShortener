package db

import (
	"github.com/MrMiM77/URLShortener.git/internal/models"
	"github.com/jinzhu/gorm"
	"sync"
)

var instance *gorm.DB
var once sync.Once



func GetInstance() *gorm.DB{
	once.Do(func() {
		var err error
		//instance, err = gorm.Open()
		if err != nil{
			panic(err)
		}
		//TODO
		instance.AutoMigrate(&models.URLMap{})
	})
	return instance
}

func RefreshDataBase(){
	GetInstance().DropTable(models.URLMap{})
	GetInstance().AutoMigrate(models.URLMap{})
}
