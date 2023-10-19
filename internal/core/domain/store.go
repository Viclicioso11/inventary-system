package domain

type Store struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"createdAt"`
	ImageKey  string `json:"imageKey"`
}
