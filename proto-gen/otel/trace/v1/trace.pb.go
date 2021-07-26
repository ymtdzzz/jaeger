// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trace/v1/trace.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v11 "github.com/jaegertracing/jaeger/proto-gen/otel/common/v1"
	v1 "github.com/jaegertracing/jaeger/proto-gen/otel/resource/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SpanKind is the type of span. Can be used to specify additional relationships between spans
// in addition to a parent/child relationship.
type Span_SpanKind int32

const (
	// Unspecified. Do NOT use as default.
	// Implementations MAY assume SpanKind to be INTERNAL when receiving UNSPECIFIED.
	Span_SPAN_KIND_UNSPECIFIED Span_SpanKind = 0
	// Indicates that the span represents an internal operation within an application,
	// as opposed to an operation happening at the boundaries. Default value.
	Span_SPAN_KIND_INTERNAL Span_SpanKind = 1
	// Indicates that the span covers server-side handling of an RPC or other
	// remote network request.
	Span_SPAN_KIND_SERVER Span_SpanKind = 2
	// Indicates that the span describes a request to some remote service.
	Span_SPAN_KIND_CLIENT Span_SpanKind = 3
	// Indicates that the span describes a producer sending a message to a broker.
	// Unlike CLIENT and SERVER, there is often no direct critical path latency relationship
	// between producer and consumer spans. A PRODUCER span ends when the message was accepted
	// by the broker while the logical processing of the message might span a much longer time.
	Span_SPAN_KIND_PRODUCER Span_SpanKind = 4
	// Indicates that the span describes consumer receiving a message from a broker.
	// Like the PRODUCER kind, there is often no direct critical path latency relationship
	// between producer and consumer spans.
	Span_SPAN_KIND_CONSUMER Span_SpanKind = 5
)

var Span_SpanKind_name = map[int32]string{
	0: "SPAN_KIND_UNSPECIFIED",
	1: "SPAN_KIND_INTERNAL",
	2: "SPAN_KIND_SERVER",
	3: "SPAN_KIND_CLIENT",
	4: "SPAN_KIND_PRODUCER",
	5: "SPAN_KIND_CONSUMER",
}

var Span_SpanKind_value = map[string]int32{
	"SPAN_KIND_UNSPECIFIED": 0,
	"SPAN_KIND_INTERNAL":    1,
	"SPAN_KIND_SERVER":      2,
	"SPAN_KIND_CLIENT":      3,
	"SPAN_KIND_PRODUCER":    4,
	"SPAN_KIND_CONSUMER":    5,
}

func (x Span_SpanKind) String() string {
	return proto.EnumName(Span_SpanKind_name, int32(x))
}

func (Span_SpanKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{2, 0}
}

type Status_DeprecatedStatusCode int32

const (
	Status_DEPRECATED_STATUS_CODE_OK                  Status_DeprecatedStatusCode = 0
	Status_DEPRECATED_STATUS_CODE_CANCELLED           Status_DeprecatedStatusCode = 1
	Status_DEPRECATED_STATUS_CODE_UNKNOWN_ERROR       Status_DeprecatedStatusCode = 2
	Status_DEPRECATED_STATUS_CODE_INVALID_ARGUMENT    Status_DeprecatedStatusCode = 3
	Status_DEPRECATED_STATUS_CODE_DEADLINE_EXCEEDED   Status_DeprecatedStatusCode = 4
	Status_DEPRECATED_STATUS_CODE_NOT_FOUND           Status_DeprecatedStatusCode = 5
	Status_DEPRECATED_STATUS_CODE_ALREADY_EXISTS      Status_DeprecatedStatusCode = 6
	Status_DEPRECATED_STATUS_CODE_PERMISSION_DENIED   Status_DeprecatedStatusCode = 7
	Status_DEPRECATED_STATUS_CODE_RESOURCE_EXHAUSTED  Status_DeprecatedStatusCode = 8
	Status_DEPRECATED_STATUS_CODE_FAILED_PRECONDITION Status_DeprecatedStatusCode = 9
	Status_DEPRECATED_STATUS_CODE_ABORTED             Status_DeprecatedStatusCode = 10
	Status_DEPRECATED_STATUS_CODE_OUT_OF_RANGE        Status_DeprecatedStatusCode = 11
	Status_DEPRECATED_STATUS_CODE_UNIMPLEMENTED       Status_DeprecatedStatusCode = 12
	Status_DEPRECATED_STATUS_CODE_INTERNAL_ERROR      Status_DeprecatedStatusCode = 13
	Status_DEPRECATED_STATUS_CODE_UNAVAILABLE         Status_DeprecatedStatusCode = 14
	Status_DEPRECATED_STATUS_CODE_DATA_LOSS           Status_DeprecatedStatusCode = 15
	Status_DEPRECATED_STATUS_CODE_UNAUTHENTICATED     Status_DeprecatedStatusCode = 16
)

