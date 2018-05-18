package mbgl

import (
	"testing"
)

func TestNewOnlineFileSource(t *testing.T) {
	
	// The Online File Source requires a Run Loop
	loop := NewRunLoop()
    defer loop.Destroy()
	
	fs := NewOnlineFileSource()
	defer fs.Destroy()
	
	if &fs.cptr == nil {
		t.Fatal("NewOnlineFileSource returned nil")
	}	
}

