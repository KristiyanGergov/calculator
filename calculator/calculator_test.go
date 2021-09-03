package calculator

import (
	"calculator_assignment/parser"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpressionCalculator_Expression(t *testing.T) {
	type args struct {
		expression []parser.Token
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{
			name:    "Test multiplication",
			args:    args{
				expression: []parser.Token{
					{
						Type:  parser.Operand,
						Value: 13.0,
					},
					{
						Type:  parser.Operand,
						Value: 2.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Multiply,
					},
				},
			},
			want:    26.0,
			wantErr: nil,
		},
		{
			name:    "Test division",
			args:    args{
				expression: []parser.Token{
					{
						Type:  parser.Operand,
						Value: 17.0,
					},
					{
						Type:  parser.Operand,
						Value: 2.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Divide,
					},
				},
			},
			want:    8.5,
			wantErr: nil,
		},
		{
			name:    "Test subtraction",
			args:    args{
				expression: []parser.Token{
					{
						Type:  parser.Operand,
						Value: 150.0,
					},
					{
						Type:  parser.Operand,
						Value: 37.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Subtract,
					},
				},
			},
			want:    113.0,
			wantErr: nil,
		},
		{
			name:    "Test addition",
			args:    args{
				expression: []parser.Token{
					{
						Type:  parser.Operand,
						Value: 4.5,
					},
					{
						Type:  parser.Operand,
						Value: 20.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Add,
					},
				},
			},
			want:    24.5,
			wantErr: nil,
		},
		{
			name:    "18|2|/|3|*|2|-| should equal 25",
			args:    args{
				expression: []parser.Token{
					{
						Type:  parser.Operand,
						Value: 18.0,
					},
					{
						Type:  parser.Operand,
						Value: 2.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Divide,
					},
					{
						Type:  parser.Operand,
						Value: 3.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Multiply,
					},
					{
						Type:  parser.Operand,
						Value: 2.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Subtract,
					},
				},
			},
			want:    25,
			wantErr: nil,
		},
		{
			name:    "Division by 0 should result in an error",
			args:    args{
				expression: []parser.Token{
					{
						Type:  parser.Operand,
						Value: 17.0,
					},
					{
						Type:  parser.Operand,
						Value: 0.0,
					},
					{
						Type:  parser.Operator,
						Value: parser.Divide,
					},
				},
			},
			want:    0,
			wantErr: fmt.Errorf("division with 0 is impossible"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := &ExpressionCalculator{}
			got, err := ex.Expression(tt.args.expression)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
