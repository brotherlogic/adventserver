package main

import (
	"testing"
)

func Test2022_16_1_Main(t *testing.T) {
	data := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	Valve BB has flow rate=13; tunnels lead to valves CC, AA
	Valve CC has flow rate=2; tunnels lead to valves DD, BB
	Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
	Valve EE has flow rate=3; tunnels lead to valves FF, DD
	Valve FF has flow rate=0; tunnels lead to valves EE, GG
	Valve GG has flow rate=0; tunnels lead to valves FF, HH
	Valve HH has flow rate=22; tunnel leads to valve GG
	Valve II has flow rate=0; tunnels lead to valves AA, JJ
	Valve JJ has flow rate=21; tunnel leads to valve II`

	val, _ := releaseGas(data, 30)

	if val != 1651 {
		t.Errorf("Bad gas release: %v (1651)", val)
	}
}

func Test2022_16_2_Main(t *testing.T) {
	data := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	Valve BB has flow rate=13; tunnels lead to valves CC, AA
	Valve CC has flow rate=2; tunnels lead to valves DD, BB
	Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
	Valve EE has flow rate=3; tunnels lead to valves FF, DD
	Valve FF has flow rate=0; tunnels lead to valves EE, GG
	Valve GG has flow rate=0; tunnels lead to valves FF, HH
	Valve HH has flow rate=22; tunnel leads to valve GG
	Valve II has flow rate=0; tunnels lead to valves AA, JJ
	Valve JJ has flow rate=21; tunnel leads to valve II`

	val := releaseGasSimple(data, 26)

	if val != 1707 {
		t.Errorf("Bad gas release: %v (1707)", val)
	}
}

func Test2022_16_2_Full(t *testing.T) {
	data := `Valve XB has flow rate=0; tunnels lead to valves YV, RP
	Valve VN has flow rate=0; tunnels lead to valves WL, ET
	Valve NT has flow rate=0; tunnels lead to valves CU, MQ
	Valve ON has flow rate=0; tunnels lead to valves AA, FP
	Valve CW has flow rate=0; tunnels lead to valves UH, WY
	Valve KN has flow rate=0; tunnels lead to valves JL, MQ
	Valve VT has flow rate=0; tunnels lead to valves CU, UI
	Valve CR has flow rate=0; tunnels lead to valves OA, QQ
	Valve YX has flow rate=0; tunnels lead to valves YJ, CI
	Valve WL has flow rate=7; tunnels lead to valves OQ, VN, PU, VF, UA
	Valve HV has flow rate=0; tunnels lead to valves OQ, OK
	Valve JM has flow rate=21; tunnels lead to valves RG, OH, JE
	Valve XF has flow rate=24; tunnels lead to valves OL, TM
	Valve VD has flow rate=0; tunnels lead to valves MY, OK
	Valve AA has flow rate=0; tunnels lead to valves KO, ON, UI, QE, VF
	Valve JE has flow rate=0; tunnels lead to valves JM, NZ
	Valve UN has flow rate=0; tunnels lead to valves UA, WY
	Valve CC has flow rate=0; tunnels lead to valves IV, CU
	Valve PU has flow rate=0; tunnels lead to valves JL, WL
	Valve UA has flow rate=0; tunnels lead to valves WL, UN
	Valve OJ has flow rate=13; tunnels lead to valves AZ, FP, MY, OL, ET
	Valve CJ has flow rate=0; tunnels lead to valves MQ, WS
	Valve IV has flow rate=0; tunnels lead to valves NZ, CC
	Valve NZ has flow rate=4; tunnels lead to valves WS, IV, IU, EQ, JE
	Valve TM has flow rate=0; tunnels lead to valves HL, XF
	Valve SG has flow rate=0; tunnels lead to valves MQ, OH
	Valve QQ has flow rate=12; tunnel leads to valve CR
	Valve WX has flow rate=15; tunnels lead to valves CI, SN
	Valve VF has flow rate=0; tunnels lead to valves WL, AA
	Valve RP has flow rate=0; tunnels lead to valves WY, XB
	Valve SN has flow rate=0; tunnels lead to valves WX, OI
	Valve HL has flow rate=0; tunnels lead to valves OK, TM
	Valve ET has flow rate=0; tunnels lead to valves OJ, VN
	Valve UI has flow rate=0; tunnels lead to valves AA, VT
	Valve FP has flow rate=0; tunnels lead to valves ON, OJ
	Valve IU has flow rate=0; tunnels lead to valves NZ, QE
	Valve JQ has flow rate=0; tunnels lead to valves HR, CU
	Valve CU has flow rate=5; tunnels lead to valves NT, VT, JQ, CC
	Valve WY has flow rate=19; tunnels lead to valves CW, UN, RP
	Valve YJ has flow rate=16; tunnel leads to valve YX
	Valve HR has flow rate=0; tunnels lead to valves JQ, JL
	Valve RM has flow rate=11; tunnels lead to valves OI, AZ
	Valve RG has flow rate=0; tunnels lead to valves YV, JM
	Valve MY has flow rate=0; tunnels lead to valves VD, OJ
	Valve QE has flow rate=0; tunnels lead to valves AA, IU
	Valve OK has flow rate=17; tunnels lead to valves HL, UH, VD, HV
	Valve CI has flow rate=0; tunnels lead to valves WX, YX
	Valve OL has flow rate=0; tunnels lead to valves XF, OJ
	Valve WS has flow rate=0; tunnels lead to valves CJ, NZ
	Valve OH has flow rate=0; tunnels lead to valves JM, SG
	Valve OQ has flow rate=0; tunnels lead to valves WL, HV
	Valve OA has flow rate=0; tunnels lead to valves CR, MQ
	Valve OI has flow rate=0; tunnels lead to valves SN, RM
	Valve YV has flow rate=25; tunnels lead to valves RG, XB
	Valve JL has flow rate=3; tunnels lead to valves KO, HR, PU, KN, EQ
	Valve AZ has flow rate=0; tunnels lead to valves OJ, RM
	Valve UH has flow rate=0; tunnels lead to valves CW, OK
	Valve KO has flow rate=0; tunnels lead to valves AA, JL
	Valve EQ has flow rate=0; tunnels lead to valves NZ, JL
	Valve MQ has flow rate=10; tunnels lead to valves CJ, OA, NT, SG, KN`

	val := releaseGasSimple(data, 26)

	if val != 1707 {
		t.Errorf("Bad gas release: %v (1707)", val)
	}
}
