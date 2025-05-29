package filterue

import (
	"github.com/apigear-io/cli/pkg/gen/filters/common"
)

//TODO: add test including prefix for all filters

func makeString(project string, moduleName string, submodule string, delimeter string) string {
	return "com" + delimeter +
		common.CamelLowerCase(project) + delimeter +
		ueGetModuleName(moduleName, submodule)
}

func ueGetModuleName(moduleName string, submodule string) string {
	return common.CamelLowerCase(moduleName) + common.CamelLowerCase(submodule)
}

func ueJavaPckgName(project string, moduleName string, submodule string) string {
	return makeString(project, moduleName, submodule, ".")
}

func ueJavaPath(project string, moduleName string, submodule string) string {
	return makeString(project, moduleName, submodule, "/")
}

func javaOnSingalStyle(signalName string) string {
	return "on" + common.CamelTitleCase(signalName)
}

func javaOnPropertyChangedStyle(signalName string) string {
	return "on" + common.CamelTitleCase(signalName) + "Changed"
}

func jniNameSetProperty(propertyName string) string {
	return "nativeSet" + common.CamelTitleCase(propertyName)
}

func jniNameGetProperty(propertyName string) string {
	return "nativeGet" + common.CamelTitleCase(propertyName)
}

func jniNameOperation(methodName string) string {
	return "nativet" + common.CamelTitleCase(methodName)
}

func ueJniClassPathPrefix(project string, moduleName string, submodule string, className string) string {
	return "Java" + makeString(project, moduleName, submodule, "_")
}
