package system

// SystemData represents system-level data, similar to the Dart class.
type SystemData struct {
	LoggedUser string
}

// BaseFetcher defines the interface for fetching system data,
// similar to your Dart interface.
type BaseFetcher interface {
	FetchSystemData() SystemData
}
