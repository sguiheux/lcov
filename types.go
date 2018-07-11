package lcov

// Parser represents the lcov parser
type Parser struct {
	path string
}

//Report represents the result of  LcovParser.parse
type Report struct {
	Files            []FileReport
	TotalLines       int
	CoveredLines     int
	TotalFunctions   int
	CoveredFunctions int
	TotalBranches    int
	CoveredBranches  int
}

// FileReport contains all informations about a file
type FileReport struct {
	Path             string
	TotalLines       int
	CoveredLines     int
	TotalFunctions   int
	CoveredFunctions int
	TotalBranches    int
	CoveredBranches  int
}
