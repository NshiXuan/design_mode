package chain

import (
	"testing"
)

func TestChain(t *testing.T) {
	// if 时间 {}
	// if 分组 {}
	c := OrderChain{}
	c.add(&TimeFilter{})
	c.add(&GroupFilter{})
	if err := c.Filter(); err != nil {
		t.Log(err)
	}
}
