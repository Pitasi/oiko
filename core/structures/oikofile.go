package structures

type Oikofile struct {
	ProjectName string `yaml:"project_name"`
	Namespace   string
	Version     string
	Owner       string
	Email       string
	License     string
	Vcs struct {
		name string
		url  string
	}
	Dependencies []string
}
