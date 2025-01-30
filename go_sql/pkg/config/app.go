package config


import(
	"fmt"
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"
)

var (
	db * gorm.DB
)

func Connect(){
	dsn := "host=127.0.0.1 user=postgres password=Kristan_009 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Calcutta"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil{
		panic(err)
	}
	
	db = d
	fmt.Printf("DB connected successfully")
}


func GetDB() *gorm.DB{
	return db
}

