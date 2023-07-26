package project

// TerraformProject represents a terraform project
type TerraformProject struct {
	BaseDir      string
	RelativePath string
}

// NewTerraformProject TerraformProject constructor
func NewTerraformProject(baseDir string, relativePath string) TerraformProject {
	return TerraformProject{
		BaseDir:      baseDir,
		RelativePath: relativePath,
	}
}
