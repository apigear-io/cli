package filterue

import (
	"github.com/apigear-io/cli/pkg/gen/filters/common"
)

//TODO: add test including prefix for all filters

func makeString3(project string, moduleName string, submodule string, delimeter string) string {
	return common.CamelLowerCase(project) + delimeter + moduleName + delimeter + submodule
}

func makeString2(project string, moduleName string, delimeter string) string {
	return common.CamelLowerCase(project) + delimeter + moduleName
}

func ueJavaPckgName(project string, moduleName string, submodule string) string {
	if submodule != "" {
		return makeString3(project, moduleName, submodule, ".")
	}
	return makeString2(project, moduleName, ".")
}

func ueJavaPath(project string, moduleName string, submodule string) string {
	if submodule != "" {
		return makeString3(project, moduleName, submodule, "/")
	}
	return makeString2(project, moduleName, "/")
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
	return "native" + common.CamelTitleCase(methodName)
}

func ueJniClassPathPrefix(project string, moduleName string, submodule string, className string) string {
	if submodule != "" {
		return "Java_" + makeString3(project, moduleName, submodule, "_") + "_" + className
	}
	return "Java_" + makeString3(project, moduleName, className, "_")
}
