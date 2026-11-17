package entities

type Graph struct {
	Nodes       []*Node       `yaml:"nodes"`
	Connections []*Connection `yaml:"connections"`
}
