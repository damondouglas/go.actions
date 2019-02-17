package google

// PushMessage holds structured data to push for the Actions Fulfillment API.
type PushMessage struct {
	Target           *Target           `json:"target"`
	OrderUpdate      *OrderUpdate      `json:"orderUpdate"`
	UserNotification *UserNotification `json:"userNotification"`
}

// Target for the push request.
type Target struct {
	UserID   string    `json:"userId"`
	Intent   string    `json:"intent"`
	Argument *Argument `json:"argument"`
	Locale   string    `json:"locale"`
}

// OrderUpdate to an order.
type OrderUpdate struct {
	// Todo: to complete later.
}

// UserNotification specifies title and text displayed to user in message.
type UserNotification struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
