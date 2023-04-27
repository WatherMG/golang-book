package tar

import (
	"archive/tar"
	"errors"
	"io"
	"os"

	"GolangBook/chapter10/lesson5/ex10.2/archive"
)

func init() {
	archive.InitFormats(archive.Format{Name: "tar", Str: "ustar\x0000", Offset: 257, List: list})
}

func list(f *os.File) ([]archive.FileHeader, error) {
	var headers []archive.FileHeader
	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		headers = append(headers,
			archive.FileHeader{Name: hdr.Name, Size: uint64(hdr.Size)})
	}
	return headers, nil
}
