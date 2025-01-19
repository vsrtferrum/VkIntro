package engine

import (
	"math"

	"github.com/idsulik/go-collections/queue"
	"github.com/vsrtferrum/VkIntro/internal/field"
)

type List struct {
	X, Y       int
	Next, Last *List
}
type Pair struct {
	A, B int
}

type EngineImplemetation interface {
	getPos(int, int, int) int
	logic(*[]int, *field.Field, [][]*List, Pair, int, int, int, int, bool) bool
	Deikstra(*field.Field) []*List
}

func getPos(x, y, h int) int {
	return x*h + y
}
func logic(mass *[]int, field *field.Field, path *[]*List, temp Pair, x, y, h int) bool {
	xs, ys, _, _ := field.GetStartAndEnd()
	if x < 0 || y < 0 || x >= field.GetLenght() || y >= field.GetHeight() || (x == xs && y == ys) {
		return false
	}
	if field.GetField(x, y) == 0 {
		return false
	}
	if uint64((*mass)[getPos(temp.A, temp.B, h)])+uint64(field.GetField(x, y)) < uint64((*mass)[getPos(x, y, h)]) {
		(*mass)[getPos(x, y, h)] = (*mass)[getPos(temp.A, temp.B, h)] + field.GetField(x, y)
		for temp2 := (*path)[getPos(temp.A, temp.B, h)]; temp2 != nil; temp2 = temp2.Next {
			if (*path)[getPos(x, y, h)] == nil {
				(*path)[getPos(x, y, h)] = &List{X: temp2.X, Y: temp2.Y, Next: nil, Last: nil}
				(*path)[getPos(x, y, h)].Last = (*path)[getPos(x, y, h)]
			} else {
				(*path)[getPos(x, y, h)].Last.Next = &List{X: temp2.X, Y: temp2.Y, Next: nil, Last: nil}
				(*path)[getPos(x, y, h)].Last = (*path)[getPos(x, y, h)].Last.Next
			}
		}
		(*path)[getPos(x, y, h)].Last.Next = &List{X: x, Y: y, Next: nil}
		(*path)[getPos(x, y, h)].Last = (*path)[getPos(x, y, h)].Last.Next
		return true
	}
	return false
}

func Deikstra(field *field.Field) *List {
	xStart, yStart, xEnd, yEnd := field.GetStartAndEnd()
	h := field.GetHeight()
	mass := make([]int, field.GetHeight()*field.GetLenght())
	path := make([]*List, field.GetHeight()*field.GetLenght())

	for i := range mass {
		mass[i] = math.MaxInt
	}

	mass[getPos(xStart, yStart, h)] = 0
	path[getPos(xStart, yStart, h)] = &List{X: xStart, Y: yStart}
	path[getPos(xStart, yStart, h)].Last = path[getPos(xStart, yStart, h)]
	q := queue.New[Pair](0)
	q.Enqueue(Pair{xStart, yStart})
	for q.Len() != 0 {
		temp, _ := q.Dequeue()
		if temp.A < 0 || temp.B < 0 || temp.A >= field.GetLenght() || temp.B >= field.GetHeight() {
			continue
		}
		if field.GetField(temp.A, temp.B) == 0 {
			continue
		}
		xUp, xDown, yUp, yDown := temp.A+1, temp.A-1, temp.B+1, temp.B-1
		if xUp < field.GetLenght() {
			if logic(&mass, field, &path, temp, xUp, temp.B, h) {
				q.Enqueue(Pair{xUp, temp.B})
			}
		}
		if xDown > -1 {
			if logic(&mass, field, &path, temp, xDown, temp.B, h) {
				q.Enqueue(Pair{xDown, temp.B})
			}
		}
		if yUp < field.GetHeight() {
			if logic(&mass, field, &path, temp, temp.A, yUp, h) {
				q.Enqueue(Pair{temp.A, yUp})
			}
		}
		if yDown > -1 {
			if logic(&mass, field, &path, temp, temp.A, yUp, h) {
				q.Enqueue(Pair{temp.A, yDown})
			}
		}
	}
	return path[getPos(xEnd, yEnd, h)]
}
