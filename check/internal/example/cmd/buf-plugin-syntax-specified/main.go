// Copyright 2024 Buf Technologies, Inc.
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

// Package main implements a simple plugin that checks that syntax is specified in every file.
//
// This is just demonstrating the additional functionality that check.Files have
// over FileDescriptors and FileDescriptorProtos.
//
// To use this plugin:
//
//	# buf.yaml
//	version: v2
//	lint:
//	  use:
//	   - STANDARD # omit if you do not want to use the rules builtin to buf
//	   - PLUGIN_SYNTAX_SPECIFIED
//	plugins:
//	  - plugin: buf-plugin-syntax-specified
//
// Note that the buf CLI implements this check by as a builtin rule, but this is just for example.
package main

import (
	"context"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
)

// syntaxSpecifiedRuleID is the Rule ID of the syntax specified Rule.
//
// This has a "PLUGIN_" prefix as the buf CLI has a rule "SYNTAX_SPECIFIED" builtin,
// and plugins/the buf CLI must have unique Rule IDs.
const syntaxSpecifiedRuleID = "PLUGIN_SYNTAX_SPECIFIED"

var (
	// syntaxSpecifiedRuleSpec is the RuleSpec for the syntax specified Rule.
	syntaxSpecifiedRuleSpec = &check.RuleSpec{
		ID:      syntaxSpecifiedRuleID,
		Default: true,
		Purpose: "Checks that syntax is specified.",
		Type:    check.RuleTypeLint,
		Handler: checkutil.NewFileRuleHandler(checkSyntaxSpecified, checkutil.WithoutImports()),
	}

	// spec is the Spec for the syntax specified plugin.
	spec = &check.Spec{
		Rules: []*check.RuleSpec{
			syntaxSpecifiedRuleSpec,
		},
	}
)

func main() {
	check.Main(spec)
}

func checkSyntaxSpecified(
	_ context.Context,
	responseWriter check.ResponseWriter,
	_ check.Request,
	file check.File,
) error {
	if file.IsSyntaxUnspecified() {
		syntax := file.FileDescriptorProto().GetSyntax()
		responseWriter.AddAnnotation(
			check.WithMessagef("Syntax should be specified but was %q.", syntax),
			check.WithDescriptor(file.FileDescriptor()),
		)
	}
	return nil
}
