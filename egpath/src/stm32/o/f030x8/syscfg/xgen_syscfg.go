package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x8/mmap"
)

type SYSCFG_Periph struct {
	CFGR1  RCFGR1
	_      uint32
	EXTICR [4]REXTICR
	CFGR2  RCFGR2
}

func (p *SYSCFG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SYSCFG = (*SYSCFG_Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE)))

type CFGR1 uint32

func (b CFGR1) Field(mask CFGR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR1) J(v int) CFGR1 {
	return CFGR1(bits.Make32(v, uint32(mask)))
}

type RCFGR1 struct{ mmio.U32 }

func (r *RCFGR1) Bits(mask CFGR1) CFGR1   { return CFGR1(r.U32.Bits(uint32(mask))) }
func (r *RCFGR1) StoreBits(mask, b CFGR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR1) SetBits(mask CFGR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR1) ClearBits(mask CFGR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR1) Load() CFGR1             { return CFGR1(r.U32.Load()) }
func (r *RCFGR1) Store(b CFGR1)           { r.U32.Store(uint32(b)) }

func (r *RCFGR1) AtomicStoreBits(mask, b CFGR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR1) AtomicSetBits(mask CFGR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR1) AtomicClearBits(mask CFGR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR1 struct{ mmio.UM32 }

func (rm RMCFGR1) Load() CFGR1   { return CFGR1(rm.UM32.Load()) }
func (rm RMCFGR1) Store(b CFGR1) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) MEM_MODE() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(MEM_MODE)}}
}

func (p *SYSCFG_Periph) DMA_RMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(DMA_RMP)}}
}

func (p *SYSCFG_Periph) I2C_FMP_PB6() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_FMP_PB6)}}
}

func (p *SYSCFG_Periph) I2C_FMP_PB7() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_FMP_PB7)}}
}

func (p *SYSCFG_Periph) I2C_FMP_PB8() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_FMP_PB8)}}
}

func (p *SYSCFG_Periph) I2C_FMP_PB9() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_FMP_PB9)}}
}

type EXTICR uint32

func (b EXTICR) Field(mask EXTICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EXTICR) J(v int) EXTICR {
	return EXTICR(bits.Make32(v, uint32(mask)))
}

type REXTICR struct{ mmio.U32 }

func (r *REXTICR) Bits(mask EXTICR) EXTICR  { return EXTICR(r.U32.Bits(uint32(mask))) }
func (r *REXTICR) StoreBits(mask, b EXTICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REXTICR) SetBits(mask EXTICR)      { r.U32.SetBits(uint32(mask)) }
func (r *REXTICR) ClearBits(mask EXTICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *REXTICR) Load() EXTICR             { return EXTICR(r.U32.Load()) }
func (r *REXTICR) Store(b EXTICR)           { r.U32.Store(uint32(b)) }

func (r *REXTICR) AtomicStoreBits(mask, b EXTICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *REXTICR) AtomicSetBits(mask EXTICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *REXTICR) AtomicClearBits(mask EXTICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMEXTICR struct{ mmio.UM32 }

func (rm RMEXTICR) Load() EXTICR   { return EXTICR(rm.UM32.Load()) }
func (rm RMEXTICR) Store(b EXTICR) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) EXTI0(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI0)}}
}

func (p *SYSCFG_Periph) EXTI1(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI1)}}
}

func (p *SYSCFG_Periph) EXTI2(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI2)}}
}

func (p *SYSCFG_Periph) EXTI3(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI3)}}
}

type CFGR2 uint32

func (b CFGR2) Field(mask CFGR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR2) J(v int) CFGR2 {
	return CFGR2(bits.Make32(v, uint32(mask)))
}

type RCFGR2 struct{ mmio.U32 }

func (r *RCFGR2) Bits(mask CFGR2) CFGR2   { return CFGR2(r.U32.Bits(uint32(mask))) }
func (r *RCFGR2) StoreBits(mask, b CFGR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) SetBits(mask CFGR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR2) ClearBits(mask CFGR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR2) Load() CFGR2             { return CFGR2(r.U32.Load()) }
func (r *RCFGR2) Store(b CFGR2)           { r.U32.Store(uint32(b)) }

func (r *RCFGR2) AtomicStoreBits(mask, b CFGR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) AtomicSetBits(mask CFGR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR2) AtomicClearBits(mask CFGR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR2 struct{ mmio.UM32 }

func (rm RMCFGR2) Load() CFGR2   { return CFGR2(rm.UM32.Load()) }
func (rm RMCFGR2) Store(b CFGR2) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) LOCKUP_LOCK() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(LOCKUP_LOCK)}}
}

func (p *SYSCFG_Periph) SRAM_PARITY_LOCK() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(SRAM_PARITY_LOCK)}}
}

func (p *SYSCFG_Periph) SRAM_PEF() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(SRAM_PEF)}}
}
