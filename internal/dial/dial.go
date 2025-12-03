package dial

type Dial struct {
	Position       int
	KeyValueSeen   int
	keyValue       int
	KeyValuePassed int
	NumPositions   int
}

func NewDial(startingPosition int, numPositions int, keyValue int) (d *Dial) {
	return &Dial{
		Position:     startingPosition,
		NumPositions: numPositions,
		keyValue:     keyValue,
	}
}

func (d *Dial) TimesPassedZero(distanceFromZero int, steps int) int {
	timesPassed := 0
	stepsRemaining := steps

	// Take off the surplus
	if stepsRemaining >= distanceFromZero {
		timesPassed += 1
		stepsRemaining -= distanceFromZero
	}

	// stepsRemaining is now from pos 0
	// Just divide by NumPositions
	if stepsRemaining >= d.NumPositions {
		timesPassed += stepsRemaining / d.NumPositions
	}

	return timesPassed
}

func (d *Dial) RotateRight(steps int) *Dial {
	distanceFromZero := d.NumPositions - d.Position
	d.KeyValuePassed += d.TimesPassedZero(distanceFromZero, steps)

	// For a Right rotation we must increment steps then modulo num positions
	d.Position = (d.Position + steps) % d.NumPositions
	if d.Position == d.keyValue {
		d.KeyValueSeen += 1
	}
	return d
}

func (d *Dial) RotateLeft(steps int) *Dial {
	distanceFromZero := d.Position
	// if starting position is 0 then we are actually NumPositions from passing zero in the rotation
	if distanceFromZero == 0 {
		distanceFromZero = d.NumPositions
	}
	d.KeyValuePassed += d.TimesPassedZero(distanceFromZero, steps)
	
	// For a left rotation decrement steps and then modulo num positions
	d.Position = (((d.Position - steps) % d.NumPositions) + d.NumPositions) % d.NumPositions
	if d.Position == d.keyValue {
		d.KeyValueSeen += 1
	}
	return d
}
