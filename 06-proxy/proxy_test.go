package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	user := NewUser()
	userProxy := NewUserCacheProxy(user)
	userProxy.GetInfo()
}
