package field

type StartAndEnd struct {
	xStart, yStart, xEnd, yEnd int
}

type StartAndEndImplementations interface {
	GetStartAndEnd() (int, int, int, int)
}

func NewStartAndEnd(xStart, yStart, xEnd, yEnd int) StartAndEnd {
	return StartAndEnd{xStart: xStart, yStart: yStart, xEnd: xEnd, yEnd: yEnd}
}

func (startAndEnd *StartAndEnd) GetStartAndEnd() (xStart, yStart, xEnd, yEnd int) {
	return startAndEnd.xStart, startAndEnd.yStart, startAndEnd.xEnd, startAndEnd.yEnd

}
