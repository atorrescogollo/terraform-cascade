package orchestration

import (
	"io/fs"
	"sort"
	"strings"
	"testing"

	"github.com/atorrescogollo/terraform-cascade/internal/project"
	"github.com/atorrescogollo/terraform-cascade/internal/shared/utils"
	"github.com/stretchr/testify/assert"
)

type MockedFolder struct {
	DirEntry fs.DirEntry
	Folders  map[string]MockedFolder
	Files    map[string]fs.DirEntry
}

var mockedFileHierarchy map[string]MockedFolder = map[string]MockedFolder{
	".": {
		DirEntry: utils.NewDirEntry(".", true),
		Folders: map[string]MockedFolder{
			"base": {
				DirEntry: utils.NewDirEntry("base", true),
				Folders: map[string]MockedFolder{
					"network": {
						DirEntry: utils.NewDirEntry("network", true),
						Folders: map[string]MockedFolder{
							"peerings": {
								DirEntry: utils.NewDirEntry("peerings", true),
								Folders:  map[string]MockedFolder{},
								Files: map[string]fs.DirEntry{
									"main.tf":    utils.NewDirEntry("main.tf", false),
									"backend.tf": utils.NewDirEntry("backend.tf", false),
								},
							},
						},
						Files: map[string]fs.DirEntry{
							"backend.tf": utils.NewDirEntry("backend.tf", false),
							"main.tf":    utils.NewDirEntry("main.tf", false),
						},
					},
					"permissions": {
						DirEntry: utils.NewDirEntry("permissions", true),
						Folders: map[string]MockedFolder{
							"base": {
								DirEntry: utils.NewDirEntry("base", true),
								Folders: map[string]MockedFolder{
									"iam": {
										DirEntry: utils.NewDirEntry("iam", true),
										Folders:  map[string]MockedFolder{},
										Files: map[string]fs.DirEntry{
											"main.tf":    utils.NewDirEntry("main.tf", false),
											"backend.tf": utils.NewDirEntry("backend.tf", false),
										},
									},
								},
								Files: map[string]fs.DirEntry{
									"main.tf":    utils.NewDirEntry("main.tf", false),
									"backend.tf": utils.NewDirEntry("backend.tf", false),
								},
							},
						},
						Files: map[string]fs.DirEntry{},
					},
				},
				Files: map[string]fs.DirEntry{},
			},
			"dev": {
				DirEntry: utils.NewDirEntry("dev", true),
				Folders: map[string]MockedFolder{
					"base": {
						DirEntry: utils.NewDirEntry("base", true),
						Folders: map[string]MockedFolder{
							"network": {
								DirEntry: utils.NewDirEntry("network", true),
								Folders: map[string]MockedFolder{
									"vpc": {
										DirEntry: utils.NewDirEntry("vpc", true),
										Folders:  map[string]MockedFolder{},
										Files: map[string]fs.DirEntry{
											"main.tf":    utils.NewDirEntry("main.tf", false),
											"backend.tf": utils.NewDirEntry("backend.tf", false),
										},
									},
								},
								Files: map[string]fs.DirEntry{
									"backend.tf": utils.NewDirEntry("backend.tf", false),
									"main.tf":    utils.NewDirEntry("main.tf", false),
								},
							},
							"permissions": {
								DirEntry: utils.NewDirEntry("permissions", true),
								Folders: map[string]MockedFolder{
									"base": {
										DirEntry: utils.NewDirEntry("base", true),
										Folders: map[string]MockedFolder{
											"iam": {
												DirEntry: utils.NewDirEntry("iam", true),
												Folders:  map[string]MockedFolder{},
												Files: map[string]fs.DirEntry{
													"main.tf":    utils.NewDirEntry("main.tf", false),
													"backend.tf": utils.NewDirEntry("backend.tf", false),
												},
											},
										},
										Files: map[string]fs.DirEntry{
											"main.tf":    utils.NewDirEntry("main.tf", false),
											"backend.tf": utils.NewDirEntry("backend.tf", false),
										},
									},
								},
								Files: map[string]fs.DirEntry{
									"main.tf":    utils.NewDirEntry("main.tf", false),
									"backend.tf": utils.NewDirEntry("backend.tf", false),
								},
							},
						},
						Files: map[string]fs.DirEntry{},
					},
					"s3": {
						DirEntry: utils.NewDirEntry("s3", true),
						Folders:  map[string]MockedFolder{},
						Files: map[string]fs.DirEntry{
							"main.tf":    utils.NewDirEntry("main.tf", false),
							"backend.tf": utils.NewDirEntry("backend.tf", false),
						},
					},
					"eks": {
						DirEntry: utils.NewDirEntry("eks", true),
						Folders: map[string]MockedFolder{
							"addons": {
								DirEntry: utils.NewDirEntry("addons", true),
								Folders:  map[string]MockedFolder{},
								Files: map[string]fs.DirEntry{
									"main.tf":    utils.NewDirEntry("main.tf", false),
									"backend.tf": utils.NewDirEntry("backend.tf", false),
								},
							},
						},
						Files: map[string]fs.DirEntry{
							"main.tf":    utils.NewDirEntry("main.tf", false),
							"backend.tf": utils.NewDirEntry("backend.tf", false),
						},
					},
				},
				Files: map[string]fs.DirEntry{},
			},
		},
		Files: map[string]fs.DirEntry{},
	},
}

func TestResolver(t *testing.T) {
	r := &ProjectOrderResolver{
		ReadDirFn: func(name string) ([]fs.DirEntry, error) {
			layers := strings.Split(name, "/")
			for _, layer := range layers {
				if layer == "." {
					layers = layers[1:]
					continue
				}
				break
			}
			layerMockedFolder := mockedFileHierarchy["."]
			for _, layer := range layers {
				layerMockedFolder = layerMockedFolder.Folders[layer]
			}
			result := make([]fs.DirEntry, 0)
			for _, folder := range layerMockedFolder.Folders {
				result = append(result, folder.DirEntry)
			}
			for _, file := range layerMockedFolder.Files {
				result = append(result, file)
			}
			sort.Slice(result, func(i, j int) bool { // Sort by name
				return result[i].Name() < result[j].Name()
			})
			return result, nil
		},
	}
	projects, err := r.Resolve(".")
	assert.Nil(t, err)
	assert.Equal(t, []project.TerraformProject{
		{BaseDir: ".", RelativePath: "base/network"},
		{BaseDir: ".", RelativePath: "base/network/peerings"},
		{BaseDir: ".", RelativePath: "base/permissions/base"},
		{BaseDir: ".", RelativePath: "base/permissions/base/iam"},
		{BaseDir: ".", RelativePath: "dev/base/network"},
		{BaseDir: ".", RelativePath: "dev/base/network/vpc"},
		{BaseDir: ".", RelativePath: "dev/base/permissions"},
		{BaseDir: ".", RelativePath: "dev/base/permissions/base"},
		{BaseDir: ".", RelativePath: "dev/base/permissions/base/iam"},
		{BaseDir: ".", RelativePath: "dev/eks"},
		{BaseDir: ".", RelativePath: "dev/eks/addons"},
		{BaseDir: ".", RelativePath: "dev/s3"},
	}, projects)
}
