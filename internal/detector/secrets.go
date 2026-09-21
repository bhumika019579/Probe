package detector

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var Keypatterns = []*regexp.Regexp{
	regexp.MustCompile(`AKIA[0-9A-Z]{16}`),                    
	regexp.MustCompile(`sk_live_[0-9a-zA-Z]{20,}`),             
	regexp.MustCompile(`(?i)(api[_-]?key|secret|token)\s*=\s*["'][0-9a-zA-Z]{16,}["']`), 
}
type SecretsDetector struct{}
func isTrackedByGit(path string) bool {
	cmd := exec.Command("git", "ls-files", "--error-unmatch", filepath.Base(path))
	cmd.Dir = filepath.Dir(path)
	return cmd.Run() == nil
}
func (s SecretsDetector) Name() string {
	return "Secrets Management Detector"
}

func (s SecretsDetector) Category() string {
	return "secrets"
}
func(s SecretsDetector)Detect(files[]string)[]Evidence{
	var evidence[]Evidence
	for _,path:=range files{
		filename:=filepath.Base(path)
		if filename==".env"{
			if isTrackedByGit(path){
			evidence = append(evidence, Evidence{
				File:        path,
				Description: ".env file committed to repository — risk of exposed secrets",
				Confidence:  "high",
				Category:    "secrets",
			})
		}
			continue

		}
		if filename==".env.example"||filename==".env.sample"{
			evidence = append(evidence, Evidence{
				File: path,
				Description: "environment variable template present — good secrets hygiene",
				Confidence: "high",
				Category: "secrets",
			})
            continue
		}
		ext:=filepath.Ext(path)
		if ext==""||strings.Contains(filepath.ToSlash(path),".git/"){
			continue
		}
		content,err:=os.ReadFile(path)
		if err!=nil{
			continue
		}
		lines:=strings.Split(string(content),"\n")
		for lineNum,line:=range lines{
			for _,patterns:=range Keypatterns{
				if patterns.MatchString(line){
					evidence = append(evidence, Evidence{
						File: path,
						Line: lineNum+1,
						Description:  "possible hardcoded secret/key-shaped string found",
						Confidence: "low",
						Category: "secrets",
					})
				}
			}
		}
	}
	return evidence
}