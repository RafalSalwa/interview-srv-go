// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: auth_service.proto

package intrvproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xdf, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e,
	0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x74,
	0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x69,
	0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x1e, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69,
	0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x6c, 0x53, 0x61, 0x6c, 0x77, 0x61, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x61, 0x70, 0x70, 0x2d, 0x73, 0x72, 0x76, 0x2f,
	0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_auth_service_proto_goTypes = []interface{}{
	(*SignUpUserInput)(nil),          // 0: intrvproto.SignUpUserInput
	(*SignInUserInput)(nil),          // 1: intrvproto.SignInUserInput
	(*SignInByCodeUserInput)(nil),    // 2: intrvproto.SignInByCodeUserInput
	(*VerificationCodeRequest)(nil),  // 3: intrvproto.VerificationCodeRequest
	(*SignUpUserResponse)(nil),       // 4: intrvproto.SignUpUserResponse
	(*SignInUserResponse)(nil),       // 5: intrvproto.SignInUserResponse
	(*VerificationCodeResponse)(nil), // 6: intrvproto.VerificationCodeResponse
}
var file_auth_service_proto_depIdxs = []int32{
	0, // 0: intrvproto.AuthService.SignUpUser:input_type -> intrvproto.SignUpUserInput
	1, // 1: intrvproto.AuthService.SignInUser:input_type -> intrvproto.SignInUserInput
	2, // 2: intrvproto.AuthService.SignInByCode:input_type -> intrvproto.SignInByCodeUserInput
	3, // 3: intrvproto.AuthService.GetVerificationKey:input_type -> intrvproto.VerificationCodeRequest
	4, // 4: intrvproto.AuthService.SignUpUser:output_type -> intrvproto.SignUpUserResponse
	5, // 5: intrvproto.AuthService.SignInUser:output_type -> intrvproto.SignInUserResponse
	5, // 6: intrvproto.AuthService.SignInByCode:output_type -> intrvproto.SignInUserResponse
	6, // 7: intrvproto.AuthService.GetVerificationKey:output_type -> intrvproto.VerificationCodeResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_service_proto_init() }
func file_auth_service_proto_init() {
	if File_auth_service_proto != nil {
		return
	}
	file_rpc_signin_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_proto_depIdxs,
	}.Build()
	File_auth_service_proto = out.File
	file_auth_service_proto_rawDesc = nil
	file_auth_service_proto_goTypes = nil
	file_auth_service_proto_depIdxs = nil
}
