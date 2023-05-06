package fx

type Predicate[T any] func(T) bool

type Func[T, U any] func(T) U

type Consumer[T any] func(T)

func Identity[T any](t T) T {
	return t
}

func IdentityInt(t int) int {
	return t
}
