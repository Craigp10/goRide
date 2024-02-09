// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/routeRaydar.proto

package routeRaydar

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Grid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Grid) Reset() {
	*x = Grid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grid) ProtoMessage() {}

func (x *Grid) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grid.ProtoReflect.Descriptor instead.
func (*Grid) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{0}
}

func (x *Grid) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []int64 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetValues() []int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{2}
}

func (x *Coordinates) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinates) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

// The request message containing the route ID.
type GetRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId string `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
}

func (x *GetRouteRequest) Reset() {
	*x = GetRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteRequest) ProtoMessage() {}

func (x *GetRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteRequest.ProtoReflect.Descriptor instead.
func (*GetRouteRequest) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{3}
}

func (x *GetRouteRequest) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

// The response message containing the route details.
type GetRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId string `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
}

func (x *GetRouteResponse) Reset() {
	*x = GetRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteResponse) ProtoMessage() {}

func (x *GetRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteResponse.ProtoReflect.Descriptor instead.
func (*GetRouteResponse) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{4}
}

func (x *GetRouteResponse) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

type SubmitGridRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define the structure of the grid request message
	// Add fields to represent the grid data
	Width  int64 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *SubmitGridRequest) Reset() {
	*x = SubmitGridRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitGridRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitGridRequest) ProtoMessage() {}

func (x *SubmitGridRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitGridRequest.ProtoReflect.Descriptor instead.
func (*SubmitGridRequest) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{5}
}

func (x *SubmitGridRequest) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *SubmitGridRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type SubmitGridResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define the structure of the grid response message
	// Add fields as needed
	Grid *Grid `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
}

func (x *SubmitGridResponse) Reset() {
	*x = SubmitGridResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitGridResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitGridResponse) ProtoMessage() {}

func (x *SubmitGridResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitGridResponse.ProtoReflect.Descriptor instead.
func (*SubmitGridResponse) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{6}
}

func (x *SubmitGridResponse) GetGrid() *Grid {
	if x != nil {
		return x.Grid
	}
	return nil
}

type SendCoordinatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define the structure of the new points request message
	// Add fields to represent the new points data
	Start *Coordinates `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *Coordinates `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *SendCoordinatesRequest) Reset() {
	*x = SendCoordinatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCoordinatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCoordinatesRequest) ProtoMessage() {}

func (x *SendCoordinatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCoordinatesRequest.ProtoReflect.Descriptor instead.
func (*SendCoordinatesRequest) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{7}
}

func (x *SendCoordinatesRequest) GetStart() *Coordinates {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *SendCoordinatesRequest) GetEnd() *Coordinates {
	if x != nil {
		return x.End
	}
	return nil
}

type SendCoordinatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define the structure of the new points response message
	// Add fields as needed
	Route    []*Coordinates `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty"`
	Distance int64          `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *SendCoordinatesResponse) Reset() {
	*x = SendCoordinatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_routeRaydar_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCoordinatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCoordinatesResponse) ProtoMessage() {}

func (x *SendCoordinatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_routeRaydar_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCoordinatesResponse.ProtoReflect.Descriptor instead.
func (*SendCoordinatesResponse) Descriptor() ([]byte, []int) {
	return file_api_routeRaydar_proto_rawDescGZIP(), []int{8}
}

func (x *SendCoordinatesResponse) GetRoute() []*Coordinates {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *SendCoordinatesResponse) GetDistance() int64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_api_routeRaydar_proto protoreflect.FileDescriptor

var file_api_routeRaydar_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61,
	0x79, 0x64, 0x61, 0x72, 0x22, 0x2c, 0x0a, 0x04, 0x47, 0x72, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x22, 0x1d, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x22, 0x29, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x22, 0x2c, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3b, 0x0a, 0x12,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x67, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x2e, 0x47,
	0x72, 0x69, 0x64, 0x52, 0x04, 0x67, 0x72, 0x69, 0x64, 0x22, 0x74, 0x0a, 0x16, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72,
	0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x2e, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22,
	0x65, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x84, 0x02, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x47, 0x72, 0x69, 0x64, 0x12, 0x1e,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5c, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x23, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x61, 0x79, 0x64, 0x61, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72, 0x61, 0x69,
	0x67, 0x70, 0x31, 0x30, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x79, 0x64, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_routeRaydar_proto_rawDescOnce sync.Once
	file_api_routeRaydar_proto_rawDescData = file_api_routeRaydar_proto_rawDesc
)

func file_api_routeRaydar_proto_rawDescGZIP() []byte {
	file_api_routeRaydar_proto_rawDescOnce.Do(func() {
		file_api_routeRaydar_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_routeRaydar_proto_rawDescData)
	})
	return file_api_routeRaydar_proto_rawDescData
}

var file_api_routeRaydar_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_routeRaydar_proto_goTypes = []interface{}{
	(*Grid)(nil),                    // 0: routeRaydar.Grid
	(*Row)(nil),                     // 1: routeRaydar.Row
	(*Coordinates)(nil),             // 2: routeRaydar.Coordinates
	(*GetRouteRequest)(nil),         // 3: routeRaydar.GetRouteRequest
	(*GetRouteResponse)(nil),        // 4: routeRaydar.GetRouteResponse
	(*SubmitGridRequest)(nil),       // 5: routeRaydar.SubmitGridRequest
	(*SubmitGridResponse)(nil),      // 6: routeRaydar.SubmitGridResponse
	(*SendCoordinatesRequest)(nil),  // 7: routeRaydar.SendCoordinatesRequest
	(*SendCoordinatesResponse)(nil), // 8: routeRaydar.SendCoordinatesResponse
}
var file_api_routeRaydar_proto_depIdxs = []int32{
	1, // 0: routeRaydar.Grid.rows:type_name -> routeRaydar.Row
	0, // 1: routeRaydar.SubmitGridResponse.grid:type_name -> routeRaydar.Grid
	2, // 2: routeRaydar.SendCoordinatesRequest.start:type_name -> routeRaydar.Coordinates
	2, // 3: routeRaydar.SendCoordinatesRequest.end:type_name -> routeRaydar.Coordinates
	2, // 4: routeRaydar.SendCoordinatesResponse.route:type_name -> routeRaydar.Coordinates
	3, // 5: routeRaydar.RouteService.GetRoute:input_type -> routeRaydar.GetRouteRequest
	5, // 6: routeRaydar.RouteService.SubmitGrid:input_type -> routeRaydar.SubmitGridRequest
	7, // 7: routeRaydar.RouteService.SendCoordinates:input_type -> routeRaydar.SendCoordinatesRequest
	4, // 8: routeRaydar.RouteService.GetRoute:output_type -> routeRaydar.GetRouteResponse
	6, // 9: routeRaydar.RouteService.SubmitGrid:output_type -> routeRaydar.SubmitGridResponse
	8, // 10: routeRaydar.RouteService.SendCoordinates:output_type -> routeRaydar.SendCoordinatesResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_routeRaydar_proto_init() }
func file_api_routeRaydar_proto_init() {
	if File_api_routeRaydar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_routeRaydar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitGridRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitGridResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCoordinatesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_routeRaydar_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCoordinatesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_routeRaydar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_routeRaydar_proto_goTypes,
		DependencyIndexes: file_api_routeRaydar_proto_depIdxs,
		MessageInfos:      file_api_routeRaydar_proto_msgTypes,
	}.Build()
	File_api_routeRaydar_proto = out.File
	file_api_routeRaydar_proto_rawDesc = nil
	file_api_routeRaydar_proto_goTypes = nil
	file_api_routeRaydar_proto_depIdxs = nil
}
