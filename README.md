# golas
### Geophysical well log lexer and parser

 - Easily convert geophysical well logs to JSON
 - Currently, we only support the [CWLS LAS 2.0](http://www.cwls.org/wp-content/uploads/2014/09/LAS_20_Update_Jan2014.pdf) standard

--- 

# Example 

This test case is a good "hello world" example (using [this](/samples/unwrapped.las) .las file as input)

```golang
func TestNewLexer(t *testing.T) {
	r, e := os.Open("samples/unwrapped.las")
	if e != nil {
		panic("Unable to open file")
	}

	las := Parse(r)

	printData := func(data []Line) {
		for _, line := range data {
			fmt.Println("Mnemonic\t==", line.Mnem, "\nData\t\t==", line.Data, "\nUnits\t\t==", line.Units, "\nDesc\t\t==", line.Description)
			fmt.Println("-----------------------------------------------------------------------------------")
		}
	}

	fmt.Printf("\n\n===========\nVersion Info\n===========\n\n")
	printData(las.VersionInformation.Data)

	fmt.Printf("\n\n===========\nWell Info\n===========\n\n")
	printData(las.WellInformation.Data)

	fmt.Printf("\n\n===========\nCurve Info\n===========\n\n")
	printData(las.CurveInformation.Data)

	fmt.Printf("\n\n===========\nParam Info\n===========\n\n")
	printData(las.ParameterInformation.Data)

	fmt.Printf("\n\n===========\nOther Info\n===========\n\n")
	printData(las.Other.Data)

	for _, customSectn := range las.CustomSections {
		fmt.Printf("\n\n===========\n%s\n===========\n\n", customSectn.Name)
		printData(customSectn.Data)
	}
}
```

Which outputs

```
===========
Version Info
===========

Mnemonic	== VERS. 
Data		== 2.0: 
Units		==  
Desc		== CWLS LOG ASCII STANDARD -VERSION 2.0
-----------------------------------------------------------------------------------
Mnemonic	== WRAP. 
Data		== NO: 
Units		==  
Desc		== ONE LINE PER DEPTH STEP
-----------------------------------------------------------------------------------


===========
Well Info
===========

Mnemonic	== WELL. 
Data		== NORVEHC MGSU 1 MITSUE 01-01    : 
Units		==  
Desc		== Well_name    - WELL
-----------------------------------------------------------------------------------
Mnemonic	== LOC . 
Data		== 00/01-01-073-05W5/0            : 
Units		==  
Desc		== Location     - LOCATION
-----------------------------------------------------------------------------------
Mnemonic	== UWI . 
Data		== 00/01-01-073-05W5/0            : 
Units		==  
Desc		== Uwi          - UNIQUE WELL ID
-----------------------------------------------------------------------------------
Mnemonic	== ENTR. 
Data		== JOHN                           : 
Units		==  
Desc		== Entered      - ENTERED BY
-----------------------------------------------------------------------------------
Mnemonic	== SRVC. 
Data		== REGREBMULHCS                   : 
Units		==  
Desc		== Scn          - SERVICE COMPANY
-----------------------------------------------------------------------------------
Mnemonic	== DATE. 
Data		== 01 JAN 70                      : 
Units		==  
Desc		== Date         - LOG DATE
-----------------------------------------------------------------------------------
Mnemonic	== STRT. 
Data		== 390                            : 
Units		== M 
Desc		== top_depth    - START DEPTH
-----------------------------------------------------------------------------------
Mnemonic	== STOP. 
Data		== 650                            : 
Units		== M 
Desc		== bot_depth    - STOP DEPTH
-----------------------------------------------------------------------------------
Mnemonic	== STEP. 
Data		== 0.25                           : 
Units		== M 
Desc		== increment    - STEP LENGTH
-----------------------------------------------------------------------------------
Mnemonic	== NULL. 
Data		== -999.2500: 
Units		==  
Desc		== NULL Value
-----------------------------------------------------------------------------------


===========
Curve Info
===========

Mnemonic	== DEPT. 
Data		== 00 001 00 00         : 
Units		== M 
Desc		== DEPTH        - DEPTH
-----------------------------------------------------------------------------------
Mnemonic	== DPHI. 
Data		== 00 890 00 00         : 
Units		== V/V 
Desc		== PHID         - DENSITY POROSITY (SANDSTONE)
-----------------------------------------------------------------------------------
Mnemonic	== NPHI. 
Data		== 00 330 00 00         : 
Units		== V/V 
Desc		== PHIN         - NEUTRON POROSITY (SANDSTONE)
-----------------------------------------------------------------------------------
Mnemonic	== GR  . 
Data		== 00 310 00 00         : 
Units		== API 
Desc		== GR           - GAMMA RAY
-----------------------------------------------------------------------------------
Mnemonic	== CALI. 
Data		== 00 280 01 00         : 
Units		== MM 
Desc		== CAL          - CALIPER
-----------------------------------------------------------------------------------
Mnemonic	== ILD . 
Data		== 00 120 00 00         : 
Units		== OHMM 
Desc		== RESD         - DEEP RESISTIVITY (DIL)
-----------------------------------------------------------------------------------


===========
Param Info
===========

Mnemonic	== GL  . 
Data		== 583.3                : 
Units		== M 
Desc		== gl           - GROUND LEVEL ELEVATION
-----------------------------------------------------------------------------------
Mnemonic	== EREF. 
Data		== 589                  : 
Units		== M 
Desc		== kb           - ELEVATION OF DEPTH REFERENCE
-----------------------------------------------------------------------------------
Mnemonic	== DATM. 
Data		== 583.3                : 
Units		== M 
Desc		== datum        - DATUM ELEVATION
-----------------------------------------------------------------------------------
Mnemonic	== TDD . 
Data		== 733.4                : 
Units		== M 
Desc		== tdd          - TOTAL DEPTH DRILLER
-----------------------------------------------------------------------------------
Mnemonic	== RUN . 
Data		== ONE                  : 
Units		==  
Desc		== Run          - RUN NUMBER
-----------------------------------------------------------------------------------
Mnemonic	== ENG . 
Data		== SIMMONS              : 
Units		==  
Desc		== Engineer     - RECORDING ENGINEER
-----------------------------------------------------------------------------------
Mnemonic	== WIT . 
Data		== SANK                 : 
Units		==  
Desc		== Witness      - WITNESSED BY
-----------------------------------------------------------------------------------
Mnemonic	== BASE. 
Data		== S.L.                 : 
Units		==  
Desc		== Branch       - HOME BASE OF LOGGING UNIT
-----------------------------------------------------------------------------------
Mnemonic	== MUD . 
Data		== GEL CHEM             : 
Units		==  
Desc		== Mud_type     - MUD TYPE
-----------------------------------------------------------------------------------
Mnemonic	== MATR. 
Data		== SANDSTONE            : 
Units		==  
Desc		== Logunit      - NEUTRON MATRIX
-----------------------------------------------------------------------------------
Mnemonic	== TMAX. 
Data		== 41                   : 
Units		== C 
Desc		== BHT          - MAXIMUM RECORDED TEMPERATURE
-----------------------------------------------------------------------------------
Mnemonic	== BHTD. 
Data		== 733.8                : 
Units		== M 
Desc		== BHTDEP       - MAXIMUM RECORDED TEMPERATURE
-----------------------------------------------------------------------------------
Mnemonic	== RMT . 
Data		== 17                   : 
Units		== C 
Desc		== MDTP         - TEMPERATURE OF MUD
-----------------------------------------------------------------------------------
Mnemonic	== MUDD. 
Data		== 1100                 : 
Units		== KG/M 
Desc		== MWT          - MUD DENSITY
-----------------------------------------------------------------------------------
Mnemonic	== NEUT. 
Data		== 1                    : 
Units		==  
Desc		== NEUTRON      - NEUTRON TYPE
-----------------------------------------------------------------------------------
Mnemonic	== RESI. 
Data		== 0                    : 
Units		==  
Desc		== RESIST       - RESISTIVITY TYPE
-----------------------------------------------------------------------------------
Mnemonic	== RM  . 
Data		== 2.62                 : 
Units		== OHMM 
Desc		== RM           - RESISTIVITY OF MUD
-----------------------------------------------------------------------------------
Mnemonic	== RMC . 
Data		== 0                    : 
Units		== OHMM 
Desc		== RMC          - RESISTIVITY OF MUD CAKE
-----------------------------------------------------------------------------------
Mnemonic	== RMF . 
Data		== 1.02                 : 
Units		== OHMM 
Desc		== RMF          - RESISTIVITY OF MUD FILTRATE
-----------------------------------------------------------------------------------
Mnemonic	== SUFT. 
Data		== 0                    : 
Units		== C 
Desc		== SUFT         - SURFACE TEMPERATURE
-----------------------------------------------------------------------------------


===========
Other Info
===========
```