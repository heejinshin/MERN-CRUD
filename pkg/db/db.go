package db

import (
	"github.com/jackbalageru/MERN-CRUD/go-server/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	gDB *gorm.DB
}


// post요청 받으면 동작하는 함수인건가? automigrate가 그 역할인듯. (DB에 테이블 write해주는)
func NewAndConnectGorm(dsn string) (*DBHandler, error) {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  // 데이터베이스 생성 
	// 왜 dsn 그대로 받지??? dsn이 어디서 선언??? 
	gormDB.AutoMigrate(&model.Board{})  // 테이블을 db에 write

	dbHandler := &DBHandler{   // DB핸들러 구조체의 인스턴스 생성 
		gDB: gormDB,  // 콜론으로 나타내는 이 문법이 이해가 안갔다. gDB가 gormDB 데이터베이스 라고 알려준것. 
		// 이 gDB를 통해서 데이터베이스에 접근할 예정, 일단 main에서 이걸 가져다가 NEWAPI에 전달한다. 
	}

	return dbHandler, err
}
