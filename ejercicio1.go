package main


import (
	"fmt"      //imprimir en pantalla
	"sync"     //WaitGroup (sincronización)
	"time"     //pausas con Sleep
)

func main() {
	// Creamos una variable del tipo WaitGroup.
	// hasta que todas las tareas (goroutines) terminen.
	var wg sync.WaitGroup

	//3 tareas diferentes, por eso sumamos 3 al WaitGroup.
	wg.Add(3)

	// ---------- TAREA 1: Imprimir números del 1 al 5 ----------
	go func() {
		// Cuando esta tarea termine, avisamos con Done().
		defer wg.Done()

		// Recorremos del 1 al 5 e imprimimos cada número.
		for i := 1; i <= 5; i++ {
			fmt.Println("Número:", i)
			// Esperamos medio segundo (0.5 segundos)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// ---------- TAREA 2: Imprimir letras de la A a la E ----------
	go func() {
		defer wg.Done()

		// Creamos una lista de letras
		letras := []string{"A", "B", "C", "D", "E"}

		// Recorremos cada letra y la imprimimos
		for _, letra := range letras {
			fmt.Println("Letra:", letra)
			// Esperamos un poquito más, 0.7 segundos
			time.Sleep(700 * time.Millisecond)
		}
	}()

	// ---------- TAREA 3: Mostrar mensaje de estado ----------
	go func() {
		defer wg.Done()

		// Mostramos el mensaje 5 veces
		for i := 0; i < 5; i++ {
			fmt.Println("Procesando...")
			// Esperamos 1 segundo antes de mostrar el siguiente
			time.Sleep(1 * time.Second)
		}
	}()

	// ---------- ESPERAR QUE TERMINEN TODAS LAS TAREAS ----------
	wg.Wait() // Aquí el programa se detiene hasta que las 3 tareas terminen

	// Cuando todas las tareas acaban, se imprime este mensaje
	fmt.Println(" ¡Todas las tareas han terminado!")
}
