// Copyright 2019, OpenTelemetry Authors
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

// metric package provides an API for reporting diagnostic
// measurements using three basic kinds of instruments (or four, if
// calling one special case a separate one).
//
// The three basic kinds are:
//
// - counters
// - gauges
// - measures
//
// There is also a special case of a gauge instrument called
// observer. It will be discussed later.
//
// All instruments report either float64 or int64 values.
//
// The primary object that handles metrics is Meter. The
// implementation of the Meter is provided by SDK. Normally, the Meter
// is used directly only for the LabelSet generation, batch recording
// and the handle destruction.
//
// LabelSet is a set of keys and values that are in a suitable,
// optimized form to be used by Meter.
//
// Counters are instruments that are reporting a quantity or a sum. An
// example could be bank account balance or bytes downloaded. Counters
// can be created with either NewFloat64Counter or
// NewInt64Counter. Counters expect non-negative values by default to
// be reported. This can be changed with the WithNonMonotonic option
// passed to the New*Counter function - this allows reporting negative
// values. To report the new value, use an Add function.
//
// Gauges are instruments that are reporting a current state of a
// value. An example could be voltage or temperature. Gauges can be
// created with either NewFloat64Gauge or NewInt64Gauge. Gauges by
// default have no limitations about reported values - they can be
// less or greater than the last reported value. This can be changed
// with the WithMonotonic option passed to the New*Gauge function -
// this permits the reported values only to go up. To report a new
// value, use the Set function.
//
// Measures are instruments that are reporting values that are
// recorded separately to figure out some statistical properties from
// those values (like average). An example could be temperature over
// time or lines of code in the project over time. Measures can be
// created with either NewFloat64Measure or NewInt64Measure. Measures
// by default take only non-negative values. This can be changed with
// the WithSigned option passed to the New*Measure function - this
// allows reporting negative values too. To report a new value, use
// the Record function.
//
// A special case of a gauge is observer. It has the same role as
// gauge, but reports values in a different way. The observer can be
// created with NewFloat64Observer or NewInt64Observer and then
// registered within a Meter with either RegisterInt64Observer or
// RegisterFloat64Observer functions. These take a callback for
// reporting the values. The callback will be called by the Meter when
// it deems it necessary.
//
// All the basic kinds of instruments (so, not observers) also support
// creating handles for a potentially more efficient reporting. The
// handles have the same function names as the instruments (so counter
// handle has Add, gauge handle has Set and measure handle has
// Record). Handles can be created with the GetHandle function of the
// respective instrument.
package metric // import "go.opentelemetry.io/api/metric"
