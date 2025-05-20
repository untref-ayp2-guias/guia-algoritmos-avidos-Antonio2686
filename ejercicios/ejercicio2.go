package ejercicios

import (
	"fmt"
	"sort"
)

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	// Ordenamos las tareas por duración (tiempo de ejecución)
	sort.Slice(tareas, func(i, j int) bool {
		return tareas[i].Tiempo < tareas[j].Tiempo
	})

	// Variables para calcular los tiempos de finalización
	tiempoActual := 0
	var tiemposFinalizacion []int

	// Procesamos las tareas en el orden óptimo
	for _, tarea := range tareas {
		tiempoActual += tarea.Tiempo
		tiemposFinalizacion = append(tiemposFinalizacion, tiempoActual)
		fmt.Printf("Tarea %s finaliza en %d\n", tarea.Nombre, tiempoActual)
	}

	// Calculamos el tiempo medio de finalización
	sumaFinalizaciones := 0
	for _, tf := range tiemposFinalizacion {
		sumaFinalizaciones += tf
	}
	promedioFinalizacion := float64(sumaFinalizaciones) / float64(len(tareas))

	fmt.Printf("Tiempo medio de finalización: %.2f\n", promedioFinalizacion)
}
