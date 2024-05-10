// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package subscription

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type SubscriptionService interface {
	SubscriptionList(ctx context.Context, request *SubscriptionListRequest) (r *SubscriptionListResponse, err error)
}

type SubscriptionServiceClient struct {
	c thrift.TClient
}

func NewSubscriptionServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SubscriptionServiceClient {
	return &SubscriptionServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewSubscriptionServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SubscriptionServiceClient {
	return &SubscriptionServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewSubscriptionServiceClient(c thrift.TClient) *SubscriptionServiceClient {
	return &SubscriptionServiceClient{
		c: c,
	}
}

func (p *SubscriptionServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *SubscriptionServiceClient) SubscriptionList(ctx context.Context, request *SubscriptionListRequest) (r *SubscriptionListResponse, err error) {
	var _args SubscriptionServiceSubscriptionListArgs
	_args.Request = request
	var _result SubscriptionServiceSubscriptionListResult
	if err = p.Client_().Call(ctx, "SubscriptionList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type SubscriptionServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SubscriptionService
}

func (p *SubscriptionServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SubscriptionServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SubscriptionServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSubscriptionServiceProcessor(handler SubscriptionService) *SubscriptionServiceProcessor {
	self := &SubscriptionServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("SubscriptionList", &subscriptionServiceProcessorSubscriptionList{handler: handler})
	return self
}
func (p *SubscriptionServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type subscriptionServiceProcessorSubscriptionList struct {
	handler SubscriptionService
}

func (p *subscriptionServiceProcessorSubscriptionList) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SubscriptionServiceSubscriptionListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("SubscriptionList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := SubscriptionServiceSubscriptionListResult{}
	var retval *SubscriptionListResponse
	if retval, err2 = p.handler.SubscriptionList(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing SubscriptionList: "+err2.Error())
		oprot.WriteMessageBegin("SubscriptionList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("SubscriptionList", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type SubscriptionServiceSubscriptionListArgs struct {
	Request *SubscriptionListRequest `thrift:"request,1" frugal:"1,default,SubscriptionListRequest" json:"request"`
}

func NewSubscriptionServiceSubscriptionListArgs() *SubscriptionServiceSubscriptionListArgs {
	return &SubscriptionServiceSubscriptionListArgs{}
}

func (p *SubscriptionServiceSubscriptionListArgs) InitDefault() {
	*p = SubscriptionServiceSubscriptionListArgs{}
}

var SubscriptionServiceSubscriptionListArgs_Request_DEFAULT *SubscriptionListRequest

func (p *SubscriptionServiceSubscriptionListArgs) GetRequest() (v *SubscriptionListRequest) {
	if !p.IsSetRequest() {
		return SubscriptionServiceSubscriptionListArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *SubscriptionServiceSubscriptionListArgs) SetRequest(val *SubscriptionListRequest) {
	p.Request = val
}

var fieldIDToName_SubscriptionServiceSubscriptionListArgs = map[int16]string{
	1: "request",
}

func (p *SubscriptionServiceSubscriptionListArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *SubscriptionServiceSubscriptionListArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubscriptionServiceSubscriptionListArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubscriptionServiceSubscriptionListArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewSubscriptionListRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Request = _field
	return nil
}

func (p *SubscriptionServiceSubscriptionListArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SubscriptionList_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SubscriptionServiceSubscriptionListArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SubscriptionServiceSubscriptionListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubscriptionServiceSubscriptionListArgs(%+v)", *p)

}

func (p *SubscriptionServiceSubscriptionListArgs) DeepEqual(ano *SubscriptionServiceSubscriptionListArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *SubscriptionServiceSubscriptionListArgs) Field1DeepEqual(src *SubscriptionListRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type SubscriptionServiceSubscriptionListResult struct {
	Success *SubscriptionListResponse `thrift:"success,0,optional" frugal:"0,optional,SubscriptionListResponse" json:"success,omitempty"`
}

func NewSubscriptionServiceSubscriptionListResult() *SubscriptionServiceSubscriptionListResult {
	return &SubscriptionServiceSubscriptionListResult{}
}

func (p *SubscriptionServiceSubscriptionListResult) InitDefault() {
	*p = SubscriptionServiceSubscriptionListResult{}
}

var SubscriptionServiceSubscriptionListResult_Success_DEFAULT *SubscriptionListResponse

func (p *SubscriptionServiceSubscriptionListResult) GetSuccess() (v *SubscriptionListResponse) {
	if !p.IsSetSuccess() {
		return SubscriptionServiceSubscriptionListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SubscriptionServiceSubscriptionListResult) SetSuccess(x interface{}) {
	p.Success = x.(*SubscriptionListResponse)
}

var fieldIDToName_SubscriptionServiceSubscriptionListResult = map[int16]string{
	0: "success",
}

func (p *SubscriptionServiceSubscriptionListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SubscriptionServiceSubscriptionListResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubscriptionServiceSubscriptionListResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubscriptionServiceSubscriptionListResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewSubscriptionListResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *SubscriptionServiceSubscriptionListResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SubscriptionList_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SubscriptionServiceSubscriptionListResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *SubscriptionServiceSubscriptionListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubscriptionServiceSubscriptionListResult(%+v)", *p)

}

func (p *SubscriptionServiceSubscriptionListResult) DeepEqual(ano *SubscriptionServiceSubscriptionListResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *SubscriptionServiceSubscriptionListResult) Field0DeepEqual(src *SubscriptionListResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}