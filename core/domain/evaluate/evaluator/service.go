package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/models"
	"strings"
)

type Service struct {
}

func (s Service) Evaluate(node models.JClassNode) {
	var methodNameArray [][]string
	for _, method := range node.Methods {
		methodNameArray = append(methodNameArray, SplitCamelcase(method.Name))
	}

	lifecycleMap := s.buildLifecycle(methodNameArray)
	if len(lifecycleMap) > 0 {
		for key, value := range lifecycleMap {
			fmt.Println(key, value)
		}
	}

	if s.hasSameBehavior() {

	}
	if s.hasModelLike() {

	}
	if s.hasSameReturnType() {

	}
}

func (s Service) buildLifecycle(methodNameArray [][]string) map[string][]string {
	var hadLifecycle = make(map[string][]string)
	var nameMap = make(map[string][]string)
	for _, nameArray := range methodNameArray {
		firstWord := nameArray[0]
		if !(firstWord == "set" || firstWord == "get") {
			nameMap[firstWord] = append(nameMap[firstWord], strings.Join(nameArray, ""))
		}
		if len(nameMap[firstWord]) > 1 {
			hadLifecycle[firstWord] = nameMap[firstWord]
		}
	}

	return hadLifecycle
}

func (s Service) hasSameBehavior() bool {
	return false
}

func (s Service) hasModelLike() bool {
	return false
}

func (s Service) hasSameReturnType() bool {
	return false
}