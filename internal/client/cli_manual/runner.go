package cli_manual

import (
	"payment/internal/client"
)

func Run(jsonPathFromExeDir ...string) error {
	return client.Run(parseParamsAndRun, jsonPathFromExeDir...)
}
