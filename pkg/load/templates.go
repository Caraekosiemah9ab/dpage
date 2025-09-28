package load

import (
	"fmt"
	"html/template"
	"path/filepath"

	"dpage/types"
	"dpage/utils"
)

func Templates() error {
	templateDir, err := utils.GetCurrentDir()
	if err != nil {
		return err
	}

	templatesPath := filepath.Join(templateDir, "templates", "*.html")

	matches, err := filepath.Glob(templatesPath)
	if err != nil {
		return err
	}
	if len(matches) == 0 {
		return fmt.Errorf("fail template: %s", templatesPath)
	}

	fmt.Printf("Find templates: %d\n", len(matches))
	for _, match := range matches {
		fmt.Printf("-> %s\n", match)
	}

	types.Tmpl = template.Must(template.ParseGlob(templatesPath))
	return nil
}
