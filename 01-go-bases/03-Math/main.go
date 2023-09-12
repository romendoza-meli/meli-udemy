package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	jugar()
}

func jugar() {
	numeroAleatorio := rand.Intn(100) + 1
	var numeroIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numeroIngresado)

		if numeroIngresado == numeroAleatorio {
			fmt.Println("¡Felicitaciones, adivinaste el número!")
			jugarNuevamente()
			return
		} else if numeroIngresado < numeroAleatorio {
			fmt.Println("El número a adivinar es mayor.")
		} else if numeroIngresado > numeroAleatorio {
			fmt.Println("El número a adivinar es menor.")
		}
	}

	fmt.Println("Se acabaron los intentos. El número era:", numeroAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Print("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Elección inválida. Inténtalo nuevamente.")
		jugarNuevamente()
	}
}