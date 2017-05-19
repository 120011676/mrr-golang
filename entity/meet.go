package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Meet struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	City           string        `json:"city"`
	Floor          string        `json:"floor"`
	Room           string        `json:"room"`
	ConferenceName string        `json:"conferenceName"`
	People         string        `json:"people"`
	Phone          string        `json:"phone"`
	Password       string        `json:"-"`
	StartDate      string        `json:"startDate"`
	EndDate        string        `json:"endDate"`
	Status         bool          `json:"status"`
	CreateDate     time.Time     `json:"createDate"`
}
