package sprig

import (
	"fmt"
	"path/filepath"
	"strconv"
)

func globMap(input string) map[string]string {
	files, err := filepath.Glob(input)
	if err != nil {
		return map[string]string{
			"_0": fmt.Sprintf("failed to parse glob pattern: %s", err),
		}
	}

	res := make(map[string]string, len(files))
	for i, v := range files {
		res["_"+strconv.Itoa(i)] = v
	}

	return res
}

func glob(input string) []string {
	files, err := filepath.Glob(input)
	if err != nil {
		return []string{
			fmt.Sprintf("failed to parse glob pattern: %s", err),
		}
	}
	return files
}
