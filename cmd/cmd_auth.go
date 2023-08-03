package cmd

import (
	"fmt"
)

const (
	Auth_API = "https://asciinema.org/connect/%s"
)

var info string = `Open the following URL in a web browser to link your install ID with your https://asciinema.org user account:
%s
This will associate all recordings uploaded from this machine (past and future ones) to your account, 
and allow you to manage them (change title/theme, delete) at https://asciinema.org.`

func (r *Runner) Auth() string {
	return fmt.Sprintf(info, fmt.Sprintf(Auth_API, cfg.ApiToken()))
}
