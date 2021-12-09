package klz77

import (
	"bytes"
	"encoding/binary"
	"math"
)

// Decompress return decompressed data of input.
// The returned array of bytes is decompressed data of input, err is io.EOF when decompression buffer exhausted due invalid input, otherwise nil.
func Decompress(input []byte) ([]byte, error) {
	var result []byte
	buffer := bytes.NewBuffer(input)

	for buffer.Len() > 0 {
		flag, err := buffer.ReadByte()
		if err != nil {
			return result, err
		}

		for i := 0; i < 8; i++ {
			if (flag>>i)&1 == 1 {
				b, err := buffer.ReadByte()
				if err != nil {
					return result, err
				}

				result = append(result, b)
				continue
			}

			data := make([]byte, 2)
			if _, err = buffer.Read(data); err != nil {
				return result, err
			}

			token := int(binary.BigEndian.Uint16(data))
			position := token >> 4
			length := (token & 0x0F) + threshold

			if position == 0 {
				return result, nil
			}

			if position > len(result) {
				diff := int(math.Min(float64(position-len(result)), float64(length)))
				result = append(result, make([]byte, diff)...)
				length -= diff
			}

			if -position+length < 0 {
				a := len(result) - position
				b := len(result) - (position - length)

				if a >= 0 && b >= 0 {
					result = append(result, result[a:b]...)
				}
			} else {
				for i := 0; i < length; i++ {
					result = append(result, result[len(result)-position])
				}
			}
		}
	}

	return result, nil
}
