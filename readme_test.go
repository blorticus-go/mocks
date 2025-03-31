package mocks_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/blorticus-go/mocks"
	"github.com/go-test/deep"
)

func TestReaderMethod(t *testing.T) {
	mockReader := mocks.NewReader().
		AddGoodRead([]byte{0x01, 0xff, 0x20}).
		AddGoodRead([]byte("this is a string")).
		AddEOF()

	if err := testARead(mockReader, false, false, []byte{0x01, 0xff, 0x20}); err != nil {
		t.Error(err)
	}
	if err := testARead(mockReader, false, false, []byte("this is a string")); err != nil {
		t.Error(err)
	}
	if err := testARead(mockReader, true, false, nil); err != nil {
		t.Error(err)
	}
}

func testARead(reader io.Reader, expectEOF bool, expectError bool, expectedReadBytes []byte) error {
	readBuf := make([]byte, 100)
	bytesRead, err := reader.Read(readBuf)

	if err != nil {
		if err == io.EOF {
			if expectEOF {
				return nil
			}
			return fmt.Errorf("got EOF, but expected error = (%s)", err.Error())
		}

		switch {
		case expectEOF:
			return fmt.Errorf("expected EOF, but got error = (%s)", err.Error())
		case expectError:
			return nil
		default:
			return fmt.Errorf("expected no error, but got error = (%s)", err.Error())
		}
	}

	if len(expectedReadBytes) != bytesRead {
		return fmt.Errorf("expected (%d) bytes on Read, but got (%d) bytes", len(expectedReadBytes), bytesRead)
	}

	if diff := deep.Equal(readBuf[:bytesRead], expectedReadBytes); diff != nil {
		return fmt.Errorf("%s", diff)
	}

	return nil
}
