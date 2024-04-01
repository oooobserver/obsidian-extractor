package main

func strip_file_extension(base string) string {
	i := 0
	for ; i < len(base); i++ {
		if base[i] == '.' {
			break
		}
	}

	return base[:i]
}
