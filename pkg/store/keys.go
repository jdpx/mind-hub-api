package store

import "fmt"

const (
	user_key     = "USER#%s"
	progress_key = "PROGRESS#%s"
	note_key     = "NOTE#%s"
	timemap_key  = "TIMEMAP"
)

func UserPK(id string) string {
	return fmt.Sprintf(user_key, id)
}

func ProgressSK(id string) string {
	return fmt.Sprintf(progress_key, id)
}

func NoteSK(id string) string {
	return fmt.Sprintf(note_key, id)
}

func TimemapSK() string {
	return timemap_key
}
