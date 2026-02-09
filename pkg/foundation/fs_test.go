package foundation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJoin(t *testing.T) {
	t.Run("join relative paths", func(t *testing.T) {
		result := Join("a", "b", "c")
		expected := filepath.Join("a", "b", "c")
		assert.Equal(t, expected, result)
	})

	t.Run("last element is absolute", func(t *testing.T) {
		absPath := "/absolute/path"
		result := Join("a", "b", absPath)
		assert.Equal(t, absPath, result)
	})

	t.Run("single element", func(t *testing.T) {
		result := Join("single")
		assert.Equal(t, "single", result)
	})

	t.Run("empty elements", func(t *testing.T) {
		result := Join("", "a", "b")
		expected := filepath.Join("", "a", "b")
		assert.Equal(t, expected, result)
	})
}

func TestBaseName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"file with extension", "/path/to/file.txt", "file.txt"},
		{"directory", "/path/to/dir/", "dir"},
		{"single file", "file.txt", "file.txt"},
		{"no extension", "/path/to/file", "file"},
		{"root", "/", "/"},
		{"current dir", ".", "."},
		{"parent dir", "..", ".."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BaseName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsDir(t *testing.T) {
	dir := t.TempDir()

	t.Run("existing directory", func(t *testing.T) {
		assert.True(t, IsDir(dir))
	})

	t.Run("non-existing path", func(t *testing.T) {
		assert.False(t, IsDir(filepath.Join(dir, "nonexistent")))
	})

	t.Run("file is not directory", func(t *testing.T) {
		file := filepath.Join(dir, "file.txt")
		err := os.WriteFile(file, []byte("content"), 0644)
		require.NoError(t, err)
		assert.False(t, IsDir(file))
	})
}

func TestIsFile(t *testing.T) {
	dir := t.TempDir()

	t.Run("existing file", func(t *testing.T) {
		file := filepath.Join(dir, "test.txt")
		err := os.WriteFile(file, []byte("content"), 0644)
		require.NoError(t, err)
		assert.True(t, IsFile(file))
	})

	t.Run("directory is not file", func(t *testing.T) {
		assert.False(t, IsFile(dir))
	})

	t.Run("non-existing path", func(t *testing.T) {
		assert.False(t, IsFile(filepath.Join(dir, "nonexistent.txt")))
	})
}

func TestIsExist(t *testing.T) {
	dir := t.TempDir()

	t.Run("existing directory", func(t *testing.T) {
		assert.True(t, IsExist(dir))
	})

	t.Run("existing file", func(t *testing.T) {
		file := filepath.Join(dir, "test.txt")
		err := os.WriteFile(file, []byte("content"), 0644)
		require.NoError(t, err)
		assert.True(t, IsExist(file))
	})

	t.Run("non-existing path", func(t *testing.T) {
		assert.False(t, IsExist(filepath.Join(dir, "nonexistent")))
	})
}

func TestReadWriteDocument(t *testing.T) {
	dir := t.TempDir()

	t.Run("JSON document", func(t *testing.T) {
		type TestData struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}

		data := TestData{Name: "test", Value: 42}
		path := filepath.Join(dir, "test.json")

		// Write
		err := WriteDocument(path, data)
		require.NoError(t, err)

		// Read
		var result TestData
		err = ReadDocument(path, &result)
		require.NoError(t, err)

		assert.Equal(t, data.Name, result.Name)
		assert.Equal(t, data.Value, result.Value)
	})

	t.Run("YAML document", func(t *testing.T) {
		type TestData struct {
			Name  string `yaml:"name"`
			Value int    `yaml:"value"`
		}

		data := TestData{Name: "test", Value: 42}
		path := filepath.Join(dir, "test.yaml")

		// Write
		err := WriteDocument(path, data)
		require.NoError(t, err)

		// Read
		var result TestData
		err = ReadDocument(path, &result)
		require.NoError(t, err)

		assert.Equal(t, data.Name, result.Name)
		assert.Equal(t, data.Value, result.Value)
	})

	t.Run("YML extension", func(t *testing.T) {
		type TestData struct {
			Name string `yaml:"name"`
		}

		data := TestData{Name: "test"}
		path := filepath.Join(dir, "test.yml")

		err := WriteDocument(path, data)
		require.NoError(t, err)

		var result TestData
		err = ReadDocument(path, &result)
		require.NoError(t, err)

		assert.Equal(t, data.Name, result.Name)
	})

	t.Run("unsupported extension write", func(t *testing.T) {
		path := filepath.Join(dir, "test.txt")
		err := WriteDocument(path, "data")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported file extension")
	})

	t.Run("unsupported extension read", func(t *testing.T) {
		path := filepath.Join(dir, "test.txt")
		err := os.WriteFile(path, []byte("content"), 0644)
		require.NoError(t, err)

		var result string
		err = ReadDocument(path, &result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported file extension")
	})

	t.Run("read non-existing file", func(t *testing.T) {
		path := filepath.Join(dir, "nonexistent.json")
		var result map[string]interface{}
		err := ReadDocument(path, &result)
		assert.Error(t, err)
	})
}

func TestIsDocument(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"JSON file", "test.json", true},
		{"YAML file", "test.yaml", true},
		{"YML file", "test.yml", true},
		{"text file", "test.txt", false},
		{"no extension", "test", false},
		{"Go file", "test.go", false},
		{"hidden JSON", ".config.json", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsDocument(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFindDocuments(t *testing.T) {
	dir := t.TempDir()

	t.Run("find documents in directory", func(t *testing.T) {
		// Create test files
		files := []string{
			"doc1.json",
			"doc2.yaml",
			"doc3.yml",
			"notdoc.txt",
			"script.go",
		}

		for _, file := range files {
			path := filepath.Join(dir, file)
			err := os.WriteFile(path, []byte("content"), 0644)
			require.NoError(t, err)
		}

		// Create a subdirectory (should be ignored)
		subdir := filepath.Join(dir, "subdir")
		err := os.Mkdir(subdir, 0755)
		require.NoError(t, err)

		// Find documents
		docs, err := FindDocuments(dir)
		require.NoError(t, err)

		assert.Len(t, docs, 3)
		// Check that only document files are returned
		for _, doc := range docs {
			assert.True(t, IsDocument(doc))
		}
	})

	t.Run("empty directory", func(t *testing.T) {
		emptyDir := filepath.Join(dir, "empty")
		err := os.Mkdir(emptyDir, 0755)
		require.NoError(t, err)

		docs, err := FindDocuments(emptyDir)
		require.NoError(t, err)
		assert.Empty(t, docs)
	})

	t.Run("non-existing directory", func(t *testing.T) {
		docs, err := FindDocuments(filepath.Join(dir, "nonexistent"))
		assert.Error(t, err)
		assert.Empty(t, docs)
	})
}

func TestMakeDirRemoveDir(t *testing.T) {
	dir := t.TempDir()

	t.Run("create and remove directory", func(t *testing.T) {
		newDir := filepath.Join(dir, "testdir")

		// Create
		err := MakeDir(newDir)
		require.NoError(t, err)
		assert.True(t, IsDir(newDir))

		// Remove
		err = RemoveDir(newDir)
		require.NoError(t, err)
		assert.False(t, IsExist(newDir))
	})

	t.Run("create nested directories", func(t *testing.T) {
		nestedDir := filepath.Join(dir, "a", "b", "c")
		err := MakeDir(nestedDir)
		require.NoError(t, err)
		assert.True(t, IsDir(nestedDir))
	})

	t.Run("remove directory with contents", func(t *testing.T) {
		testDir := filepath.Join(dir, "withcontent")
		err := MakeDir(testDir)
		require.NoError(t, err)

		// Add files
		file := filepath.Join(testDir, "file.txt")
		err = os.WriteFile(file, []byte("content"), 0644)
		require.NoError(t, err)

		// Remove should remove all contents
		err = RemoveDir(testDir)
		require.NoError(t, err)
		assert.False(t, IsExist(testDir))
	})
}

func TestDir(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"file path", "/path/to/file.txt", "/path/to"},
		{"directory", "/path/to/dir/", "/path/to/dir"},
		{"no directory", "file.txt", "."},
		{"root file", "/file.txt", "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Dir(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWriteFile(t *testing.T) {
	dir := t.TempDir()

	t.Run("write file", func(t *testing.T) {
		path := filepath.Join(dir, "test.txt")
		data := []byte("test content")

		err := WriteFile(path, data)
		require.NoError(t, err)

		// Verify
		content, err := os.ReadFile(path)
		require.NoError(t, err)
		assert.Equal(t, data, content)
	})

	t.Run("overwrite existing file", func(t *testing.T) {
		path := filepath.Join(dir, "overwrite.txt")

		// First write
		err := WriteFile(path, []byte("first"))
		require.NoError(t, err)

		// Second write
		err = WriteFile(path, []byte("second"))
		require.NoError(t, err)

		// Verify
		content, err := os.ReadFile(path)
		require.NoError(t, err)
		assert.Equal(t, []byte("second"), content)
	})
}

func TestHasExt(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		exts     []string
		expected bool
	}{
		{"single match", "file.txt", []string{".txt"}, true},
		{"multiple match first", "file.go", []string{".go", ".txt"}, true},
		{"multiple match second", "file.txt", []string{".go", ".txt"}, true},
		{"no match", "file.md", []string{".txt", ".go"}, false},
		{"no extension", "file", []string{".txt"}, false},
		{"empty extensions", "file.txt", []string{}, false},
		{"case sensitive", "file.TXT", []string{".txt"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasExt(tt.file, tt.exts...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExt(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected string
	}{
		{"text file", "file.txt", ".txt"},
		{"go file", "file.go", ".go"},
		{"no extension", "file", ""},
		{"multiple dots", "file.tar.gz", ".gz"},
		{"hidden file", ".gitignore", ".gitignore"},
		{"hidden with ext", ".config.json", ".json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ext(tt.file)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFallbackDir(t *testing.T) {
	dir := t.TempDir()

	t.Run("first directory exists", func(t *testing.T) {
		dir1 := filepath.Join(dir, "first")
		target := filepath.Join(dir1, "target")
		err := os.MkdirAll(target, 0755)
		require.NoError(t, err)

		result, err := FallbackDir("target", dir1, "/nonexistent")
		require.NoError(t, err)
		assert.Equal(t, target, result)
	})

	t.Run("second directory exists", func(t *testing.T) {
		dir1 := filepath.Join(dir, "first2")
		dir2 := filepath.Join(dir, "second2")
		target := filepath.Join(dir2, "target")

		err := os.MkdirAll(dir1, 0755)
		require.NoError(t, err)
		err = os.MkdirAll(target, 0755)
		require.NoError(t, err)

		result, err := FallbackDir("target", dir1, dir2)
		require.NoError(t, err)
		assert.Equal(t, target, result)
	})

	t.Run("no directory exists", func(t *testing.T) {
		result, err := FallbackDir("target", "/nonexistent1", "/nonexistent2")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
		assert.Empty(t, result)
	})
}

func TestScanFile(t *testing.T) {
	dir := t.TempDir()

	t.Run("scan file with multiple lines", func(t *testing.T) {
		content := "line1\nline2\nline3\n"
		path := filepath.Join(dir, "scan.txt")
		err := os.WriteFile(path, []byte(content), 0644)
		require.NoError(t, err)

		lines, err := ScanFile(path)
		require.NoError(t, err)
		assert.Len(t, lines, 3)
		assert.Equal(t, []byte("line1"), lines[0])
		assert.Equal(t, []byte("line2"), lines[1])
		assert.Equal(t, []byte("line3"), lines[2])
	})

	t.Run("scan empty file", func(t *testing.T) {
		path := filepath.Join(dir, "empty.txt")
		err := os.WriteFile(path, []byte(""), 0644)
		require.NoError(t, err)

		lines, err := ScanFile(path)
		require.NoError(t, err)
		assert.Empty(t, lines)
	})

	t.Run("scan non-existing file", func(t *testing.T) {
		lines, err := ScanFile(filepath.Join(dir, "nonexistent.txt"))
		assert.Error(t, err)
		assert.Nil(t, lines)
	})
}

func TestYamlToJson(t *testing.T) {
	t.Run("convert yaml to json", func(t *testing.T) {
		yaml := []byte(`
name: test
value: 42
enabled: true
`)
		json, err := YamlToJson(yaml)
		require.NoError(t, err)
		assert.NotEmpty(t, json)
		assert.Contains(t, string(json), "test")
		assert.Contains(t, string(json), "42")
	})

	t.Run("invalid yaml", func(t *testing.T) {
		yaml := []byte("invalid: yaml: data:")
		_, err := YamlToJson(yaml)
		assert.Error(t, err)
	})
}

func TestReadYamlFromString(t *testing.T) {
	t.Run("read valid yaml", func(t *testing.T) {
		type Config struct {
			Name  string `yaml:"name"`
			Value int    `yaml:"value"`
		}

		yamlStr := `name: test
value: 42`

		var config Config
		err := ReadYamlFromString(yamlStr, &config)
		require.NoError(t, err)
		assert.Equal(t, "test", config.Name)
		assert.Equal(t, 42, config.Value)
	})
}

func TestReadYamlFromData(t *testing.T) {
	t.Run("read valid yaml", func(t *testing.T) {
		type Config struct {
			Name string `yaml:"name"`
		}

		yamlData := []byte("name: test")

		var config Config
		err := ReadYamlFromData(yamlData, &config)
		require.NoError(t, err)
		assert.Equal(t, "test", config.Name)
	})
}
