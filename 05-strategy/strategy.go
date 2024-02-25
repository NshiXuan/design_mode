package strategy

import "fmt"

type SaveStrategy interface {
	Save()
}

type SaveFile struct {
}

func (s *SaveFile) Save() {
	fmt.Println("SaveFile")
}

type SaveDB struct {
}

func (s *SaveDB) Save() {
	fmt.Println("SaveDB")
}

var strategy = map[string]SaveStrategy{
	"file": &SaveFile{},
	"db":   &SaveDB{},
}
