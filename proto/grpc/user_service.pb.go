// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: user_service.proto

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

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe8,
	0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43,
	0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x17, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x15, 0x2e, 0x69, 0x6e, 0x74,
	0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x2e, 0x69, 0x6e,
	0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x72,
	0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e,
	0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x17, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x6c, 0x53, 0x61, 0x6c,
	0x77, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x61, 0x70, 0x70,
	0x2d, 0x73, 0x72, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_user_service_proto_goTypes = []interface{}{
	(*StringValue)(nil),            // 0: intrvproto.StringValue
	(*GetUserRequest)(nil),         // 1: intrvproto.GetUserRequest
	(*VerifyUserRequest)(nil),      // 2: intrvproto.VerifyUserRequest
	(*ChangePasswordRequest)(nil),  // 3: intrvproto.ChangePasswordRequest
	(*GetUserSignInRequest)(nil),   // 4: intrvproto.GetUserSignInRequest
	(*VerificationCode)(nil),       // 5: intrvproto.VerificationCode
	(*BoolValue)(nil),              // 6: intrvproto.BoolValue
	(*UserDetails)(nil),            // 7: intrvproto.UserDetails
	(*VerificationResponse)(nil),   // 8: intrvproto.VerificationResponse
	(*ChangePasswordResponse)(nil), // 9: intrvproto.ChangePasswordResponse
}
var file_user_service_proto_depIdxs = []int32{
	0, // 0: intrvproto.UserService.CheckUserExists:input_type -> intrvproto.StringValue
	1, // 1: intrvproto.UserService.GetUserById:input_type -> intrvproto.GetUserRequest
	1, // 2: intrvproto.UserService.GetUserDetails:input_type -> intrvproto.GetUserRequest
	2, // 3: intrvproto.UserService.VerifyUser:input_type -> intrvproto.VerifyUserRequest
	3, // 4: intrvproto.UserService.ChangePassword:input_type -> intrvproto.ChangePasswordRequest
	4, // 5: intrvproto.UserService.GetUser:input_type -> intrvproto.GetUserSignInRequest
	5, // 6: intrvproto.UserService.GetUserByCode:input_type -> intrvproto.VerificationCode
	1, // 7: intrvproto.UserService.GetUserByToken:input_type -> intrvproto.GetUserRequest
	6, // 8: intrvproto.UserService.CheckUserExists:output_type -> intrvproto.BoolValue
	7, // 9: intrvproto.UserService.GetUserById:output_type -> intrvproto.UserDetails
	7, // 10: intrvproto.UserService.GetUserDetails:output_type -> intrvproto.UserDetails
	8, // 11: intrvproto.UserService.VerifyUser:output_type -> intrvproto.VerificationResponse
	9, // 12: intrvproto.UserService.ChangePassword:output_type -> intrvproto.ChangePasswordResponse
	7, // 13: intrvproto.UserService.GetUser:output_type -> intrvproto.UserDetails
	7, // 14: intrvproto.UserService.GetUserByCode:output_type -> intrvproto.UserDetails
	7, // 15: intrvproto.UserService.GetUserByToken:output_type -> intrvproto.UserDetails
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	file_user_proto_init()
	file_basic_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
