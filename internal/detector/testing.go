package detector

import (
	"path/filepath"
	"strconv"
	"strings"
)

type testingDetector struct{}

func (t testingDetector) Name() string {
	return "testing detector"
}
func (t testingDetector) Category() string {
	return "testing"
}
func (t testingDetector) Detect(files []string) []Evidence {
	var evidence []Evidence
	var testfiles []string
	for _, path := range files {
		filename := strings.ToLower(filepath.Base(path))
		isTestFile := strings.HasSuffix(filename, "_test.go") ||
			strings.HasPrefix(filename, "test_") ||
			strings.HasSuffix(filename, "_test.py") ||
			strings.HasSuffix(filename, ".test.js") ||
			strings.HasSuffix(filename, ".test.ts") ||
			strings.HasSuffix(filename, ".spec.js") ||
			strings.HasSuffix(filename, ".spec.ts") ||
			strings.Contains(path, "src/test/java")
			if isTestFile{
				testfiles=append(testfiles,path)
			}

	}
	if len(testfiles)==0{
		return evidence
	}
	desc:="found"+strconv.Itoa(len(testfiles))+"test file(s),e.g"+testfiles[0]
	evidence = append(evidence, Evidence{
		File: testfiles[0],
		Description: desc,
		Confidence: "medium",
		Category: "testing",
	})
	return  evidence

}