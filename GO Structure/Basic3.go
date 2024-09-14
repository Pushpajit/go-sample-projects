package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Id     int
	Name   string
	Stream string
	Age    int
}

// This is a way of creating method (Just like a method of a class object)
func (s *Student) toUpper() {
	s.Name = strings.ToUpper(s.Name)
	s.Stream = strings.ToUpper(s.Stream)
}

func main() {
	s1 := Student{1, "pushpajit", "bca", 25}
	s1.toUpper()

	fmt.Printf("%+#v", s1)
}
