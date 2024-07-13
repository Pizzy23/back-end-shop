package interfaces

type StorageCashBack struct {
	ID       string  `json:"id"`
	Cashback float64 `json:"cashback"`
}
type StorageBuying struct {
	ID     string `json:"id"`
	Buying int64  `json:"buying"`
}
type StorageUse struct {
	ID           string `json:"id"`
	CashBacksUse int64  `json:"buying"`
}

type Purchase struct {
	Value   float64 `json:"value"`
	Email   string  `json:"email"`
	Storage string  `json:"storage"`
}

type MktInput struct {
	Product     string `json:"product"`
	Value       int64  `json:"value"`
	Description string `json:"description"`
	Img         string `json:"img"`
}
