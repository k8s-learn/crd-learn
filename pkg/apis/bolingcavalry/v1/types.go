package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Student struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec StudentSpec				`json:"spec"`
}

type StudentSpec struct {
	Name   string `json:"name"`
	School string `json:"school"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StudentList is a list of Student resources
type StudentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items []Student `json:"items"`
}

//cd $GOPATH/src/k8s.io/code-generator
//do not use absolute path
//./generate-groups.sh all $GOPATH/src/git.pacloud.io/crd-learn/pkg/client $GOPATH/src/git.pacloud.io/crd-learn/pkg/apis bolingcavalry:v1
//./generate-groups.sh all git.pacloud.io/crd-learn/pkg/client git.pacloud.io/crd-learn/pkg/apis bolingcavalry:v1