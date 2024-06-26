// Code generated by Kitex v0.9.1. DO NOT EDIT.

package sc_subscription_api

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *OrderDTO) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 9:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField9(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 10:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField10(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_OrderDTO[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *OrderDTO) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Id = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.CreateTime = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.UpdateTime = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.No = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ReaderId = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ColumnId = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ColumnQuoteId = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField8(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PaymentId = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField9(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Price = v

	}
	return offset, nil
}

func (p *OrderDTO) FastReadField10(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Status = OrderStatus(v)

	}
	return offset, nil
}

// for compatibility
func (p *OrderDTO) FastWrite(buf []byte) int {
	return 0
}

func (p *OrderDTO) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "OrderDTO")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
		offset += p.fastWriteField8(buf[offset:], binaryWriter)
		offset += p.fastWriteField9(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField10(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("OrderDTO")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
		l += p.field9Length()
		l += p.field10Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *OrderDTO) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Id)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "createTime", thrift.I64, 2)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.CreateTime)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "updateTime", thrift.I64, 3)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.UpdateTime)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "no", thrift.STRING, 4)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.No)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "readerId", thrift.I64, 5)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.ReaderId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "columnId", thrift.I64, 6)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.ColumnId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "columnQuoteId", thrift.I64, 7)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.ColumnQuoteId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField8(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "paymentId", thrift.I64, 8)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.PaymentId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField9(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "price", thrift.I64, 9)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Price)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) fastWriteField10(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status", thrift.I32, 10)
	offset += bthrift.Binary.WriteI32(buf[offset:], int32(p.Status))

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *OrderDTO) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.Id)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("createTime", thrift.I64, 2)
	l += bthrift.Binary.I64Length(p.CreateTime)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("updateTime", thrift.I64, 3)
	l += bthrift.Binary.I64Length(p.UpdateTime)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("no", thrift.STRING, 4)
	l += bthrift.Binary.StringLengthNocopy(p.No)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("readerId", thrift.I64, 5)
	l += bthrift.Binary.I64Length(p.ReaderId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("columnId", thrift.I64, 6)
	l += bthrift.Binary.I64Length(p.ColumnId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field7Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("columnQuoteId", thrift.I64, 7)
	l += bthrift.Binary.I64Length(p.ColumnQuoteId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field8Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("paymentId", thrift.I64, 8)
	l += bthrift.Binary.I64Length(p.PaymentId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field9Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("price", thrift.I64, 9)
	l += bthrift.Binary.I64Length(p.Price)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *OrderDTO) field10Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("status", thrift.I32, 10)
	l += bthrift.Binary.I32Length(int32(p.Status))

	l += bthrift.Binary.FieldEndLength()
	return l
}
