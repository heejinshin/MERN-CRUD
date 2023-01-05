package api

import "github.com/jackbalageru/MERN-CRUD/go-server/pkg/db"

type APIs struct {
	db *db.DBHandler  // 필드 인자에 DBHandler 포인터변수 넣어줌 
}

func NewAPI(h *db.DBHandler) *APIs {
	// a := APIS{db: h}
		// a.db = h 와 같은 의미이다. 
	// return a  
	return &APIs{db: h}   // DBHandler 포인터 인스턴스를 위 API구조체의 db 에 넣어준다. 
}
