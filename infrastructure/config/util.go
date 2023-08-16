package config

import "time"

// GetString 설정파일에서 문자열을 가져온다.
func GetString(key string) string {
	val := v.GetString(key)
	if val == "" {
		panic("invalid config key:" + key)
	}
	return val
}

// GetInt 설정파일에서 정수를 가져온다.
func GetInt(key string) int {
	val := v.GetInt(key)
	if val == 0 {
		panic("invalid config key:" + key)
	}
	return val
}

// GetInt64 설정파일에서 정수64를 가져온다.
func GetInt64(key string) int64 {
	val := v.GetInt64(key)
	if val == 0 {
		panic("invalid config key:" + key)
	}
	return val
}

// GetIntSlice 설정파일에서 정수 슬라이스를 가져온다.
func GetIntSlice(key string) []int {
	val := v.GetIntSlice(key)
	if len(val) == 0 {
		panic("invalid config key:" + key)
	}
	return val
}

// GetBool 설정파일에서 불리언을 가져온다.
func GetBool(key string) bool {
	return v.GetBool(key)
}

// GetDuration 설정파일에서 시간을 가져온다.
func GetDuration(key string) time.Duration {
	val := v.GetDuration(key)
	if val == 0 {
		panic("invalid config key:" + key)
	}
	return val
}
