package activity

import "github.com/apigear-io/cli/pkg/log"

var C = make(chan *Action)

func SendAction(action *Action) {
	select {
	case C <- action:
		log.Infof("action %s sent", action)
	default:
		log.Infof("action %s dropped", action)
	}
}
