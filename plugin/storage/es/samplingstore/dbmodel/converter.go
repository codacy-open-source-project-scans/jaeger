// Copyright (c) 2024 The Jaeger Authors.
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
	"github.com/jaegertracing/jaeger/cmd/collector/app/sampling/model"
)

func FromThroughputs(throughputs []*model.Throughput) []Throughput {
	if throughputs == nil {
		return nil
	}
	ret := make([]Throughput, len(throughputs))
	for i, d := range throughputs {
		ret[i] = Throughput{
			Service:       d.Service,
			Operation:     d.Operation,
			Count:         d.Count,
			Probabilities: d.Probabilities,
		}
	}
	return ret
}

func ToThroughputs(throughputs []Throughput) []*model.Throughput {
	if throughputs == nil {
		return nil
	}
	ret := make([]*model.Throughput, len(throughputs))
	for i, d := range throughputs {
		ret[i] = &model.Throughput{
			Service:       d.Service,
			Operation:     d.Operation,
			Count:         d.Count,
			Probabilities: d.Probabilities,
		}
	}
	return ret
}
