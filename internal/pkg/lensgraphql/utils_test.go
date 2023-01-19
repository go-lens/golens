package lensgraphql

import "testing"

func TestIsValidHex(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "valid hex",
			s:    "0x0D",
			want: true,
		},
		{
			name: "invalid hex with wrong characters",
			s:    "0x0DFZ",
			want: false,
		},
		{
			name: "invalid hex with odd length",
			s:    "0x0DD",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidHex(tt.s); got != tt.want {
				t.Errorf("isValidHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidProfileId(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "valid profileId",
			s:    "0x0D",
			want: true,
		},
		{
			name: "invalid profileId with wrong characters",
			s:    "000D",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidProfileId(tt.s); got != tt.want {
				t.Errorf("isValidProfileId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirstArg(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want string
	}{
		{
			name: "valid first arg",
			s:    []string{"0x0D"},
			want: "0x0D",
		},
		{
			name: "empty arg",
			s:    []string{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstArg(tt.s); got != tt.want {
				t.Errorf("isValidHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
