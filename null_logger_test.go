package level

import (
	"os"
	"os/exec"
	"testing"
)

func TestNullLogger_Fatal(t *testing.T) {
	if os.Getenv("TESTLEVELLOGGER_FATAL") == "1" {
		(NullLogger{}).Fatal("")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestNullLogger_Fatal")
	cmd.Env = append(os.Environ(), "TESTLEVELLOGGER_FATAL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
