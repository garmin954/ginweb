package model

type User struct {
	Id       int64   `xrom:"pk autoincr" json:"id"`
	UserName string  `xrom:"varchar(20)" json:"user_name"`
	Mobile   string  `xrom:"varchar(11)" json:"mobile"`
	Password string  `xrom:"varchar(255)" json:"password"`
	CreateAt int64   `xrom:"bigint" json:"create_at"`
	Avatar   string  `xrom:"varchar(255)" json:"avatar"`
	Balance  float64 `xrom:"double" json:"balance"`
	IsActive int8    `xrom:"tinyint" json:"is_active"`
	City     string  `xrom:"varchar(10)" json:"city"`
}
