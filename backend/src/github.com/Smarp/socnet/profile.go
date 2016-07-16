package socnet

type Profile struct {
	Socnet     string
	Id         string
	AvatarUrl  string
	ProfileUrl string
	FullName   string
	Email      string
	FirstName  string
	LastName   string
	AgeRange   struct {
		Min int
		Max int
	}
	Gender      string
	Locale      string
	Timezone    int
	UpdatedTime string
	Verified    bool
}

type Profiler interface {
	Profile() (*Profile, error)
}
