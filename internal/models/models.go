package models

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type CreateQuoteDTO struct {
	Author string `json:"author,omitempty"`
	Text   string `json:"text,omitempty"`
}

func (c CreateQuoteDTO) Valid() bool {
	return c.Author != "" && c.Text != ""
}
