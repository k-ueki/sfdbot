package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func DBConnection() (*gorm.DB, error) {
	//db, err := gorm.Open(config.Config.SQLDriver, fmt.Sprintf("root:@/%s?parseTime=true", config.Config.DBName))
	db, err := gorm.Open("mysql", "root:@/sfdbot?parseTime=true")
	if err != nil {
		log.Fatal("err:", err)
		return nil, err
	}
	return db, nil
}

func init() {
	db, err := DBConnection()
	if err != nil {
		log.Fatal("cannot connect to db")
		os.Exit(1)
	}

	//ProductCode         string
	//Side                string
	//Price               float64
	//Size                float64
	//Commission          float64
	//SwapPointAccumulate float64
	//RequireCollateral   float64
	//OpenDate            string
	//Leverage            float64
	//Pnl                 float64
	//Std                 float64
	if err := db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS position (
				product_code    VARCHAR(55) NOT NULL,
				side      		VARCHAR(55) NOT NULL,
				price 			INT UNSIGNED NOT NULL,
				size			FLOAT UNSIGNED NOT NULL,
				date			DATE NOT NULL)`)).Error; err != nil {
		log.Fatalf("cannot create table: position, err: %v", err)
		os.Exit(1)
	}
}
