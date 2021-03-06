// Code generated by protoc-gen-go. DO NOT EDIT.
// source: public/api.proto

/*
Package conduit_public is a generated protocol buffer package.

It is generated from these files:
	public/api.proto

It has these top-level messages:
	HistogramValue
	Histogram
	MetricValue
	MetricDatapoint
	MetricSeries
	MetricMetadata
	MetricResponse
	MetricRequest
	Empty
	VersionInfo
	ListPodsResponse
	Pod
	TapRequest
	ApiError
*/
package conduit_public

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import conduit_common "github.com/runconduit/conduit/controller/gen/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MetricName int32

const (
	MetricName_REQUEST_RATE MetricName = 0
	MetricName_LATENCY      MetricName = 1
	MetricName_SUCCESS_RATE MetricName = 2
)

var MetricName_name = map[int32]string{
	0: "REQUEST_RATE",
	1: "LATENCY",
	2: "SUCCESS_RATE",
}
var MetricName_value = map[string]int32{
	"REQUEST_RATE": 0,
	"LATENCY":      1,
	"SUCCESS_RATE": 2,
}

func (x MetricName) String() string {
	return proto.EnumName(MetricName_name, int32(x))
}
func (MetricName) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TimeWindow int32

const (
	TimeWindow_TEN_SEC  TimeWindow = 0
	TimeWindow_ONE_MIN  TimeWindow = 1
	TimeWindow_TEN_MIN  TimeWindow = 2
	TimeWindow_ONE_HOUR TimeWindow = 3
)

var TimeWindow_name = map[int32]string{
	0: "TEN_SEC",
	1: "ONE_MIN",
	2: "TEN_MIN",
	3: "ONE_HOUR",
}
var TimeWindow_value = map[string]int32{
	"TEN_SEC":  0,
	"ONE_MIN":  1,
	"TEN_MIN":  2,
	"ONE_HOUR": 3,
}

func (x TimeWindow) String() string {
	return proto.EnumName(TimeWindow_name, int32(x))
}
func (TimeWindow) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AggregationType int32

const (
	AggregationType_TARGET_POD    AggregationType = 0
	AggregationType_TARGET_DEPLOY AggregationType = 1
	AggregationType_SOURCE_POD    AggregationType = 2
	AggregationType_SOURCE_DEPLOY AggregationType = 3
	AggregationType_MESH          AggregationType = 4
	AggregationType_PATH          AggregationType = 5
)

var AggregationType_name = map[int32]string{
	0: "TARGET_POD",
	1: "TARGET_DEPLOY",
	2: "SOURCE_POD",
	3: "SOURCE_DEPLOY",
	4: "MESH",
	5: "PATH",
}
var AggregationType_value = map[string]int32{
	"TARGET_POD":    0,
	"TARGET_DEPLOY": 1,
	"SOURCE_POD":    2,
	"SOURCE_DEPLOY": 3,
	"MESH":          4,
	"PATH":          5,
}

