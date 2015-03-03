package goutils

import (
	"runtime"
)

func IsOSLinux() bool {
	return runtime.GOOS == "linux"
}
func IsOSWindows() bool {
	return runtime.GOOS == "windows"
}
func IsOSFreeBSD() bool {
	return runtime.GOOS == "freebsd"
}
func IsOSMac() bool {
	return runtime.GOOS == "darwin"
}
