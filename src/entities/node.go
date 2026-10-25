package entities

type Node struct {
	ID          int     `yaml:"-"`
	Type        string  `yaml:"name"`
	NwFunctions []VNF   `yaml:"-"`
	X, Y        float32 `yaml:"-"`
}
