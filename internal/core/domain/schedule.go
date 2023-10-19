package domain

type Schedule struct {
	Id          int        `json:"id"`
	DayOfWeekId int        `json:"dayOfWeekId"`
	StoreId     string     `json:"storeId"`
	OpenAt      string     `json:"openAt"`
	CloseAt     string     `json:"closeAt"`
	DayOfWeek   *DayOfWeek `json:"dayOfWeek"`
	Store       *Store     `json:"store"`
}