func (x AggregationType) String() string {
	return proto.EnumName(AggregationType_name, int32(x))
}
func (AggregationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type HistogramLabel int32

const (
	HistogramLabel_MIN HistogramLabel = 0
	HistogramLabel_P50 HistogramLabel = 1
	HistogramLabel_P95 HistogramLabel = 2
	HistogramLabel_P99 HistogramLabel = 3
	HistogramLabel_MAX HistogramLabel = 4
)

var HistogramLabel_name = map[int32]string{
	0: "MIN",
	1: "P50",
	2: "P95",
	3: "P99",
	4: "MAX",
}
var HistogramLabel_value = map[string]int32{
	"MIN": 0,
	"P50": 1,
	"P95": 2,
	"P99": 3,
	"MAX": 4,
}

func (x HistogramLabel) String() string {
	return proto.EnumName(HistogramLabel_name, int32(x))
}
func (HistogramLabel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type HistogramValue struct {
	Label HistogramLabel `protobuf:"varint,1,opt,name=label,enum=conduit.public.HistogramLabel" json:"label,omitempty"`
	Value int64          `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *HistogramValue) Reset()                    { *m = HistogramValue{} }
func (m *HistogramValue) String() string            { return proto.CompactTextString(m) }
func (*HistogramValue) ProtoMessage()               {}
func (*HistogramValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HistogramValue) GetLabel() HistogramLabel {
	if m != nil {
		return m.Label
	}
	return HistogramLabel_MIN
}

func (m *HistogramValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Histogram struct {
	Values []*HistogramValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *Histogram) Reset()                    { *m = Histogram{} }
func (m *Histogram) String() string            { return proto.CompactTextString(m) }
func (*Histogram) ProtoMessage()               {}
func (*Histogram) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Histogram) GetValues() []*HistogramValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type MetricValue struct {
	// Types that are valid to be assigned to Value:
	//	*MetricValue_Counter
	//	*MetricValue_Gauge
	//	*MetricValue_Histogram
	Value isMetricValue_Value `protobuf_oneof:"value"`
}

func (m *MetricValue) Reset()                    { *m = MetricValue{} }
func (m *MetricValue) String() string            { return proto.CompactTextString(m) }
func (*MetricValue) ProtoMessage()               {}
func (*MetricValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isMetricValue_Value interface {
	isMetricValue_Value()
}

type MetricValue_Counter struct {
	Counter int64 `protobuf:"varint,1,opt,name=counter,oneof"`
}
type MetricValue_Gauge struct {
	Gauge float64 `protobuf:"fixed64,2,opt,name=gauge,oneof"`
}
type MetricValue_Histogram struct {
	Histogram *Histogram `protobuf:"bytes,3,opt,name=histogram,oneof"`
}

func (*MetricValue_Counter) isMetricValue_Value()   {}
func (*MetricValue_Gauge) isMetricValue_Value()     {}
func (*MetricValue_Histogram) isMetricValue_Value() {}

func (m *MetricValue) GetValue() isMetricValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricValue) GetCounter() int64 {
	if x, ok := m.GetValue().(*MetricValue_Counter); ok {
		return x.Counter
	}
	return 0
}

func (m *MetricValue) GetGauge() float64 {
	if x, ok := m.GetValue().(*MetricValue_Gauge); ok {
		return x.Gauge
	}
	return 0
}

func (m *MetricValue) GetHistogram() *Histogram {
	if x, ok := m.GetValue().(*MetricValue_Histogram); ok {
		return x.Histogram
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetricValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetricValue_OneofMarshaler, _MetricValue_OneofUnmarshaler, _MetricValue_OneofSizer, []interface{}{
		(*MetricValue_Counter)(nil),
		(*MetricValue_Gauge)(nil),
		(*MetricValue_Histogram)(nil),
	}
}

func _MetricValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetricValue)
	// value
	switch x := m.Value.(type) {
	case *MetricValue_Counter:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Counter))
	case *MetricValue_Gauge:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Gauge))
	case *MetricValue_Histogram:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Histogram); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MetricValue.Value has unexpected type %T", x)
	}
	return nil
}

func _MetricValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetricValue)
	switch tag {
	case 1: // value.counter
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricValue_Counter{int64(x)}
		return true, err
	case 2: // value.gauge
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &MetricValue_Gauge{math.Float64frombits(x)}
		return true, err
	case 3: // value.histogram
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Histogram)
		err := b.DecodeMessage(msg)
		m.Value = &MetricValue_Histogram{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MetricValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetricValue)
	// value
	switch x := m.Value.(type) {
	case *MetricValue_Counter:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Counter))
	case *MetricValue_Gauge:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *MetricValue_Histogram:
		s := proto.Size(x.Histogram)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type MetricDatapoint struct {
	Value       *MetricValue `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64        `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *MetricDatapoint) Reset()                    { *m = MetricDatapoint{} }
func (m *MetricDatapoint) String() string            { return proto.CompactTextString(m) }
func (*MetricDatapoint) ProtoMessage()               {}
func (*MetricDatapoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetricDatapoint) GetValue() *MetricValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricDatapoint) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

type MetricSeries struct {
	Name       MetricName         `protobuf:"varint,1,opt,name=name,enum=conduit.public.MetricName" json:"name,omitempty"`
	Metadata   *MetricMetadata    `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	Datapoints []*MetricDatapoint `protobuf:"bytes,3,rep,name=datapoints" json:"datapoints,omitempty"`
}

func (m *MetricSeries) Reset()                    { *m = MetricSeries{} }
func (m *MetricSeries) String() string            { return proto.CompactTextString(m) }
func (*MetricSeries) ProtoMessage()               {}
func (*MetricSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MetricSeries) GetName() MetricName {
	if m != nil {
		return m.Name
	}
	return MetricName_REQUEST_RATE
}

func (m *MetricSeries) GetMetadata() *MetricMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MetricSeries) GetDatapoints() []*MetricDatapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

type MetricMetadata struct {
	TargetPod    string `protobuf:"bytes,1,opt,name=targetPod" json:"targetPod,omitempty"`
	TargetDeploy string `protobuf:"bytes,2,opt,name=targetDeploy" json:"targetDeploy,omitempty"`
	SourcePod    string `protobuf:"bytes,3,opt,name=sourcePod" json:"sourcePod,omitempty"`
	SourceDeploy string `protobuf:"bytes,4,opt,name=sourceDeploy" json:"sourceDeploy,omitempty"`
	Component    string `protobuf:"bytes,5,opt,name=component" json:"component,omitempty"`
	Path         string `protobuf:"bytes,6,opt,name=path" json:"path,omitempty"`
}

func (m *MetricMetadata) Reset()                    { *m = MetricMetadata{} }
func (m *MetricMetadata) String() string            { return proto.CompactTextString(m) }
func (*MetricMetadata) ProtoMessage()               {}
func (*MetricMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MetricMetadata) GetTargetPod() string {
	if m != nil {
		return m.TargetPod
	}
	return ""
}

func (m *MetricMetadata) GetTargetDeploy() string {
	if m != nil {
		return m.TargetDeploy
	}
	return ""
}

func (m *MetricMetadata) GetSourcePod() string {
	if m != nil {
		return m.SourcePod
	}
	return ""
}

func (m *MetricMetadata) GetSourceDeploy() string {
	if m != nil {
		return m.SourceDeploy
	}
	return ""
}

func (m *MetricMetadata) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *MetricMetadata) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type MetricResponse struct {
	Metrics []*MetricSeries `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *MetricResponse) Reset()                    { *m = MetricResponse{} }
func (m *MetricResponse) String() string            { return proto.CompactTextString(m) }
func (*MetricResponse) ProtoMessage()               {}
func (*MetricResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MetricResponse) GetMetrics() []*MetricSeries {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type MetricRequest struct {
	Metrics   []MetricName    `protobuf:"varint,1,rep,packed,name=metrics,enum=conduit.public.MetricName" json:"metrics,omitempty"`
	Window    TimeWindow      `protobuf:"varint,2,opt,name=window,enum=conduit.public.TimeWindow" json:"window,omitempty"`
	GroupBy   AggregationType `protobuf:"varint,3,opt,name=groupBy,enum=conduit.public.AggregationType" json:"groupBy,omitempty"`
	FilterBy  *MetricMetadata `protobuf:"bytes,4,opt,name=filterBy" json:"filterBy,omitempty"`
	Summarize bool            `protobuf:"varint,5,opt,name=summarize" json:"summarize,omitempty"`
}

func (m *MetricRequest) Reset()                    { *m = MetricRequest{} }
func (m *MetricRequest) String() string            { return proto.CompactTextString(m) }
func (*MetricRequest) ProtoMessage()               {}
func (*MetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MetricRequest) GetMetrics() []MetricName {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *MetricRequest) GetWindow() TimeWindow {
	if m != nil {
		return m.Window
	}
	return TimeWindow_TEN_SEC
}

func (m *MetricRequest) GetGroupBy() AggregationType {
	if m != nil {
		return m.GroupBy
	}
	return AggregationType_TARGET_POD
}

func (m *MetricRequest) GetFilterBy() *MetricMetadata {
	if m != nil {
		return m.FilterBy
	}
	return nil
}

func (m *MetricRequest) GetSummarize() bool {
	if m != nil {
		return m.Summarize
	}
	return false
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type VersionInfo struct {
	GoVersion      string `protobuf:"bytes,1,opt,name=goVersion" json:"goVersion,omitempty"`
	BuildDate      string `protobuf:"bytes,2,opt,name=buildDate" json:"buildDate,omitempty"`
	ReleaseVersion string `protobuf:"bytes,3,opt,name=releaseVersion" json:"releaseVersion,omitempty"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *VersionInfo) GetGoVersion() string {
	if m != nil {
		return m.GoVersion
	}
	return ""
}

