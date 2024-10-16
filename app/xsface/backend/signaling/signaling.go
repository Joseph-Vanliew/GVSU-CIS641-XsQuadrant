package signaling

// Package signal contains utilities for exchanging SDP (Session Description Protocol)
// descriptions between examples. These functions handle encoding, decoding, and optionally
// compressing data to facilitate easier transmission.

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Allows compressing offer/answer to bypass terminal input limits.
// When set to 'true', the input will be compressed before encoding.
const compress = false

// MustReadStdin blocks execution until input is received from standard input (stdin).
// It reads and returns a trimmed line of input as a string.
func MustReadStdin() string {
	// Create a buffered reader for stdin
	r := bufio.NewReader(os.Stdin)

	var in string
	for {
		var err error
		// Read input until a newline character is encountered
		in, err = r.ReadString('\n')
		// If end-of-file is not reached and there's an error, panic
		if err != io.EOF {
			if err != nil {
				panic(err)
			}
		}
		// Remove any leading/trailing whitespace characters
		in = strings.TrimSpace(in)
		// If input is non-empty, break the loop
		if len(in) > 0 {
			break
		}
	}

	fmt.Println("")

	return in
}

// Encode converts an object to a JSON string, encodes it using base64, and
// optionally compresses the data before encoding if 'compress' is enabled.
func Encode(obj interface{}) string {
	// Marshal the object into a JSON byte array
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	// Compress the byte array if 'compress' is enabled
	if compress {
		b = zip(b)
	}

	// Encode the byte array to a base64 string and return
	return base64.StdEncoding.EncodeToString(b)
}

// Decode decodes a base64-encoded string back into an object.
// It also optionally decompresses the data after decoding if 'compress' is enabled.
func Decode(in string, obj interface{}) {
	// Decode the base64 string back to a byte array
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		panic(err)
	}

	// Decompress the byte array if 'compress' is enabled
	if compress {
		b = unzip(b)
	}

	// Unmarshal the JSON data into the provided object
	err = json.Unmarshal(b, obj)
	if err != nil {
		panic(err)
	}
}

// zip compresses a byte array using gzip and returns the compressed byte array.
func zip(in []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err := gz.Write(in)
	if err != nil {
		panic(err)
	}
	// Flush and close the gzip writer
	err = gz.Flush()
	if err != nil {
		panic(err)
	}
	err = gz.Close()
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

// unzip decompresses a gzip-compressed byte array and returns the decompressed byte array.
func unzip(in []byte) []byte {
	var b bytes.Buffer
	// Write the compressed byte array to the buffer
	_, err := b.Write(in)
	if err != nil {
		panic(err)
	}
	// Create a new gzip reader from the buffer
	r, err := gzip.NewReader(&b)
	if err != nil {
		panic(err)
	}
	// Read all data from the gzip reader (decompressed)
	res, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return res
}
