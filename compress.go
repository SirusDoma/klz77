package klz77

import (
	"bytes"
	"encoding/binary"
	"math"
)

const (
	threshold       = 3
	windowSize      = 0x1000
	windowMask      = windowSize - 1
	maxBufferLength = 0x0F + threshold
)

// Compress return compressed data of input.
// The returned array of bytes is compressed data of input; error is always nil but included to match with Decompress.
func Compress(input []byte) ([]byte, error) {
	var result = make([]byte, 0)
	cursor := windowSize
	bit := 0

	input = append(make([]byte, windowSize), input...)
	for cursor < len(input) {
		flag := 0
		buffer := bytes.NewBuffer(make([]byte, 0))

		for i := 0; i < 8; i++ {
			if cursor < len(input) {
				if position, length := findMatchWindow(input, cursor); position >= 0 && length >= 0 {
					data := (position << 4) | ((length - threshold) & 0x0F)
					encoded := make([]byte, 2)
					binary.BigEndian.PutUint16(encoded, uint16(data))

					if _, err := buffer.Write(encoded); err != nil {
						return result, nil
					}

					bit = 0
					cursor += length
				} else {
					if err := buffer.WriteByte(input[cursor]); err != nil {
						return result, nil
					}

					bit = 1
					cursor++
				}
			} else {
				bit = 0
			}

			flag = (flag >> 1) | ((bit & 1) << 7)
		}

		result = append(result, byte(flag))
		result = append(result, buffer.Bytes()...)
	}

	result = append(result, make([]byte, 2)...)
	return result, nil
}

func findMatchWindow(input []byte, cursor int) (int, int) {
	start := int(math.Max(float64(cursor-windowMask), 0))

	for n := maxBufferLength; n >= threshold; n-- {
		end := int(math.Min(float64(cursor+n), float64(len(input))))
		if end-cursor < threshold {
			break
		}

		criteria := input[cursor:end]
		maxLookup := int(math.Min(float64(end-n), float64(len(input))))
		if index := bytes.LastIndex(input[start:maxLookup], criteria); index != -1 {
			index += start
			position := cursor - index
			length := len(criteria)

			return position, length
		}
	}

	return -1, -1
}
