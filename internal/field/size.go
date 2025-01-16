package field

type Size struct {
	lenght, height int
	SizeImplemetation
}

type SizeImplemetation interface {
	GetLenght() int
	GetHeight() int
}

func (size Size) GetLenght() int {
	return size.lenght
}
func (size Size) GetHeight() int {
	return size.height
}

func NewSize(lenght, height int) Size {
	return Size{lenght: lenght, height: height}
}
