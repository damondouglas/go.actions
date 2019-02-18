package helper

// Storer stores object in storage.
type Storer interface {
	Store(obj interface{}) error
}
