package level

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mattmeyers/assert"
)

type mockWriter struct {
	calls [][]byte
}

func (w *mockWriter) Write(b []byte) (int, error) {
	w.calls = append(w.calls, b)
	return len(b), nil
}

func TestNewBasicLoggerValidatesLevel(t *testing.T) {
	tests := []struct {
		name    string
		level   Level
		wantErr bool
	}{
		{
			name:    "Valid level",
			level:   LevelDebug,
			wantErr: false,
		},
		{
			name:    "Invalid level",
			level:   Level(-1),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBasicLogger(tt.level, nil)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestNewBasicLoggerDefaultsToStdout(t *testing.T) {
	l, err := NewBasicLogger(LevelDebug, nil)
	assert.NoError(assert.Fatal(t), err)
	assert.DeepEqual(t, l.w, os.Stdout)
}

func TestBasicLogger_Debug(t *testing.T) {
	type fields struct {
		level Level
	}
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantWrites int
	}{
		{
			name:       "Logs when LevelDebug - no newline",
			fields:     fields{level: LevelDebug},
			args:       args{format: "%s", args: []interface{}{"foo"}},
			wantWrites: 3,
		},
		{
			name:       "Logs when LevelDebug - newline",
			fields:     fields{level: LevelDebug},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 2,
		},
		{
			name:       "Does not log when higher level",
			fields:     fields{level: LevelInfo},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &mockWriter{calls: make([][]byte, 0)}
			l := &BasicLogger{
				w:     w,
				level: tt.fields.level,
			}
			l.Debug(tt.args.format, tt.args.args...)

			assert.Equal(t, len(w.calls), tt.wantWrites)
		})
	}
}

func TestBasicLogger_Info(t *testing.T) {
	type fields struct {
		level Level
	}
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantWrites int
	}{
		{
			name:       "Logs when LevelInfo - no newline",
			fields:     fields{level: LevelInfo},
			args:       args{format: "%s", args: []interface{}{"foo"}},
			wantWrites: 3,
		},
		{
			name:       "Logs when LevelInfo - newline",
			fields:     fields{level: LevelInfo},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 2,
		},
		{
			name:       "Logs when lower level",
			fields:     fields{level: LevelDebug},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 2,
		},
		{
			name:       "Does not log when higher level",
			fields:     fields{level: LevelFatal},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &mockWriter{calls: make([][]byte, 0)}
			l := &BasicLogger{
				w:     w,
				level: tt.fields.level,
			}
			l.Info(tt.args.format, tt.args.args...)

			assert.Equal(t, len(w.calls), tt.wantWrites)
		})
	}
}

func TestBasicLogger_Warn(t *testing.T) {
	type fields struct {
		level Level
	}
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantWrites int
	}{
		{
			name:       "Logs when LevelWarn - no newline",
			fields:     fields{level: LevelWarn},
			args:       args{format: "%s", args: []interface{}{"foo"}},
			wantWrites: 3,
		},
		{
			name:       "Logs when LevelWarn - newline",
			fields:     fields{level: LevelWarn},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 2,
		},
		{
			name:       "Logs when lower level",
			fields:     fields{level: LevelDebug},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 2,
		},
		{
			name:       "Does not log when higher level",
			fields:     fields{level: LevelFatal},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &mockWriter{calls: make([][]byte, 0)}
			l := &BasicLogger{
				w:     w,
				level: tt.fields.level,
			}
			l.Warn(tt.args.format, tt.args.args...)

			assert.Equal(t, len(w.calls), tt.wantWrites)
		})
	}
}

func TestBasicLogger_Error(t *testing.T) {
	type fields struct {
		level Level
	}
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantWrites int
	}{
		{
			name:       "Logs when LevelError - no newline",
			fields:     fields{level: LevelError},
			args:       args{format: "%s", args: []interface{}{"foo"}},
			wantWrites: 3,
		},
		{
			name:       "Logs when LevelError - newline",
			fields:     fields{level: LevelError},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 2,
		},
		{
			name:       "Logs when lower level",
			fields:     fields{level: LevelDebug},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 2,
		},
		{
			name:       "Does not log when higher level",
			fields:     fields{level: LevelFatal},
			args:       args{format: "%s\n", args: []interface{}{"foo"}},
			wantWrites: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &mockWriter{calls: make([][]byte, 0)}
			l := &BasicLogger{
				w:     w,
				level: tt.fields.level,
			}
			l.Error(tt.args.format, tt.args.args...)

			assert.Equal(t, len(w.calls), tt.wantWrites)
		})
	}
}

func TestBasicLogger_Fatal(t *testing.T) {
	if os.Getenv("TESTLEVELLOGGER_FATAL") == "1" {
		(&BasicLogger{
			w:     &mockWriter{calls: make([][]byte, 0)},
			level: LevelFatal,
		}).Fatal("")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestBasicLogger_Fatal")
	cmd.Env = append(os.Environ(), "TESTLEVELLOGGER_FATAL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
