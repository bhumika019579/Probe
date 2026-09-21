package detector

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var authLibraries = map[string]string{
	"golang-jwt":          "Go JWT library",
	"golang.org/x/oauth2": "Go OAuth2 library",
	"passport":            "Node.js Passport auth middleware",
	"next-auth":           "Next.js authentication library",
	"jsonwebtoken":        "Node.js JWT library",
	"django.contrib.auth": "Django built-in auth",
	"spring-security":     "Spring Security (Java)",
	"pyjwt":               "Python JWT library",
	"oauth":               "OAuth library (any provider)", 
	"goth":                "Go multi-provider login library", 
	"authlib":             "Python OAuth library", 
	"auth0":               "Auth0 login service", 
	"bcrypt":              "Password hashing library", 
}
var dependencyFiles = map[string]bool{
	"go.mod":           true,
	"package.json":     true,
	"requirements.txt": true,
	"pyproject.toml":   true,
	"pom.xml":          true,
}
var authSourceExts = map[string]bool{
	".go": true, ".js": true, ".jsx": true, ".ts": true,
	".tsx": true, ".py": true, ".java": true,
}
var authRegex = regexp.MustCompile(
	`(?i)client_secret|redirect_uri|/(oauth|authorize|callback)\b|jwt\.(parse|new|sign|verify)|bearer |bcrypt|generatefrompassword`,
)

type authDetector struct{}

func (a authDetector) Name() string {
	return "authentication detector"
}
func (a authDetector) Category() string {
	return "authentication"
}
func (a authDetector) Detect(files []string) []Evidence {
	var evidence []Evidence
	for _, path := range files {
		filename := filepath.Base(path)
		if !dependencyFiles[filename]{
			if !authSourceExts[strings.ToLower(filepath.Ext(path))]{
				continue
			}
			src,err:=os.ReadFile(path)
			if err!=nil{
				continue
			}
			for i,line:=range strings.Split(string(src),"\n"){
				t:=strings.TrimSpace(line)
				if strings.HasPrefix(t,"//")||strings.HasPrefix(t,"#"){
					continue
				}
				if authRegex.MatchString(line){
					evidence = append(evidence, Evidence{
						File: path,
						Line: i+1,
						Description: "auth code pattern detected",
						Confidence: "low",
						Category: "authentication",
					})
					break
				}
			}
			continue
		}
		content,err:=os.ReadFile(path)
		if err!=nil{
			continue
		}
		text:=strings.ToLower(string(content))
		for lib,description:=range authLibraries{
			if strings.Contains(text,lib){
				evidence=append(evidence, Evidence{
					File: path,
					Description: description + "(\""+lib+"\")found in dependencies",
					Confidence: "high",
					Category: "authentication",
				})
			}
		}
	}
	return evidence
}