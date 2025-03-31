# mocks

Simple mocks for a set of built-in interfaces.

## Overview

Currently, the following mock is implemented:
* `mock.Reader`: simulates an `io.Reader`

## Examples

```go
package mypackage_test

import (
    "testing"
    "github.com/go-test/deep"
    "github.com/blorticus-go/mocks"
)

func TestReaderMethod(t *testing.T) {
    mockReader := mocks.NewReader().
        AddGoodRead([]bytes{0x01, 0xff, 0x20}).
        AddGoodRead([]bytes("this is a string")).
        AddEOF()

    if err := testARead(mockReader, false, false, []bytes{0x01, 0xff, 0x20}); err != nil {
        t.Error(err)
    }
    if err := testARead(mockReader, false, false, []bytes("this is a string")); err != nil {
        t.Error(err)
    }
    if err := testARead(mockReader, true, false, nil); err != nil {
        t.Error(err)
    }
}

func testARead(reader io.Reader, expectEOF bool, expectError bool, expectedReadBytes) error {
    readBuf := make([]bytes, 100)
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
        return fmt.Errorf(diff)
    }

    return nil
}

```
