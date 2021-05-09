package store

import "fmt"

const (
	usersKey         = "USER#%s"
	progressKey      = "PROGRESS#%s"
	noteKey          = "NOTE#%s"
	timemapKey       = "COURSE#%s#TIMEMAP#%s"
	courseTimemapKey = "COURSE#%s#TIMEMAP"
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

func CourseTimemapsSK(cID string) string {
	return fmt.Sprintf(courseTimemapKey, cID)
}

func TimemapSK(cID, tID string) string {
	return fmt.Sprintf(timemapKey, cID, tID)
}
