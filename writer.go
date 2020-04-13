package binrw

import (
	"encoding/binary"
	"io"
	"math"
	"os"
)

func NewWriter(w io.WriteSeeker, bo binary.ByteOrder) *Writer {
	return &Writer{w: w, bo: bo}
}

type Writer struct {
	w   io.WriteSeeker
	err error
	bo  binary.ByteOrder
}

func (w *Writer) Error() error {
	return w.err
}

func (w *Writer) Seek(offset int64, whence int) (result int64) {
	result, w.err = w.w.Seek(offset, whence)
	return result
}

func (w *Writer) Skip(amount int64) (result int64) {
	return w.Seek(amount, os.SEEK_CUR)
}

func (w *Writer) Offset() (result int64) {
	return w.Seek(0, os.SEEK_CUR)
}

// Write

func (w *Writer) Write(p []byte) (n int) {
	n, w.err = w.w.Write(p)
	return n
}

func (w *Writer) WriteU8(v uint8) {
	var buf [1]byte
	buf[0] = v
	w.Write(buf[:])
}
func (w *Writer) WriteI8(v int8) { w.WriteU8(uint8(v)) }

func (w *Writer) WriteU16(v uint16) {
	var buf [2]byte
	w.bo.PutUint16(buf[:], v)
	w.Write(buf[:])
}
func (w *Writer) WriteI16(v int16) { w.WriteU16(uint16(v)) }

func (w *Writer) WriteU32(v uint32) {
	var buf [4]byte
	w.bo.PutUint32(buf[:], v)
	w.Write(buf[:])
}
func (w *Writer) WriteI32(v int32) { w.WriteU32(uint32(v)) }

func (w *Writer) WriteU64(v uint64) {
	var buf [8]byte
	w.bo.PutUint64(buf[:], v)
	w.Write(buf[:])
}
func (w *Writer) WriteI64(v int64) { w.WriteU64(uint64(v)) }

func (w *Writer) WriteF32(v float32) { w.WriteU32(math.Float32bits(v)) }
func (w *Writer) WriteF64(v float64) { w.WriteU64(math.Float64bits(v)) }

// TODO: decide float always little endian or no?
// Or create WriteF{32,64}{L,B} helpers
