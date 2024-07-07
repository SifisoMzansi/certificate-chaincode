package certificatechaincode

type Ceremony struct {
	CeremonyID  string `json:"ceremonyID"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Venue       string `json:"Venue"`
	Description string `json:"description"`
}
