package github

import (
	"archive/tar"
	"compress/gzip"
	"io"

	"github.com/pkg/errors"
)

type Decompresser func(io.Reader) (io.Reader, error)

var (
	DefaultDecompresser              = TarGZIPDecompresser
	TarDecompresser     Decompresser = tarDecompresser
	TarGZIPDecompresser Decompresser = tarGZIPDecompresser
)

func tarDecompresser(r io.Reader) (io.Reader, error) {
	tr := tar.NewReader(r)
	_, err := tr.Next()
	return tr, errors.Wrap(err, "failed to create tar reader")
}

func tarGZIPDecompresser(r io.Reader) (io.Reader, error) {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create GZIP reader")
	}
	return TarDecompresser(gr)
}
