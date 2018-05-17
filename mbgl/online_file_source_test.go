package mbgl

import (
	"testing"
)

func TestNewOnlineFileSource(t *testing.T) {
	
	NewRunLoop()
	fs := NewOnlineFileSource()
	
	if &fs.cptr == nil {
		t.Fatal("NewOnlineFileSource returned nil")
	}	
}

