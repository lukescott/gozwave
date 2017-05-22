package database

type Value struct {
	From int64
	To int64
	Desc string
	Unit string
}
type parameter struct {
	ID int
	Name string
	Type string
	Description string
	Size int
	Default string

	Values []Value
}
type Device struct{
	Brand string
	Product string
	Description string

	CommandClasses []*CommandClass
	Parameters []*parameter

	ManufacturerID string
	ProductType string
	ProductID string
}
func New(manufacturerID, productType, productID string) *Device{
	dev := manufacturerID+"_"+productType+"_"+productID
	switch dev {
	case "0081_00a0_0001":
		return New0081_00A0_0001()
	case "0068_0000_0010":
		return New0068_0000_0010()
	case "010a_1100_4b00":
		return New010A_1100_4B00()
	case "0001_4450_3030":
		return New0001_4450_3030()
	case "0001_444d_3330":
		return New0001_444D_3330()
	case "0001_4457_3033":
		return New0001_4457_3033()
	case "0001_4457_3230":
		return New0001_4457_3230()
	case "0001_4952_3130":
		return New0001_4952_3130()
	case "0001_524d_3330":
		return New0001_524D_3330()
	case "0001_7fff_7fff":
		return New0001_7FFF_7FFF()
	case "0001_5257_3330":
		return New0001_5257_3330()
	case "0001_544d_3330":
		return New0001_544D_3330()
	case "0086_0002_0078":
		return New0086_0002_0078()
	case "0086_0102_0078":
		return New0086_0102_0078()
	case "0086_0001_0003":
		return New0086_0001_0003()
	case "0086_0006_0002":
		return New0086_0006_0002()
	case "0086_0001_0026":
		return New0086_0001_0026()
	case "0086_0002_0005":
		return New0086_0002_0005()
	case "0086_0002_0009":
		return New0086_0002_0009()
	case "0086_0002_001c":
		return New0086_0002_001C()
	case "0086_0002_0004":
		return New0086_0002_0004()
	case "0086_0002_001d":
		return New0086_0002_001D()
	case "0086_0002_002d":
		return New0086_0002_002D()
	case "0086_0002_0036":
		return New0086_0002_0036()
	case "0086_0003_0006":
		return New0086_0003_0006()
	case "0086_0003_0008":
		return New0086_0003_0008()
	case "0086_0003_000a":
		return New0086_0003_000A()
	case "0086_0003_000b":
		return New0086_0003_000B()
	case "0086_0003_000c":
		return New0086_0003_000C()
	case "0086_0003_000d":
		return New0086_0003_000D()
	case "0086_0003_000e":
		return New0086_0003_000E()
	case "0086_0003_0011":
		return New0086_0003_0011()
	case "0086_0003_0012":
		return New0086_0003_0012()
	case "0086_0003_0013":
		return New0086_0003_0013()
	case "0086_0003_0018":
		return New0086_0003_0018()
	case "0086_0103_004b":
		return New0086_0103_004B()
	case "0086_0003_0019":
		return New0086_0003_0019()
	case "0086_0003_001a":
		return New0086_0003_001A()
	case "0086_0003_001b":
		return New0086_0003_001B()
	case "0086_0003_0023":
		return New0086_0003_0023()
	case "0086_0004_0050":
		return New0086_0004_0050()
	case "0086_0004_0025":
		return New0086_0004_0025()
	case "0086_0019_0004":
		return New0086_0019_0004()
	case "0086_0004_0038":
		return New0086_0004_0038()
	case "0086_0104_0038":
		return New0086_0104_0038()
	case "0086_0204_0038":
		return New0086_0204_0038()
	case "0086_0304_0038":
		return New0086_0304_0038()
	case "0086_0904_0038":
		return New0086_0904_0038()
	case "0086_0a04_0038":
		return New0086_0A04_0038()
	case "0086_1a04_0038":
		return New0086_1A04_0038()
	case "0086_1d04_0038":
		return New0086_1D04_0038()
	case "0086_0003_003e":
		return New0086_0003_003E()
	case "0086_0103_003e":
		return New0086_0103_003E()
	case "0086_0203_003e":
		return New0086_0203_003E()
	case "0086_0002_004a":
		return New0086_0002_004A()
	case "0086_0102_004a":
		return New0086_0102_004A()
	case "0086_0202_004a":
		return New0086_0202_004A()
	case "0086_0003_004b":
		return New0086_0003_004B()
	case "0086_0203_004b":
		return New0086_0203_004B()
	case "0086_0003_004e":
		return New0086_0003_004E()
	case "0086_0103_004e":
		return New0086_0103_004E()
	case "0086_0203_004e":
		return New0086_0203_004E()
	case "0086_0104_0050":
		return New0086_0104_0050()
	case "0086_0204_0050":
		return New0086_0204_0050()
	case "0086_0001_0058":
		return New0086_0001_0058()
	case "0086_0101_0058":
		return New0086_0101_0058()
	case "0086_0201_0058":
		return New0086_0201_0058()
	case "0086_0002_0059":
		return New0086_0002_0059()
	case "0086_0102_0059":
		return New0086_0102_0059()
	case "0086_0202_0059":
		return New0086_0202_0059()
	case "0086_0002_005f":
		return New0086_0002_005F()
	case "0086_0102_005f":
		return New0086_0102_005F()
	case "0086_0003_0060":
		return New0086_0003_0060()
	case "0086_0103_0060":
		return New0086_0103_0060()
	case "0086_0203_0060":
		return New0086_0203_0060()
	case "0086_1d03_0060":
		return New0086_1D03_0060()
	case "0086_0002_0061":
		return New0086_0002_0061()
	case "0086_0102_0061":
		return New0086_0102_0061()
	case "0086_0202_0061":
		return New0086_0202_0061()
	case "0086_0003_0062":
		return New0086_0003_0062()
	case "0086_0103_0062":
		return New0086_0103_0062()
	case "0086_0203_0062":
		return New0086_0203_0062()
	case "0086_0003_0063":
		return New0086_0003_0063()
	case "0086_0103_0063":
		return New0086_0103_0063()
	case "0086_0203_0063":
		return New0086_0203_0063()
	case "0086_1d03_0063":
		return New0086_1D03_0063()
	case "0086_0002_0064":
		return New0086_0002_0064()
	case "0086_0102_0064":
		return New0086_0102_0064()
	case "0086_0202_0064":
		return New0086_0202_0064()
	case "0086_0103_006f":
		return New0086_0103_006F()
	case "0086_0203_006f":
		return New0086_0203_006F()
	case "0086_0002_0070":
		return New0086_0002_0070()
	case "0086_0102_0070":
		return New0086_0102_0070()
	case "0086_0202_0070":
		return New0086_0202_0070()
	case "0086_1d02_0070":
		return New0086_1D02_0070()
	case "0086_0004_0075":
		return New0086_0004_0075()
	case "0086_0002_0082":
		return New0086_0002_0082()
	case "0086_0102_0082":
		return New0086_0102_0082()
	case "0086_0202_0082":
		return New0086_0202_0082()
	case "0111_8200_0200":
		return New0111_8200_0200()
	case "008a_0005_0101":
		return New008A_0005_0101()
	case "008a_000d_0100":
		return New008A_000D_0100()
	case "008a_0004_0100":
		return New008A_0004_0100()
	case "008a_0004_0101":
		return New008A_0004_0101()
	case "008a_0020_0001":
		return New008A_0020_0001()
	case "008a_0021_0102":
		return New008A_0021_0102()
	case "008a_0003_0101":
		return New008A_0003_0101()
	case "008a_0003_0100":
		return New008A_0003_0100()
	case "008a_002f_0100":
		return New008A_002F_0100()
	case "008a_0018_0100":
		return New008A_0018_0100()
	case "008a_0008_0101":
		return New008A_0008_0101()
	case "008a_0007_0101":
		return New008A_0007_0101()
	case "0169_0001_0001":
		return New0169_0001_0001()
	case "0138_0001_0002":
		return New0138_0001_0002()
	case "0138_0001_0001":
		return New0138_0001_0001()
	case "0166_0100_0100":
		return New0166_0100_0100()
	case "0116_0002_0001":
		return New0116_0002_0001()
	case "0116_0001_0001":
		return New0116_0001_0001()
	case "0200_0003_0002":
		return New0200_0003_0002()
	case "0179_0013_0021":
		return New0179_0013_0021()
	case "001a_5342_0000":
		return New001A_5342_0000()
	case "001a_534c_0000":
		return New001A_534C_0000()
	case "001a_5352_0000":
		return New001A_5352_0000()
	case "001a_4449_0002":
		return New001A_4449_0002()
	case "001a_4449_0101":
		return New001A_4449_0101()
	case "001a_4441_0000":
		return New001A_4441_0000()
	case "001a_5244_0000":
		return New001A_5244_0000()
	case "001a_574d_0000":
		return New001A_574D_0000()
	case "0248_0003_0001":
		return New0248_0003_0001()
	case "0002_0115_a010":
		return New0002_0115_A010()
	case "0002_0003_8010":
		return New0002_0003_8010()
	case "0002_0248_a020":
		return New0002_0248_A020()
	case "0002_0248_a010":
		return New0002_0248_A010()
	case "0002_0005_0004":
		return New0002_0005_0004()
	case "0002_8005_0001":
		return New0002_8005_0001()
	case "0002_8005_0002":
		return New0002_8005_0002()
	case "0002_0005_0003":
		return New0002_0005_0003()
	case "0002_0005_0175":
		return New0002_0005_0175()
	case "0175_0001_0011":
		return New0175_0001_0011()
	case "0175_0002_000d":
		return New0175_0002_000D()
	case "0175_0002_000e":
		return New0175_0002_000E()
	case "0175_0100_0101":
		return New0175_0100_0101()
	case "0175_0100_0102":
		return New0175_0100_0102()
	case "0175_0002_0020":
		return New0175_0002_0020()
	case "0175_0002_0021":
		return New0175_0002_0021()
	case "0175_0001_0001":
		return New0175_0001_0001()
	case "0175_0004_000a":
		return New0175_0004_000A()
	case "0175_0002_0018":
		return New0175_0002_0018()
	case "0108_0002_000e":
		return New0108_0002_000E()
	case "0108_0002_000d":
		return New0108_0002_000D()
	case "0108_0004_000a":
		return New0108_0004_000A()
	case "020e_4c42_3135":
		return New020E_4C42_3135()
	case "020e_4754_3038":
		return New020E_4754_3038()
	case "020e_4c42_3134":
		return New020E_4C42_3134()
	case "0184_4447_3031":
		return New0184_4447_3031()
	case "0184_4447_3034":
		return New0184_4447_3034()
	case "0184_4744_3032":
		return New0184_4744_3032()
	case "0184_4447_3033":
		return New0184_4447_3033()
	case "014a_0001_0002":
		return New014A_0001_0002()
	case "014a_0004_0002":
		return New014A_0004_0002()
	case "014a_0004_0001":
		return New014A_0004_0001()
	case "014a_0001_0001":
		return New014A_0001_0001()
	case "014a_0004_0003":
		return New014A_0004_0003()
	case "014a_0001_0003":
		return New014A_0001_0003()
	case "0157_0003_0512":
		return New0157_0003_0512()
	case "0157_0100_0100":
		return New0157_0100_0100()
	case "0157_0003_0002":
		return New0157_0003_0002()
	case "021f_0003_0088":
		return New021F_0003_0088()
	case "011a_0101_0102":
		return New011A_0101_0102()
	case "011a_0101_0103":
		return New011A_0101_0103()
	case "011a_0101_0603":
		return New011A_0101_0603()
	case "011a_0111_0101":
		return New011A_0111_0101()
	case "011a_0102_0201":
		return New011A_0102_0201()
	case "011a_0111_0201":
		return New011A_0111_0201()
	case "011a_0111_0605":
		return New011A_0111_0605()
	case "011a_0101_0104":
		return New011A_0101_0104()
	case "011a_0601_0901":
		return New011A_0601_0901()
	case "011a_0101_5605":
		return New011A_0101_5605()
	case "011a_0101_5606":
		return New011A_0101_5606()
	case "011a_0801_0b03":
		return New011A_0801_0B03()
	case "0148_0002_0001":
		return New0148_0002_0001()
	case "0148_0001_0001":
		return New0148_0001_0001()
	case "0135_000b_0001":
		return New0135_000B_0001()
	case "0060_0003_0001":
		return New0060_0003_0001()
	case "0060_0003_0002":
		return New0060_0003_0002()
	case "0060_0003_0003":
		return New0060_0003_0003()
	case "0060_0004_0001":
		return New0060_0004_0001()
	case "0060_0104_0001":
		return New0060_0104_0001()
	case "0060_0004_0002":
		return New0060_0004_0002()
	case "0060_0004_0005":
		return New0060_0004_0005()
	case "0060_0004_0008":
		return New0060_0004_0008()
	case "0060_0011_0002":
		return New0060_0011_0002()
	case "0060_0004_0007":
		return New0060_0004_0007()
	case "0060_0004_0006":
		return New0060_0004_0006()
	case "0060_0010_0001":
		return New0060_0010_0001()
	case "0060_0202_0001":
		return New0060_0202_0001()
	case "0060_0002_0002":
		return New0060_0002_0002()
	case "0060_0001_0003":
		return New0060_0001_0003()
	case "0060_000c_0001":
		return New0060_000C_0001()
	case "0060_000d_0001":
		return New0060_000D_0001()
	case "0060_0002_0001":
		return New0060_0002_0001()
	case "0060_0101_0001":
		return New0060_0101_0001()
	case "0060_0001_0002":
		return New0060_0001_0002()
	case "0060_000b_0001":
		return New0060_000B_0001()
	case "0060_0006_0001":
		return New0060_0006_0001()
	case "0060_0007_0001":
		return New0060_0007_0001()
	case "0113_4450_3030":
		return New0113_4450_3030()
	case "0113_5246_3133":
		return New0113_5246_3133()
	case "0113_5250_3030":
		return New0113_5250_3030()
	case "0113_4457_3034":
		return New0113_4457_3034()
	case "0113_5257_3533":
		return New0113_5257_3533()
	case "0113_4556_5434":
		return New0113_4556_5434()
	case "0085_0002_0002":
		return New0085_0002_0002()
	case "0085_0002_0001":
		return New0085_0002_0001()
	case "0085_0003_0001":
		return New0085_0003_0001()
	case "010f_0501_0102":
		return New010F_0501_0102()
	case "010f_0501_1002":
		return New010F_0501_1002()
	case "010f_0501_2002":
		return New010F_0501_2002()
	case "010f_0501_3002":
		return New010F_0501_3002()
	case "010f_0d01_1000":
		return New010F_0D01_1000()
	case "010f_0d01_3000":
		return New010F_0D01_3000()
	case "010f_0100_0104":
		return New010F_0100_0104()
	case "010f_0100_0106":
		return New010F_0100_0106()
	case "010f_0100_0107":
		return New010F_0100_0107()
	case "010f_0100_0109":
		return New010F_0100_0109()
	case "010f_0100_100a":
		return New010F_0100_100A()
	case "010f_0100_300a":
		return New010F_0100_300A()
	case "010f_0102_1000":
		return New010F_0102_1000()
	case "010f_0102_2000":
		return New010F_0102_2000()
	case "010f_0102_3000":
		return New010F_0102_3000()
	case "010f_0102_4000":
		return New010F_0102_4000()
	case "010f_0b00_1001":
		return New010F_0B00_1001()
	case "010f_0b00_2001":
		return New010F_0B00_2001()
	case "010f_0b00_3001":
		return New010F_0B00_3001()
	case "010f_0b01_1002":
		return New010F_0B01_1002()
	case "010f_0b01_2002":
		return New010F_0B01_2002()
	case "010f_0700_1000":
		return New010F_0700_1000()
	case "010f_0700_2000":
		return New010F_0700_2000()
	case "010f_0700_3000":
		return New010F_0700_3000()
	case "010f_0701_1001":
		return New010F_0701_1001()
	case "010f_0701_3001":
		return New010F_0701_3001()
	case "010f_1001_1000":
		return New010F_1001_1000()
	case "010f_0800_1001":
		return New010F_0800_1001()
	case "010f_0800_2001":
		return New010F_0800_2001()
	case "010f_0800_3001":
		return New010F_0800_3001()
	case "010f_0801_1001":
		return New010F_0801_1001()
	case "010f_0801_2001":
		return New010F_0801_2001()
	case "010f_0801_3001":
		return New010F_0801_3001()
	case "010f_8800_3001":
		return New010F_8800_3001()
	case "010f_0f01_1000":
		return New010F_0F01_1000()
	case "010f_0f01_2000":
		return New010F_0F01_2000()
	case "010f_0f01_3000":
		return New010F_0F01_3000()
	case "010f_fb10_1014":
		return New010F_FB10_1014()
	case "010f_0300_0104":
		return New010F_0300_0104()
	case "010f_0300_0106":
		return New010F_0300_0106()
	case "010f_0300_0107":
		return New010F_0300_0107()
	case "010f_0300_100a":
		return New010F_0300_100A()
	case "010f_0300_0109":
		return New010F_0300_0109()
	case "010f_0900_1000":
		return New010F_0900_1000()
	case "010f_0900_2000":
		return New010F_0900_2000()
	case "010f_0900_3000":
		return New010F_0900_3000()
	case "010f_0900_4000":
		return New010F_0900_4000()
	case "010f_0301_1001":
		return New010F_0301_1001()
	case "010f_0301_3001":
		return New010F_0301_3001()
	case "010f_0302_1000":
		return New010F_0302_1000()
	case "010f_0302_3000":
		return New010F_0302_3000()
	case "010f_0400_0102":
		return New010F_0400_0102()
	case "010f_0400_0104":
		return New010F_0400_0104()
	case "010f_0400_0105":
		return New010F_0400_0105()
	case "010f_0400_0106":
		return New010F_0400_0106()
	case "010f_0400_0107":
		return New010F_0400_0107()
	case "010f_0400_0108":
		return New010F_0400_0108()
	case "010f_0400_0109":
		return New010F_0400_0109()
	case "010f_0400_100a":
		return New010F_0400_100A()
	case "010f_0400_300a":
		return New010F_0400_300A()
	case "010f_0402_3002":
		return New010F_0402_3002()
	case "010f_0401_100a":
		return New010F_0401_100A()
	case "010f_0402_1002":
		return New010F_0402_1002()
	case "010f_0403_1000":
		return New010F_0403_1000()
	case "010f_0403_3000":
		return New010F_0403_3000()
	case "010f_0200_0102":
		return New010F_0200_0102()
	case "010f_0200_0103":
		return New010F_0200_0103()
	case "010f_0200_0104":
		return New010F_0200_0104()
	case "010f_0200_0105":
		return New010F_0200_0105()
	case "010f_0200_0106":
		return New010F_0200_0106()
	case "010f_0200_0107":
		return New010F_0200_0107()
	case "010f_0200_0109":
		return New010F_0200_0109()
	case "010f_0200_100a":
		return New010F_0200_100A()
	case "010f_0200_300a":
		return New010F_0200_300A()
	case "010f_0202_3002":
		return New010F_0202_3002()
	case "010f_0202_1002":
		return New010F_0202_1002()
	case "010f_0202_4002":
		return New010F_0202_4002()
	case "010f_0203_1000":
		return New010F_0203_1000()
	case "010f_0203_2000":
		return New010F_0203_2000()
	case "010f_0203_3000":
		return New010F_0203_3000()
	case "010f_0c02_1002":
		return New010F_0C02_1002()
	case "010f_0c00_1000":
		return New010F_0C00_1000()
	case "010f_0c00_3000":
		return New010F_0C00_3000()
	case "010f_0600_1000":
		return New010F_0600_1000()
	case "010f_0602_1001":
		return New010F_0602_1001()
	case "0207_0027_0100":
		return New0207_0027_0100()
	case "0084_0451_010e":
		return New0084_0451_010E()
	case "0084_0453_010e":
		return New0084_0453_010E()
	case "0084_0453_010f":
		return New0084_0453_010F()
	case "0084_0453_0110":
		return New0084_0453_0110()
	case "0084_0453_0111":
		return New0084_0453_0111()
	case "0084_0311_010b":
		return New0084_0311_010B()
	case "0084_0313_0108":
		return New0084_0313_0108()
	case "0084_0313_010b":
		return New0084_0313_010B()
	case "0084_0213_0215":
		return New0084_0213_0215()
	case "0084_0061_010c":
		return New0084_0061_010C()
	case "0084_0063_010c":
		return New0084_0063_010C()
	case "0063_5044_3031":
		return New0063_5044_3031()
	case "0063_4952_3033":
		return New0063_4952_3033()
	case "0063_4944_3033":
		return New0063_4944_3033()
	case "0063_4c42_3031":
		return New0063_4C42_3031()
	case "0063_5052_3031":
		return New0063_5052_3031()
	case "0063_5250_3030":
		return New0063_5250_3030()
	case "0063_5250_3130":
		return New0063_5250_3130()
	case "0063_4f50_3031":
		return New0063_4F50_3031()
	case "0063_5252_3530":
		return New0063_5252_3530()
	case "0063_4457_3230":
		return New0063_4457_3230()
	case "0063_4457_3033":
		return New0063_4457_3033()
	case "0063_5257_3533":
		return New0063_5257_3533()
	case "0063_6363_3533":
		return New0063_6363_3533()
	case "0063_4944_3032":
		return New0063_4944_3032()
	case "0063_4952_3036":
		return New0063_4952_3036()
	case "0063_4944_3038":
		return New0063_4944_3038()
	case "0063_4944_3135":
		return New0063_4944_3135()
	case "0063_5052_3033":
		return New0063_5052_3033()
	case "0063_4953_3032":
		return New0063_4953_3032()
	case "0063_4952_3031":
		return New0063_4952_3031()
	case "0063_4944_3031":
		return New0063_4944_3031()
	case "0063_4450_3030":
		return New0063_4450_3030()
	case "0063_4944_3034":
		return New0063_4944_3034()
	case "0063_4944_3131":
		return New0063_4944_3131()
	case "0063_4952_3032":
		return New0063_4952_3032()
	case "0063_6363_3030":
		return New0063_6363_3030()
	case "0063_4953_3133":
		return New0063_4953_3133()
	case "026e_4353_5a31":
		return New026E_4353_5A31()
	case "026e_5253_5a31":
		return New026E_5253_5A31()
	case "0099_0002_0002":
		return New0099_0002_0002()
	case "0099_0003_0003":
		return New0099_0003_0003()
	case "0099_0003_0004":
		return New0099_0003_0004()
	case "0208_0101_0005":
		return New0208_0101_0005()
	case "0208_0200_0009":
		return New0208_0200_0009()
	case "0208_0100_000a":
		return New0208_0100_000A()
	case "0260_0168_0168":
		return New0260_0168_0168()
	case "0260_8002_1000":
		return New0260_8002_1000()
	case "0260_8006_1000":
		return New0260_8006_1000()
	case "0293_0004_0000":
		return New0293_0004_0000()
	case "0293_0003_0014":
		return New0293_0003_0014()
	case "001e_0002_0002":
		return New001E_0002_0002()
	case "001e_0002_0001":
		return New001E_0002_0001()
	case "001e_0004_0001":
		return New001E_0004_0001()
	case "000c_4447_3034":
		return New000C_4447_3034()
	case "000c_4447_3033":
		return New000C_4447_3033()
	case "0039_0011_0001":
		return New0039_0011_0001()
	case "0059_0003_0001":
		return New0059_0003_0001()
	case "0059_0001_0003":
		return New0059_0001_0003()
	case "0059_0004_0001":
		return New0059_0004_0001()
	case "0059_000d_0002":
		return New0059_000D_0002()
	case "0059_0010_0001":
		return New0059_0010_0001()
	case "0059_0010_0002":
		return New0059_0010_0002()
	case "0059_0001_0004":
		return New0059_0001_0004()
	case "0059_0003_0002":
		return New0059_0003_0002()
	case "0059_000f_0001":
		return New0059_000F_0001()
	case "0230_0003_0001":
		return New0230_0003_0001()
	case "011f_0001_0002":
		return New011F_0001_0002()
	case "0077_0011_0001":
		return New0077_0011_0001()
	case "0077_0010_0001":
		return New0077_0010_0001()
	case "0005_4341_0600":
		return New0005_4341_0600()
	case "0005_4341_3750":
		return New0005_4341_3750()
	case "0005_4341_8900":
		return New0005_4341_8900()
	case "0005_0001_0003":
		return New0005_0001_0003()
	case "0005_0002_0003":
		return New0005_0002_0003()
	case "0005_0003_0003":
		return New0005_0003_0003()
	case "0005_0004_0003":
		return New0005_0004_0003()
	case "0005_0005_0003":
		return New0005_0005_0003()
	case "0005_4841_0020":
		return New0005_4841_0020()
	case "0005_4341_3000":
		return New0005_4341_3000()
	case "0214_0002_0002":
		return New0214_0002_0002()
	case "0114_b005_3b60":
		return New0114_B005_3B60()
	case "0090_0001_0001":
		return New0090_0001_0001()
	case "001d_3401_0001":
		return New001D_3401_0001()
	case "001d_3201_0001":
		return New001D_3201_0001()
	case "001d_1b03_0334":
		return New001D_1B03_0334()
	case "001d_1a02_0334":
		return New001D_1A02_0334()
	case "001d_3601_0001":
		return New001D_3601_0001()
	case "001d_3501_0001":
		return New001D_3501_0001()
	case "001d_1902_0334":
		return New001D_1902_0334()
	case "001d_1d04_0334":
		return New001D_1D04_0334()
	case "001d_1c02_0334":
		return New001D_1C02_0334()
	case "001d_0201_0206":
		return New001D_0201_0206()
	case "001d_1102_0243":
		return New001D_1102_0243()
	case "001d_0e01_0334":
		return New001D_0E01_0334()
	case "001d_1001_0334":
		return New001D_1001_0334()
	case "001d_1001_0209":
		return New001D_1001_0209()
	case "001d_0401_0209":
		return New001D_0401_0209()
	case "001d_0401_0334":
		return New001D_0401_0334()
	case "001d_0602_0334":
		return New001D_0602_0334()
	case "001d_0202_030b":
		return New001D_0202_030B()
	case "001d_1805_0334":
		return New001D_1805_0334()
	case "001d_1706_0334":
		return New001D_1706_0334()
	case "001d_0f01_0209":
		return New001D_0F01_0209()
	case "001d_0301_0209":
		return New001D_0301_0209()
	case "001d_0301_0334":
		return New001D_0301_0334()
	case "014f_5246_3133":
		return New014F_5246_3133()
	case "014f_4742_3030":
		return New014F_4742_3030()
	case "014f_4754_3038":
		return New014F_4754_3038()
	case "014f_4744_3030":
		return New014F_4744_3030()
	case "014f_4744_3530":
		return New014F_4744_3530()
	case "014f_4450_3030":
		return New014F_4450_3030()
	case "014f_5250_3030":
		return New014F_5250_3030()
	case "014f_5442_5431":
		return New014F_5442_5431()
	case "014f_2009_0903":
		return New014F_2009_0903()
	case "014f_2005_0503":
		return New014F_2005_0503()
	case "014f_2001_0102":
		return New014F_2001_0102()
	case "014f_2002_0203":
		return New014F_2002_0203()
	case "014f_4457_3034":
		return New014F_4457_3034()
	case "014f_5252_3530":
		return New014F_5252_3530()
	case "014f_5257_3533":
		return New014F_5257_3533()
	case "014f_5457_3033":
		return New014F_5457_3033()
	case "0234_0002_010a":
		return New0234_0002_010A()
	case "0234_0003_010a":
		return New0234_0003_010A()
	case "015f_0000_0000":
		return New015F_0000_0000()
	case "015f_0905_0201":
		return New015F_0905_0201()
	case "015f_4121_1302":
		return New015F_4121_1302()
	case "015f_0a05_0100":
		return New015F_0A05_0100()
	case "015f_0a05_0201":
		return New015F_0A05_0201()
	case "015f_0702_3102":
		return New015F_0702_3102()
	case "015f_0702_5102":
		return New015F_0702_5102()
	case "015f_0801_3102":
		return New015F_0801_3102()
	case "015f_3102_0202":
		return New015F_3102_0202()
	case "015f_3102_0204":
		return New015F_3102_0204()
	case "015f_3141_1302":
		return New015F_3141_1302()
	case "015f_4102_0201":
		return New015F_4102_0201()
	case "015f_4102_0202":
		return New015F_4102_0202()
	case "015f_4111_1302":
		return New015F_4111_1302()
	case "007a_0001_0002":
		return New007A_0001_0002()
	case "007a_8001_8001":
		return New007A_8001_8001()
	case "007a_8001_8002":
		return New007A_8001_8002()
	case "007a_8001_8003":
		return New007A_8001_8003()
	case "007a_4003_0002":
		return New007A_4003_0002()
	case "007a_0002_0001":
		return New007A_0002_0001()
	case "007a_4005_0001":
		return New007A_4005_0001()
	case "007a_4004_0001":
		return New007A_4004_0001()
	case "0178_4442_3130":
		return New0178_4442_3130()
	case "0165_0001_0001":
		return New0165_0001_0001()
	case "0165_0002_0001":
		return New0165_0002_0001()
	case "0165_0002_0003":
		return New0165_0002_0003()
	case "0165_0001_0003":
		return New0165_0001_0003()
	case "0165_0002_0002":
		return New0165_0002_0002()
	case "0096_0010_0001":
		return New0096_0010_0001()
	case "0096_0001_0001":
		return New0096_0001_0001()
	case "0096_0001_0002":
		return New0096_0001_0002()
	case "013c_0006_001a":
		return New013C_0006_001A()
	case "013c_0005_0031":
		return New013C_0005_0031()
	case "013c_0001_000f":
		return New013C_0001_000F()
	case "013c_0001_0003":
		return New013C_0001_0003()
	case "013c_0001_0012":
		return New013C_0001_0012()
	case "013c_0001_0014":
		return New013C_0001_0014()
	case "013c_0001_0004":
		return New013C_0001_0004()
	case "013c_0001_0013":
		return New013C_0001_0013()
	case "013c_0001_0006":
		return New013C_0001_0006()
	case "013c_0001_0015":
		return New013C_0001_0015()
	case "013c_0001_0001":
		return New013C_0001_0001()
	case "013c_0001_0011":
		return New013C_0001_0011()
	case "013c_0001_0030":
		return New013C_0001_0030()
	case "013c_0002_001f":
		return New013C_0002_001F()
	case "013c_0002_0020":
		return New013C_0002_0020()
	case "013c_0002_0021":
		return New013C_0002_0021()
	case "013c_0004_000a":
		return New013C_0004_000A()
	case "013c_0002_0002":
		return New013C_0002_0002()
	case "013c_0009_0022":
		return New013C_0009_0022()
	case "013c_0002_000c":
		return New013C_0002_000C()
	case "013c_0002_000d":
		return New013C_0002_000D()
	case "013c_0002_000e":
		return New013C_0002_000E()
	case "010e_0008_0002":
		return New010E_0008_0002()
	case "0154_0100_0201":
		return New0154_0100_0201()
	case "0154_0004_0003":
		return New0154_0004_0003()
	case "0154_0000_0000":
		return New0154_0000_0000()
	case "0154_0004_0002":
		return New0154_0004_0002()
	case "0154_0100_0400":
		return New0154_0100_0400()
	case "0154_0100_0103":
		return New0154_0100_0103()
	case "0154_0100_0101":
		return New0154_0100_0101()
	case "0154_0004_0004":
		return New0154_0004_0004()
	case "0154_0005_0001":
		return New0154_0005_0001()
	case "0154_0003_0001":
		return New0154_0003_0001()
	case "0154_1100_0002":
		return New0154_1100_0002()
	case "0154_1100_0001":
		return New0154_1100_0001()
	case "0154_0001_0001":
		return New0154_0001_0001()
	case "0154_0004_0011":
		return New0154_0004_0011()
	case "0128_0000_0000":
		return New0128_0000_0000()
	case "0095_3103_0001":
		return New0095_3103_0001()
	case "0095_0002_0001":
		return New0095_0002_0001()
	case "0159_0002_0002":
		return New0159_0002_0002()
	case "0159_0002_0052":
		return New0159_0002_0052()
	case "0159_0002_0001":
		return New0159_0002_0001()
	case "0159_0002_0051":
		return New0159_0002_0051()
	case "0159_0003_0002":
		return New0159_0003_0002()
	case "0159_0003_0052":
		return New0159_0003_0052()
	case "0159_0001_0001":
		return New0159_0001_0001()
	case "0159_0001_0051":
		return New0159_0001_0051()
	case "0159_0005_0001":
		return New0159_0005_0001()
	case "0159_0005_0051":
		return New0159_0005_0051()
	case "0159_0004_0001":
		return New0159_0004_0001()
	case "0159_0004_0051":
		return New0159_0004_0051()
	case "0159_0005_0003":
		return New0159_0005_0003()
	case "0159_0005_0053":
		return New0159_0005_0053()
	case "0159_0002_0053":
		return New0159_0002_0053()
	case "0159_0003_0053":
		return New0159_0003_0053()
	case "0159_0001_0052":
		return New0159_0001_0052()
	case "0159_0001_0053":
		return New0159_0001_0053()
	case "0159_0001_0054":
		return New0159_0001_0054()
	case "0159_0007_0053":
		return New0159_0007_0053()
	case "0010_5442_5432":
		return New0010_5442_5432()
	case "0010_0001_0002":
		return New0010_0001_0002()
	case "0064_2001_0000":
		return New0064_2001_0000()
	case "0064_0001_0000":
		return New0064_0001_0000()
	case "0064_3001_0000":
		return New0064_3001_0000()
	case "0064_3002_0000":
		return New0064_3002_0000()
	case "0064_5002_0000":
		return New0064_5002_0000()
	case "0064_4001_0000":
		return New0064_4001_0000()
	case "0064_1001_0000":
		return New0064_1001_0000()
	case "0064_5003_0000":
		return New0064_5003_0000()
	case "0064_0166_2007":
		return New0064_0166_2007()
	case "0064_1000_0009":
		return New0064_1000_0009()
	case "5254_8000_0002":
		return New5254_8000_0002()
	case "5254_0001_8380":
		return New5254_0001_8380()
	case "5254_0002_8380":
		return New5254_0002_8380()
	case "5254_0000_8510":
		return New5254_0000_8510()
	case "5254_0001_8510":
		return New5254_0001_8510()
	case "5254_0002_8510":
		return New5254_0002_8510()
	case "5254_0200_8031":
		return New5254_0200_8031()
	case "5254_0202_8031":
		return New5254_0202_8031()
	case "5254_0102_8377":
		return New5254_0102_8377()
	case "5254_0101_8377":
		return New5254_0101_8377()
	case "5254_0100_8377":
		return New5254_0100_8377()
	case "0147_0400_0001":
		return New0147_0400_0001()
	case "0147_0400_0002":
		return New0147_0400_0002()
	case "0098_6401_0015":
		return New0098_6401_0015()
	case "0098_6401_0105":
		return New0098_6401_0105()
	case "0098_6401_0106":
		return New0098_6401_0106()
	case "0098_6401_0107":
		return New0098_6401_0107()
	case "0098_6401_01fd":
		return New0098_6401_01FD()
	case "0098_6501_000b":
		return New0098_6501_000B()
	case "0098_6501_000c":
		return New0098_6501_000C()
	case "0098_6e02_0101":
		return New0098_6E02_0101()
	case "0098_0001_001e":
		return New0098_0001_001E()
	case "0098_1e10_0158":
		return New0098_1E10_0158()
	case "0098_1e12_015e":
		return New0098_1E12_015E()
	case "0098_3200_015e":
		return New0098_3200_015E()
	case "0098_1e12_015c":
		return New0098_1E12_015C()
	case "0098_0000_0000":
		return New0098_0000_0000()
	case "0098_2002_0100":
		return New0098_2002_0100()
	case "0098_2002_0102":
		return New0098_2002_0102()
	case "0098_0002_0100":
		return New0098_0002_0100()
	case "003b_634b_5044":
		return New003B_634B_5044()
	case "003b_6349_5044":
		return New003B_6349_5044()
	case "003b_6341_5044":
		return New003B_6341_5044()
	case "003b_634b_504c":
		return New003B_634B_504C()
	case "0301_0048_010f":
		return New0301_0048_010F()
	case "019a_0003_0003":
		return New019A_0003_0003()
	case "0258_0003_3082":
		return New0258_0003_3082()
	case "0258_0003_1082":
		return New0258_0003_1082()
	case "0258_0003_1085":
		return New0258_0003_1085()
	case "0258_0003_1083":
		return New0258_0003_1083()
	case "0258_0003_3083":
		return New0258_0003_3083()
	case "0258_0003_0082":
		return New0258_0003_0082()
	case "0258_0003_1087":
		return New0258_0003_1087()
	case "0258_0003_1088":
		return New0258_0003_1088()
	case "0000_0003_0003":
		return New0000_0003_0003()
	case "0047_5a52_5400":
		return New0047_5A52_5400()
	case "0047_5a52_5401":
		return New0047_5A52_5401()
	case "0047_5a52_5402":
		return New0047_5A52_5402()
	case "0047_5a52_5403":
		return New0047_5A52_5403()
	case "0047_5a52_5404":
		return New0047_5A52_5404()
	case "0047_5a52_5405":
		return New0047_5A52_5405()
	case "0047_5a52_5406":
		return New0047_5A52_5406()
	case "0047_5a52_5407":
		return New0047_5A52_5407()
	case "0047_5a52_5408":
		return New0047_5A52_5408()
	case "0047_5a52_5409":
		return New0047_5A52_5409()
	case "0047_5a52_540a":
		return New0047_5A52_540A()
	case "0047_5a52_540b":
		return New0047_5A52_540B()
	case "0047_5a52_540c":
		return New0047_5A52_540C()
	case "0047_5a52_540d":
		return New0047_5A52_540D()
	case "0047_5a52_540e":
		return New0047_5A52_540E()
	case "0047_5a52_540f":
		return New0047_5A52_540F()
	case "0047_5a52_5410":
		return New0047_5A52_5410()
	case "0239_0001_0001":
		return New0239_0001_0001()
	case "0276_0139_0001":
		return New0276_0139_0001()
	case "0176_0005_0001":
		return New0176_0005_0001()
	case "0176_0003_0001":
		return New0176_0003_0001()
	case "019b_0001_0001":
		return New019B_0001_0001()
	case "0118_0002_0002":
		return New0118_0002_0002()
	case "0118_0001_0006":
		return New0118_0001_0006()
	case "0118_0311_0203":
		return New0118_0311_0203()
	case "0118_0003_0004":
		return New0118_0003_0004()
	case "0118_0004_0003":
		return New0118_0004_0003()
	case "0118_0808_0808":
		return New0118_0808_0808()
	case "0118_0102_1020":
		return New0118_0102_1020()
	case "0118_0003_0008":
		return New0118_0003_0008()
	case "0118_0202_0611":
		return New0118_0202_0611()
	case "0118_0003_0002":
		return New0118_0003_0002()
	case "0118_0101_0103":
		return New0118_0101_0103()
	case "0118_0004_0002":
		return New0118_0004_0002()
	case "0118_0001_0001":
		return New0118_0001_0001()
	case "0118_0001_0011":
		return New0118_0001_0011()
	case "008b_5452_5435":
		return New008B_5452_5435()
	case "008b_5452_5431":
		return New008B_5452_5431()
	case "008b_5452_5433":
		return New008B_5452_5433()
	case "008b_5452_5436":
		return New008B_5452_5436()
	case "008b_5452_5439":
		return New008B_5452_5439()
	case "008b_5452_5443":
		return New008B_5452_5443()
	case "008b_5452_5442":
		return New008B_5452_5442()
	case "0152_0003_0512":
		return New0152_0003_0512()
	case "0152_0500_0004":
		return New0152_0500_0004()
	case "0152_0500_0003":
		return New0152_0500_0003()
	case "0109_2007_0703":
		return New0109_2007_0703()
	case "0109_201a_1aa4":
		return New0109_201A_1AA4()
	case "0109_2001_0102":
		return New0109_2001_0102()
	case "0109_2001_0105":
		return New0109_2001_0105()
	case "0109_2001_0106":
		return New0109_2001_0106()
	case "0109_2022_2201":
		return New0109_2022_2201()
	case "0109_201f_1f10":
		return New0109_201F_1F10()
	case "0109_200a_0a02":
		return New0109_200A_0A02()
	case "0109_2008_0801":
		return New0109_2008_0801()
	case "0109_200c_0c02":
		return New0109_200C_0C02()
	case "0109_200c_0c06":
		return New0109_200C_0C06()
	case "0109_2016_1616":
		return New0109_2016_1616()
	case "0109_2017_1717":
		return New0109_2017_1717()
	case "0109_2005_0503":
		return New0109_2005_0503()
	case "0109_2005_0508":
		return New0109_2005_0508()
	case "0109_2009_0903":
		return New0109_2009_0903()
	case "0109_2009_0908":
		return New0109_2009_0908()
	case "0109_2006_0620":
		return New0109_2006_0620()
	case "0109_2006_0621":
		return New0109_2006_0621()
	case "0109_2002_0201":
		return New0109_2002_0201()
	case "0109_2002_0202":
		return New0109_2002_0202()
	case "0109_2002_0203":
		return New0109_2002_0203()
	case "0109_2002_0204":
		return New0109_2002_0204()
	case "0109_2002_0205":
		return New0109_2002_0205()
	case "0109_2003_0302":
		return New0109_2003_0302()
	case "0109_2003_0306":
		return New0109_2003_0306()
	case "0109_2003_0307":
		return New0109_2003_0307()
	case "0109_2004_0403":
		return New0109_2004_0403()
	case "0109_200f_0f03":
		return New0109_200F_0F03()
	case "0109_2021_2101":
		return New0109_2021_2101()
	case "0109_200d_0d03":
		return New0109_200D_0D03()
	case "0149_0012_0104":
		return New0149_0012_0104()
	case "0149_1214_0504":
		return New0149_1214_0504()
	case "0149_1214_0304":
		return New0149_1214_0304()
	case "015d_0651_f51c":
		return New015D_0651_F51C()
	case "015d_0111_231c":
		return New015D_0111_231C()
	case "015d_1111_1e1c":
		return New015D_1111_1E1C()
	case "015d_0111_1f1c":
		return New015D_0111_1F1C()
	case "015d_0221_251c":
		return New015D_0221_251C()
	case "0097_1122_5501":
		return New0097_1122_5501()
	case "0097_6943_5501":
		return New0097_6943_5501()
	case "0097_1200_5502":
		return New0097_1200_5502()
	case "0097_6943_4501":
		return New0097_6943_4501()
	case "0097_0024_0045":
		return New0097_0024_0045()
	case "0097_1182_4501":
		return New0097_1182_4501()
	case "0097_1188_4501":
		return New0097_1188_4501()
	case "0097_6941_5501":
		return New0097_6941_5501()
	case "0097_1180_5501":
		return New0097_1180_5501()
	case "0097_6131_4501":
		return New0097_6131_4501()
	case "0129_0006_0000":
		return New0129_0006_0000()
	case "0129_0004_0800":
		return New0129_0004_0800()
	case "0129_0002_0800":
		return New0129_0002_0800()
	case "0129_0004_0000":
		return New0129_0004_0000()
	case "0129_0004_aa00":
		return New0129_0004_AA00()
	case "0129_0002_0000":
		return New0129_0002_0000()
	case "0129_0002_0209":
		return New0129_0002_0209()
	case "0129_0002_aa00":
		return New0129_0002_AA00()
	case "0129_0002_ffff":
		return New0129_0002_FFFF()
	case "0129_0001_0409":
		return New0129_0001_0409()
	case "0129_842a_3cac":
		return New0129_842A_3CAC()
	case "0131_0002_0002":
		return New0131_0002_0002()
	case "027a_0101_000a":
		return New027A_0101_000A()
	case "027a_0003_0087":
		return New027A_0003_0087()
	case "027a_0003_0082":
		return New027A_0003_0082()
	case "027a_0083_0003":
		return New027A_0083_0003()
	case "027a_0003_0000":
		return New027A_0003_0000()
	case "027a_0003_0085":
		return New027A_0003_0085()
	case "027a_2021_2101":
		return New027A_2021_2101()
	case "0115_1000_0003":
		return New0115_1000_0003()
	case "0115_0100_0001":
		return New0115_0100_0001()
	case "0115_0100_0002":
		return New0115_0100_0002()
	case "0115_0000_0000":
		return New0115_0000_0000()
	case "0115_0100_0101":
		return New0115_0100_0101()
	case "0115_0100_0004":
		return New0115_0100_0004()
	case "0115_1000_0001":
		return New0115_1000_0001()
	case "0115_1000_0002":
		return New0115_1000_0002()
	case "0115_1000_0009":
		return New0115_1000_0009()
	case "0115_1000_0100":
		return New0115_1000_0100()
	case "0115_1000_0200":
		return New0115_1000_0200()
	case "0115_0400_0004":
		return New0115_0400_0004()
	case "0115_1000_0004":
		return New0115_1000_0004()
	case "0115_0024_0001":
		return New0115_0024_0001()
	case "0115_0100_0102":
		return New0115_0100_0102()
	case "0115_0100_0103":
		return New0115_0100_0103()
	case "0115_1000_0300":
		return New0115_1000_0300()
	case "0115_0110_0001":
		return New0115_0110_0001()
	case "0115_0100_0400":
		return New0115_0100_0400()
	}

	return nil
}
func New0081_00A0_0001() *Device{
	return &Device{
		Brand: "SIEGENIA-AUBI KG",
		Product: "Sensoair",
		Description: "Air quality sensor for indoor use",

		ManufacturerID: "0081",
		ProductType: "00A0",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Device Configuration Value 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 143,
					},
				},
			},
		},
	}
}
func New0068_0000_0010() *Device{
	return &Device{
		Brand: "Good Way Technology Co., Ltd",
		Product: "TD1311",
		Description: "In wall Power Monitor Switch",

		ManufacturerID: "0068",
		ProductType: "0000",
		ProductID: "0010",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New010A_1100_4B00() *Device{
	return &Device{
		Brand: "VDA",
		Product: "Vitrum III EU Dimmer",
		Description: "",

		ManufacturerID: "010A",
		ProductType: "1100",
		ProductID: "4B00",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x91,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0001_4450_3030() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "45602",
		Description: "Lamp Module with Dimmer Control",

		ManufacturerID: "0001",
		ProductType: "4450",
		ProductID: "3030",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 5,
				Name: "Ignore Start Level When Receiving Dim Commands",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Respect Start Level",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Ignore Start Level (default)",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dim Rate Step Count - ZWave Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dim Rate Step Duration - ZWave Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Dim Rate Step Count - Manual Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Dim Rate Step Duration - manual Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Dim Rate Step Count - ZWave All on/off",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Dim Rate Step Duration - ZWave All on/off",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Load Sensing",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled (default)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
		},
	}
}
func New0001_444D_3330() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZDM230",
		Description: "HomePro Wall Dimmer ZDM230",

		ManufacturerID: "0001",
		ProductType: "444D",
		ProductID: "3330",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Ignore Start-Level (Transmitting)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Suspend group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
		},
	}
}
func New0001_4457_3033() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZDW103",
		Description: "Wall Mounted 3-Way Dimmer Receiver",

		ManufacturerID: "0001",
		ProductType: "4457",
		ProductID: "3033",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set Ignore Start Level Bit When Transmitting Dim Commands",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Use transmisison start level",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Ignore transmisison start level",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Suspend Group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable transmission to Group 4",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable transmission to Group 4",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED on when switch is ON",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED on when switch is OFF",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Invert Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Ignore Start-Level (Receiving)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Don’t Send Level Command After Transmitting Dim Commands",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send level command",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Don’t send level command",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "On/Off Command Dim Step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "On/Off Command Dim Rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Local Control Dim Step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Local Control Dim Rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "ALL ON/ALL OFF Dim Step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "ALL ON/ALL OFF Dim Rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Disable Group 4 During a Dim Command",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable transmission to Group 4",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable transmission to Group 4",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Enable Shade Control Group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Enable Shade Control Group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "LED Transmission Indication",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No flicker",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Flicker during all transmission",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flicker first second of transmission",
					},
				},
			},
		},
	}
}
func New0001_4457_3230() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZDW120",
		Description: "Two Wire Wall Dimer",

		ManufacturerID: "0001",
		ProductType: "4457",
		ProductID: "3230",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set Ignore Start Level Bit When Transmitting Dim Commands",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Use transmisison start level",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Ignore transmisison start level",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Suspend Group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable transmission to Group 4",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable transmission to Group 4",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED on when switch is OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED on when switch is ON",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Invert Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Ignore Start-Level (Receiving)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Don’t Send Level Command After Transmitting Dim Commands",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send level command",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Don’t send level command",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "On/Off Command Dim Step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "On/Off Command Dim Rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Local Control Dim Step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Local Control Dim Rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "ALL ON/ALL OFF Dim Step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "ALL ON/ALL OFF Dim Rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Disable Group 4 During a Dim Command",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable transmission to Group 4",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable transmission to Group 4",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "LED Transmission Indication",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No flicker",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Flicker during all transmission",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flicker first second of transmission",
					},
				},
			},
		},
	}
}
func New0001_4952_3130() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZIR010",
		Description: "PIR Motion Sensor",

		ManufacturerID: "0001",
		ProductType: "4952",
		ProductID: "3130",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 17,
				Name: "Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Lightning",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Alarm",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Sensor",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Off/Idle delay",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Unsolicited Commands",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Awake Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 15,
						To: 45,
					},
				},
			},
		},
	}
}
func New0001_524D_3330() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZRM230",
		Description: "Wall Switch/Transmitter (2-gang)",

		ManufacturerID: "0001",
		ProductType: "524D",
		ProductID: "3330",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Ignore start level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore start level",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Do not ignore start level",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Suspend Group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Transmit to group 4 devices",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Stop transmit to group 4 devices",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Led when load on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Led when load off",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Invert Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Invert off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Invert on",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Enable Shade Control Group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Enable Shade Control Group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "LED Transmission Indication",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No flicker",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Flicker for 1 second",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flicker for 2 second",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Poll Group 2 Interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Poll Group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0001_7FFF_7FFF() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZRP200",
		Description: "HomePro Applicance Module ZRP200",

		ManufacturerID: "0001",
		ProductType: "7FFF",
		ProductID: "7FFF",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0001_5257_3330() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZRW230",
		Description: "Wall Mounted 3-Way Switch",

		ManufacturerID: "0001",
		ProductType: "5257",
		ProductID: "3330",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set Ignore Start Level Bit",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Dimmer will not ignore the start level",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Dimmer will ignore the start level",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Suspend Group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable transmitting to devices “associated” into Group 4",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Stop transmitting to devices “associated” into Group 4",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Turn LED ON when the load is turned ON",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn LED ON when the load is turned OFF",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Invert Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Top of switch is ON, bottom is OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Top of switch is OFF, bottom is ON",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Enable Shade Control Group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable operation of shade control devices via group 2",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable operation of shade control devices via group 2",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Enable Shade Control Group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable operation of shade control devices via group 3",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable operation of shade control devices via group 3",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "LED Transmission Indication",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED will not flicker",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED will flicker the entire time while transmitting",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "LED will flicker 1 second when it begins transmitting",
					},
				},
			},
		},
	}
}
func New0001_544D_3330() *Device{
	return &Device{
		Brand: "ACT - Advanced Control Technologies",
		Product: "ZTM230",
		Description: "Dual paddle wall mounted transmitter",

		ManufacturerID: "0001",
		ProductType: "544D",
		ProductID: "3330",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Ignore Start Level Bit",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not ignore start level",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Ignore start level",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Suspend Group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Invert Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Ignore Start Level Bit",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "On/Off Command dim step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "On/Off Command dim timer",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Local Control dim rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Dim timer Parameter",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "ALL ON/ALL OFF dim step",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "ALL ON/ALL OFF dim rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Disable group 4 during dim command",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Enable shade control group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Enable shade control group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Enable shade control group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Transmission LED",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 2,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Poll group 1 interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Poll group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Poll Group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0086_0002_0078() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW120",
		Description: "Door/Window sensor Gen5",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0078",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Which value of the Sensor Binary Report or Basic Set will be sent when the Magnet is triggered On/Off.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On=Sensor Binary Report/Basic Set 0xFF; Off=Sensor Binary Report/Basic",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "On=Sensor Binary Report/Basic Set 0x00; Off=Sensor Binary Report/Basic",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable wake up 10 minutes when re-power on the sensor",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Basic Set Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Open: 0xFF, Close: 0x00",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Open: 0x00, Close: 0xFF",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Set the low battery value.",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Set the interval time of low battery checking.",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Report type",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset To Factory Defaults.",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset to factory default setting",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Normal",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset to factory default",
					},
				},
			},
		},
	}
}
func New0086_0102_0078() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW120",
		Description: "Door/Window sensor Gen5",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "0078",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Which value of the Sensor Binary Report or Basic Set will be sent when the Magnet is triggered On/Off.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On=Sensor Binary Report/Basic Set 0xFF; Off=Sensor Binary Report/Basic",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "On=Sensor Binary Report/Basic Set 0x00; Off=Sensor Binary Report/Basic",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable wake up 10 minutes when re-power on the sensor",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Basic Set Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Open: 0xFF, Close: 0x00",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Open: 0x00, Close: 0xFF",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Set the low battery value.",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Set the interval time of low battery checking.",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Report type",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset To Factory Defaults.",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset to factory default setting",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Normal",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset to factory default",
					},
				},
			},
		},
	}
}
func New0086_0001_0003() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSA03202",
		Description: "Minimote 4 button remote control",

		ManufacturerID: "0086",
		ProductType: "0001",
		ProductID: "0003",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 241,
				Name: "Mode of Button 1 (upper left)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 242,
				Name: "Mode of Button 2 (upper right)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 243,
				Name: "Mode of Button 3 (lower left)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 244,
				Name: "Mode of Button 4 (lower right)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 245,
				Name: "Mode of Button 5 (Include)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 246,
				Name: "Mode of Button 6 (Exclude)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 247,
				Name: "Mode of Button 7 (Association)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 248,
				Name: "Mode of Button 8 (Learn)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 250,
				Name: "Secondary Controller Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Group Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
				},
			},
		},
	}
}
func New0086_0006_0002() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSA03202",
		Description: "Minimote 4 button remote control",

		ManufacturerID: "0086",
		ProductType: "0006",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 241,
				Name: "Mode of Button 1 (upper left)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 242,
				Name: "Mode of Button 2 (upper right)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 243,
				Name: "Mode of Button 3 (lower left)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 244,
				Name: "Mode of Button 4 (lower right)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 245,
				Name: "Mode of Button 5 (Include)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 246,
				Name: "Mode of Button 6 (Exclude)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 247,
				Name: "Mode of Button 7 (Association)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 248,
				Name: "Mode of Button 8 (Learn)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Add Mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Remove Mode",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Association Mode",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Learn Mode",
					},
				},
			},
			&parameter{
				ID: 250,
				Name: "Secondary Controller Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Group Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
				},
			},
		},
	}
}
func New0086_0001_0026() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSA38",
		Description: "Panic Button Key Fob",

		ManufacturerID: "0086",
		ProductType: "0001",
		ProductID: "0026",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 250,
				Name: "Mode of Button 1 ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Factory Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
				},
			},
		},
	}
}
func New0086_0002_0005() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSB05",
		Description: "4 in One MultiSensor",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0005",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Wake up for 10 minutes when batteries are inserted",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable (default)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "On time",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 15300,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion sensor",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Command to send when movement",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "send Sensor Binary report",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports to send to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 225,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 225,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 225,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to the default Configuration of all parameters",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0086_0002_0009() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSB09",
		Description: "Home Energy Meter",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0009",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Voltage",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32000,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Current Threshold - Whole HEM",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Current Threshold - Clamp 1",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Current Threshold - Clamp 2",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Current Threshold - Clamp 3",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Percent Change - Whole HEM",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Percent Change - Clamp 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Percent Change - Clamp 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Percent Change - Clamp 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
		},
	}
}
func New0086_0002_001C() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSB28",
		Description: "Home Energy Meter G2",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "001C",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Reverse clamping",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 16,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Current Threshold - Whole HEM",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Current Threshold - Clamp 1",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Current Threshold - Clamp 2",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Current Threshold - Clamp 3",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Percent Change - Whole HEM",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Percent Change - Clamp 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Percent Change - Clamp 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Percent Change - Clamp 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1061109568,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1061109568,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1061109568,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
		},
	}
}
func New0086_0002_0004() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSB29",
		Description: "Door/Window sensor Gen2",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0004",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable wake up 10 minutes when power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "command class to use for open/close",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "not set",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "BATTERY report",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "SENSOR BINARY report",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "SENSOR BINARY and BATTERY report",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "command class to use for open/close",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "not set",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "BASIC report",
					},
				},
			},
		},
	}
}
func New0086_0002_001D() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSB29",
		Description: "Door/Window sensor Gen2",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "001D",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable wake up 10 minutes when power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "command class to use for open/close",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "not set",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "BATTERY report",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "SENSOR BINARY report",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "SENSOR BINARY and BATTERY report",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "command class to use for open/close",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "not set",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "BASIC report",
					},
				},
			},
		},
	}
}
func New0086_0002_002D() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSB45",
		Description: "Water Sensor",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "002D",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Sensor Binary Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Open: 00, Close: FF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Open: FF, Close: 00",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Wake up",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Basic set value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Open: 00, Close: FF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Open: FF, Close: 00",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Reports that will be sent",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not send anything",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send battery report",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Send Sensor Binary report",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Send Sensor Binary and Battery reports",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Parameter 121 Value 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not send anything",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Send ALARM",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Send Basic Set and ALARM",
					},
				},
			},
		},
	}
}
func New0086_0002_0036() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSB54",
		Description: "Recessed Door/Window Sensor",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0036",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Flag values for triggered magnet switch",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Permit other configurations",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0003_0006() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC06",
		Description: "Smart Energy Switch",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0006",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Multilevel Sensor Report Content",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Power",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Automatic Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable/disable Parameter 91 and 92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum change in wattage",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum change in wattage %",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 100 to 103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Group 1 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "Meter Report Watts and kWh",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Group 2 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "Meter Report Watts and kWh",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Group 3 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "Meter Report Watts and kWh",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111 to 113 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Device Reset",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0003_0008() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC08",
		Description: "Smart Energy Illuminator",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0008",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Multilevel Sensor Report Content",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Power",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Make SES blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Automatic Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable Parameter 91",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum change in wattage",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum change in wattage %",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 100 to 103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Group 1 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Group 2 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Group 3 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111 to 113 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Device Reset",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0003_000A() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC10",
		Description: "Heavy Duty Smart Switch",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "000A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Activate Overload Protection",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Deactivate Overload Protection",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output Load Status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always OFF",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enables/disables parameter 91/92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable parameters 91 and 92",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable parameters 91 and 92",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Aeon Labs",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Other",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock/unlock Configuration Changes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlocked",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Locked",
					},
				},
			},
		},
	}
}
func New0086_0003_000B() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC11",
		Description: "Smart Strip",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "000B",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Voltage to calculate power",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32000,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 16,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Wattage Threshold - Whole strip",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Wattage Threshold - Socket 1",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Wattage Threshold - Socket 2",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Wattage Threshold - Socket 3",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Wattage Threshold - Socket 4",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Wattage Threshold - Socket 5",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Wattage Threshold - Socket 6",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Percent Change - Whole Strip",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Percent Change - Socket 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Percent Change - Socket 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Percent Change - Socket 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Percent Change - Socket 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "Percent Change - Socket 5",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Percent Change - Socket 6",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Status of load changed report",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Not sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Basic",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Get Temperature",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Reset 0x65~0x67 to default value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "select channels to report consumtion as item 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Select channels to report power as item 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Item 2 (kWh)",
				Size: 4,
				Values: []Value{
					Value{
						From: 255,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Item 2 (power)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
		},
	}
}
func New0086_0003_000C() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC12",
		Description: "Micro Smart Energy Switch",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "000C",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Make Micro Smart Switch Blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Automated sending of a Report triggered by minimal change of value.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deavtivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically in timing intervals for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically in timing intervals for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically in timing intervals for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
		},
	}
}
func New0086_0003_000D() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC13",
		Description: "Micro Smart Energy Illuminator",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "000D",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Type of Sensor Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Power",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Blinking Behavior",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Disable Minimum Change report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Reports to send to Group 1 ",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Reports to send to Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Reports to send to Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New0086_0003_000E() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC14",
		Description: "Micro Motor Controller",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "000E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Type of Sensor Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Power",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Blinking Behavior",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Disable automated reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "reports to send automatically",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports to send automatically to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports to send automatically to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New0086_0003_0011() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC17",
		Description: "Micro Double Smart Switch",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0011",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enables/disables parameter 91 and 92 ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deavtivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 101-103 to default.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Automatic reports for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Automatic reports for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Automatic reports for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111-113 to default.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Set External Switch/Button Control mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary button mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2-state switch mode",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Unidentified mode",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to the default Configuration.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
		},
	}
}
func New0086_0003_0012() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC18",
		Description: "Micro Smart Energy Switch G2",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0012",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Make Micro Smart Switch Blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Automated sending of a Report triggered by minimal change of value.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deavtivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically in timing intervals for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically in timing intervals for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically in timing intervals for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
		},
	}
}
func New0086_0003_0013() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC19",
		Description: "Micro Smart Energy Illuminator G2",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0013",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Make Micro Smart Dimmer 2nd Edition Blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Enable/Disable CRC16 encapsulation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Sending of a Report triggered by minimal change of value.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deavtivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Reports for Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Reports for Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Reports for Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
		},
	}
}
func New0086_0003_0018() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC24",
		Description: "Smart Energy Switch G2",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0018",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Multilevel Sensor Report Content",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Make SES blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Automatic Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable Parameter 91",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum change in wattage",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum change in wattage %",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 100 to 103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Group 1 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Group 2 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Group 3 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111 to 113 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Device Reset",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0103_004B() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC24",
		Description: "Smart Energy Switch G2",

		ManufacturerID: "0086",
		ProductType: "0103",
		ProductID: "004B",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Multilevel Sensor Report Content",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Make SES blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Automatic Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable Parameter 91",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum change in wattage",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum change in wattage %",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 100 to 103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Group 1 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Group 2 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Group 3 report",
				Size: 4,
				Values: []Value{
					Value{
						From: 2,
						To: 2,
						Desc: "Multisensor Report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Meter Report (Watts)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Meter Report (kWh)",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111 to 113 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Device Reset",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0003_0019() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC25",
		Description: "Smart Energy Illuminator G2",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0019",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0086_0003_001A() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC26",
		Description: "Micro Switch G2",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "001A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Make Micro Switch 2nd Edition Blink",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically in timing intervals for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically in timing intervals for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically in timing intervals for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Aeon Labs Standard Product",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "AT&T",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Lock Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset configuration set up to default setting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0086_0003_001B() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC27",
		Description: "Micro Illuminator G2",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "001B",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
		},
	}
}
func New0086_0003_0023() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSC35",
		Description: "AEON Labs Micro Double Switch",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0023",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0086_0004_0050() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSD31",
		Description: "Outlet Plugable Siren",

		ManufacturerID: "0086",
		ProductType: "0004",
		ProductID: "0050",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 37,
				Name: "Sirensound and Volume",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Siren Sound",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Siren Sound 1",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Siren Sound 2",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Siren Sound 3",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Siren Sound 4",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Siren Sound 5",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Siren Volume",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "88dB",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "100dB",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "105dB",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Send Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/Disable 'Lock' - Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
		},
	}
}
func New0086_0004_0025() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSD37",
		Description: "Range Extender",

		ManufacturerID: "0086",
		ProductType: "0004",
		ProductID: "0025",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0086_0019_0004() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "DSD37",
		Description: "Range Extender",

		ManufacturerID: "0086",
		ProductType: "0019",
		ProductID: "0004",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0086_0004_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "0004",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_0104_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "0104",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_0204_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "0204",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_0304_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "0304",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_0904_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "0904",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_0A04_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "0A04",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_1A04_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "1A04",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_1D04_0038() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW056",
		Description: "Doorbell",

		ManufacturerID: "0086",
		ProductType: "1D04",
		ProductID: "0038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Repetitions of ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Default ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Ringtone to play",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "play the default ringtone",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control items",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Ignore",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Play",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Stop",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Pause",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Next",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Previous",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Ringtone volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Define button- and button+ action",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "“Button ‐” is previous, “Button+” is next",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "“Button ‐” is next, “Button+” is previous",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Define long press button- and button+",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "short press changes vol, long press changes default ringtone",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "long press changes vol, short press changes default ringtone",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Get wireless button battery status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Battery Level Good",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Low Battery",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Enable/disable battery group notification",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration report parameter 42",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair doorbell with button",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New0086_0003_003E() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW062",
		Description: "Aeon Labs Garage Door Controller  Gen5 ",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "003E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x66,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 32,
				Name: "Startup ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Calibration not active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Begin calibration",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Calibration timout",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "Number of alarm musics",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Disable opening alarm",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable alarm prompt",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable alarm prompt",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm volume",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm choice",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm LED mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Disable closing alarm",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable alarm prompt",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable alarm prompt",
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm volume",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm choice",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm LED mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Unknown state alarm mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Closed alarm mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Tamper switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Battery state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 16,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Play or Pause ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Ringtone test volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "Temperature",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 500,
					},
				},
			},
			&parameter{
				ID: 47,
				Name: "Button definition",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mode 0",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Mode 1",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Door state change report type",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Send hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send barrier operator report CC",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair the Sensor",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop sensor pairing",
					},
					Value{
						From: 1431655681,
						To: 1431655681,
						Desc: "Start sensor pairing",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Configuration enabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration disabled (locked)",
					},
				},
			},
		},
	}
}
func New0086_0103_003E() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW062",
		Description: "Aeon Labs Garage Door Controller  Gen5 ",

		ManufacturerID: "0086",
		ProductType: "0103",
		ProductID: "003E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x66,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 32,
				Name: "Startup ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Calibration not active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Begin calibration",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Calibration timout",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "Number of alarm musics",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Disable opening alarm",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable alarm prompt",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable alarm prompt",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm volume",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm choice",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm LED mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Disable closing alarm",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable alarm prompt",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable alarm prompt",
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm volume",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm choice",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm LED mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Unknown state alarm mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Closed alarm mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Tamper switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Battery state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 16,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Play or Pause ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Ringtone test volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "Temperature",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 500,
					},
				},
			},
			&parameter{
				ID: 47,
				Name: "Button definition",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mode 0",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Mode 1",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Door state change report type",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Send hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send barrier operator report CC",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair the Sensor",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop sensor pairing",
					},
					Value{
						From: 1431655681,
						To: 1431655681,
						Desc: "Start sensor pairing",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Configuration enabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration disabled (locked)",
					},
				},
			},
		},
	}
}
func New0086_0203_003E() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW062",
		Description: "Aeon Labs Garage Door Controller  Gen5 ",

		ManufacturerID: "0086",
		ProductType: "0203",
		ProductID: "003E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x66,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 32,
				Name: "Startup ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Calibration not active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Begin calibration",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Calibration timout",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "Number of alarm musics",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Disable opening alarm",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable alarm prompt",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable alarm prompt",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm volume",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm choice",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Opening alarm LED mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Disable closing alarm",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable alarm prompt",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable alarm prompt",
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm volume",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm choice",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Closing alarm LED mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Unknown state alarm mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Closed alarm mode",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Tamper switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Battery state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 16,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Play or Pause ringtone",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Ringtone test volume",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "Temperature",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 500,
					},
				},
			},
			&parameter{
				ID: 47,
				Name: "Button definition",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mode 0",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Mode 1",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Door state change report type",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Send hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send barrier operator report CC",
					},
				},
			},
			&parameter{
				ID: 241,
				Name: "Pair the Sensor",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Stop sensor pairing",
					},
					Value{
						From: 1431655681,
						To: 1431655681,
						Desc: "Start sensor pairing",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Configuration enabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Configuration disabled (locked)",
					},
				},
			},
		},
	}
}
func New0086_0002_004A() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW074",
		Description: "4 in One MultiSensor (G5)",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "004A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Wake up for 10 minutes when batteries are inserted",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "On time",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 15300,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion sensor",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Enable/disable the selective reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Threshold change in temperature to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Threshold change in humidity to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Threshold change in luminance to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Threshold change in battery level to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "Enable/disable to send the alarm report of low temperature(‐15C)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled (The MultiSensor will report the Multi Level Temperature CC wi",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically in timing intervals for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically in timing intervals for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically in timing intervals for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Configuration Locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default factory setting.",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Resets all configuration parameters to default setting.",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset the product to default factory setting and be excluded from the ",
					},
				},
			},
		},
	}
}
func New0086_0102_004A() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW074",
		Description: "4 in One MultiSensor (G5)",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "004A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Wake up for 10 minutes when batteries are inserted",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "On time",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 15300,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion sensor",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Enable/disable the selective reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Threshold change in temperature to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Threshold change in humidity to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Threshold change in luminance to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Threshold change in battery level to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "Enable/disable to send the alarm report of low temperature(‐15C)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled (The MultiSensor will report the Multi Level Temperature CC wi",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically in timing intervals for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically in timing intervals for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically in timing intervals for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Configuration Locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default factory setting.",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Resets all configuration parameters to default setting.",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset the product to default factory setting and be excluded from the ",
					},
				},
			},
		},
	}
}
func New0086_0202_004A() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW074",
		Description: "4 in One MultiSensor (G5)",

		ManufacturerID: "0086",
		ProductType: "0202",
		ProductID: "004A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Wake up for 10 minutes when batteries are inserted",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "On time",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 15300,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion sensor",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Enable/disable the selective reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Threshold change in temperature to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Threshold change in humidity to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Threshold change in luminance to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Threshold change in battery level to induce an automatic report.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "Enable/disable to send the alarm report of low temperature(‐15C)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled (The MultiSensor will report the Multi Level Temperature CC wi",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Which reports need to send automatically in timing intervals for group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Which reports need to send automatically in timing intervals for group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Which reports need to send automatically in timing intervals for group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Configuration Locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default factory setting.",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Resets all configuration parameters to default setting.",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset the product to default factory setting and be excluded from the ",
					},
				},
			},
		},
	}
}
func New0086_0003_004B() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW075",
		Description: "Smart Energy Switch 3rd Edition",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "004B",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Blinking Behaviour",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output Load Status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always On",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always Off",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Configure the state of red LED ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "The LED will follow the status of its load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch off after 5 seconds",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enables/disables parameter 91/92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 101‐103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Set to Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Nothing",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Report type to send to Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Report type to send to Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Report type to send to Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports to Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Configuration Lock",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
		},
	}
}
func New0086_0203_004B() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW075",
		Description: "Smart Energy Switch 3rd Edition",

		ManufacturerID: "0086",
		ProductType: "0203",
		ProductID: "004B",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Blinking Behaviour",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output Load Status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always On",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always Off",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Configure the state of red LED ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "The LED will follow the status of its load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch off after 5 seconds",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enables/disables parameter 91/92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 101‐103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Set to Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Nothing",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Report type to send to Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Report type to send to Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Report type to send to Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports to Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Configuration Lock",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
		},
	}
}
func New0086_0003_004E() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW078",
		Description: "Heavy Duty Switch",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "004E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0xEF,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output Load Status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always OFF",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Send report on value change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deavtivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "reports to send automatically",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "which reports to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "which reports to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
		},
	}
}
func New0086_0103_004E() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW078",
		Description: "Heavy Duty Switch",

		ManufacturerID: "0086",
		ProductType: "0103",
		ProductID: "004E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0xEF,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output Load Status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always OFF",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Send report on value change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deavtivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "reports to send automatically",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "which reports to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "which reports to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
		},
	}
}
func New0086_0203_004E() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW078",
		Description: "Heavy Duty Switch",

		ManufacturerID: "0086",
		ProductType: "0203",
		ProductID: "004E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0xEF,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output Load Status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always OFF",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deactivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail CC Sent",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report Sent",
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Send report on value change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Deavtivate",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Activate",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "reports to send automatically",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "which reports to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "which reports to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Turn External Button Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Momentary Button Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 State Switch Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 Way Switch Mode",
					},
				},
			},
		},
	}
}
func New0086_0104_0050() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW080",
		Description: "Siren Gen5",

		ManufacturerID: "0086",
		ProductType: "0104",
		ProductID: "0050",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 37,
				Name: "Siren Sound",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not change the current Siren sound",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Siren sound 1 is selected",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Siren sound 2 is selected",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Siren sound 3 is selected",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Siren sound 4 is selected",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Siren sound 5 is selected",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Siren Volume",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not change the current volume",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Set the volume to 88 dB",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Set the volume to 100 dB",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Set the volume to 105 dB",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications to Group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Lock Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Setting of configuration parameters is allowed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "All configuration parameters cannot be set (Locked)",
					},
				},
			},
		},
	}
}
func New0086_0204_0050() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW080",
		Description: "Siren Gen5",

		ManufacturerID: "0086",
		ProductType: "0204",
		ProductID: "0050",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 37,
				Name: "Siren Sound",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not change the current Siren sound",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Siren sound 1 is selected",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Siren sound 2 is selected",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Siren sound 3 is selected",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Siren sound 4 is selected",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Siren sound 5 is selected",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Siren Volume",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not change the current volume",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Set the volume to 88 dB",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Set the volume to 100 dB",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Set the volume to 105 dB",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable Notifications to Group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable Lock Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Setting of configuration parameters is allowed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "All configuration parameters cannot be set (Locked)",
					},
				},
			},
		},
	}
}
func New0086_0001_0058() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW088",
		Description: "4 button keyfob - Gen 5",

		ManufacturerID: "0086",
		ProductType: "0001",
		ProductID: "0058",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 250,
				Name: "Use Mode setting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Group Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
				},
			},
		},
	}
}
func New0086_0101_0058() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW088",
		Description: "4 button keyfob - Gen 5",

		ManufacturerID: "0086",
		ProductType: "0101",
		ProductID: "0058",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 250,
				Name: "Use Mode setting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Group Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
				},
			},
		},
	}
}
func New0086_0201_0058() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW088",
		Description: "4 button keyfob - Gen 5",

		ManufacturerID: "0086",
		ProductType: "0201",
		ProductID: "0058",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 250,
				Name: "Use Mode setting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Group Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene Mode",
					},
				},
			},
		},
	}
}
func New0086_0002_0059() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW089",
		Description: "Recessed Door Sensor Gen5",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0059",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 240,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Configuration Value 4(LSB) SENSOR BINARY",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable SENSOR BINARY (factory)",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "enable SENSOR BINARY",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Config value 3 - BASIC SET",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "do not send BASIC SET",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "send BASIC SET (factory default)",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Permit other configurations",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0086_0102_0059() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW089",
		Description: "Recessed Door Sensor Gen5",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "0059",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 240,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Configuration Value 4(LSB) SENSOR BINARY",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable SENSOR BINARY (factory)",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "enable SENSOR BINARY",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Config value 3 - BASIC SET",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "do not send BASIC SET",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "send BASIC SET (factory default)",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Permit other configurations",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0086_0202_0059() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW089",
		Description: "Recessed Door Sensor Gen5",

		ManufacturerID: "0086",
		ProductType: "0202",
		ProductID: "0059",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 240,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Configuration Value 4(LSB) SENSOR BINARY",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable SENSOR BINARY (factory)",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "enable SENSOR BINARY",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Config value 3 - BASIC SET",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "do not send BASIC SET",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "send BASIC SET (factory default)",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Permit other configurations",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0086_0002_005F() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW095",
		Description: "Home Energy Meter - Gen 5 ",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "005F",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0xEF,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Energy detection mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 16,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Current Threshold - Whole HEM",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Current Threshold - Clamp 1",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Current Threshold - Clamp 2",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Current Threshold - Clamp 3",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Percent Change - Whole HEM",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Percent Change - Clamp 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Percent Change - Clamp 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Percent Change - Clamp 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Report Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Report Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Report Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
		},
	}
}
func New0086_0102_005F() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW095",
		Description: "Home Energy Meter - Gen 5 ",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "005F",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0xEF,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Energy detection mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 16,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Current Threshold - Whole HEM",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Current Threshold - Clamp 1",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Current Threshold - Clamp 2",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Current Threshold - Clamp 3",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Percent Change - Whole HEM",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Percent Change - Clamp 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Percent Change - Clamp 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Percent Change - Clamp 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Report Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Report Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Report Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Interval to send out reports of group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Interval to send out reports of group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Interval to send out reports of group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
		},
	}
}
func New0086_0003_0060() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW096",
		Description: "Smart Switch 6",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0060",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current overload protection enable",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output load after re-power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "last status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "always on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "always off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED color testing",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable send to associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Configure LED state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED follows load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED follows load for 5 seconds",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night light mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Color in night light mode",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Blue night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Green night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Red night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Color in energy mode ",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Green brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Yellow brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Red brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable items 91 and 92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Wattage Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Wattage Percent Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 101‐103 to default. ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Values to send to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Values to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Values to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111‐113 to default.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Time interval for sending to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Time interval for sending to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Time interval for sending to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 600,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "RESET",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0086_0103_0060() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW096",
		Description: "Smart Switch 6",

		ManufacturerID: "0086",
		ProductType: "0103",
		ProductID: "0060",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current overload protection enable",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output load after re-power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "last status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "always on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "always off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED color testing",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable send to associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Configure LED state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED follows load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED follows load for 5 seconds",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night light mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Color in night light mode",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Blue night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Green night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Red night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Color in energy mode ",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Green brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Yellow brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Red brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable items 91 and 92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Wattage Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Wattage Percent Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 101‐103 to default. ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Values to send to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Values to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Values to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111‐113 to default.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Time interval for sending to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Time interval for sending to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Time interval for sending to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 600,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "RESET",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0086_0203_0060() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW096",
		Description: "Smart Switch 6",

		ManufacturerID: "0086",
		ProductType: "0203",
		ProductID: "0060",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current overload protection enable",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output load after re-power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "last status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "always on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "always off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED color testing",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable send to associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Configure LED state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED follows load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED follows load for 5 seconds",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night light mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Color in night light mode",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Blue night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Green night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Red night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Color in energy mode ",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Green brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Yellow brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Red brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable items 91 and 92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Wattage Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Wattage Percent Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 101‐103 to default. ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Values to send to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Values to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Values to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111‐113 to default.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Time interval for sending to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Time interval for sending to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Time interval for sending to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 600,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "RESET",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0086_1D03_0060() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW096",
		Description: "Smart Switch 6",

		ManufacturerID: "0086",
		ProductType: "1D03",
		ProductID: "0060",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current overload protection enable",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Output load after re-power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "last status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "always on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "always off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED color testing",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable send to associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "basic CC report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Configure LED state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED follows load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED follows load for 5 seconds",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night light mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Color in night light mode",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Blue night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Green night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Red night light color value",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Color in energy mode ",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Green brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Yellow brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "Red brightness in energy mode (%)",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable items 91 and 92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Wattage Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Wattage Percent Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set 101‐103 to default. ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Values to send to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Values to send to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Values to send to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Voltage",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Current",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Wattage",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "kWh",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "All Values",
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set 111‐113 to default.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "False",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "True",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Time interval for sending to group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Time interval for sending to group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Time interval for sending to group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 600,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "RESET",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0086_0002_0061() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW097",
		Description: "Dry Contact Sensor",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0061",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0xEF,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable wake up 10 minutes when power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Send battery report when less than percentage",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Set the interval time of battery report",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Send Sensor Binary report to associated devices",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Send Basic report to associated devices",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 122,
				Name: "Notification type to send",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Smoke alarm",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "CO alarm",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "CO2 alarm",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Heat alarm",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Water alarm",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Access control",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Home security",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Power management",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "System",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "Emergency alarm",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Timer ended",
					},
				},
			},
		},
	}
}
func New0086_0102_0061() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW097",
		Description: "Dry Contact Sensor",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "0061",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0xEF,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable wake up 10 minutes when power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Send battery report when less than percentage",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Set the interval time of battery report",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Send Sensor Binary report to associated devices",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Send Basic report to associated devices",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 122,
				Name: "Notification type to send",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Smoke alarm",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "CO alarm",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "CO2 alarm",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Heat alarm",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Water alarm",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Access control",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Home security",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Power management",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "System",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "Emergency alarm",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Timer ended",
					},
				},
			},
		},
	}
}
func New0086_0202_0061() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW097",
		Description: "Dry Contact Sensor",

		ManufacturerID: "0086",
		ProductType: "0202",
		ProductID: "0061",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0xEF,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Sensor binary report on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable wake up 10 minutes when power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Send Basic Set on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Send battery report when less than percentage",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Set the interval time of battery report",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Send Sensor Binary report to associated devices",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Send Basic report to associated devices",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 122,
				Name: "Notification type to send",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Smoke alarm",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "CO alarm",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "CO2 alarm",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Heat alarm",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Water alarm",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Access control",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Home security",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Power management",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "System",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "Emergency alarm",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Timer ended",
					},
				},
			},
		},
	}
}
func New0086_0003_0062() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW098",
		Description: "LED Bulb",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0062",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x33,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 32,
				Name: "Toggle send report on color change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable color updates",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "send HAIL CC",
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Allow external switch to toggle bulb",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Allow external switch to change color",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "Control color mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "reboot normal mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "boot into color programming",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "exit color programming",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "save and exit color programming",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Colour Residence Time",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Colour Change Speed",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Cycle count",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Color transitions (Value 1 MSB) see Overview",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "single color mode with smooth color transitions (default)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "rainbow mode with smooth color transitions",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "multi color mode with smooth color transitions",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "random mode with smooth color transitions",
					},
					Value{
						From: 64,
						To: 64,
						Desc: "single color mode with direct color transitions",
					},
					Value{
						From: 65,
						To: 65,
						Desc: "rainbow color mode with direct color transitions",
					},
					Value{
						From: 66,
						To: 66,
						Desc: "multi color mode with direct color transitions",
					},
					Value{
						From: 67,
						To: 67,
						Desc: "random color mode with direct color transitions",
					},
					Value{
						From: 128,
						To: 128,
						Desc: "single color mode with fade out/fade in transitions",
					},
					Value{
						From: 129,
						To: 129,
						Desc: "rainbow color mode with fade out/fade in transitions",
					},
					Value{
						From: 130,
						To: 130,
						Desc: "multi color mode with fade out/fade in transitions",
					},
					Value{
						From: 131,
						To: 131,
						Desc: "random color mode with fade out/fade in transitions",
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (1,2)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (3,4)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (5.6)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (7,8)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
		},
	}
}
func New0086_0103_0062() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW098",
		Description: "LED Bulb",

		ManufacturerID: "0086",
		ProductType: "0103",
		ProductID: "0062",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x33,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 32,
				Name: "Toggle send report on color change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable color updates",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "send HAIL CC",
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Allow external switch to toggle bulb",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Allow external switch to change color",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "Control color mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "reboot normal mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "boot into color programming",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "exit color programming",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "save and exit color programming",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Colour Residence Time",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Colour Change Speed",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Cycle count",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Color transitions (Value 1 MSB) see Overview",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "single color mode with smooth color transitions (default)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "rainbow mode with smooth color transitions",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "multi color mode with smooth color transitions",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "random mode with smooth color transitions",
					},
					Value{
						From: 64,
						To: 64,
						Desc: "single color mode with direct color transitions",
					},
					Value{
						From: 65,
						To: 65,
						Desc: "rainbow color mode with direct color transitions",
					},
					Value{
						From: 66,
						To: 66,
						Desc: "multi color mode with direct color transitions",
					},
					Value{
						From: 67,
						To: 67,
						Desc: "random color mode with direct color transitions",
					},
					Value{
						From: 128,
						To: 128,
						Desc: "single color mode with fade out/fade in transitions",
					},
					Value{
						From: 129,
						To: 129,
						Desc: "rainbow color mode with fade out/fade in transitions",
					},
					Value{
						From: 130,
						To: 130,
						Desc: "multi color mode with fade out/fade in transitions",
					},
					Value{
						From: 131,
						To: 131,
						Desc: "random color mode with fade out/fade in transitions",
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (1,2)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (3,4)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (5.6)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (7,8)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
		},
	}
}
func New0086_0203_0062() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW098",
		Description: "LED Bulb",

		ManufacturerID: "0086",
		ProductType: "0203",
		ProductID: "0062",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x33,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 32,
				Name: "Toggle send report on color change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable color updates",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "send HAIL CC",
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Allow external switch to toggle bulb",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Allow external switch to change color",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "Control color mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "reboot normal mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "boot into color programming",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "exit color programming",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "save and exit color programming",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Colour Residence Time",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Colour Change Speed",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Cycle count",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Color transitions (Value 1 MSB) see Overview",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "single color mode with smooth color transitions (default)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "rainbow mode with smooth color transitions",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "multi color mode with smooth color transitions",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "random mode with smooth color transitions",
					},
					Value{
						From: 64,
						To: 64,
						Desc: "single color mode with direct color transitions",
					},
					Value{
						From: 65,
						To: 65,
						Desc: "rainbow color mode with direct color transitions",
					},
					Value{
						From: 66,
						To: 66,
						Desc: "multi color mode with direct color transitions",
					},
					Value{
						From: 67,
						To: 67,
						Desc: "random color mode with direct color transitions",
					},
					Value{
						From: 128,
						To: 128,
						Desc: "single color mode with fade out/fade in transitions",
					},
					Value{
						From: 129,
						To: 129,
						Desc: "rainbow color mode with fade out/fade in transitions",
					},
					Value{
						From: 130,
						To: 130,
						Desc: "multi color mode with fade out/fade in transitions",
					},
					Value{
						From: 131,
						To: 131,
						Desc: "random color mode with fade out/fade in transitions",
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (1,2)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (3,4)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (5.6)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 38,
				Name: "Color index configuration for multi mode (7,8)",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 136,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Enable notifications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
		},
	}
}
func New0086_0003_0063() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW099",
		Description: "Smart Dimmer 6",

		ManufacturerID: "0086",
		ProductType: "0003",
		ProductID: "0063",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0xEF,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Initial Load State",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always On",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always Off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED Color Test",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Indicator Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Energy Mode (Load State)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night Light Mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "LED Color for Night Light",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "LED Brightness",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 6579300,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable Report triggered by minimal change of value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Reset Report Assignments for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Reports for Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Reports for Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Reports for Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Reset Report Time Intervals for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Report Group 1 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Report Group 2 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Report Group 3 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Aeon Labs Standard Product",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Others",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Lock",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to Factory Default Settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0103_0063() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW099",
		Description: "Smart Dimmer 6",

		ManufacturerID: "0086",
		ProductType: "0103",
		ProductID: "0063",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0xEF,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Initial Load State",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always On",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always Off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED Color Test",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Indicator Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Energy Mode (Load State)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night Light Mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "LED Color for Night Light",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "LED Brightness",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 6579300,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable Report triggered by minimal change of value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Reset Report Assignments for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Reports for Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Reports for Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Reports for Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Reset Report Time Intervals for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Report Group 1 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Report Group 2 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Report Group 3 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Aeon Labs Standard Product",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Others",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Lock",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to Factory Default Settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0203_0063() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW099",
		Description: "Smart Dimmer 6",

		ManufacturerID: "0086",
		ProductType: "0203",
		ProductID: "0063",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0xEF,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Initial Load State",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always On",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always Off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED Color Test",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Indicator Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Energy Mode (Load State)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night Light Mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "LED Color for Night Light",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "LED Brightness",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 6579300,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable Report triggered by minimal change of value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Reset Report Assignments for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Reports for Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Reports for Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Reports for Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Reset Report Time Intervals for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Report Group 1 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Report Group 2 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Report Group 3 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Aeon Labs Standard Product",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Others",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Lock",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to Factory Default Settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_1D03_0063() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW099",
		Description: "Smart Dimmer 6",

		ManufacturerID: "0086",
		ProductType: "1D03",
		ProductID: "0063",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x33,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0xEF,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Initial Load State",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Last Status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always On",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always Off",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "RGB LED Color Test",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Notification on Status Change",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hail",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Basic CC Report",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Indicator Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Energy Mode (Load State)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary Mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Night Light Mode",
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "LED Color for Night Light",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 84,
				Name: "LED Brightness",
				Size: 3,
				Values: []Value{
					Value{
						From: 0,
						To: 6579300,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "Enable Report triggered by minimal change of value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Minimum Change to send Report (Watt)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Minimum Change to send Report (%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Reset Report Assignments for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Reports for Group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Reports for Group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Reports for Group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Reset Report Time Intervals for All Groups",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reset to Defaults",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Report Group 1 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Report Group 2 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Report Group 3 Time Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Aeon Labs Standard Product",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Others",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Configuration Lock",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to Factory Default Settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset",
					},
				},
			},
		},
	}
}
func New0086_0002_0064() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW100",
		Description: "MultiSensor 6",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0064",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Stay Awake in Battery Mode			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Motion Sensor reset timeout",
				Size: 2,
				Values: []Value{
					Value{
						From: 5,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion sensor sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable, sensitivity level 1 (minimum)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Enable, sensitivity level 2",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Enable, sensitivity level 3",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Enable, sensitivity level 4",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Enable, sensitivity level 5 (maximum)",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Motion Sensor Triggered Command",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Sensor Binary Report CC",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low Battery Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Temperature Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Humidity Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Luminance Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Battery Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "Ultraviolet Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "Send Alarm Report if low temperature",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set parameters 101-103 to default.			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Group 1 Periodic Reports",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Group 2 Periodic Reports			",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Group 3 Periodic Reports			",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set parameters 111-113 to default.			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Group 1 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Group 2 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Group 3 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 201,
				Name: "Temperature Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 202,
				Name: "Humidity Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -50,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 203,
				Name: "Luminance Sensor Calibration",
				Size: 2,
				Values: []Value{
					Value{
						From: -1000,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 204,
				Name: "Ultraviolet Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -10,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock Configuration Parameters			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default factory settings",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Resets all configuration parameters to defaults",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset to default factory settings and be excluded",
					},
				},
			},
		},
	}
}
func New0086_0102_0064() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW100",
		Description: "MultiSensor 6",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "0064",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Stay Awake in Battery Mode			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Motion Sensor reset timeout",
				Size: 2,
				Values: []Value{
					Value{
						From: 5,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion sensor sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable, sensitivity level 1 (minimum)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Enable, sensitivity level 2",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Enable, sensitivity level 3",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Enable, sensitivity level 4",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Enable, sensitivity level 5 (maximum)",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Motion Sensor Triggered Command",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Sensor Binary Report CC",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low Battery Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Temperature Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Humidity Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Luminance Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Battery Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "Ultraviolet Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "Send Alarm Report if low temperature",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set parameters 101-103 to default.			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Group 1 Periodic Reports",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Group 2 Periodic Reports			",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Group 3 Periodic Reports			",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set parameters 111-113 to default.			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Group 1 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Group 2 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Group 3 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 201,
				Name: "Temperature Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 202,
				Name: "Humidity Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -50,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 203,
				Name: "Luminance Sensor Calibration",
				Size: 2,
				Values: []Value{
					Value{
						From: -1000,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 204,
				Name: "Ultraviolet Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -10,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock Configuration Parameters			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default factory settings",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Resets all configuration parameters to defaults",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset to default factory settings and be excluded",
					},
				},
			},
		},
	}
}
func New0086_0202_0064() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW100",
		Description: "MultiSensor 6",

		ManufacturerID: "0086",
		ProductType: "0202",
		ProductID: "0064",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Stay Awake in Battery Mode			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Motion Sensor reset timeout",
				Size: 2,
				Values: []Value{
					Value{
						From: 5,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion sensor sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable, sensitivity level 1 (minimum)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Enable, sensitivity level 2",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Enable, sensitivity level 3",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Enable, sensitivity level 4",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Enable, sensitivity level 5 (maximum)",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Motion Sensor Triggered Command",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Sensor Binary Report CC",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low Battery Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Selective Reporting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Temperature Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Humidity Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Luminance Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Battery Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "Ultraviolet Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "Send Alarm Report if low temperature",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Set parameters 101-103 to default.			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Group 1 Periodic Reports",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Group 2 Periodic Reports			",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "Group 3 Periodic Reports			",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Set parameters 111-113 to default.			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Group 1 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Group 2 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Group 3 Report Interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 5,
						To: 2678400,
					},
				},
			},
			&parameter{
				ID: 201,
				Name: "Temperature Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 202,
				Name: "Humidity Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -50,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 203,
				Name: "Luminance Sensor Calibration",
				Size: 2,
				Values: []Value{
					Value{
						From: -1000,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 204,
				Name: "Ultraviolet Sensor Calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: -10,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock Configuration Parameters			",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to default factory settings",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Resets all configuration parameters to defaults",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset to default factory settings and be excluded",
					},
				},
			},
		},
	}
}
func New0086_0103_006F() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW111",
		Description: "Nano Dimmer ",

		ManufacturerID: "0086",
		ProductType: "0103",
		ProductID: "006F",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Enable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Overheat protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Configure the output status after re-power on it",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Last status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always off",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Group 1 notificaiton",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Send Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Basic CC report",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send Multilevel Switch report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Hail CC when using the manual switch",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Group 3 notificaiton",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Send Basic Set CC",
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Group 4 notificaiton",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Send Basic Set CC",
					},
				},
			},
			&parameter{
				ID: 85,
				Name: "Set appointment 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16843197,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "Set appointment 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16843197,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "(En/dis)able 91 and 92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Set the threshold value of wattage",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Set the threshold value of wattage",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Reset 101-103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset the parameter 101-103",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Set group 2 reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "To set which reports need to be sent in Report group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Reset 111-113 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset the parameter 111-113",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Automatic report interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Automatic report interval group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Automatic report interval group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Configure the external switch mode for S1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Enter automatic identification mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary push button mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 way switch mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "2-state switch mode",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Configure the external switch mode for S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Enter automatic identification mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary push button mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 way switch mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "2-state switch mode",
					},
				},
			},
			&parameter{
				ID: 122,
				Name: "Get the state of touch panel port",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) the touch panel is not connected.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "the touch panel is connected",
					},
				},
			},
			&parameter{
				ID: 123,
				Name: "Set the control destination for external switch S1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "control the output loads of itself",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control the other nodes",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "(Default) control the output loads of itself and other nodes",
					},
				},
			},
			&parameter{
				ID: 124,
				Name: "Set the control destination for external switch S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "control the output loads of itself",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control the other nodes",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "(Default) control the output loads of itself and other nodes",
					},
				},
			},
			&parameter{
				ID: 125,
				Name: "Set the default dimming rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 128,
				Name: "Get the current working mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Unknown working mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2-wire mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3-wire mode",
					},
				},
			},
			&parameter{
				ID: 129,
				Name: "Set the dimming principle",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Trailing edge mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Leading edge mode",
					},
				},
			},
			&parameter{
				ID: 130,
				Name: "To get what type of load the Dimmer is connected to",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Unknown load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Resistive load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Capacitive load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Inductive load",
					},
				},
			},
			&parameter{
				ID: 131,
				Name: "Set the min brightness level that the load can reach to",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 132,
				Name: "Set the max brightness level that the load can reach to",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 249,
				Name: "Set the recognition way of load",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Never recognize the load when power on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only recognize once when first power on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "(Default) Recognize the load once power on",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock/unlock configuration parameters",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset the Nano Dimmer",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) factory default",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "factory default and remove",
					},
				},
			},
		},
	}
}
func New0086_0203_006F() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW111",
		Description: "Nano Dimmer ",

		ManufacturerID: "0086",
		ProductType: "0203",
		ProductID: "006F",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Current Overload Protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Enable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Overheat protection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Configure the output status after re-power on it",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Last status",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Always on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Always off",
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "Group 1 notificaiton",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Send Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Hail CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Basic CC report",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send Multilevel Switch report",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Hail CC when using the manual switch",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "Group 3 notificaiton",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Send Basic Set CC",
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Group 4 notificaiton",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send Nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Send Basic Set CC",
					},
				},
			},
			&parameter{
				ID: 85,
				Name: "Set appointment 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16843197,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "Set appointment 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16843197,
					},
				},
			},
			&parameter{
				ID: 90,
				Name: "(En/dis)able 91 and 92",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 91,
				Name: "Set the threshold value of wattage",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 60000,
					},
				},
			},
			&parameter{
				ID: 92,
				Name: "Set the threshold value of wattage",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "Reset 101-103 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset the parameter 101-103",
					},
				},
			},
			&parameter{
				ID: 102,
				Name: "Set group 2 reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 103,
				Name: "To set which reports need to be sent in Report group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 110,
				Name: "Reset 111-113 to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset the parameter 111-113",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Automatic report interval group 1",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 112,
				Name: "Automatic report interval group 2",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 113,
				Name: "Automatic report interval group 3",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
			&parameter{
				ID: 120,
				Name: "Configure the external switch mode for S1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Enter automatic identification mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary push button mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 way switch mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "2-state switch mode",
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Configure the external switch mode for S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Enter automatic identification mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Momentary push button mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 way switch mode",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "2-state switch mode",
					},
				},
			},
			&parameter{
				ID: 122,
				Name: "Get the state of touch panel port",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) the touch panel is not connected.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "the touch panel is connected",
					},
				},
			},
			&parameter{
				ID: 123,
				Name: "Set the control destination for external switch S1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "control the output loads of itself",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control the other nodes",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "(Default) control the output loads of itself and other nodes",
					},
				},
			},
			&parameter{
				ID: 124,
				Name: "Set the control destination for external switch S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "control the output loads of itself",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control the other nodes",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "(Default) control the output loads of itself and other nodes",
					},
				},
			},
			&parameter{
				ID: 125,
				Name: "Set the default dimming rate",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 128,
				Name: "Get the current working mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Unknown working mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2-wire mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3-wire mode",
					},
				},
			},
			&parameter{
				ID: 129,
				Name: "Set the dimming principle",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Trailing edge mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "(Default) Leading edge mode",
					},
				},
			},
			&parameter{
				ID: 130,
				Name: "To get what type of load the Dimmer is connected to",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Unknown load",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Resistive load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Capacitive load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Inductive load",
					},
				},
			},
			&parameter{
				ID: 131,
				Name: "Set the min brightness level that the load can reach to",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 132,
				Name: "Set the max brightness level that the load can reach to",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 249,
				Name: "Set the recognition way of load",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Never recognize the load when power on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only recognize once when first power on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "(Default) Recognize the load once power on",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock/unlock configuration parameters",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset the Nano Dimmer",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "(Default) factory default",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "factory default and remove",
					},
				},
			},
		},
	}
}
func New0086_0002_0070() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW112",
		Description: "Door/Window Sensor 6",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0070",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Value that will be sent on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "	Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low battery threshold (10% to 50%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 240,
						To: 134217727,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Reporting mode on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Sensor Binary Report CC",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send Basic Set CC and Sensor Binary Report CC",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock/Unlock all configuration parameters",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset sensor to factory default and remove from network",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
		},
	}
}
func New0086_0102_0070() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW112",
		Description: "Door/Window Sensor 6",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "0070",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Value that will be sent on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "	Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low battery threshold (10% to 50%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 240,
						To: 134217727,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Reporting mode on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Sensor Binary Report CC",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send Basic Set CC and Sensor Binary Report CC",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock/Unlock all configuration parameters",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset sensor to factory default and remove from network",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
		},
	}
}
func New0086_0202_0070() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW112",
		Description: "Door/Window Sensor 6",

		ManufacturerID: "0086",
		ProductType: "0202",
		ProductID: "0070",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Value that will be sent on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "	Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low battery threshold (10% to 50%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 240,
						To: 134217727,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Reporting mode on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Sensor Binary Report CC",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send Basic Set CC and Sensor Binary Report CC",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock/Unlock all configuration parameters",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset sensor to factory default and remove from network",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
		},
	}
}
func New0086_1D02_0070() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW112",
		Description: "Door/Window Sensor 6",

		ManufacturerID: "0086",
		ProductType: "1D02",
		ProductID: "0070",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Value that will be sent on open/close events",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On for opened, Off for closed",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "	Off for opened, On for closed",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low battery threshold (10% to 50%)",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 101,
				Name: "Low battery voltage check",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 111,
				Name: "Low battery voltage check time",
				Size: 4,
				Values: []Value{
					Value{
						From: 240,
						To: 134217727,
					},
				},
			},
			&parameter{
				ID: 121,
				Name: "Reporting mode on open/close event",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send nothing",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send Basic Set CC",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send Sensor Binary Report CC",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send Basic Set CC and Sensor Binary Report CC",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Lock/Unlock all configuration parameters",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Unlock",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Lock",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset sensor to factory default and remove from network",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 2147483647,
					},
				},
			},
		},
	}
}
func New0086_0004_0075() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW117",
		Description: "Range Extender 6",

		ManufacturerID: "0086",
		ProductType: "0004",
		ProductID: "0075",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x33,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				Secure: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 33,
				Name: "RGB value",
				Size: 3,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Reserved",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Red value",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Green value",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Blue value",
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Default status of the LED",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "The green LED remains On for 2 seconds",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off",
					},
				},
			},
			&parameter{
				ID: 200,
				Name: "Partner ID",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Aeon Labs Standard Product",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "others",
					},
				},
			},
			&parameter{
				ID: 252,
				Name: "Enable/disable the Configuration to be locked",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enable",
					},
				},
			},
			&parameter{
				ID: 254,
				Name: "Device Tag.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset to factory default",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "DELETE ME",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset to factory default",
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset configuration",
					},
				},
			},
		},
	}
}
func New0086_0002_0082() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW130",
		Description: "WallMote Quad",

		ManufacturerID: "0086",
		ProductType: "0002",
		ProductID: "0082",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Touch sound",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Touch vibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Button slide",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Notification report",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Central scene",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Central scene and config",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low battery value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset the WallMote Quad",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset to factory default",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset and remove",
					},
				},
			},
		},
	}
}
func New0086_0102_0082() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW130",
		Description: "WallMote Quad",

		ManufacturerID: "0086",
		ProductType: "0102",
		ProductID: "0082",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Touch sound",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Touch vibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Button slide",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Notification report",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Central scene",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Central scene and config",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low battery value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset the WallMote Quad",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset to factory default",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset and remove",
					},
				},
			},
		},
	}
}
func New0086_0202_0082() *Device{
	return &Device{
		Brand: "AEON Labs",
		Product: "ZW130",
		Description: "WallMote Quad",

		ManufacturerID: "0086",
		ProductType: "0202",
		ProductID: "0082",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Touch sound",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Touch vibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Button slide",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Notification report",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Central scene",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Central scene and config",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Low battery value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 50,
					},
				},
			},
			&parameter{
				ID: 255,
				Name: "Reset the WallMote Quad",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reset to factory default",
					},
					Value{
						From: 1431655765,
						To: 1431655765,
						Desc: "Reset and remove",
					},
				},
			},
		},
	}
}
func New0111_8200_0200() *Device{
	return &Device{
		Brand: "Airline Mechanical Co., Ltd.",
		Product: "ZDS-UD10",
		Description: "Dimming Switch Module",

		ManufacturerID: "0111",
		ProductType: "8200",
		ProductID: "0200",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New008A_0005_0101() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "BeNext Alarm Sound",
		Description: "Alarm sound",

		ManufacturerID: "008A",
		ProductType: "0005",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set to Default",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Destination routine on/off",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 25,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Desitnation routine succes time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Destination routine failed time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Select index sound/light mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 6,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "power offline sound/light mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 6,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "sound/light index 1",
				Size: 18,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 28,
				Name: "sound/light index 2",
				Size: 18,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "sound/light index 3",
				Size: 18,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "sound/light index 4",
				Size: 18,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "sound/light index 5",
				Size: 18,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 100,
				Name: "sound/light index 6",
				Size: 18,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New008A_000D_0100() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "builtInDimmer",
		Description: "BeNext Built-in Dimmer",

		ManufacturerID: "008A",
		ProductType: "000D",
		ProductID: "0100",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Reset to factory settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "The way how the button reacts when press/released",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Toggle light when button is pressed, no action when button is released",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Toggle light when button is released, Start dimming when button presse",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Toggle light when button is pressed and when button is Released, start",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "The way the Built-in Dimmer reacts when light is turned on/off with button",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Switch own light on and send a Z-Wave message to all associated nodes ",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Don.t switch own light on but only send a Z-Wave message to associated",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Enable dimming",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Maximum load",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimming speed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Fading up speed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Toggle time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Auto meter report: percentage",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Auto meter report: watt",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Auto meter report: time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Last known status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Fading down speed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New008A_0004_0100() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "doorSensor",
		Description: "BeNext Door Sensor",

		ManufacturerID: "008A",
		ProductType: "0004",
		ProductID: "0100",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Reset to factory settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "External contact behavior",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send alarm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send basic frame",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Operating mode.",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Normal",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Mode 1 report",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Always-on",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Temperature offset",
				Size: 2,
				Values: []Value{
					Value{
						From: -32768,
						To: 32767,
					},
				},
			},
		},
	}
}
func New008A_0004_0101() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "doorSensor",
		Description: "BeNext Door Sensor",

		ManufacturerID: "008A",
		ProductType: "0004",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Reset to factory settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "External contact behavior",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send alarm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send basic frame",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Operating mode.",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Normal",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Mode 1 report",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Always-on",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Temperature offset",
				Size: 2,
				Values: []Value{
					Value{
						From: -32768,
						To: 32767,
					},
				},
			},
		},
	}
}
func New008A_0020_0001() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "Energy Switch +",
		Description: "Energy Switch +",

		ManufacturerID: "008A",
		ProductType: "0020",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set to Default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Amount of decimals",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Power Limit ",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 4500,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Start up with last known socket status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Auto report %",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Auto report Watt",
				Size: 1,
				Values: []Value{
					Value{
						From: 5,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Auto report time",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Relais delay time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Led indicator",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 2,
					},
				},
			},
		},
	}
}
func New008A_0021_0102() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "Heating Control",
		Description: "Thermostat for controlling the opentherm protocol",

		ManufacturerID: "008A",
		ProductType: "0021",
		ProductID: "0102",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "6",
			},
			&CommandClass{
				ID: 0x40,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				Version: "3",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set to default",
				Size: 1,
				Values: []Value{
					Value{
						From: 255,
						To: 255,
						Desc: "Reset",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Data request interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "T room update difference",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "T setpoint update difference",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Unsolicited CRC",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Off",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "On",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Type of 'special' thermostat",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No special",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Remeha Celcia 20",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Honeywell (rounded temperatures)",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Status auto report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable boiler/thermostat status messages auto report",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Enable boiler/thermostat status messages auto report",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Enable/Disable thermostat schedule",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable thermostat schedule",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Enable thermostat schedule",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Readout manual setpoint thermostat",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Readout OFF",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Readout ON",
					},
				},
			},
		},
	}
}
func New008A_0003_0101() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "Molite",
		Description: "Movement sensor with temperature and light sensor",

		ManufacturerID: "008A",
		ProductType: "0003",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Reset to factory settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Mode timeout",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Switch off time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "The mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Mode 1: no detection possible. Battery save mode.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Mode 2: normal operation mode: send on after detection and off after g",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Mode 3: Z-Wave chip is always on to request e.g. version or manufactur",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "The temperature offset",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Light table 100 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Light table 90 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Light table 80 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Light table 70 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Light table 60 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Light table 50 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Light table 40 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Light table 30 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Light table 20 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Light table 10 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New008A_0003_0100() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "Molite",
		Description: "Movement sensor with temperature and light sensor",

		ManufacturerID: "008A",
		ProductType: "0003",
		ProductID: "0100",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Reset to factory settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Mode timeout",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Switch off time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "The mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Mode 1: no detection possible. Battery save mode.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Mode 2: normal operation mode: send on after detection and off after g",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Mode 3: Z-Wave chip is always on to request e.g. version or manufactur",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "The temperature offset",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Light table 100 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Light table 90 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Light table 80 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Light table 70 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Light table 60 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Light table 50 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Light table 40 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Light table 30 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Light table 20 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Light table 10 %",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
		},
	}
}
func New008A_002F_0100() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "P1-dongle",
		Description: "P1-dongle",

		ManufacturerID: "008A",
		ProductType: "002F",
		ProductID: "0100",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 4,
				Name: "Baud Rate",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 40000,
					},
				},
			},
		},
	}
}
func New008A_0018_0100() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "plugInDimmer",
		Description: "BeNext Plug-in Dimmer",

		ManufacturerID: "008A",
		ProductType: "0018",
		ProductID: "0100",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Reset to factory settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Enable dimming",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Maximum powerload",
				Size: 2,
				Values: []Value{
					Value{
						From: 25,
						To: 150,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimming speed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Fading up speed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Auto meter report: percentage",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Auto meter report: watt",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Auto meter report: time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Last known status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Fading down speed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New008A_0008_0101() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "powerSwitch",
		Description: "BeNext Power Switch EU",

		ManufacturerID: "008A",
		ProductType: "0008",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set to Default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Start up with last known socket status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Relay delay time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Led indicator",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 2,
					},
				},
			},
		},
	}
}
func New008A_0007_0101() *Device{
	return &Device{
		Brand: "BeNext",
		Product: "Tag Reader",
		Description: "Tag Reader",

		ManufacturerID: "008A",
		ProductType: "0007",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x63,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x76,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Set to Default",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Feedback time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Feedback time-out",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "feedback beeps per second",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 5,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Mode ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "RFID circuit start-up time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0169_0001_0001() *Device{
	return &Device{
		Brand: "Bönig und Kallenbach oHG",
		Product: "POPE005206",
		Description: "Z Weather",

		ManufacturerID: "0169",
		ProductType: "0001",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "6",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x89,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9B,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Wind Speed Action Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 30,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Twilight Action Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
		},
	}
}
func New0138_0001_0002() *Device{
	return &Device{
		Brand: "BRK Brands, Inc.",
		Product: "ZCOMBO",
		Description: "Smoke and Carbon Monoxide Alarm",

		ManufacturerID: "0138",
		ProductType: "0001",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send double alarms",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Single Alarm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Double Alarm",
					},
				},
			},
		},
	}
}
func New0138_0001_0001() *Device{
	return &Device{
		Brand: "BRK Brands, Inc.",
		Product: "ZSMOKE",
		Description: "Smoke Alarm",

		ManufacturerID: "0138",
		ProductType: "0001",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send double alarms",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Single Alarm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Double Alarm",
					},
				},
			},
		},
	}
}
func New0166_0100_0100() *Device{
	return &Device{
		Brand: "CBCC Domotique SAS",
		Product: "SW-ZCS-1",
		Description: "Cord Switch",

		ManufacturerID: "0166",
		ProductType: "0100",
		ProductID: "0100",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Switch All",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch on only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch off only",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Fully enabled (Default)",
					},
				},
			},
		},
	}
}
func New0116_0002_0001() *Device{
	return &Device{
		Brand: "Chromagic Technologies Corporation",
		Product: "HSM02",
		Description: "Door Window Sensor",

		ManufacturerID: "0116",
		ProductType: "0002",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Configuring the OFF Delay",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
		},
	}
}
func New0116_0001_0001() *Device{
	return &Device{
		Brand: "Chromagic Technologies Corporation",
		Product: "SM103",
		Description: "Door/Window Contact",

		ManufacturerID: "0116",
		ProductType: "0001",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Configuring the OFF Delay",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
		},
	}
}
func New0200_0003_0002() *Device{
	return &Device{
		Brand: "Cloud Media",
		Product: "A803N",
		Description: "Motion Sensor",

		ManufacturerID: "0200",
		ProductType: "0003",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Param 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Param 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Param 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Param 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Param 5",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Param 6",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Param 7",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Param 8",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Param 9",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Param 10",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Param 11",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Param 12",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0179_0013_0021() *Device{
	return &Device{
		Brand: "ConnectHome",
		Product: "CH-201",
		Description: "CH-201 Thermostat",

		ManufacturerID: "0179",
		ProductType: "0013",
		ProductID: "0021",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x40,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x42,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New001A_5342_0000() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RF9500",
		Description: "Battery Switch",

		ManufacturerID: "001A",
		ProductType: "5342",
		ProductID: "0000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New001A_534C_0000() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RF9501",
		Description: "15A Light Switch",

		ManufacturerID: "001A",
		ProductType: "534C",
		ProductID: "0000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Delayed OFF time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Panic ON time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Panic OFF time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Basic Set Value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Power Up State",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ON",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Last State",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Panic mode enable",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ON",
					},
				},
			},
		},
	}
}
func New001A_5352_0000() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RF9517",
		Description: "Accessory Switch",

		ManufacturerID: "001A",
		ProductType: "5352",
		ProductID: "0000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Delayed OFF time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New001A_4449_0002() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RF9534WS",
		Description: "Smart Dimmer",

		ManufacturerID: "001A",
		ProductType: "4449",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New001A_4449_0101() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RF9540-N",
		Description: "All Load Dimmer Light Switch",

		ManufacturerID: "001A",
		ProductType: "4449",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Delayed OFF time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Panic ON time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Panic OFF time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Basic Set Value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Power Up State",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ON",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Last State",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Panic mode enable",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ON",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimmer Ramp Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Kickstart / Rapid Start",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ON",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Minimum Dimmer Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 4,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Maximum Dimmer Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 4,
						To: 99,
					},
				},
			},
		},
	}
}
func New001A_4441_0000() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RF9542",
		Description: "Dimmer Accessory Switch",

		ManufacturerID: "001A",
		ProductType: "4441",
		ProductID: "0000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Delayed OFF time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Panic ON time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Panic OFF time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Basic Set Value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Power Up State",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ON",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Last State",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Panic mode enable",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ON",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimmer Ramp Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New001A_5244_0000() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RFTR9505",
		Description: "Duplex receptical",

		ManufacturerID: "001A",
		ProductType: "5244",
		ProductID: "0000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x82,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New001A_574D_0000() *Device{
	return &Device{
		Brand: "Cooper Wiring Devices",
		Product: "RFWC5",
		Description: "5-Scene Keypad",

		ManufacturerID: "001A",
		ProductType: "574D",
		ProductID: "0000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x21,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x22,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x2D,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0248_0003_0001() *Device{
	return &Device{
		Brand: "Coqon",
		Product: "PSMZ0001",
		Description: "Plug module",

		ManufacturerID: "0248",
		ProductType: "0003",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x59,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0002_0115_A010() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "010101",
		Description: "Popp Wireless Thermostatic Valve TRV",

		ManufacturerID: "0002",
		ProductType: "0115",
		ProductID: "A010",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "6",
			},
			&CommandClass{
				ID: 0x43,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x46,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0002_0003_8010() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "014G0160",
		Description: "Z-Wave room sensor",

		ManufacturerID: "0002",
		ProductType: "0003",
		ProductID: "8010",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x43,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x87,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Temperature Report threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Set-point display resolution",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Min Set-point and override limit",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 40,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Max Set-point and override limit",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 40,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "LED Flash period",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Set-point control function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Temporarily override scheduler",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Set-point type in Thermostat_Setpoint_Reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Heating",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Cooling",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "Auto-Changeover",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "LED on time",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 5,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Number of LED flashes (duration)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "LED color",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Green",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Red",
					},
				},
			},
		},
	}
}
func New0002_0248_A020() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "DTHERMZ5",
		Description: "Z-Wave room sensor",

		ManufacturerID: "0002",
		ProductType: "0248",
		ProductID: "A020",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "6",
			},
			&CommandClass{
				ID: 0x43,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x53,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x87,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Temperature Report threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Set-point display resolution",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Min Set-point and override limit",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 40,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Max Set-point and override limit",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 40,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "LED Flash period",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Set-point control function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Temporarily override scheduler",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "enabled",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Set-point type in Thermostat_Setpoint_Reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Heating",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Cooling",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "Auto-Changeover",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "LED on time",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 5,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Number of LED flashes (duration)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "LED color",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Green",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Red",
					},
				},
			},
		},
	}
}
func New0002_0248_A010() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "DTHERMZ6",
		Description: "Living Connect Z Thermostat",

		ManufacturerID: "0002",
		ProductType: "0248",
		ProductID: "A010",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "6",
			},
			&CommandClass{
				ID: 0x43,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x46,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0002_0005_0004() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "LC-13",
		Description: "Living Connect Z Thermostat",

		ManufacturerID: "0002",
		ProductType: "0005",
		ProductID: "0004",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x81,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0002_8005_0001() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "LC-13",
		Description: "Living Connect Z Thermostat",

		ManufacturerID: "0002",
		ProductType: "8005",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x81,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0002_8005_0002() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "LC-13",
		Description: "Living Connect Z Thermostat",

		ManufacturerID: "0002",
		ProductType: "8005",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x81,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0002_0005_0003() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "LCZ251",
		Description: "Living Connect Z Thermostat 2.51",

		ManufacturerID: "0002",
		ProductType: "0005",
		ProductID: "0003",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0002_0005_0175() *Device{
	return &Device{
		Brand: "Danfoss",
		Product: "MT02650",
		Description: "Devolo Thermostat (09356)",

		ManufacturerID: "0002",
		ProductType: "0005",
		ProductID: "0175",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x43,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0175_0001_0011() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "MT02646",
		Description: "Home Control Metering Plug",

		ManufacturerID: "0175",
		ProductType: "0001",
		ProductID: "0011",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Watt Meter Report Period",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "KWH Meter Report Period",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Threshold of current for Load Caution",
				Size: 2,
				Values: []Value{
					Value{
						From: 10,
						To: 1300,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Threshold of KWh for Load Caution",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Restore switch state mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Upon the return of AC power, the switch will remain Off.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "resume previous state",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Upon the return of AC power, the switch will turn On.",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Mode of switch off function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "SWITCH OFF commands received are ignored.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "SWITCH OFF commands received are honored.",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "LED indication mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "The LED follows the state of the switch.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "LED off with load ON",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Flash mode",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Auto off timer",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "RF off command mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "When SWITCH ALL OFF is received, the MT02646 turns Off.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Ignore ALL OFF",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "If ALL OFF is received, turn ON if it is OFF",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "When SWITCH ALL OFF is received, the MT02646 turns On.",
					},
				},
			},
		},
	}
}
func New0175_0002_000D() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "MT02647",
		Description: "Devolo Motion Sensor",

		ManufacturerID: "0175",
		ProductType: "0002",
		ProductID: "000D",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x62,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x63,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				Secure: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x76,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Off",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "On",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Light Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Operation Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Multi-Sensor Function Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Re-Detect Interval Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Turn Off Light Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Auto Report Battery Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Auto Report Door/Window State Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Auto Report Illumination Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Auto Report Temperature Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Auto Report Tick Interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Temperature Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Illumination Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
		},
	}
}
func New0175_0002_000E() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "MT02648",
		Description: "Door/Window Contact",

		ManufacturerID: "0175",
		ProductType: "0002",
		ProductID: "000E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable PIR",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Light Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable (Light OFF)",
					},
					Value{
						From: 100,
						To: 100,
						Desc: "Disable (Light ON)",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Operation Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 10,
						Desc: "Preset: Celsius and LED on = Bits: 00001010 = 10",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Multi-Sensor Function Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 20,
						To: 20,
						Desc: "Preset: PIR Super Sensitivity and Binary Report = Bits: 00010100 = 20",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Re-Detect Interval Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Turn Off Light Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Auto Report Battery Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Auto Report Door/Window State Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Auto Report Illumination Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Auto Report Temperature Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Auto Report Tick Interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Temperature Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Illumination Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
		},
	}
}
func New0175_0100_0101() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "MT2652",
		Description: "Scene Switch",

		ManufacturerID: "0175",
		ProductType: "0100",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2D,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Button 1 and 3 pair mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Separately",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "In pair without double clicks",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "In pair with double clicks",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Button 2 and 4 pair mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Separately",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "In pair without double clicks",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "In pair with double clicks",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Action on group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Action on group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Action on group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Action on group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Typical click timeout",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Send the following Switch All commands",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "Switch all on and off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "NC (Normal Close)",
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Invert buttons",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "LED confirmation mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No confirmations",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Confirm button press",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Confirm button press and delivery",
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Send unsolicited Battery Report on Wake Up",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "To same node as wake up notification",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Broadcast to neighbours",
					},
				},
			},
		},
	}
}
func New0175_0100_0102() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "MT2653",
		Description: "Keyfob",

		ManufacturerID: "0175",
		ProductType: "0100",
		ProductID: "0102",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2D,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Button 1 and 3 pair mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Separately",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "In pair without double clicks",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "In pair with double clicks",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Button 2 and 4 pair mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Separately",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "In pair without double clicks",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "In pair with double clicks",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Action on group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Control door lock",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Central scene to gateway",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Action on group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Control door lock",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Central scene to gateway",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Action on group 3",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Control door lock",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Central scene to gateway",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Action on group 4",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Switch On/Off and Dim (send Basic Set and Switch Multilevel)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch On/Off only (send Basic Set)",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Switch all",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Send Scenes",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Send Preconfigured Scenes",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Control devices in proximity",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Control door lock",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Central scene to gateway",
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Send the following Switch All commands",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Switch off only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Switch on only",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Switch all on and off",
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Invert buttons",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Block Wake up",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Wake up is BLOCKED",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Wake up is POSSIBLE according to Intervall",
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Send unsolicited Battery Report on Wake Up",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "To same node as wake up notification",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Broadcast to neighbours",
					},
				},
			},
		},
	}
}
func New0175_0002_0020() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "MT02755",
		Description: "Humidity Sensor",

		ManufacturerID: "0175",
		ProductType: "0002",
		ProductID: "0020",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 5,
				Name: "Operation Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Fahrenheit",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Celsius",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Multi Sensor Function Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 32,
						To: 32,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Using Notification Report",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Using Sensor Binary Report",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 32,
						To: 32,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Auto Report Battery Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Auto Report Temperature Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Auto Report Humidity Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Auto Report Tick Interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Temperature Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Humidity Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 60,
					},
				},
			},
		},
	}
}
func New0175_0002_0021() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "MT2756",
		Description: "Devolo MT2756 Flood Sensor",

		ManufacturerID: "0175",
		ProductType: "0002",
		ProductID: "0021",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set OFF Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Basic Set ON Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Disable the Flood function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Flood function is enabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Flood function is disabled",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Auto Report Battery Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Auto Report Flood Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Auto Report Tick Interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0175_0001_0001() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "PAN11",
		Description: "Smart Energy Plug-in Switch",

		ManufacturerID: "0175",
		ProductType: "0001",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0175_0004_000A() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "ph-pse02",
		Description: "Multisound indoor siren (Zipato/Devolo)",

		ManufacturerID: "0175",
		ProductType: "0004",
		ProductID: "000A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				Version: "0",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x59,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 7,
				Name: "Costumer  Function",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Disable Alarm",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Alarm Duration",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
		},
	}
}
func New0175_0002_0018() *Device{
	return &Device{
		Brand: "Devolo",
		Product: "PST02-1B",
		Description: "Multisensor",

		ManufacturerID: "0175",
		ProductType: "0002",
		ProductID: "0018",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0108_0002_000E() *Device{
	return &Device{
		Brand: "D-Link",
		Product: "DCH-Z110",
		Description: "Door & Window Sensor",

		ManufacturerID: "0108",
		ProductType: "0002",
		ProductID: "000E",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Off",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "On",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Light Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Operation Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 10,
						To: 10,
						Desc: "Preset: Celsius and LED on = Bits: 00001010 = 10",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Multi-Sensor Function Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 20,
						To: 20,
						Desc: "Preset: PIR Super Sensitivity and Binary Report = Bits: 00010100 = 20",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Disable send of BASIC OFF after door closed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Notification Type",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Notification Report",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Sensor Binary Report",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Disable Multi CC in auto report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Disable to report battery state when the device triggered",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Re-Detect Interval Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Turn Off Light Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Auto Report Battery Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Auto Report Door/Window State Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Auto Report Illumination Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Auto Report Temperature Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Auto Report Tick Interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Temperature Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Illumination Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
		},
	}
}
func New0108_0002_000D() *Device{
	return &Device{
		Brand: "D-Link",
		Product: "DCH-Z120",
		Description: "Battery Motion Sensor",

		ManufacturerID: "0108",
		ProductType: "0002",
		ProductID: "000D",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 2,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Off",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "On",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Light Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Operation Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Multi-Sensor Function Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Disable send of BASIC OFF after door closed",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Notification Type",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Notification Report",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Sensor Binary Report",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Disable Multi CC in auto report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Disable to report battery state when the device triggered",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Re-Detect Interval Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Turn Off Light Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Auto Report Battery Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Auto Report Door/Window State Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Auto Report Illumination Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Auto Report Temperature Time",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Auto Report Tick Interval",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Temperature Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Illumination Differential Report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
		},
	}
}
func New0108_0004_000A() *Device{
	return &Device{
		Brand: "D-Link",
		Product: "DCH-Z510",
		Description: "Siren",

		ManufacturerID: "0108",
		ProductType: "0004",
		ProductID: "000A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x71,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 7,
				Name: "Customer Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Using Notification Report",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Using Sensor Binary Report",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Disable Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Enable Alarm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Disable Alarm",
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Alarm Duration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
		},
	}
}
func New020E_4C42_3135() *Device{
	return &Device{
		Brand: "Domitech Products, LLC",
		Product: "ZB22",
		Description: "Smart LED Light Bulb",

		ManufacturerID: "020E",
		ProductType: "4C42",
		ProductID: "3135",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Dim level when the light bulb is turned On",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Full brightness",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Last setting",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Dimming/Brightening Step Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Dimming/Brightening Step Timing",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
		},
	}
}
func New020E_4754_3038() *Device{
	return &Device{
		Brand: "Domitech Products, LLC",
		Product: "DTA19",
		Description: "Smart LED Light",

		ManufacturerID: "020E",
		ProductType: "4754",
		ProductID: "3038",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Dim level when the bulb is turned on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Full brightness",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Last setting",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Dimming / brightening step level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Dimming / brightening step timing",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
		},
	}
}
func New020E_4C42_3134() *Device{
	return &Device{
		Brand: "Domitech Products, LLC",
		Product: "DTA19",
		Description: "Smart LED Light",

		ManufacturerID: "020E",
		ProductType: "4C42",
		ProductID: "3134",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Dim level when the bulb is turned on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Full brightness",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Last setting",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Dimming / brightening step level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Dimming / brightening step timing",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
		},
	}
}
func New0184_4447_3031() *Device{
	return &Device{
		Brand: "Dragon Tech Industrial, Ltd.",
		Product: "PA-100",
		Description: "Plug-in On/Off Switch",

		ManufacturerID: "0184",
		ProductType: "4447",
		ProductID: "3031",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "LED Indicator",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "off/on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "on/off",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "off",
					},
				},
			},
		},
	}
}
func New0184_4447_3034() *Device{
	return &Device{
		Brand: "Dragon Tech Industrial, Ltd.",
		Product: "WD-100",
		Description: "Wall Dimmer Switch",

		ManufacturerID: "0184",
		ProductType: "4447",
		ProductID: "3034",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 4,
				Name: "Orientation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Normal",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Inverted",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Dim Level Increment",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Step Duration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0184_4744_3032() *Device{
	return &Device{
		Brand: "Dragon Tech Industrial, Ltd.",
		Product: "WD-100",
		Description: "Wall Dimmer Switch",

		ManufacturerID: "0184",
		ProductType: "4744",
		ProductID: "3032",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 4,
				Name: "Orientation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Normal",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Inverted",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Dim Level Increment",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Step Duration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New0184_4447_3033() *Device{
	return &Device{
		Brand: "Dragon Tech Industrial, Ltd.",
		Product: "WS-100",
		Description: "Wall On/Off Switch",

		ManufacturerID: "0184",
		ProductType: "4447",
		ProductID: "3033",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "LED Indicator",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "off/on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "on/off",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "off",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Orientation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Normal",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Inverted",
					},
				},
			},
		},
	}
}
func New014A_0001_0002() *Device{
	return &Device{
		Brand: "Ecolink",
		Product: "DWZWAVE2",
		Description: "Z-Wave Door/Window Sensor",

		ManufacturerID: "014A",
		ProductType: "0001",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New014A_0004_0002() *Device{
	return &Device{
		Brand: "Ecolink",
		Product: "DWZWAVE2",
		Description: "Z-Wave Door/Window Sensor",

		ManufacturerID: "014A",
		ProductType: "0004",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New014A_0004_0001() *Device{
	return &Device{
		Brand: "Ecolink",
		Product: "Ecolink PIR v2.5",
		Description: "Z-Wave PIR Motion Sensor v2.5",

		ManufacturerID: "014A",
		ProductType: "0004",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x59,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Send Basic Sets",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not send",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Send Basic Sets",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Send Binary Reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Send Sensor Binary Reports",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Send only Notification Reports",
					},
				},
			},
		},
	}
}
func New014A_0001_0001() *Device{
	return &Device{
		Brand: "Ecolink",
		Product: "Ecolink PIR",
		Description: "Z-Wave PIR Motion Sensor",

		ManufacturerID: "014A",
		ProductType: "0001",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 99,
				Name: "Toggle sending clear message to group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable sending clear message",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send clear message to group 2",
					},
				},
			},
		},
	}
}
func New014A_0004_0003() *Device{
	return &Device{
		Brand: "Ecolink",
		Product: "TILT-ZWAVE2.5-ECO",
		Description: "Z-wave Plus Gold Plated Reliability Garage Door Tilt Sensor",

		ManufacturerID: "014A",
		ProductType: "0004",
		ProductID: "0003",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New014A_0001_0003() *Device{
	return &Device{
		Brand: "Ecolink",
		Product: "TILTZWAVE2",
		Description: "Z-Wave Garage Door Sensor",

		ManufacturerID: "014A",
		ProductType: "0001",
		ProductID: "0003",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0157_0003_0512() *Device{
	return &Device{
		Brand: "EcoNet Controls",
		Product: "EBV-105",
		Description: "Wireless Water Shutoff Valve",

		ManufacturerID: "0157",
		ProductType: "0003",
		ProductID: "0512",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0157_0100_0100() *Device{
	return &Device{
		Brand: "EcoNet Controls",
		Product: "EV100",
		Description: "Z-Vent ZWave Controlled HVAC Air Register",

		ManufacturerID: "0157",
		ProductType: "0100",
		ProductID: "0100",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0157_0003_0002() *Device{
	return &Device{
		Brand: "EcoNet Controls",
		Product: "GR-105N",
		Description: "Z-Wave Motor for water/gas ball valves",

		ManufacturerID: "0157",
		ProductType: "0003",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New021F_0003_0088() *Device{
	return &Device{
		Brand: "Elexa Consumer Products Inc.",
		Product: "DMS01",
		Description: "Dome Wireless Siren",

		ManufacturerID: "021F",
		ProductType: "0003",
		ProductID: "0088",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "6",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x87,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New011A_0101_0102() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "Enerwave ZW15S",
		Description: "Enerwave ZW15S Binary Switch",

		ManufacturerID: "011A",
		ProductType: "0101",
		ProductID: "0102",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New011A_0101_0103() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZW15R ",
		Description: "Duplex receptacle",

		ManufacturerID: "011A",
		ProductType: "0101",
		ProductID: "0103",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New011A_0101_0603() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZW20R",
		Description: "20A Tamper Resistant Duplex Receptacle",

		ManufacturerID: "011A",
		ProductType: "0101",
		ProductID: "0603",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Toggle LED mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "0: LED is on when switch is off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "1: LED is on when switch is on",
					},
				},
			},
		},
	}
}
func New011A_0111_0101() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZW20RM",
		Description: "In-wall Smart Meter  Duplex Receptacle",

		ManufacturerID: "011A",
		ProductType: "0111",
		ProductID: "0101",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Power on, LED off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Power on, LED on",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Not report (METER_REPORT)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send (METER_REPORT) only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send (SENSOR_MULTI_LEVEL_REPORT) only",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send both",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New011A_0102_0201() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZW500D",
		Description: "Dimmer",

		ManufacturerID: "011A",
		ProductType: "0102",
		ProductID: "0201",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "configure LED light state",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED off with load on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED on with load on",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Reverse installed direction",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "normal install - up is on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "invert switch",
					},
				},
			},
		},
	}
}
func New011A_0111_0201() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZW500DM",
		Description: "In-wall Smart Meter Dimmer Switch",

		ManufacturerID: "011A",
		ProductType: "0111",
		ProductID: "0201",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Power on, LED off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Power on, LED on",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Not report (METER_REPORT)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send (METER_REPORT) only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Send (SENSOR_MULTI_LEVEL_REPORT) only",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Send both",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New011A_0111_0605() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZWM-RSM1",
		Description: "Smart Single Relay Switch Module",

		ManufacturerID: "011A",
		ProductType: "0111",
		ProductID: "0605",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New011A_0101_0104() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZWN-333",
		Description: "Plug-in Appliance Module",

		ManufacturerID: "011A",
		ProductType: "0101",
		ProductID: "0104",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New011A_0601_0901() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZWN-BPC",
		Description: "PIR Sensor",

		ManufacturerID: "011A",
		ProductType: "0601",
		ProductID: "0901",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				Name: "Off time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New011A_0101_5605() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZWN-RSM1-PLUS",
		Description: "Smart Relay Switch Module",

		ManufacturerID: "011A",
		ProductType: "0101",
		ProductID: "5605",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Unsolicited Report Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New011A_0101_5606() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZWN-RSM2",
		Description: "Smart Dual Relay Switch Module",

		ManufacturerID: "011A",
		ProductType: "0101",
		ProductID: "5606",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Unsolicited Report Configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
		},
	}
}
func New011A_0801_0B03() *Device{
	return &Device{
		Brand: "Wenzhou MTLC Electric Appliances Co.,Ltd.",
		Product: "ZWN-SC7",
		Description: "Scene Controller",

		ManufacturerID: "011A",
		ProductType: "0801",
		ProductID: "0B03",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0148_0002_0001() *Device{
	return &Device{
		Brand: "Eurotronics",
		Product: "CometZ",
		Description: "Thermostatic Valve",

		ManufacturerID: "0148",
		ProductType: "0002",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x40,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x77,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0148_0001_0001() *Device{
	return &Device{
		Brand: "Eurotronics",
		Product: "StellaZ",
		Description: "Thermostatic Valve",

		ManufacturerID: "0148",
		ProductType: "0001",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x40,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x43,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0135_000B_0001() *Device{
	return &Device{
		Brand: "Everspring (135)",
		Product: "ST812",
		Description: "Flood sensor",

		ManufacturerID: "0135",
		ProductType: "000B",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Level Set",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "OFF",
					},
				},
			},
		},
	}
}
func New0060_0003_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AD130",
		Description: "Dimmer Plugin",

		ManufacturerID: "0060",
		ProductType: "0003",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0060_0003_0002() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AD146",
		Description: "In-Wall Dimmer Module",

		ManufacturerID: "0060",
		ProductType: "0003",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Command value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Report time delay group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 3,
						To: 25,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Remember last status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not remember",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Remember",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Switching type",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Edge Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Toogle Mode",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Output mode setting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Dimming",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "On/Off",
					},
				},
			},
		},
	}
}
func New0060_0003_0003() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AD147",
		Description: "Z-Wave Dimmer Plug",

		ManufacturerID: "0060",
		ProductType: "0003",
		ProductID: "0003",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Command value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Delaying Time to report to Group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 3,
						To: 25,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Remember the last status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Output mode setting: Dimming, On/Off",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0060_0004_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AN157",
		Description: "Switch Plugin",

		ManufacturerID: "0060",
		ProductType: "0004",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0060_0104_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AN157",
		Description: "Switch Plugin",

		ManufacturerID: "0060",
		ProductType: "0104",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0060_0004_0002() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AN158",
		Description: "Switch Meter Plugin",

		ManufacturerID: "0060",
		ProductType: "0004",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "True Period",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 254,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Send Out Basic Command",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Meter Report Period",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 3240,
					},
				},
			},
		},
	}
}
func New0060_0004_0005() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AN163",
		Description: "Metering Power Switch",

		ManufacturerID: "0060",
		ProductType: "0004",
		ProductID: "0005",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "True Period",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 120,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Send Basic Command to Group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Meter Report Period (W)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 3240,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Meter Report Period (kWh)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 10080,
					},
				},
			},
		},
	}
}
func New0060_0004_0008() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AN179",
		Description: "In-Wall Switch Module",

		ManufacturerID: "0060",
		ProductType: "0004",
		ProductID: "0008",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Command value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Report time delay group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 3,
						To: 25,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Remember last status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not remember",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Remember",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Switching type",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Edge Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Toogle Mode",
					},
				},
			},
		},
	}
}
func New0060_0011_0002() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AN179",
		Description: "In-Wall Switch Module",

		ManufacturerID: "0060",
		ProductType: "0011",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Command value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Report time delay group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 3,
						To: 25,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Remember last status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not remember",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Remember",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Switching type",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Edge Mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Toogle Mode",
					},
				},
			},
		},
	}
}
func New0060_0004_0007() *Device{
	return &Device{
		Brand: "Everspring",
		Product: " AN1802",
		Description: "Switch without metering",

		ManufacturerID: "0060",
		ProductType: "0004",
		ProductID: "0007",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0060_0004_0006() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "AN181",
		Description: "Mini Plug Switch with Metering (Gen5)",

		ManufacturerID: "0060",
		ProductType: "0004",
		ProductID: "0006",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x32,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Command value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No Basic Set Command will be sent",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "Basic Set Command ON will be sent",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Delaying time",
				Size: 1,
				Values: []Value{
					Value{
						From: 3,
						To: 25,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Remember status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not remember",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Remember",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Interval for wattage auto report",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Does not report automatically",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Interval for kW*h auto report",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Does not report automatically",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Auto report load surpasses value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Does not report automatically",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Auto report change percentage",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Does not report automatically",
					},
				},
			},
		},
	}
}
func New0060_0010_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "HAC01",
		Description: "In-Wall Remote Insert",

		ManufacturerID: "0060",
		ProductType: "0010",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Amount Of Delay",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
		},
	}
}
func New0060_0202_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "HSM02",
		Description: "Door/Window Contact",

		ManufacturerID: "0060",
		ProductType: "0202",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Configuring the OFF Delay",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Staying Awake (for testing)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0060_0002_0002() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "HSM02",
		Description: "Door/Window Contact",

		ManufacturerID: "0060",
		ProductType: "0002",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Configuring the OFF Delay",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Staying Awake (for testing)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
		},
	}
}
func New0060_0001_0003() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "HSP02",
		Description: "Motion Detector",

		ManufacturerID: "0060",
		ProductType: "0001",
		ProductID: "0003",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Sensor Detecting Function",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 Enabled / Group 2 Enabled",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 1 Enabled / Group 2 Disabled",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Group 1 Disabled / Group 2 Disabled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Sensitivity Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Re-trigger Interval Setting",
				Size: 1,
				Values: []Value{
					Value{
						From: 5,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "LUX Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "On-Off Duration",
				Size: 1,
				Values: []Value{
					Value{
						From: 5,
						To: 255,
					},
				},
			},
		},
	}
}
func New0060_000C_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "SE812",
		Description: "Siren",

		ManufacturerID: "0060",
		ProductType: "000C",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0060_000D_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "SF812",
		Description: "Smoke Sensor",

		ManufacturerID: "0060",
		ProductType: "000D",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0060_0002_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "SM103",
		Description: "Door/Window Contact",

		ManufacturerID: "0060",
		ProductType: "0002",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Configuring the OFF Delay",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 127,
					},
				},
			},
		},
	}
}
func New0060_0101_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "SP103",
		Description: "Motion Detector",

		ManufacturerID: "0060",
		ProductType: "0101",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Phase level on command",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Set OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Set ON",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Set LAST Level",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable/Disable Sleeping",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
		},
	}
}
func New0060_0001_0002() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "SP814",
		Description: "Motion Detector",

		ManufacturerID: "0060",
		ProductType: "0001",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Enable/Disable Detecting",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Sensitivity Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 10,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Re-trigger Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 5,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Lux Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "On-Off Duration",
				Size: 1,
				Values: []Value{
					Value{
						From: 5,
						To: 255,
					},
				},
			},
		},
	}
}
func New0060_000B_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "ST812",
		Description: "Flood Sensor",

		ManufacturerID: "0060",
		ProductType: "000B",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Level Set",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
		},
	}
}
func New0060_0006_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "ST814",
		Description: "Temperature and Humidity Sensor",

		ManufacturerID: "0060",
		ProductType: "0006",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Temperature Trigger ON value",
				Size: 1,
				Values: []Value{
					Value{
						From: 99,
						To: 99,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Temperature Trigger OFF value",
				Size: 1,
				Values: []Value{
					Value{
						From: 99,
						To: 99,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Humidity Trigger ON value",
				Size: 1,
				Values: []Value{
					Value{
						From: 99,
						To: 99,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Humidity Trigger OFF value",
				Size: 1,
				Values: []Value{
					Value{
						From: 99,
						To: 99,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Auto report time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Auto report temperature",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Auto report humidity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
				},
			},
		},
	}
}
func New0060_0007_0001() *Device{
	return &Device{
		Brand: "Everspring",
		Product: "ST815",
		Description: "Illumination Sensor",

		ManufacturerID: "0060",
		ProductType: "0007",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Basic Set Level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Lux Trigger On Value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Lux Trigger Off Value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Lux Trigger Off Timer Value.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 480,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Auto Report Time Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1439,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Auto Report Lux Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
		},
	}
}
func New0113_4450_3030() *Device{
	return &Device{
		Brand: "Evolve",
		Product: "LDM-15W",
		Description: "Lamp Module",

		ManufacturerID: "0113",
		ProductType: "4450",
		ProductID: "3030",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0113_5246_3133() *Device{
	return &Device{
		Brand: "Evolve",
		Product: "LFM-20",
		Description: "",

		ManufacturerID: "0113",
		ProductType: "5246",
		ProductID: "3133",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
		},
	}
}
func New0113_5250_3030() *Device{
	return &Device{
		Brand: "Evolve",
		Product: "LPM-15",
		Description: "Appliance module",

		ManufacturerID: "0113",
		ProductType: "5250",
		ProductID: "3030",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0113_4457_3034() *Device{
	return &Device{
		Brand: "Evolve",
		Product: "LRM-AS",
		Description: "Wall Mounted Dimmer",

		ManufacturerID: "0113",
		ProductType: "4457",
		ProductID: "3034",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Invert Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Yes",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Load Sense",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
		},
	}
}
func New0113_5257_3533() *Device{
	return &Device{
		Brand: "Evolve",
		Product: "LSM-15",
		Description: "Wall Mounted Switch",

		ManufacturerID: "0113",
		ProductType: "5257",
		ProductID: "3533",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 3,
				Name: "Night Light",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "On when load is on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Off when load is on",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Invert Switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Top-ON, Bottom-OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bottom--ON, Top-OFF",
					},
				},
			},
		},
	}
}
func New0113_4556_5434() *Device{
	return &Device{
		Brand: "Evolve",
		Product: "T100",
		Description: "Thermostat",

		ManufacturerID: "0113",
		ProductType: "4556",
		ProductID: "5434",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x40,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x42,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x43,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x44,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x45,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x81,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0085_0002_0002() *Device{
	return &Device{
		Brand: "Fakro",
		Product: "ARZ",
		Description: "Roller Shutter Module",

		ManufacturerID: "0085",
		ProductType: "0002",
		ProductID: "0002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0085_0002_0001() *Device{
	return &Device{
		Brand: "Fakro",
		Product: "ZWS12",
		Description: "Chain Actuator",

		ManufacturerID: "0085",
		ProductType: "0002",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New0085_0003_0001() *Device{
	return &Device{
		Brand: "Fakro",
		Product: "ZWS12",
		Description: "Chain Actuator",

		ManufacturerID: "0085",
		ProductType: "0003",
		ProductID: "0001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
		},
	}
}
func New010F_0501_0102() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGBS001",
		Description: "Universal Binary Sensor",

		ManufacturerID: "010F",
		ProductType: "0501",
		ProductID: "0102",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "IN1 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "IN2 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Type of input no. 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Type of input no. 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of transmitted control frame for association group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Type of transmitted control frame for association group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Forced Level of Dimming group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Deactivate transmission of frame canceling alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Groups 1 and 2 sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 sent, Group 2 not sent.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 1 not sent, Group 2 sent.",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Interval between successive readings of temperature sensors",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Interval betw. forcing to send report concerning the temp.",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Insensitiveness to temperature changes.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Transmitting the alarm or control frame broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "IN1 and IN2 Broadcast inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "IN1 broadcast mode active, Sensor 2 broadcast mode inactive",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "IN1 broadcast mode inactive, Sensor 2 broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "INI and IN2 broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Scene activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scenes disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scenes enabled",
					},
				},
			},
		},
	}
}
func New010F_0501_1002() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGBS001",
		Description: "Universal Binary Sensor",

		ManufacturerID: "010F",
		ProductType: "0501",
		ProductID: "1002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "IN1 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "IN2 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Type of input no. 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Type of input no. 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of transmitted control frame for association group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Type of transmitted control frame for association group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Forced Level of Dimming group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Deactivate transmission of frame canceling alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Groups 1 and 2 sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 sent, Group 2 not sent.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 1 not sent, Group 2 sent.",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Interval between successive readings of temperature sensors",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Interval betw. forcing to send report concerning the temp.",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Insensitiveness to temperature changes.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Transmitting the alarm or control frame broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "IN1 and IN2 Broadcast inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "IN1 broadcast mode active, Sensor 2 broadcast mode inactive",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "IN1 broadcast mode inactive, Sensor 2 broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "INI and IN2 broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Scene activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scenes disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scenes enabled",
					},
				},
			},
		},
	}
}
func New010F_0501_2002() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGBS001",
		Description: "Universal Binary Sensor",

		ManufacturerID: "010F",
		ProductType: "0501",
		ProductID: "2002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "IN1 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "IN2 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Type of input no. 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Type of input no. 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of transmitted control frame for association group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Type of transmitted control frame for association group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Forced Level of Dimming group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Deactivate transmission of frame canceling alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Groups 1 and 2 sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 sent, Group 2 not sent.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 1 not sent, Group 2 sent.",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Interval between successive readings of temperature sensors",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Interval betw. forcing to send report concerning the temp.",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Insensitiveness to temperature changes.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Transmitting the alarm or control frame broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "IN1 and IN2 Broadcast inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "IN1 broadcast mode active, Sensor 2 broadcast mode inactive",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "IN1 broadcast mode inactive, Sensor 2 broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "INI and IN2 broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Scene activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scenes disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scenes enabled",
					},
				},
			},
		},
	}
}
func New010F_0501_3002() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGBS001",
		Description: "Universal Binary Sensor",

		ManufacturerID: "010F",
		ProductType: "0501",
		ProductID: "3002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "IN1 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "IN2 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Type of input no. 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Type of input no. 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "NO (Normal Open)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "NC (Normal Close)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "MONOSTABLE",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "BISTABLE",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of transmitted control frame for association group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Type of transmitted control frame for association group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Forced Level of Dimming group 2",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Deactivate transmission of frame canceling alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Groups 1 and 2 sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 sent, Group 2 not sent.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 1 not sent, Group 2 sent.",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Interval between successive readings of temperature sensors",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Interval betw. forcing to send report concerning the temp.",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Insensitiveness to temperature changes.",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Transmitting the alarm or control frame broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "IN1 and IN2 Broadcast inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "IN1 broadcast mode active, Sensor 2 broadcast mode inactive",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "IN1 broadcast mode inactive, Sensor 2 broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "INI and IN2 broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Scene activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scenes disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scenes enabled",
					},
				},
			},
		},
	}
}
func New010F_0D01_1000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGGC001",
		Description: "Fibaro Swipe Scene Controller",

		ManufacturerID: "010F",
		ProductType: "0D01",
		ProductID: "1000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Device Orientation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Default Orientation",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "180° rotation",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "90° clockwise rotation",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "90° counter-clockwise rotation",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Buzzer - acoustic signal settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Gestures detection is not signalled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Gestures detection is signalled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: " LED diode - visual indicator settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Gestures detection is not indicated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: " Gestures detection is indicated",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Buzzer - signalling result of gesture recognition",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Only successful recognition is signalled",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only failed recognition is signalled",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Successful and failed recognition is signalled",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Powering mode - interval of updating the current mode",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Powering mode is not updated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Time interval (in minutes)",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Power saving mode (battery mode)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Standby mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Simple mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Swipe does not enter power saving mode",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Hold gesture to enter the menu",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Hold gesture to enter the menu enabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hold gesture to enter the menu disabled",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Scenes sent to the controller",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Scenes for flick UP gesture enabled",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Scenes for flick DOWN gesture enabled",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Scenes for flick LEFT gesture enabled",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Scenes for flick RIGHT gesture enabled",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Scenes for clockwise circular gesture enabled",
					},
					Value{
						From: 32,
						To: 32,
						Desc: "Scenes for counter-clockwise circular gesture enabled",
					},
					Value{
						From: 63,
						To: 63,
						Desc: "All commands enabled",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group ”Flick UP” sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group ”Flick DOWN” sent as secure",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "4th group \"Flick LEFT” sent as secure",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "5th group \"Flick RIGHT” sent as secure",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "6th group \"Circular AirWheel” sent as secure",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Control mode of 2nd - 5th \"Flick UP/DOWN/LEFT/RIGHT” association groups and scenes",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Toggle Mode enabled for 2nd association group",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Toggle Mode enabled for 3rd association group",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Toggle Mode enabled for 4th association group",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Toggle Mode enabled for 5th association group",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Rate of smooth level control",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "SWITCH ON control frame value for FLICK UP gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "SWITCH OFF control frame value for FLICK UP gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "SWITCH ON control frame value for FLICK DOWN gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "SWITCH OFF control frame value for FLICK DOWN gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "SWITCH ON control frame value for FLICK LEFT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "SWITCH OFF control frame value for FLICK LEFT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "SWITCH ON control frame value for FLICK RIGHT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 27,
				Name: "SWITCH OFF control frame value for FLICK RIGHT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Sequence learning mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 6,
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "1st gestures sequence (SLOT 1)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 32,
				Name: "2nd gestures sequence (SLOT 2)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "3rd gestures sequence (SLOT 3)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "4th gestures sequence (SLOT 4)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "5th gestures sequence (SLOT 5)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "6th gestures sequence (SLOT 6)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
		},
	}
}
func New010F_0D01_3000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGGC001",
		Description: "Fibaro Swipe Scene Controller",

		ManufacturerID: "010F",
		ProductType: "0D01",
		ProductID: "3000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Device Orientation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Default Orientation",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "180° rotation",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "90° clockwise rotation",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "90° counter-clockwise rotation",
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Buzzer - acoustic signal settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Gestures detection is not signalled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Gestures detection is signalled",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: " LED diode - visual indicator settings",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Gestures detection is not indicated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: " Gestures detection is indicated",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Buzzer - signalling result of gesture recognition",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Only successful recognition is signalled",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only failed recognition is signalled",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Successful and failed recognition is signalled",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Powering mode - interval of updating the current mode",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Powering mode is not updated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Time interval (in minutes)",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Power saving mode (battery mode)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Standby mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Simple mode",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Swipe does not enter power saving mode",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Hold gesture to enter the menu",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Hold gesture to enter the menu enabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Hold gesture to enter the menu disabled",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Scenes sent to the controller",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Scenes for flick UP gesture enabled",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Scenes for flick DOWN gesture enabled",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Scenes for flick LEFT gesture enabled",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Scenes for flick RIGHT gesture enabled",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Scenes for clockwise circular gesture enabled",
					},
					Value{
						From: 32,
						To: 32,
						Desc: "Scenes for counter-clockwise circular gesture enabled",
					},
					Value{
						From: 63,
						To: 63,
						Desc: "All commands enabled",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group ”Flick UP” sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group ”Flick DOWN” sent as secure",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "4th group \"Flick LEFT” sent as secure",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "5th group \"Flick RIGHT” sent as secure",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "6th group \"Circular AirWheel” sent as secure",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Control mode of 2nd - 5th \"Flick UP/DOWN/LEFT/RIGHT” association groups and scenes",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 1,
						Desc: "Toggle Mode enabled for 2nd association group",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Toggle Mode enabled for 3rd association group",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Toggle Mode enabled for 4th association group",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Toggle Mode enabled for 5th association group",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Rate of smooth level control",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "SWITCH ON control frame value for FLICK UP gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "SWITCH OFF control frame value for FLICK UP gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "SWITCH ON control frame value for FLICK DOWN gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "SWITCH OFF control frame value for FLICK DOWN gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "SWITCH ON control frame value for FLICK LEFT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "SWITCH OFF control frame value for FLICK LEFT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "SWITCH ON control frame value for FLICK RIGHT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 27,
				Name: "SWITCH OFF control frame value for FLICK RIGHT gesture",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Sequence learning mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 6,
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "1st gestures sequence (SLOT 1)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 32,
				Name: "2nd gestures sequence (SLOT 2)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "3rd gestures sequence (SLOT 3)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "4th gestures sequence (SLOT 4)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "5th gestures sequence (SLOT 5)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
			&parameter{
				ID: 36,
				Name: "6th gestures sequence (SLOT 6)",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1076,
					},
				},
			},
		},
	}
}
func New010F_0100_0104() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD211",
		Description: "Universal Dimmer 500W",

		ManufacturerID: "010F",
		ProductType: "0100",
		ProductID: "0104",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Enable/Disable ALL ON/OFF",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "ALL ON active / ALL OFF active",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON disabled / ALL OFF disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON disabled / ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active / ALL OFF disabled",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Separation of association sending (key 1)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Map status to all devices in group 1",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control key #2 behaviour",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device status is not checked",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status is checked",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming step at automatic control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Time of MANUALLY moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Time of AUTOMATIC moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Dimming step at manual control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Maximum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 2,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Minimum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Parm 15",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Synchronizing light level for associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Parm 20",
				Size: 1,
				Values: []Value{
					Value{
						From: 100,
						To: 170,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Relay 1: Response to General Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "DEACTIVATION - no response to alarm frames",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM RELAY ON - relay will turn ON upon receipt of alarm frame",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM RELAY OFF - relay will turn OFF upon receipt of alarm frame",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM FLASHING - relay will turn ON and OFF periodically (see param.39",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "ALARM FLASHING alarm time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New010F_0100_0106() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD211",
		Description: "Universal Dimmer 500W",

		ManufacturerID: "010F",
		ProductType: "0100",
		ProductID: "0106",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Enable/Disable ALL ON/OFF",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "ALL ON active / ALL OFF active",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON disabled / ALL OFF disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON disabled / ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active / ALL OFF disabled",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Separation of association sending (key 1)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Map status to all devices in group 1",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control key #2 behaviour",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device status is not checked",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status is checked",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming step at automatic control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Time of MANUALLY moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Time of AUTOMATIC moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Dimming step at manual control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Maximum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 2,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Minimum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Parm 15",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Synchronizing light level for associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Parm 20",
				Size: 1,
				Values: []Value{
					Value{
						From: 100,
						To: 170,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Relay 1: Response to General Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "DEACTIVATION - no response to alarm frames",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM RELAY ON - relay will turn ON upon receipt of alarm frame",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM RELAY OFF - relay will turn OFF upon receipt of alarm frame",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM FLASHING - relay will turn ON and OFF periodically (see param.39",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "ALARM FLASHING alarm time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New010F_0100_0107() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD211",
		Description: "Universal Dimmer 500W",

		ManufacturerID: "010F",
		ProductType: "0100",
		ProductID: "0107",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Enable/Disable ALL ON/OFF",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "ALL ON active / ALL OFF active",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON disabled / ALL OFF disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON disabled / ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active / ALL OFF disabled",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Separation of association sending (key 1)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Map status to all devices in group 1",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control key #2 behaviour",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device status is not checked",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status is checked",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming step at automatic control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Time of MANUALLY moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Time of AUTOMATIC moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Dimming step at manual control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Maximum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 2,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Minimum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Parm 15",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Synchronizing light level for associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Parm 20",
				Size: 1,
				Values: []Value{
					Value{
						From: 100,
						To: 170,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Relay 1: Response to General Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "DEACTIVATION - no response to alarm frames",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM RELAY ON - relay will turn ON upon receipt of alarm frame",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM RELAY OFF - relay will turn OFF upon receipt of alarm frame",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM FLASHING - relay will turn ON and OFF periodically (see param.39",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "ALARM FLASHING alarm time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New010F_0100_0109() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD211",
		Description: "Universal Dimmer 500W",

		ManufacturerID: "010F",
		ProductType: "0100",
		ProductID: "0109",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Enable/Disable ALL ON/OFF",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "ALL ON active / ALL OFF active",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON disabled / ALL OFF disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON disabled / ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active / ALL OFF disabled",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Separation of association sending (key 1)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Map status to all devices in group 1",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control key #2 behaviour",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device status is not checked",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status is checked",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming step at automatic control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Time of MANUALLY moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Time of AUTOMATIC moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Dimming step at manual control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Maximum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 2,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Minimum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Parm 15",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Synchronizing light level for associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Parm 20",
				Size: 1,
				Values: []Value{
					Value{
						From: 100,
						To: 170,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Relay 1: Response to General Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "DEACTIVATION - no response to alarm frames",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM RELAY ON - relay will turn ON upon receipt of alarm frame",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM RELAY OFF - relay will turn OFF upon receipt of alarm frame",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM FLASHING - relay will turn ON and OFF periodically (see param.39",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "ALARM FLASHING alarm time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New010F_0100_100A() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD211",
		Description: "Universal Dimmer 500W",

		ManufacturerID: "010F",
		ProductType: "0100",
		ProductID: "100A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Enable/Disable ALL ON/OFF",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "ALL ON active / ALL OFF active",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON disabled / ALL OFF disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON disabled / ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active / ALL OFF disabled",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Separation of association sending (key 1)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Map status to all devices in group 1",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control key #2 behaviour",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device status is not checked",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status is checked",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming step at automatic control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Time of MANUALLY moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Time of AUTOMATIC moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Dimming step at manual control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Maximum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 2,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Minimum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Parm 15",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Synchronizing light level for associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Parm 20",
				Size: 1,
				Values: []Value{
					Value{
						From: 100,
						To: 170,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Relay 1: Response to General Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "DEACTIVATION - no response to alarm frames",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM RELAY ON - relay will turn ON upon receipt of alarm frame",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM RELAY OFF - relay will turn OFF upon receipt of alarm frame",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM FLASHING - relay will turn ON and OFF periodically (see param.39",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "ALARM FLASHING alarm time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New010F_0100_300A() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD211",
		Description: "Universal Dimmer 500W",

		ManufacturerID: "010F",
		ProductType: "0100",
		ProductID: "300A",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Enable/Disable ALL ON/OFF",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "ALL ON active / ALL OFF active",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON disabled / ALL OFF disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON disabled / ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active / ALL OFF disabled",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Separation of association sending (key 1)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Map status to all devices in group 1",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Map OFF status to all devices in group 1, Double click on key 1 will s",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Control key #2 behaviour",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device status is not checked",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status is checked",
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming step at automatic control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Time of MANUALLY moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Time of AUTOMATIC moving between the extreme dimming values",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Dimming step at manual control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Maximum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 2,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Minimum dimmer level control",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Parm 15",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Synchronizing light level for associated devices",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Parm 20",
				Size: 1,
				Values: []Value{
					Value{
						From: 100,
						To: 170,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Relay 1: Response to General Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "DEACTIVATION - no response to alarm frames",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM RELAY ON - relay will turn ON upon receipt of alarm frame",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM RELAY OFF - relay will turn OFF upon receipt of alarm frame",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM FLASHING - relay will turn ON and OFF periodically (see param.39",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "ALARM FLASHING alarm time",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
		},
	}
}
func New010F_0102_1000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD212",
		Description: "Dimmer 2",

		ManufacturerID: "010F",
		ProductType: "0102",
		ProductID: "1000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Minimum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 98,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Maximum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Incandescence level of dimmable compact fluorescent lamps",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Incandescence time of dimmable compact fluorescent lamps",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Dimming Step at Automatic Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Dimming Step Timing at Automatic Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimming Step at Manual Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming Step Timing at Manual Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Timer functionality (auto - off). 0 disables the function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "ALL ON/ALL OFF function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON not active, ALL OFF not active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON not active, ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active, ALL OFF not active",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "ALL ON active, ALL OFF active",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Force auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "readout",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "force auto-calibration of the load without Fibaro Bypass",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "force auto-calibration of the load with Fibaro Bypass",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Auto-calibration status (read-only parameter)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "calibration procedure not performed or Dimmer operates on manual setti",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Dimmer operates on auto-calibration settings",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Burnt out bulb detection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Time delay of a burnt out bulb (parameter 15) or overload (parameter 39) detection",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Forced switch on brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Roller blind switch (UP / DOWN)",
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Value sent to associated devices on single click",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "0xFF value is sent, which will set associated devices to their last sa",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "current Dimmer 2 state is sent, which will synchronize brightness leve",
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Double click option",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Command frames sent in 2nd and 3rd association groups (S1 associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Command frames sent in 4th and 5th association groups (S2associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "The function of 3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "3-way switch function for S2 disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "3-way switch function for S2 enabled",
					},
				},
			},
			&parameter{
				ID: 27,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all groups (II-V) sent as non-secure",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group sent as secure",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "4th group sent as secure",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "5th group sent as secure",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "all groups (II-V) sent as secure",
					},
				},
			},
			&parameter{
				ID: 28,
				Name: "Scene activation functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scene functionality deactivated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene functionality activated",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Switch functionality of S1 and S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "standard mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "S1 operates as S2, S2 operates as S1",
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Load control mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "forced leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "forced trailing edge control",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control mode selected automatically (based on auto-calibration)",
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Load control mode recognized during auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "trailing edge control",
					},
				},
			},
			&parameter{
				ID: 32,
				Name: "On/Off mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "on/off mode disabled (dimming is possible)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "on/off mode enabled (dimming is not possible)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "mode selected automatically",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "Dimmability of the load (read only)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Load recognized as dimmable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Load recognized as non-dimmable",
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Soft-Start functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "no soft-start",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "short soft-start (0,1s)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "long soft-start (0,5s)",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Auto-calibration after power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No auto-calibration of the load after power on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Auto-calibration performed after first power on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Auto-calibration performed after each power on",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Auto-calibration performed after first power on or after each LOAD ERR",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Auto-calibration performed after each power on or after each LOAD ERRO",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Behaviour of the Dimmer after overload or surge",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "device permanently disabled until re-enabling by comand or push-button",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "three attempts to turn on the load",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Power limit - OVERLOAD",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 250,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "General Purpose Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Water Flooding Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Smoke, CO or CO2 Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Temperature Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Time of alarm state",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "OVERLOAD alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "LOAD ERROR alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 47,
				Name: "OVERCURRENT alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 48,
				Name: "SURGE alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 49,
				Name: "OVERHEAT (critical temperature) and VOLTAGE DROP (low voltage) alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Active power reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 52,
				Name: "Periodic active power and energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 53,
				Name: "Energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 54,
				Name: "Self-measurement",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Self-measurement inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Self-measurement active",
					},
				},
			},
			&parameter{
				ID: 58,
				Name: "Method of calculating the active power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "measurement based on the standard algorithm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "approximation based on the calibration data",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "approximation based on the control angle",
					},
				},
			},
			&parameter{
				ID: 59,
				Name: "Approximated power at the maximum brightness level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 500,
					},
				},
			},
		},
	}
}
func New010F_0102_2000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD212",
		Description: "Dimmer 2",

		ManufacturerID: "010F",
		ProductType: "0102",
		ProductID: "2000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Minimum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 98,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Maximum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Incandescence level of dimmable compact fluorescent lamps",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Incandescence time of dimmable compact fluorescent lamps",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Dimming Step at Automatic Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Dimming Step Timing at Automatic Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimming Step at Manual Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming Step Timing at Manual Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Timer functionality (auto - off). 0 disables the function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "ALL ON/ALL OFF function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON not active, ALL OFF not active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON not active, ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active, ALL OFF not active",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "ALL ON active, ALL OFF active",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Force auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "readout",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "force auto-calibration of the load without Fibaro Bypass",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "force auto-calibration of the load with Fibaro Bypass",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Auto-calibration status (read-only parameter)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "calibration procedure not performed or Dimmer operates on manual setti",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Dimmer operates on auto-calibration settings",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Burnt out bulb detection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Time delay of a burnt out bulb (parameter 15) or overload (parameter 39) detection",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Forced switch on brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Roller blind switch (UP / DOWN)",
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Value sent to associated devices on single click",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "0xFF value is sent, which will set associated devices to their last sa",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "current Dimmer 2 state is sent, which will synchronize brightness leve",
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Double click option",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Command frames sent in 2nd and 3rd association groups (S1 associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Command frames sent in 4th and 5th association groups (S2associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "The function of 3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "3-way switch function for S2 disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "3-way switch function for S2 enabled",
					},
				},
			},
			&parameter{
				ID: 27,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all groups (II-V) sent as non-secure",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group sent as secure",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "4th group sent as secure",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "5th group sent as secure",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "all groups (II-V) sent as secure",
					},
				},
			},
			&parameter{
				ID: 28,
				Name: "Scene activation functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scene functionality deactivated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene functionality activated",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Switch functionality of S1 and S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "standard mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "S1 operates as S2, S2 operates as S1",
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Load control mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "forced leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "forced trailing edge control",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control mode selected automatically (based on auto-calibration)",
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Load control mode recognized during auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "trailing edge control",
					},
				},
			},
			&parameter{
				ID: 32,
				Name: "On/Off mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "on/off mode disabled (dimming is possible)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "on/off mode enabled (dimming is not possible)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "mode selected automatically",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "Dimmability of the load (read only)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Load recognized as dimmable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Load recognized as non-dimmable",
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Soft-Start functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "no soft-start",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "short soft-start (0,1s)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "long soft-start (0,5s)",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Auto-calibration after power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No auto-calibration of the load after power on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Auto-calibration performed after first power on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Auto-calibration performed after each power on",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Auto-calibration performed after first power on or after each LOAD ERR",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Auto-calibration performed after each power on or after each LOAD ERRO",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Behaviour of the Dimmer after overload or surge",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "device permanently disabled until re-enabling by comand or push-button",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "three attempts to turn on the load",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Power limit - OVERLOAD",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 250,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "General Purpose Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Water Flooding Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Smoke, CO or CO2 Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Temperature Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Time of alarm state",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "OVERLOAD alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "LOAD ERROR alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 47,
				Name: "OVERCURRENT alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 48,
				Name: "SURGE alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 49,
				Name: "OVERHEAT (critical temperature) and VOLTAGE DROP (low voltage) alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Active power reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 52,
				Name: "Periodic active power and energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 53,
				Name: "Energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 54,
				Name: "Self-measurement",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Self-measurement inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Self-measurement active",
					},
				},
			},
			&parameter{
				ID: 58,
				Name: "Method of calculating the active power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "measurement based on the standard algorithm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "approximation based on the calibration data",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "approximation based on the control angle",
					},
				},
			},
			&parameter{
				ID: 59,
				Name: "Approximated power at the maximum brightness level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 500,
					},
				},
			},
		},
	}
}
func New010F_0102_3000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD212",
		Description: "Dimmer 2",

		ManufacturerID: "010F",
		ProductType: "0102",
		ProductID: "3000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Minimum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 98,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Maximum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Incandescence level of dimmable compact fluorescent lamps",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Incandescence time of dimmable compact fluorescent lamps",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Dimming Step at Automatic Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Dimming Step Timing at Automatic Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimming Step at Manual Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming Step Timing at Manual Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Timer functionality (auto - off). 0 disables the function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "ALL ON/ALL OFF function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON not active, ALL OFF not active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON not active, ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active, ALL OFF not active",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "ALL ON active, ALL OFF active",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Force auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "readout",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "force auto-calibration of the load without Fibaro Bypass",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "force auto-calibration of the load with Fibaro Bypass",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Auto-calibration status (read-only parameter)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "calibration procedure not performed or Dimmer operates on manual setti",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Dimmer operates on auto-calibration settings",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Burnt out bulb detection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Time delay of a burnt out bulb (parameter 15) or overload (parameter 39) detection",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Forced switch on brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Roller blind switch (UP / DOWN)",
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Value sent to associated devices on single click",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "0xFF value is sent, which will set associated devices to their last sa",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "current Dimmer 2 state is sent, which will synchronize brightness leve",
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Double click option",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Command frames sent in 2nd and 3rd association groups (S1 associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Command frames sent in 4th and 5th association groups (S2associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "The function of 3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "3-way switch function for S2 disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "3-way switch function for S2 enabled",
					},
				},
			},
			&parameter{
				ID: 27,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all groups (II-V) sent as non-secure",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group sent as secure",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "4th group sent as secure",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "5th group sent as secure",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "all groups (II-V) sent as secure",
					},
				},
			},
			&parameter{
				ID: 28,
				Name: "Scene activation functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scene functionality deactivated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene functionality activated",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Switch functionality of S1 and S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "standard mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "S1 operates as S2, S2 operates as S1",
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Load control mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "forced leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "forced trailing edge control",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control mode selected automatically (based on auto-calibration)",
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Load control mode recognized during auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "trailing edge control",
					},
				},
			},
			&parameter{
				ID: 32,
				Name: "On/Off mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "on/off mode disabled (dimming is possible)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "on/off mode enabled (dimming is not possible)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "mode selected automatically",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "Dimmability of the load (read only)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Load recognized as dimmable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Load recognized as non-dimmable",
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Soft-Start functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "no soft-start",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "short soft-start (0,1s)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "long soft-start (0,5s)",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Auto-calibration after power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No auto-calibration of the load after power on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Auto-calibration performed after first power on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Auto-calibration performed after each power on",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Auto-calibration performed after first power on or after each LOAD ERR",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Auto-calibration performed after each power on or after each LOAD ERRO",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Behaviour of the Dimmer after overload or surge",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "device permanently disabled until re-enabling by comand or push-button",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "three attempts to turn on the load",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Power limit - OVERLOAD",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 250,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "General Purpose Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Water Flooding Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Smoke, CO or CO2 Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Temperature Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Time of alarm state",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "OVERLOAD alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "LOAD ERROR alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 47,
				Name: "OVERCURRENT alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 48,
				Name: "SURGE alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 49,
				Name: "OVERHEAT (critical temperature) and VOLTAGE DROP (low voltage) alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Active power reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 52,
				Name: "Periodic active power and energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 53,
				Name: "Energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 54,
				Name: "Self-measurement",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Self-measurement inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Self-measurement active",
					},
				},
			},
			&parameter{
				ID: 58,
				Name: "Method of calculating the active power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "measurement based on the standard algorithm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "approximation based on the calibration data",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "approximation based on the control angle",
					},
				},
			},
			&parameter{
				ID: 59,
				Name: "Approximated power at the maximum brightness level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 500,
					},
				},
			},
		},
	}
}
func New010F_0102_4000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGD212",
		Description: "Dimmer 2",

		ManufacturerID: "010F",
		ProductType: "0102",
		ProductID: "4000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				NonSecure: true,
				Version: "4",
			},
			&CommandClass{
				ID: 0x32,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Minimum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 98,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Maximum brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Incandescence level of dimmable compact fluorescent lamps",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Incandescence time of dimmable compact fluorescent lamps",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Dimming Step at Automatic Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Dimming Step Timing at Automatic Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Dimming Step at Manual Control",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Dimming Step Timing at Manual Control",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Saving state before power failure",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "State NOT saved at power failure, all outputs are set to OFF upon powe",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "State saved at power failure, all outputs are set to previous state up",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Timer functionality (auto - off). 0 disables the function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "ALL ON/ALL OFF function",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON not active, ALL OFF not active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON not active, ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active, ALL OFF not active",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "ALL ON active, ALL OFF active",
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Force auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "readout",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "force auto-calibration of the load without Fibaro Bypass",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "force auto-calibration of the load with Fibaro Bypass",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Auto-calibration status (read-only parameter)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "calibration procedure not performed or Dimmer operates on manual setti",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Dimmer operates on auto-calibration settings",
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Burnt out bulb detection",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Time delay of a burnt out bulb (parameter 15) or overload (parameter 39) detection",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Forced switch on brightness level",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 99,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Roller blind switch (UP / DOWN)",
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Value sent to associated devices on single click",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "0xFF value is sent, which will set associated devices to their last sa",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "current Dimmer 2 state is sent, which will synchronize brightness leve",
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Change [On-Off] bi-stable keys",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Device changes status on key status change",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Device status depends on key status; ON when the key is ON, OFF when t",
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Double click option",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disable double click",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enable double click",
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Command frames sent in 2nd and 3rd association groups (S1 associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Command frames sent in 4th and 5th association groups (S2associations)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all actions send to association groups",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "do not send when switching ON (single click)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "do not send when switching OFF (single click)",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "do not send when changing dimming level (holding and releasing)",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "do not send on double click",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "send 0xFF value on double click",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "The function of 3-way switch",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "3-way switch function for S2 disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "3-way switch function for S2 enabled",
					},
				},
			},
			&parameter{
				ID: 27,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "all groups (II-V) sent as non-secure",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group sent as secure",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "4th group sent as secure",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "5th group sent as secure",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "all groups (II-V) sent as secure",
					},
				},
			},
			&parameter{
				ID: 28,
				Name: "Scene activation functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Scene functionality deactivated",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Scene functionality activated",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Switch functionality of S1 and S2",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "standard mode",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "S1 operates as S2, S2 operates as S1",
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Load control mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "forced leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "forced trailing edge control",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "control mode selected automatically (based on auto-calibration)",
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Load control mode recognized during auto-calibration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "leading edge control",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "trailing edge control",
					},
				},
			},
			&parameter{
				ID: 32,
				Name: "On/Off mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "on/off mode disabled (dimming is possible)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "on/off mode enabled (dimming is not possible)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "mode selected automatically",
					},
				},
			},
			&parameter{
				ID: 33,
				Name: "Dimmability of the load (read only)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Load recognized as dimmable",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Load recognized as non-dimmable",
					},
				},
			},
			&parameter{
				ID: 34,
				Name: "Soft-Start functionality",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "no soft-start",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "short soft-start (0,1s)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "long soft-start (0,5s)",
					},
				},
			},
			&parameter{
				ID: 35,
				Name: "Auto-calibration after power on",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No auto-calibration of the load after power on",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Auto-calibration performed after first power on",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Auto-calibration performed after each power on",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Auto-calibration performed after first power on or after each LOAD ERR",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Auto-calibration performed after each power on or after each LOAD ERRO",
					},
				},
			},
			&parameter{
				ID: 37,
				Name: "Behaviour of the Dimmer after overload or surge",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "device permanently disabled until re-enabling by comand or push-button",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "three attempts to turn on the load",
					},
				},
			},
			&parameter{
				ID: 39,
				Name: "Power limit - OVERLOAD",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 250,
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "General Purpose Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 41,
				Name: "Water Flooding Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Smoke, CO or CO2 Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 43,
				Name: "Temperature Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turn on the load",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Turn off the load",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Load blinking",
					},
				},
			},
			&parameter{
				ID: 44,
				Name: "Time of alarm state",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 45,
				Name: "OVERLOAD alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 46,
				Name: "LOAD ERROR alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 47,
				Name: "OVERCURRENT alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 48,
				Name: "SURGE alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 49,
				Name: "OVERHEAT (critical temperature) and VOLTAGE DROP (low voltage) alarm report",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "No reaction",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send an alarm frame",
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Active power reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 52,
				Name: "Periodic active power and energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 53,
				Name: "Energy reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 54,
				Name: "Self-measurement",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Self-measurement inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Self-measurement active",
					},
				},
			},
			&parameter{
				ID: 58,
				Name: "Method of calculating the active power",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "measurement based on the standard algorithm",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "approximation based on the calibration data",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "approximation based on the control angle",
					},
				},
			},
			&parameter{
				ID: 59,
				Name: "Approximated power at the maximum brightness level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 500,
					},
				},
			},
		},
	}
}
func New010F_0B00_1001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGFS101",
		Description: "Flood Sensor",

		ManufacturerID: "010F",
		ProductType: "0B00",
		ProductID: "1001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Acoustic and visual signals On / Off in case of flooding",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Acoustic and visual alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Acoustic alarm inactive, visual alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Acoustic alarm active, visual alarm inactive",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Acoustic and visual alarms active",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of alarm frame sent to 1-st association group (FLOOD)",
				Size: 1,
				Values: []Value{
					Value{
						From: -127,
						To: -127,
						Desc: "BASIC_SET",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM WATER",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Alarm cancelling or turning a device off (Basic) command frame deactivation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Alarm (flooding) cancellation inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Alarm (flooding) cancellation active",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Interval between successive readings of temperature sensors",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Insensitiveness to temperature changes.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Alarm BROADCAST",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Broadcast inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Flood broadcast mode active, Tamper broadcast mode inactive",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Tamper broadcast mode inactive, Tamper broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Flood and Tamper broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Low temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 51,
				Name: "High temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 61,
				Name: "Low temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "High temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 63,
				Name: "Managing a LED indicator under standard operation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED indicator doesn’t indicate the temperature",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicator indicates the temperature (blink) every Temperature Meas",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "LED indicator indicates the temperature continuously, only in constant",
					},
				},
			},
			&parameter{
				ID: 73,
				Name: "Temperature measurement compensation",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 74,
				Name: "Alarm frame sent to 2-nd Association Group activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Button tamper alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Movement tamper alarm active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Button and movement tampers alarm active",
					},
				},
			},
			&parameter{
				ID: 75,
				Name: "Visual and audible alarms duration",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 76,
				Name: "Visual and audible alarms duration",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 77,
				Name: "Flood sensor functionality turned off",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Default flood sensor operation (flood detection, reactions)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Built in flood sensor TURNED OFF (doesn’t change its state in the main",
					},
				},
			},
		},
	}
}
func New010F_0B00_2001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGFS101",
		Description: "Flood Sensor",

		ManufacturerID: "010F",
		ProductType: "0B00",
		ProductID: "2001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Acoustic and visual signals On / Off in case of flooding",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Acoustic and visual alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Acoustic alarm inactive, visual alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Acoustic alarm active, visual alarm inactive",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Acoustic and visual alarms active",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of alarm frame sent to 1-st association group (FLOOD)",
				Size: 1,
				Values: []Value{
					Value{
						From: -127,
						To: -127,
						Desc: "BASIC_SET",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM WATER",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Alarm cancelling or turning a device off (Basic) command frame deactivation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Alarm (flooding) cancellation inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Alarm (flooding) cancellation active",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Interval between successive readings of temperature sensors",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Insensitiveness to temperature changes.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Alarm BROADCAST",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Broadcast inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Flood broadcast mode active, Tamper broadcast mode inactive",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Tamper broadcast mode inactive, Tamper broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Flood and Tamper broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Low temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 51,
				Name: "High temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 61,
				Name: "Low temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "High temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 63,
				Name: "Managing a LED indicator under standard operation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED indicator doesn’t indicate the temperature",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicator indicates the temperature (blink) every Temperature Meas",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "LED indicator indicates the temperature continuously, only in constant",
					},
				},
			},
			&parameter{
				ID: 73,
				Name: "Temperature measurement compensation",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 74,
				Name: "Alarm frame sent to 2-nd Association Group activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Button tamper alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Movement tamper alarm active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Button and movement tampers alarm active",
					},
				},
			},
			&parameter{
				ID: 75,
				Name: "Visual and audible alarms duration",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 76,
				Name: "Visual and audible alarms duration",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 77,
				Name: "Flood sensor functionality turned off",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Default flood sensor operation (flood detection, reactions)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Built in flood sensor TURNED OFF (doesn’t change its state in the main",
					},
				},
			},
		},
	}
}
func New010F_0B00_3001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGFS101",
		Description: "Flood Sensor",

		ManufacturerID: "010F",
		ProductType: "0B00",
		ProductID: "3001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x60,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Acoustic and visual signals On / Off in case of flooding",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Acoustic and visual alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Acoustic alarm inactive, visual alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Acoustic alarm active, visual alarm inactive",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Acoustic and visual alarms active",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of alarm frame sent to 1-st association group (FLOOD)",
				Size: 1,
				Values: []Value{
					Value{
						From: -127,
						To: -127,
						Desc: "BASIC_SET",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM WATER",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Alarm cancelling or turning a device off (Basic) command frame deactivation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Alarm (flooding) cancellation inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Alarm (flooding) cancellation active",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Interval between successive readings of temperature sensors",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Insensitiveness to temperature changes.",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Alarm BROADCAST",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Broadcast inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Flood broadcast mode active, Tamper broadcast mode inactive",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Tamper broadcast mode inactive, Tamper broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Flood and Tamper broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Low temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 51,
				Name: "High temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 61,
				Name: "Low temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "High temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 63,
				Name: "Managing a LED indicator under standard operation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED indicator doesn’t indicate the temperature",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicator indicates the temperature (blink) every Temperature Meas",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "LED indicator indicates the temperature continuously, only in constant",
					},
				},
			},
			&parameter{
				ID: 73,
				Name: "Temperature measurement compensation",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 74,
				Name: "Alarm frame sent to 2-nd Association Group activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Button tamper alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Movement tamper alarm active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Button and movement tampers alarm active",
					},
				},
			},
			&parameter{
				ID: 75,
				Name: "Visual and audible alarms duration",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 76,
				Name: "Visual and audible alarms duration",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 77,
				Name: "Flood sensor functionality turned off",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Default flood sensor operation (flood detection, reactions)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Built in flood sensor TURNED OFF (doesn’t change its state in the main",
					},
				},
			},
		},
	}
}
func New010F_0B01_1002() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGFS101",
		Description: "Flood Sensor",

		ManufacturerID: "010F",
		ProductType: "0B01",
		ProductID: "1002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Select alarm type (visual/acoustic)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Acoustic and visual alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Acoustic alarm inactive, visual alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Acoustic alarm active, visual alarm inactive",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Acoustic and visual alarms active",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Requested dimming level to 2nd assoc. group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Device off and alarm cancellations",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Alarm (flooding) cancellation inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Alarm (flooding) cancellation active",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Temperature measurement interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Temperature measurement hysteresis",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Low temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 51,
				Name: "High temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 61,
				Name: "Low temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "High temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 63,
				Name: "Visual temperature indicator",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Visual indicator does not indicate the temperature",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Indicator at every temperature measurement and wake up",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Visual indicator indicates the temperature continuously",
					},
				},
			},
			&parameter{
				ID: 73,
				Name: "Temperature measurement compensation",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 74,
				Name: "Alarm frame for movement/tamper",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Button tamper alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Movement tamper alarm active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Button and movement tampers alarm active",
					},
				},
			},
			&parameter{
				ID: 75,
				Name: "Alarms signalization duration",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 76,
				Name: "Alarm signalization reactivation period",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 77,
				Name: "Flood sensor functionality turned off",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Default flood sensor operation (flood detection, reactions)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Built-in flood sensor TURNED OFF",
					},
				},
			},
			&parameter{
				ID: 78,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None of the groups are sent as secure",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group ”Control” sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group ”Alarm” sent as secure",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "4th group \"Tamper” sent as secure",
					},
				},
			},
		},
	}
}
func New010F_0B01_2002() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGFS101",
		Description: "Flood Sensor",

		ManufacturerID: "010F",
		ProductType: "0B01",
		ProductID: "2002",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 3600,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Select alarm type (visual/acoustic)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Acoustic and visual alarms inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Acoustic alarm inactive, visual alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Acoustic alarm active, visual alarm inactive",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Acoustic and visual alarms active",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Requested dimming level to 2nd assoc. group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Device off and alarm cancellations",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Alarm (flooding) cancellation inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Alarm (flooding) cancellation active",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Temperature measurement interval",
				Size: 4,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Temperature measurement hysteresis",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Low temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 51,
				Name: "High temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 61,
				Name: "Low temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "High temperature alarm indicator colour",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 16777215,
					},
				},
			},
			&parameter{
				ID: 63,
				Name: "Visual temperature indicator",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Visual indicator does not indicate the temperature",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Indicator at every temperature measurement and wake up",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Visual indicator indicates the temperature continuously",
					},
				},
			},
			&parameter{
				ID: 73,
				Name: "Temperature measurement compensation",
				Size: 2,
				Values: []Value{
					Value{
						From: -10000,
						To: 10000,
					},
				},
			},
			&parameter{
				ID: 74,
				Name: "Alarm frame for movement/tamper",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Button tamper alarm active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Movement tamper alarm active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Button and movement tampers alarm active",
					},
				},
			},
			&parameter{
				ID: 75,
				Name: "Alarms signalization duration",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 76,
				Name: "Alarm signalization reactivation period",
				Size: 4,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 77,
				Name: "Flood sensor functionality turned off",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Default flood sensor operation (flood detection, reactions)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Built-in flood sensor TURNED OFF",
					},
				},
			},
			&parameter{
				ID: 78,
				Name: "Associations in Z-Wave network security mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "None of the groups are sent as secure",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2nd group ”Control” sent as secure",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3rd group ”Alarm” sent as secure",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "4th group \"Tamper” sent as secure",
					},
				},
			},
		},
	}
}
func New010F_0700_1000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGK101",
		Description: "Door Opening Sensor",

		ManufacturerID: "010F",
		ProductType: "0700",
		ProductID: "1000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "IN1 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Status change signalled by LED",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED turned Off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED turned On",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Type of input no. 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Normally Open",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Normally Closed",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Monostable",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Bistable",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of transmitted control frame for association",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Enable/Disable transmission of frame cancelling alarm 	",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Groups 1 and 2 sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 sent, Group 2 not sent.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 2 sent, Group 1 not sent.",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Sensitivity to temperature changes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Transmitting the alarm or control frame broadcast",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "IN and TMP Broadcast mode inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "IN broadcast mode active, TMP broadcast mode inactive	",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "IN broadcast mode inactive, TMP broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "IN and TMP broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Scene activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
		},
	}
}
func New010F_0700_2000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGK101",
		Description: "Door Opening Sensor",

		ManufacturerID: "010F",
		ProductType: "0700",
		ProductID: "2000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "IN1 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Status change signalled by LED",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED turned Off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED turned On",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Type of input no. 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Normally Open",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Normally Closed",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Monostable",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Bistable",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of transmitted control frame for association",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Enable/Disable transmission of frame cancelling alarm 	",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Groups 1 and 2 sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 sent, Group 2 not sent.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 2 sent, Group 1 not sent.",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Sensitivity to temperature changes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Transmitting the alarm or control frame broadcast",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "IN and TMP Broadcast mode inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "IN broadcast mode active, TMP broadcast mode inactive	",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "IN broadcast mode inactive, TMP broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "IN and TMP broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Scene activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
		},
	}
}
func New010F_0700_3000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGK101",
		Description: "Door Opening Sensor",

		ManufacturerID: "010F",
		ProductType: "0700",
		ProductID: "3000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "IN1 Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Status change signalled by LED",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED turned Off",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED turned On",
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Type of input no. 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Normally Open",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Normally Closed",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Monostable",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Bistable",
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Type of transmitted control frame for association",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "ALARM GENERIC",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALARM SMOKE",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALARM CO",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "ALARM CO2",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "ALARM HEAT",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "ALARM WATER",
					},
					Value{
						From: 255,
						To: 255,
						Desc: "BASIC_SET",
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Forced Level of Dimming group 1",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Enable/Disable transmission of frame cancelling alarm 	",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Groups 1 and 2 sent",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Group 1 sent, Group 2 not sent.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Group 2 sent, Group 1 not sent.",
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Sensitivity to temperature changes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Transmitting the alarm or control frame broadcast",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "IN and TMP Broadcast mode inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "IN broadcast mode active, TMP broadcast mode inactive	",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "IN broadcast mode inactive, TMP broadcast mode active",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "IN and TMP broadcast mode active",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Scene activation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Enabled",
					},
				},
			},
		},
	}
}
func New010F_0701_1001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGK101",
		Description: "Door Opening Sensor",

		ManufacturerID: "010F",
		ProductType: "0701",
		ProductID: "1001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Operations mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Door/Window or alarm status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Visual LED indications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 6,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Range test after double click",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "2nd association group triggers",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 2,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Commands sent to 2nd association group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 2,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Value of ON command frame sent to 2nd association group",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Value of OFF command frame sent to 2nd association group",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Time delay of ON command frame",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Time delay of OFF command frame",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Type of sent alarm frames",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 5,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Delay of tamper alarm cancellation",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Reporting tamper alarm cancellation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Interval of temperature measurements",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 51,
				Name: "Temperature reports threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 300,
					},
				},
			},
			&parameter{
				ID: 52,
				Name: "Interval of temperature reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 53,
				Name: "Temperature offset",
				Size: 4,
				Values: []Value{
					Value{
						From: -1000,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 54,
				Name: "Temperature alarm reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 55,
				Name: "High temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 56,
				Name: "Low temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -300,
						To: 700,
					},
				},
			},
			&parameter{
				ID: 70,
				Name: "Scene activation functionality",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 2048,
					},
				},
			},
			&parameter{
				ID: 71,
				Name: "Alarm broadcast",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 72,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
		},
	}
}
func New010F_0701_3001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGK101",
		Description: "Door Opening Sensor",

		ManufacturerID: "010F",
		ProductType: "0701",
		ProductID: "3001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x2B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Operations mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Door/Window or alarm status",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Visual LED indications",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 6,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Range test after double click",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "2nd association group triggers",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 2,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Commands sent to 2nd association group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 2,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Value of ON command frame sent to 2nd association group",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Value of OFF command frame sent to 2nd association group",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Time delay of ON command frame",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Time delay of OFF command frame",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Type of sent alarm frames",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 5,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Delay of tamper alarm cancellation",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 31,
				Name: "Reporting tamper alarm cancellation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 1,
					},
				},
			},
			&parameter{
				ID: 50,
				Name: "Interval of temperature measurements",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 51,
				Name: "Temperature reports threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 300,
					},
				},
			},
			&parameter{
				ID: 52,
				Name: "Interval of temperature reports",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32400,
					},
				},
			},
			&parameter{
				ID: 53,
				Name: "Temperature offset",
				Size: 4,
				Values: []Value{
					Value{
						From: -1000,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 54,
				Name: "Temperature alarm reports",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 55,
				Name: "High temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1000,
					},
				},
			},
			&parameter{
				ID: 56,
				Name: "Low temperature alarm threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: -300,
						To: 700,
					},
				},
			},
			&parameter{
				ID: 70,
				Name: "Scene activation functionality",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 2048,
					},
				},
			},
			&parameter{
				ID: 71,
				Name: "Alarm broadcast",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 4,
					},
				},
			},
			&parameter{
				ID: 72,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
		},
	}
}
func New010F_1001_1000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGKF601",
		Description: "Keyfob",

		ManufacturerID: "010F",
		ProductType: "1001",
		ProductID: "1000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x5E,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x75,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x7A,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x8E,
				NonSecure: true,
				Version: "3",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Lock Mode - unlocking sequence",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 28086,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Lock Mode - time to lock and locking button",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 1791,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "First scene sequence",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 28086,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Second scene sequence",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 28086,
					},
				},
			},
			&parameter{
				ID: 5,
				Name: "Third scene sequence",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 28086,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Fourth scene sequence",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 28086,
					},
				},
			},
			&parameter{
				ID: 7,
				Name: "Fifth scene sequence",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 28086,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Sixth scene sequence",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 28086,
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Sequences - timeout",
				Size: 1,
				Values: []Value{
					Value{
						From: 5,
						To: 30,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Single button associations - operating mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "single click switches state to opposite",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "single click => opposite, double click => max",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "single click turns on, double click turns off",
					},
				},
			},
			&parameter{
				ID: 11,
				Name: "Value sent to □ association group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "Value sent to ○ association group	",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Value sent to x association group	",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Value sent to Δ association group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Value sent to - association group	",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "Value sent to + association group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 17,
				Name: "Paired buttons association for □ and ○",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "paired buttons association inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "paired buttons association active",
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Paired buttons association for x and Δ",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "paired buttons association inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "paired buttons association active",
					},
				},
			},
			&parameter{
				ID: 19,
				Name: "Paired buttons association for - and +",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "paired buttons association inactive",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "paired buttons association active",
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Scene activation for □ button",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Scene activation for ○ button",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Scene activation for x button	",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Scene activation for Δ button	",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Scene activation for - button	",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "Scene activation for + button	",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Associations in Z-Wave network security mode",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 4095,
					},
				},
			},
		},
	}
}
func New010F_0800_1001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGMS001",
		Description: "Motion Sensor",

		ManufacturerID: "010F",
		ProductType: "0800",
		ProductID: "1001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Motion Sensor's Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 8,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Motion Sensor's Blind Time (Insensitive)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensor’s “PULSE COUNTER”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "PIR Sensor’s “WINDOW TIME”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Motion Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Sensor Operating Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "PIR sensor always active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "PIR sensor active during the day only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "PIR sensor active during the night only",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Night / Day",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "BASIC command class configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "BASIC ON and BASIC OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only the BASIC ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only the BASIC OFF",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "BASIC ON command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "BASIC OFF command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Tamper Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 122,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Tamper - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Tamper Operating Modes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is reported  / Cancellation is not reported.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm is reported / Cancellation is reported.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Same as Value 0 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Same as Value 1 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "The maximum level of vibrations recorded in the time period.",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "Tamper Alarm Broadcast Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is not sent in broadcast mode.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm sent in broadcast mode.",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Illumination Report Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Illumination Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 60,
				Name: "Temperature Report Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "Interval of Temperature Measuring",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "Temperature Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 66,
				Name: "Temperature Offset",
				Size: 2,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "LED Signalling Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED inactive.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flashlight mode - LED glows in white for 10 seconds.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "White.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Red.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Green.",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Blue.",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Yellow.",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Cyan.",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Magenta.",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Flashlight mode",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "White.",
					},
					Value{
						From: 13,
						To: 13,
						Desc: "Red.",
					},
					Value{
						From: 14,
						To: 14,
						Desc: "Green.",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "Blue.",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Yellow.",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Cyan.",
					},
					Value{
						From: 18,
						To: 18,
						Desc: "Magenta.",
					},
					Value{
						From: 19,
						To: 19,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 20,
						To: 20,
						Desc: "White.",
					},
					Value{
						From: 21,
						To: 21,
						Desc: "Red.",
					},
					Value{
						From: 22,
						To: 22,
						Desc: "Green.",
					},
					Value{
						From: 23,
						To: 23,
						Desc: "Blue.",
					},
					Value{
						From: 24,
						To: 24,
						Desc: "Yellow.",
					},
					Value{
						From: 25,
						To: 25,
						Desc: "Cyan.",
					},
					Value{
						From: 26,
						To: 26,
						Desc: "Magenta.",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Brightness",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "Min Temperature Resulting in Blue LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 87,
				Name: "Max Temperature Resulting in Red LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 89,
				Name: "LED indicating Tamper Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED does not indicate tamper alarm.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicates tamper alarm.",
					},
				},
			},
		},
	}
}
func New010F_0800_2001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGMS001",
		Description: "Motion Sensor",

		ManufacturerID: "010F",
		ProductType: "0800",
		ProductID: "2001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Motion Sensor's Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 8,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Motion Sensor's Blind Time (Insensitive)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensor’s “PULSE COUNTER”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "PIR Sensor’s “WINDOW TIME”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Motion Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Sensor Operating Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "PIR sensor always active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "PIR sensor active during the day only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "PIR sensor active during the night only",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Night / Day",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "BASIC command class configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "BASIC ON and BASIC OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only the BASIC ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only the BASIC OFF",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "BASIC ON command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "BASIC OFF command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Tamper Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 122,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Tamper - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Tamper Operating Modes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is reported  / Cancellation is not reported.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm is reported / Cancellation is reported.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Same as Value 0 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Same as Value 1 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "The maximum level of vibrations recorded in the time period.",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "Tamper Alarm Broadcast Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is not sent in broadcast mode.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm sent in broadcast mode.",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Illumination Report Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Illumination Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 60,
				Name: "Temperature Report Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "Interval of Temperature Measuring",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "Temperature Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 66,
				Name: "Temperature Offset",
				Size: 2,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "LED Signalling Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED inactive.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flashlight mode - LED glows in white for 10 seconds.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "White.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Red.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Green.",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Blue.",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Yellow.",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Cyan.",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Magenta.",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Flashlight mode",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "White.",
					},
					Value{
						From: 13,
						To: 13,
						Desc: "Red.",
					},
					Value{
						From: 14,
						To: 14,
						Desc: "Green.",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "Blue.",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Yellow.",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Cyan.",
					},
					Value{
						From: 18,
						To: 18,
						Desc: "Magenta.",
					},
					Value{
						From: 19,
						To: 19,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 20,
						To: 20,
						Desc: "White.",
					},
					Value{
						From: 21,
						To: 21,
						Desc: "Red.",
					},
					Value{
						From: 22,
						To: 22,
						Desc: "Green.",
					},
					Value{
						From: 23,
						To: 23,
						Desc: "Blue.",
					},
					Value{
						From: 24,
						To: 24,
						Desc: "Yellow.",
					},
					Value{
						From: 25,
						To: 25,
						Desc: "Cyan.",
					},
					Value{
						From: 26,
						To: 26,
						Desc: "Magenta.",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Brightness",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "Min Temperature Resulting in Blue LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 87,
				Name: "Max Temperature Resulting in Red LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 89,
				Name: "LED indicating Tamper Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED does not indicate tamper alarm.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicates tamper alarm.",
					},
				},
			},
		},
	}
}
func New010F_0800_3001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGMS001",
		Description: "Motion Sensor",

		ManufacturerID: "010F",
		ProductType: "0800",
		ProductID: "3001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Motion Sensor's Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 8,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Motion Sensor's Blind Time (Insensitive)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensor’s “PULSE COUNTER”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "PIR Sensor’s “WINDOW TIME”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Motion Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Sensor Operating Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "PIR sensor always active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "PIR sensor active during the day only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "PIR sensor active during the night only",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Night / Day",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "BASIC command class configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "BASIC ON and BASIC OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only the BASIC ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only the BASIC OFF",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "BASIC ON command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "BASIC OFF command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Tamper Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 122,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Tamper - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Tamper Operating Modes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is reported  / Cancellation is not reported.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm is reported / Cancellation is reported.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Same as Value 0 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Same as Value 1 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "The maximum level of vibrations recorded in the time period.",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "Tamper Alarm Broadcast Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is not sent in broadcast mode.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm sent in broadcast mode.",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Illumination Report Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Illumination Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 60,
				Name: "Temperature Report Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "Interval of Temperature Measuring",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "Temperature Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 66,
				Name: "Temperature Offset",
				Size: 2,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "LED Signalling Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED inactive.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flashlight mode - LED glows in white for 10 seconds.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "White.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Red.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Green.",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Blue.",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Yellow.",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Cyan.",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Magenta.",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Flashlight mode",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "White.",
					},
					Value{
						From: 13,
						To: 13,
						Desc: "Red.",
					},
					Value{
						From: 14,
						To: 14,
						Desc: "Green.",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "Blue.",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Yellow.",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Cyan.",
					},
					Value{
						From: 18,
						To: 18,
						Desc: "Magenta.",
					},
					Value{
						From: 19,
						To: 19,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 20,
						To: 20,
						Desc: "White.",
					},
					Value{
						From: 21,
						To: 21,
						Desc: "Red.",
					},
					Value{
						From: 22,
						To: 22,
						Desc: "Green.",
					},
					Value{
						From: 23,
						To: 23,
						Desc: "Blue.",
					},
					Value{
						From: 24,
						To: 24,
						Desc: "Yellow.",
					},
					Value{
						From: 25,
						To: 25,
						Desc: "Cyan.",
					},
					Value{
						From: 26,
						To: 26,
						Desc: "Magenta.",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Brightness",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "Min Temperature Resulting in Blue LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 87,
				Name: "Max Temperature Resulting in Red LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 89,
				Name: "LED indicating Tamper Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED does not indicate tamper alarm.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicates tamper alarm.",
					},
				},
			},
		},
	}
}
func New010F_0801_1001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGMS001",
		Description: "Motion Sensor",

		ManufacturerID: "010F",
		ProductType: "0801",
		ProductID: "1001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Motion Sensor's Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 8,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Motion Sensor's Blind Time (Insensitive)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensor’s “PULSE COUNTER”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "PIR Sensor’s “WINDOW TIME”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Motion Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Sensor Operating Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "PIR sensor always active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "PIR sensor active during the day only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "PIR sensor active during the night only",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Night / Day",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "BASIC command class configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "BASIC ON and BASIC OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only the BASIC ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only the BASIC OFF",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "BASIC ON command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "BASIC OFF command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Tamper Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 122,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Tamper - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Tamper Operating Modes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is reported  / Cancellation is not reported.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm is reported / Cancellation is reported.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Same as Value 0 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Same as Value 1 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "The maximum level of vibrations recorded in the time period.",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "Tamper Alarm Broadcast Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is not sent in broadcast mode.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm sent in broadcast mode.",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Illumination Report Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Illumination Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 60,
				Name: "Temperature Report Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "Interval of Temperature Measuring",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "Temperature Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 66,
				Name: "Temperature Offset",
				Size: 2,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "LED Signalling Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED inactive.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flashlight mode - LED glows in white for 10 seconds.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "White.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Red.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Green.",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Blue.",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Yellow.",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Cyan.",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Magenta.",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Flashlight mode",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "White.",
					},
					Value{
						From: 13,
						To: 13,
						Desc: "Red.",
					},
					Value{
						From: 14,
						To: 14,
						Desc: "Green.",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "Blue.",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Yellow.",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Cyan.",
					},
					Value{
						From: 18,
						To: 18,
						Desc: "Magenta.",
					},
					Value{
						From: 19,
						To: 19,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 20,
						To: 20,
						Desc: "White.",
					},
					Value{
						From: 21,
						To: 21,
						Desc: "Red.",
					},
					Value{
						From: 22,
						To: 22,
						Desc: "Green.",
					},
					Value{
						From: 23,
						To: 23,
						Desc: "Blue.",
					},
					Value{
						From: 24,
						To: 24,
						Desc: "Yellow.",
					},
					Value{
						From: 25,
						To: 25,
						Desc: "Cyan.",
					},
					Value{
						From: 26,
						To: 26,
						Desc: "Magenta.",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Brightness",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "Min Temperature Resulting in Blue LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 87,
				Name: "Max Temperature Resulting in Red LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 89,
				Name: "LED indicating Tamper Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED does not indicate tamper alarm.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicates tamper alarm.",
					},
				},
			},
		},
	}
}
func New010F_0801_2001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGMS001",
		Description: "Motion Sensor",

		ManufacturerID: "010F",
		ProductType: "0801",
		ProductID: "2001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "5",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x8F,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Motion Sensor's Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 8,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Motion Sensor's Blind Time (Insensitive)",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "PIR Sensor’s “PULSE COUNTER”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "PIR Sensor’s “WINDOW TIME”",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Motion Alarm Cancellation Delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "PIR Sensor Operating Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "PIR sensor always active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "PIR sensor active during the day only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "PIR sensor active during the night only",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Night / Day",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "BASIC command class configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "BASIC ON and BASIC OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only the BASIC ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only the BASIC OFF",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "BASIC ON command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "BASIC OFF command frame value",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Tamper Sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 122,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Tamper - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Tamper Operating Modes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is reported  / Cancellation is not reported.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm is reported / Cancellation is reported.",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Same as Value 0 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "Same as Value 1 / with unsupported Fibar Command Class.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "The maximum level of vibrations recorded in the time period.",
					},
				},
			},
			&parameter{
				ID: 26,
				Name: "Tamper Alarm Broadcast Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Tamper alarm is not sent in broadcast mode.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper alarm sent in broadcast mode.",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Illumination Report Threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Illumination Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 60,
				Name: "Temperature Report Threshold",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "Interval of Temperature Measuring",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "Temperature Reports Interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 66,
				Name: "Temperature Offset",
				Size: 2,
				Values: []Value{
					Value{
						From: -100,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "LED Signalling Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED inactive.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flashlight mode - LED glows in white for 10 seconds.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "White.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Red.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Green.",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Blue.",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Yellow.",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Cyan.",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Magenta.",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Flashlight mode",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "White.",
					},
					Value{
						From: 13,
						To: 13,
						Desc: "Red.",
					},
					Value{
						From: 14,
						To: 14,
						Desc: "Green.",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "Blue.",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Yellow.",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Cyan.",
					},
					Value{
						From: 18,
						To: 18,
						Desc: "Magenta.",
					},
					Value{
						From: 19,
						To: 19,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 20,
						To: 20,
						Desc: "White.",
					},
					Value{
						From: 21,
						To: 21,
						Desc: "Red.",
					},
					Value{
						From: 22,
						To: 22,
						Desc: "Green.",
					},
					Value{
						From: 23,
						To: 23,
						Desc: "Blue.",
					},
					Value{
						From: 24,
						To: 24,
						Desc: "Yellow.",
					},
					Value{
						From: 25,
						To: 25,
						Desc: "Cyan.",
					},
					Value{
						From: 26,
						To: 26,
						Desc: "Magenta.",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED Brightness",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "Ambient Illumination Level",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "Min Temperature Resulting in Blue LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 87,
				Name: "Max Temperature Resulting in Red LED Illumination",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 89,
				Name: "LED indicating Tamper Alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED does not indicate tamper alarm.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicates tamper alarm.",
					},
				},
			},
		},
	}
}
func New010F_0801_3001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGMS001",
		Description: "Motion Sensor",

		ManufacturerID: "010F",
		ProductType: "0801",
		ProductID: "3001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "7",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Motion detection - sensitivity",
				Size: 2,
				Values: []Value{
					Value{
						From: 8,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Motion detection - blind time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Motion detection - pulse counter",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "1 pulse",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 pulses",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 pulses",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "4 pulses",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion detection - window time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "4 seconds",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "8 seconds",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "12 seconds",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "16 seconds",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Motion detection - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Motion detection - operating mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "PIR sensor always active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "PIR sensor active during the day only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "PIR sensor active during the night only",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Motion detection - night/day",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "BASIC command class configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "BASIC ON and BASIC OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only the BASIC ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only the BASIC OFF",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "BASIC ON command frame value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "BASIC OFF command frame value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Tamper - sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 121,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Tamper - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Tamper - operating modes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Only tamper (default)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper and earthquake detector",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Tamper and orientation in space",
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Tamper - alarm cancellation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not send tamper cancellation report",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send tamper cancellation report",
					},
				},
			},
			&parameter{
				ID: 28,
				Name: "Tamper - broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reported to association groups.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reported in broadcast mode to 3rd association group.",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Tamper - backward compatible broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reported to association groups.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reported in broadcast mode to 5th association group.",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Illuminance report - threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Illuminance report - interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 60,
				Name: "Temperature report - threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "Temperature measuring - interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "Temperature report - interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 66,
				Name: "Temperature offset",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "LED - signalling mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED inactive.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flashlight mode - LED glows in white for 10 seconds.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "White.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Red.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Green.",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Blue.",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Yellow.",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Cyan.",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Magenta.",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Flashlight mode",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "White.",
					},
					Value{
						From: 13,
						To: 13,
						Desc: "Red.",
					},
					Value{
						From: 14,
						To: 14,
						Desc: "Green.",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "Blue.",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Yellow.",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Cyan.",
					},
					Value{
						From: 18,
						To: 18,
						Desc: "Magenta.",
					},
					Value{
						From: 19,
						To: 19,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 20,
						To: 20,
						Desc: "White.",
					},
					Value{
						From: 21,
						To: 21,
						Desc: "Red.",
					},
					Value{
						From: 22,
						To: 22,
						Desc: "Green.",
					},
					Value{
						From: 23,
						To: 23,
						Desc: "Blue.",
					},
					Value{
						From: 24,
						To: 24,
						Desc: "Yellow.",
					},
					Value{
						From: 25,
						To: 25,
						Desc: "Cyan.",
					},
					Value{
						From: 26,
						To: 26,
						Desc: "Magenta.",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED - brightness",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "LED - illuminance for low indicator brightness",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "LED - illuminance for high indicator brightness",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "LED - temperature for blue colour",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 87,
				Name: "LED - temperature for red colour",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 89,
				Name: "LED - tamper alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED does not indicate tamper alarm.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicates tamper alarm.",
					},
				},
			},
		},
	}
}
func New010F_8800_3001() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGMS001",
		Description: "Motion Sensor",

		ManufacturerID: "010F",
		ProductType: "8800",
		ProductID: "3001",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x22,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x30,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x31,
				InNIF: true,
				NonSecure: true,
				Version: "7",
			},
			&CommandClass{
				ID: 0x56,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x59,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5E,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x7A,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				InNIF: true,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x98,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
			&CommandClass{
				ID: 0x9C,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Motion detection - sensitivity",
				Size: 2,
				Values: []Value{
					Value{
						From: 8,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 2,
				Name: "Motion detection - blind time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Motion detection - pulse counter",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "1 pulse",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "2 pulses",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "3 pulses",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "4 pulses",
					},
				},
			},
			&parameter{
				ID: 4,
				Name: "Motion detection - window time",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "4 seconds",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "8 seconds",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "12 seconds",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "16 seconds",
					},
				},
			},
			&parameter{
				ID: 6,
				Name: "Motion detection - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 8,
				Name: "Motion detection - operating mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "PIR sensor always active",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "PIR sensor active during the day only",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "PIR sensor active during the night only",
					},
				},
			},
			&parameter{
				ID: 9,
				Name: "Motion detection - night/day",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: "BASIC command class configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "BASIC ON and BASIC OFF",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Only the BASIC ON",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Only the BASIC OFF",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "BASIC ON command frame value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 16,
				Name: "BASIC OFF command frame value",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 18,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 15,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: "Tamper - sensitivity",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 121,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: "Tamper - alarm cancellation delay",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Tamper - operating modes",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Only tamper (default)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Tamper and earthquake detector",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Tamper and orientation in space",
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Tamper - alarm cancellation",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Do not send tamper cancellation report",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Send tamper cancellation report",
					},
				},
			},
			&parameter{
				ID: 28,
				Name: "Tamper - broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reported to association groups.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reported in broadcast mode to 3rd association group.",
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Tamper - backward compatible broadcast mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Reported to association groups.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Reported in broadcast mode to 5th association group.",
					},
				},
			},
			&parameter{
				ID: 40,
				Name: "Illuminance report - threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 42,
				Name: "Illuminance report - interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 60,
				Name: "Temperature report - threshold",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 62,
				Name: "Temperature measuring - interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 64,
				Name: "Temperature report - interval",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 66,
				Name: "Temperature offset",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 65535,
					},
				},
			},
			&parameter{
				ID: 80,
				Name: "LED - signalling mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED inactive.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Flashlight mode - LED glows in white for 10 seconds.",
					},
					Value{
						From: 3,
						To: 3,
						Desc: "White.",
					},
					Value{
						From: 4,
						To: 4,
						Desc: "Red.",
					},
					Value{
						From: 5,
						To: 5,
						Desc: "Green.",
					},
					Value{
						From: 6,
						To: 6,
						Desc: "Blue.",
					},
					Value{
						From: 7,
						To: 7,
						Desc: "Yellow.",
					},
					Value{
						From: 8,
						To: 8,
						Desc: "Cyan.",
					},
					Value{
						From: 9,
						To: 9,
						Desc: "Magenta.",
					},
					Value{
						From: 10,
						To: 10,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 11,
						To: 11,
						Desc: "Flashlight mode",
					},
					Value{
						From: 12,
						To: 12,
						Desc: "White.",
					},
					Value{
						From: 13,
						To: 13,
						Desc: "Red.",
					},
					Value{
						From: 14,
						To: 14,
						Desc: "Green.",
					},
					Value{
						From: 15,
						To: 15,
						Desc: "Blue.",
					},
					Value{
						From: 16,
						To: 16,
						Desc: "Yellow.",
					},
					Value{
						From: 17,
						To: 17,
						Desc: "Cyan.",
					},
					Value{
						From: 18,
						To: 18,
						Desc: "Magenta.",
					},
					Value{
						From: 19,
						To: 19,
						Desc: "LED colour depends on the temperature",
					},
					Value{
						From: 20,
						To: 20,
						Desc: "White.",
					},
					Value{
						From: 21,
						To: 21,
						Desc: "Red.",
					},
					Value{
						From: 22,
						To: 22,
						Desc: "Green.",
					},
					Value{
						From: 23,
						To: 23,
						Desc: "Blue.",
					},
					Value{
						From: 24,
						To: 24,
						Desc: "Yellow.",
					},
					Value{
						From: 25,
						To: 25,
						Desc: "Cyan.",
					},
					Value{
						From: 26,
						To: 26,
						Desc: "Magenta.",
					},
				},
			},
			&parameter{
				ID: 81,
				Name: "LED - brightness",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 100,
					},
				},
			},
			&parameter{
				ID: 82,
				Name: "LED - illuminance for low indicator brightness",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 83,
				Name: "LED - illuminance for high indicator brightness",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 32767,
					},
				},
			},
			&parameter{
				ID: 86,
				Name: "LED - temperature for blue colour",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 87,
				Name: "LED - temperature for red colour",
				Size: 2,
				Values: []Value{
					Value{
						From: 0,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 89,
				Name: "LED - tamper alarm",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "LED does not indicate tamper alarm.",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "LED indicates tamper alarm.",
					},
				},
			},
		},
	}
}
func New010F_0F01_1000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGPB101",
		Description: "A real push button (switch) available in several colors.",

		ManufacturerID: "010F",
		ProductType: "0F01",
		ProductID: "1000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: " Scenes sent to the controller",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 7,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: " Key Pressed 1 time – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: " Key Pressed 1 time – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: " Key Pressed 2 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Key Pressed 3 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: " Key Pressed 1 time – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Key Pressed 1 time – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: " Key Pressed 2 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Key Pressed 3 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Key Held Down – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Alarm frame triggers",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
		},
	}
}
func New010F_0F01_2000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGPB101",
		Description: "A real push button (switch) available in several colors.",

		ManufacturerID: "010F",
		ProductType: "0F01",
		ProductID: "2000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: " Scenes sent to the controller",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 7,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: " Key Pressed 1 time – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: " Key Pressed 1 time – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: " Key Pressed 2 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Key Pressed 3 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: " Key Pressed 1 time – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Key Pressed 1 time – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: " Key Pressed 2 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Key Pressed 3 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Key Held Down – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Alarm frame triggers",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
		},
	}
}
func New010F_0F01_3000() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGPB101",
		Description: "A real push button (switch) available in several colors.",

		ManufacturerID: "010F",
		ProductType: "0F01",
		ProductID: "3000",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: " Scenes sent to the controller",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 7,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: " Key Pressed 1 time – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: " Key Pressed 1 time – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: " Key Pressed 2 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Key Pressed 3 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: " Key Pressed 1 time – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Key Pressed 1 time – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: " Key Pressed 2 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Key Pressed 3 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Key Held Down – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Alarm frame triggers",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
		},
	}
}
func New010F_FB10_1014() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGPB101",
		Description: "A real push button (switch) available in several colors.",

		ManufacturerID: "010F",
		ProductType: "FB10",
		ProductID: "1014",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x56,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x5B,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x71,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x80,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x84,
				NonSecure: true,
				Version: "2",
			},
			&CommandClass{
				ID: 0x85,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				NonSecure: true,
				Version: "1",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: " Scenes sent to the controller",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
			&parameter{
				ID: 3,
				Name: "Associations in Z-Wave network Security Mode",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 7,
					},
				},
			},
			&parameter{
				ID: 10,
				Name: " Key Pressed 1 time – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 11,
				Name: " Key Pressed 1 time – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 12,
				Name: " Key Pressed 2 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 13,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Key Pressed 3 times – command sent to 2nd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 15,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 2nd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 20,
				Name: " Key Pressed 1 time – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 21,
				Name: "Key Pressed 1 time – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 22,
				Name: " Key Pressed 2 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 23,
				Name: "Key Pressed 2 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 24,
				Name: "Key Pressed 3 times – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 25,
				Name: "Key Pressed 3 times – value of SWITCH ON sent to 3rd group",
				Size: 2,
				Values: []Value{
					Value{
						From: 1,
						To: 255,
					},
				},
			},
			&parameter{
				ID: 29,
				Name: "Key Held Down – command sent to 3rd group",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 3,
					},
				},
			},
			&parameter{
				ID: 30,
				Name: "Alarm frame triggers",
				Size: 1,
				Values: []Value{
					Value{
						From: 1,
						To: 127,
					},
				},
			},
		},
	}
}
func New010F_0300_0104() *Device{
	return &Device{
		Brand: "Fibargroup",
		Product: "FGR221",
		Description: "Roller Shutter Controller",

		ManufacturerID: "010F",
		ProductType: "0300",
		ProductID: "0104",
		CommandClasses: []*CommandClass{
			&CommandClass{
				ID: 0x00,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x20,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x25,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x26,
				InNIF: true,
				NonSecure: true,
				Version: "3",
			},
			&CommandClass{
				ID: 0x27,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x50,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x70,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x72,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x73,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x85,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x86,
				InNIF: true,
				NonSecure: true,
				Version: "1",
			},
			&CommandClass{
				ID: 0x8E,
				InNIF: true,
				NonSecure: true,
				Version: "0",
			},
		},
		Parameters: []*parameter{
			&parameter{
				ID: 1,
				Name: "Enable/Disable ALL ON/OFF",
				Size: 1,
				Values: []Value{
					Value{
						From: -1,
						To: -1,
						Desc: "ALL ON active / ALL OFF active",
					},
					Value{
						From: 0,
						To: 0,
						Desc: "ALL ON disabled/ ALL OFF disabled",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "ALL ON disabled/ ALL OFF active",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "ALL ON active / ALL OFF disabled",
					},
				},
			},
			&parameter{
				ID: 10,
				Name: "Control key #2 behaviour",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Turning on the shutter positioning function - Default",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Turning off the shutter positioning function",
					},
				},
			},
			&parameter{
				ID: 14,
				Name: "Inputs Button/Switch configuration",
				Size: 1,
				Values: []Value{
					Value{
						From: 0,
						To: 0,
						Desc: "Mono-stable input (button)",
					},
					Value{
						From: 1,
						To: 1,
						Desc: "Bi-stable input (switch)",
					},
					Value{
						From: 2,
						To: 2,
						Desc: "Single Mono-stable switch",
					},
				},
			},
			&parameter{
					},
}