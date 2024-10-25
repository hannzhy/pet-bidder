package types

import "time"

type AdCampaign struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Budget    float64   `json:"budget"`
}
