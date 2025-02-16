package constants

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

// constants defines some file paths that are shared outside of the
// MCO package; and thus consumed by other users

const (
	// APIServerURLFile is the path to the apiserver url environment file.
	// See templates/master/00-master/_base/files/apiserver-url-env.yaml
	APIServerURLFile = "/etc/kubernetes/apiserver-url.env"
)

var (
	// NodeUpdateBackoff is the backoff time before asking APIServer to update node
	// object again.
	NodeUpdateBackoff = wait.Backoff{
		Steps:    5,
		Duration: 100 * time.Millisecond,
		Jitter:   1.0,
	}
	RenderBackoff = wait.Backoff{
		Steps:    5,
		Duration: time.Second,
		Jitter:   1.0,
	}
	// NodeUpdateInProgressTaint is a taint applied by MCC when the update of node starts.
	NodeUpdateInProgressTaint = &corev1.Taint{
		Key:    "UpdateInProgress",
		Effect: corev1.TaintEffectPreferNoSchedule,
	}
	// ConstantsByName is a map of constants for ease of templating
	ConstantsByName = map[string]string{
		"APIServerURLFile": APIServerURLFile,
	}
)
