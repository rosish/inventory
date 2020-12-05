package urls

type (
	urls struct {
		FETCH_INV_PATH    string
	}
)

func ReturnURLS() urls {
	var urlPatterns urls
	urlPatterns.FETCH_INV_PATH = "/inv/"
	return urlPatterns
}