var Status_DeprecatedStatusCode_name = map[int32]string{
	0:  "DEPRECATED_STATUS_CODE_OK",
	1:  "DEPRECATED_STATUS_CODE_CANCELLED",
	2:  "DEPRECATED_STATUS_CODE_UNKNOWN_ERROR",
	3:  "DEPRECATED_STATUS_CODE_INVALID_ARGUMENT",
	4:  "DEPRECATED_STATUS_CODE_DEADLINE_EXCEEDED",
	5:  "DEPRECATED_STATUS_CODE_NOT_FOUND",
	6:  "DEPRECATED_STATUS_CODE_ALREADY_EXISTS",
	7:  "DEPRECATED_STATUS_CODE_PERMISSION_DENIED",
	8:  "DEPRECATED_STATUS_CODE_RESOURCE_EXHAUSTED",
	9:  "DEPRECATED_STATUS_CODE_FAILED_PRECONDITION",
	10: "DEPRECATED_STATUS_CODE_ABORTED",
	11: "DEPRECATED_STATUS_CODE_OUT_OF_RANGE",
	12: "DEPRECATED_STATUS_CODE_UNIMPLEMENTED",
	13: "DEPRECATED_STATUS_CODE_INTERNAL_ERROR",
	14: "DEPRECATED_STATUS_CODE_UNAVAILABLE",
	15: "DEPRECATED_STATUS_CODE_DATA_LOSS",
	16: "DEPRECATED_STATUS_CODE_UNAUTHENTICATED",
}

var Status_DeprecatedStatusCode_value = map[string]int32{
	"DEPRECATED_STATUS_CODE_OK":                  0,
	"DEPRECATED_STATUS_CODE_CANCELLED":           1,
	"DEPRECATED_STATUS_CODE_UNKNOWN_ERROR":       2,
	"DEPRECATED_STATUS_CODE_INVALID_ARGUMENT":    3,
	"DEPRECATED_STATUS_CODE_DEADLINE_EXCEEDED":   4,
	"DEPRECATED_STATUS_CODE_NOT_FOUND":           5,
	"DEPRECATED_STATUS_CODE_ALREADY_EXISTS":      6,
	"DEPRECATED_STATUS_CODE_PERMISSION_DENIED":   7,
	"DEPRECATED_STATUS_CODE_RESOURCE_EXHAUSTED":  8,
	"DEPRECATED_STATUS_CODE_FAILED_PRECONDITION": 9,
	"DEPRECATED_STATUS_CODE_ABORTED":             10,
	"DEPRECATED_STATUS_CODE_OUT_OF_RANGE":        11,
	"DEPRECATED_STATUS_CODE_UNIMPLEMENTED":       12,
	"DEPRECATED_STATUS_CODE_INTERNAL_ERROR":      13,
	"DEPRECATED_STATUS_CODE_UNAVAILABLE":         14,
	"DEPRECATED_STATUS_CODE_DATA_LOSS":           15,
	"DEPRECATED_STATUS_CODE_UNAUTHENTICATED":     16,
}

func (x Status_DeprecatedStatusCode) String() string {
	return proto.EnumName(Status_DeprecatedStatusCode_name, int32(x))
}

func (Status_DeprecatedStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{3, 0}
}

// For the semantics of status codes see
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/api.md#set-status
type Status_StatusCode int32

const (
	// The default status.
	Status_STATUS_CODE_UNSET Status_StatusCode = 0
	// The Span has been validated by an Application developers or Operator to have
	// completed successfully.
	Status_STATUS_CODE_OK Status_StatusCode = 1
	// The Span contains an error.
	Status_STATUS_CODE_ERROR Status_StatusCode = 2
)

var Status_StatusCode_name = map[int32]string{
	0: "STATUS_CODE_UNSET",
	1: "STATUS_CODE_OK",
	2: "STATUS_CODE_ERROR",
}

var Status_StatusCode_value = map[string]int32{
	"STATUS_CODE_UNSET": 0,
	"STATUS_CODE_OK":    1,
	"STATUS_CODE_ERROR": 2,
}

func (x Status_StatusCode) String() string {
	return proto.EnumName(Status_StatusCode_name, int32(x))
}

func (Status_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{3, 1}
}

