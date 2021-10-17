package clistate

import (
	"testing"

	"github.com/Eitol/terraform-internals/pkg/command/arguments"
	"github.com/Eitol/terraform-internals/pkg/command/views"
	"github.com/Eitol/terraform-internals/pkg/states/statemgr"
	"github.com/Eitol/terraform-internals/pkg/terminal"
)

func TestUnlock(t *testing.T) {
	streams, _ := terminal.StreamsForTesting(t)
	view := views.NewView(streams)

	l := NewLocker(0, views.NewStateLocker(arguments.ViewHuman, view))
	l.Lock(statemgr.NewUnlockErrorFull(nil, nil), "test-lock")

	diags := l.Unlock()
	if diags.HasErrors() {
		t.Log(diags.Err().Error())
	} else {
		t.Error("expected error")
	}
}
