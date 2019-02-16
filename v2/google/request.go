package google

// Request is the HTTP request body from Google Assistant Actions.
type Request struct {
	User              *User
	Device            *Device
	Surface           *Surface
	Conversation      *Conversation
	Inputs            []*Input
	IsInSandbox       bool
	AvailableSurfaces []*Surface
}
