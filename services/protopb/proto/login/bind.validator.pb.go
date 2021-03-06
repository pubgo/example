// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/login/bind.proto

// 绑定手机号

package login

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CheckRequest) Validate() error {
	return nil
}
func (this *CheckResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *BindVerifyRequest) Validate() error {
	return nil
}
func (this *BindVerifyResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *BindData) Validate() error {
	return nil
}
func (this *BindChangeRequest) Validate() error {
	return nil
}
func (this *BindChangeResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AutomaticBindRequest) Validate() error {
	return nil
}
func (this *AutomaticBindResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *BindPhoneParseRequest) Validate() error {
	return nil
}
func (this *BindPhoneParseResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *BindPhoneParseByOneClickRequest) Validate() error {
	return nil
}
func (this *BindPhoneParseByOneClickResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
