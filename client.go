package main

import (
	"fmt"
	"net/rpc"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Agregar calificación")
		fmt.Println("2) Obtener el promedio del alumno")
		fmt.Println("3) Obtener el promedio de todos los alumnos")
		fmt.Println("4) Obtener el promedio por materia")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		if op == 1 { //agregar calificación de un alumno por materia
			var name string
			var subject string
			var grade string
			fmt.Print("Name: ")
			fmt.Scanln(&name)
			fmt.Print("Materia: ")
			fmt.Scanln(&subject)
			fmt.Print("Calificación: ")
			fmt.Scanln(&grade)
			slice := []string{name, subject, grade}

			var result string
			err = c.Call("Server.Recibe", slice, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Recibe =", result)
			}
		} else if op == 2 { //obtener el promedio del alumno
			var name string
			fmt.Print("Nombre: ")
			fmt.Scanln(&name)

			var result float64
			err = c.Call("Server.Promedio", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Promedio =", result)
			}
		} else if op == 3 { // obtener el promedio de todos los alumnos
			name := "si"
			var result float64
			err = c.Call("Server.PromedioGral", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioGral =", result)
			}
		} else if op == 4 { // obtener el promedio por materia
			var subject string
			fmt.Print("Materia: ")
			fmt.Scanln(&subject)

			var result float64
			err = c.Call("Server.PromedioSub", subject, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioSub =", result)
			}
		} else if op == 0 {
			return
		}

	}
}

func main() {
	client()
}
