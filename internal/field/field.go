package field

type Field struct {
	Size
	field *[][]int
}

type FieldImplementation interface {
}

func NewField(size Size, field *[][]int) Field {
	return Field{Size: size, field: field}
}
