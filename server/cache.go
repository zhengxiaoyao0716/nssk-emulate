package server

var _cache = map[string]interface{}{}

// PushCache 置入缓存
func PushCache(key string, value interface{}) {
	_cache[key] = value
}

// GetStrCache 读出缓存
func GetStrCache(key string) string {
	value := _cache[key]
	if value == nil {
		return ""
	}
	return value.(string)
}
