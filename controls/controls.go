package controls

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Evidence struct {
	File        string `json:"file"`
	Line        int    `json:"line"`
	Description string `json:"description"`
	Confidence  string `json:"confidence"`
	Category    string `json:"category"`
}
type Control struct {
	ID          string `yaml:"id"`
	Domain      string `yaml:"domain"`
	Name        string `yaml:"name"`
	Detector    string `yaml:"detector"`
	Description string `yaml:"description"`
}
type controlsFile struct {
	Controls []Control `yaml:"controls"`
}

func LoadControls(path string)([]Control,error){
	data,err:=os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading controls file %q: %w", path, err)
	}
	var cf controlsFile
	if err:=yaml.Unmarshal(data,&cf);err!=nil{
		return nil,fmt.Errorf("parsing controls file %q:%w",path,err)
	}
    if len(cf.Controls)==0{
		return nil, fmt.Errorf("no controls found in %q", path)
	}
	seenID:=make(map[string]bool)
	seenDetector:=make(map[string]bool)
	for _,c:=range cf.Controls{
		if c.ID==""{
			return nil,fmt.Errorf("control missing id ")
		}
		if seenID[c.ID]{
			return nil, fmt.Errorf("duplicate control ID %s",c.ID)
		}
		seenID[c.ID]=true
		if c.Detector==""{
			return nil,fmt.Errorf("control %s missing detector key",c.ID)
		}
		if seenDetector[c.Detector]{
			return nil,fmt.Errorf("detector %q is mapped to more than one control",c.Detector)
		}
		seenDetector[c.Detector]=true

	}
	return cf.Controls,nil


}