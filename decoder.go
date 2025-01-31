package hdrx

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Decoder struct {
	r   *bufio.Reader
	pos Position
}

type DecodeError struct {
	Header string
	Pos    Position
	Err    error
}

var ErrEndOfHeaders = errors.New("end of headers")

func (err *DecodeError) Error() string {
	return fmt.Sprintf("decoding line %d, column %d: %v", err.Pos.Line, err.Pos.Column, err.Err)
}

func (err *DecodeError) Unwrap() error {
	return err.Err
}

type Position struct {
	Line, Column int
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:   bufio.NewReader(r),
		pos: Position{1, 1},
	}
}

func (dec *Decoder) ReadHeaders() (Headers, error) {
	res := Headers{}
	for {
		hdr, err := dec.ReadHeader()
		if err == ErrEndOfHeaders {
			return res, nil
		}
		if err != nil {
			return res, err
		}
		res = append(res, hdr)
	}
}

func (dec *Decoder) ReadHeader() (hdr Header, err error) {
	var block bool
	hdr.Key, block, err = dec.readKey()
	if err != nil {
		return
	}
	if block {
		hdr.Value, err = dec.readBlockValue()
	} else {
		hdr.Value, err = dec.readLineValue()
	}
	return
}

func (dec *Decoder) readKey() (key string, block bool, err error) {
	key, err = dec.r.ReadString(':') // XXX Consider " {" too.
	if err == io.EOF {
		err = ErrEndOfHeaders
	} else {
		key = key[:len(key)-1]
	}
	return
}

func (dec *Decoder) readBlockValue() (string, error) {
	panic("TODO: readBlockValue")
}

func (dec *Decoder) readLineValue() (val string, err error) {
	val, err = dec.r.ReadString('\n') // XXX Unescaping.
	if err == io.EOF {
		err = nil
	}
	val = strings.TrimSpace(val)
	return
}
