package input

import (
	"fmt"
	"strconv"

	"github.com/vsrtferrum/VkIntro/internal/errors"
	"github.com/vsrtferrum/VkIntro/internal/field"
)

type Filter interface {
	GetSize() (field.Size, error)
	GetFiled(field.Size) (field.Size, error)
}

func GetSize() (field.Size, error) {
	var lenght, height string
	fmt.Scan(&lenght, &height)
	len, err := strconv.Atoi(lenght)
	if err != nil || len < 1 {
		return field.Size{}, errors.ErrParseSize
	}

	high, err := strconv.Atoi(height)
	if err != nil || high < 1 {
		return field.Size{}, errors.ErrParseSize
	}
	return field.NewSize(len, high), nil
}

func GetFiled() (field.Field, error) {
	var temp string
	size, err := GetSize()
	if err != nil {
		return field.Field{}, err
	}
	gameField := make([][]int, size.GetHeight())
	for i := 0; i < size.GetHeight(); i++ {
		gameField[i] = make([]int, size.GetLenght())
		for j := 0; j < size.GetLenght(); j++ {
			fmt.Scan(&temp)
			num, err := strconv.Atoi(temp)
			if err != nil || num < 0 || num > 9 {
				return field.Field{}, errors.ErrParseField
			}
			gameField[i][j] = num
		}
	}
	return field.NewField(size, &gameField), nil
}
