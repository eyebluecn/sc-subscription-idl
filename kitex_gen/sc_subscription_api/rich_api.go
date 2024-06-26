// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package sc_subscription_api

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_base"
)

type SubscriptionRichPageRequest struct {
	PageNum  int64                      `thrift:"pageNum,1" frugal:"1,default,i64" json:"pageNum"`
	PageSize int64                      `thrift:"pageSize,2" frugal:"2,default,i64" json:"pageSize"`
	ReaderId *int64                     `thrift:"readerId,3,optional" frugal:"3,optional,i64" json:"readerId,omitempty"`
	Base     *sc_subscription_base.Base `thrift:"base,255,optional" frugal:"255,optional,sc_subscription_base.Base" json:"base,omitempty"`
}

func NewSubscriptionRichPageRequest() *SubscriptionRichPageRequest {
	return &SubscriptionRichPageRequest{}
}

func (p *SubscriptionRichPageRequest) InitDefault() {
	*p = SubscriptionRichPageRequest{}
}

func (p *SubscriptionRichPageRequest) GetPageNum() (v int64) {
	return p.PageNum
}

func (p *SubscriptionRichPageRequest) GetPageSize() (v int64) {
	return p.PageSize
}

var SubscriptionRichPageRequest_ReaderId_DEFAULT int64

func (p *SubscriptionRichPageRequest) GetReaderId() (v int64) {
	if !p.IsSetReaderId() {
		return SubscriptionRichPageRequest_ReaderId_DEFAULT
	}
	return *p.ReaderId
}

var SubscriptionRichPageRequest_Base_DEFAULT *sc_subscription_base.Base

func (p *SubscriptionRichPageRequest) GetBase() (v *sc_subscription_base.Base) {
	if !p.IsSetBase() {
		return SubscriptionRichPageRequest_Base_DEFAULT
	}
	return p.Base
}
func (p *SubscriptionRichPageRequest) SetPageNum(val int64) {
	p.PageNum = val
}
func (p *SubscriptionRichPageRequest) SetPageSize(val int64) {
	p.PageSize = val
}
func (p *SubscriptionRichPageRequest) SetReaderId(val *int64) {
	p.ReaderId = val
}
func (p *SubscriptionRichPageRequest) SetBase(val *sc_subscription_base.Base) {
	p.Base = val
}

var fieldIDToName_SubscriptionRichPageRequest = map[int16]string{
	1:   "pageNum",
	2:   "pageSize",
	3:   "readerId",
	255: "base",
}

func (p *SubscriptionRichPageRequest) IsSetReaderId() bool {
	return p.ReaderId != nil
}

func (p *SubscriptionRichPageRequest) IsSetBase() bool {
	return p.Base != nil
}

func (p *SubscriptionRichPageRequest) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubscriptionRichPageRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubscriptionRichPageRequest) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.PageNum = _field
	return nil
}
func (p *SubscriptionRichPageRequest) ReadField2(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.PageSize = _field
	return nil
}
func (p *SubscriptionRichPageRequest) ReadField3(iprot thrift.TProtocol) error {

	var _field *int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.ReaderId = _field
	return nil
}
func (p *SubscriptionRichPageRequest) ReadField255(iprot thrift.TProtocol) error {
	_field := sc_subscription_base.NewBase()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Base = _field
	return nil
}

