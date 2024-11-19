package sprig

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlob(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sprig_test")
	assert.NoError(t, err, "Failed to create temporary directory")
	defer os.RemoveAll(tempDir)

	// Create some test files
	files := []string{
		filepath.Join(tempDir, "file1.txt"),
		filepath.Join(tempDir, "file2.log"),
		filepath.Join(tempDir, "file3.txt"),
	}
	for _, file := range files {
		_, err := os.Create(file)
		assert.NoError(t, err, "Failed to create test file")
	}

	tests := []struct {
		name     string
		tpl      string
		expected string
	}{
		{
			name: "Valid glob pattern - Match .txt files",
			tpl:  fmt.Sprintf(`{{ glob "%s/*.txt" }}`, tempDir),
			expected: fmt.Sprintf(`[%s %s]`,
				filepath.Join(tempDir, "file1.txt"),
				filepath.Join(tempDir, "file3.txt"),
			),
		},
		{
			name:     "Valid glob pattern - No matches",
			tpl:      fmt.Sprintf(`{{ glob "%s/*.csv" }}`, tempDir),
			expected: "[]",
		},
		{
			name:     "invalid glob pattern",
			tpl:      `{{ glob "a[" }}`,
			expected: "[failed to parse glob pattern: syntax error in pattern]",
		},
		{
			name: "Valid glob pattern - Match all files",
			tpl:  fmt.Sprintf(`{{ glob "%s/*" }}`, tempDir),
			expected: fmt.Sprintf(`[%s %s %s]`,
				filepath.Join(tempDir, "file1.txt"),
				filepath.Join(tempDir, "file2.log"),
				filepath.Join(tempDir, "file3.txt"),
			),
		},
		{
			name:     "Valid glob pattern - Directory match",
			tpl:      fmt.Sprintf(`{{ glob "%s/nonexistent*" }}`, tempDir),
			expected: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			// result := glob(tt.given)
			res, err := runRaw(tt.tpl, nil)
			assert.NoError(t, err)

			// Validate the result
			assert.Equal(t, tt.expected, res, "Result should match expected output")
		})
	}
}

func TestGlobMap(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sprig_test")
	assert.NoError(t, err, "Failed to create temporary directory")
	defer os.RemoveAll(tempDir)

	// Create some test files
	files := []string{
		filepath.Join(tempDir, "file1.txt"),
		filepath.Join(tempDir, "file2.log"),
		filepath.Join(tempDir, "file3.txt"),
	}
	for _, file := range files {
		_, err := os.Create(file)
		assert.NoError(t, err, "Failed to create test file")
	}

	tests := []struct {
		name     string
		tpl      string
		expected map[string]string
	}{
		{
			name: "Valid glob pattern - Match .txt files",
			tpl:  fmt.Sprintf(`{{ glob "%s/*.txt" }}`, tempDir),
			expected: map[string]string{
				"_0": filepath.Join(tempDir, "file1.txt"),
				"_1": filepath.Join(tempDir, "file3.txt"),
			},
		},
		// {
		// 	name:     "Valid glob pattern - No matches",
		// 	given:    filepath.Join(tempDir, "*.csv"),
		// 	expected: map[string]string{
		// 		// No matches, so the map should be empty
		// 	},
		// },
		// {
		// 	name:  "Invalid glob pattern",
		// 	given: "[invalid][",
		// 	expected: map[string]string{
		// 		"_0": "failed to parse glob pattern: syntax error in pattern",
		// 	},
		// },
		// {
		// 	name:  "Valid glob pattern - Match all files",
		// 	given: filepath.Join(tempDir, "*"),
		// 	expected: map[string]string{
		// 		"_0": filepath.Join(tempDir, "file1.txt"),
		// 		"_1": filepath.Join(tempDir, "file2.log"),
		// 		"_2": filepath.Join(tempDir, "file3.txt"),
		// 	},
		// },
		// {
		// 	name:     "Valid glob pattern - Directory match",
		// 	given:    filepath.Join(tempDir, "nonexistent*"),
		// 	expected: map[string]string{
		// 		// No matches, so the map should be empty
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			// result := glob(tt.given)
			res, err := runRaw(tt.tpl, nil)
			assert.NoError(t, err)

			// Validate the result
			assert.Equal(t, tt.expected, res, "Result should match expected output")
		})
	}
}
