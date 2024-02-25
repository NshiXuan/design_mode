package chain

import "fmt"

type OrderFilter interface {
	Filter() error
}

type OrderChain struct {
	filters []OrderFilter
}

func (o *OrderChain) add(filter OrderFilter) {
	o.filters = append(o.filters, filter)
}

func (o *OrderChain) Filter() error {
	for _, f := range o.filters {
		if err := f.Filter(); err != nil {
			return err
		}
	}
	return nil
}

type TimeFilter struct {
}

func (t *TimeFilter) Filter() error {
	fmt.Println("Time Filter")
	return nil
}

type GroupFilter struct {
}

func (t *GroupFilter) Filter() error {
	fmt.Println("Group Filter")
	return nil
}
