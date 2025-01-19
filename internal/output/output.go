package output

import (
	"fmt"

	"github.com/vsrtferrum/VkIntro/internal/engine"
)

func Output(data *engine.List) {

	for temp := data; temp != nil; temp = temp.Next {
		fmt.Printf("%d %d\n", temp.X, temp.Y)
	}
	fmt.Print(".\n")

}
