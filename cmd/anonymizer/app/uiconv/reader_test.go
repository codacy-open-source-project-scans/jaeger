// Copyright (c) 2020 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uiconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestReaderTraceSuccess(t *testing.T) {
	inputFile := "fixtures/trace_success.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	s1, err := r.NextSpan()
	require.NoError(t, err)
	assert.Equal(t, "a071653098f9250d", s1.OperationName)
	assert.Equal(t, 1, r.spansRead)
	assert.False(t, r.eofReached)

	r.spansRead = 999

	s2, err := r.NextSpan()
	require.NoError(t, err)
	assert.Equal(t, "471418097747d04a", s2.OperationName)
	assert.Equal(t, 1000, r.spansRead)
	assert.True(t, r.eofReached)

	_, err = r.NextSpan()
	require.Equal(t, errNoMoreSpans, err)
	assert.Equal(t, 1000, r.spansRead)
	assert.True(t, r.eofReached)
}

func TestReaderTraceNonExistent(t *testing.T) {
	inputFile := "fixtures/trace_non_existent.json"
	_, err := newSpanReader(inputFile, zap.NewNop())
	require.Contains(t, err.Error(), "cannot open captured file")
}

func TestReaderTraceEmpty(t *testing.T) {
	inputFile := "fixtures/trace_empty.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	_, err = r.NextSpan()
	require.Contains(t, err.Error(), "cannot read file")
	assert.Equal(t, 0, r.spansRead)
	assert.True(t, r.eofReached)
}

func TestReaderTraceWrongFormat(t *testing.T) {
	inputFile := "fixtures/trace_wrong_format.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	_, err = r.NextSpan()
	require.Equal(t, "file must begin with '['", err.Error())
	assert.Equal(t, 0, r.spansRead)
	assert.True(t, r.eofReached)
}

func TestReaderTraceInvalidJson(t *testing.T) {
	inputFile := "fixtures/trace_invalid_json.json"
	r, err := newSpanReader(inputFile, zap.NewNop())
	require.NoError(t, err)

	_, err = r.NextSpan()
	require.Contains(t, err.Error(), "cannot unmarshal span")
	assert.Equal(t, 0, r.spansRead)
	assert.True(t, r.eofReached)
}
