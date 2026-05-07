package link

import (
	"github.com/cilium/ebpf/internal/tracefs"
)

// SetTracefsPath overrides the tracefs mount point used when attaching
// Kprobe, Uprobe, and Tracepoint links. Pass an empty string to clear a
// previous override and resume auto-detection of /sys/kernel/tracing or
// /sys/kernel/debug/tracing.
//
// The path must be an existing tracefs or debugfs mount; SetTracefsPath
// validates the filesystem type and returns an error otherwise.
//
// This is intended for containerized environments where the canonical
// tracefs path inside the container is something other than the kernel
// default — for example, a /host bind mount.
//
// SetTracefsPath should be called once before any tracefs-backed link is
// attached. Concurrent calls are safe; established links are unaffected.
func SetTracefsPath(path string) error {
	return tracefs.SetPath(path)
}
