package immortal

import (
	"time"
)

// GetJSON unix socket web client
func GetJSON(spath, path string, target interface{}) error {
	_ = "STUB: not implemented"
	// http socket client
	return nil
}

// AbsSince return time since in format [days]d[hours]h[minutes]m[seconds.decisecond]s
func AbsSince(t time.Time) string { _ = "STUB: not implemented"; return "" }

// md5sum return md5 checksum of given file
func md5sum(filePath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// inSlice find if item is in slice
func inSlice(s []string, item string) bool { _ = "STUB: not implemented"; return false }

// GetSdir return the main supervise directory, defaults to /var/run/immortal
func GetSdir() string {
	_ = "STUB: not implemented"
	// if IMMORTAL_SDIR env is set, use it as default sdir
	return ""
}

// GetUserSdir returns the $HOME/.immortal
func GetUserSdir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// isDir return true if path is a dir
func isDir(path string) bool { _ = "STUB: not implemented"; return false }

// isFile return true if path is a regular file
func isFile(path string) bool { _ = "STUB: not implemented"; return false }
