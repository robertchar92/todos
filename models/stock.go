package models

type Stock struct {
	Open    int64 `json:"open_price"`
	High    int64 `json:"high_price"`
	Low     int64 `json:"low_price"`
	Close   int64 `json:"close_price"`
	Volume  int64 `json:"volume"`
	Value   int64 `json:"value"`
	PrevCP  int64 `json:"prev_price"`
	Average int64 `json:"average"`
}
