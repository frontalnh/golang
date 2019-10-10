package module1

import "fmt"

type Module1 struct{}

func (m *Module1) Hello() {
	fmt.Println("hello")
}
