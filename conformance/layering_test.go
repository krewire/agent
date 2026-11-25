package conformance_test

import (
	"os/exec"
	"testing"
)

func TestAgent_PureGoBuild(t *testing.T) {
	cmd := exec.Command("go", "build", "./core/...", "./conformance/...", "./version/...")
	cmd.Dir = ".."
	cmd.Env = append(cmd.Environ(), "CGO_ENABLED=0")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("pure Go build failed: %v\nOutput: %s", err, string(out))
	}
}
