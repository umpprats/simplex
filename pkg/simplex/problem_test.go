package simplex

import (
	"strconv"
	"strings"
	"testing"
)

func TestNewProblem(t *testing.T) {
	problem := NewProblem()
	
	if problem == nil {
		t.Fatal("NewProblem() returned nil")
	}
	
	if len(problem.Constraints) != 0 {
		t.Errorf("Expected empty constraints, got %d", len(problem.Constraints))
	}
	
	if len(problem.Variables) != 0 {
		t.Errorf("Expected empty variables, got %d", len(problem.Variables))
	}
}

func TestDisplayEquation(t *testing.T) {
	problem := NewProblem()
	
	// Capture output using a string builder approach
	variables := []Variable{
		{Name: "x", Coefficient: 5},
		{Name: "y", Coefficient: 3},
		{Name: "z", Coefficient: -2},
	}
	
	// Test that displayEquation doesn't panic
	problem.displayEquation(variables)
}

func TestVariableCreation(t *testing.T) {
	variable := Variable{
		Name:        "x",
		Coefficient: 5.5,
	}
	
	if variable.Name != "x" {
		t.Errorf("Expected variable name 'x', got '%s'", variable.Name)
	}
	
	if variable.Coefficient != 5.5 {
		t.Errorf("Expected coefficient 5.5, got %f", variable.Coefficient)
	}
}

func TestConstraintCreation(t *testing.T) {
	variables := []Variable{
		{Name: "x", Coefficient: 2},
		{Name: "y", Coefficient: 1},
	}
	
	constraint := Constraint{
		Variables: variables,
		Operator:  "<=",
		Value:     20,
	}
	
	if len(constraint.Variables) != 2 {
		t.Errorf("Expected 2 variables, got %d", len(constraint.Variables))
	}
	
	if constraint.Operator != "<=" {
		t.Errorf("Expected operator '<=', got '%s'", constraint.Operator)
	}
	
	if constraint.Value != 20 {
		t.Errorf("Expected value 20, got %f", constraint.Value)
	}
}

func TestObjectiveFunctionCreation(t *testing.T) {
	variables := []Variable{
		{Name: "x", Coefficient: 5},
		{Name: "y", Coefficient: 3},
	}
	
	objective := ObjectiveFunction{
		Variables: variables,
		Type:      "max",
	}
	
	if len(objective.Variables) != 2 {
		t.Errorf("Expected 2 variables, got %d", len(objective.Variables))
	}
	
	if objective.Type != "max" {
		t.Errorf("Expected type 'max', got '%s'", objective.Type)
	}
}

func TestProblemWithSampleData(t *testing.T) {
	// Test creating a problem similar to the one described in issue #1
	problem := NewProblem()
	
	// Set up objective function: 5x + 3y (max)
	problem.Objective = ObjectiveFunction{
		Variables: []Variable{
			{Name: "x", Coefficient: 5},
			{Name: "y", Coefficient: 3},
		},
		Type: "max",
	}
	
	// Set up constraints: 2x + y <= 20, x + y <= 12
	problem.Constraints = []Constraint{
		{
			Variables: []Variable{
				{Name: "x", Coefficient: 2},
				{Name: "y", Coefficient: 1},
			},
			Operator: "<=",
			Value:    20,
		},
		{
			Variables: []Variable{
				{Name: "x", Coefficient: 1},
				{Name: "y", Coefficient: 1},
			},
			Operator: "<=",
			Value:    12,
		},
	}
	
	problem.Variables = []string{"x", "y"}
	
	// Verify the problem structure
	if problem.Objective.Type != "max" {
		t.Errorf("Expected objective type 'max', got '%s'", problem.Objective.Type)
	}
	
	if len(problem.Constraints) != 2 {
		t.Errorf("Expected 2 constraints, got %d", len(problem.Constraints))
	}
	
	if len(problem.Variables) != 2 {
		t.Errorf("Expected 2 variables, got %d", len(problem.Variables))
	}
	
	// Test the display function (should not panic)
	problem.Display()
}

func TestDisplayFunction(t *testing.T) {
	problem := NewProblem()
	
	// Set up a simple problem
	problem.Objective = ObjectiveFunction{
		Variables: []Variable{
			{Name: "x", Coefficient: 1},
			{Name: "y", Coefficient: 1},
		},
		Type: "max",
	}
	
	problem.Constraints = []Constraint{
		{
			Variables: []Variable{
				{Name: "x", Coefficient: 1},
				{Name: "y", Coefficient: 1},
			},
			Operator: "<=",
			Value:    10,
		},
	}
	
	problem.Variables = []string{"x", "y"}
	
	// This should not panic
	problem.Display()
}

// Test helper functions for input validation
func TestValidateNumericInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"5", true},
		{"5.5", true},
		{"-3", true},
		{"-3.14", true},
		{"0", true},
		{"abc", false},
		{"5x", false},
		{"", false},
		{" ", false},
	}
	
	for _, tc := range testCases {
		isValid := isNumericInput(tc.input)
		if isValid != tc.expected {
			t.Errorf("Input '%s': expected %v, got %v", tc.input, tc.expected, isValid)
		}
	}
}

// Helper function to validate numeric input (used for testing)
func isNumericInput(input string) bool {
	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}
	
	// Try to parse as float
	_, err := parseFloat(input)
	return err == nil
}

// Helper function to parse float (used for testing)  
func parseFloat(s string) (float64, error) {
	// Use the actual strconv.ParseFloat for proper validation
	// Import strconv at the top of the file
	return strconv.ParseFloat(s, 64)
}