// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2018 Uber Technologies, Inc.
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

package dbmodel

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CompareJSONSpans(t *testing.T, expected *Span, actual *Span) {
	sortJSONSpan(expected)
	sortJSONSpan(actual)

	if !assert.EqualValues(t, expected, actual) {
		for _, err := range pretty.Diff(expected, actual) {
			t.Log(err)
		}
		out, err := json.Marshal(actual)
		require.NoError(t, err)
		t.Logf("Actual trace: %s", string(out))
	}
}

func sortJSONSpan(span *Span) {
	sortJSONTags(span.Tags)
	sortJSONLogs(span.Logs)
	sortJSONProcess(span.Process)
}

type JSONTagByKey []KeyValue

func (t JSONTagByKey) Len() int           { return len(t) }
func (t JSONTagByKey) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t JSONTagByKey) Less(i, j int) bool { return t[i].Key < t[j].Key }

func sortJSONTags(tags []KeyValue) {
	sort.Sort(JSONTagByKey(tags))
}

type JSONLogByTimestamp []Log

func (t JSONLogByTimestamp) Len() int           { return len(t) }
func (t JSONLogByTimestamp) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t JSONLogByTimestamp) Less(i, j int) bool { return t[i].Timestamp < t[j].Timestamp }

func sortJSONLogs(logs []Log) {
	sort.Sort(JSONLogByTimestamp(logs))
	for i := range logs {
		sortJSONTags(logs[i].Fields)
	}
}

func sortJSONProcess(process Process) {
	sortJSONTags(process.Tags)
}
