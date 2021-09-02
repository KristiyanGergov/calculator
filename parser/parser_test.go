package parser

import (
	"github.com/stretchr/testify/require"
	"reflect"
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := &ExpressionParser{}
			got, err := ex.Expression(tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expression() got = %v, wantOperand %v", got, tt.want)
			}
		})
	}
}

func Test_handleOperatorToken(t *testing.T) {
	type args struct {
		token     string
		tokens    []Token
		operators *stack
	}
	tests := []struct {
		name    string
		args    args
		want    []Token
		want1   *stack
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := handleOperatorToken(tt.args.token, tt.args.tokens, tt.args.operators)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleOperatorToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleOperatorToken() got = %v, wantOperand %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("handleOperatorToken() got1 = %v, wantOperand %v", got1, tt.want1)
			}
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
