/*
 * Copyright (c) 2019 Alec Randazzo
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 */

package binarytransforms

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

// LittleEndianBinaryToInt64 converts a little endian byte slice to int64.
func LittleEndianBinaryToInt64(inBytes []byte) (outInt64 int64, err error) {
	inBytesLength := len(inBytes)
	if inBytesLength > 8 || inBytesLength == 0 {
		err = fmt.Errorf("LittleEndianBinaryToInt64() received %v bytes but expected 1-8 bytes", inBytesLength)
		outInt64 = 0
		return
	}

	// Pad it to get to 8 bytes
	if inBytes[inBytesLength-1] >= 0x80 {
		if inBytesLength < 8 {
			bytesToPad := 8 - inBytesLength
			for i := 0; i < bytesToPad; i++ {
				inBytes = append(inBytes, 0xff)
			}
		}
	} else {
		if inBytesLength < 8 {
			bytesToPad := 8 - inBytesLength
			for i := 0; i < bytesToPad; i++ {
				inBytes = append(inBytes, 0x00)
			}
		}
	}
	outInt64 = int64(binary.LittleEndian.Uint64(inBytes))
	return
}

// LittleEndianBinaryToUInt64 converts a little endian byte slice to uint64.
func LittleEndianBinaryToUInt64(inBytes []byte) (outUInt64 uint64, err error) {
	inBytesLength := len(inBytes)
	if inBytesLength > 8 || inBytesLength == 0 {
		err = fmt.Errorf("LittleEndianBinaryToUInt64() received %v bytes but expected 1-8 bytes", inBytesLength)
		outUInt64 = 0
		return
	}

	// Pad it to get to 8 bytes
	if inBytesLength < 8 {
		bytesToPad := 8 - inBytesLength
		for i := 0; i < bytesToPad; i++ {
			inBytes = append(inBytes, 0x00)
		}
	}
	outUInt64 = binary.LittleEndian.Uint64(inBytes)
	return
}

// LittleEndianBinaryToUInt16 converts a little endian byte slice to uint16.
func LittleEndianBinaryToUInt16(inBytes []byte) (outUInt16 uint16, err error) {
	inBytesLength := len(inBytes)
	if inBytesLength != 2 {
		err = fmt.Errorf("LittleEndianBinaryToUInt16() received %v bytes but expected 2 bytes", inBytesLength)
		outUInt16 = 0
		return
	}
	outUInt16 = binary.LittleEndian.Uint16(inBytes)
	return
}

// LittleEndianBinaryToUInt32 converts a little endian byte slice to uint32.
func LittleEndianBinaryToUInt32(inBytes []byte) (outUInt32 uint32, err error) {
	inBytesLength := len(inBytes)
	if inBytesLength != 4 {
		err = fmt.Errorf("LittleEndianBinaryToUInt32() received %v bytes but expected 4 bytes", inBytesLength)
		outUInt32 = 0
		return
	}
	outUInt32 = binary.LittleEndian.Uint32(inBytes)
	return
}

// UnicodeBytesToASCII converts a byte slice of unicode characters to ASCII
func UnicodeBytesToASCII(unicodeBytes []byte) (asciiString string, err error) {
	inBytesLength := len(unicodeBytes)
	if inBytesLength == 0 {
		err = errors.New("UnicodeBytesToASCII() received no bytes")
		asciiString = ""
		return
	}

	unicodeString := string(unicodeBytes)
	asciiString = strings.Replace(unicodeString, "\x00", "", -1)
	return
}
