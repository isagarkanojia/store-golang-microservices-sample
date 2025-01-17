// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api_docs/catalogs/protobuf/products/kafka_messages/product_kafka_messages.proto

package kafka_messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   string  `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ProductCreate) Reset() {
	*x = ProductCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreate) ProtoMessage() {}

func (x *ProductCreate) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreate.ProtoReflect.Descriptor instead.
func (*ProductCreate) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCreate) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *ProductCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductCreate) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   string  `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ProductUpdate) Reset() {
	*x = ProductUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdate) ProtoMessage() {}

func (x *ProductUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdate.ProtoReflect.Descriptor instead.
func (*ProductUpdate) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ProductUpdate) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *ProductUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductUpdate) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   string                 `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64                `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Product) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ProductCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=Product,proto3" json:"Product,omitempty"`
}

func (x *ProductCreated) Reset() {
	*x = ProductCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreated) ProtoMessage() {}

func (x *ProductCreated) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreated.ProtoReflect.Descriptor instead.
func (*ProductCreated) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ProductCreated) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProductUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=Product,proto3" json:"Product,omitempty"`
}

func (x *ProductUpdated) Reset() {
	*x = ProductUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdated) ProtoMessage() {}

func (x *ProductUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdated.ProtoReflect.Descriptor instead.
func (*ProductUpdated) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ProductUpdated) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProductDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
}

func (x *ProductDelete) Reset() {
	*x = ProductDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDelete) ProtoMessage() {}

func (x *ProductDelete) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDelete.ProtoReflect.Descriptor instead.
func (*ProductDelete) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ProductDelete) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

type ProductDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
}

func (x *ProductDeleted) Reset() {
	*x = ProductDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDeleted) ProtoMessage() {}

func (x *ProductDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDeleted.ProtoReflect.Descriptor instead.
func (*ProductDeleted) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP(), []int{6}
}

func (x *ProductDeleted) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

var File_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto protoreflect.FileDescriptor

var file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x61, 0x70, 0x69, 0x5f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x79, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xe7, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x43, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x43, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x2d, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x2e, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2f, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescOnce sync.Once
	file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescData = file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDesc
)

func file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescGZIP() []byte {
	file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescOnce.Do(func() {
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescData)
	})
	return file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDescData
}

var file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_goTypes = []interface{}{
	(*ProductCreate)(nil),         // 0: kafka_messages.ProductCreate
	(*ProductUpdate)(nil),         // 1: kafka_messages.ProductUpdate
	(*Product)(nil),               // 2: kafka_messages.Product
	(*ProductCreated)(nil),        // 3: kafka_messages.ProductCreated
	(*ProductUpdated)(nil),        // 4: kafka_messages.ProductUpdated
	(*ProductDelete)(nil),         // 5: kafka_messages.ProductDelete
	(*ProductDeleted)(nil),        // 6: kafka_messages.ProductDeleted
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_depIdxs = []int32{
	7, // 0: kafka_messages.Product.CreatedAt:type_name -> google.protobuf.Timestamp
	7, // 1: kafka_messages.Product.UpdatedAt:type_name -> google.protobuf.Timestamp
	2, // 2: kafka_messages.ProductCreated.Product:type_name -> kafka_messages.Product
	2, // 3: kafka_messages.ProductUpdated.Product:type_name -> kafka_messages.Product
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_init()
}
func file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_init() {
	if File_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreate); i {
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
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductUpdate); i {
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
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreated); i {
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
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductUpdated); i {
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
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDelete); i {
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
		file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDeleted); i {
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
			RawDescriptor: file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_goTypes,
		DependencyIndexes: file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_depIdxs,
		MessageInfos:      file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_msgTypes,
	}.Build()
	File_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto = out.File
	file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_rawDesc = nil
	file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_goTypes = nil
	file_api_docs_catalogs_protobuf_products_kafka_messages_product_kafka_messages_proto_depIdxs = nil
}
