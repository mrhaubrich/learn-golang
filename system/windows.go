package system

// WindowsFetcher is a struct that implements the BaseFetcher interface.
type WindowsFetcher struct{}

// FetchSystemData implements the BaseFetcher interface for Windows.
// It returns a SystemData instance with the logged-in user set.
func (w WindowsFetcher) FetchSystemData() SystemData {
	return SystemData{
		LoggedUser: "windows user", // Simulate fetching the logged user in Windows
	}
}
