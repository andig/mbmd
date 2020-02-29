// Code generated by "stringer -type=Measurement"; DO NOT EDIT.

package meters

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Frequency-1]
	_ = x[Current-2]
	_ = x[CurrentL1-3]
	_ = x[CurrentL2-4]
	_ = x[CurrentL3-5]
	_ = x[Voltage-6]
	_ = x[VoltageL1-7]
	_ = x[VoltageL2-8]
	_ = x[VoltageL3-9]
	_ = x[Power-10]
	_ = x[PowerL1-11]
	_ = x[PowerL2-12]
	_ = x[PowerL3-13]
	_ = x[ImportPower-14]
	_ = x[ImportPowerL1-15]
	_ = x[ImportPowerL2-16]
	_ = x[ImportPowerL3-17]
	_ = x[ExportPower-18]
	_ = x[ExportPowerL1-19]
	_ = x[ExportPowerL2-20]
	_ = x[ExportPowerL3-21]
	_ = x[ReactivePower-22]
	_ = x[ReactivePowerL1-23]
	_ = x[ReactivePowerL2-24]
	_ = x[ReactivePowerL3-25]
	_ = x[ApparentPower-26]
	_ = x[ApparentPowerL1-27]
	_ = x[ApparentPowerL2-28]
	_ = x[ApparentPowerL3-29]
	_ = x[Cosphi-30]
	_ = x[CosphiL1-31]
	_ = x[CosphiL2-32]
	_ = x[CosphiL3-33]
	_ = x[THD-34]
	_ = x[THDL1-35]
	_ = x[THDL2-36]
	_ = x[THDL3-37]
	_ = x[Sum-38]
	_ = x[SumT1-39]
	_ = x[SumT2-40]
	_ = x[SumL1-41]
	_ = x[SumL2-42]
	_ = x[SumL3-43]
	_ = x[Import-44]
	_ = x[ImportT1-45]
	_ = x[ImportT2-46]
	_ = x[ImportL1-47]
	_ = x[ImportL2-48]
	_ = x[ImportL3-49]
	_ = x[Export-50]
	_ = x[ExportT1-51]
	_ = x[ExportT2-52]
	_ = x[ExportL1-53]
	_ = x[ExportL2-54]
	_ = x[ExportL3-55]
	_ = x[ReactiveSum-56]
	_ = x[ReactiveSumT1-57]
	_ = x[ReactiveSumT2-58]
	_ = x[ReactiveSumL1-59]
	_ = x[ReactiveSumL2-60]
	_ = x[ReactiveSumL3-61]
	_ = x[ReactiveImport-62]
	_ = x[ReactiveImportT1-63]
	_ = x[ReactiveImportT2-64]
	_ = x[ReactiveImportL1-65]
	_ = x[ReactiveImportL2-66]
	_ = x[ReactiveImportL3-67]
	_ = x[ReactiveExport-68]
	_ = x[ReactiveExportT1-69]
	_ = x[ReactiveExportT2-70]
	_ = x[ReactiveExportL1-71]
	_ = x[ReactiveExportL2-72]
	_ = x[ReactiveExportL3-73]
	_ = x[DCCurrent-74]
	_ = x[DCVoltage-75]
	_ = x[DCPower-76]
	_ = x[HeatSinkTemp-77]
	_ = x[DCCurrentS1-78]
	_ = x[DCVoltageS1-79]
	_ = x[DCPowerS1-80]
	_ = x[DCEnergyS1-81]
	_ = x[DCCurrentS2-82]
	_ = x[DCVoltageS2-83]
	_ = x[DCPowerS2-84]
	_ = x[DCEnergyS2-85]
	_ = x[DCCurrentS3-86]
	_ = x[DCVoltageS3-87]
	_ = x[DCPowerS3-88]
	_ = x[DCEnergyS3-89]
	_ = x[ChargeState-90]
	_ = x[BatteryVoltage-91]
}

const _Measurement_name = "FrequencyCurrentCurrentL1CurrentL2CurrentL3VoltageVoltageL1VoltageL2VoltageL3PowerPowerL1PowerL2PowerL3ImportPowerImportPowerL1ImportPowerL2ImportPowerL3ExportPowerExportPowerL1ExportPowerL2ExportPowerL3ReactivePowerReactivePowerL1ReactivePowerL2ReactivePowerL3ApparentPowerApparentPowerL1ApparentPowerL2ApparentPowerL3CosphiCosphiL1CosphiL2CosphiL3THDTHDL1THDL2THDL3SumSumT1SumT2SumL1SumL2SumL3ImportImportT1ImportT2ImportL1ImportL2ImportL3ExportExportT1ExportT2ExportL1ExportL2ExportL3ReactiveSumReactiveSumT1ReactiveSumT2ReactiveSumL1ReactiveSumL2ReactiveSumL3ReactiveImportReactiveImportT1ReactiveImportT2ReactiveImportL1ReactiveImportL2ReactiveImportL3ReactiveExportReactiveExportT1ReactiveExportT2ReactiveExportL1ReactiveExportL2ReactiveExportL3DCCurrentDCVoltageDCPowerHeatSinkTempDCCurrentS1DCVoltageS1DCPowerS1DCEnergyS1DCCurrentS2DCVoltageS2DCPowerS2DCEnergyS2DCCurrentS3DCVoltageS3DCPowerS3DCEnergyS3ChargeStateBatteryVoltage"

var _Measurement_index = [...]uint16{0, 9, 16, 25, 34, 43, 50, 59, 68, 77, 82, 89, 96, 103, 114, 127, 140, 153, 164, 177, 190, 203, 216, 231, 246, 261, 274, 289, 304, 319, 325, 333, 341, 349, 352, 357, 362, 367, 370, 375, 380, 385, 390, 395, 401, 409, 417, 425, 433, 441, 447, 455, 463, 471, 479, 487, 498, 511, 524, 537, 550, 563, 577, 593, 609, 625, 641, 657, 671, 687, 703, 719, 735, 751, 760, 769, 776, 788, 799, 810, 819, 829, 840, 851, 860, 870, 881, 892, 901, 911, 922, 936}

func (i Measurement) String() string {
	i -= 1
	if i < 0 || i >= Measurement(len(_Measurement_index)-1) {
		return "Measurement(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Measurement_name[_Measurement_index[i]:_Measurement_index[i+1]]
}
