package dto

type TrackingResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Summary Summary        `json:"summary"`
	Detail  Detail         `json:"detail"`
	History []HistoryEntry `json:"history"`
}

type Summary struct {
	Awb     string `json:"awb"`
	Courier string `json:"courier"`
	Service string `json:"service"`
	Status  string `json:"status"`
	Date    string `json:"date"`
	Desc    string `json:"desc"`
	Amount  string `json:"amount"`
	Weight  string `json:"weight"`
}

type Detail struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Shipper     string `json:"shipper"`
	Receiver    string `json:"receiver"`
}

type HistoryEntry struct {
	Date     string `json:"date"`
	Desc     string `json:"desc"`
	Location string `json:"location"`
}
