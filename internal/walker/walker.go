package walker

import (
	"io/fs"
	"path/filepath"
)

var skipDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	"dist":         true,
	"build":        true,
}

func Walk(rootPath string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(rootPath,func(path string,d fs.DirEntry,err error)error{
		if err!=nil{
			return err
		}
		if d.IsDir(){
			if skipDirs[d.Name()]{
				return filepath.SkipDir
			}
			return nil
		}
		files=append(files,path)
		return nil
	})
	return files,err
}