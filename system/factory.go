package system

import (
	"fmt"
	"runtime"
)

// GetFetcher is a factory method that returns the appropriate BaseFetcher
// depending on the current operating system.
func GetFetcher() (BaseFetcher, error) {
	switch os := runtime.GOOS; os {
	case "windows":
		return WindowsFetcher{}, nil
	case "linux":
		return LinuxFetcher{}, nil
	case "darwin":
		// Assuming macOS fetcher would be implemented later
		return nil, fmt.Errorf("fetcher for %s is not implemented yet", os)
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", os)
	}
}
