package x

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCommand(t *testing.T) {
	t.Run("creates root x command", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "x", cmd.Use)
		assert.Contains(t, cmd.Aliases, "experimental")
		assert.Contains(t, cmd.Short, "Experimental")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Equal(t, []string{"experimental"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.Contains(t, cmd.Long, "experimental")
	})

	t.Run("adds doc subcommand", func(t *testing.T) {
		cmd := NewRootCommand()
		assert.True(t, cmd.HasSubCommands())

		// Find doc subcommand
		docCmd, _, err := cmd.Find([]string{"doc"})
		assert.NoError(t, err)
		assert.NotNil(t, docCmd)
		assert.Equal(t, "doc", docCmd.Use)
	})

	t.Run("adds json2yaml subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find json2yaml subcommand
		j2yCmd, _, err := cmd.Find([]string{"json2yaml"})
		assert.NoError(t, err)
		assert.NotNil(t, j2yCmd)
		assert.Equal(t, "json2yaml", j2yCmd.Use)
	})

	t.Run("adds yaml2json subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find yaml2json subcommand
		y2jCmd, _, err := cmd.Find([]string{"yaml2json"})
		assert.NoError(t, err)
		assert.NotNil(t, y2jCmd)
		assert.Contains(t, y2jCmd.Use, "yaml2json")
	})

	t.Run("adds yaml2idl subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find yaml2idl subcommand
		y2iCmd, _, err := cmd.Find([]string{"yaml2idl"})
		assert.NoError(t, err)
		assert.NotNil(t, y2iCmd)
		assert.Contains(t, y2iCmd.Use, "yaml2idl")
	})

	t.Run("adds idl2yaml subcommand", func(t *testing.T) {
		cmd := NewRootCommand()

		// Find idl2yaml subcommand
		i2yCmd, _, err := cmd.Find([]string{"idl2yaml"})
		assert.NoError(t, err)
		assert.NotNil(t, i2yCmd)
		assert.Contains(t, i2yCmd.Use, "idl2yaml")
	})

	t.Run("has all five subcommands", func(t *testing.T) {
		cmd := NewRootCommand()
		subcommands := cmd.Commands()

		assert.Len(t, subcommands, 5)

		// Check that we have all expected subcommands
		subcommandNames := make([]string, 0, len(subcommands))
		for _, subcmd := range subcommands {
			subcommandNames = append(subcommandNames, subcmd.Use)
		}

		// Check each subcommand exists (may have args in Use field)
		hasDoc := false
		hasJson2Yaml := false
		hasYaml2Json := false
		hasYaml2Idl := false
		hasIdl2Yaml := false

		for _, use := range subcommandNames {
			if use == "doc" {
				hasDoc = true
			}
			if use == "json2yaml" {
				hasJson2Yaml = true
			}
			if len(use) >= 10 && use[:10] == "yaml2json " {
				hasYaml2Json = true
			} else if use == "yaml2json" {
				hasYaml2Json = true
			}
			if len(use) >= 9 && use[:9] == "yaml2idl " {
				hasYaml2Idl = true
			} else if use == "yaml2idl" {
				hasYaml2Idl = true
			}
			if len(use) >= 9 && use[:9] == "idl2yaml " {
				hasIdl2Yaml = true
			} else if use == "idl2yaml" {
				hasIdl2Yaml = true
			}
		}

		assert.True(t, hasDoc, "doc subcommand not found")
		assert.True(t, hasJson2Yaml, "json2yaml subcommand not found")
		assert.True(t, hasYaml2Json, "yaml2json subcommand not found")
		assert.True(t, hasYaml2Idl, "yaml2idl subcommand not found")
		assert.True(t, hasIdl2Yaml, "idl2yaml subcommand not found")
	})

	t.Run("json2yaml has j2y alias", func(t *testing.T) {
		cmd := NewRootCommand()
		j2yCmd, _, err := cmd.Find([]string{"j2y"})
		assert.NoError(t, err)
		assert.NotNil(t, j2yCmd)
		assert.Equal(t, "json2yaml", j2yCmd.Use)
	})

	t.Run("yaml2json has y2j alias", func(t *testing.T) {
		cmd := NewRootCommand()
		y2jCmd, _, err := cmd.Find([]string{"y2j"})
		assert.NoError(t, err)
		assert.NotNil(t, y2jCmd)
		assert.Contains(t, y2jCmd.Use, "yaml2json")
	})
}

