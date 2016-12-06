// +build f746xx

// Peripheral: DMA_Periph  DMA Controller.
// Instances:
//  DMA1  mmap.DMA1_BASE
//  DMA2  mmap.DMA2_BASE
// Registers:
//  0x00 32  LISR  Low interrupt status register.
//  0x04 32  HISR  High interrupt status register.
//  0x08 32  LIFCR Low interrupt flag clear register.
//  0x0C 32  HIFCR High interrupt flag clear register.
// Import:
//  stm32/o/f746xx/mmap
package dma

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	TCIF3  LISR_Bits = 0x01 << 27 //+
	HTIF3  LISR_Bits = 0x01 << 26 //+
	TEIF3  LISR_Bits = 0x01 << 25 //+
	DMEIF3 LISR_Bits = 0x01 << 24 //+
	FEIF3  LISR_Bits = 0x01 << 22 //+
	TCIF2  LISR_Bits = 0x01 << 21 //+
	HTIF2  LISR_Bits = 0x01 << 20 //+
	TEIF2  LISR_Bits = 0x01 << 19 //+
	DMEIF2 LISR_Bits = 0x01 << 18 //+
	FEIF2  LISR_Bits = 0x01 << 16 //+
	TCIF1  LISR_Bits = 0x01 << 11 //+
	HTIF1  LISR_Bits = 0x01 << 10 //+
	TEIF1  LISR_Bits = 0x01 << 9  //+
	DMEIF1 LISR_Bits = 0x01 << 8  //+
	FEIF1  LISR_Bits = 0x01 << 6  //+
	TCIF0  LISR_Bits = 0x01 << 5  //+
	HTIF0  LISR_Bits = 0x01 << 4  //+
	TEIF0  LISR_Bits = 0x01 << 3  //+
	DMEIF0 LISR_Bits = 0x01 << 2  //+
	FEIF0  LISR_Bits = 0x01 << 0  //+
)

const (
	TCIF3n  = 27
	HTIF3n  = 26
	TEIF3n  = 25
	DMEIF3n = 24
	FEIF3n  = 22
	TCIF2n  = 21
	HTIF2n  = 20
	TEIF2n  = 19
	DMEIF2n = 18
	FEIF2n  = 16
	TCIF1n  = 11
	HTIF1n  = 10
	TEIF1n  = 9
	DMEIF1n = 8
	FEIF1n  = 6
	TCIF0n  = 5
	HTIF0n  = 4
	TEIF0n  = 3
	DMEIF0n = 2
	FEIF0n  = 0
)

const (
	TCIF7  HISR_Bits = 0x01 << 27 //+
	HTIF7  HISR_Bits = 0x01 << 26 //+
	TEIF7  HISR_Bits = 0x01 << 25 //+
	DMEIF7 HISR_Bits = 0x01 << 24 //+
	FEIF7  HISR_Bits = 0x01 << 22 //+
	TCIF6  HISR_Bits = 0x01 << 21 //+
	HTIF6  HISR_Bits = 0x01 << 20 //+
	TEIF6  HISR_Bits = 0x01 << 19 //+
	DMEIF6 HISR_Bits = 0x01 << 18 //+
	FEIF6  HISR_Bits = 0x01 << 16 //+
	TCIF5  HISR_Bits = 0x01 << 11 //+
	HTIF5  HISR_Bits = 0x01 << 10 //+
	TEIF5  HISR_Bits = 0x01 << 9  //+
	DMEIF5 HISR_Bits = 0x01 << 8  //+
	FEIF5  HISR_Bits = 0x01 << 6  //+
	TCIF4  HISR_Bits = 0x01 << 5  //+
	HTIF4  HISR_Bits = 0x01 << 4  //+
	TEIF4  HISR_Bits = 0x01 << 3  //+
	DMEIF4 HISR_Bits = 0x01 << 2  //+
	FEIF4  HISR_Bits = 0x01 << 0  //+
)

const (
	TCIF7n  = 27
	HTIF7n  = 26
	TEIF7n  = 25
	DMEIF7n = 24
	FEIF7n  = 22
	TCIF6n  = 21
	HTIF6n  = 20
	TEIF6n  = 19
	DMEIF6n = 18
	FEIF6n  = 16
	TCIF5n  = 11
	HTIF5n  = 10
	TEIF5n  = 9
	DMEIF5n = 8
	FEIF5n  = 6
	TCIF4n  = 5
	HTIF4n  = 4
	TEIF4n  = 3
	DMEIF4n = 2
	FEIF4n  = 0
)

