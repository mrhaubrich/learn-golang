package system

// LinuxFetcher is a struct that implements the BaseFetcher interface.
type LinuxFetcher struct{}

// FetchSystemData implements the BaseFetcher interface for Windows.
// It returns a SystemData instance with the logged-in user set.
func (l LinuxFetcher) FetchSystemData() SystemData {
	return SystemData{
		LoggedUser: "linux user", // Simulate fetching the logged user in Windows
	}
}
