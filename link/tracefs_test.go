package link

import (
	"testing"

	"github.com/go-quicktest/qt"
)

func TestSetTracefsPath(t *testing.T) {
	t.Cleanup(func() { _ = SetTracefsPath("") })

	// Clearing is always a no-op success.
	qt.Assert(t, qt.IsNil(SetTracefsPath("")))

	// Non-existent paths and non-tracefs directories must be rejected.
	qt.Assert(t, qt.IsNotNil(SetTracefsPath("/does/not/exist")))
	qt.Assert(t, qt.IsNotNil(SetTracefsPath(t.TempDir())))
}
