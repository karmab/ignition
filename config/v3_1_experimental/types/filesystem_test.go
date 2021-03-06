// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"reflect"
	"testing"

	"github.com/coreos/ignition/v2/config/shared/errors"
	"github.com/coreos/ignition/v2/config/util"
	"github.com/coreos/ignition/v2/config/validate/report"
)

func TestFilesystemValidateFormat(t *testing.T) {
	type in struct {
		filesystem Filesystem
	}
	type out struct {
		err error
	}

	tests := []struct {
		in  in
		out out
	}{
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("ext4")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("btrfs")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: nil}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr(""), Path: util.StrToPtr("/")}},
			out: out{errors.ErrFormatNilWithOthers},
		},
		{
			in:  in{filesystem: Filesystem{Format: nil, Path: util.StrToPtr("/")}},
			out: out{errors.ErrFormatNilWithOthers},
		},
	}

	for i, test := range tests {
		err := test.in.filesystem.ValidateFormat()
		if !reflect.DeepEqual(report.ReportFromError(test.out.err, report.EntryError), err) {
			t.Errorf("#%d: bad error: want %v, got %v", i, test.out.err, err)
		}
	}
}

func TestFilesystemValidatePath(t *testing.T) {
	type in struct {
		filesystem Filesystem
	}
	type out struct {
		err error
	}

	tests := []struct {
		in  in
		out out
	}{
		{
			in:  in{filesystem: Filesystem{Path: util.StrToPtr("/foo")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Path: util.StrToPtr("")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Path: nil}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Path: util.StrToPtr("foo")}},
			out: out{err: errors.ErrPathRelative},
		},
	}

	for i, test := range tests {
		err := test.in.filesystem.ValidatePath()
		if !reflect.DeepEqual(report.ReportFromError(test.out.err, report.EntryError), err) {
			t.Errorf("#%d: bad error: want %v, got %v", i, test.out.err, err)
		}
	}
}

func TestLabelValidate(t *testing.T) {
	type in struct {
		filesystem Filesystem
	}
	type out struct {
		err error
	}

	tests := []struct {
		in  in
		out out
	}{
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("ext4"), Label: nil}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("ext4"), Label: util.StrToPtr("data")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("ext4"), Label: util.StrToPtr("thislabelistoolong")}},
			out: out{err: errors.ErrExt4LabelTooLong},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("btrfs"), Label: nil}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("btrfs"), Label: util.StrToPtr("thislabelisnottoolong")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("btrfs"), Label: util.StrToPtr("thislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolongthislabelistoolong")}},
			out: out{err: errors.ErrBtrfsLabelTooLong},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("xfs"), Label: nil}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("xfs"), Label: util.StrToPtr("data")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("xfs"), Label: util.StrToPtr("thislabelistoolong")}},
			out: out{err: errors.ErrXfsLabelTooLong},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("swap"), Label: nil}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("swap"), Label: util.StrToPtr("data")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("swap"), Label: util.StrToPtr("thislabelistoolong")}},
			out: out{err: errors.ErrSwapLabelTooLong},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("vfat"), Label: nil}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("vfat"), Label: util.StrToPtr("data")}},
			out: out{},
		},
		{
			in:  in{filesystem: Filesystem{Format: util.StrToPtr("vfat"), Label: util.StrToPtr("thislabelistoolong")}},
			out: out{err: errors.ErrVfatLabelTooLong},
		},
	}

	for i, test := range tests {
		err := test.in.filesystem.ValidateLabel()
		if !reflect.DeepEqual(report.ReportFromError(test.out.err, report.EntryError), err) {
			t.Errorf("#%d: bad error: want %v, got %v", i, test.out.err, err)
		}
	}
}
