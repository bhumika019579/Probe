package detector


func AllDetectors() []Detector {
	return []Detector{
		authDetector{},
		accessControlDetector{},
		SecretsDetector{},
		cicdDetector{},
		testingDetector{},
		dependenciesDetector{},
		docsDetector{},
		hardeningDetector{},
	}
}