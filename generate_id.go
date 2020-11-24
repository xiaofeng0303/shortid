package shortid

// GenerateID  生成短ID算法 generate a id with prefix
func GenerateID(prefix string) string {
	id, _ := Generate()
	return prefix + "-" + id
}
