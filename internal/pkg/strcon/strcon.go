package strcon

func New() Service {
	var (
		rp  = newRepository()
		svc = NewService(rp)
	)

	return svc
}
