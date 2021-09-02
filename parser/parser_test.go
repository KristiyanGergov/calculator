package parser

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpressionParser_Expression(t *testing.T) {
	type args struct {
		expression []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Token
		wantErr error
	}{
		{
			name: "Parsing expression with two consecutive operators with the same priority",
			args: args{
				expression: []string{"5", "/", "2", "*", "3"},
			},
			want: []Token{
				{Type: Operand, Value: 5},
				{Type: Operand, Value: 2},
				{Type: Operator, Value: Divide},
				{Type: Operand, Value: 3},
				{Type: Operator, Value: Multiply},
			},
			wantErr: nil,
		},
		{
			name: "Parsing expression with two consecutive operators where the first is with higher priority",
			args: args{
				expression: []string{"5", "/", "2", "+", "3"},
			},
			want: []Token{
				{Type: Operand, Value: 5},
				{Type: Operand, Value: 2},
				{Type: Operator, Value: Divide},
				{Type: Operand, Value: 3},
				{Type: Operator, Value: Add},
			},
			wantErr: nil,
		},
		{
			name: "Parsing expression with two consecutive operators where the second is with higher priority",
			args: args{
				expression: []string{"5", "+", "2", "/", "3"},
			},
			want: []Token{
				{Type: Operand, Value: 5},
				{Type: Operand, Value: 2},
				{Type: Operand, Value: 3},
				{Type: Operator, Value: Divide},
				{Type: Operator, Value: Add},
			},
			wantErr: nil,
		},
		{
			name: "Invalid operator passed",
			args: args{
				expression: []string{"5", "&", "2", "/", "3"},
			},
			want:    nil,
			wantErr: fmt.Errorf("invalid input. the character & is not valid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := &ExpressionParser{}
			got, err := ex.Expression(tt.args.expression)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_parseTokenIntoOperand(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name                 string
		args                 args
		wantOperand          int
		wantTokenIsAnOperand bool
	}{
		{
			name: "Token is an integer",
			args: args{
				token: "5123",
			},
			wantOperand:          5123,
			wantTokenIsAnOperand: true,
		},
		{
			name: "Token is not an integer",
			args: args{
				token: "gasdfasdf",
			},
			wantOperand:          0,
			wantTokenIsAnOperand: false,
		},
		{
			name: "Token contains numbers and letters",
			args: args{
				token: "g15asd1234fasdf",
			},
			wantOperand:          0,
			wantTokenIsAnOperand: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseTokenIntoOperand(tt.args.token)
			require.Equal(t, tt.wantOperand, got)
			require.Equal(t, tt.wantTokenIsAnOperand, got1)
		})
	}
}
