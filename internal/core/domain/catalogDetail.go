package domain

type CatalogDetail struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Code        string   `json:"code"`
	Enabled     bool     `json:"enabled"`
	CreatedAt   string   `json:"createdAt"`
	CatalogId   int      `json:"catalogId"`
	Catalog     *Catalog `json:"catalog"`
}
