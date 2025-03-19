package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\r\n")
	if len(lines) < 1 {
		return nil, errors.New("invalid request")
	}

	requestLine, err := requestLineFromString(lines[0])
	if err != nil {
		return nil, err
	}

	return &Request{RequestLine: requestLine}, nil
}

func requestLineFromString(line string) (RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return RequestLine{}, errors.New("invalid request line")
	}

	method := parts[0]
	path := parts[1]
	httpVersion := parts[2]

	if strings.ContainsFunc(method, func(r rune) bool {
		return r < 'A' || r > 'Z'
	}) {
		return RequestLine{}, errors.New("invalid method")
	}

	if !strings.HasPrefix(path, "/") {
		return RequestLine{}, errors.New("invalid path")
	}

	if httpVersion != "HTTP/1.1" {
		return RequestLine{}, errors.New("unsupported http version")
	}

	return RequestLine{
		HttpVersion:   "1.1",
		RequestTarget: path,
		Method:        method,
	}, nil
}
