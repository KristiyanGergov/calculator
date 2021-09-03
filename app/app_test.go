package app

import (
	"calculator_assignment/calculator"
	"calculator_assignment/parser"
	"calculator_assignment/reader"
	"github.com/stretchr/testify/require"
	"testing"
)

type MockReader struct {
	expression []string
}

func NewMockReader(expression []string) *MockReader {
	return &MockReader{expression: expression}
}

func (m *MockReader) ReadExpression() []string {
	return m.expression
}

func TestApp_Start(t *testing.T) {
	type fields struct {
		reader     reader.Reader
		parser     parser.Parser
		calculator calculator.Calculator
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "End to end",
			fields: fields{
				reader:     NewMockReader([]string{"5", "*", "2", "/", "10"}),
				parser:     parser.NewExpressionParser(),
				calculator: calculator.NewExpressionCalculator(),
			},
			want: "Expression: [5 * 2 / 10] = 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				reader:     tt.fields.reader,
				parser:     tt.fields.parser,
				calculator: tt.fields.calculator,
			}
			got := app.Start()
			require.Equal(t, tt.want, got)
		})
	}
}
