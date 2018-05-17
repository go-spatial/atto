package mbgl

import (
	"testing"
)

func TestNewOnlineFileSource(t *testing.T) {
	
	loop := NewRunLoop()
	defer loop.Destroy()
	
	fs := NewOnlineFileSource()
	defer fs.Destroy()
	
	if &fs.cptr == nil {
		t.Fatal("NewOnlineFileSource returned nil")
	}	
}

