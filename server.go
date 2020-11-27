package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

type Server struct{}

type Student struct {
	name    string
	subject string
	grade   float64
}

var Students []Student

func (this *Server) Recibe(s []string, reply *string) error {
	g, _ := strconv.ParseFloat(s[2], 8)
	//fmt.Println(Student{name: s[0], subject: s[1], grade: g})
	Students = append(Students, Student{name: s[0], subject: s[1], grade: g})
	*reply = "alumno recibido"
	return nil
}

func (this *Server) Promedio(a string, reply *float64) error {
	mean := 0.000

	if len(Students) > 0 {
		for i := 0; i < len(Students); i++ {
			if Students[i].name == a {
				mean = Students[i].grade
			}
		}

		*reply = mean
		return nil
	} else {
		return errors.New("primero registre alumnos jeje")
	}
}

func (this *Server) PromedioGral(a string, reply *float64) error {
	sum := 0.000
	mean := 0.000
	cant := float64(len(Students))

	if len(Students) > 0 {
		for i := 0; i < len(Students); i++ {
			sum += Students[i].grade
		}
		mean = sum / cant

		*reply = mean
		return nil
	} else {
		return errors.New("primero registre alumnos jeje")
	}
}

func (this *Server) PromedioSub(a string, reply *float64) error {
	sum := 0.000
	mean := 0.000
	cant := 0.000

	if len(Students) > 0 {
		for i := 0; i < len(Students); i++ {
			if Students[i].subject == a {
				sum += Students[i].grade
				cant++
			}
		}
		mean = sum / cant

		*reply = mean
		return nil
	} else {
		return errors.New("primero registre alumnos jeje")
	}
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
