package service

import (
	"../entity"
	"../mongodb"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	TABLE = "meet"
)

func Reserve(m entity.Meet) error {
	if m.StartDate.Before(m.EndDate) {
		session, database := mongodb.Connect()
		defer session.Close()
		c, _ := database.C(TABLE).Find(bson.M{"status": true, "city": m.City, "floor": m.Floor, "room": m.Room, "startdate": bson.M{"$not": bson.M{"$gte": m.EndDate}}, "enddate": bson.M{"$not": bson.M{"$lte": m.StartDate}}}).Count()
		log.Print(c)
		if c == 0 {
			err := database.C(TABLE).Insert(&m)
			if err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			return errors.New("预约时间冲突")
		}
	} else {
		return errors.New("预约开始时间必须小于结束时间")
	}
	return nil
}

func Cancel(id string, password string) error {
	session, database := mongodb.Connect()
	defer session.Close()
	err := database.C(TABLE).Update(bson.M{"_id": bson.ObjectIdHex(id), "password": password}, bson.M{"$set": bson.M{"status": false}})
	if err != nil {
		return errors.New("取消失败，请检查密码是否正确")
	}
	return nil
}

func Query(city string, floor string, room string, startDate time.Time, endDate time.Time) ([]entity.Meet, error) {
	session, database := mongodb.Connect()
	defer session.Close()
	var ms []entity.Meet
	m := bson.M{}
	if city != "" {
		m["city"] = city
	}
	if floor != "" {
		m["floor"] = floor
	}
	if room != "" {
		m["room"] = room
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		if startDate.Before(endDate) {
			m["startdate"] = bson.M{"$not": bson.M{"$gte": endDate}}
			m["enddate"] = bson.M{"$not": bson.M{"$lte": startDate}}
		} else {
			return nil, errors.New("开始时间必须小于结束时间")
		}
	} else if !startDate.IsZero() {
		m["startdate"] = bson.M{"$lte": startDate}
	} else if !endDate.IsZero() {
		m["enddate"] = bson.M{"$gte": endDate}
	}
	m["status"] = true
	log.Print(m)
	err := database.C(TABLE).Find(m).Sort("status", "startdate").All(&ms)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("查询错误")
	}
	return ms, nil
}
