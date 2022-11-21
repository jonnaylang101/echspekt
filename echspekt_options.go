package echspekt

func StatusCodeEqualTo(wantStatusCode int) EchspektOption {
	return func(t *Then) {
		t.wantStatusCode = wantStatusCode
	}
}
