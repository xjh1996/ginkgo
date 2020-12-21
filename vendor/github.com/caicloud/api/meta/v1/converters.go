package v1

import (
	"encoding/json"

	kubemeta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConvertKubeObjectMetaToObjectMeta converts Kubernetes metadata to an API ObjectMeta object.
func ConvertKubeObjectMetaToObjectMeta(typeMeta *kubemeta.TypeMeta, objectMeta *kubemeta.ObjectMeta) ObjectMeta {
	var ret ObjectMeta
	if typeMeta != nil {
		ret.Kind = typeMeta.Kind
		ret.APIVersion = typeMeta.APIVersion
	}
	if objectMeta != nil {
		ret.UID = string(objectMeta.UID)
		ret.CreationTimestamp = objectMeta.CreationTimestamp.Time
		ret.Name = objectMeta.Name
		ret.Namespace = objectMeta.Namespace
		ret.ClusterName = objectMeta.ClusterName
		if objectMeta.DeletionTimestamp != nil {
			ret.DeletionTimestamp = &objectMeta.DeletionTimestamp.Time
		}
		if objectMeta.Labels != nil {
			labels, _ := json.Marshal(objectMeta.Labels)
			ret.Labels = string(labels)
		}
		if objectMeta.Annotations != nil {
			annotations, _ := json.Marshal(objectMeta.Annotations)
			ret.Annotations = string(annotations)
		}
	}
	return ret
}
