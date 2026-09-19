package main

import "fmt"
var nome string
var Batman string

func main () {

     fmt.Println("Diga o seu nome:" ,nome)
     fmt.Scan(&nome)

     if nome == "Batman" { 
     fmt.Println("Olá,Batman!!!!!")
     } else {
     fmt.Println("Olá" ,nome)

}


}
