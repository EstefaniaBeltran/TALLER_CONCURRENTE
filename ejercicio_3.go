package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Estructura Tarea con los campos requeridos
type Tarea struct {
	id        int
	prioridad int           // 1 = alta, 3 = baja
	duracion  time.Duration // tiempo de ejecución simulado
}

// Función que ejecuta una tarea y reporta su finalización
func ejecutarTarea(t Tarea, resultados chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la tarea como completada al terminar
	time.Sleep(t.duracion)
	mensaje := fmt.Sprintf("Tarea %d (Prioridad %d) completada en %v", t.id, t.prioridad, t.duracion)
	resultados <- mensaje
}

// Función principal
func main() {
	rand.Seed(time.Now().UnixNano())

	// Crear una lista de tareas con distintas prioridades y duraciones aleatorias
	tareas := []Tarea{
		{1, 2, time.Duration(rand.Intn(1000)+500) * time.Millisecond},
		{2, 1, time.Duration(rand.Intn(1000)+500) * time.Millisecond},
		{3, 3, time.Duration(rand.Intn(1000)+500) * time.Millisecond},
		{4, 1, time.Duration(rand.Intn(1000)+500) * time.Millisecond},
		{5, 2, time.Duration(rand.Intn(1000)+500) * time.Millisecond},
	}

	//Canales de comunicación
	canalTareas := make(chan Tarea)
	resultados := make(chan string)

	var wg sync.WaitGroup

	// Goroutine planificadora: envía las tareas por prioridad
	go func() {
		// Ordenar las tareas por prioridad (1 = más alta)
		sort.Slice(tareas, func(i, j int) bool {
			return tareas[i].prioridad < tareas[j].prioridad
		})

		fmt.Println("📋 Planificador: enviando tareas por prioridad...")
		for _, t := range tareas {
			fmt.Printf("Enviando tarea %d (Prioridad %d)\n", t.id, t.prioridad)
			canalTareas <- t
		}
		close(canalTareas) // Indica que no hay más tareas
	}()

	// Lanzar varias goroutines worker para ejecutar tareas concurrentemente
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			for t := range canalTareas {
				fmt.Printf("Worker %d ejecutando tarea %d...\n", id, t.id)
				wg.Add(1)
				ejecutarTarea(t, resultados, &wg)
			}
		}(i)
	}

	// Goroutine para mostrar los resultados a medida que llegan
	go func() {
		for msg := range resultados {
			fmt.Println(msg)
		}
	}()

	// Esperar que todas las tareas finalicen
	wg.Wait()
	close(resultados)

	fmt.Println("\n Todas las tareas han sido completadas.")
}