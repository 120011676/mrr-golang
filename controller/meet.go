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
	result := entity.Result{Status: false, Date: time.Now().Format("2006-01-02 15:04:05")}
	if startDate == "" {
		result.Code = -1
		result.Msg = "预约开始时间不能为空"
	} else if endDate == "" {
		result.Code = -2
		result.Msg = "预约结束时间不能为空"
	} else {
		sd, sdErr := time.Parse("2006-01-02 15:04:05", startDate)
		if sdErr != nil {
			result.Code = -3
			result.Msg = "预约开始时间格式不正确，必须是yyyy-MM-dd HH:mm:ss格式"
		} else {
			sd = time.Unix(sd.UnixNano()/1e9, 0)
		}
		ed, edErr := time.Parse("2006-01-02 15:04:05", endDate)
		if edErr != nil {
			result.Code = -4
			result.Msg = "预约结束时间格式不正确，必须是yyyy-MM-dd HH:mm:ss格式"
		} else {
			ed = time.Unix(ed.UnixNano()/1e9, 0)
		}
		if sdErr == nil && edErr == nil {
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
			if e != nil {
				result.Code = -100
				result.Msg = e.Error()
			} else {
				result.Status = true
				result.Msg = "成功"
			}
		}
	}
	json, _ := json.Marshal(result)
	io.WriteString(w, string(json))
}

func Cancel(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	password := r.FormValue("password")
	result := entity.Result{Status: false, Date: time.Now().Format("2006-01-02 15:04:05")}
	if id == "" {
		result.Code = -1
		result.Msg = "id不能为空"
	} else if password == "" {
		result.Code = -2
		result.Msg = "密码不能为空"
	} else {
		e := service.Cancel(id, password)
		if e != nil {
			result.Code = -100
			result.Msg = e.Error()
		} else {
			result.Status = true
			result.Msg = "成功"
		}
	}
	json, _ := json.Marshal(result)
	io.WriteString(w, string(json))
}

func Query(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	floor := r.FormValue("floor")
	room := r.FormValue("room")
	startDate := r.FormValue("startDate")
	endDate := r.FormValue("endDate")
	var sd time.Time
	var ed time.Time
	var sdErr error
	var edErr error
	result := entity.Result{Status: false, Date: time.Now().Format("2006-01-02 15:04:05")}
	if startDate != "" {
		sd, sdErr = time.Parse("2006-01-02 15:04:05", startDate)
		if sdErr != nil {
			result.Code = -1
			result.Msg = "开始时间格式不正确，必须是yyyy-MM-dd HH:mm:ss格式"
		} else {
			sd = time.Unix(sd.UnixNano()/1e9, 0)
		}
	}
	if endDate != "" {
		ed, edErr = time.Parse("2006-01-02 15:04:05", endDate)
		if edErr != nil {
			result.Code = -2
			result.Msg = "结束时间格式不正确，必须是yyyy-MM-dd HH:mm:ss格式"
		} else {
			ed = time.Unix(ed.UnixNano()/1e9, 0)
		}
	}
	if sdErr == nil && edErr == nil {
		ms, e := service.Query(city, floor, room, sd, ed)
		jms := make([]entity.MeetJson, len(ms))
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
		if e != nil {
			result.Code = -100
			result.Msg = e.Error()
		} else {
			result.Status = true
			result.Msg = "成功"
			result.Data = jms
		}
	}
	json, _ := json.Marshal(result)
	io.WriteString(w, string(json))
}
