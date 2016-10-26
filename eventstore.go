// Copyright (c) 2014 - Max Ekman <max@looplab.se>
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

package eventhorizon

import (
	"errors"
)

// ErrNoEventsToAppend is when no events are available to append.
var ErrNoEventsToAppend = errors.New("no events to append")

// EventStore is an interface for an event sourcing event store.
type EventStore interface {
	// Save appends all events in the event stream to the store.
	Save(events []Event, originalVersion int) error

	// Load loads all events for the aggregate id from the store.
	Load(ID) ([]Event, error)
}

// AggregateRecord is a stored record of an aggregate in form of its events.
// NOTE: Not currently used.
type AggregateRecord interface {
	AggregateID() ID
	Version() int
	EventRecords() []EventRecord
}

// EventRecord is a single event record with timestamp
// NOTE: Not currently used.
type EventRecord interface {
	Type() string
	Version() int
	Events() []Event
}
