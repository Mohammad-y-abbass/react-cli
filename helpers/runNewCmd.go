package helpers

import "github.com/Mohammad-y-abbass/react-cli/types"

func RunNewCmd(appName string, useTs bool, useTailwind bool, data types.TemplateData) {

	ExecViteCmd(appName, useTs)

	if useTailwind {
		ConfigTailwind(appName, data, useTs)
	}

	SetInitialCode(appName, useTs, data)

	InstallDeps(appName, useTailwind)
}
