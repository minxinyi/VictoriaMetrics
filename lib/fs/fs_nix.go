//go:build linux || darwin || freebsd

package fs

import (
	"fmt"
	"os"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"golang.org/x/sys/unix"
)

type statfs_t = unix.Statfs_t

func freeSpace(stat statfs_t) uint64 {
	return uint64(stat.Bavail) * uint64(stat.Bsize)
}

func mustRemoveDirAtomic(dir string) {
	n := atomicDirRemoveCounter.Add(1)
	tmpDir := fmt.Sprintf("%s.must-remove.%d", dir, n)
	if err := os.Rename(dir, tmpDir); err != nil {
		logger.Panicf("FATAL: cannot move %s to %s: %s", dir, tmpDir, err)
	}
	MustRemoveAll(tmpDir)
}

func statfs(path string, stat *statfs_t) (err error) {
	return unix.Statfs(path, stat)
}
