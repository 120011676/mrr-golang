package entity

import ()

type Result struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Date   string `json:"date"`
	Data   string `json:"data"`
}
