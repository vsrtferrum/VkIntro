package field

type Field struct {
	Size
	StartAndEnd
	field *[][]int
}

type FieldImplementation interface {
	GetField(a, b int) int
}

func (field Field) GetField(a, b int) int {
	return (*field.field)[a][b]
}

func NewField(size Size, startAndEnd StartAndEnd, field *[][]int) Field {
	return Field{Size: size, StartAndEnd: startAndEnd, field: field}
}
