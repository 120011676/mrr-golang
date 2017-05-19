package service

import (
	"../entity"
	"../mongodb"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	TABLE = "meet"
)

func Reserve(m entity.Meet) error {
	session, database := mongodb.Connect()
	defer session.Close()
	err := database.C(TABLE).Insert(&m)
	if err != nil {
		log.Fatal(err)
		return err
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

func Query(city string, floor string, room string) ([]entity.Meet, error) {
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
	err := database.C(TABLE).Find(m).All(&ms)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	return ms, nil
}
