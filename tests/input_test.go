package tests

import (
	errorsGoPack "errors"
	"os"
	"testing"

	"github.com/vsrtferrum/VkIntro/internal/errors"
	"github.com/vsrtferrum/VkIntro/internal/input"

	field "github.com/vsrtferrum/VkIntro/internal/field"
)

func TestGetSize(t *testing.T) {
	tests := []struct {
		input        string
		expectedSize field.Size
		expectedErr  error
	}{
		{"5 10\n", field.NewSize(5, 10), nil},
		{"0 10\n", field.Size{}, errors.ErrParseSize},
		{"5 0\n", field.Size{}, errors.ErrParseSize},
		{"-5 10\n", field.Size{}, errors.ErrParseSize},
		{"5 -10\n", field.Size{}, errors.ErrParseSize},
		{"abc 10\n", field.Size{}, errors.ErrParseSize},
		{"5 xyz\n", field.Size{}, errors.ErrParseSize},
		{"1000000000000 10000000000\n", field.Size{}, errors.ErrSizeOfSize},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			tempFile, err := os.CreateTemp("", "testinput")
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tempFile.Name())
			if _, err := tempFile.WriteString(tt.input); err != nil {
				t.Fatalf("failed to write to temp file: %v", err)
			}
			tempFile.Close()
			originalStdin := os.Stdin
			defer func() { os.Stdin = originalStdin }()
			file, err := os.Open(tempFile.Name())
			if err != nil {
				t.Fatalf("failed to open temp file: %v", err)
			}
			defer file.Close()

			os.Stdin = file

			gotSize, err := input.GetSize()

			if !errorsGoPack.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}

			if gotSize != tt.expectedSize {
				t.Errorf("expected size %v, got %v", tt.expectedSize, gotSize)
			}
		})
	}
}
func TestGetField(t *testing.T) {
	tests := []struct {
		input         string
		expectedField field.Field
		expectedErr   error
	}{
		{
			"3 4\n0 1 2 3\n4 5 6 7\n8 9 0 1\n1 1 2 2\n",
			field.NewField(field.NewSize(3, 4), field.NewStartAndEnd(1, 1, 2, 2), &[][]int{
				{0, 1, 2, 3},
				{4, 5, 6, 7},
				{8, 9, 0, 1},
			}),
			nil,
		},
		{
			"3 4\n0 1 2 3\n4 5 -1 7\n8 9 0 1\n1 1 2 2\n",
			field.Field{},
			errors.ErrParseField,
		},
		{
			"3 4\n0 1 2 3\n4 5 6 7\n8 9 0 1\n-1 -1 -1 -1\n",
			field.Field{},
			errors.ErrStartOutOfRange,
		},
		{
			"3 4\n0 1 2 3\n4 5 6 7\n8 9 0 1\n3 3 4 4\n",
			field.Field{},
			errors.ErrEndOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "testinput")
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tempFile.Name())
			if _, err := tempFile.WriteString(tt.input); err != nil {
				t.Fatalf("failed to write to temp file: %v", err)
			}
			tempFile.Close()

			originalStdin := os.Stdin
			defer func() { os.Stdin = originalStdin }()

			file, err := os.Open(tempFile.Name())
			if err != nil {
				t.Fatalf("failed to open temp file: %v", err)
			}
			defer file.Close()

			os.Stdin = file

			gotField, err := input.GetField()

			if !errorsGoPack.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				return
			}

			if gotField.Size != tt.expectedField.Size || gotField.StartAndEnd != tt.expectedField.StartAndEnd {
				t.Errorf("expected field %v, got %v", tt.expectedField, gotField)
				return
			}
			for i := 0; i < gotField.GetLenght(); i++ {
				for j := 0; j < gotField.GetHeight(); j++ {
					if gotField.GetField(i, j) != tt.expectedField.GetField(i, j) {
						t.Errorf("expected field %v, got %v", tt.expectedField.GetField(i, j), gotField.GetField(i, j))
						return
					}

				}
			}
		})
	}
}

func TestGetStartAndEnd(t *testing.T) {
	tests := []struct {
		input       string
		size        field.Size
		expected    field.StartAndEnd
		expectedErr error
	}{
		{"1 1 2 2\n", field.NewSize(3, 4), field.NewStartAndEnd(1, 1, 2, 2), nil},
		{"-1 -1 -1 -1\n", field.NewSize(3, 4), field.StartAndEnd{}, errors.ErrStartOutOfRange},
		{"3 3 4 4\n", field.NewSize(3, 4), field.StartAndEnd{}, errors.ErrEndOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			tempFile, err := os.CreateTemp("", "testinput")
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tempFile.Name())
			if _, err := tempFile.WriteString(tt.input); err != nil {
				t.Fatalf("failed to write to temp file: %v", err)
			}
			tempFile.Close()
			originalStdin := os.Stdin
			defer func() { os.Stdin = originalStdin }()

			file, err := os.Open(tempFile.Name())
			if err != nil {
				t.Fatalf("failed to open temp file: %v", err)
			}
			defer file.Close()

			os.Stdin = file

			gotStartAndEnd, err := input.GetStartAndEnd(tt.size)

			if !errorsGoPack.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}

			if gotStartAndEnd != tt.expected {
				t.Errorf("expected StartAndEnd %v, got %v", tt.expected, gotStartAndEnd)
			}
		})
	}
}
