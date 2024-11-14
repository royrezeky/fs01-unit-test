package ewallet

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	isTest = true
}

func TestRun(t *testing.T) {
	type args struct {
		perintah []string
	}
	tests := []struct {
		name     string
		args     args
		expected float64
		err      error
	}{
		// error case 1
		{
			name: "tarik tanpa setor",
			args: args{
				perintah: []string{"withdraw"},
			},
			expected: 0,
			err:      errors.New("saldo anda tidak mencukupi"),
		},
		// test case 1
		{
			name: "setor sekali",
			args: args{
				perintah: []string{"deposit"},
			},
			expected: 50000,
			err:      nil,
		},
		// test case 2
		{
			name: "setor 2x",
			args: args{
				perintah: []string{"deposit", "deposit"},
			},
			expected: 150000,
			err:      nil,
		},
		// test case 3
		{
			name: "setor 3x",
			args: args{
				perintah: []string{"deposit", "deposit", "deposit"},
			},
			expected: 300000,
			err:      nil,
		},
		// test case 4
		{
			name: "setor 5x",
			args: args{
				perintah: []string{"deposit", "deposit", "deposit", "deposit", "deposit"},
			},
			expected: 550000,
			err:      nil,
		},
		// test case 5 
		{
			name: "setor 8x",
			args: args{
				perintah: []string{"deposit", "deposit", "deposit", "deposit", "deposit", "deposit", "deposit", "deposit"},
			},
			expected: 950000,
			err:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Run(tt.args.perintah)
			assert.Equal(t, res, tt.expected)
			if err != nil {
				assert.Equal(t, err, tt.err)
			}
		})
	}
}
