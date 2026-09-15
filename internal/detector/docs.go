package detector

import "path/filepath"

var docFiles = map[string]string{
	"SECURITY.md":     "security policy documentation",
	"CONTRIBUTING.md": "contribution guidelines",
	"LICENSE":         "license file",
	"LICENSE.md":      "license file",
	"LICENSE.txt":     "license file",
}

type docsDetector struct{}

func (d docsDetector) Name() string {
	return " security documentation  detector"
}
func (d docsDetector) Category() string {
	return "docs"
}
func (d docsDetector) Detect(files []string) []Evidence {
	var evidence []Evidence
	for _, path := range files {
		filename := filepath.Base(path)
		if desc,found:=docFiles[filename];found{
			evidence = append(evidence, Evidence{
				File: path,
				Description: desc+"found",
				Confidence: "high",
				Category: "docs",
			})
		}
	}
	return evidence
}