package entity

import ()

type MeetJson struct {
	Id             string `json:"id"`
	City           string `json:"city"`
	Floor          string `json:"floor"`
	Room           string `json:"room"`
	ConferenceName string `json:"conferenceName"`
	People         string `json:"people"`
	Phone          string `json:"phone"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	Status         bool   `json:"status"`
	CreateDate     string `json:"createDate"`
}
