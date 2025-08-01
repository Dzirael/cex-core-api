// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: crypto/v1/service.proto

package crypto_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_crypto_v1_service_proto protoreflect.FileDescriptor

const file_crypto_v1_service_proto_rawDesc = "" +
	"\n" +
	"\x17crypto/v1/service.proto\x12\tcrypto.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x18crypto/v1/requests.proto\x1a\x18crypto/v1/response.proto2\x85\x02\n" +
	"\rCryptoService\x12y\n" +
	"\x12GetSupportedChains\x12$.crypto.v1.GetSupportedChainsRequest\x1a%.crypto.v1.GetSupportedChainsResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/api/v1/chains\x12y\n" +
	"\x12GetSupportedTokens\x12$.crypto.v1.GetSupportedTokensRequest\x1a%.crypto.v1.GetSupportedTokensResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/api/v1/tokensB\x88\x01\n" +
	"\rcom.crypto.v1B\fServiceProtoP\x01Z$cex-core-api/gen/crypto/v1;crypto_v1\xa2\x02\x03CXX\xaa\x02\tCrypto.V1\xca\x02\tCrypto\\V1\xe2\x02\x15Crypto\\V1\\GPBMetadata\xea\x02\n" +
	"Crypto::V1b\x06proto3"

var file_crypto_v1_service_proto_goTypes = []any{
	(*GetSupportedChainsRequest)(nil),  // 0: crypto.v1.GetSupportedChainsRequest
	(*GetSupportedTokensRequest)(nil),  // 1: crypto.v1.GetSupportedTokensRequest
	(*GetSupportedChainsResponse)(nil), // 2: crypto.v1.GetSupportedChainsResponse
	(*GetSupportedTokensResponse)(nil), // 3: crypto.v1.GetSupportedTokensResponse
}
var file_crypto_v1_service_proto_depIdxs = []int32{
	0, // 0: crypto.v1.CryptoService.GetSupportedChains:input_type -> crypto.v1.GetSupportedChainsRequest
	1, // 1: crypto.v1.CryptoService.GetSupportedTokens:input_type -> crypto.v1.GetSupportedTokensRequest
	2, // 2: crypto.v1.CryptoService.GetSupportedChains:output_type -> crypto.v1.GetSupportedChainsResponse
	3, // 3: crypto.v1.CryptoService.GetSupportedTokens:output_type -> crypto.v1.GetSupportedTokensResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypto_v1_service_proto_init() }
func file_crypto_v1_service_proto_init() {
	if File_crypto_v1_service_proto != nil {
		return
	}
	file_crypto_v1_requests_proto_init()
	file_crypto_v1_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_crypto_v1_service_proto_rawDesc), len(file_crypto_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_v1_service_proto_goTypes,
		DependencyIndexes: file_crypto_v1_service_proto_depIdxs,
	}.Build()
	File_crypto_v1_service_proto = out.File
	file_crypto_v1_service_proto_goTypes = nil
	file_crypto_v1_service_proto_depIdxs = nil
}
