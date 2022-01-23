// Copyright © 2021 Rak Laptudirm <raklaptudirm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package completion

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/raklaptudirm/krypt/internal/cmd/completion/bash"
	"github.com/raklaptudirm/krypt/internal/cmd/completion/fish"
	"github.com/raklaptudirm/krypt/internal/cmd/completion/pwsh"
	"github.com/raklaptudirm/krypt/internal/cmd/completion/zsh"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "completion",
		Short: "completion generates completion scripts for shells",
		Args:  cobra.NoArgs,
		Long: heredoc.Doc(`
			Generate the autocompletion script for krypt for the specified shell.      
			See each sub-command's help for details on how to use the generated script.
		`),
		Hidden: true,
	}

	// supported shells
	cmd.AddCommand(zsh.NewCmd())
	cmd.AddCommand(bash.NewCmd())
	cmd.AddCommand(fish.NewCmd())
	cmd.AddCommand(pwsh.NewCmd())

	return cmd
}