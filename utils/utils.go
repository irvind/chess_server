package utils

import "os"

func GetenvOrDefault(varname string, defaultVal string) string {
	val, ok := os.LookupEnv(varname)
	if ok {
		return val
	} else {
		return defaultVal
	}	
}