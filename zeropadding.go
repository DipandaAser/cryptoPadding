package crypto_padding

import (
    "bytes"
    "errors"
    "fmt"
)

type ZeroPadding struct {}

func (padding ZeroPadding) Pad(data []byte, blockSize int) (output []byte, err error) {
    if blockSize < 1 || blockSize >= 256 {
        return output, errors.New(fmt.Sprintf("blocksize is out of bounds: %v", blockSize))
    }
    var paddingBytes = padSize(len(data), blockSize)
    paddingSlice := bytes.Repeat([]byte{byte(0)}, paddingBytes)
    output = append(data, paddingSlice...)
    fmt.Println(paddingSlice)
    return output, nil
}

// May not behave properly if the last character of the unpadded data is a zero.
func (padding ZeroPadding) Unpad(data []byte, blockSize int) (output []byte, err error) {
    var dataLen = len(data)
    if dataLen % blockSize != 0 {
        return output, errors.New("data's length isn't a multiple of blockSize")
    }
    var paddingBytes = 0
    for data[dataLen - 1 - paddingBytes] == 0 {
        paddingBytes++
    }
    if paddingBytes > blockSize || paddingBytes <= 0 {
        return output, errors.New(fmt.Sprintf("invalid padding found: %v", paddingBytes))
    }
    output = data[0:dataLen - paddingBytes]
    return output, nil
}
