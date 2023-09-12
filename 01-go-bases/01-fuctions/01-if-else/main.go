package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	
	

	if t := time.Now(); t.Hour()< 12 {
		fmt.Println(t.Hour(),"¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println(t.Hour(),"¡Tarde!")
	} else {
		fmt.Println(t.Hour(),"¡Noche!")
	}

	/*
		t := time.Now()
	hour := t.Hour()

	if hour < 12 {
		fmt.Println("¡Mañana!")
	} else if hour < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}

	*/

	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> MacOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}
}