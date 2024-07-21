package pointer

// Of returns a pointer to the input. In most cases, callers should just do &t. However, in some cases
// Go cannot take a pointer. For example, `ptr.Of(f())`.
func Of[T any](t T) *T {
	return &t
}
