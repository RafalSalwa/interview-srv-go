package swagger

import "embed"

var static embed.FS

func GetStaticFiles() embed.FS {
	return static
}
