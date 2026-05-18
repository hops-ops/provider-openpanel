/*
Copyright 2024 The Crossplane Authors.
*/

// Package v1alpha1 is a placeholder for resources promoted to v1beta1 (or
// not yet generated). Upjet's codegen emits an import of this package
// even when no v1alpha1 resources exist; keeping an empty SchemeBuilder
// keeps the scheme-registration code in zz_register.go satisfied.
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

const (
	Group   = "openpanel.crossplane.io"
	Version = "v1alpha1"
)

var (
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}
	SchemeBuilder      = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)
