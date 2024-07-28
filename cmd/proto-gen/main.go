package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/iancoleman/strcase"
)

func main() {
	var dir string
	var dst string

	flag.StringVar(&dir, "dir", "", "Directory of proto files")
	flag.StringVar(&dst, "dst", "", "Destination directory of generated files")

	flag.Parse()

	if dir == "" || dst == "" {
		fmt.Println("Usage: proto-gen -dir <dir> -dst <dst>")
		return
	}

	var err error

	sizeMap, err = readSizeMap(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	// convertProtoId(dir, dst, "proto")
	// convertProtoId(dir, dst, "battle_proto")

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(files))
	for _, file := range files {
		go func(file os.DirEntry) {
			defer wg.Done()
			convertProto(file, dir, dst)
		}(file)
	}

	wg.Wait()

	fmt.Println("Done")
}

var typeMap = map[string]string{
	"Byte":   "byte",
	"Uint8":  "uint8",
	"UInt16": "uint16",
	"UInt32": "uint32",
	"UInt64": "uint64",
	"Int8":   "int8",
	"Int16":  "int16",
	"Int32":  "int32",
	"Int64":  "int64",
	"String": "string",
	"Single": "float32",
	"Double": "float64",
	"SByte":  "byte",
}

var sizeMap map[string]int

// read proto csv file
func readProtoFile(name string) (map[string]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	result := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		field := parts[0]
		typ := parts[1]
		trimmedTyp := strings.ReplaceAll(typ, "[]", "")
		mappedType, ok := typeMap[trimmedTyp]
		if ok {
			trimmedTyp = mappedType
		}

		for i := 0; i < strings.Count(typ, "[]"); i++ {
			trimmedTyp = "[]" + trimmedTyp
		}

		result[field] = trimmedTyp
	}

	return result, nil
}

// generate proto file
func generateProtoFile(dir, name, pkg string, size int, fields map[string]string) error {
	name = strcase.ToCamel(name)

	fp := path.Join(dir, name+".go")
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("package %s\n\n", pkg))

	f.WriteString(fmt.Sprintf("type %s struct {\n", name))

	for field, typ := range fields {
		// if is Chinese, leave as is. if not, convert to camel case
		if utf8.RuneCountInString(field) == len(field) {
			field = strcase.ToCamel(field)
		}
		f.WriteString(fmt.Sprintf("\t%s %s\n", field, typ))
	}

	f.WriteString("}\n\n")

	f.WriteString(fmt.Sprintf("const %sSize = %d\n\n", name, size))

	f.WriteString(fmt.Sprintf("func (s *%s) Serialize() []byte {\n", name))

	f.WriteString(fmt.Sprintf("\tbytes := make([]byte, %sSize)\n", name))

	f.WriteString("\t// TODO: Implement serialization\n")

	f.WriteString("\treturn bytes\n")

	f.WriteString("}\n\n")

	f.WriteString(fmt.Sprintf("func (s *%s) Deserialize(bytes []byte) error {\n", name))

	f.WriteString("\t// TODO: Implement deserialization\n")

	f.WriteString("\treturn nil\n")

	f.WriteString("}\n")

	return nil
}

func readSizeMap(dir string) (map[string]int, error) {
	sizeMap := make(map[string]int)
	f := path.Join(dir, "sizes.csv")

	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		sizeMap[parts[0]] = size
	}

	return sizeMap, nil
}

func convertProto(file os.DirEntry, dir, dst string) {
	if file.IsDir() {
		return
	}

	name := file.Name()
	if !strings.HasSuffix(name, ".csv") {
		return
	}

	if name == "sizes.csv" {
		return
	}

	if name == "proto.csv" {
		return
	}

	if name == "battle_proto.csv" {
		return
	}

	parts := strings.Split(name, ".")
	if len(parts) != 2 {
		fmt.Println("Invalid file name:", name)
		return
	}

	name = parts[0]

	fields, err := readProtoFile(path.Join(dir, file.Name()))
	if err != nil {
		fmt.Println(err)
		return
	}

	size, ok := sizeMap[name]
	if !ok {
		fmt.Println("Size not found for", name)
		size = 0
	}

	err = generateProtoFile(dst, name, dst, size, fields)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func convertProtoId(dir, dst, name string) {
	f, err := os.Open(path.Join(dir, name+".csv"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	idMap := make(map[string]int)

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		idMap[parts[0]] = id
	}

	fp := path.Join(dst, name+"_id"+".go")
	f, err = os.OpenFile(fp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("package %s\n\n", dst))

	f.WriteString("const (\n")

	for field, id := range idMap {
		f.WriteString(fmt.Sprintf("\t%s = %d\n", field, id))
	}

	f.WriteString(")\n")

	fmt.Println("Done")
}
