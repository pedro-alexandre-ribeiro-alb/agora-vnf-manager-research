package types

type Chart struct {
	Name    string   `yaml:"name"`
	Version string   `yaml:"version"`
	Urls    []string `yaml:"urls"`
}
