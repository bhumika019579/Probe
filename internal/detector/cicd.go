package detector

import (
	"os"
	"path/filepath"
	"strings"
)

var cifiles = map[string]bool{
	".gitlab-ci.yml": true,
	"Jenkinsfile":    true,
}

type cicdDetector struct{}

func (c cicdDetector) Name() string {
	return "Change Management / CI-CD Detector"
}
func (c cicdDetector) Category() string {
	return "cicd"
}
func (c cicdDetector) Detect(files []string) []Evidence {
	var evidence []Evidence
	for _, path := range files {
		filename := filepath.Base(path)
		if strings.Contains(path,".github/workflow/")&&(strings.HasSuffix(filename,".yml")||strings.HasSuffix(filename,".yaml")){
			content,err:=os.ReadFile(path)
			hasTestStep:=false
			if err==nil&&strings.Contains(string(content),"test"){
				hasTestStep=true
			}
			desc:="Github Actions Workflow Found"
			confidence:="medium"
			if hasTestStep{
				desc="GitHub Actions workflow found, appears to run a test step"
				confidence="high"
			}
			evidence = append(evidence, Evidence{
				File: path,
				Description: desc,
				Confidence: confidence,
				Category: "cicd",
			})
			continue

		}
		if cifiles[filename]{
			evidence = append(evidence, Evidence{
				File: path,
				Description: filename+"Ci/Cd configuration found",
				Confidence: "high",
				Category: "cicd",
			})
			continue
		}
		if filename=="PULL_REQUEST_TEMPLATE.md"{
			evidence = append(evidence, Evidence{
				File: path,
				Description: "pull request template found — indicates a defined review process",
				Confidence: "high",
				Category: "cicd",
			})
			continue
		}
	}
	return evidence

}