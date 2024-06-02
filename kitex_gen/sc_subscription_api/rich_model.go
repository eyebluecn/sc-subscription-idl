// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package sc_subscription_api

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type RichSubscriptionDTO struct {
	Subscription *SubscriptionDTO `thrift:"subscription,1" frugal:"1,default,SubscriptionDTO" json:"subscription"`
	Column       *ColumnDTO       `thrift:"column,2" frugal:"2,default,ColumnDTO" json:"column"`
	Reader       *ReaderDTO       `thrift:"reader,3" frugal:"3,default,ReaderDTO" json:"reader"`
	Order        *OrderDTO        `thrift:"order,4" frugal:"4,default,OrderDTO" json:"order"`
}

func NewRichSubscriptionDTO() *RichSubscriptionDTO {
	return &RichSubscriptionDTO{}
}

func (p *RichSubscriptionDTO) InitDefault() {
	*p = RichSubscriptionDTO{}
}

var RichSubscriptionDTO_Subscription_DEFAULT *SubscriptionDTO

func (p *RichSubscriptionDTO) GetSubscription() (v *SubscriptionDTO) {
	if !p.IsSetSubscription() {
		return RichSubscriptionDTO_Subscription_DEFAULT
	}
	return p.Subscription
}

var RichSubscriptionDTO_Column_DEFAULT *ColumnDTO

func (p *RichSubscriptionDTO) GetColumn() (v *ColumnDTO) {
	if !p.IsSetColumn() {
		return RichSubscriptionDTO_Column_DEFAULT
	}
	return p.Column
}

var RichSubscriptionDTO_Reader_DEFAULT *ReaderDTO

func (p *RichSubscriptionDTO) GetReader() (v *ReaderDTO) {
	if !p.IsSetReader() {
		return RichSubscriptionDTO_Reader_DEFAULT
	}
	return p.Reader
}

var RichSubscriptionDTO_Order_DEFAULT *OrderDTO

func (p *RichSubscriptionDTO) GetOrder() (v *OrderDTO) {
	if !p.IsSetOrder() {
		return RichSubscriptionDTO_Order_DEFAULT
	}
	return p.Order
}
func (p *RichSubscriptionDTO) SetSubscription(val *SubscriptionDTO) {
	p.Subscription = val
}
func (p *RichSubscriptionDTO) SetColumn(val *ColumnDTO) {
	p.Column = val
}
func (p *RichSubscriptionDTO) SetReader(val *ReaderDTO) {
	p.Reader = val
}
func (p *RichSubscriptionDTO) SetOrder(val *OrderDTO) {
	p.Order = val
}

var fieldIDToName_RichSubscriptionDTO = map[int16]string{
	1: "subscription",
	2: "column",
	3: "reader",
	4: "order",
}

func (p *RichSubscriptionDTO) IsSetSubscription() bool {
	return p.Subscription != nil
}

func (p *RichSubscriptionDTO) IsSetColumn() bool {
	return p.Column != nil
}

func (p *RichSubscriptionDTO) IsSetReader() bool {
	return p.Reader != nil
}

func (p *RichSubscriptionDTO) IsSetOrder() bool {
	return p.Order != nil
}

func (p *RichSubscriptionDTO) Read(iprot thrift.TProtocol) (err error) {

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
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField4(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_RichSubscriptionDTO[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *RichSubscriptionDTO) ReadField1(iprot thrift.TProtocol) error {
	_field := NewSubscriptionDTO()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Subscription = _field
	return nil
}
func (p *RichSubscriptionDTO) ReadField2(iprot thrift.TProtocol) error {
	_field := NewColumnDTO()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Column = _field
	return nil
}
func (p *RichSubscriptionDTO) ReadField3(iprot thrift.TProtocol) error {
	_field := NewReaderDTO()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Reader = _field
	return nil
}
func (p *RichSubscriptionDTO) ReadField4(iprot thrift.TProtocol) error {
	_field := NewOrderDTO()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Order = _field
	return nil
}

func (p *RichSubscriptionDTO) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("RichSubscriptionDTO"); err != nil {
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
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
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

func (p *RichSubscriptionDTO) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("subscription", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Subscription.Write(oprot); err != nil {
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

func (p *RichSubscriptionDTO) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("column", thrift.STRUCT, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Column.Write(oprot); err != nil {
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

func (p *RichSubscriptionDTO) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("reader", thrift.STRUCT, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Reader.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *RichSubscriptionDTO) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("order", thrift.STRUCT, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Order.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *RichSubscriptionDTO) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RichSubscriptionDTO(%+v)", *p)

}

func (p *RichSubscriptionDTO) DeepEqual(ano *RichSubscriptionDTO) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Subscription) {
		return false
	}
	if !p.Field2DeepEqual(ano.Column) {
		return false
	}
	if !p.Field3DeepEqual(ano.Reader) {
		return false
	}
	if !p.Field4DeepEqual(ano.Order) {
		return false
	}
	return true
}

func (p *RichSubscriptionDTO) Field1DeepEqual(src *SubscriptionDTO) bool {

	if !p.Subscription.DeepEqual(src) {
		return false
	}
	return true
}
func (p *RichSubscriptionDTO) Field2DeepEqual(src *ColumnDTO) bool {

	if !p.Column.DeepEqual(src) {
		return false
	}
	return true
}
func (p *RichSubscriptionDTO) Field3DeepEqual(src *ReaderDTO) bool {

	if !p.Reader.DeepEqual(src) {
		return false
	}
	return true
}
func (p *RichSubscriptionDTO) Field4DeepEqual(src *OrderDTO) bool {

	if !p.Order.DeepEqual(src) {
		return false
	}
	return true
}
