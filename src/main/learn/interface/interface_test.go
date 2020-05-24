package _interface

import "testing"

//定义接口
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"java Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
