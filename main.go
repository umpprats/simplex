package main

import (
	"fmt"
	"log"
	"os"

	"github.com/umpprats/simplex/pkg/simplex"
)

func main() {
	fmt.Println("Simplex Algorithm Solver")
	fmt.Println("========================")
	
	// Check for command line arguments
	if len(os.Args) > 1 && os.Args[1] == "--example" {
		runExampleDemo()
		return
	}
	
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		showHelp()
		return
	}
	
	// Create a new problem instance
	problem := simplex.NewProblem()
	
	// Start data input interface
	err := problem.InputData()
	if err != nil {
		log.Fatalf("Error durante el ingreso de datos: %v", err)
	}
	
	// Validate the entered problem
	if err := problem.ValidateProblem(); err != nil {
		log.Fatalf("Problema inválido: %v", err)
	}
	
	// Display the entered problem
	problem.Display()
	
	// Show problem summary
	fmt.Println("\n" + problem.GetProblemSummary())
	
	fmt.Println("\nNota: La resolución del algoritmo Simplex se implementará en futuras versiones.")
	fmt.Println("Este módulo implementa la integración con la entrada de datos según el issue #5.")
}

func runExampleDemo() {
	fmt.Println("Modo Demo - Problema de Ejemplo")
	fmt.Println("===============================")
	fmt.Println()
	
	problem := simplex.NewProblem()
	problem.SetFromExample()
	
	fmt.Println("Cargando problema de ejemplo del issue #1:")
	fmt.Println("Maximizar: 5x + 3y")
	fmt.Println("Sujeto a:")
	fmt.Println("  2x + y <= 20")
	fmt.Println("  x + y <= 12")
	fmt.Println("  x, y >= 0")
	fmt.Println()
	fmt.Println("Solución esperada: x=8, y=4, valor objetivo=52")
	fmt.Println()
	
	problem.Display()
	fmt.Println("\n" + problem.GetProblemSummary())
}

func showHelp() {
	fmt.Println("Uso del Simplex Algorithm Solver")
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("Opciones:")
	fmt.Println("  ./simplex           Modo interactivo para ingresar datos")
	fmt.Println("  ./simplex --example Ejecutar con problema de ejemplo")
	fmt.Println("  ./simplex --help    Mostrar esta ayuda")
	fmt.Println()
	fmt.Println("El programa implementa la integración con la entrada de datos")
	fmt.Println("para problemas de programación lineal usando el método Simplex.")
	fmt.Println()
	fmt.Println("Características implementadas:")
	fmt.Println("- Entrada interactiva de función objetivo")
	fmt.Println("- Entrada de restricciones con validación")
	fmt.Println("- Validación de datos numéricos")
	fmt.Println("- Visualización clara de ecuaciones")
	fmt.Println("- Restricciones de no negatividad automáticas")
}