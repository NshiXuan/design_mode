package observer

import "testing"

func TestObserver(t *testing.T) {
	order := Order{}
	order.Register(&ScoreObserver{})
	order.Notify("score observer")
	order.Register(&EmailObserver{})
	order.Notify("email observer")
}
