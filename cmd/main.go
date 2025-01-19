package main

import (
	"github.com/vsrtferrum/VkIntro/internal/engine"
	"github.com/vsrtferrum/VkIntro/internal/input"
	"github.com/vsrtferrum/VkIntro/internal/output"
)

func main() {

	field, err := input.GetField()
	if err != nil {
		panic(err)
	}
	path := engine.Deikstra(&field)
	output.Output(path)

}
