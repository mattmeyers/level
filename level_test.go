package level

import (
	"testing"

	"github.com/mattmeyers/assert"
)

func TestLevel_validate(t *testing.T) {
	tests := []struct {
		name    string
		l       Level
		wantErr bool
	}{
		{
			name:    "Valid - Debug",
			l:       LevelDebug,
			wantErr: false,
		},
		{
			name:    "Valid - Info",
			l:       LevelInfo,
			wantErr: false,
		},
		{
			name:    "Valid - Warn",
			l:       LevelWarn,
			wantErr: false,
		},
		{
			name:    "Valid - Error",
			l:       LevelError,
			wantErr: false,
		},
		{
			name:    "Valid - Fatal",
			l:       LevelFatal,
			wantErr: false,
		},
		{
			name:    "Invalid level",
			l:       Level(-1),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.l.validate()
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name    string
		l       string
		want    Level
		wantErr bool
	}{
		{
			name:    "Parse debug",
			l:       "DEBUG",
			want:    LevelDebug,
			wantErr: false,
		},
		{
			name:    "Parse info",
			l:       "info",
			want:    LevelInfo,
			wantErr: false,
		},
		{
			name:    "Parse warn",
			l:       "warn",
			want:    LevelWarn,
			wantErr: false,
		},
		{
			name:    "Parse error",
			l:       "error",
			want:    LevelError,
			wantErr: false,
		},
		{
			name:    "Parse fatal",
			l:       "fatal",
			want:    LevelFatal,
			wantErr: false,
		},
		{
			name:    "Invalid level",
			l:       "Foo",
			want:    Level(-1),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLevel(tt.l)
			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}
