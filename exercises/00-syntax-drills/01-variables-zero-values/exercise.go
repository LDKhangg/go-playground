package zerovalues

type Profile struct {
	Name     string
	Active   bool
	Attempts int
}

func DefaultProfile() Profile {
	return Profile{}
}
