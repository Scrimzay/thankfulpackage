package thankfulpackage

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Scans Go files in the given dir for github.com imports
// and creates a README.md thanking those packages
func GenerateThanks(dir string) error {
	imports, err := collectGithubImports(dir)
	if err != nil {
		return fmt.Errorf("failed to collect imports: %v", err)
	}
	if len(imports) == 0 {
		return nil
	}

	readmeContent := generateReadmeContent(imports)
	err = os.WriteFile(filepath.Join(dir, "README.md"), []byte(readmeContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write README.md: %v", err)
	}

	return nil
}

// walks the directory and collects unique github.com imports
func collectGithubImports(dir string) (map[string]struct{}, error) {
	imports := make(map[string]struct{})
	fset := token.NewFileSet()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %v", path, err)
		}

		for _, imp := range file.Imports {
			// Remove quotes from import path
			path := strings.Trim(imp.Path.Value, `"`)
			if strings.HasPrefix(path, "github.com") {
				imports[path] = struct{}{}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return imports, nil
}

func generateReadmeContent(imports map[string]struct{}) string {
	var buf bytes.Buffer
	buf.WriteString("# Thanks to Open Source\n\n")
	buf.WriteString("This project uses the following awesome open-source packages from GitHub:\n\n")

	for imp := range imports {
		// Format as a Markdown link
		repoURL := fmt.Sprintf("https://%s", imp)
		buf.WriteString(fmt.Sprintf("- [%s](%s)\n", imp, repoURL))
	}

	buf.WriteString("\nThank you to all the maintainers and contributors of these packages!\n")
	return buf.String()
}