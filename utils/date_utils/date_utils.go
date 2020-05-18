package date_utils

import "time"

// Constant for datetime format while saving in database
const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// this func return current datetime
func GetNow() time.Time{
	return time.Now().UTC()
}


// Two functions which return current datetime with giving format
func GetNowString() string{
	return GetNow().Format(apiDateLayout)
}

func GetNowDbFormat() string{
	return GetNow().Format(apiDbLayout)
}
