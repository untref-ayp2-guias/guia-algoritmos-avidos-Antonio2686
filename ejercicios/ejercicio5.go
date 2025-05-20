package ejercicios

var billetes = []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

// Cambiar recibe una cantidad y devuelve un mapa con la cantidad de billetes necesarios
// para hacer el cambio. Utiliza un enfoque codicioso para seleccionar los billetes.
// La función toma como entrada una cantidad entera y devuelve un mapa donde la clave es el
// valor del billete y el valor es la cantidad de billetes de ese valor.
// La función utiliza recursión para encontrar la combinación óptima de billetes.
// Si la cantidad es 0, devuelve un mapa vacío.
// Si la cantidad no se puede cambiar con los billetes disponibles, devuelve nil.
// La función asume que la lista de billetes está ordenada de mayor a menor.
// La función utiliza un enfoque recursivo para encontrar la combinación óptima de billetes.

func Cambiar(cantidad int) map[int]int {
	if cantidad == 0 {
		return map[int]int{}
	}

	// Encontrar el billete más grande que aún cabe en la cantidad restante
	for _, billete := range billetes {
		if billete <= cantidad {
			resultado := Cambiar(cantidad - billete)
			resultado[billete]++
			return resultado
		}
	}

	return nil // Esto no debería ocurrir si el sistema de billetes es válido
}
