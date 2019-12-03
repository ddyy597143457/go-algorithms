package main

import (
	"ddyy/go-algorithms/timer"
	"time"
)

func main() {
	arr := make([]*timer.Pitem,0)
	t := timer.NewPQueue(arr,true)
	//now := time.Now()
	day1 := time.Date(2019,12,03,23,23,55,0,time.Local)
	t.Add(timer.NewPitem(day1,"睡觉"))
	day2 := time.Date(2019,12,03,23,24,00,0,time.Local)
	t.Add(timer.NewPitem(day2,"吃饭"))
	day3 := time.Date(2019,12,03,23,37,30,0,time.Local)
	t.Add(timer.NewPitem(day3,"看电影"))
	day4 := time.Date(2019,12,03,23,41,10,0,time.Local)
	t.Add(timer.NewPitem(day4,"累了"))
	t.JobStart()
}
