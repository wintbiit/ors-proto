package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
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
	convertProtoId(dir, dst)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	genFile, err := os.OpenFile(path.Join(dst, "proto.gen.go"),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}

	genFile.WriteString("package " + dst + "\n\n")
	defer genFile.Close()
	for _, file := range files {
		convertProto(file, dir, dst, genFile)
	}

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
func readProtoFile(name string) ([]Pair, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	result := make([]Pair, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		field := parts[0]
		typ := parts[1]
		trimmedTyp := strings.ReplaceAll(typ, "[]", "")
		mappedType, ok := typeMap[trimmedTyp]
		if ok {
			trimmedTyp = mappedType
		} else if utf8.RuneCountInString(trimmedTyp) == len(trimmedTyp) {
			trimmedTyp = strcase.ToCamel(trimmedTyp)
		}

		for i := 0; i < strings.Count(typ, "[]"); i++ {
			trimmedTyp = "[]" + trimmedTyp
		}

		// result[field] = trimmedTyp
		result = append(result, Pair{field, trimmedTyp})
	}

	return result, nil
}

// generate proto file
func generateProtoFile(dir, name, pkg string, size int, genFile *os.File, fields []Pair) error {
	name = strcase.ToCamel(name)

	genFile.WriteString(fmt.Sprintf("type %s struct {\n", name))

	for _, e := range fields {
		field := e.Key
		typ := e.Value

		if strings.HasPrefix(field, "_") {
			field = "S" + field[1:]
		}
		// if is Chinese, leave as is. if not, convert to camel case
		if utf8.RuneCountInString(field) == len(field) {
			field = strcase.ToCamel(field)
		}
		genFile.WriteString(fmt.Sprintf("\t%s %s\n", field, typ))
	}

	genFile.WriteString("}\n\n")

	fp := path.Join(dir, name+".go")

	if _, err := os.Stat(fp); err == nil {
		return nil
	}

	f, err := os.OpenFile(fp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("package %s\n\n", pkg))

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

func convertProto(file os.DirEntry, dir, dst string, genFile *os.File) {
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

	err = generateProtoFile(dst, name, dst, size, genFile, fields)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func convertProtoId(dir, dst string) {
	f, err := os.Open(path.Join(dir, "proto.csv"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	idMap := make(map[int]string)

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
		protoName := parts[0]
		if strings.HasPrefix(protoName, "ID_") {
			protoName = strings.Replace(protoName, "ID_", "", 1)
		}
		protoName = "ProtoID" + protoName

		idMap[id] = protoName
	}

	f, err = os.Open(path.Join(dir, "battle_proto.csv"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	scanner = bufio.NewScanner(f)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		protoName := parts[0]
		if strings.HasPrefix(protoName, "ID_") {
			protoName = strings.Replace(protoName, "ID_", "", 1)
		}
		protoName = "ProtoID" + protoName

		idMap[id] = protoName
	}

	fp := path.Join(dst, "proto_id"+".gen.go")
	f, err = os.OpenFile(fp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("package %s\n\n", dst))

	f.WriteString("const (\n")

	keys := make([]int, 0, len(idMap))
	for k := range idMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		f.WriteString(fmt.Sprintf("\t%s = %d\n", idMap[k], k))
	}

	f.WriteString(")\n\n")

	f.WriteString(fmt.Sprintf("var ProtoIdMap = map[int]string{\n"))

	for _, k := range keys {
		f.WriteString(fmt.Sprintf("\t%d: \"%s\",\n", k, idMap[k]))
	}

	f.WriteString("}\n")

	fmt.Println("Done")
}

type Pair struct {
	Key   string
	Value string
}
