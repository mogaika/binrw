package binrw

import (
	"encoding/binary"
	"io"
	"math"
	"os"
)

// TODO: separate reader and readseeker
func NewReader(r io.ReadSeeker, bo binary.ByteOrder) *Reader {
	return &Reader{r: r, bo: bo}
}

type Reader struct {
	r   io.ReadSeeker
	err error
	bo  binary.ByteOrder
}

func (r *Reader) Error() error {
	return r.err
}

func (r *Reader) Seek(offset int64, whence int) (result int64) {
	result, r.err = r.r.Seek(offset, whence)
	return result
}

func (r *Reader) Skip(amount int64) (result int64) {
	return r.Seek(amount, os.SEEK_CUR)
}

func (r *Reader) Offset() (result int64) {
	return r.Seek(0, os.SEEK_CUR)
}

// Read

func (r *Reader) Read(p []byte) (n int) {
	n, r.err = r.r.Read(p)
	return n
}

func (r *Reader) ReadBuf(size int) []byte {
	buf := make([]byte, size)
	r.Read(buf)
	return buf
}

func (r *Reader) ReadU8() uint8 {
	var buf [1]byte
	r.Read(buf[:])
	return uint8(buf[0])
}
func (r *Reader) ReadI8() int8 { return int8(r.ReadU8()) }

func (r *Reader) ReadU16() uint16 {
	var buf [2]byte
	r.Read(buf[:])
	return r.bo.Uint16(buf[:])
}
func (r *Reader) ReadI16() int16 { return int16(r.ReadU16()) }

func (r *Reader) ReadU32() uint32 {
	var buf [4]byte
	r.Read(buf[:])
	return r.bo.Uint32(buf[:])
}
func (r *Reader) ReadI32() int32 { return int32(r.ReadU32()) }

func (r *Reader) ReadU64() uint64 {
	var buf [8]byte
	r.Read(buf[:])
	return r.bo.Uint64(buf[:])
}
func (r *Reader) ReadI64() int64 { return int64(r.ReadU64()) }

// TODO: decide float always little endian or no?
// Or create {Read,Peek}F{32,64}{L,B} helpers
func (r *Reader) ReadF32() float32 { return math.Float32frombits(r.ReadU32()) }
func (r *Reader) ReadF64() float64 { return math.Float64frombits(r.ReadU64()) }

// Peek

func (r *Reader) Peek(p []byte) {
	r.Seek(-int64(r.Read(p)), os.SEEK_CUR)
}

func (r *Reader) PeekBuf(size int) []byte {
	buf := make([]byte, size)
	r.Peek(buf)
	return buf
}

func (r *Reader) PeekU8() uint8 {
	var buf [1]byte
	r.Peek(buf[:])
	return uint8(buf[0])
}
func (r *Reader) PeekI8() int8 { return int8(r.PeekU8()) }

func (r *Reader) PeekU16() uint16 {
	var buf [2]byte
	r.Peek(buf[:])
	return r.bo.Uint16(buf[:])
}
func (r *Reader) PeekI16() int16 { return int16(r.PeekU16()) }

func (r *Reader) PeekU32() uint32 {
	var buf [4]byte
	r.Peek(buf[:])
	return r.bo.Uint32(buf[:])
}
func (r *Reader) PeekI32() int32 { return int32(r.PeekU32()) }

func (r *Reader) PeekU64() uint64 {
	var buf [8]byte
	r.Peek(buf[:])
	return r.bo.Uint64(buf[:])
}
func (r *Reader) PeekI64() int64 { return int64(r.PeekU64()) }

func (r *Reader) PeekF32() float32 { return math.Float32frombits(r.PeekU32()) }
func (r *Reader) PeekF64() float64 { return math.Float64frombits(r.PeekU64()) }
