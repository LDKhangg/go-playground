package zerovalues

type Profile struct {
	Name     string
	Active   bool
	Attempts int
}

func DefaultProfile() Profile {
	DefaultProfile().Name = ""
	DefaultProfile().Active = false
	DefaultProfile().Attempts = 0
}
