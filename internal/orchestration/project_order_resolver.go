package orchestration

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/atorrescogollo/terraform-cascade/internal/project"
)

type ProjectOrderResolver struct {
}

func NewProjectOrderResolver() *ProjectOrderResolver {
	return &ProjectOrderResolver{}
}

func (r ProjectOrderResolver) Resolve(cwd string) ([]project.TerraformProject, error) {
	resolution, err := r.resolveDir(cwd, ".")
	if err != nil {
		return nil, err
	}
	/*
	* Flatten the map so that it follows the following order:
	* - current directory
	* - base projects
	* - other projects
	 */
	result := make([]project.TerraformProject, 0)
	for i := 1; i <= len(resolution); i++ {
		bases := make([]project.TerraformProject, 0)
		other := make([]project.TerraformProject, 0)
		for _, p := range resolution[i] {
			if strings.HasSuffix("/"+p.RelativePath, "/base") {
				bases = append(bases, p)
			} else {
				other = append(other, p)
			}
		}
		result = append(result, bases...)
		result = append(result, other...)
	}
	return result, nil
}

func (r ProjectOrderResolver) resolveDir(workDir string, dir string) (map[int][]project.TerraformProject, error) {
	files, err := os.ReadDir(workDir + "/" + dir)
	if err != nil {
		return nil, err
	}

	// The key is the depth of the directory
	result := make(map[int][]project.TerraformProject, 0)

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if file.IsDir() {
			partial, err := r.resolveDir(
				workDir,
				filepath.Join(dir, file.Name()),
			)
			if err != nil {
				return nil, err
			}
			for k, v := range partial {
				// Recursion happened, so the real depth is k+1
				result[k+1] = append(result[k+1], v...)
			}
		} else if file.Name() == "backend.tf" {
			result[0] = append(result[0], project.NewTerraformProject(
				workDir,
				dir,
			))
		}
	}

	return result, nil
}
