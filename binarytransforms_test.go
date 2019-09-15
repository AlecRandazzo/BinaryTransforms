/*
 * Copyright (c) 2019 Alec Randazzo
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 */

package BinaryTransforms

import (
	"reflect"
	"testing"
)

func TestLittleEndianBinaryToInt64(t *testing.T) {
	type args struct {
		inBytes []byte
	}
	tests := []struct {
		name         string
		args         args
		wantOutInt64 int64
	}{
		{
			name:         "Testing with a random byte slice.",
			args:         args{inBytes: []byte{20, 21, 255}},
			wantOutInt64: -60140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutInt64 := LittleEndianBinaryToInt64(tt.args.inBytes); !reflect.DeepEqual(gotOutInt64, tt.wantOutInt64) {
				t.Errorf("LittleEndianBinaryToInt64() = %v, want %v", gotOutInt64, tt.wantOutInt64)
			}
		})
	}
}

func TestLittleEndianBinaryToUInt16(t *testing.T) {
	type args struct {
		inBytes []byte
	}
	tests := []struct {
		name          string
		args          args
		wantOutUInt16 uint16
	}{
		{
			name:          "Testing with a random byte slice.",
			args:          args{inBytes: []byte{246, 201}},
			wantOutUInt16: 51702,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutUInt16, err := LittleEndianBinaryToUInt16(tt.args.inBytes); !reflect.DeepEqual(gotOutUInt16, tt.wantOutUInt16) && err == nil {
				t.Errorf("LittleEndianBinaryToUInt16() = %v, want %v", gotOutUInt16, tt.wantOutUInt16)
			}
		})
	}
}

func TestLittleEndianBinaryToUInt32(t *testing.T) {
	type args struct {
		inBytes []byte
	}
	tests := []struct {
		name          string
		args          args
		wantOutUInt32 uint32
	}{
		{
			name:          "Testing with a random byte slice.",
			args:          args{inBytes: []byte{246, 201, 201, 0}},
			wantOutUInt32: 13224438,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutUInt32, err := LittleEndianBinaryToUInt32(tt.args.inBytes); !reflect.DeepEqual(gotOutUInt32, tt.wantOutUInt32) && err == nil {
				t.Errorf("LittleEndianBinaryToUInt32() = %v, want %v", gotOutUInt32, tt.wantOutUInt32)
			}
		})
	}
}

func TestLittleEndianBinaryToUInt64(t *testing.T) {
	type args struct {
		inBytes []byte
	}
	tests := []struct {
		name          string
		args          args
		wantOutUInt64 uint64
	}{
		{
			name:          "Testing with a random byte slice.",
			args:          args{inBytes: []byte{20, 21, 255}},
			wantOutUInt64: 16717076,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutUInt64 := LittleEndianBinaryToUInt64(tt.args.inBytes); !reflect.DeepEqual(gotOutUInt64, tt.wantOutUInt64) {
				t.Errorf("LittleEndianBinaryToUInt64() = %v, want %v", gotOutUInt64, tt.wantOutUInt64)
			}
		})
	}
}

func TestUnicodeBytesToASCII(t *testing.T) {
	type args struct {
		unicodeBytes []byte
	}
	tests := []struct {
		name            string
		args            args
		wantAsciiString string
	}{
		{
			name:            "Testing with a random byte slice.",
			args:            args{unicodeBytes: []byte{116, 0, 101, 0, 115, 0, 116}},
			wantAsciiString: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAsciiString := UnicodeBytesToASCII(tt.args.unicodeBytes); !reflect.DeepEqual(gotAsciiString, tt.wantAsciiString) {
				t.Errorf("UnicodeBytesToASCII() = %v, want %v", gotAsciiString, tt.wantAsciiString)
			}
		})
	}
}
