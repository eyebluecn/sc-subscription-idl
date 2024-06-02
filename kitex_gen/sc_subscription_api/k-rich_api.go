// Code generated by Kitex v0.9.1. DO NOT EDIT.

package sc_subscription_api

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"

	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_base"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = sc_subscription_base.KitexUnusedProtection
)

func (p *SubscriptionRichPageRequest) FastRead(buf []byte) (int, error) {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubscriptionRichPageRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubscriptionRichPageRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PageNum = v

	}
	return offset, nil
}

func (p *SubscriptionRichPageRequest) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PageSize = v

	}
	return offset, nil
}

func (p *SubscriptionRichPageRequest) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.ReaderId = &v

	}
	return offset, nil
}

func (p *SubscriptionRichPageRequest) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_subscription_base.NewBase()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Base = tmp
	return offset, nil
}

// for compatibility
func (p *SubscriptionRichPageRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *SubscriptionRichPageRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "SubscriptionRichPageRequest")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *SubscriptionRichPageRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("SubscriptionRichPageRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *SubscriptionRichPageRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "pageNum", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.PageNum)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *SubscriptionRichPageRequest) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "pageSize", thrift.I64, 2)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.PageSize)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *SubscriptionRichPageRequest) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetReaderId() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "readerId", thrift.I64, 3)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.ReaderId)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *SubscriptionRichPageRequest) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetBase() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "base", thrift.STRUCT, 255)
		offset += p.Base.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *SubscriptionRichPageRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("pageNum", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.PageNum)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *SubscriptionRichPageRequest) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("pageSize", thrift.I64, 2)
	l += bthrift.Binary.I64Length(p.PageSize)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *SubscriptionRichPageRequest) field3Length() int {
	l := 0
	if p.IsSetReaderId() {
		l += bthrift.Binary.FieldBeginLength("readerId", thrift.I64, 3)
		l += bthrift.Binary.I64Length(*p.ReaderId)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *SubscriptionRichPageRequest) field255Length() int {
	l := 0
	if p.IsSetBase() {
		l += bthrift.Binary.FieldBeginLength("base", thrift.STRUCT, 255)
		l += p.Base.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *SubscriptionRichPageResponse) FastRead(buf []byte) (int, error) {
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
			if fieldTypeId == thrift.LIST {
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
			if fieldTypeId == thrift.STRUCT {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubscriptionRichPageResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubscriptionRichPageResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Data = make([]*RichSubscriptionDTO, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewRichSubscriptionDTO()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.Data = append(p.Data, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

func (p *SubscriptionRichPageResponse) FastReadField2(buf []byte) (int, error) {
	offset := 0

	tmp := sc_subscription_base.NewPagination()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Pagination = tmp
	return offset, nil
}

func (p *SubscriptionRichPageResponse) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_subscription_base.NewBaseResp()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.BaseResp = tmp
	return offset, nil
}

// for compatibility
func (p *SubscriptionRichPageResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *SubscriptionRichPageResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "SubscriptionRichPageResponse")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *SubscriptionRichPageResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("SubscriptionRichPageResponse")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *SubscriptionRichPageResponse) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "data", thrift.LIST, 1)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
	var length int
	for _, v := range p.Data {
		length++
		offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *SubscriptionRichPageResponse) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "pagination", thrift.STRUCT, 2)
	offset += p.Pagination.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *SubscriptionRichPageResponse) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "baseResp", thrift.STRUCT, 255)
	offset += p.BaseResp.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *SubscriptionRichPageResponse) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("data", thrift.LIST, 1)
	l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.Data))
	for _, v := range p.Data {
		l += v.BLength()
	}
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *SubscriptionRichPageResponse) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("pagination", thrift.STRUCT, 2)
	l += p.Pagination.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *SubscriptionRichPageResponse) field255Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("baseResp", thrift.STRUCT, 255)
	l += p.BaseResp.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}
