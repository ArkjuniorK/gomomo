package strcon

func New() Service {
	var rp = newRepository()
	return newService(rp)
}
