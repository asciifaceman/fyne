package app

import (
	"os"
	"os/exec"
	"testing"

	"fyne.io/fyne"
	_ "fyne.io/fyne/test"
	"github.com/stretchr/testify/assert"
)

func TestDummyApp(t *testing.T) {
	app := New()

	app.Quit()
}

func TestCurrentApp(t *testing.T) {
	app := New()

	assert.Equal(t, app, fyne.CurrentApp())
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// some code here to check arguments perhaps?
	//fmt.Fprintf(os.Stdout, dockerRunResult)
	os.Exit(0)
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}
