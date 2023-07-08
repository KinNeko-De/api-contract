// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: kinnekode/restaurant/document/v1/invoice.proto

package v1

import (
	protobuf "github.com/kinneko-de/api-contract/golang/kinnekode/protobuf"
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

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveredOn *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=delivered_on,json=deliveredOn,proto3" json:"delivered_on,omitempty"`
	// three character currency code as specified in ISO 4217 ( see https://de.wikipedia.org/wiki/ISO_4217 )
	CurrencyCode string             `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Recipient    *Invoice_Recipient `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Items        []*Invoice_Item    `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_kinnekode_restaurant_document_v1_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetDeliveredOn() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveredOn
	}
	return nil
}

func (x *Invoice) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Invoice) GetRecipient() *Invoice_Recipient {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *Invoice) GetItems() []*Invoice_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Invoice_Recipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Street   string `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	PostCode string `protobuf:"bytes,4,opt,name=postCode,proto3" json:"postCode,omitempty"`
	Country  string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Invoice_Recipient) Reset() {
	*x = Invoice_Recipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice_Recipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice_Recipient) ProtoMessage() {}

func (x *Invoice_Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice_Recipient.ProtoReflect.Descriptor instead.
func (*Invoice_Recipient) Descriptor() ([]byte, []int) {
	return file_kinnekode_restaurant_document_v1_invoice_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Invoice_Recipient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Invoice_Recipient) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Invoice_Recipient) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Invoice_Recipient) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *Invoice_Recipient) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Invoice_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string            `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Quantity    int64             `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	NetAmount   *protobuf.Decimal `protobuf:"bytes,3,opt,name=netAmount,proto3" json:"netAmount,omitempty"`
	Taxation    *protobuf.Decimal `protobuf:"bytes,4,opt,name=taxation,proto3" json:"taxation,omitempty"`
	TotalAmount *protobuf.Decimal `protobuf:"bytes,5,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	Sum         *protobuf.Decimal `protobuf:"bytes,6,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *Invoice_Item) Reset() {
	*x = Invoice_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice_Item) ProtoMessage() {}

func (x *Invoice_Item) ProtoReflect() protoreflect.Message {
	mi := &file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice_Item.ProtoReflect.Descriptor instead.
func (*Invoice_Item) Descriptor() ([]byte, []int) {
	return file_kinnekode_restaurant_document_v1_invoice_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Invoice_Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Invoice_Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Invoice_Item) GetNetAmount() *protobuf.Decimal {
	if x != nil {
		return x.NetAmount
	}
	return nil
}

func (x *Invoice_Item) GetTaxation() *protobuf.Decimal {
	if x != nil {
		return x.Taxation
	}
	return nil
}

func (x *Invoice_Item) GetTotalAmount() *protobuf.Decimal {
	if x != nil {
		return x.TotalAmount
	}
	return nil
}

func (x *Invoice_Item) GetSum() *protobuf.Decimal {
	if x != nil {
		return x.Sum
	}
	return nil
}

var File_kinnekode_restaurant_document_v1_invoice_proto protoreflect.FileDescriptor

var file_kinnekode_restaurant_document_v1_invoice_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x05, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6b, 0x69, 0x6e, 0x6e, 0x65,
	0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b,
	0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x81,
	0x01, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x1a, 0xa6, 0x02, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x09, 0x6e, 0x65, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b,
	0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x63, 0x69,
	0x6d, 0x61, 0x6c, 0x52, 0x08, 0x74, 0x61, 0x78, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x03,
	0x73, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x6e, 0x6e,
	0x65, 0x6b, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x42, 0x4c, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b,
	0x6f, 0x2d, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x69, 0x6e, 0x6e, 0x65, 0x6b, 0x6f,
	0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kinnekode_restaurant_document_v1_invoice_proto_rawDescOnce sync.Once
	file_kinnekode_restaurant_document_v1_invoice_proto_rawDescData = file_kinnekode_restaurant_document_v1_invoice_proto_rawDesc
)

func file_kinnekode_restaurant_document_v1_invoice_proto_rawDescGZIP() []byte {
	file_kinnekode_restaurant_document_v1_invoice_proto_rawDescOnce.Do(func() {
		file_kinnekode_restaurant_document_v1_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_kinnekode_restaurant_document_v1_invoice_proto_rawDescData)
	})
	return file_kinnekode_restaurant_document_v1_invoice_proto_rawDescData
}

var file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kinnekode_restaurant_document_v1_invoice_proto_goTypes = []interface{}{
	(*Invoice)(nil),               // 0: kinnekode.restaurant.document.v1.Invoice
	(*Invoice_Recipient)(nil),     // 1: kinnekode.restaurant.document.v1.Invoice.Recipient
	(*Invoice_Item)(nil),          // 2: kinnekode.restaurant.document.v1.Invoice.Item
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*protobuf.Decimal)(nil),      // 4: kinnekode.protobuf.Decimal
}
var file_kinnekode_restaurant_document_v1_invoice_proto_depIdxs = []int32{
	3, // 0: kinnekode.restaurant.document.v1.Invoice.delivered_on:type_name -> google.protobuf.Timestamp
	1, // 1: kinnekode.restaurant.document.v1.Invoice.recipient:type_name -> kinnekode.restaurant.document.v1.Invoice.Recipient
	2, // 2: kinnekode.restaurant.document.v1.Invoice.items:type_name -> kinnekode.restaurant.document.v1.Invoice.Item
	4, // 3: kinnekode.restaurant.document.v1.Invoice.Item.netAmount:type_name -> kinnekode.protobuf.Decimal
	4, // 4: kinnekode.restaurant.document.v1.Invoice.Item.taxation:type_name -> kinnekode.protobuf.Decimal
	4, // 5: kinnekode.restaurant.document.v1.Invoice.Item.totalAmount:type_name -> kinnekode.protobuf.Decimal
	4, // 6: kinnekode.restaurant.document.v1.Invoice.Item.sum:type_name -> kinnekode.protobuf.Decimal
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_kinnekode_restaurant_document_v1_invoice_proto_init() }
func file_kinnekode_restaurant_document_v1_invoice_proto_init() {
	if File_kinnekode_restaurant_document_v1_invoice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice_Recipient); i {
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
		file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice_Item); i {
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
			RawDescriptor: file_kinnekode_restaurant_document_v1_invoice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kinnekode_restaurant_document_v1_invoice_proto_goTypes,
		DependencyIndexes: file_kinnekode_restaurant_document_v1_invoice_proto_depIdxs,
		MessageInfos:      file_kinnekode_restaurant_document_v1_invoice_proto_msgTypes,
	}.Build()
	File_kinnekode_restaurant_document_v1_invoice_proto = out.File
	file_kinnekode_restaurant_document_v1_invoice_proto_rawDesc = nil
	file_kinnekode_restaurant_document_v1_invoice_proto_goTypes = nil
	file_kinnekode_restaurant_document_v1_invoice_proto_depIdxs = nil
}
