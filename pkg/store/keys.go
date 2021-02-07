package store

import "fmt"

const (
	usersKey    = "USER#%s"
	progressKey = "PROGRESS#%s"
	noteKey     = "NOTE#%s"
	timemapKey  = "TIMEMAP"
)

func UserPK(id string) string {
	return fmt.Sprintf(usersKey, id)
}

func ProgressSK(id string) string {
	return fmt.Sprintf(progressKey, id)
}

func NoteSK(id string) string {
	return fmt.Sprintf(noteKey, id)
}

func TimemapSK() string {
	return timemapKey
}
