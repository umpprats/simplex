package simplex

import (
	"fmt"
	"strings"
)

// ValidateProblem checks if the entered problem is valid
func (p *Problem) ValidateProblem() error {
	// Check if objective function is set
	if len(p.Objective.Variables) == 0 {
		return fmt.Errorf("función objetivo no definida")
	}
	
	// Check if objective type is valid
	if p.Objective.Type != "max" && p.Objective.Type != "min" {
		return fmt.Errorf("tipo de objetivo inválido: debe ser 'max' o 'min'")
	}
	
	// Check if constraints exist
	if len(p.Constraints) == 0 {
		return fmt.Errorf("no se han definido restricciones")
	}
	
	// Check if variables are consistent across objective and constraints
	objVars := make(map[string]bool)
	for _, v := range p.Objective.Variables {
		objVars[v.Name] = true
	}
	
	for i, constraint := range p.Constraints {
		for _, v := range constraint.Variables {
			if !objVars[v.Name] {
				return fmt.Errorf("variable '%s' en restricción %d no está en la función objetivo", v.Name, i+1)
			}
		}
	}
	
	return nil
}

// GetProblemSummary returns a text summary of the problem
func (p *Problem) GetProblemSummary() string {
	var summary strings.Builder
	
	summary.WriteString("Resumen del Problema:\n")
	summary.WriteString("====================\n")
	
	// Objective function summary
	summary.WriteString(fmt.Sprintf("Objetivo: %s ", strings.ToUpper(p.Objective.Type)))
	for i, v := range p.Objective.Variables {
		if i > 0 {
			if v.Coefficient >= 0 {
				summary.WriteString(" + ")
			} else {
				summary.WriteString(" ")
			}
		}
		summary.WriteString(fmt.Sprintf("%.2g%s", v.Coefficient, v.Name))
	}
	summary.WriteString("\n")
	
	// Variables count
	summary.WriteString(fmt.Sprintf("Variables: %d (%s)\n", len(p.Variables), strings.Join(p.Variables, ", ")))
	
	// Constraints count
	summary.WriteString(fmt.Sprintf("Restricciones: %d\n", len(p.Constraints)))
	
	return summary.String()
}

// ExportToJSON exports the problem to JSON format (for future use)
func (p *Problem) ExportToJSON() (string, error) {
	// This is a placeholder for future JSON export functionality
	// to support the "Save and Load Problems" feature from issue #4
	return `{"message": "JSON export feature coming soon"}`, nil
}

// SetFromExample creates the example problem from issue #1
func (p *Problem) SetFromExample() {
	// Clear existing data
	p.Constraints = make([]Constraint, 0)
	p.Variables = []string{"x", "y"}
	
	// Set objective function: maximize 5x + 3y
	p.Objective = ObjectiveFunction{
		Variables: []Variable{
			{Name: "x", Coefficient: 5},
			{Name: "y", Coefficient: 3},
		},
		Type: "max",
	}
	
	// Add constraints: 2x + y <= 20, x + y <= 12
	p.Constraints = append(p.Constraints, Constraint{
		Variables: []Variable{
			{Name: "x", Coefficient: 2},
			{Name: "y", Coefficient: 1},
		},
		Operator: "<=",
		Value:    20,
	})
	
	p.Constraints = append(p.Constraints, Constraint{
		Variables: []Variable{
			{Name: "x", Coefficient: 1},
			{Name: "y", Coefficient: 1},
		},
		Operator: "<=",
		Value:    12,
	})
}