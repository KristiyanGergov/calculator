package reader

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type MockConsoleReader struct {
	input string
}

func (m *MockConsoleReader) ReadString(byte) (string, error) {
	return m.input, nil
}

func NewMockConsoleReader(input string) *MockConsoleReader {
	return &MockConsoleReader{input: input}
}

func TestExpressionReader_ReadExpression(t *testing.T) {
	type fields struct {
		consoleReader ConsoleReader
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "No spaces between tokens",
			fields: fields{
				consoleReader: NewMockConsoleReader("5/2*3"),
			},
			want: []string{"5", "/", "2", "*", "3"},
		},
		{
			name: "Spaces between tokens",
			fields: fields{
				consoleReader: NewMockConsoleReader("5   / 2 *     3"),
			},
			want: []string{"5", "/", "2", "*", "3"},
		},
		{
			name: "Tokens with more than one character",
			fields: fields{
				consoleReader: NewMockConsoleReader("1225   / 255 *     38888888"),
			},
			want: []string{"1225", "/", "255", "*", "38888888"},
		},
		{
			name: "Reading Float Tokens",
			fields: fields{
				consoleReader: NewMockConsoleReader("1.2312 * 4123.2 / 1.222"),
			},
			want: []string{"1.2312", "*", "4123.2", "/", "1.222"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ExpressionReader{
				consoleReader: tt.fields.consoleReader,
			}
			got := c.ReadExpression()
			require.Equal(t, tt.want, got)
		})
	}
}
