/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tribe_event

import "github.com/intelsdi-x/snap/core"

const (
	PluginAdded = "Tribe.PluginAdded"
)

type AddPluginEvent struct {
	Agreement struct {
		Name string
	}
	Plugin struct {
		Name    string
		Type    core.PluginType
		Version int
	}
}

func (e AddPluginEvent) Namespace() string {
	return PluginAdded
}
