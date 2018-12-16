package model

type Users struct {
	Model
	FirstName  string
	MiddleName string
	LastName   string
	Username   string `sql:",notnull, unique"`
	Password   string `sql:",notnull"`

	lat          float64
	long         float64
	TimeZone     string
	IsVerified   bool
	isActive     bool
	MembershipId int64
	Membership   *Membership

	Bank       float64
	ExchangeId int64
	Exchange   *Exchange
	Rating     int `sql:"-"`
	Deleatables
}

func (users *Users) GetUser(userId string) {

}
