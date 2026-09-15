package detector

import (
	"os"
	"path/filepath"
	"strings"
)

var hardeningLibraries = map[string]string{
	"express-rate-limit":     "rate limiting middleware (Node)",
	"golang.org/x/time/rate": "rate limiting package (Go)",
	"django-ratelimit":       "rate limiting package (Django)",
	"helmet":                 "security headers middleware (Node)",
	"secure":                 "security headers package (Go)",
	"joi":                    "input validation library (Node)",
	"validator":              "input validation library (Go)",
	"django-csp":               "security headers middleware (Django)",
	"pydantic":                 "input validation library (Python)",
	"spring-security":          "authentication and authorization framework (Java)",
	"bucket4j":                 "rate limiting library (Java)",
	"hibernate-validator":      "input validation library (Java)",
}

type hardeningDetector struct{}

func (h hardeningDetector) Name() string {
	return "application hardening detector"
}
func (h hardeningDetector) Category() string {
	return "hardening"
}
func (h hardeningDetector) Detect(files []string) []Evidence {
	var evidence []Evidence
	for _, path := range files {
		filename := filepath.Base(path)
		if !dependencyFiles[filename]{
			continue
		}
		content,err:=os.ReadFile(path)
		if err!=nil{
			continue
		}
		text:=string(content)
		for lib,description:=range hardeningLibraries{
			if strings.Contains(text,lib){
				evidence = append(evidence, Evidence{
					File: path,
					Description: description + " (\"" + lib + "\") found code level signal only, not proof of runtime protection",
					Confidence: "low",
					Category: "hardening",
					
				})
				continue
			}
		}
	}
	return evidence
}