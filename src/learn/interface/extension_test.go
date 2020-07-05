package _interface

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	//不支持强制类型转换
	//var dog Pet = new(Dog)
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao")
}
