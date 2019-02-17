package google

// User represents user data in Google payload of request.
type User struct {
	UserID              string
	IDToken             string
	Profile             *UserProfile
	AccessToken         string
	Permissions         []string // Todo: make enum type.
	Locale              string
	LastSeen            string
	UserStorage         string
	PackageEntitlements []*PackageEntitlement
}

// UserProfile represents user name information.
type UserProfile struct {
	DisplayName string
	GivenName   string
	FamilyName  string
}

// PackageEntitlement represents list of entitlements related to a package name.
type PackageEntitlement struct {
	PackageName  string
	Entitlements []struct {
		SKU          string
		SKUType      string // Todo: make enum type
		InAppDetails struct {
		}
	}
}
