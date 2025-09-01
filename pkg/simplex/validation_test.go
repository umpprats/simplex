package simplex

import (
	"testing"
)

func TestValidateProblem(t *testing.T) {
	// Test valid problem
	problem := NewProblem()
	problem.SetFromExample()
	
	err := problem.ValidateProblem()
	if err != nil {
		t.Errorf("Expected valid problem, got error: %v", err)
	}
}

func TestValidateProblemEmpty(t *testing.T) {
	// Test empty problem
	problem := NewProblem()
	
	err := problem.ValidateProblem()
	if err == nil {
		t.Error("Expected error for empty problem, got nil")
	}
}

func TestValidateProblemInvalidObjectiveType(t *testing.T) {
	problem := NewProblem()
	problem.Objective = ObjectiveFunction{
		Variables: []Variable{{Name: "x", Coefficient: 1}},
		Type:      "invalid",
	}
	problem.Variables = []string{"x"}
	problem.Constraints = []Constraint{
		{
			Variables: []Variable{{Name: "x", Coefficient: 1}},
			Operator:  "<=",
			Value:     10,
		},
	}
	
	err := problem.ValidateProblem()
	if err == nil {
		t.Error("Expected error for invalid objective type, got nil")
	}
}

func TestGetProblemSummary(t *testing.T) {
	problem := NewProblem()
	problem.SetFromExample()
	
	summary := problem.GetProblemSummary()
	if summary == "" {
		t.Error("Expected non-empty summary")
	}
	
	// Check that summary contains expected elements
	if !contains(summary, "MAX") {
		t.Error("Summary should contain objective type")
	}
	
	if !contains(summary, "Variables: 2") {
		t.Error("Summary should contain variable count")
	}
	
	if !contains(summary, "Restricciones: 2") {
		t.Error("Summary should contain constraint count")
	}
}

func TestSetFromExample(t *testing.T) {
	problem := NewProblem()
	problem.SetFromExample()
	
	// Verify objective function
	if problem.Objective.Type != "max" {
		t.Errorf("Expected objective type 'max', got '%s'", problem.Objective.Type)
	}
	
	if len(problem.Objective.Variables) != 2 {
		t.Errorf("Expected 2 objective variables, got %d", len(problem.Objective.Variables))
	}
	
	// Verify constraints
	if len(problem.Constraints) != 2 {
		t.Errorf("Expected 2 constraints, got %d", len(problem.Constraints))
	}
	
	// Verify variables
	if len(problem.Variables) != 2 {
		t.Errorf("Expected 2 variables, got %d", len(problem.Variables))
	}
	
	expectedVars := map[string]bool{"x": true, "y": true}
	for _, varName := range problem.Variables {
		if !expectedVars[varName] {
			t.Errorf("Unexpected variable: %s", varName)
		}
	}
}

func TestExportToJSON(t *testing.T) {
	problem := NewProblem()
	problem.SetFromExample()
	
	json, err := problem.ExportToJSON()
	if err != nil {
		t.Errorf("Unexpected error in ExportToJSON: %v", err)
	}
	
	if json == "" {
		t.Error("Expected non-empty JSON output")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s != substr && s[0:len(substr)] == substr || 
		   len(s) > len(substr) && findSubstring(s, substr)
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}