func TestNewDocsCommand(t *testing.T) {
	t.Run("creates doc command", func(t *testing.T) {
		cmd := NewDocsCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "doc", cmd.Use)
		assert.Contains(t, cmd.Short, "docs")
		assert.Contains(t, cmd.Short, "markdown")
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewDocsCommand()
		assert.Contains(t, cmd.Long, "markdown")
	})

	t.Run("has force flag", func(t *testing.T) {
		cmd := NewDocsCommand()
		flag := cmd.Flags().Lookup("force")
		assert.NotNil(t, flag)
		assert.Equal(t, "f", flag.Shorthand)
		assert.Equal(t, "false", flag.DefValue)
	})

	t.Run("force flag defaults to false", func(t *testing.T) {
		cmd := NewDocsCommand()
		force, err := cmd.Flags().GetBool("force")
		assert.NoError(t, err)
		assert.False(t, force)
	})

	t.Run("accepts force flag", func(t *testing.T) {
		cmd := NewDocsCommand()
		err := cmd.ParseFlags([]string{"--force"})
		assert.NoError(t, err)

		force, err := cmd.Flags().GetBool("force")
		assert.NoError(t, err)
		assert.True(t, force)
	})

	t.Run("accepts short force flag", func(t *testing.T) {
		cmd := NewDocsCommand()
		err := cmd.ParseFlags([]string{"-f"})
		assert.NoError(t, err)

		force, err := cmd.Flags().GetBool("force")
		assert.NoError(t, err)
		assert.True(t, force)
	})

	t.Run("has Run function", func(t *testing.T) {
		cmd := NewDocsCommand()
		assert.NotNil(t, cmd.Run)
	})

	t.Run("accepts maximum 1 argument", func(t *testing.T) {
		cmd := NewDocsCommand()
		assert.NotNil(t, cmd.Args)

		// Test with no arguments (should pass)
		err := cmd.Args(cmd, []string{})
		assert.NoError(t, err)

		// Test with one argument (should pass)
		err = cmd.Args(cmd, []string{"docs"})
		assert.NoError(t, err)

		// Test with two arguments (should fail)
		err = cmd.Args(cmd, []string{"docs", "extra"})
		assert.Error(t, err)
	})
}

func TestNewJson2YamlCommand(t *testing.T) {
	t.Run("creates json2yaml command", func(t *testing.T) {
		cmd := NewJson2YamlCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "json2yaml", cmd.Use)
		assert.Contains(t, cmd.Aliases, "j2y")
		assert.Contains(t, cmd.Short, "json")
		assert.Contains(t, cmd.Short, "yaml")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewJson2YamlCommand()
		assert.Equal(t, []string{"j2y"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewJson2YamlCommand()
		assert.Contains(t, cmd.Long, "json")
		assert.Contains(t, cmd.Long, "yaml")
	})

	t.Run("requires exactly one argument", func(t *testing.T) {
		cmd := NewJson2YamlCommand()
		assert.NotNil(t, cmd.Args)

		// Test with no arguments
		err := cmd.Args(cmd, []string{})
		assert.Error(t, err)

		// Test with one argument (should pass)
		err = cmd.Args(cmd, []string{"file.json"})
		assert.NoError(t, err)

		// Test with two arguments
		err = cmd.Args(cmd, []string{"file1.json", "file2.json"})
		assert.Error(t, err)
	})

	t.Run("has Run function", func(t *testing.T) {
		cmd := NewJson2YamlCommand()
		assert.NotNil(t, cmd.Run)
	})
}

func TestNewYaml2JsonCommand(t *testing.T) {
	t.Run("creates yaml2json command", func(t *testing.T) {
		cmd := NewYaml2JsonCommand()
		assert.NotNil(t, cmd)
		assert.Contains(t, cmd.Use, "yaml2json")
		assert.Contains(t, cmd.Aliases, "y2j")
		assert.Contains(t, cmd.Short, "yaml")
		assert.Contains(t, cmd.Short, "json")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewYaml2JsonCommand()
		assert.Equal(t, []string{"y2j"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewYaml2JsonCommand()
		assert.Contains(t, cmd.Long, "yaml")
		assert.Contains(t, cmd.Long, "json")
	})

	t.Run("requires exactly one argument", func(t *testing.T) {
		cmd := NewYaml2JsonCommand()
		assert.NotNil(t, cmd.Args)

		// Test with no arguments
		err := cmd.Args(cmd, []string{})
		assert.Error(t, err)

		// Test with one argument (should pass)
		err = cmd.Args(cmd, []string{"file.yaml"})
		assert.NoError(t, err)

		// Test with two arguments
		err = cmd.Args(cmd, []string{"file1.yaml", "file2.yaml"})
		assert.Error(t, err)
	})

	t.Run("has Run function", func(t *testing.T) {
		cmd := NewYaml2JsonCommand()
		assert.NotNil(t, cmd.Run)
	})
}
