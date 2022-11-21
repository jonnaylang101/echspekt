package echspekt

type Then struct {
	wantStatusCode int
	wantErr        bool
	bindObject     interface{}
	wantObject     interface{}
}

type EchspektOption func(*Then)

func AResponseBodyEqualTo(match interface{}) EchspektOption {
	return func(e *Then) {
		e.wantObject = match
	}
}
