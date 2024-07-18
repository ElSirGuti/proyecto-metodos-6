package main

import (
	"fmt"
	"math"
	"os/exec"
)

// objectiveFunction calcula el valor de la función f(x) = (x-3)^2
func objectiveFunction(x float64) float64 {
	return math.Pow(x-3, 2)
}

// derivative calcula la derivada de la función f'(x) = 2(x-3)
func derivative(x float64) float64 {
	return 2 * (x - 3)
}

// gradientDescent implementa el algoritmo de descenso de gradiente
func gradientDescent(x0, alpha float64, maxIter int) (float64, float64) {
	x := x0
	for i := 0; i < maxIter; i++ {
		grad := derivative(x)
		x = x - alpha*grad
	}
	return x, objectiveFunction(x)
}

func main() {
	// Casos de prueba
	// Quise probar con varios a ver si por alguna razon daba distinto con otros parámetros, luego me acordé que no sé de matemáticas y parece que siempre va a dar igual
	casosDePrueba := []struct {
		x0      float64 // El punto inicial
		alpha   float64 // El tamaño de paso
		maxIter int     // El número máximo de iteraciones
	}{ // {x0, alpha, maxIter}
		{0.0, 0.1, 1000},     // Caso 1
		{5.0, 0.01, 500},     // Caso 2
		{-10.0, 0.05, 2000},  // Caso 3
		{10.0, 0.001, 10000}, // Caso 4
		{2.0, 0.5, 100},      // Caso 5
	}

	// Ejecutar cada caso de prueba
	for i, caso := range casosDePrueba {
		minX, minVal := gradientDescent(caso.x0, caso.alpha, caso.maxIter)
		fmt.Printf("Caso de Prueba %d:\n", i+1)
		fmt.Printf("Punto Inicial (x0): %.2f, Tamaño de Paso (alpha): %.3f, Número Máximo de Iteraciones: %d\n", caso.x0, caso.alpha, caso.maxIter)
		fmt.Printf("El valor de x que minimiza la función es: %.6f\n", minX)
		fmt.Printf("El valor mínimo de la función en ese punto es: %.6f\n", minVal)
		fmt.Println("")
		fmt.Println()
	}

	videoPath := "sorpresa.mp4" // Nombre del archivo de video en la carpeta raíz del script
	err := playVideo(videoPath)
	if err != nil {
		fmt.Println("Error al reproducir el video:", err)
	}
}

func playVideo(videoPath string) error {
	cmd := exec.Command("cmd", "/c", "start", videoPath)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
