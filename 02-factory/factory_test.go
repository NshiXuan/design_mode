package factory

import "testing"

func TestFactory(t *testing.T) {
	// jsonParse := NewParse("json")
	// jsonParse.Parse()
	// xmlParse := NewParse("xml")
	// xmlParse.Parse()

	// 工厂方法也就是比简单工厂多封装了一层
	jsonFactory := NewParseFactory("json")
	jsonParse := jsonFactory.CreateParser()
	jsonParse.Parse()
	xmlFactory := NewParseFactory("xml")
	xmlFactory.CreateParser().Parse()
}