func (p *SubscriptionRichPageRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SubscriptionRichPageRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
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

func (p *SubscriptionRichPageRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("pageNum", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.PageNum); err != nil {
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

func (p *SubscriptionRichPageRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("pageSize", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.PageSize); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *SubscriptionRichPageRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetReaderId() {
		if err = oprot.WriteFieldBegin("readerId", thrift.I64, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI64(*p.ReaderId); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *SubscriptionRichPageRequest) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *SubscriptionRichPageRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubscriptionRichPageRequest(%+v)", *p)

}

func (p *SubscriptionRichPageRequest) DeepEqual(ano *SubscriptionRichPageRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.PageNum) {
		return false
	}
	if !p.Field2DeepEqual(ano.PageSize) {
		return false
	}
	if !p.Field3DeepEqual(ano.ReaderId) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *SubscriptionRichPageRequest) Field1DeepEqual(src int64) bool {

	if p.PageNum != src {
		return false
	}
	return true
}
func (p *SubscriptionRichPageRequest) Field2DeepEqual(src int64) bool {

	if p.PageSize != src {
		return false
	}
	return true
}
func (p *SubscriptionRichPageRequest) Field3DeepEqual(src *int64) bool {

	if p.ReaderId == src {
		return true
	} else if p.ReaderId == nil || src == nil {
		return false
	}
	if *p.ReaderId != *src {
		return false
	}
	return true
}
func (p *SubscriptionRichPageRequest) Field255DeepEqual(src *sc_subscription_base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type SubscriptionRichPageResponse struct {
	Data       []*RichSubscriptionDTO           `thrift:"data,1" frugal:"1,default,list<RichSubscriptionDTO>" json:"data"`
	Pagination *sc_subscription_base.Pagination `thrift:"pagination,2" frugal:"2,default,sc_subscription_base.Pagination" json:"pagination"`
	BaseResp   *sc_subscription_base.BaseResp   `thrift:"baseResp,255" frugal:"255,default,sc_subscription_base.BaseResp" json:"baseResp"`
}

func NewSubscriptionRichPageResponse() *SubscriptionRichPageResponse {
	return &SubscriptionRichPageResponse{}
}

func (p *SubscriptionRichPageResponse) InitDefault() {
	*p = SubscriptionRichPageResponse{}
}

func (p *SubscriptionRichPageResponse) GetData() (v []*RichSubscriptionDTO) {
	return p.Data
}

var SubscriptionRichPageResponse_Pagination_DEFAULT *sc_subscription_base.Pagination

func (p *SubscriptionRichPageResponse) GetPagination() (v *sc_subscription_base.Pagination) {
	if !p.IsSetPagination() {
		return SubscriptionRichPageResponse_Pagination_DEFAULT
	}
	return p.Pagination
}

var SubscriptionRichPageResponse_BaseResp_DEFAULT *sc_subscription_base.BaseResp

func (p *SubscriptionRichPageResponse) GetBaseResp() (v *sc_subscription_base.BaseResp) {
	if !p.IsSetBaseResp() {
		return SubscriptionRichPageResponse_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *SubscriptionRichPageResponse) SetData(val []*RichSubscriptionDTO) {
	p.Data = val
}
func (p *SubscriptionRichPageResponse) SetPagination(val *sc_subscription_base.Pagination) {
	p.Pagination = val
}
func (p *SubscriptionRichPageResponse) SetBaseResp(val *sc_subscription_base.BaseResp) {
	p.BaseResp = val
}

var fieldIDToName_SubscriptionRichPageResponse = map[int16]string{
	1:   "data",
	2:   "pagination",
	255: "baseResp",
}

func (p *SubscriptionRichPageResponse) IsSetPagination() bool {
	return p.Pagination != nil
}

func (p *SubscriptionRichPageResponse) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *SubscriptionRichPageResponse) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubscriptionRichPageResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubscriptionRichPageResponse) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	_field := make([]*RichSubscriptionDTO, 0, size)
	values := make([]RichSubscriptionDTO, size)
	for i := 0; i < size; i++ {
		_elem := &values[i]

		if err := _elem.Read(iprot); err != nil {
			return err
		}

		_field = append(_field, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	p.Data = _field
	return nil
}
func (p *SubscriptionRichPageResponse) ReadField2(iprot thrift.TProtocol) error {
	_field := sc_subscription_base.NewPagination()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Pagination = _field
	return nil
}
func (p *SubscriptionRichPageResponse) ReadField255(iprot thrift.TProtocol) error {
	_field := sc_subscription_base.NewBaseResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.BaseResp = _field
	return nil
}

func (p *SubscriptionRichPageResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SubscriptionRichPageResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
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

func (p *SubscriptionRichPageResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.LIST, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Data)); err != nil {
		return err
	}
	for _, v := range p.Data {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
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

func (p *SubscriptionRichPageResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("pagination", thrift.STRUCT, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Pagination.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *SubscriptionRichPageResponse) writeField255(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("baseResp", thrift.STRUCT, 255); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.BaseResp.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *SubscriptionRichPageResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubscriptionRichPageResponse(%+v)", *p)

}

func (p *SubscriptionRichPageResponse) DeepEqual(ano *SubscriptionRichPageResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Data) {
		return false
	}
	if !p.Field2DeepEqual(ano.Pagination) {
		return false
	}
	if !p.Field255DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *SubscriptionRichPageResponse) Field1DeepEqual(src []*RichSubscriptionDTO) bool {

	if len(p.Data) != len(src) {
		return false
	}
	for i, v := range p.Data {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}
func (p *SubscriptionRichPageResponse) Field2DeepEqual(src *sc_subscription_base.Pagination) bool {

	if !p.Pagination.DeepEqual(src) {
		return false
	}
	return true
}
func (p *SubscriptionRichPageResponse) Field255DeepEqual(src *sc_subscription_base.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
