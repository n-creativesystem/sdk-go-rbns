// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: organization.proto

package proto

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

var File_organization_proto protoreflect.FileDescriptor

var file_organization_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xf2, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x6e,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x4b,
	0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x6e, 0x63, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x1a, 0x20, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x07, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x6e, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x45, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x6e, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x1a, 0x13, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x1a,
	0x13, 0x2e, 0x6e, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_organization_proto_goTypes = []interface{}{
	(*OrganizationEntity)(nil),       // 0: ncs.protobuf.organizationEntity
	(*OrganizationKey)(nil),          // 1: ncs.protobuf.organizationKey
	(*Empty)(nil),                    // 2: ncs.protobuf.empty
	(*OrganizationUpdateEntity)(nil), // 3: ncs.protobuf.organizationUpdateEntity
	(*OrganizationEntities)(nil),     // 4: ncs.protobuf.organizationEntities
}
var file_organization_proto_depIdxs = []int32{
	0, // 0: ncs.protobuf.Organization.Create:input_type -> ncs.protobuf.organizationEntity
	1, // 1: ncs.protobuf.Organization.FindById:input_type -> ncs.protobuf.organizationKey
	2, // 2: ncs.protobuf.Organization.FindAll:input_type -> ncs.protobuf.empty
	3, // 3: ncs.protobuf.Organization.Update:input_type -> ncs.protobuf.organizationUpdateEntity
	1, // 4: ncs.protobuf.Organization.Delete:input_type -> ncs.protobuf.organizationKey
	0, // 5: ncs.protobuf.Organization.Create:output_type -> ncs.protobuf.organizationEntity
	0, // 6: ncs.protobuf.Organization.FindById:output_type -> ncs.protobuf.organizationEntity
	4, // 7: ncs.protobuf.Organization.FindAll:output_type -> ncs.protobuf.organizationEntities
	2, // 8: ncs.protobuf.Organization.Update:output_type -> ncs.protobuf.empty
	2, // 9: ncs.protobuf.Organization.Delete:output_type -> ncs.protobuf.empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_organization_proto_init() }
func file_organization_proto_init() {
	if File_organization_proto != nil {
		return
	}
	file_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_organization_proto_goTypes,
		DependencyIndexes: file_organization_proto_depIdxs,
	}.Build()
	File_organization_proto = out.File
	file_organization_proto_rawDesc = nil
	file_organization_proto_goTypes = nil
	file_organization_proto_depIdxs = nil
}