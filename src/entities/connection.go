package entities

type Connection struct {
	ID                 int     `yaml:"-"`
	From               *Node   `yaml:"from"`
	To                 *Node   `yaml:"to"`
	Slices             []Slice `yaml:"-"`
	status             Status  `yaml:"-"`
	maxBandwidth       int     `yaml:"-"`
	avaliableBandwidth int     `yaml:"-"`
	latency            int     `yaml:"-"`
	jitter             int     `yaml:"-"`
	packetLossRate     float64 `yaml:"-"`
	cost               float64 `yaml:"-"`
	currentLoad        float32 `yaml:"-"`
}

func NewConnection(ID int, From *Node, To *Node, maxBandwidth int, latency int, status Status) *Connection {
	return &Connection{
		ID:                 ID,
		From:               From,
		To:                 To,
		maxBandwidth:       maxBandwidth,
		latency:            latency,
		status:             status,
		avaliableBandwidth: maxBandwidth,
	}
}

func (c *Connection) GetavaliableBandwidth() int {
	avaliableBandwidth := c.avaliableBandwidth
	for _, slice := range c.Slices {
		avaliableBandwidth -= slice.Bandwidth
	}
	return avaliableBandwidth
}

func (c *Connection) AddSlice(slice Slice) bool {
	if slice.Bandwidth > c.GetavaliableBandwidth() {
		return false
	}
	c.Slices = append(c.Slices, slice)
	return true
}
