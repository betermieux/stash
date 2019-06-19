// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "kmodules.xyz/openshift/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DeploymentConfigs returns a DeploymentConfigInformer.
	DeploymentConfigs() DeploymentConfigInformer
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

// DeploymentConfigs returns a DeploymentConfigInformer.
func (v *version) DeploymentConfigs() DeploymentConfigInformer {
	return &deploymentConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}