// A collection of InstrumentationLibrarySpans from a Resource.
type ResourceSpans struct {
	// The resource for the spans in this message.
	// If this field is not set then no resource info is known.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of InstrumentationLibrarySpans that originate from a resource.
	InstrumentationLibrarySpans []*InstrumentationLibrarySpans `protobuf:"bytes,2,rep,name=instrumentation_library_spans,json=instrumentationLibrarySpans,proto3" json:"instrumentation_library_spans,omitempty"`
	// This schema_url applies to the data in the "resource" field. It does not apply
	// to the data in the "instrumentation_library_spans" field which have their own
	// schema_url field.
	SchemaUrl            string   `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schema_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceSpans) Reset()         { *m = ResourceSpans{} }
func (m *ResourceSpans) String() string { return proto.CompactTextString(m) }
func (*ResourceSpans) ProtoMessage()    {}
func (*ResourceSpans) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{0}
}
func (m *ResourceSpans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceSpans.Unmarshal(m, b)
}
func (m *ResourceSpans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceSpans.Marshal(b, m, deterministic)
}
func (m *ResourceSpans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceSpans.Merge(m, src)
}
func (m *ResourceSpans) XXX_Size() int {
	return xxx_messageInfo_ResourceSpans.Size(m)
}
func (m *ResourceSpans) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceSpans.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceSpans proto.InternalMessageInfo

func (m *ResourceSpans) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceSpans) GetInstrumentationLibrarySpans() []*InstrumentationLibrarySpans {
	if m != nil {
		return m.InstrumentationLibrarySpans
	}
	return nil
}

func (m *ResourceSpans) GetSchemaUrl() string {
	if m != nil {
		return m.SchemaUrl
	}
	return ""
}

// A collection of Spans produced by an InstrumentationLibrary.
type InstrumentationLibrarySpans struct {
	// The instrumentation library information for the spans in this message.
	// Semantically when InstrumentationLibrary isn't set, it is equivalent with
	// an empty instrumentation library name (unknown).
	InstrumentationLibrary *v11.InstrumentationLibrary `protobuf:"bytes,1,opt,name=instrumentation_library,json=instrumentationLibrary,proto3" json:"instrumentation_library,omitempty"`
	// A list of Spans that originate from an instrumentation library.
	Spans []*Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	// This schema_url applies to all spans and span events in the "spans" field.
	SchemaUrl            string   `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schema_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentationLibrarySpans) Reset()         { *m = InstrumentationLibrarySpans{} }
func (m *InstrumentationLibrarySpans) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibrarySpans) ProtoMessage()    {}
func (*InstrumentationLibrarySpans) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{1}
}
func (m *InstrumentationLibrarySpans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibrarySpans.Unmarshal(m, b)
}
func (m *InstrumentationLibrarySpans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibrarySpans.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibrarySpans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibrarySpans.Merge(m, src)
}
func (m *InstrumentationLibrarySpans) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibrarySpans.Size(m)
}
func (m *InstrumentationLibrarySpans) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibrarySpans.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibrarySpans proto.InternalMessageInfo

func (m *InstrumentationLibrarySpans) GetInstrumentationLibrary() *v11.InstrumentationLibrary {
	if m != nil {
		return m.InstrumentationLibrary
	}
	return nil
}

func (m *InstrumentationLibrarySpans) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *InstrumentationLibrarySpans) GetSchemaUrl() string {
	if m != nil {
		return m.SchemaUrl
	}
	return ""
}

