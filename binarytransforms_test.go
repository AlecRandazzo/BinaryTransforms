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
	"fmt"
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
		wantErr      bool
	}{
		{
			name:         "Testing with a random byte slice.",
			args:         args{inBytes: []byte{20, 21, 255}},
			wantOutInt64: -60140,
			wantErr:      false,
		},
		{
			name:         "Testing with null byte slice.",
			args:         args{inBytes: []byte{}},
			wantOutInt64: 0,
			wantErr:      true,
		},
		{
			name:         "Testing with byte slice larger than 8 bytes.",
			args:         args{inBytes: []byte{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
			wantOutInt64: 0,
			wantErr:      true,
		},
		{
			name:         "Testing with byte slice with 8 bytes.",
			args:         args{inBytes: []byte{1}},
			wantOutInt64: 1,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutInt64, err := LittleEndianBinaryToInt64(tt.args.inBytes)
			fmt.Println(gotOutInt64, tt.wantOutInt64)
			if gotOutInt64 != tt.wantOutInt64 || (err != nil) != tt.wantErr {
				t.Errorf("got = %v, want %v", gotOutInt64, tt.wantOutInt64)
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
		wantErr       bool
	}{
		{
			name:          "Testing with a random byte slice.",
			args:          args{inBytes: []byte{246, 201}},
			wantOutUInt16: 51702,
			wantErr:       false,
		},
		{
			name:          "Testing with null byte slice.",
			args:          args{inBytes: []byte{}},
			wantOutUInt16: 0,
			wantErr:       true,
		},
		{
			name:          "Testing with byte slice larger than 8 bytes.",
			args:          args{inBytes: []byte{20, 20, 20, 20, 20, 20, 20, 20, 20, 20}},
			wantOutUInt16: 0,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutUInt16, err := LittleEndianBinaryToUInt16(tt.args.inBytes); !reflect.DeepEqual(gotOutUInt16, tt.wantOutUInt16) || (err != nil) != tt.wantErr {
				t.Errorf("got = %v, want %v", gotOutUInt16, tt.wantOutUInt16)
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
		wantErr       bool
	}{
		{
			name:          "Testing with a random byte slice.",
			args:          args{inBytes: []byte{246, 201, 201, 0}},
			wantOutUInt32: 13224438,
			wantErr:       false,
		},
		{
			name:          "Testing with null byte slice.",
			args:          args{inBytes: []byte{}},
			wantOutUInt32: 0,
			wantErr:       true,
		},
		{
			name:          "Testing with byte slice larger than 8 bytes.",
			args:          args{inBytes: []byte{20, 20, 20, 20, 20, 20, 20, 20, 20, 20}},
			wantOutUInt32: 0,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutUInt32, err := LittleEndianBinaryToUInt32(tt.args.inBytes); !reflect.DeepEqual(gotOutUInt32, tt.wantOutUInt32) || (err != nil) != tt.wantErr {
				t.Errorf("got = %v, want %v", gotOutUInt32, tt.wantOutUInt32)
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
		wantErr       bool
	}{
		{
			name:          "Testing with a random byte slice.",
			args:          args{inBytes: []byte{20, 21, 255}},
			wantOutUInt64: 16717076,
			wantErr:       false,
		},
		{
			name:          "Testing with null byte slice.",
			args:          args{inBytes: []byte{}},
			wantOutUInt64: 0,
			wantErr:       true,
		},
		{
			name:          "Testing with byte slice larger than 8 bytes.",
			args:          args{inBytes: []byte{20, 20, 20, 20, 20, 20, 20, 20, 20, 20}},
			wantOutUInt64: 0,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutUInt64, err := LittleEndianBinaryToUInt64(tt.args.inBytes); !reflect.DeepEqual(gotOutUInt64, tt.wantOutUInt64) || (err != nil) != tt.wantErr {
				t.Errorf("got = %v, want %v", gotOutUInt64, tt.wantOutUInt64)
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
		wantASCIIString string
		wantErr         bool
	}{
		{
			name:            "Testing with a random byte slice.",
			args:            args{unicodeBytes: []byte{116, 0, 101, 0, 115, 0, 116}},
			wantASCIIString: "test",
			wantErr:         false,
		},
		{
			name:            "Testing with null byte slice.",
			args:            args{unicodeBytes: []byte{}},
			wantASCIIString: "",
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotASCIIString, err := UnicodeBytesToASCII(tt.args.unicodeBytes); !reflect.DeepEqual(gotASCIIString, tt.wantASCIIString) || (err != nil) != tt.wantErr {
				t.Errorf("got = %v, want %v", gotASCIIString, tt.wantASCIIString)
			}
		})
	}
}
