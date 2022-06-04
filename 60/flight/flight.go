package flight

// Position 飞机的位置信息
type Position struct {
	latitude  float32
	longitude float32
	height    float32
}

func NewPosition(lat, lon, hei float32) Position {
	return Position{
		latitude:  lat,
		longitude: lon,
		height:    hei,
	}

}