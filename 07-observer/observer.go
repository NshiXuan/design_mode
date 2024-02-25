package observer

import "fmt"

type IOrder interface {
	Register(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

type Order struct {
	obs []IObserver
}

func (o *Order) Register(observer IObserver) {
	o.obs = append(o.obs, observer)
}

func (o *Order) Notify(msg string) {
	for _, ob := range o.obs {
		ob.Update(msg)
	}
}

type ScoreObserver struct {
}

func (s *ScoreObserver) Update(msg string) {
	fmt.Println(msg + ": 积分增加10")
}

type EmailObserver struct {
}

func (s *EmailObserver) Update(msg string) {
	fmt.Println(msg + ": 发送邮件")
}
