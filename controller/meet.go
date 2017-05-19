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
	m := entity.Meet{
		Id:             bson.NewObjectId(),
		City:           city,
		Floor:          floor,
		Room:           room,
		ConferenceName: conferenceName,
		People:         people,
		Phone:          phone,
		Password:       password,
		StartDate:      startDate,
		EndDate:        endDate,
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
	result := entity.Result{Status: false, Date: time.Now().Format("2006-01-02 15:04:05")}
	if e != nil {
		result.Msg = e.Error()
	} else {
		result.Status = true
		result.Msg = "成功"
		data, _ := json.Marshal(ms)
		result.Data = string(data)
	}
	json, _ := json.Marshal(result)
	io.WriteString(w, string(json))
}
