package project

type TerraformProject struct {
	BaseDir      string
	RelativePath string
}

func NewTerraformProject(baseDir string, relativePath string) TerraformProject {
	return TerraformProject{
		BaseDir:      baseDir,
		RelativePath: relativePath,
	}
}
