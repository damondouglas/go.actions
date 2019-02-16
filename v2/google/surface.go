package google

// Surface represents information specific to the Google Assistant client surface the user is interacting with.
// Surface is distinguished from Device by the fact that multiple Assistant surfaces may live on the same device.
type Surface struct {
	Capabilities []struct {
		Name string
	}
}
