package core

import (
	"fmt"
	"regexp"
	"strings"
)

const PseudocodeToPythonPrompt = `# Pseudocode to Python Conversion Prompt

You will convert pseudocode into valid, executable Python 3 code.

Here is the pseudocode you need to convert:

<pseudocode>
{{PSEUDOCODE}}
</pseudocode>

Your task is to interpret this pseudocode and generate Python code that can be executed with ` + "`python file.py`" + `.

## Conversion Requirements

- Convert all comments (whether using "#" or "//") to Python's "#" format
- Handle mixed language syntax (Python, C, etc.) and convert to proper Python syntax
- Convert function definitions to Python's "def" syntax with proper indentation
- Convert control structures (if/else, loops, etc.) to Python syntax
- Convert data types to appropriate Python equivalents
- Use Python's print() function for output statements
- Use Python's input() function for input operations when appropriate
- Only use imports from Python's standard library
- Ensure the code is executable without syntax errors
- Preserve the original logic and functionality
- Make reasonable assumptions for ambiguous pseudocode elements
- Include appropriate error handling if the pseudocode suggests it
- Generate fully correct, working Python code

## Process

First, analyze the pseudocode systematically in <conversion_analysis> tags. In your analysis:

1. Go through the pseudocode line by line, identifying what each line contains and what specific conversions are needed
2. List all syntax transformations required (e.g., function definitions, variable declarations, control structures, operators, etc.)
3. Note any data type conversions needed and what Python equivalents you'll use
4. Identify any input/output operations and plan the appropriate Python functions
5. Note any ambiguous parts and clearly state your assumptions for handling them
6. Plan the overall structure and indentation of the final Python code

**Be concise but thorough in your analysis.**

After your analysis, provide the converted Python code in <code> tags.

## Output Format

` + "```" + `
<conversion_analysis>
[Your systematic line-by-line analysis of the pseudocode and detailed conversion plan]
</conversion_analysis>

<code>
[Your converted Python code here]
</code>
` + "```" + `
`

// ExtractPythonCode parses the LLM response and extracts the Python code from <code> tags
func ExtractPythonCode(response string) (string, error) {
	re := regexp.MustCompile(`(?s)<code>\s*(.*?)\s*</code>`)
	matches := re.FindStringSubmatch(response)

	if len(matches) < 2 {
		return "", fmt.Errorf("no <code> tags found in response")
	}

	code := strings.TrimSpace(matches[1])
	if code == "" {
		return "", fmt.Errorf("code block is empty")
	}

	return code, nil
}

// BuildPseudocodePrompt replaces the {{PSEUDOCODE}} placeholder with actual input
func BuildPseudocodePrompt(pseudocode string) string {
	return strings.Replace(PseudocodeToPythonPrompt, "{{PSEUDOCODE}}", pseudocode, 1)
}
