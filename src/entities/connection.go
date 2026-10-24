package entities

type Connection struct {
	ID                 int
	From               *Node
	To                 *Node
	Slices             []Slice
	status             Status
	maxBandwidth       int
	avaliableBandwidth int
	latency            int
	jitter             int
	packetLossRate     float64
	cost               float64
	currentLoad        float32
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
