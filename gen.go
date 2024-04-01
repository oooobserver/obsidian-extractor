package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"regexp"
)

const start = `## Content
***
`

type title struct {
	level int
	name  string
}

func gen_dir(filename string) []byte {
	files, err := os.ReadDir(filename)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}

	res := start
	for _, file := range files {
		if file.IsDir() {
			// If directory name is imgs, no need to generate
			if file.Name() != "imgs" {
				gen_dir_helper(filename, file, 0, &res)
			}
		} else {
			tmp := fmt.Sprintf("[[%s]]", strip_file_extension(file.Name()))
			res = res + "* " + tmp + "\n"
		}
	}

	return []byte(res)
}

func gen_dir_helper(base string, file fs.DirEntry, indent int, res *string) {
	path := base + "/" + file.Name()

	// Add this directory name
	ind := ""
	for i := 0; i < indent; i++ {
		ind += "\t"
	}
	*res += "* " + ind + file.Name() + "\n"

	dir, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer dir.Close()

	entries, err := dir.ReadDir(0)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			indent += 1
			gen_dir_helper(path, entry, indent, res)
		} else {
			tmp := fmt.Sprintf("[[%s]]", strip_file_extension(entry.Name()))
			*res += "* " + ind + tmp + "\n"
		}
	}
}

func gen_file(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %v", filename)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", filename)
	}
	file.Close()

	res := start

	lt := get_all_titles(string(content))
	for _, t := range lt {
		indent := ""
		for i := 2; i < t.level; i++ {
			indent += "\t"
		}

		tmp := fmt.Sprintf("[[#%s]]", t.name)

		res = res + indent + "* " + tmp + "\n"
	}

	return append([]byte(res), content...)
}

func get_all_titles(content string) []title {
	titles := regular_get(content)
	res := []title{}
	if len(titles) == 0 {
		return res
	}

	for _, t := range titles {
		level := 2
		i := 0
		for ; t[i] != byte(' '); i++ {
			level++
		}

		res = append(res, title{level, t[i+1:]})
	}

	return res

}

func regular_get(content string) []string {
	re := regexp.MustCompile(`##(.*)`)

	matches := re.FindAllStringSubmatch(content, -1)

	res := []string{}

	for _, match := range matches {
		if len(match) >= 2 {
			res = append(res, match[1])
		}
	}

	return res
}
