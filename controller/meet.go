package controller

import (
	"../entity"
	"../service"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
	"time"
)

func Reserve(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	floor := r.FormValue("floor")
	room := r.FormValue("room")
	conferenceName := r.FormValue("conferenceName")
	people := r.FormValue("people")
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	startDate := r.FormValue("startDate")
	endDate := r.FormValue("endDate")
	sd, _ := time.Parse("2006-01-02 15:04:05", startDate)
	ed, _ := time.Parse("2006-01-02 15:04:05", endDate)
	m := entity.Meet{
		Id:             bson.NewObjectId(),
		City:           city,
		Floor:          floor,
		Room:           room,
		ConferenceName: conferenceName,
		People:         people,
		Phone:          phone,
		Password:       password,
		StartDate:      sd,
		EndDate:        ed,
		Status:         true,
		CreateDate:     time.Now(),
	}
	e := service.Reserve(m)
	result := entity.Result{Status: false, Date: time.Now().Format("2006-01-02 15:04:05")}
	if e != nil {
		result.Msg = e.Error()
	} else {
		result.Status = true
		result.Msg = "成功"
	}
	json, _ := json.Marshal(result)
	io.WriteString(w, string(json))
}

func Cancel(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	password := r.FormValue("password")
	e := service.Cancel(id, password)
	result := entity.Result{Status: false, Date: time.Now().Format("2006-01-02 15:04:05")}
	if e != nil {
		result.Msg = e.Error()
	} else {
		result.Status = true
		result.Msg = "成功"
	}
	json, _ := json.Marshal(result)
	io.WriteString(w, string(json))
}

func Query(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	floor := r.FormValue("floor")
	room := r.FormValue("room")
	ms, e := service.Query(city, floor, room)
	jms := make([]entity.MeetJson,len(ms))
	if ms != nil {
		for i := 0; i < len(ms); i++ {
			m := entity.MeetJson{
				Id:             ms[i].Id.Hex(),
				City:           ms[i].City,
				Floor:          ms[i].Floor,
				Room:           ms[i].Room,
				ConferenceName: ms[i].ConferenceName,
				People:         ms[i].People,
				Phone:          ms[i].Phone,
				StartDate:      ms[i].StartDate.Format("2006-01-02 15:04:05"),
				EndDate:        ms[i].EndDate.Format("2006-01-02 15:04:05"),
				Status:         ms[i].Status,
				CreateDate:     ms[i].CreateDate.Format("2006-01-02 15:04:05"),
			}
			jms[i] = m
		}
	}
	result := entity.Result{Status: false, Date: time.Now().Format("2006-01-02 15:04:05")}
	if e != nil {
		result.Msg = e.Error()
	} else {
		result.Status = true
		result.Msg = "成功"
		result.Data = jms
	}
	json, _ := json.Marshal(result)
	io.WriteString(w, string(json))
}
