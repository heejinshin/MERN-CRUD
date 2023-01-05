package model

type Board struct {   // db의 컬럼 항목들을 필드로 가지고 있다. snakeCase를 사용했다. 
	ID          uint   `gorm:"primarykey;autoIncrement" json:"_id"`
	CompanyName string `json:"companyName"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Location    string `json:"location"`
	Link        string `json:"link"`
	Description string `json:"description"`
}


// 클라이언트가 GET요청 보낼때 Payload에 보면 위 항목들이 있잖아. 그걸 바인딩 
	// 시키려고 json 지정 해주는 거. (서버에서 받을 수 있어야하니까)
	// 응답을 줄때도 마찬가지로 위와 같이 줘야 한다. 

// view요청 때 _id 가 쓰여서 (프론트에서) _id를 키값으로 넣어줘야한다. (프론트에서 없애주던지 so 맞춰줘야 )