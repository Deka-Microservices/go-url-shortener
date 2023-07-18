package consts

import "fmt"

const (
	VERSION_MAJOR = 0
	VERSION_MINOR = 4
	VERSION_PATCH = 1
)

func Version() string {
	return fmt.Sprintf("%d.%d.%d", VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH)
}
