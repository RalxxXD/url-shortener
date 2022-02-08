package msgpack

import (
	"github.com/RalxxXD/url-shortener/shortener"
	errs "github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrap(err, "serialize.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serialize.Redirect.Encode")
	}
	return rawMsg, nil
}
