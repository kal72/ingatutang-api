package entity

type Base struct {
	Status  int
	Message string
	Data    interface{}
}

type User struct {
	Id          int
	Username    string `orm:"size(40)"`
	Password    string `orm:"size(40)"`
	ProfileName string `orm:"size(60)"`
	Phone       string `orm:"size(15)"`
	Image       string `orm:"size(255)"`
	Token       string `orm:"-"`
}

type Debt struct {
	Id          int
	NoTransaksi string `orm:"size(11)"`
	DebtFrom    *User   `orm:"rel(fk)"`
	DebtTo      *User   `orm:"rel(fk)"`
	Total       int    `orm:"size(11)"`
	StartDate   string `orm:"size(10)`
	ReturnDate  string `orm:"size(10)`
}
