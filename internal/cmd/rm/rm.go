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

package rm

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/raklaptudirm/krypt/internal/auth"
	"github.com/raklaptudirm/krypt/internal/cmdutil"
	"github.com/raklaptudirm/krypt/pkg/pass"
	"github.com/spf13/cobra"
)

func NewCmd(c *cmdutil.Context) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rm [name]",
		Short: "remove a password from krypt",
		Args:  cobra.ExactArgs(1),
		Long: heredoc.Doc(`
			Logout clears the file which stores your database key,
			so that accessing the passwords requires logging in with
			the master password.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			pass, err := pass.GetS(c.PassManager, args[0], c.Creds.Key)
			if err != nil {
				return err
			}

			return rm(c.PassManager, c.Creds, pass.Checksum)
		},
	}

	return cmd
}

func rm(passMan pass.Manager, creds *auth.Creds, checksum []byte) error {
	if !creds.LoggedIn() {
		return cmdutil.ErrNoLogin
	}

	err := passMan.Delete(checksum)
	if err != nil {
		return err
	}

	fmt.Println("Deleted password.")
	return nil
}
