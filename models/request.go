package models

type URLRequest struct {
	URL         string `json:"url" binding:"required"`
	CustomShort string `json:"custom_short"`
	Expiry      int    `json:"expiry"` // in hours
}
