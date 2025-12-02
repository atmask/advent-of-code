package dial

type Dial struct {
	Position     int
	KeyValueSeen int
	keyValue     int
	NumPositions int
}

func NewDial(startingPosition int, numPositions int, keyValue int) (d *Dial) {
	return &Dial{
		Position:     startingPosition,
		NumPositions: numPositions,
		keyValue:     keyValue,
	}
}

func (d *Dial) RotateRight(steps int) *Dial {
	// For a Right rotation we must increment steps then modulo num positions
	d.Position = (d.Position + steps) % d.NumPositions
	if d.Position == d.keyValue {
		d.KeyValueSeen += 1
	}
	return d
}

func (d *Dial) RotateLeft(steps int) *Dial {
	// For a left rotation decrement steps and then modulo num positions
	d.Position = (d.Position - steps + d.NumPositions) % d.NumPositions
	if d.Position == d.keyValue {
		d.KeyValueSeen += 1
	}
	return d
}
