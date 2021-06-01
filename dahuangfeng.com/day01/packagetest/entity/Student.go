package entity

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
	Addr string
}

func (s *Student) PrintAge() {
	if s != nil {
		fmt.Println(s.Age)
	}
}
func (s Student) SumAge() int {
	ret := s.Age + 1
	return ret
}
