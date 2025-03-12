package app

func NewServer() {

}

func PacakgeInitializer(i *Infra) {
	Initializer := NewInitializer(i)

	Initializer.InitFeedService()
}
