package strategy

import "testing"

func TestStrategy(t *testing.T) {
	saveStr := "db"
	switch saveStr {
	case "file":
		saveStrategy := strategy["file"]
		saveStrategy.Save()
	case "db":
		saveStrategy := strategy["db"]
		saveStrategy.Save()
	}
}
