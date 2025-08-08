package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func jniJavaSignatureParams(nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("ueJniJavaParams called with nil nodes")
	}
	var text = ""
	for _, p := range nodes {
		r, err := jniSignatureType(p)
		if err != nil {
			return "xxx", err
		}
		text += r
	}
	return text, nil
}
