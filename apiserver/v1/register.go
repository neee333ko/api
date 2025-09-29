package v1

import "github.com/neee333ko/component-base/pkg/scheme"

const GroupName = "iam.api"

var schemeGroupVersion = scheme.GroupVersion{Group: GroupName, Version: "v1"}

func Resource(resource string) *scheme.GroupResource {
	return schemeGroupVersion.WithResource(resource).GroupResource()
}
