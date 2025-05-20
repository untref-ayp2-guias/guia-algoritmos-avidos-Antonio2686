package ejercicios

import (
	"sort"
)

// SelectorRecursivo selecciona un conjunto máximo de actividades compatibles
// utilizando un enfoque recursivo.
// La función toma una lista de actividades y devuelve una lista de actividades seleccionadas.
type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorRecursivo(actividades []Actividad) []Actividad {
	// Caso base: si no hay actividades, retornar una lista vacía
	if len(actividades) == 0 {
		return []Actividad{}
	}

	// Ordenamos las actividades por tiempo de finalización
	sort.Slice(actividades, func(i, j int) bool {
		return actividades[i].Fin < actividades[j].Fin
	})

	// Seleccionamos la primera actividad
	seleccionada := actividades[0]

	// Filtramos las actividades que comienzan después de que la actividad seleccionada finalice
	var restantes []Actividad
	for _, actividad := range actividades[1:] {
		if actividad.Inicio >= seleccionada.Fin {
			restantes = append(restantes, actividad)
		}
	}

	// Llamada recursiva con el conjunto restante de actividades
	return append([]Actividad{seleccionada}, SelectorRecursivo(restantes)...)
}
