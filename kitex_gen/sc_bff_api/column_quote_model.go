// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package sc_bff_api

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type ColumnQuoteDTO struct {
	Id         int64             `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	CreateTime int64             `thrift:"createTime,2" frugal:"2,default,i64" json:"createTime"`
	UpdateTime int64             `thrift:"updateTime,3" frugal:"3,default,i64" json:"updateTime"`
	ColumnId   int64             `thrift:"columnId,4" frugal:"4,default,i64" json:"columnId"`
	Editor     int64             `thrift:"editor,5" frugal:"5,default,i64" json:"editor"`
	Price      int64             `thrift:"price,6" frugal:"6,default,i64" json:"price"`
	Status     ColumnQuoteStatus `thrift:"status,7" frugal:"7,default,ColumnQuoteStatus" json:"status"`
}

func NewColumnQuoteDTO() *ColumnQuoteDTO {
	return &ColumnQuoteDTO{}
}

func (p *ColumnQuoteDTO) InitDefault() {
	*p = ColumnQuoteDTO{}
}

func (p *ColumnQuoteDTO) GetId() (v int64) {
	return p.Id
}

func (p *ColumnQuoteDTO) GetCreateTime() (v int64) {
	return p.CreateTime
}

func (p *ColumnQuoteDTO) GetUpdateTime() (v int64) {
	return p.UpdateTime
}

func (p *ColumnQuoteDTO) GetColumnId() (v int64) {
	return p.ColumnId
}

func (p *ColumnQuoteDTO) GetEditor() (v int64) {
	return p.Editor
}

func (p *ColumnQuoteDTO) GetPrice() (v int64) {
	return p.Price
}

func (p *ColumnQuoteDTO) GetStatus() (v ColumnQuoteStatus) {
	return p.Status
}
func (p *ColumnQuoteDTO) SetId(val int64) {
	p.Id = val
}
func (p *ColumnQuoteDTO) SetCreateTime(val int64) {
	p.CreateTime = val
}
func (p *ColumnQuoteDTO) SetUpdateTime(val int64) {
	p.UpdateTime = val
}
func (p *ColumnQuoteDTO) SetColumnId(val int64) {
	p.ColumnId = val
}
func (p *ColumnQuoteDTO) SetEditor(val int64) {
	p.Editor = val
}
func (p *ColumnQuoteDTO) SetPrice(val int64) {
	p.Price = val
}
func (p *ColumnQuoteDTO) SetStatus(val ColumnQuoteStatus) {
	p.Status = val
}

var fieldIDToName_ColumnQuoteDTO = map[int16]string{
	1: "id",
	2: "createTime",
	3: "updateTime",
	4: "columnId",
	5: "editor",
	6: "price",
	7: "status",
}

func (p *ColumnQuoteDTO) Read(iprot thrift.TProtocol) (err error) {

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
		case 4:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 7:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField7(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnQuoteDTO[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnQuoteDTO) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *ColumnQuoteDTO) ReadField2(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.CreateTime = _field
	return nil
}
func (p *ColumnQuoteDTO) ReadField3(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UpdateTime = _field
	return nil
}
func (p *ColumnQuoteDTO) ReadField4(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ColumnId = _field
	return nil
}
func (p *ColumnQuoteDTO) ReadField5(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Editor = _field
	return nil
}
func (p *ColumnQuoteDTO) ReadField6(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Price = _field
	return nil
}
func (p *ColumnQuoteDTO) ReadField7(iprot thrift.TProtocol) error {

	var _field ColumnQuoteStatus
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		_field = ColumnQuoteStatus(v)
	}
	p.Status = _field
	return nil
}

func (p *ColumnQuoteDTO) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ColumnQuoteDTO"); err != nil {
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
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
		if err = p.writeField7(oprot); err != nil {
			fieldId = 7
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

func (p *ColumnQuoteDTO) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
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

func (p *ColumnQuoteDTO) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("createTime", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CreateTime); err != nil {
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

func (p *ColumnQuoteDTO) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("updateTime", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UpdateTime); err != nil {
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

func (p *ColumnQuoteDTO) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("columnId", thrift.I64, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.ColumnId); err != nil {
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

func (p *ColumnQuoteDTO) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("editor", thrift.I64, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Editor); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *ColumnQuoteDTO) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("price", thrift.I64, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Price); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *ColumnQuoteDTO) writeField7(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status", thrift.I32, 7); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 end error: ", p), err)
}

func (p *ColumnQuoteDTO) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ColumnQuoteDTO(%+v)", *p)

}

func (p *ColumnQuoteDTO) DeepEqual(ano *ColumnQuoteDTO) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.CreateTime) {
		return false
	}
	if !p.Field3DeepEqual(ano.UpdateTime) {
		return false
	}
	if !p.Field4DeepEqual(ano.ColumnId) {
		return false
	}
	if !p.Field5DeepEqual(ano.Editor) {
		return false
	}
	if !p.Field6DeepEqual(ano.Price) {
		return false
	}
	if !p.Field7DeepEqual(ano.Status) {
		return false
	}
	return true
}

func (p *ColumnQuoteDTO) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *ColumnQuoteDTO) Field2DeepEqual(src int64) bool {

	if p.CreateTime != src {
		return false
	}
	return true
}
func (p *ColumnQuoteDTO) Field3DeepEqual(src int64) bool {

	if p.UpdateTime != src {
		return false
	}
	return true
}
func (p *ColumnQuoteDTO) Field4DeepEqual(src int64) bool {

	if p.ColumnId != src {
		return false
	}
	return true
}
func (p *ColumnQuoteDTO) Field5DeepEqual(src int64) bool {

	if p.Editor != src {
		return false
	}
	return true
}
func (p *ColumnQuoteDTO) Field6DeepEqual(src int64) bool {

	if p.Price != src {
		return false
	}
	return true
}
func (p *ColumnQuoteDTO) Field7DeepEqual(src ColumnQuoteStatus) bool {

	if p.Status != src {
		return false
	}
	return true
}
