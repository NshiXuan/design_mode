# 01-单例模式
两种模式:
  1.饿汉式模式: 在包加载的时候把结构体给创建了,返回的是同一个地址的变量
```go
  type Person struct {
	t time.Time
}

var p *Person

func init() {
	p = &Person{t: time.Now()}
}

func GetInstance() *Person {
	return p
}
```

  2.懒汉式模式: 只有结构体用到时候才创建结构体,返回的是同一个地址的变量
```go
type Person struct {
	t time.Time
}

var lazyPerson *Person
var once = sync.Once{} // 只执行一次

func GetInstance() *Person {
	// Do 中的函数只有第一次调用时执行 之后的会忽略
	once.Do(func() {
		lazyPerson = &Person{t: time.Now()}
	})
	return lazyPerson
}
```
- 测试
```go
func TestSingleon(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			p := GetInstance()
			fmt.Printf("p: %v\n", p)
			wg.Done()
		}()
	}
	wg.Wait()
}
```

# 02-工厂模式

```go
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
	fmt.Println("JsonParse")
}

func NewParse(s string) Parser {
	switch s {
	case "json":
		return &JsonParse{}
	case "xml":
		return &XmlParse{}
	default:
		return nil
	}
}
```
- 测试
```go
func TestFactory(t *testing.T) {
	jsonParse := NewParse("json")
	jsonParse.Parse()
	xmlParse := NewParse("xml")
	xmlParse.Parse()
}
```