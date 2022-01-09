package api

import "net/http"

// Optional interfaces for request type

//WithHeader sets headers to request
type WithHeader interface {
	WithHeader(header http.Header)
}

//WithMethod sets method to request
type WithMethod interface {
	WithMethod(method string)
}

// Optional interfaces for response type

//Renderer renders response to byte slice
type Renderer interface {
	Render() ([]byte, error)
}

//WithContentType returns custom content type for response
type WithContentType interface {
	ContentType() string
}

//WithHTTPStatus returns custom status code
type WithHTTPStatus interface {
	Status() int
}
