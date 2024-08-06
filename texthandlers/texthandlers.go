package texthandlers

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const readWriteMode = 666 // -rw-rw-rw-
// NOTE: if file exists, file permissions do not overwrites

type text struct {
	absolutePath string
	size         int
}

func RandomFromAlphabet(alphabet []byte, size int, absolutePath string) (text, error) {
	var builder strings.Builder
	for i := 0; i < size; i++ {
		randIndex := rand.Intn(len(alphabet))
		letter := alphabet[randIndex]
		builder.WriteByte(letter)
	}
	if err := os.WriteFile(
		absolutePath,
		[]byte(builder.String()),
		readWriteMode,
	); err != nil {
		return text{}, err
	}
	return text{absolutePath, size}, nil
}

type formulaCoords struct {
	left, right int
}

func findMath(data []byte) ([]formulaCoords, error) {
	var indexes []int
	isRightDoubleDollarIndex := false
	i := 0
	for i < len(data)-1 {
		current := data[i]
		lookahead := data[i+1]

		if current == '\\' && lookahead == '$' {
			i += 2
			continue
		} else if current == '$' && lookahead == '$' {
			if isRightDoubleDollarIndex {
				indexes = append(indexes, i+1)
			} else {
				indexes = append(indexes, i)
			}
			isRightDoubleDollarIndex = !isRightDoubleDollarIndex
			i += 2
			continue
		} else if current == '$' && lookahead != '$' {
			indexes = append(indexes, i)
			i += 2
			continue
		}

		i++
	}

	if len(indexes)%2 != 0 {
		return nil, fmt.Errorf("unclosed \"$\" or \"$$\"")
	}

	coords := make([]formulaCoords, len(indexes)/2)
	for i := 0; i < len(coords); i++ {
		coords[i] = formulaCoords{
			left:  indexes[2*i],
			right: indexes[2*i+1],
		}
	}
	return coords, nil
}

// fix dumb code (formula to hash-string conversion)
func getMathMapping(data []byte, formulasIndexes []formulaCoords) map[string]string {
	hasher := fnv.New64()
	formulasMapping := make(map[string]string)

	for _, formulaIndexes := range formulasIndexes {
		formula := data[formulaIndexes.left : formulaIndexes.right+1]

		hasher.Write(formula)
		hash := strconv.FormatUint(hasher.Sum64(), 10)

		lettersHash := make([]byte, len(hash))
		alphabet := "abcdefghijklmstupidity" // :)
		for i := 0; i < len(hash); i++ {
			// no error check since hash[i] is always just a digit
			res, _ := strconv.Atoi(string(hash[i])) //nolint:all
			lettersHash[i] = alphabet[res]
		}

		formulasMapping[string(formula)] = "hash" + string(lettersHash)

		hasher.Reset()
	}

	return formulasMapping
}

func FromArticle(arcticlePath, destPath string) (text, error) {
	if err := exec.Command(
		"pandoc",
		arcticlePath,
		"-o",
		destPath,
	).Run(); err != nil {
		return text{}, err
	}

	data, err := os.ReadFile(destPath)
	if err != nil {
		return text{}, err
	}

	formulasIndexes, err := findMath(data)
	if err != nil {
		return text{}, err
	}

	mapping := getMathMapping(data, formulasIndexes)
	var builder strings.Builder
	left := 0
	for _, formulaIndexes := range formulasIndexes {
		formula := data[formulaIndexes.left : formulaIndexes.right+1]
		builder.Write(data[left:formulaIndexes.left])
		builder.Write([]byte(mapping[string(formula)]))
		left = formulaIndexes.right + 1
	}
	builder.Write(data[left:])

	if err := os.Truncate(destPath, 0); err != nil {
		return text{}, err
	}

	if err := os.WriteFile(
		destPath,
		[]byte(builder.String()),
		readWriteMode,
	); err != nil {
		return text{}, err
	}

	return text{
		absolutePath: destPath,
		size:         len(data),
	}, nil
}

func (t text) getData() (string, error) {
	data, err := os.ReadFile(t.absolutePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (t text) toAscii() error {
	if err := exec.Command(
		"pandoc",
		t.absolutePath,
		"--ascii",
		"-o",
		t.absolutePath,
	).Run(); err != nil {
		return err
	}

	return nil
}
