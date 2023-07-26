package orchestration

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/atorrescogollo/terraform-cascade/internal/project"
)

// ProjectOrderResolver resolves the terraform projects in order
// of precedence
type ProjectOrderResolver struct {
	ReadDirFn func(name string) ([]fs.DirEntry, error)
}

// NewProjectOrderResolver ProjectOrderResolver constructor
func NewProjectOrderResolver() *ProjectOrderResolver {
	return &ProjectOrderResolver{
		ReadDirFn: os.ReadDir,
	}
}

// Resolve resolves the terraform projects in order of precedence
func (r ProjectOrderResolver) Resolve(cwd string) ([]project.TerraformProject, error) {
	return r.resolve(cwd, ".")
}
func (r ProjectOrderResolver) resolve(baseDir string, cwd string) ([]project.TerraformProject, error) {
	files, err := r.ReadDirFn(cwd)
	if err != nil {
		return nil, err
	}

	projectResolutions := make([]project.TerraformProject, 0)
	baseResolutions := make([]project.TerraformProject, 0)
	otherResolutions := make([]project.TerraformProject, 0)

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			// Hidden file, ignore
			continue
		}
		if file.IsDir() {
			if file.Name() == "base" {
				// Base project, whole project takes precedence
				res, err := r.resolve(baseDir, filepath.Join(cwd, file.Name()))
				if err != nil {
					return nil, err
				}
				baseResolutions = append(
					baseResolutions,
					res...,
				)
				continue
			}
			// Other project, normal precedence
			res, err := r.resolve(baseDir, filepath.Join(cwd, file.Name()))
			if err != nil {
				return nil, err
			}
			otherResolutions = append(
				otherResolutions,
				res...,
			)
			continue
		}
		if file.Name() == "backend.tf" {
			// Backend file, project over folders
			projectResolutions = append(projectResolutions, project.NewTerraformProject(
				baseDir,
				cwd,
			))
			continue
		}
	}

	result := make([]project.TerraformProject, 0)
	result = append(result, projectResolutions...)
	result = append(result, baseResolutions...)
	result = append(result, otherResolutions...)
	return result, nil
}
