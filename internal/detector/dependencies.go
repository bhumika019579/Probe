package detector

import "path/filepath"

var lockFiles = map[string]string{
	"go.sum":            "Go module checksums",
	"package-lock.json": "npm lockfile",
	 "package.json":    "Node.js",
	"yarn.lock":         "Yarn lockfile",
	"requirements.txt":  "Python pip requirements",
	"poetry.lock":       "Python Poetry lockfile",
	"pom.xml":           "Java Maven config",
}

type dependenciesDetector struct{}

func (d dependenciesDetector) Name() string {
	return "dependency management detector"
}
func (d dependenciesDetector) Category() string {
	return "dependencies"
}
func (d dependenciesDetector) Detect(files []string) []Evidence {
	var evidence []Evidence
	for _, path := range files {
		filename := filepath.Base(path)
		if desc,found:=lockFiles[filename];found{
			evidence = append(evidence, Evidence{
				File: path,
				Description: desc+"found dependencies are pinned",
				Confidence: "high",
				Category: "dependencies",
			})
			continue
		}
	}
	return evidence
}