// Span represents a single operation within a trace. Spans can be
// nested to form a trace tree. Spans may also be linked to other spans
// from the same or different trace and form graphs. Often, a trace
// contains a root span that describes the end-to-end latency, and one
// or more subspans for its sub-operations. A trace can also contain
// multiple root spans, or none at all. Spans do not need to be
// contiguous - there may be gaps or overlaps between spans in a trace.
//
// The next available field id is 17.
type Span struct {
	// A unique identifier for a trace. All spans from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random trace_id if empty or invalid trace_id was received.
	//
	// This field is required.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random span_id if empty or invalid span_id was received.
	//
	// This field is required.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// trace_state conveys information about request position in multiple distributed tracing graphs.
	// It is a trace_state in w3c-trace-context format: https://www.w3.org/TR/trace-context/#tracestate-header
	// See also https://github.com/w3c/distributed-tracing for more details about this field.
	TraceState string `protobuf:"bytes,3,opt,name=trace_state,json=traceState,proto3" json:"trace_state,omitempty"`
	// The `span_id` of this span's parent span. If this is a root span, then this
	// field must be empty. The ID is an 8-byte array.
	ParentSpanId []byte `protobuf:"bytes,4,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// A description of the span's operation.
	//
	// For example, the name can be a qualified method name or a file name
	// and a line number where the operation is called. A best practice is to use
	// the same display name at the same call point in an application.
	// This makes it easier to correlate spans in different traces.
	//
	// This field is semantically required to be set to non-empty string.
	// When null or empty string received - receiver may use string "name"
	// as a replacement. There might be smarted algorithms implemented by
	// receiver to fix the empty span name.
	//
	// This field is required.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `CLIENT` (caller)
	// and `SERVER` (callee) to identify queueing latency associated with the span.
	Kind Span_SpanKind `protobuf:"varint,6,opt,name=kind,proto3,enum=jaeger.trace.v1.Span_SpanKind" json:"kind,omitempty"`
	// start_time_unix_nano is the start time of the span. On the client side, this is the time
	// kept by the local machine where the span execution starts. On the server side, this
	// is the time when the server's application handler starts running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	StartTimeUnixNano uint64 `protobuf:"fixed64,7,opt,name=start_time_unix_nano,json=startTimeUnixNano,proto3" json:"start_time_unix_nano,omitempty"`
	// end_time_unix_nano is the end time of the span. On the client side, this is the time
	// kept by the local machine where the span execution ends. On the server side, this
	// is the time when the server application handler stops running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	EndTimeUnixNano uint64 `protobuf:"fixed64,8,opt,name=end_time_unix_nano,json=endTimeUnixNano,proto3" json:"end_time_unix_nano,omitempty"`
	// attributes is a collection of key/value pairs. The value can be a string,
	// an integer, a double or the Boolean values `true` or `false`. Note, global attributes
	// like server name can be set using the resource API. Examples of attributes:
	//
	//     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	//     "/http/server_latency": 300
	//     "abc.com/myattribute": true
	//     "abc.com/score": 10.239
	Attributes []*v11.KeyValue `protobuf:"bytes,9,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of attributes that were discarded. Attributes
	// can be discarded because their keys are too long or because there are too many
	// attributes. If this value is 0, then no attributes were dropped.
	DroppedAttributesCount uint32 `protobuf:"varint,10,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// events is a collection of Event items.
	Events []*Span_Event `protobuf:"bytes,11,rep,name=events,proto3" json:"events,omitempty"`
	// dropped_events_count is the number of dropped events. If the value is 0, then no
	// events were dropped.
	DroppedEventsCount uint32 `protobuf:"varint,12,opt,name=dropped_events_count,json=droppedEventsCount,proto3" json:"dropped_events_count,omitempty"`
	// links is a collection of Links, which are references from this span to a span
	// in the same or different trace.
	Links []*Span_Link `protobuf:"bytes,13,rep,name=links,proto3" json:"links,omitempty"`
	// dropped_links_count is the number of dropped links after the maximum size was
	// enforced. If this value is 0, then no links were dropped.
	DroppedLinksCount uint32 `protobuf:"varint,14,opt,name=dropped_links_count,json=droppedLinksCount,proto3" json:"dropped_links_count,omitempty"`
	// An optional final status for this span. Semantically when Status isn't set, it means
	// span's status code is unset, i.e. assume STATUS_CODE_UNSET (code = 0).
	Status               *Status  `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{2}
}
func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *Span) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *Span) GetTraceState() string {
	if m != nil {
		return m.TraceState
	}
	return ""
}

func (m *Span) GetParentSpanId() []byte {
	if m != nil {
		return m.ParentSpanId
	}
	return nil
}

func (m *Span) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span) GetKind() Span_SpanKind {
	if m != nil {
		return m.Kind
	}
	return Span_SPAN_KIND_UNSPECIFIED
}

func (m *Span) GetStartTimeUnixNano() uint64 {
	if m != nil {
		return m.StartTimeUnixNano
	}
	return 0
}

func (m *Span) GetEndTimeUnixNano() uint64 {
	if m != nil {
		return m.EndTimeUnixNano
	}
	return 0
}

func (m *Span) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Span) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func (m *Span) GetEvents() []*Span_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Span) GetDroppedEventsCount() uint32 {
	if m != nil {
		return m.DroppedEventsCount
	}
	return 0
}

func (m *Span) GetLinks() []*Span_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Span) GetDroppedLinksCount() uint32 {
	if m != nil {
		return m.DroppedLinksCount
	}
	return 0
}

func (m *Span) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Event is a time-stamped annotation of the span, consisting of user-supplied
// text description and key-value pairs.
type Span_Event struct {
	// time_unix_nano is the time the event occurred.
	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	// name of the event.
	// This field is semantically required to be set to non-empty string.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// attributes is a collection of attribute key/value pairs on the event.
	Attributes []*v11.KeyValue `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32   `protobuf:"varint,4,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Span_Event) Reset()         { *m = Span_Event{} }
func (m *Span_Event) String() string { return proto.CompactTextString(m) }
func (*Span_Event) ProtoMessage()    {}
func (*Span_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{2, 0}
}
func (m *Span_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_Event.Unmarshal(m, b)
}
func (m *Span_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_Event.Marshal(b, m, deterministic)
}
func (m *Span_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_Event.Merge(m, src)
}
func (m *Span_Event) XXX_Size() int {
	return xxx_messageInfo_Span_Event.Size(m)
}
func (m *Span_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Span_Event proto.InternalMessageInfo

func (m *Span_Event) GetTimeUnixNano() uint64 {
	if m != nil {
		return m.TimeUnixNano
	}
	return 0
}

func (m *Span_Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span_Event) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Span_Event) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

// A pointer from the current span to another span in the same trace or in a
// different trace. For example, this can be used in batching operations,
// where a single batch handler processes multiple requests from different
// traces or when the handler receives a request from a different project.
type Span_Link struct {
	// A unique identifier of a trace that this linked span is part of. The ID is a
	// 16-byte array.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for the linked span. The ID is an 8-byte array.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// The trace_state associated with the link.
	TraceState string `protobuf:"bytes,3,opt,name=trace_state,json=traceState,proto3" json:"trace_state,omitempty"`
	// attributes is a collection of attribute key/value pairs on the link.
	Attributes []*v11.KeyValue `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32   `protobuf:"varint,5,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Span_Link) Reset()         { *m = Span_Link{} }
func (m *Span_Link) String() string { return proto.CompactTextString(m) }
func (*Span_Link) ProtoMessage()    {}
func (*Span_Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{2, 1}
}
func (m *Span_Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span_Link.Unmarshal(m, b)
}
func (m *Span_Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span_Link.Marshal(b, m, deterministic)
}
func (m *Span_Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span_Link.Merge(m, src)
}
func (m *Span_Link) XXX_Size() int {
	return xxx_messageInfo_Span_Link.Size(m)
}
func (m *Span_Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Span_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Span_Link proto.InternalMessageInfo

func (m *Span_Link) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *Span_Link) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *Span_Link) GetTraceState() string {
	if m != nil {
		return m.TraceState
	}
	return ""
}

func (m *Span_Link) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Span_Link) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

