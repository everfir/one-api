package anthropic

var ModelList = []string{
	"claude-instant-1.2", "claude-2.0", "claude-2.1",
	"claude-3-haiku-20240307",
	"claude-3-5-haiku-20241022",
	"claude-3-5-haiku-latest",
	"claude-3-sonnet-20240229",
	"claude-3-opus-20240229",
	"claude-3-5-sonnet-20240620",
	"claude-3-5-sonnet-20241022",
	"claude-3-5-sonnet-latest",
	"claude-3-7-sonnet-20250219",
	"claude-sonnet-4-20250514",
}

var ModelCacheControlMap = map[string]bool{
	"claude-3-7-sonnet-20250219": true,
	"claude-sonnet-4-20250514":   true,
}
