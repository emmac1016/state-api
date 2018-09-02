package handlers

import "testing"

func TestExecute(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		expectErr bool
	}{
		{
			name:      "execute returns valid parsed response",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
