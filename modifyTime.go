package main

import (
	"time"
	"fmt"
	"strconv"
	"reflect"
)

func modify(h string)  {
	t:=time.Now()
	fmt.Println("t:",t)
	d, d1 := time.ParseDuration(h)
	fmt.Println("d:",d,d1)
	t2 := t.Add(d)
	t2.UnixNano()
	fmt.Println("t2:",t2)
	fmt.Println(strconv.FormatInt(t2.UnixNano()/1e6,10))
	timeStr := t2.Format("2006-01-02 15:04:05")
	fmt.Println("timeStr:", timeStr)
	price:="15.9"
	p1,_:=strconv.ParseFloat(price, 64)
	p2:=p1*100
	fmt.Println("啊啊啊啊",p2)
}
func main()  {

	a:=30.0
	fmt.Println(reflect.TypeOf(a))
	day:=time.Now().Day()
	d:=strconv.Itoa(day)
	var s string
	s+=d
	s+="39448"
	fmt.Println(s)
	fmt.Println("======day:",day,reflect.TypeOf(day))
	var hour string
	hour="-10m"
	modify(hour)
	fmt.Println("===")

	var usernewid =1200005414
	fmt.Println("几号表：",usernewid%32)

}