// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kyma-project.io/compass-runtime-agent/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CompassConnections returns a CompassConnectionInformer.
	CompassConnections() CompassConnectionInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CompassConnections returns a CompassConnectionInformer.
func (v *version) CompassConnections() CompassConnectionInformer {
	return &compassConnectionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
