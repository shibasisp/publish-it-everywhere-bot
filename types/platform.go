package types

// Platform defines the publishing platform
// e.g. LinkedIn, Facebook, Twitter etc.
type Platform string

const (
	// LinkedinPlatform ...
	LinkedinPlatform Platform = "linkedin"
	// TwitterPlatform ...
	TwitterPlatform Platform = "twitter"
)

func (p Platform) String() string {
	return string(p)
}
