package detector

import (
	"os"
	"path/filepath"
	"strings"
)

var accessControlKeywords = map[string]string{
	"isAdmin":       "admin-check logic found",
	"hasRole":       "role-check logic found",
	"hasPermission": "permission-check logic found",
	"RequireAuth":   "auth-guard middleware pattern found",
	"middleware":    "middleware usage found (possible route guarding)",
	"RBAC":          "explicit RBAC reference found",
}
var sourceExtensions = map[string]bool{
	".go":   true,
	".js":   true,
	".ts":   true,
	".py":   true,
	".java": true,
}

type accessControlDetector struct{}

func (a accessControlDetector) Name() string {
	return "access control detector"
}
func (a accessControlDetector) Category() string {
	return "access control"
}
func (a accessControlDetector) Detect(files []string) []Evidence {
	var evidence []Evidence
	for _, path := range files {
		filename := filepath.Base(path)
		if filename=="CODEOWNERS"{
			evidence = append(evidence, Evidence{
				File: path,
				Description: "CODEOWNERS file present — ownership/access boundaries defined",
				Confidence: "high",
				Category: "access_control",

			})
			continue
		}
		ext:=filepath.Ext(path)
		if !sourceExtensions[ext]{
			continue
		}
		content,err:=os.ReadFile(path)
		if err!=nil{
			continue
		}
		lines:=strings.Split(string(content),"\n")
		for lineNum, line:=range lines{
			for keyword,description:=range accessControlKeywords{
				if strings.Contains(line,keyword){
					evidence = append(evidence, Evidence{
						File:        path,
						Line:        lineNum + 1,
						Description: description + " (\"" + keyword + "\")",
						Confidence:  "low",
						Category:    "access-control",
					})

				}
			}

		}
	}
	return evidence
}