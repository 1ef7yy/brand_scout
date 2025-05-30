package models

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote   string `json:"quote"`
}

type CreateQuoteDTO struct {
	Author string `json:"author"`
	Quote   string `json:"quote"`
}

func (c CreateQuoteDTO) Valid() bool {
	return c.Author != "" && c.Quote != ""
}
