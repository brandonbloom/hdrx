package hdrx

import "io"

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

func (dec *Decoder) ReadHeaders() (Headers, error) {
	return Headers{}, nil // XXX
}

func (dec *Decoder) ReadHeader() (Header, error) {
	return Header{}, nil // XXX
}
