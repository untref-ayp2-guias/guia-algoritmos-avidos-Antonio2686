package ejercicios

import (
	"errors"
)

type Farolas struct {
	Posicion  int
	Radio     int
	Encendida bool
}

// EncenderFarolas recibe una lista de farolas y una lista de puntos a iluminar.
// Devuelve una lista de farolas encendidas que iluminan todos los puntos.
// Si no es posible iluminar todos los puntos, devuelve un error.
// La función utiliza un enfoque codicioso para seleccionar las farolas que cubren los puntos.
// Se asume que las farolas están ordenadas por posición y que los puntos también están ordenados.
// La función toma como entrada una lista de farolas y una lista de puntos a iluminar.
// Devuelve una lista de farolas encendidas y un error si no es posible iluminar todos los puntos.

func EncenderFarolas(farolas []Farolas, puntos []int) ([]Farolas, error) {
	encendidas := []Farolas{}
	indiceFarola := 0

	for _, punto := range puntos {
		// Buscar la farola más a la derecha que ilumine el punto
		var mejorFarola *Farolas

		for indiceFarola < len(farolas) && farolas[indiceFarola].Posicion-farolas[indiceFarola].Radio <= punto {
			if farolas[indiceFarola].Posicion+farolas[indiceFarola].Radio >= punto {
				mejorFarola = &farolas[indiceFarola]
			}
			indiceFarola++
		}

		// Si no se encontró una farola para iluminar el punto, el problema no tiene solución
		if mejorFarola == nil {
			return nil, errors.New("EncenderFarolas no debería retornar error")
		}

		// Encender la farola seleccionada si aún no estaba encendida
		if !mejorFarola.Encendida {
			mejorFarola.Encendida = true
			encendidas = append(encendidas, *mejorFarola)
		}
	}

	return encendidas, nil
}
