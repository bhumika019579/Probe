package detector


type Evidence struct{
	File string
    Line  int
	Description string
	Confidence string
	Category string

}
type Detector interface{
	Name() string
	Category() string
	Detect(files []string)Evidence
}