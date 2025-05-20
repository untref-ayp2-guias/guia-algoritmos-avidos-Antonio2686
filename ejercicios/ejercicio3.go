package ejercicios

import (
	"fmt"
	"sort"
)

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

// Ejercicio3 resuelve el problema de la mochila fraccionaria.
// Dada una lista de objetos con un peso y un valor, y una capacidad máxima de la mochila,
// devuelve un mapa que indica qué fracción de cada objeto se incluye en la mochila.
// La función utiliza un enfoque codicioso para seleccionar los objetos de mayor valor por unidad de peso.
// La función toma una lista de objetos y la capacidad máxima de la mochila como argumentos.

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	// Ordenar los objetos por valor por unidad de peso (de mayor a menor)
	sort.Slice(objetos, func(i, j int) bool {
		return float64(objetos[i].Valor)/float64(objetos[i].Peso) >
			float64(objetos[j].Valor)/float64(objetos[j].Peso)
	})

	resultado := make(map[Objeto]float32)
	pesoActual := 0

	for _, obj := range objetos {
		if pesoActual+obj.Peso <= capacidad {
			// Se puede incluir el objeto completo
			resultado[obj] = 1.0
			pesoActual += obj.Peso
		} else {
			// Se incluye solo una fracción del objeto
			fraccion := float32(capacidad-pesoActual) / float32(obj.Peso)
			resultado[obj] = fraccion
			break // La mochila ya está llena
		}
	}

	// Imprimir resultado
	fmt.Println("Objetos seleccionados:")
	for obj, fraccion := range resultado {
		fmt.Printf("- %s (%.2f%% del objeto)\n", obj.Nombre, fraccion*100)
	}

	return resultado
}
