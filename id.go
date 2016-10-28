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

import "github.com/looplab/eventhorizon/ids"

// ID is an unique identifier interface
type ID interface {
	String() string
}

// IDFactory is a factory to generates new ids
type IDFactory interface {
	// NewID generates a new ID
	NewID() ID

	// CreateID generates a blank ID
	CreateID() ID
}

type defaultFactory struct{}

func (d *defaultFactory) NewID() ID {
	return ids.NewUUID()
}

func (d *defaultFactory) CreateID() ID {
	return ids.UUID("")
}

// default factory to type ids.UUID
var factoryID IDFactory = new(defaultFactory)

// RegisterID registers an id factory.
func RegisterID(factory IDFactory) {
	// test id factory
	id := factory.NewID()
	if id == nil {
		panic("eventhorizon: created id is nil")
	}

	factoryID = factory
}

// NewID creates a new id using the factory
// registered with RegisterID.
func NewID() ID {
	return factoryID.NewID()
}

// CreateID creates a blank id using the factory
// registered with RegisterID.
func CreateID() ID {
	return factoryID.CreateID()
}
