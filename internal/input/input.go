package input

import (
	"fmt"
	"math"
	"strconv"

	"github.com/vsrtferrum/VkIntro/internal/errors"
	"github.com/vsrtferrum/VkIntro/internal/field"
)

type Filter interface {
	GetSize() (field.Size, error)
	GetField(field.Size) (field.Size, error)
	GetStartAndEnd(field.StartAndEnd) field.StartAndEnd
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
	if uint64(high)*uint64(len) > uint64(math.MaxInt) || float64(high)*float64(len) > float64(math.MaxInt) {
		return field.Size{}, errors.ErrSizeOfSize
	}
	return field.NewSize(len, high), nil
}

func GetField() (field.Field, error) {
	var temp string
	size, err := GetSize()
	if err != nil {
		return field.Field{}, err
	}

	gameField := make([][]int, size.GetHeight())
	for i := 0; i < size.GetLenght(); i++ {
		gameField[i] = make([]int, size.GetHeight())
		for j := 0; j < size.GetHeight(); j++ {
			fmt.Scan(&temp)
			num, err := strconv.Atoi(temp)
			if err != nil || num < 0 || num > 9 {
				return field.Field{}, errors.ErrParseField
			}
			gameField[i][j] = num
		}
	}

	startAndEnd, err := GetStartAndEnd(size)
	if err != nil {
		return field.Field{}, err
	}
	return field.NewField(size, startAndEnd, &gameField), nil
}

func GetStartAndEnd(size field.Size) (field.StartAndEnd, error) {
	var xStart, yStart, xEnd, yEnd int
	var xS, yS, xE, yE string
	fmt.Scan(&xS, &yS, &xE, &yE)
	xStart, err := strconv.Atoi(xS)
	if err != nil {
		return field.StartAndEnd{}, errors.ErrParseStart
	}
	if xStart < 0 || xStart > size.GetHeight() {
		return field.StartAndEnd{}, errors.ErrStartOutOfRange
	}

	yStart, err = strconv.Atoi(yS)
	if err != nil {
		return field.StartAndEnd{}, errors.ErrParseStart
	}
	if yStart < 0 || yStart > size.GetLenght() {
		return field.StartAndEnd{}, errors.ErrStartOutOfRange
	}

	xEnd, err = strconv.Atoi(xE)
	if err != nil {
		return field.StartAndEnd{}, errors.ErrParseEnd
	}
	if xEnd < 0 || xEnd > size.GetHeight() {
		return field.StartAndEnd{}, errors.ErrEndOutOfRange
	}

	yEnd, err = strconv.Atoi(yE)
	if err != nil {
		return field.StartAndEnd{}, errors.ErrParseEnd
	}
	if yEnd < 0 || yEnd > size.GetLenght() {
		return field.StartAndEnd{}, errors.ErrEndOutOfRange
	}
	return field.NewStartAndEnd(xStart, yStart, xEnd, yEnd), nil
}
