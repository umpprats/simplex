package simplex

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Variable represents a variable in the linear programming problem
type Variable struct {
	Name        string
	Coefficient float64
}

// Constraint represents a constraint in the linear programming problem
type Constraint struct {
	Variables []Variable
	Operator  string // "<=", ">=", "="
	Value     float64
}

// ObjectiveFunction represents the objective function to maximize or minimize
type ObjectiveFunction struct {
	Variables []Variable
	Type      string // "max" or "min"
}

// Problem represents a complete linear programming problem
type Problem struct {
	Objective   ObjectiveFunction
	Constraints []Constraint
	Variables   []string
}

// NewProblem creates a new empty linear programming problem
func NewProblem() *Problem {
	return &Problem{
		Constraints: make([]Constraint, 0),
		Variables:   make([]string, 0),
	}
}

// InputData prompts the user to enter all problem data
func (p *Problem) InputData() error {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("\nIngreso de datos para el problema de programación lineal")
	fmt.Println("=========================================================")
	
	// Input objective function
	if err := p.inputObjectiveFunction(scanner); err != nil {
		return fmt.Errorf("error al ingresar función objetivo: %v", err)
	}
	
	// Input constraints
	if err := p.inputConstraints(scanner); err != nil {
		return fmt.Errorf("error al ingresar restricciones: %v", err)
	}
	
	return nil
}

// inputObjectiveFunction prompts user to enter the objective function
func (p *Problem) inputObjectiveFunction(scanner *bufio.Scanner) error {
	fmt.Println("\n1. Función Objetivo")
	fmt.Println("-------------------")
	
	// Get objective type
	for {
		fmt.Print("¿Desea maximizar o minimizar? (max/min): ")
		if scanner.Scan() {
			objType := strings.ToLower(strings.TrimSpace(scanner.Text()))
			if objType == "max" || objType == "min" {
				p.Objective.Type = objType
				break
			}
		}
		fmt.Println("Por favor, ingrese 'max' o 'min'")
	}
	
	// Get variables and coefficients
	fmt.Println("\nIngrese la función objetivo.")
	fmt.Println("Ejemplo: Para 5x + 3y, ingrese los coeficientes cuando se soliciten")
	
	variables, err := p.inputVariables(scanner, "función objetivo")
	if err != nil {
		return err
	}
	
	p.Objective.Variables = variables
	return nil
}

// inputConstraints prompts user to enter all constraints
func (p *Problem) inputConstraints(scanner *bufio.Scanner) error {
	fmt.Println("\n2. Restricciones")
	fmt.Println("----------------")
	
	constraintNum := 1
	for {
		fmt.Printf("\nRestricción %d:\n", constraintNum)
		
		// Get variables for this constraint
		variables, err := p.inputVariables(scanner, fmt.Sprintf("restricción %d", constraintNum))
		if err != nil {
			return err
		}
		
		// Get operator
		operator, err := p.inputOperator(scanner)
		if err != nil {
			return err
		}
		
		// Get constraint value
		value, err := p.inputFloat(scanner, "valor del lado derecho")
		if err != nil {
			return err
		}
		
		constraint := Constraint{
			Variables: variables,
			Operator:  operator,
			Value:     value,
		}
		
		p.Constraints = append(p.Constraints, constraint)
		
		// Ask if user wants to add more constraints
		fmt.Print("\n¿Desea agregar otra restricción? (s/n): ")
		if scanner.Scan() {
			response := strings.ToLower(strings.TrimSpace(scanner.Text()))
			if response != "s" && response != "si" && response != "sí" {
				break
			}
		}
		constraintNum++
	}
	
	return nil
}

// inputVariables prompts user to enter variables and their coefficients
func (p *Problem) inputVariables(scanner *bufio.Scanner, context string) ([]Variable, error) {
	if len(p.Variables) == 0 {
		// First time - ask for variable names
		fmt.Print("Ingrese los nombres de las variables separados por espacios (ej: x y z): ")
		if scanner.Scan() {
			varNames := strings.Fields(scanner.Text())
			if len(varNames) == 0 {
				return nil, fmt.Errorf("debe ingresar al menos una variable")
			}
			p.Variables = varNames
		}
	}
	
	var variables []Variable
	fmt.Printf("\nIngrese los coeficientes para la %s:\n", context)
	
	for _, varName := range p.Variables {
		coeff, err := p.inputFloat(scanner, fmt.Sprintf("coeficiente de %s", varName))
		if err != nil {
			return nil, err
		}
		
		variables = append(variables, Variable{
			Name:        varName,
			Coefficient: coeff,
		})
	}
	
	return variables, nil
}

// inputOperator prompts user to enter constraint operator
func (p *Problem) inputOperator(scanner *bufio.Scanner) (string, error) {
	for {
		fmt.Print("Ingrese el operador (<=, >=, =): ")
		if scanner.Scan() {
			operator := strings.TrimSpace(scanner.Text())
			if operator == "<=" || operator == ">=" || operator == "=" {
				return operator, nil
			}
		}
		fmt.Println("Por favor, ingrese un operador válido: <=, >=, o =")
	}
}

// inputFloat prompts user to enter a floating point number with validation
func (p *Problem) inputFloat(scanner *bufio.Scanner, prompt string) (float64, error) {
	for {
		fmt.Printf("Ingrese %s: ", prompt)
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if value, err := strconv.ParseFloat(input, 64); err == nil {
				return value, nil
			}
		}
		fmt.Println("Por favor, introduce solo valores numéricos")
	}
}

// Display shows the complete problem in a clear format
func (p *Problem) Display() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("PROBLEMA DE PROGRAMACIÓN LINEAL INGRESADO")
	fmt.Println(strings.Repeat("=", 60))
	
	// Display objective function
	fmt.Printf("\nFunción Objetivo (%s):\n", strings.ToUpper(p.Objective.Type))
	fmt.Print("  ")
	p.displayEquation(p.Objective.Variables)
	fmt.Println()
	
	// Display constraints
	fmt.Println("\nRestricciones:")
	for i, constraint := range p.Constraints {
		fmt.Printf("  %d. ", i+1)
		p.displayEquation(constraint.Variables)
		fmt.Printf(" %s %.2f\n", constraint.Operator, constraint.Value)
	}
	
	// Display non-negativity constraints
	fmt.Println("\nRestricciones de no negatividad:")
	for _, varName := range p.Variables {
		fmt.Printf("  %s >= 0\n", varName)
	}
	
	fmt.Println(strings.Repeat("=", 60))
}

// displayEquation formats and displays a linear equation
func (p *Problem) displayEquation(variables []Variable) {
	for i, variable := range variables {
		if i > 0 {
			if variable.Coefficient >= 0 {
				fmt.Print(" + ")
			} else {
				fmt.Print(" ")
			}
		}
		
		if variable.Coefficient == 1 {
			fmt.Printf("%s", variable.Name)
		} else if variable.Coefficient == -1 {
			fmt.Printf("-%s", variable.Name)
		} else {
			fmt.Printf("%.2g%s", variable.Coefficient, variable.Name)
		}
	}
}