package azure

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"strings"
)

const (
	autoDiscovererTypeLabel        = "label"
)

// A labelAutoDiscoveryConfig specifies how to auto-discover Azure scale sets.
type labelAutoDiscoveryConfig struct {
	// Key-values to match on.
	Selector map[string]string
}

// ParseLabelAutoDiscoverySpecs returns any provided NodeGroupAutoDiscoverySpecs
// parsed into configuration appropriate for ASG autodiscovery.
func ParseLabelAutoDiscoverySpecs(o cloudprovider.NodeGroupDiscoveryOptions) ([]labelAutoDiscoveryConfig, error) {
	cfgs := make([]labelAutoDiscoveryConfig, len(o.NodeGroupAutoDiscoverySpecs))
	var err error
	for i, spec := range o.NodeGroupAutoDiscoverySpecs {
		cfgs[i], err = parseLabelAutoDiscoverySpec(spec)
		if err != nil {
			return nil, err
		}
	}
	return cfgs, nil
}

// parseLabelAutoDiscoverySpec parses a single spec and returns the corresponding node group spec.
func parseLabelAutoDiscoverySpec(spec string) (labelAutoDiscoveryConfig, error) {
	cfg := labelAutoDiscoveryConfig{
		Selector: make(map[string]string),
	}

	tokens := strings.Split(spec, ":")
	if len(tokens) != 2 {
		return cfg, fmt.Errorf("spec \"%s\" should be discoverer:key=value,key=value", spec)
	}
	discoverer := tokens[0]
	if discoverer != autoDiscovererTypeLabel {
		return cfg, fmt.Errorf("unsupported discoverer specified: %s", discoverer)
	}

	for _, arg := range strings.Split(tokens[1], ",") {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			return cfg, fmt.Errorf("invalid key=value pair %s", kv)
		}
		k, v := kv[0], kv[1]
		if k == "" || v == "" {
			return cfg, fmt.Errorf("empty value not allowed in key=value tag pairs")
		}
		cfg.Selector[k] = v
	}
	return cfg, nil
}

func matchDiscoveryConfig(labels map[string]*string, configs []labelAutoDiscoveryConfig) bool {
	if len(configs) == 0 {
		return false
	}

	for _, c := range configs {
		if len(c.Selector) == 0 {
			return false
		}

		for k, v := range c.Selector {
			value, ok := labels[k]
			if !ok {
				return false
			}

			if len(v) > 0 {
				if value == nil || *value != v {
					return false
				}
			}
		}
	}

	return true
}