const (
	CTCIF3  LIFCR_Bits = 0x01 << 27 //+
	CHTIF3  LIFCR_Bits = 0x01 << 26 //+
	CTEIF3  LIFCR_Bits = 0x01 << 25 //+
	CDMEIF3 LIFCR_Bits = 0x01 << 24 //+
	CFEIF3  LIFCR_Bits = 0x01 << 22 //+
	CTCIF2  LIFCR_Bits = 0x01 << 21 //+
	CHTIF2  LIFCR_Bits = 0x01 << 20 //+
	CTEIF2  LIFCR_Bits = 0x01 << 19 //+
	CDMEIF2 LIFCR_Bits = 0x01 << 18 //+
	CFEIF2  LIFCR_Bits = 0x01 << 16 //+
	CTCIF1  LIFCR_Bits = 0x01 << 11 //+
	CHTIF1  LIFCR_Bits = 0x01 << 10 //+
	CTEIF1  LIFCR_Bits = 0x01 << 9  //+
	CDMEIF1 LIFCR_Bits = 0x01 << 8  //+
	CFEIF1  LIFCR_Bits = 0x01 << 6  //+
	CTCIF0  LIFCR_Bits = 0x01 << 5  //+
	CHTIF0  LIFCR_Bits = 0x01 << 4  //+
	CTEIF0  LIFCR_Bits = 0x01 << 3  //+
	CDMEIF0 LIFCR_Bits = 0x01 << 2  //+
	CFEIF0  LIFCR_Bits = 0x01 << 0  //+
)

const (
	CTCIF3n  = 27
	CHTIF3n  = 26
	CTEIF3n  = 25
	CDMEIF3n = 24
	CFEIF3n  = 22
	CTCIF2n  = 21
	CHTIF2n  = 20
	CTEIF2n  = 19
	CDMEIF2n = 18
	CFEIF2n  = 16
	CTCIF1n  = 11
	CHTIF1n  = 10
	CTEIF1n  = 9
	CDMEIF1n = 8
	CFEIF1n  = 6
	CTCIF0n  = 5
	CHTIF0n  = 4
	CTEIF0n  = 3
	CDMEIF0n = 2
	CFEIF0n  = 0
)

const (
	CTCIF7  HIFCR_Bits = 0x01 << 27 //+
	CHTIF7  HIFCR_Bits = 0x01 << 26 //+
	CTEIF7  HIFCR_Bits = 0x01 << 25 //+
	CDMEIF7 HIFCR_Bits = 0x01 << 24 //+
	CFEIF7  HIFCR_Bits = 0x01 << 22 //+
	CTCIF6  HIFCR_Bits = 0x01 << 21 //+
	CHTIF6  HIFCR_Bits = 0x01 << 20 //+
	CTEIF6  HIFCR_Bits = 0x01 << 19 //+
	CDMEIF6 HIFCR_Bits = 0x01 << 18 //+
	CFEIF6  HIFCR_Bits = 0x01 << 16 //+
	CTCIF5  HIFCR_Bits = 0x01 << 11 //+
	CHTIF5  HIFCR_Bits = 0x01 << 10 //+
	CTEIF5  HIFCR_Bits = 0x01 << 9  //+
	CDMEIF5 HIFCR_Bits = 0x01 << 8  //+
	CFEIF5  HIFCR_Bits = 0x01 << 6  //+
	CTCIF4  HIFCR_Bits = 0x01 << 5  //+
	CHTIF4  HIFCR_Bits = 0x01 << 4  //+
	CTEIF4  HIFCR_Bits = 0x01 << 3  //+
	CDMEIF4 HIFCR_Bits = 0x01 << 2  //+
	CFEIF4  HIFCR_Bits = 0x01 << 0  //+
)

const (
	CTCIF7n  = 27
	CHTIF7n  = 26
	CTEIF7n  = 25
	CDMEIF7n = 24
	CFEIF7n  = 22
	CTCIF6n  = 21
	CHTIF6n  = 20
	CTEIF6n  = 19
	CDMEIF6n = 18
	CFEIF6n  = 16
	CTCIF5n  = 11
	CHTIF5n  = 10
	CTEIF5n  = 9
	CDMEIF5n = 8
	CFEIF5n  = 6
	CTCIF4n  = 5
	CHTIF4n  = 4
	CTEIF4n  = 3
	CDMEIF4n = 2
	CFEIF4n  = 0
)