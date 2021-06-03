package Models

type Address struct {
	Id            int
	UserId        int
	ReceiveName   string
	PhoneType     int
	Phone         string
	DetailAddress string
	Suburb        string
	State         string
	Country       string
	PostCode      string
	IsDefault     int
	AddTime       string
	UpdateTime    string
	Type          int
}
