package factory

import "fmt"

// type Parser interface {
// 	Parse()
// }

// type JsonParse struct {
// }

// func (p *JsonParse) Parse() {
// 	fmt.Println("JsonParse")
// }

// type XmlParse struct {
// }

// func (p *XmlParse) Parse() {
// 	fmt.Println("XmlParse")
// }

// func NewParse(s string) Parser {
// 	switch s {
// 	case "json":
// 		return &JsonParse{}
// 	case "xml":
// 		return &XmlParse{}
// 	default:
// 		return nil
// 	}
// }

type ParserFactory interface {
	CreateParser() Parser
}

type Parser interface {
	Parse()
}

type JsonParse struct {
}

func (p *JsonParse) Parse() {
	fmt.Println("JsonParse")
}

type XmlParse struct {
}

func (p *XmlParse) Parse() {
	fmt.Println("XmlParse")
}

type JsonParseFactory struct {
}

func (f *JsonParseFactory) CreateParser() Parser {
	// TODO(nsx):
	return &JsonParse{}
}

type XmlParseFactory struct {
}

func (f *XmlParseFactory) CreateParser() Parser {
	// TODO(nsx):
	return &XmlParse{}
}

func NewParseFactory(s string) ParserFactory {
	switch s {
	case "json":
		return &JsonParseFactory{}
	case "xml":
		return &XmlParseFactory{}
	default:
		return nil
	}
}
