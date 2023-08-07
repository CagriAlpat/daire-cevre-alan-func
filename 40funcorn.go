package main

 import (
 	"fmt"
 	"math"
 )

 const PI = 3.14 // sabit değer atadık

 func daire(yaricap float64) (alan float64, cevre float64) {
 	alan = PI * (math.Pow(yaricap, 2.0))
 	cevre = PI * 2 * yaricap
 	return
 }
 func main() {
 	dairealan, dairecevre := daire(5.0) //dairealan = alan ' eşit oldu (returne).       dairecevre = cevre (returne )

 	fmt.Println("daire alanı = ", dairealan)
 	fmt.Println("daire cevre = ", dairecevre)

 }
