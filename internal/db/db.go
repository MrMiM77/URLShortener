package db

import (
	"github.com/MrMiM77/URLShortener.git/config"
	"github.com/MrMiM77/URLShortener.git/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
	"sync"
)

var instance *gorm.DB
var once sync.Once



func GetInstance() *gorm.DB{
	once.Do(func() {
		var err error
		instance, err = gorm.Open("postgres","host="+config.GetInstance().DATABASE.HOST+" port="+
			strconv.Itoa(config.GetInstance().DATABASE.PORT)+
			" user="+config.GetInstance().DATABASE.USER+
			" dbname="+config.GetInstance().DATABASE.NAME+
			" password="+config.GetInstance().DATABASE.PASSWORD+
			" sslmode=disable")
		if err != nil{
			panic(err)
		}
		instance.AutoMigrate(&models.URLMap{})
	})
	return instance
}

func RefreshDataBase(){
	GetInstance().DropTable(models.URLMap{})
	GetInstance().AutoMigrate(models.URLMap{})
}
