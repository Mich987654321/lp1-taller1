package main

import (
	"fmt"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
//completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		time.Sleep(500 * time.Millisecond)
		ch <- x * x
		//simular trabajo

		ch <- x * x
	}()
	return ch
}

func main() {
	//crea varios futuros y recolecta sus resultados: f1, f2, f3
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(4)
	f3 := asyncCuadrado(6)

	//Opción 1: esperar cada futuro secuencialmente
	fmt.Println("Resultado 1:", <-f1)
	fmt.Println("Resultado 2:", <-f2)
	fmt.Println("Resultado 3:", <-f3)

	//Opción 2: fan-in (combinar múltiples canales)
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.
	fmt.Println("Recolectando resultados mediante Fan-In...")
	todoEnUno := fanIn(f1, f2, f3)
	for res := range todoEnUno {
		fmt.Printf("Recibido futuro: %d\n", res)
	}
	fmt.Println("Todos los futuros recolectados.")

}
