// Code generated by protoc-gen-go.
// source: plugin.proto
// DO NOT EDIT!

/*
Package openapi_plugin_v1 is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	Version
	Parameter
	Request
	Response
	File
	Wrapper
*/
package openapi_plugin_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The version number of OpenAPI compiler.
type Version struct {
	Major int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch int32 `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix string `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A parameter passed to the plugin from (or through) the OpenAPI compiler.
type Parameter struct {
	// The name of the parameter as specified in the option string
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The parameter value as specified in the option string
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Parameter) Reset()                    { *m = Parameter{} }
func (m *Parameter) String() string            { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()               {}
func (*Parameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// An encoded Request is written to the plugin's stdin.
type Request struct {
	// A wrapped OpenAPI document to process.
	Wrapper *Wrapper `protobuf:"bytes,1,opt,name=wrapper" json:"wrapper,omitempty"`
	// Plugin parameters parsed from the invocation string.
	Parameters []*Parameter `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
	// The version number of openapi compiler.
	CompilerVersion *Version `protobuf:"bytes,3,opt,name=compiler_version,json=compilerVersion" json:"compiler_version,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Request) GetWrapper() *Wrapper {
	if m != nil {
		return m.Wrapper
	}
	return nil
}

func (m *Request) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Request) GetCompilerVersion() *Version {
	if m != nil {
		return m.CompilerVersion
	}
	return nil
}

// The plugin writes an encoded Response to stdout.
type Response struct {
	// Error message.  If non-empty, the plugin failed.
	// The plugin process should exit with status code zero
	// even if it reports an error in this way.
	//
	// This should be used to indicate errors which prevent the plugin from
	// operating as intended.  Errors which indicate a problem in openapic
	// itself -- such as the input Document being unparseable -- should be
	// reported by writing a message to stderr and exiting with a non-zero
	// status code.
	Errors []string `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
	// file output, each file will be written by openapic to an appropriate location.
	Files []*File `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

// File describes a file generated by a plugin.
type File struct {
	// name of the file
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// data to be written to the file
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Wrapper wraps an OpenAPI document with its version.
type Wrapper struct {
	// filename or URL of the wrapped document
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// version of the OpenAPI specification that is used by the wrapped document
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// valid serialized protocol buffer of the named OpenAPI specification version
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Wrapper) Reset()                    { *m = Wrapper{} }
func (m *Wrapper) String() string            { return proto.CompactTextString(m) }
func (*Wrapper) ProtoMessage()               {}
func (*Wrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Version)(nil), "openapi.plugin.v1.Version")
	proto.RegisterType((*Parameter)(nil), "openapi.plugin.v1.Parameter")
	proto.RegisterType((*Request)(nil), "openapi.plugin.v1.Request")
	proto.RegisterType((*Response)(nil), "openapi.plugin.v1.Response")
	proto.RegisterType((*File)(nil), "openapi.plugin.v1.File")
	proto.RegisterType((*Wrapper)(nil), "openapi.plugin.v1.Wrapper")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xbb, 0x4e, 0xc3, 0x40,
	0x10, 0x94, 0xe3, 0x24, 0xc6, 0x4b, 0xa4, 0xc0, 0x89, 0x87, 0x85, 0x28, 0x90, 0x2b, 0x1a, 0x2c,
	0x9e, 0x1d, 0x4d, 0x12, 0x81, 0x44, 0x81, 0x62, 0xae, 0x80, 0x12, 0x1d, 0xe6, 0x12, 0x8c, 0x12,
	0xdf, 0x71, 0xb6, 0x03, 0xdf, 0xc3, 0xb7, 0xf0, 0x61, 0xac, 0xef, 0x01, 0x48, 0xb8, 0xdb, 0x19,
	0xef, 0xec, 0xce, 0xce, 0x19, 0x06, 0x72, 0x51, 0xcf, 0xf3, 0x22, 0x91, 0x4a, 0x54, 0x82, 0x6c,
	0x0a, 0xc9, 0x0b, 0x26, 0xf3, 0xc4, 0xb2, 0xab, 0x93, 0x38, 0x83, 0xe0, 0x9e, 0xab, 0x32, 0x17,
	0x05, 0xd9, 0x82, 0xde, 0x92, 0xbd, 0x0a, 0x15, 0x79, 0x07, 0xde, 0x61, 0x8f, 0x1a, 0xa0, 0xd9,
	0xbc, 0x40, 0xb6, 0x63, 0xd9, 0x06, 0x34, 0xac, 0x64, 0x55, 0xf6, 0x12, 0xf9, 0x86, 0xd5, 0x80,
	0xec, 0x40, 0xbf, 0xac, 0x67, 0xb3, 0xfc, 0x23, 0xea, 0x22, 0x1d, 0x52, 0x8b, 0xe2, 0x0b, 0x08,
	0x53, 0xa6, 0xd8, 0x92, 0x57, 0x5c, 0x11, 0x02, 0xdd, 0x02, 0x4b, 0xbd, 0x25, 0xa4, 0xba, 0x6e,
	0xc6, 0xad, 0xd8, 0xa2, 0xe6, 0x7a, 0x49, 0x48, 0x0d, 0x88, 0xbf, 0x3c, 0x08, 0x28, 0x7f, 0xab,
	0x79, 0x59, 0x91, 0x73, 0x08, 0xde, 0x15, 0x93, 0x92, 0x1b, 0x7b, 0xeb, 0xa7, 0x7b, 0xc9, 0xbf,
	0x63, 0x92, 0x07, 0xd3, 0x41, 0x5d, 0x2b, 0xb9, 0x04, 0x90, 0x6e, 0x71, 0x89, 0xc3, 0x7d, 0x14,
	0xee, 0xb7, 0x08, 0x7f, 0xdc, 0xd1, 0x3f, 0xfd, 0xe4, 0x0a, 0x36, 0x32, 0xb1, 0x94, 0xf9, 0x82,
	0xab, 0xc7, 0x95, 0x09, 0x49, 0xdf, 0xdb, 0xbe, 0xdc, 0xc6, 0x48, 0x87, 0x4e, 0x63, 0x89, 0xf8,
	0x0e, 0xd6, 0x28, 0x2f, 0xa5, 0x28, 0x4a, 0xde, 0x24, 0xc4, 0x95, 0x12, 0x68, 0xc6, 0x43, 0x33,
	0x98, 0x90, 0x41, 0xe4, 0x08, 0x7a, 0x33, 0xd4, 0x38, 0x8f, 0xbb, 0x2d, 0xf3, 0xaf, 0xf1, 0x3b,
	0x35, 0x5d, 0x71, 0x02, 0xdd, 0x06, 0xb6, 0x66, 0x89, 0xdc, 0x33, 0xab, 0x98, 0x8e, 0x72, 0x40,
	0x75, 0x1d, 0xdf, 0x42, 0x60, 0xb3, 0x69, 0x95, 0x44, 0x10, 0xb8, 0xfb, 0xcc, 0x03, 0x38, 0xf8,
	0xfb, 0x30, 0xbe, 0x9e, 0x66, 0xc0, 0xf8, 0x18, 0x86, 0x42, 0xcd, 0x9d, 0xc7, 0x0c, 0xdd, 0x8d,
	0xb7, 0xa7, 0x08, 0x46, 0xe9, 0xcd, 0xc4, 0x1e, 0x9f, 0x6a, 0xdf, 0xa9, 0xf7, 0xd9, 0xf1, 0xa7,
	0xa3, 0xc9, 0x53, 0x5f, 0xff, 0x80, 0x67, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0x7a, 0xf1,
	0x03, 0x90, 0x02, 0x00, 0x00,
}
