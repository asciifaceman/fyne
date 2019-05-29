package app

import (
	"net/url"
	"testing"
)

func TestOpenURL(t *testing.T) {
	app := New()
	app.setExec(fakeExecCommand)

	testURL, _ := url.Parse("fyne.io")
	err := app.OpenURL(testURL)
	if err != nil {
		t.Errorf("expected no error but received %v", err)
	}
}
