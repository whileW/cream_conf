package cream_conf

import "strings"

// splitConfigKey split config key to key and sub key
func splitConfigKey(key string) []string {
	key = strings.ToLower(key)
	key = strings.ReplaceAll(key, "/", ".")
	return strings.Split(key, ".")
}
