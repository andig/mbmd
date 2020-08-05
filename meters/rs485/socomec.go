package rs485

import . "github.com/volkszaehler/mbmd/meters"

func init() {
	Register(NewSocomecProducer)
}

const (
	METERTYPE_SOCOMEC = "SOCOMEC"
)

type SocomecProducer struct {
	Opcodes
}

func NewSocomecProducer() Producer {
	/**
	 * Opcodes as defined by Socomec DIRIS A10
	 * See https://www.socomec.us/wp-content/uploads/2019/03/DIRIS-A10-HTML_COMMUNICATION-TABLE_2017-08_CMT_MULTI.html
	 */
	ops := Opcodes{
		VoltageL1:     50520, // / 100
		VoltageL2:     50522, // / 100
		VoltageL3:     50524, // / 100
		CurrentL1:     50528, // / 1000
		CurrentL2:     50530, // / 1000
		CurrentL3:     50532, // / 1000
		Current:       50534, // / 1000
		PowerL1:       50544, // / 0.1
		PowerL2:       50546, // / 0.1
		PowerL3:       50546, // / 0.1
		Power:         50536, // / 0.1
		ApparentPower: 50540, // / 0.1
		ReactivePower: 50538, // / 0.1
		Import:        51311, // / 1E-06
		Export:        51313, // / 1E-06
		CosphiL1:      51305, // / 1000
		CosphiL2:      51306, // / 1000
		CosphiL3:      51307, // / 1000
		THDL1:         51539, // / 10
		THDL2:         51540, // / 10
		THDL3:         51541, // / 10
		Frequency:     51287, // / 100
	}
	return &SocomecProducer{Opcodes: ops}
}

func (p *SocomecProducer) Type() string {
	return METERTYPE_SOCOMEC
}

func (p *SocomecProducer) Description() string {
	return "Socomec DIRIS A10"
}

func (p *SocomecProducer) snip(iec Measurement) Operation {
	operation := Operation{
		FuncCode:  ReadInputReg,
		OpCode:    p.Opcode(iec),
		ReadLen:   2,
		IEC61850:  iec,
		Transform: RTUIeee754ToFloat64,
	}
	return operation
}

func (p *SocomecProducer) Probe() Operation {
	return p.snip(VoltageL1)
}

func (p *SocomecProducer) Produce() (res []Operation) {
	for op := range p.Opcodes {
		res = append(res, p.snip(op))
	}

	return res
}
