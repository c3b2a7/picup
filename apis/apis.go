package apis

type API interface {
	Up(path string) (string, error)
}
