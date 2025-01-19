package tests

import (
	"testing"

	"github.com/vsrtferrum/VkIntro/internal/engine"
	"github.com/vsrtferrum/VkIntro/internal/field"
)

func TestDeikstra(t *testing.T) {
	tests := []struct {
		flag         bool
		name         string
		Field        field.Field
		expectedPath *engine.List
	}{
		{
			name: "Normal path",
			Field: field.NewField(
				field.NewSize(3, 3),
				field.NewStartAndEnd(0, 0, 2, 2),
				&[][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			),
			expectedPath: &engine.List{X: 0, Y: 0, Next: &engine.List{X: 1, Y: 0, Next: &engine.List{X: 2, Y: 0, Next: &engine.List{X: 2, Y: 1, Next: &engine.List{X: 2, Y: 2, Next: nil}}}}},
		},
		{
			name: "Blocked path",
			Field: field.NewField(
				field.NewSize(3, 3),
				field.NewStartAndEnd(0, 0, 2, 2),
				&[][]int{
					{1, 0, 1},
					{1, 0, 1},
					{1, 0, 1},
				},
			),
			expectedPath: nil,
		},
		{
			name: "Start equals end",
			Field: field.NewField(
				field.NewSize(3, 3),
				field.NewStartAndEnd(1, 1, 1, 1),
				&[][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			),
			expectedPath: &engine.List{X: 1, Y: 1, Next: nil},
		},
		{
			name: "Empty field",
			Field: field.NewField(
				field.NewSize(3, 3),
				field.NewStartAndEnd(0, 0, 2, 2),
				&[][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			),
			expectedPath: nil,
		},
		{
			name: "Original test",
			Field: field.NewField(
				field.NewSize(3, 3),
				field.NewStartAndEnd(0, 0, 2, 1),
				&[][]int{
					{1, 2, 0},
					{2, 0, 1},
					{9, 1, 0},
				},
			),
			expectedPath: &engine.List{X: 0, Y: 0, Next: &engine.List{X: 1, Y: 0, Next: &engine.List{X: 2, Y: 0, Next: &engine.List{X: 2, Y: 1, Next: nil}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			path := engine.Deikstra(&tt.Field)
			lenpath := 0
			expectedlen := 0

			for temp := path; temp != nil; temp = temp.Next {
				lenpath++
			}

			for temp := tt.expectedPath; temp != nil; temp = temp.Next {
				expectedlen++
			}

			if lenpath != expectedlen {
				t.Errorf("expected path length %d, got %d", expectedlen, lenpath)
				return
			}

			for i := 0; i < lenpath; i++ {
				if path.X != tt.expectedPath.X || path.Y != tt.expectedPath.Y {
					t.Errorf("expected path point (%d, %d), got (%d, %d)", tt.expectedPath.X, tt.expectedPath.Y, path.X, path.Y)
					return
				}
				path = path.Next
				tt.expectedPath = tt.expectedPath.Next
			}
		})
	}
}
