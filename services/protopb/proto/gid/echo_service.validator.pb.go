// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/gid/echo_service.proto

package gid

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Embedded) Validate() error {
	return nil
}
func (this *SimpleMessage) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	if oneOfNester, ok := this.GetExt().(*SimpleMessage_No); ok {
		if oneOfNester.No != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.No); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("No", err)
			}
		}
	}
	return nil
}
func (this *DynamicMessage) Validate() error {
	if this.StructField != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StructField); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StructField", err)
		}
	}
	if this.ValueField != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ValueField); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ValueField", err)
		}
	}
	return nil
}
func (this *DynamicMessageUpdate) Validate() error {
	if this.Body != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Body); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Body", err)
		}
	}
	return nil
}