func (m *VersionInfo) GetBuildDate() string {
	if m != nil {
		return m.BuildDate
	}
	return ""
}

func (m *VersionInfo) GetReleaseVersion() string {
	if m != nil {
		return m.ReleaseVersion
	}
	return ""
}

type ListPodsResponse struct {
	Pods []*Pod `protobuf:"bytes,1,rep,name=pods" json:"pods,omitempty"`
}

func (m *ListPodsResponse) Reset()                    { *m = ListPodsResponse{} }
func (m *ListPodsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPodsResponse) ProtoMessage()               {}
func (*ListPodsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListPodsResponse) GetPods() []*Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type Pod struct {
	Name                string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PodIP               string                    `protobuf:"bytes,2,opt,name=podIP" json:"podIP,omitempty"`
	Deployment          string                    `protobuf:"bytes,3,opt,name=deployment" json:"deployment,omitempty"`
	Status              string                    `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Added               bool                      `protobuf:"varint,5,opt,name=added" json:"added,omitempty"`
	SinceLastReport     *google_protobuf.Duration `protobuf:"bytes,6,opt,name=sinceLastReport" json:"sinceLastReport,omitempty"`
	ControllerNamespace string                    `protobuf:"bytes,7,opt,name=controllerNamespace" json:"controllerNamespace,omitempty"`
	ControlPlane        bool                      `protobuf:"varint,8,opt,name=controlPlane" json:"controlPlane,omitempty"`
}

func (m *Pod) Reset()                    { *m = Pod{} }
func (m *Pod) String() string            { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()               {}
func (*Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Pod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pod) GetPodIP() string {
	if m != nil {
		return m.PodIP
	}
	return ""
}

func (m *Pod) GetDeployment() string {
	if m != nil {
		return m.Deployment
	}
	return ""
}

func (m *Pod) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Pod) GetAdded() bool {
	if m != nil {
		return m.Added
	}
	return false
}

func (m *Pod) GetSinceLastReport() *google_protobuf.Duration {
	if m != nil {
		return m.SinceLastReport
	}
	return nil
}

func (m *Pod) GetControllerNamespace() string {
	if m != nil {
		return m.ControllerNamespace
	}
	return ""
}

func (m *Pod) GetControlPlane() bool {
	if m != nil {
		return m.ControlPlane
	}
	return false
}

type TapRequest struct {
	// Types that are valid to be assigned to Target:
	//	*TapRequest_Pod
	//	*TapRequest_Deployment
	Target isTapRequest_Target `protobuf_oneof:"target"`
	// validation of these fields happens on the server
	MaxRps    float32 `protobuf:"fixed32,3,opt,name=maxRps" json:"maxRps,omitempty"`
	ToPort    uint32  `protobuf:"varint,4,opt,name=toPort" json:"toPort,omitempty"`
	ToIP      string  `protobuf:"bytes,5,opt,name=toIP" json:"toIP,omitempty"`
	FromPort  uint32  `protobuf:"varint,6,opt,name=fromPort" json:"fromPort,omitempty"`
	FromIP    string  `protobuf:"bytes,7,opt,name=fromIP" json:"fromIP,omitempty"`
	Scheme    string  `protobuf:"bytes,8,opt,name=scheme" json:"scheme,omitempty"`
	Method    string  `protobuf:"bytes,9,opt,name=method" json:"method,omitempty"`
	Authority string  `protobuf:"bytes,10,opt,name=authority" json:"authority,omitempty"`
	Path      string  `protobuf:"bytes,11,opt,name=path" json:"path,omitempty"`
}

func (m *TapRequest) Reset()                    { *m = TapRequest{} }
func (m *TapRequest) String() string            { return proto.CompactTextString(m) }
func (*TapRequest) ProtoMessage()               {}
func (*TapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type isTapRequest_Target interface {
	isTapRequest_Target()
}

type TapRequest_Pod struct {
	Pod string `protobuf:"bytes,1,opt,name=pod,oneof"`
}
type TapRequest_Deployment struct {
	Deployment string `protobuf:"bytes,2,opt,name=deployment,oneof"`
}

func (*TapRequest_Pod) isTapRequest_Target()        {}
func (*TapRequest_Deployment) isTapRequest_Target() {}

func (m *TapRequest) GetTarget() isTapRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *TapRequest) GetPod() string {
	if x, ok := m.GetTarget().(*TapRequest_Pod); ok {
		return x.Pod
	}
	return ""
}

func (m *TapRequest) GetDeployment() string {
	if x, ok := m.GetTarget().(*TapRequest_Deployment); ok {
		return x.Deployment
	}
	return ""
}

func (m *TapRequest) GetMaxRps() float32 {
	if m != nil {
		return m.MaxRps
	}
	return 0
}

func (m *TapRequest) GetToPort() uint32 {
	if m != nil {
		return m.ToPort
	}
	return 0
}

func (m *TapRequest) GetToIP() string {
	if m != nil {
		return m.ToIP
	}
	return ""
}

func (m *TapRequest) GetFromPort() uint32 {
	if m != nil {
		return m.FromPort
	}
	return 0
}

func (m *TapRequest) GetFromIP() string {
	if m != nil {
		return m.FromIP
	}
	return ""
}

func (m *TapRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *TapRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *TapRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *TapRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TapRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TapRequest_OneofMarshaler, _TapRequest_OneofUnmarshaler, _TapRequest_OneofSizer, []interface{}{
		(*TapRequest_Pod)(nil),
		(*TapRequest_Deployment)(nil),
	}
}

func _TapRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TapRequest)
	// target
	switch x := m.Target.(type) {
	case *TapRequest_Pod:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Pod)
	case *TapRequest_Deployment:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Deployment)
	case nil:
	default:
		return fmt.Errorf("TapRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _TapRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TapRequest)
	switch tag {
	case 1: // target.pod
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Target = &TapRequest_Pod{x}
		return true, err
	case 2: // target.deployment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Target = &TapRequest_Deployment{x}
		return true, err
	default:
		return false, nil
	}
}

func _TapRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TapRequest)
	// target
	switch x := m.Target.(type) {
	case *TapRequest_Pod:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Pod)))
		n += len(x.Pod)
	case *TapRequest_Deployment:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Deployment)))
		n += len(x.Deployment)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ApiError struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *ApiError) Reset()                    { *m = ApiError{} }
func (m *ApiError) String() string            { return proto.CompactTextString(m) }
func (*ApiError) ProtoMessage()               {}
func (*ApiError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ApiError) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*HistogramValue)(nil), "conduit.public.HistogramValue")
	proto.RegisterType((*Histogram)(nil), "conduit.public.Histogram")
	proto.RegisterType((*MetricValue)(nil), "conduit.public.MetricValue")
	proto.RegisterType((*MetricDatapoint)(nil), "conduit.public.MetricDatapoint")
	proto.RegisterType((*MetricSeries)(nil), "conduit.public.MetricSeries")
	proto.RegisterType((*MetricMetadata)(nil), "conduit.public.MetricMetadata")
	proto.RegisterType((*MetricResponse)(nil), "conduit.public.MetricResponse")
	proto.RegisterType((*MetricRequest)(nil), "conduit.public.MetricRequest")
	proto.RegisterType((*Empty)(nil), "conduit.public.Empty")
	proto.RegisterType((*VersionInfo)(nil), "conduit.public.VersionInfo")
	proto.RegisterType((*ListPodsResponse)(nil), "conduit.public.ListPodsResponse")
	proto.RegisterType((*Pod)(nil), "conduit.public.Pod")
	proto.RegisterType((*TapRequest)(nil), "conduit.public.TapRequest")
	proto.RegisterType((*ApiError)(nil), "conduit.public.ApiError")
	proto.RegisterEnum("conduit.public.MetricName", MetricName_name, MetricName_value)
	proto.RegisterEnum("conduit.public.TimeWindow", TimeWindow_name, TimeWindow_value)
	proto.RegisterEnum("conduit.public.AggregationType", AggregationType_name, AggregationType_value)
	proto.RegisterEnum("conduit.public.HistogramLabel", HistogramLabel_name, HistogramLabel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	Stat(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*MetricResponse, error)
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionInfo, error)
	ListPods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPodsResponse, error)
	Tap(ctx context.Context, in *TapRequest, opts ...grpc.CallOption) (Api_TapClient, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Stat(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*MetricResponse, error) {
	out := new(MetricResponse)
	err := grpc.Invoke(ctx, "/conduit.public.Api/Stat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionInfo, error) {
	out := new(VersionInfo)
	err := grpc.Invoke(ctx, "/conduit.public.Api/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListPods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := grpc.Invoke(ctx, "/conduit.public.Api/ListPods", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Tap(ctx context.Context, in *TapRequest, opts ...grpc.CallOption) (Api_TapClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Api_serviceDesc.Streams[0], c.cc, "/conduit.public.Api/Tap", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiTapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Api_TapClient interface {
	Recv() (*conduit_common.TapEvent, error)
	grpc.ClientStream
}

type apiTapClient struct {
	grpc.ClientStream
}

func (x *apiTapClient) Recv() (*conduit_common.TapEvent, error) {
	m := new(conduit_common.TapEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Api service

type ApiServer interface {
	Stat(context.Context, *MetricRequest) (*MetricResponse, error)
	Version(context.Context, *Empty) (*VersionInfo, error)
	ListPods(context.Context, *Empty) (*ListPodsResponse, error)
	Tap(*TapRequest, Api_TapServer) error
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.public.Api/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Stat(ctx, req.(*MetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.public.Api/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conduit.public.Api/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListPods(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Tap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServer).Tap(m, &apiTapServer{stream})
}

type Api_TapServer interface {
	Send(*conduit_common.TapEvent) error
	grpc.ServerStream
}

type apiTapServer struct {
	grpc.ServerStream
}

func (x *apiTapServer) Send(m *conduit_common.TapEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conduit.public.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stat",
			Handler:    _Api_Stat_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Api_Version_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _Api_ListPods_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tap",
			Handler:       _Api_Tap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "public/api.proto",
}

func init() { proto.RegisterFile("public/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x15, 0x45, 0xc9, 0x92, 0xae, 0x6c, 0x85, 0x9d, 0xa4, 0x81, 0xaa, 0xa6, 0xae, 0xca, 0x45,
	0x6b, 0x78, 0x21, 0xa7, 0x6e, 0x12, 0xc0, 0x2d, 0x82, 0x40, 0x96, 0x89, 0xc8, 0x80, 0x1f, 0xec,
	0x88, 0x4e, 0x5b, 0xa0, 0x80, 0x31, 0x16, 0xc7, 0x34, 0x51, 0x92, 0xc3, 0x90, 0xc3, 0xa4, 0xee,
	0xbe, 0xdb, 0x6e, 0xdb, 0x2f, 0xe8, 0x4f, 0x74, 0xdf, 0xef, 0x2a, 0xe6, 0x41, 0x4a, 0x56, 0x94,
	0xa0, 0x2b, 0xcf, 0x3d, 0xf7, 0xcc, 0xe1, 0x7d, 0xcd, 0xb5, 0xc0, 0x4a, 0x8b, 0xab, 0x28, 0x9c,
	0xef, 0x91, 0x34, 0x1c, 0xa5, 0x19, 0xe3, 0x0c, 0xf5, 0xe6, 0x2c, 0xf1, 0x8b, 0x90, 0x8f, 0x94,
	0x67, 0xb0, 0x1d, 0x30, 0x16, 0x44, 0x74, 0x4f, 0x7a, 0xaf, 0x8a, 0xeb, 0x3d, 0xbf, 0xc8, 0x08,
	0x0f, 0x59, 0xa2, 0xf8, 0x83, 0xfb, 0x73, 0x16, 0xc7, 0x2c, 0xd9, 0x53, 0x7f, 0x14, 0x68, 0xff,
	0x0c, 0xbd, 0x69, 0x98, 0x73, 0x16, 0x64, 0x24, 0x7e, 0x45, 0xa2, 0x82, 0xa2, 0x27, 0xd0, 0x8c,
	0xc8, 0x15, 0x8d, 0xfa, 0xc6, 0xd0, 0xd8, 0xe9, 0xed, 0x6f, 0x8f, 0xee, 0x7e, 0x66, 0x54, 0xd1,
	0x4f, 0x04, 0x0b, 0x2b, 0x32, 0x7a, 0x00, 0xcd, 0x37, 0xe2, 0x7a, 0xbf, 0x3e, 0x34, 0x76, 0x4c,
	0xac, 0x0c, 0x7b, 0x02, 0x9d, 0x8a, 0x8e, 0x9e, 0xc1, 0x86, 0x44, 0xf3, 0xbe, 0x31, 0x34, 0x77,
	0xba, 0x1f, 0x50, 0x96, 0x81, 0x60, 0xcd, 0xb6, 0x7f, 0x37, 0xa0, 0x7b, 0x4a, 0x79, 0x16, 0xce,
	0x55, 0x80, 0x03, 0x68, 0xcd, 0x59, 0x91, 0x70, 0x9a, 0xc9, 0x10, 0xcd, 0x69, 0x0d, 0x97, 0x00,
	0x7a, 0x08, 0xcd, 0x80, 0x14, 0x81, 0x0a, 0xc3, 0x98, 0xd6, 0xb0, 0x32, 0xd1, 0x01, 0x74, 0x6e,
	0x4a, 0xf5, 0xbe, 0x39, 0x34, 0x76, 0xba, 0xfb, 0x9f, 0xbc, 0xf7, 0xf3, 0xd3, 0x1a, 0x5e, 0xb0,
	0x0f, 0x5b, 0x3a, 0x33, 0x3b, 0x80, 0x7b, 0x2a, 0x8c, 0x23, 0xc2, 0x49, 0xca, 0xc2, 0x84, 0xa3,
	0xaf, 0xcb, 0xac, 0x0d, 0x29, 0xf9, 0xe9, 0xaa, 0xe4, 0x52, 0xd8, 0xba, 0x24, 0xe8, 0x0b, 0xd8,
	0xe4, 0x61, 0x4c, 0x73, 0x4e, 0xe2, 0xf4, 0x32, 0xce, 0x75, 0xbd, 0xba, 0x15, 0x76, 0x9a, 0xdb,
	0xff, 0x18, 0xb0, 0xa9, 0x6e, 0xce, 0x68, 0x16, 0xd2, 0x1c, 0x8d, 0xa0, 0x91, 0x90, 0x98, 0xea,
	0x8e, 0x0c, 0xd6, 0x7f, 0xe5, 0x8c, 0xc4, 0x14, 0x4b, 0x1e, 0xfa, 0x16, 0xda, 0x31, 0xe5, 0xc4,
	0x27, 0x9c, 0x48, 0xfd, 0x35, 0xb5, 0x56, 0x77, 0x4e, 0x35, 0x0b, 0x57, 0x7c, 0xf4, 0x02, 0xc0,
	0x2f, 0xf3, 0xcb, 0xfb, 0xa6, 0xec, 0xd4, 0xe7, 0xeb, 0x6f, 0x57, 0x75, 0xc0, 0x4b, 0x57, 0xec,
	0x7f, 0x0d, 0xe8, 0xdd, 0x55, 0x47, 0x8f, 0xa0, 0xc3, 0x49, 0x16, 0x50, 0xee, 0x32, 0x5f, 0x26,
	0xd1, 0xc1, 0x0b, 0x00, 0xd9, 0xb0, 0xa9, 0x8c, 0x23, 0x9a, 0x46, 0xec, 0x56, 0x46, 0xdc, 0xc1,
	0x77, 0x30, 0xa1, 0x90, 0xb3, 0x22, 0x9b, 0x53, 0xa1, 0x60, 0x2a, 0x85, 0x0a, 0x10, 0x0a, 0xca,
	0xd0, 0x0a, 0x0d, 0xa5, 0xb0, 0x8c, 0x09, 0x85, 0x39, 0x8b, 0x53, 0x96, 0xd0, 0x84, 0xf7, 0x9b,
	0x4a, 0xa1, 0x02, 0x10, 0x82, 0x46, 0x4a, 0xf8, 0x4d, 0x7f, 0x43, 0x3a, 0xe4, 0xd9, 0x9e, 0x96,
	0x79, 0x60, 0x9a, 0xa7, 0x2c, 0xc9, 0x29, 0x7a, 0x06, 0xad, 0x58, 0x22, 0xe5, 0x08, 0x3f, 0x5a,
	0x5f, 0x18, 0xd5, 0x36, 0x5c, 0x92, 0xed, 0x3f, 0xea, 0xb0, 0x55, 0x4a, 0xbd, 0x2e, 0x68, 0xce,
	0xd1, 0x93, 0xbb, 0x4a, 0x1f, 0x6e, 0x6a, 0x49, 0x45, 0xfb, 0xb0, 0xf1, 0x36, 0x4c, 0x7c, 0xf6,
	0x56, 0xd6, 0x68, 0xcd, 0x25, 0x2f, 0x8c, 0xe9, 0x0f, 0x92, 0x81, 0x35, 0x13, 0x1d, 0x40, 0x2b,
	0xc8, 0x58, 0x91, 0x1e, 0xde, 0xca, 0xba, 0xf5, 0xde, 0x6d, 0xe6, 0x38, 0x08, 0x32, 0x1a, 0xc8,
	0x4d, 0xe1, 0xdd, 0xa6, 0x14, 0x97, 0x7c, 0x31, 0x46, 0xd7, 0x61, 0xc4, 0x69, 0x76, 0xa8, 0x4a,
	0xfa, 0x3f, 0xc6, 0xa8, 0xe4, 0xcb, 0x86, 0x15, 0x71, 0x4c, 0xb2, 0xf0, 0x37, 0x2a, 0xcb, 0xdd,
	0xc6, 0x0b, 0xc0, 0x6e, 0x41, 0xd3, 0x89, 0x53, 0x7e, 0x6b, 0xbf, 0x86, 0xee, 0x2b, 0x9a, 0xe5,
	0x21, 0x4b, 0x8e, 0x93, 0x6b, 0x26, 0x6e, 0x05, 0x4c, 0x03, 0xe5, 0xa0, 0x54, 0x80, 0xf0, 0x5e,
	0x15, 0x61, 0xe4, 0x1f, 0x11, 0x4e, 0xf5, 0x94, 0x2c, 0x00, 0xf4, 0x25, 0xf4, 0x32, 0x1a, 0x51,
	0x92, 0xd3, 0x52, 0x40, 0xcd, 0xc9, 0x0a, 0x6a, 0x7f, 0x07, 0xd6, 0x49, 0x98, 0x8b, 0xc9, 0xcb,
	0xab, 0xc6, 0x7e, 0x05, 0x8d, 0x94, 0xf9, 0x65, 0x57, 0xef, 0xaf, 0x66, 0xe9, 0x32, 0x1f, 0x4b,
	0x82, 0xfd, 0x57, 0x1d, 0x4c, 0x31, 0x71, 0x68, 0xe9, 0x45, 0x76, 0xf4, 0xab, 0x7b, 0x00, 0xcd,
	0x94, 0xf9, 0xc7, 0xae, 0x0e, 0x4d, 0x19, 0x68, 0x1b, 0xc0, 0x97, 0x13, 0x18, 0x8b, 0xc1, 0x53,
	0x21, 0x2d, 0x21, 0xe8, 0x21, 0x6c, 0xe4, 0x9c, 0xf0, 0x22, 0xd7, 0x53, 0xab, 0x2d, 0xa1, 0x46,
	0x7c, 0x9f, 0xfa, 0xba, 0x78, 0xca, 0x40, 0x13, 0xb8, 0x97, 0x87, 0xc9, 0x9c, 0x9e, 0x90, 0x9c,
	0x63, 0x9a, 0xb2, 0x8c, 0xcb, 0x91, 0x15, 0xdb, 0x4c, 0x6d, 0xff, 0x51, 0xb9, 0xfd, 0x47, 0x47,
	0x7a, 0xfb, 0xe3, 0xd5, 0x1b, 0xe8, 0x31, 0xdc, 0x9f, 0xb3, 0x84, 0x67, 0x2c, 0x8a, 0x68, 0x26,
	0x26, 0x2c, 0x4f, 0xc9, 0x9c, 0xf6, 0x5b, 0xf2, 0xfb, 0xeb, 0x5c, 0xe2, 0x81, 0x69, 0xd8, 0x8d,
	0x48, 0x42, 0xfb, 0x6d, 0x19, 0xd3, 0x1d, 0xcc, 0xfe, 0xbb, 0x0e, 0xe0, 0x91, 0xb4, 0x9c, 0x70,
	0x04, 0x66, 0x5a, 0xbe, 0xf6, 0x69, 0x0d, 0x0b, 0x03, 0x0d, 0xef, 0xd4, 0xa2, 0xae, 0x5d, 0x2b,
	0xd5, 0x88, 0xc9, 0xaf, 0x38, 0xcd, 0x65, 0xa5, 0xea, 0x58, 0x5b, 0x02, 0xe7, 0xcc, 0x15, 0xe9,
	0x8a, 0x2a, 0x6d, 0x61, 0x6d, 0x89, 0x3e, 0x70, 0x76, 0xec, 0xea, 0x07, 0x2d, 0xcf, 0x68, 0x00,
	0xed, 0xeb, 0x8c, 0xc5, 0x6e, 0x59, 0x9c, 0x2d, 0x5c, 0xd9, 0x42, 0x47, 0x9c, 0x8f, 0x5d, 0x9d,
	0xad, 0xb6, 0x64, 0x17, 0xe6, 0x37, 0x34, 0x56, 0xa9, 0x89, 0x2e, 0x48, 0x4b, 0xc6, 0x43, 0xf9,
	0x0d, 0xf3, 0xfb, 0x1d, 0x85, 0x2b, 0x4b, 0x8c, 0x22, 0x29, 0xf8, 0x0d, 0xcb, 0x42, 0x7e, 0xdb,
	0x07, 0x35, 0x8a, 0x15, 0x50, 0x6d, 0x93, 0xee, 0x62, 0x9b, 0x1c, 0xb6, 0x61, 0x43, 0x6d, 0x34,
	0x7b, 0x08, 0xed, 0x71, 0x1a, 0x3a, 0x59, 0xc6, 0x32, 0xd1, 0x65, 0x2a, 0x0e, 0x7a, 0x90, 0x94,
	0xb1, 0xfb, 0x1c, 0x60, 0xf1, 0xfc, 0x91, 0x05, 0x9b, 0xd8, 0xf9, 0xfe, 0xc2, 0x99, 0x79, 0x97,
	0x78, 0xec, 0x39, 0x56, 0x0d, 0x75, 0xa1, 0x75, 0x32, 0xf6, 0x9c, 0xb3, 0xc9, 0x4f, 0x96, 0x21,
	0xdc, 0xb3, 0x8b, 0xc9, 0xc4, 0x99, 0xcd, 0x94, 0xbb, 0xbe, 0x3b, 0x06, 0x58, 0x2c, 0x02, 0x41,
	0xf6, 0x9c, 0xb3, 0xcb, 0x99, 0x33, 0x51, 0x37, 0xcf, 0xcf, 0x9c, 0xcb, 0xd3, 0xe3, 0x33, 0xcb,
	0x28, 0x3d, 0xc2, 0xa8, 0xa3, 0x4d, 0x68, 0x0b, 0xcf, 0xf4, 0xfc, 0x02, 0x5b, 0xe6, 0xee, 0x2f,
	0x70, 0x6f, 0x65, 0x2d, 0xa0, 0x1e, 0x80, 0x37, 0xc6, 0x2f, 0x1d, 0xef, 0xd2, 0x3d, 0x3f, 0xb2,
	0x6a, 0xe8, 0x23, 0xd8, 0xd2, 0xf6, 0x91, 0xe3, 0x9e, 0x9c, 0x8b, 0x50, 0x7a, 0x00, 0xb3, 0xf3,
	0x0b, 0x3c, 0x71, 0x24, 0xa5, 0x2e, 0x28, 0xda, 0xd6, 0x14, 0x13, 0xb5, 0xa1, 0x71, 0xea, 0xcc,
	0xa6, 0x56, 0x43, 0x9c, 0xdc, 0xb1, 0x37, 0xb5, 0x9a, 0xbb, 0xcf, 0x97, 0x7e, 0x83, 0xc8, 0x1f,
	0x15, 0xa8, 0x05, 0xa6, 0x88, 0xaa, 0x26, 0x0e, 0xee, 0xd3, 0xc7, 0x96, 0x21, 0x0f, 0x07, 0x4f,
	0xad, 0xba, 0x3a, 0x1c, 0x58, 0xa6, 0xe4, 0x8c, 0x7f, 0xb4, 0x1a, 0xfb, 0x7f, 0xd6, 0xc1, 0x1c,
	0xa7, 0x21, 0x7a, 0x09, 0x8d, 0x19, 0x27, 0x1c, 0x7d, 0xb6, 0x7e, 0x49, 0xe9, 0xc1, 0x1c, 0x6c,
	0xbf, 0xcf, 0xad, 0x76, 0x81, 0x5d, 0x43, 0x2f, 0xa0, 0x55, 0xae, 0x9c, 0x8f, 0x57, 0xc9, 0x72,
	0x6d, 0x0d, 0xde, 0xf9, 0x47, 0xbf, 0xb4, 0xc4, 0xec, 0x1a, 0x72, 0xa0, 0x5d, 0xae, 0x98, 0xf7,
	0x29, 0x0c, 0x57, 0xe1, 0xd5, 0x9d, 0x24, 0xe3, 0x30, 0x3d, 0x92, 0xa2, 0x77, 0xb7, 0x7c, 0xf5,
	0xca, 0x06, 0xfd, 0xca, 0xa7, 0x7f, 0xd5, 0x79, 0x24, 0x75, 0xde, 0xd0, 0x84, 0xdb, 0xb5, 0xc7,
	0xc6, 0xd5, 0x86, 0x5c, 0x06, 0xdf, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x94, 0xcb, 0xaa, 0x8d,
	0x3c, 0x0a, 0x00, 0x00,
}
