// Copyright 2019 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linechart

import (
	"math"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func TestFormatters(t *testing.T) {
	tests := []struct {
		desc      string
		value     float64
		formatter ValueFormatter
		want      string
	}{
		{
			desc:      "Pretty duration formatter handles zero values",
			value:     0,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "0ns",
		},
		{
			desc:      "Pretty duration formatter handles minus minute values",
			value:     -1500,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "-25m",
		},
		{
			desc:      "Pretty duration formatter handles minus minute values",
			value:     -60,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "-1m",
		},
		{
			desc:      "Pretty duration formatter handles nanoseconds",
			value:     1.23e-7,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "123ns",
		},
		{
			desc:      "Pretty duration formatter handles microseconds",
			value:     1.23e-4,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "123µs",
		},
		{
			desc:      "Pretty duration formatter handles milliseconds",
			value:     0.123,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "123ms",
		},
		{
			desc:      "Pretty duration formatter handles seconds",
			value:     12,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "12s",
		},
		{
			desc:      "Pretty duration formatter handles minutes",
			value:     60,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "1m",
		},
		{
			desc:      "Pretty duration formatter handles hours",
			value:     2 * 60 * 60,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "2h",
		},
		{
			desc:      "Pretty duration formatter handles days",
			value:     5 * 24 * 60 * 60,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "5d",
		},
		{
			desc:      "Pretty minus duration formatter handles days",
			value:     -5 * 24 * 60 * 60,
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "-5d",
		},
		{
			desc:      "Pretty custom minute formatter with decimals handles days",
			value:     135,
			formatter: ValueFormatterSingleUnitDuration(time.Minute, 2),
			want:      "2.25h",
		},
		{
			desc:      "Pretty custom millisecond formatter with decimals handles minutes",
			value:     2525789,
			formatter: ValueFormatterSingleUnitDuration(time.Millisecond, 4),
			want:      "42.0965m",
		},
		{
			desc:      "Pretty custom nanosecond formatter with decimals handles days",
			value:     999999999999999,
			formatter: ValueFormatterSingleUnitDuration(time.Nanosecond, 8),
			want:      "11.57407407d",
		},
		{
			desc:      "Pretty custom minus nanosecond formatter with decimals handles days",
			value:     -999999999999999,
			formatter: ValueFormatterSingleUnitDuration(time.Nanosecond, 8),
			want:      "-11.57407407d",
		},
		{
			desc:      "Pretty custom minus nanosecond formatter without decimals handles microseconds",
			value:     -1500,
			formatter: ValueFormatterSingleUnitDuration(time.Nanosecond, 1),
			want:      "-1.5µs",
		},
		{
			desc:      "Pretty custom millisecond formatter with negative decimals handles minutes",
			value:     2525789,
			formatter: ValueFormatterSingleUnitDuration(time.Millisecond, -4),
			want:      "42m",
		},
		{
			desc:      "Pretty Second duration formatter handles NaN values",
			value:     math.NaN(),
			formatter: ValueFormatterSingleUnitSeconds,
			want:      "",
		},
		{
			desc:      "Pretty custom duration formatter handles NaN values",
			value:     math.NaN(),
			formatter: ValueFormatterSingleUnitDuration(time.Nanosecond, 8),
			want:      "",
		},
		{
			desc:      "Round formatter handles NaN values",
			value:     math.NaN(),
			formatter: ValueFormatterRound,
			want:      "",
		},
		{
			desc:      "Round formatter handles 0 values",
			value:     0,
			formatter: ValueFormatterRound,
			want:      "0",
		},
		{
			desc:      "Round formatter handles > x.5 values",
			value:     96.7,
			formatter: ValueFormatterRound,
			want:      "97",
		},
		{
			desc:      "Round formatter handles < x.5 values",
			value:     1621.2,
			formatter: ValueFormatterRound,
			want:      "1621",
		},
		{
			desc:      "Round formatter handles x.5 values",
			value:     6.5,
			formatter: ValueFormatterRound,
			want:      "7",
		},
		{
			desc:      "Round formatter handles minus > x.5 values",
			value:     -96.7,
			formatter: ValueFormatterRound,
			want:      "-97",
		},
		{
			desc:      "Round formatter handles minus < x.5 values",
			value:     -1621.2,
			formatter: ValueFormatterRound,
			want:      "-1621",
		},
		{
			desc:      "Round formatter handles minus x.5 values",
			value:     -6.5,
			formatter: ValueFormatterRound,
			want:      "-7",
		},
		{
			desc:      "Round formatter handles values with suffix",
			value:     96.7,
			formatter: ValueFormatterRoundWithSuffix("km"),
			want:      "97km",
		},
		{
			desc:      "Suffix formatter handles values with decimals",
			value:     11234567890.71234567890,
			formatter: ValueFormatterSuffix(4, " reqps"),
			want:      "11234567890.7123 reqps",
		},
		{
			desc:      "Suffix formatter handles NaN values",
			value:     math.NaN(),
			formatter: ValueFormatterSuffix(2, "test"),
			want:      "",
		},
		{
			desc:      "Suffix formatter handles 0 values",
			value:     0,
			formatter: ValueFormatterSuffix(2, "test"),
			want:      "0.00test",
		},
		{
			desc:      "Suffix formatters handles correctly percent suffix",
			value:     96.78,
			formatter: ValueFormatterSuffix(2, "%"),
			want:      "96.78%",
		},
		{
			desc:      "Round formatter handles values with percent suffix",
			value:     96.7,
			formatter: ValueFormatterRoundWithSuffix("%"),
			want:      "97%",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.formatter(tc.value)
			if diff := pretty.Compare(tc.want, got); diff != "" {
				t.Errorf("formatter => unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}
