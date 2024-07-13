package external

type OpeningHours struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

type Store struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Address      string       `json:"address"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	ZipCode      string       `json:"zipCode"`
	Phone        string       `json:"phone"`
	Email        string       `json:"email"`
	Website      string       `json:"website"`
	Category     string       `json:"category"`
	OpeningHours OpeningHours `json:"openingHours"`
}
