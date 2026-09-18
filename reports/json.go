package reports

import (
	"encoding/json"
	"fmt"
	"os"
)

func ToJson(r Report) ([]byte, error) {
	data, err := json.MarshalIndent(r,"","")
	if err!=nil{
		return nil, fmt.Errorf("marshaling report to json: %w", err)
	}
	return data,nil
}
func WriteJson(r Report,path string)error{
	data,err:=ToJson(r)
	if err!=nil{
		return err
	}
	if err:=os.WriteFile(path,data,0644);err!=nil{
		return fmt.Errorf("writing json to %q: %w", path, err)
	}
	return nil
}