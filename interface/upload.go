package upload

type Upload interface {
	Upload(path string) (string, error)
}
