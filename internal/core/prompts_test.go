package core

import (
	"strings"
	"testing"
)

func TestExtractPythonCode(t *testing.T) {
	tests := []struct {
		name        string
		response    string
		wantCode    string
		wantErr     bool
		errContains string
	}{
		{
			name: "valid code extraction",
			response: `<conversion_analysis>
Some analysis here
</conversion_analysis>

<code>
print("Hello, World!")
</code>`,
			wantCode: `print("Hello, World!")`,
			wantErr:  false,
		},
		{
			name: "code with whitespace",
			response: `<code>
  def greet(name):
      print(f"Hello, {name}!")

  greet("World")
</code>`,
			wantCode: `def greet(name):
      print(f"Hello, {name}!")

  greet("World")`,
			wantErr: false,
		},
		{
			name: "code with special characters",
			response: `<code>
import re
pattern = r'\d+'
data = "test123"
</code>`,
			wantCode: `import re
pattern = r'\d+'
data = "test123"`,
			wantErr: false,
		},
		{
			name:        "missing code tags",
			response:    "Some response without code tags",
			wantCode:    "",
			wantErr:     true,
			errContains: "no <code> tags found",
		},
		{
			name: "empty code block",
			response: `<code>

</code>`,
			wantCode:    "",
			wantErr:     true,
			errContains: "code block is empty",
		},
		{
			name: "multiple code blocks extracts first",
			response: `<code>
print("first")
</code>

<code>
print("second")
</code>`,
			wantCode: `print("first")`,
			wantErr:  false,
		},
		{
			name: "code with newlines and indentation",
			response: `<code>
for i in range(10):
    if i % 2 == 0:
        print(i)
</code>`,
			wantCode: `for i in range(10):
    if i % 2 == 0:
        print(i)`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractPythonCode(tt.response)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ExtractPythonCode() expected error but got none")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("ExtractPythonCode() error = %v, want error containing %q", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("ExtractPythonCode() unexpected error = %v", err)
				return
			}

			if got != tt.wantCode {
				t.Errorf("ExtractPythonCode() = %q, want %q", got, tt.wantCode)
			}
		})
	}
}

func TestBuildPseudocodePrompt(t *testing.T) {
	tests := []struct {
		name       string
		pseudocode string
	}{
		{
			name:       "simple pseudocode",
			pseudocode: "print hello world",
		},
		{
			name:       "empty string",
			pseudocode: "",
		},
		{
			name: "multiline pseudocode",
			pseudocode: `function greet(name):
    print "Hello, " + name`,
		},
		{
			name:       "special characters",
			pseudocode: `var x = "test's \"quoted\" string"`,
		},
		{
			name:       "pseudocode with braces",
			pseudocode: "function test() { return true; }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildPseudocodePrompt(tt.pseudocode)

			if !strings.Contains(got, tt.pseudocode) {
				t.Errorf("BuildPseudocodePrompt() result does not contain input pseudocode")
			}

			if strings.Contains(got, "{{PSEUDOCODE}}") {
				t.Errorf("BuildPseudocodePrompt() still contains {{PSEUDOCODE}} placeholder")
			}

			if !strings.Contains(got, "Pseudocode to Python Conversion Prompt") {
				t.Errorf("BuildPseudocodePrompt() does not contain expected prompt header")
			}
		})
	}
}

func TestBuildPseudocodePromptReplacesOnlyOnce(t *testing.T) {
	pseudocode := "test {{PSEUDOCODE}} in pseudocode"
	result := BuildPseudocodePrompt(pseudocode)

	count := strings.Count(result, pseudocode)
	if count != 1 {
		t.Errorf("BuildPseudocodePrompt() replaced pseudocode %d times, want 1", count)
	}
}
