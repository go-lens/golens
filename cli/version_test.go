package cli

import (
	"bytes"
	"regexp"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestVersionCmdExecute(t *testing.T) {
	// Create buffer to capture output of command
	buffOut := new(bytes.Buffer)
	// Regex to match version output
	versionRegexp := regexp.MustCompile(`golens v[0-9][0-9|\.]+[0-9]`)

	// Set output of command to buffer
	rootCmd := RootCmd()
	rootCmd.AddCommand(VersionCmd())
	rootCmd.SetArgs([]string{"version"})
	rootCmd.SetOut(buffOut)
	log.SetOutput(buffOut)

	if err := rootCmd.Execute(); err != nil {
		t.Errorf("versionCmd.Execute() failed: %v", err)
	}
	if !versionRegexp.Match(buffOut.Bytes()) {
		t.Errorf("versionCmd.Execute() has an invalid output: %s", buffOut.String())
	}
}