// The Status type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs.
type Status struct {
	// The deprecated status code. This is an optional field.
	//
	// This field is deprecated and is replaced by the `code` field below. See backward
	// compatibility notes below. According to our stability guarantees this field
	// will be removed in 12 months, on Oct 22, 2021. All usage of old senders and
	// receivers that do not understand the `code` field MUST be phased out by then.
	DeprecatedCode Status_DeprecatedStatusCode `protobuf:"varint,1,opt,name=deprecated_code,json=deprecatedCode,proto3,enum=jaeger.trace.v1.Status_DeprecatedStatusCode" json:"deprecated_code,omitempty"` // Deprecated: Do not use.
	// A developer-facing human readable error message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The status code.
	Code                 Status_StatusCode `protobuf:"varint,3,opt,name=code,proto3,enum=jaeger.trace.v1.Status_StatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52825641200f25e, []int{3}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Status) GetDeprecatedCode() Status_DeprecatedStatusCode {
	if m != nil {
		return m.DeprecatedCode
	}
	return Status_DEPRECATED_STATUS_CODE_OK
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetCode() Status_StatusCode {
	if m != nil {
		return m.Code
	}
	return Status_STATUS_CODE_UNSET
}

func init() {
	proto.RegisterEnum("jaeger.trace.v1.Span_SpanKind", Span_SpanKind_name, Span_SpanKind_value)
	proto.RegisterEnum("jaeger.trace.v1.Status_DeprecatedStatusCode", Status_DeprecatedStatusCode_name, Status_DeprecatedStatusCode_value)
	proto.RegisterEnum("jaeger.trace.v1.Status_StatusCode", Status_StatusCode_name, Status_StatusCode_value)
	proto.RegisterType((*ResourceSpans)(nil), "jaeger.trace.v1.ResourceSpans")
	proto.RegisterType((*InstrumentationLibrarySpans)(nil), "jaeger.trace.v1.InstrumentationLibrarySpans")
	proto.RegisterType((*Span)(nil), "jaeger.trace.v1.Span")
	proto.RegisterType((*Span_Event)(nil), "jaeger.trace.v1.Span.Event")
	proto.RegisterType((*Span_Link)(nil), "jaeger.trace.v1.Span.Link")
	proto.RegisterType((*Status)(nil), "jaeger.trace.v1.Status")
}

func init() { proto.RegisterFile("trace/v1/trace.proto", fileDescriptor_a52825641200f25e) }

var fileDescriptor_a52825641200f25e = []byte{
	// 1144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x53, 0xdb, 0xd6,
	0x17, 0x8d, 0xc0, 0x7f, 0xe0, 0x02, 0x46, 0x79, 0x3f, 0x42, 0x14, 0xe7, 0x97, 0xc4, 0xe3, 0xa6,
	0x89, 0xf3, 0xcf, 0x4e, 0xc8, 0x4c, 0x9b, 0xe9, 0x4e, 0x48, 0x8f, 0x44, 0x83, 0x90, 0x3c, 0x4f,
	0x12, 0x4d, 0xba, 0x79, 0x23, 0xac, 0x37, 0x44, 0xc5, 0x96, 0x3c, 0x92, 0xcc, 0x84, 0x8f, 0xd2,
	0x4f, 0xd2, 0x75, 0xf7, 0xdd, 0x74, 0x9b, 0x65, 0x3f, 0x49, 0xe7, 0x3d, 0x49, 0x80, 0x01, 0xb9,
	0x9b, 0x74, 0xc3, 0x3c, 0x9d, 0x73, 0xee, 0x3d, 0xf7, 0xea, 0xde, 0x87, 0x05, 0x5b, 0x59, 0xe2,
	0x8f, 0xd8, 0xe0, 0xf4, 0xcd, 0x40, 0x1c, 0xfa, 0xd3, 0x24, 0xce, 0x62, 0xb4, 0xf9, 0xab, 0xcf,
	0x8e, 0x59, 0xd2, 0xcf, 0xb1, 0xd3, 0x37, 0xed, 0xed, 0x51, 0x3c, 0x99, 0xc4, 0x11, 0xd7, 0xe5,
	0xa7, 0x5c, 0xd8, 0x6e, 0x27, 0x2c, 0x8d, 0x67, 0x49, 0x9e, 0xa1, 0x3c, 0xe7, 0x5c, 0xf7, 0xab,
	0x04, 0x1b, 0xa4, 0x80, 0x9c, 0xa9, 0x1f, 0xa5, 0xe8, 0x1d, 0xac, 0x94, 0x1a, 0x45, 0xea, 0x48,
	0xbd, 0xb5, 0x9d, 0xff, 0xf7, 0x0b, 0xa7, 0xf3, 0xd8, 0xd3, 0x37, 0xfd, 0x32, 0x88, 0x9c, 0xab,
	0xd1, 0x14, 0x1e, 0x84, 0x51, 0x9a, 0x25, 0xb3, 0x09, 0x8b, 0x32, 0x3f, 0x0b, 0xe3, 0x88, 0x8e,
	0xc3, 0xa3, 0xc4, 0x4f, 0xce, 0x68, 0xca, 0x53, 0x2b, 0x4b, 0x9d, 0xe5, 0xde, 0xda, 0xce, 0xcb,
	0xfe, 0x95, 0xc2, 0xfb, 0xc6, 0x7c, 0x94, 0x99, 0x07, 0x89, 0x72, 0xc8, 0xfd, 0xb0, 0x9a, 0x44,
	0x0f, 0x00, 0xd2, 0xd1, 0x67, 0x36, 0xf1, 0xe9, 0x2c, 0x19, 0x2b, 0xcb, 0x1d, 0xa9, 0xb7, 0x4a,
	0x56, 0x73, 0xc4, 0x4b, 0xc6, 0xdd, 0x3f, 0x25, 0xb8, 0xbf, 0x20, 0x37, 0xf2, 0xe1, 0x6e, 0x45,
	0xc1, 0x45, 0xe7, 0xbd, 0xb2, 0xd4, 0xe2, 0x7d, 0x56, 0xd6, 0x4a, 0xb6, 0x6f, 0x2e, 0x13, 0xbd,
	0x80, 0xfa, 0xe5, 0xde, 0xef, 0x5c, 0xeb, 0x9d, 0x57, 0x42, 0x72, 0xcd, 0xbf, 0xb5, 0xf3, 0xc7,
	0x2a, 0xd4, 0xb8, 0x1c, 0xdd, 0x83, 0x15, 0x11, 0x4f, 0xc3, 0x40, 0x14, 0xba, 0x4e, 0x9a, 0xe2,
	0xd9, 0x08, 0xd0, 0x5d, 0x68, 0xf2, 0x5c, 0x9c, 0x59, 0x12, 0x4c, 0x83, 0x3f, 0x1a, 0x01, 0x7a,
	0x04, 0x6b, 0x79, 0x4c, 0x9a, 0xf9, 0x19, 0x2b, 0x92, 0x83, 0x80, 0x1c, 0x8e, 0xa0, 0xc7, 0xd0,
	0x9a, 0xfa, 0x09, 0x8b, 0x32, 0x5a, 0x26, 0xa8, 0x89, 0x04, 0xeb, 0x39, 0xea, 0xe4, 0x69, 0x10,
	0xd4, 0x22, 0x7f, 0xc2, 0x94, 0xba, 0x88, 0x17, 0x67, 0xb4, 0x03, 0xb5, 0x93, 0x30, 0x0a, 0x94,
	0x46, 0x47, 0xea, 0xb5, 0x76, 0x1e, 0xde, 0xd8, 0xa2, 0xf8, 0xb3, 0x1f, 0x46, 0x01, 0x11, 0x5a,
	0x34, 0x80, 0xad, 0x34, 0xf3, 0x93, 0x8c, 0x66, 0xe1, 0x84, 0xd1, 0x59, 0x14, 0x7e, 0xa1, 0x91,
	0x1f, 0xc5, 0x4a, 0xb3, 0x23, 0xf5, 0x1a, 0xe4, 0xb6, 0xe0, 0xdc, 0x70, 0xc2, 0xbc, 0x28, 0xfc,
	0x62, 0xf9, 0x51, 0x8c, 0x5e, 0x00, 0x62, 0x51, 0x70, 0x55, 0xbe, 0x22, 0xe4, 0x9b, 0x2c, 0x0a,
	0xe6, 0xc4, 0x3f, 0x01, 0xf8, 0x59, 0x96, 0x84, 0x47, 0xb3, 0x8c, 0xa5, 0xca, 0xaa, 0x78, 0xf5,
	0xed, 0xeb, 0xb3, 0xdc, 0x67, 0x67, 0x87, 0xfe, 0x78, 0xc6, 0xc8, 0x25, 0x35, 0x7a, 0x07, 0x4a,
	0x90, 0xc4, 0xd3, 0x29, 0x0b, 0xe8, 0x05, 0x4a, 0x47, 0xf1, 0x2c, 0xca, 0x14, 0xe8, 0x48, 0xbd,
	0x0d, 0xb2, 0x5d, 0xf0, 0xea, 0x39, 0xad, 0x71, 0x16, 0xbd, 0x85, 0x06, 0x3b, 0x65, 0x51, 0x96,
	0x2a, 0x6b, 0xc2, 0xf1, 0xfe, 0xcd, 0x6f, 0x02, 0x73, 0x0d, 0x29, 0xa4, 0xe8, 0x35, 0x6c, 0x95,
	0x76, 0x39, 0x52, 0x58, 0xad, 0x0b, 0x2b, 0x54, 0x70, 0x22, 0xa6, 0xb0, 0x79, 0x0d, 0xf5, 0x71,
	0x18, 0x9d, 0xa4, 0xca, 0xc6, 0x7c, 0x5f, 0xf3, 0x2e, 0x66, 0x18, 0x9d, 0x90, 0x5c, 0x88, 0xfa,
	0xf0, 0xbf, 0xd2, 0x43, 0x00, 0x85, 0x45, 0x4b, 0x58, 0xdc, 0x2e, 0x28, 0x1e, 0x50, 0x38, 0x0c,
	0xa0, 0xc1, 0xb7, 0x64, 0x96, 0x2a, 0x9b, 0xe2, 0x1a, 0xdc, 0xbd, 0x6e, 0x21, 0x68, 0x52, 0xc8,
	0xda, 0xbf, 0x4b, 0x50, 0x17, 0x25, 0xf2, 0x2d, 0xba, 0x32, 0x22, 0x49, 0x8c, 0x68, 0x3d, 0xbb,
	0x3c, 0x9f, 0x72, 0x8b, 0x96, 0x2e, 0x6d, 0xd1, 0xfc, 0xcc, 0x96, 0xbf, 0xd9, 0xcc, 0x6a, 0x8b,
	0x66, 0xd6, 0xfe, 0x4b, 0x82, 0x1a, 0xef, 0xfc, 0xbf, 0xb9, 0x53, 0xf3, 0x3d, 0xd5, 0xbe, 0x59,
	0x4f, 0xf5, 0x45, 0x3d, 0x75, 0x7f, 0x93, 0x60, 0xa5, 0xbc, 0x6e, 0xe8, 0x1e, 0xdc, 0x71, 0x86,
	0xaa, 0x45, 0xf7, 0x0d, 0x4b, 0xa7, 0x9e, 0xe5, 0x0c, 0xb1, 0x66, 0xec, 0x19, 0x58, 0x97, 0x6f,
	0xa1, 0x6d, 0x40, 0x17, 0x94, 0x61, 0xb9, 0x98, 0x58, 0xaa, 0x29, 0x4b, 0x68, 0x0b, 0xe4, 0x0b,
	0xdc, 0xc1, 0xe4, 0x10, 0x13, 0x79, 0x69, 0x1e, 0xd5, 0x4c, 0x03, 0x5b, 0xae, 0xbc, 0x3c, 0x9f,
	0x63, 0x48, 0x6c, 0xdd, 0xd3, 0x30, 0x91, 0x6b, 0xf3, 0xb8, 0x66, 0x5b, 0x8e, 0x77, 0x80, 0x89,
	0x5c, 0xef, 0xfe, 0xdd, 0x84, 0x46, 0xbe, 0x3c, 0xe8, 0x13, 0x6c, 0x06, 0x6c, 0x9a, 0xb0, 0x91,
	0x9f, 0xb1, 0x80, 0x8e, 0xe2, 0x20, 0xff, 0xbd, 0x69, 0xdd, 0xf0, 0x03, 0x91, 0x47, 0xf4, 0xf5,
	0x73, 0x79, 0x0e, 0x68, 0x71, 0xc0, 0x76, 0x97, 0x14, 0x89, 0xb4, 0x2e, 0x12, 0x71, 0x0c, 0x29,
	0xd0, 0x9c, 0xb0, 0x34, 0xf5, 0x8f, 0xcb, 0x15, 0x2b, 0x1f, 0xd1, 0x0f, 0x50, 0x13, 0x4e, 0xcb,
	0xc2, 0xa9, 0x5b, 0xe5, 0x74, 0x91, 0x9f, 0x08, 0x7d, 0xf7, 0x6b, 0x1d, 0xb6, 0x6e, 0xb2, 0x47,
	0x0f, 0xe0, 0x9e, 0x8e, 0x87, 0x04, 0x6b, 0xaa, 0x8b, 0x75, 0xea, 0xb8, 0xaa, 0xeb, 0x39, 0x54,
	0xb3, 0x75, 0x4c, 0xed, 0x7d, 0xf9, 0x16, 0x7a, 0x0c, 0x9d, 0x0a, 0x5a, 0x53, 0x2d, 0x0d, 0x9b,
	0x26, 0xd6, 0x65, 0x09, 0xf5, 0xe0, 0x71, 0x85, 0xca, 0xb3, 0xf6, 0x2d, 0xfb, 0x67, 0x8b, 0x62,
	0x42, 0x6c, 0x3e, 0x85, 0x17, 0xf0, 0xb4, 0x42, 0x69, 0x58, 0x87, 0xaa, 0x69, 0xe8, 0x54, 0x25,
	0xef, 0xbd, 0x83, 0x7c, 0x38, 0x2f, 0xa1, 0x57, 0x21, 0xd6, 0xb1, 0xaa, 0x9b, 0x86, 0x85, 0x29,
	0xfe, 0xa8, 0x61, 0xac, 0x63, 0x5d, 0xae, 0x2d, 0x28, 0xd5, 0xb2, 0x5d, 0xba, 0x67, 0x7b, 0x96,
	0x2e, 0xd7, 0xd1, 0x33, 0xf8, 0xbe, 0x42, 0xa5, 0x9a, 0x04, 0xab, 0xfa, 0x27, 0x8a, 0x3f, 0x1a,
	0x8e, 0xeb, 0xc8, 0x8d, 0x05, 0xf6, 0x43, 0x4c, 0x0e, 0x0c, 0xc7, 0x31, 0x6c, 0x8b, 0xea, 0xd8,
	0xe2, 0xdb, 0xd8, 0x44, 0xaf, 0xe0, 0x59, 0x85, 0x9a, 0x60, 0xc7, 0xf6, 0x88, 0xc6, 0x8b, 0xfd,
	0xa0, 0x7a, 0x8e, 0x8b, 0x75, 0x79, 0x05, 0xf5, 0xe1, 0x79, 0x85, 0x7c, 0x4f, 0x35, 0x4c, 0xcc,
	0x97, 0x11, 0x6b, 0xb6, 0xa5, 0x1b, 0xae, 0x61, 0x5b, 0xf2, 0x2a, 0xea, 0xc2, 0xc3, 0xaa, 0xba,
	0x77, 0x6d, 0xc2, 0x73, 0x02, 0x7a, 0x0a, 0xdf, 0x55, 0xcd, 0xd2, 0x73, 0xa9, 0xbd, 0x47, 0x89,
	0x6a, 0xbd, 0xc7, 0xf2, 0xda, 0xc2, 0x79, 0x19, 0x07, 0x43, 0x13, 0xf3, 0x01, 0x60, 0x5d, 0x5e,
	0x5f, 0xf0, 0xba, 0xca, 0x0b, 0x57, 0x8c, 0x76, 0x03, 0x3d, 0x81, 0x6e, 0x65, 0x52, 0xf5, 0x50,
	0x35, 0x4c, 0x75, 0xd7, 0xc4, 0x72, 0x6b, 0xc1, 0x9c, 0x74, 0xd5, 0x55, 0xa9, 0x69, 0x3b, 0x8e,
	0xbc, 0x89, 0x9e, 0xc3, 0x93, 0xea, 0x6c, 0x9e, 0xfb, 0x01, 0x5b, 0xae, 0x21, 0x38, 0x59, 0xee,
	0x5a, 0x00, 0x97, 0x36, 0xfa, 0x0e, 0xdc, 0x9e, 0x97, 0x3b, 0xd8, 0x95, 0x6f, 0x21, 0x04, 0xad,
	0x2b, 0xdb, 0x2d, 0x5d, 0x95, 0x16, 0x4b, 0xba, 0xfb, 0x19, 0x1e, 0x85, 0x71, 0x3f, 0x9e, 0xb2,
	0x28, 0x63, 0x63, 0x36, 0x61, 0x59, 0x72, 0x96, 0x7f, 0x6e, 0x9e, 0xdf, 0xb3, 0x5d, 0x70, 0xf9,
	0x69, 0xc8, 0xc1, 0xa1, 0xf4, 0xcb, 0x8f, 0xc7, 0x61, 0xf6, 0x79, 0x76, 0xc4, 0xff, 0x25, 0x0e,
	0xf2, 0x1b, 0xc9, 0x85, 0x61, 0x74, 0x5c, 0x3c, 0x0d, 0x44, 0xf4, 0xab, 0x63, 0x16, 0x0d, 0xe2,
	0x8c, 0x8d, 0x07, 0xe5, 0xf7, 0xf0, 0x51, 0x43, 0x10, 0x6f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x34, 0x06, 0x1c, 0x22, 0x0b, 0x00, 0x00,
}