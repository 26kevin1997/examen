package main

import (
	"fmt"
	"net/rpc"
	"bufio"
	"os"
)
type Persona struct{
	Nombre string
	Mensaje string
}
func cliente()  {
	c,err:=rpc.Dial("tcp","127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op, result int64
	var persona Persona
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Nombre: ")
	scanner.Scan()
	persona.Nombre = scanner.Text()
	err = c.Call("Server.Conexion", persona.Nombre,&result)
	if err != nil {
		fmt.Println(err)
	}
	go chat()
	for{
		fmt.Println("\t- Menu -")
		fmt.Println("1.- Enviar Mensaje")
		fmt.Println("2.- Enviar Archvio")
		fmt.Println("3.- Mensajes Servidor")
		fmt.Println("4.- Salir")
		fmt.Scan(&op)
		switch op{
		case 1:
			fmt.Println("Mensaje es: ")
			scanner.Scan()
			scanner.Scan()
			persona.Mensaje = scanner.Text()
			err = c.Call("Server.Mensaje", persona,&result)
			if err != nil {
				fmt.Println(err)}
		case 2:
				fmt.Println("Archivo ")
			
		case 3:
			err = c.Call("Server.Todo", "",&result)
			if err != nil {
				fmt.Println("Mensajes en Servidor:\n",err)
			}

			default:
				err = c.Call("Server.Desconectar", persona,&result)
				if err != nil {
					fmt.Println(err)
				}
				return
		
	}	
		}
	}

func chat()  {
	var result int64
	c,err:=rpc.Dial("tcp","127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for{
		err = c.Call("Server.Ret","",&result)
		if err != nil{
			fmt.Println(err)
		}
	}

}
func main()  {
	cliente()
}