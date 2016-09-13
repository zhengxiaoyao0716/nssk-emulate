package server

var _cache = map[string]interface{}{}

// PushCache 置入缓存
func PushCache(key string, value interface{}) {
	_cache[key] = value
}

// GetStrCache 读出缓存
func GetStrCache(key string) string {
	value, ok := _cache[key]
	if !ok {
		return ""
	}
	return value.(string)
}

// GetMapCache 读出缓存
func GetMapCache(key string) map[string]interface{} {
	value, ok := _cache[key]
	if !ok {
		return nil
	}
	return value.(map[string]interface{})
}

// GetCache 读出缓存
func GetCache(key string) interface{} {
	value, ok := _cache[key]
	if ok {
		return value
	}
	return nil
}
