// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip25146

import (
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/locations"
	"github.com/googleapis/api-linter/rules/internal/utils"
	"github.com/jhump/protoreflect/desc"
)

var objectValues = &lint.FieldRule{
	Name: lint.NewRuleName(25146, "object-values"),
	OnlyIf: func(f *desc.FieldDescriptor) bool {
		return !utils.IsCommonProto(f.GetFile()) && f.GetMapValueType() != nil
	},
	LintField: func(f *desc.FieldDescriptor) []lint.Problem {
		if f.GetMapValueType().GetMessageType() != nil {
			return []lint.Problem{{
				Message:    "Avoid using objects as map values.",
				Descriptor: f,
				Location:   locations.FieldType(f),
			}}
		}
		return nil
	},
}
