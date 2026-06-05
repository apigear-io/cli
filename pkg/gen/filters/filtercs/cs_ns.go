package filtercs

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// nsName builds the dotted, PascalCased C# namespace for a module
// (e.g. module "demo.x" -> "Demo.X").
func nsName(m *model.Module) (string, error) {
	if m == nil {
		return "", fmt.Errorf("invalid module")
	}
	parts := strings.Split(m.Name, ".")
	for i, p := range parts {
		parts[i] = common.CamelTitleCase(p)
	}
	return strings.Join(parts, "."), nil
}

// csNs returns the C# namespace name for a module.
func csNs(m *model.Module) (string, error) {
	return nsName(m)
}

// csNsOpen opens a block-scoped C# namespace. Block scope (rather than the
// file-scoped form) keeps the output compatible with older C# toolchains.
func csNsOpen(m *model.Module) (string, error) {
	ns, err := nsName(m)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("namespace %s\n{", ns), nil
}

// csNsClose closes a block opened with csNsOpen.
func csNsClose(m *model.Module) (string, error) {
	ns, err := nsName(m)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("} // namespace %s", ns), nil
}
