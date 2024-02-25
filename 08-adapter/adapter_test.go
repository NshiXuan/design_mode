package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	tencenAdapter := TencenAdapter{TSDK: &TencenSDK{}}
	tencenAdapter.Pay()
	aliAdapter := AliAdapter{ASDK: &AliSDK{}}
	aliAdapter.Pay()
}
