package pointer

var (
	f = false
	t = true
)

func Of[T any](t T) *T {
	return &t
}

func False() *bool {
	return &f
}

func True() *bool {
	return &t
}
