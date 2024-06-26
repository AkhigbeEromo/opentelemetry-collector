// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1experimental"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestSample_MoveTo(t *testing.T) {
	ms := generateTestSample()
	dest := NewSample()
	ms.MoveTo(dest)
	assert.Equal(t, NewSample(), ms)
	assert.Equal(t, generateTestSample(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newSample(&otlpprofiles.Sample{}, &sharedState)) })
	assert.Panics(t, func() { newSample(&otlpprofiles.Sample{}, &sharedState).MoveTo(dest) })
}

func TestSample_CopyTo(t *testing.T) {
	ms := NewSample()
	orig := NewSample()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSample()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newSample(&otlpprofiles.Sample{}, &sharedState)) })
}

func TestSample_LocationIndex(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, pcommon.NewUInt64Slice(), ms.LocationIndex())
	internal.FillTestUInt64Slice(internal.UInt64Slice(ms.LocationIndex()))
	assert.Equal(t, pcommon.UInt64Slice(internal.GenerateTestUInt64Slice()), ms.LocationIndex())
}

func TestSample_LocationsStartIndex(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, uint64(0), ms.LocationsStartIndex())
	ms.SetLocationsStartIndex(uint64(1))
	assert.Equal(t, uint64(1), ms.LocationsStartIndex())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSample(&otlpprofiles.Sample{}, &sharedState).SetLocationsStartIndex(uint64(1)) })
}

func TestSample_LocationsLength(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, uint64(0), ms.LocationsLength())
	ms.SetLocationsLength(uint64(1))
	assert.Equal(t, uint64(1), ms.LocationsLength())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSample(&otlpprofiles.Sample{}, &sharedState).SetLocationsLength(uint64(1)) })
}

func TestSample_StacktraceIdIndex(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, uint32(0), ms.StacktraceIdIndex())
	ms.SetStacktraceIdIndex(uint32(1))
	assert.Equal(t, uint32(1), ms.StacktraceIdIndex())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSample(&otlpprofiles.Sample{}, &sharedState).SetStacktraceIdIndex(uint32(1)) })
}

func TestSample_Value(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, pcommon.NewInt64Slice(), ms.Value())
	internal.FillTestInt64Slice(internal.Int64Slice(ms.Value()))
	assert.Equal(t, pcommon.Int64Slice(internal.GenerateTestInt64Slice()), ms.Value())
}

func TestSample_Label(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, NewLabelSlice(), ms.Label())
	fillTestLabelSlice(ms.Label())
	assert.Equal(t, generateTestLabelSlice(), ms.Label())
}

func TestSample_Attributes(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, pcommon.NewUInt64Slice(), ms.Attributes())
	internal.FillTestUInt64Slice(internal.UInt64Slice(ms.Attributes()))
	assert.Equal(t, pcommon.UInt64Slice(internal.GenerateTestUInt64Slice()), ms.Attributes())
}

func TestSample_Link(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, uint64(0), ms.Link())
	ms.SetLink(uint64(1))
	assert.Equal(t, uint64(1), ms.Link())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSample(&otlpprofiles.Sample{}, &sharedState).SetLink(uint64(1)) })
}

func TestSample_TimestampsUnixNano(t *testing.T) {
	ms := NewSample()
	assert.Equal(t, pcommon.NewUInt64Slice(), ms.TimestampsUnixNano())
	internal.FillTestUInt64Slice(internal.UInt64Slice(ms.TimestampsUnixNano()))
	assert.Equal(t, pcommon.UInt64Slice(internal.GenerateTestUInt64Slice()), ms.TimestampsUnixNano())
}

func generateTestSample() Sample {
	tv := NewSample()
	fillTestSample(tv)
	return tv
}

func fillTestSample(tv Sample) {
	internal.FillTestUInt64Slice(internal.NewUInt64Slice(&tv.orig.LocationIndex, tv.state))
	tv.orig.LocationsStartIndex = uint64(1)
	tv.orig.LocationsLength = uint64(1)
	tv.orig.StacktraceIdIndex = uint32(1)
	internal.FillTestInt64Slice(internal.NewInt64Slice(&tv.orig.Value, tv.state))
	fillTestLabelSlice(newLabelSlice(&tv.orig.Label, tv.state))
	internal.FillTestUInt64Slice(internal.NewUInt64Slice(&tv.orig.Attributes, tv.state))
	tv.orig.Link = uint64(1)
	internal.FillTestUInt64Slice(internal.NewUInt64Slice(&tv.orig.TimestampsUnixNano, tv.state))
}
