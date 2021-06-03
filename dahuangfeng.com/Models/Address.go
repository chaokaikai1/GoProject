package Models

type Address struct {
	Id            int
	UserId        int
	ReceiveName   string
	PhoneType     int
	phone         string `json:Phone`
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
