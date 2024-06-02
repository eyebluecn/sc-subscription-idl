// Code generated by Kitex v0.9.1. DO NOT EDIT.

package sc_subscription_api

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"

	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_bff_api"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = sc_bff_api.KitexUnusedProtection
)

func (p *ColumnDTO) FastRead(buf []byte) (int, error) {
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
			if fieldTypeId == thrift.I32 {
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnDTO[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnDTO) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Id = v

	}
	return offset, nil
}

func (p *ColumnDTO) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.CreateTime = v

	}
	return offset, nil
}

func (p *ColumnDTO) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.UpdateTime = v

	}
	return offset, nil
}

func (p *ColumnDTO) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Name = v

	}
	return offset, nil
}

func (p *ColumnDTO) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.AuthorId = v

	}
	return offset, nil
}

func (p *ColumnDTO) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Status = ColumnStatus(v)

	}
	return offset, nil
}

// for compatibility
func (p *ColumnDTO) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnDTO) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnDTO")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnDTO) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnDTO")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnDTO) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Id)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDTO) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "createTime", thrift.I64, 2)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.CreateTime)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDTO) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "updateTime", thrift.I64, 3)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.UpdateTime)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDTO) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "name", thrift.STRING, 4)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Name)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDTO) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "authorId", thrift.I64, 5)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.AuthorId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDTO) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status", thrift.I32, 6)
	offset += bthrift.Binary.WriteI32(buf[offset:], int32(p.Status))

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDTO) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.Id)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDTO) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("createTime", thrift.I64, 2)
	l += bthrift.Binary.I64Length(p.CreateTime)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDTO) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("updateTime", thrift.I64, 3)
	l += bthrift.Binary.I64Length(p.UpdateTime)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDTO) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("name", thrift.STRING, 4)
	l += bthrift.Binary.StringLengthNocopy(p.Name)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDTO) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("authorId", thrift.I64, 5)
	l += bthrift.Binary.I64Length(p.AuthorId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDTO) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("status", thrift.I32, 6)
	l += bthrift.Binary.I32Length(int32(p.Status))

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *RichColumnDTO) FastRead(buf []byte) (int, error) {
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
			if fieldTypeId == thrift.STRUCT {
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
		case 3:
			if fieldTypeId == thrift.STRUCT {
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
			if fieldTypeId == thrift.STRUCT {
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_RichColumnDTO[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *RichColumnDTO) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewColumnDTO()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Column = tmp
	return offset, nil
}

func (p *RichColumnDTO) FastReadField2(buf []byte) (int, error) {
	offset := 0

	tmp := sc_bff_api.NewAuthorDTO()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Author = tmp
	return offset, nil
}

func (p *RichColumnDTO) FastReadField3(buf []byte) (int, error) {
	offset := 0

	tmp := sc_bff_api.NewColumnQuoteDTO()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.ColumnQuote = tmp
	return offset, nil
}

func (p *RichColumnDTO) FastReadField4(buf []byte) (int, error) {
	offset := 0

	tmp := NewSubscriptionDTO()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Subscription = tmp
	return offset, nil
}

// for compatibility
func (p *RichColumnDTO) FastWrite(buf []byte) int {
	return 0
}

func (p *RichColumnDTO) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "RichColumnDTO")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *RichColumnDTO) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("RichColumnDTO")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *RichColumnDTO) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "column", thrift.STRUCT, 1)
	offset += p.Column.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *RichColumnDTO) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "author", thrift.STRUCT, 2)
	offset += p.Author.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *RichColumnDTO) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "columnQuote", thrift.STRUCT, 3)
	offset += p.ColumnQuote.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *RichColumnDTO) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "subscription", thrift.STRUCT, 4)
	offset += p.Subscription.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *RichColumnDTO) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("column", thrift.STRUCT, 1)
	l += p.Column.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *RichColumnDTO) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("author", thrift.STRUCT, 2)
	l += p.Author.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *RichColumnDTO) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("columnQuote", thrift.STRUCT, 3)
	l += p.ColumnQuote.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *RichColumnDTO) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("subscription", thrift.STRUCT, 4)
	l += p.Subscription.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}
