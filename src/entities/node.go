package entities

type Node struct {
	ID          int          `yaml:"-"`
	Name        string       `yaml:"name"`
	NwFunctions []VNF        `yaml:"-"`
	X, Y        float32      `yaml:"-"`
	Neighbors   []Node       `yaml:"-"`
	Connections []Connection `yaml:"-"`
}
