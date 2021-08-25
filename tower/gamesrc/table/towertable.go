package towerdefense

import (
	"fmt"
	"strconv"
	info "towerjs/tower/gamesrc/info"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var (
	Monsterkillrate   [][]float64
	Monsterkillrate95 [][]float64
	Monsterkillrate99 [][]float64

	GamePanel [][]int

	MonsterOdds [][]float64

	OddsWeight   [][]int
	OddsWeight95 [][]int
	OddsWeight99 [][]int

	BossLifekillrate [][]float64

	BossOdds [][]float64

	BossOddsWeight [][]int

	MonsterWeight [][]int

	KillTarget []int

	RewardWeight [][]int
)

var (
	MonsterWeightJson   [][]int
	MonsterkillrateJson [][]float64
	OddsWeightJson      [][]int
	MonsterOddsJson     [][]float64
	GamePanelJson       [][]int
)

func GetExcelJson() {
	xlxs, err := excelize.OpenFile("/Users/terry_hsiesh/go/game/Other/towerdefense玩家模擬/tower/gameset.xlsx")
	if err != nil {
		fmt.Println("excel open failed", err)
	}

	///取得怪物擊殺率///
	temp0 := &MonsterkillrateJson
	rowkill := xlxs.GetRows("killprob")
	*temp0 = append(*temp0, []float64{})

	for i := 0; i < info.GameLevel-1; i++ {
		var slice = []float64{0}
		for k := 10*i + 2; k < 10*i+12; k++ {
			if rowkill[k][2] == "" {
				continue
			} else {
				ele, _ := strconv.ParseFloat(rowkill[k][2], 64)
				slice = append(slice, ele)

			}

		}

		*temp0 = append(*temp0, slice)
	}
	//fmt.Println("怪物擊殺率：", MonsterkillrateJson)
	fmt.Println()

	// ///取得Boss血量擊殺率///
	// temp10 := &BossLifekillrate
	// rowbosskill:= xlxs.GetRows("bosskillprob")
	// *temp10 = append(*temp10, []float64{})
	// for i := 0; i < 10; i++ {
	// 	var slice = []float64{0}
	// 	for k := 10*i + 2; k < 10*i+12; k++ {
	// 		if rowbosskill[k][2] == "" {
	// 			continue
	// 		} else {
	// 			ele, _ := strconv.ParseFloat(rowbosskill[k][2], 64)
	// 			slice = append(slice, ele)

	// 		}

	// 	}
	// 	*temp10 = append(*temp10, slice)
	// }
	// //fmt.Println("Boss血量擊殺率：", BossLifekillrate)
	// fmt.Println()

	///取得盤面///
	temp1 := &GamePanelJson
	rowpanel := xlxs.GetRows("panel")

	for i := 1; i < len(rowpanel); i++ {
		var slice []int
		var findmonster bool
		for k := 1; k < len(rowpanel[i]); k++ {

			if rowpanel[i][k] == "" {
				continue
			} else {

				ele, _ := strconv.Atoi(rowpanel[i][k])
				if ele != 0 {
					findmonster = true
				}

				if findmonster == true {
					slice = append(slice, ele)

				}

			}
		}
		*temp1 = append(*temp1, slice)
	}
	//fmt.Println("盤面：", GamePanelJson)
	fmt.Println()

	///取得浮動賠率///
	temp2 := &MonsterOddsJson
	rowodds := xlxs.GetRows("odds")

	for i := 1; i < len(rowodds); i++ {
		var slice []float64
		for k := 2; k < len(rowodds[i]); k++ {
			if rowodds[i][k] == "" {
				continue
			} else {

				ele, _ := strconv.ParseFloat(rowodds[i][k], 64)

				slice = append(slice, ele)
			}
		}
		*temp2 = append(*temp2, slice)
	}
	//fmt.Println("浮動賠率：", MonsterOddsJson)
	fmt.Println()

	///取得浮動賠率權重///
	temp3 := &OddsWeightJson
	rowweight := xlxs.GetRows("賠率權重")

	for i := 1; i < len(rowweight); i++ {
		var slice []int
		for k := 2; k < len(rowweight[i]); k++ {
			if rowweight[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowweight[i][k])
				slice = append(slice, ele)
			}
		}
		*temp3 = append(*temp3, slice)
	}
	//fmt.Println("浮動賠率權重：", OddsWeightJson)
	fmt.Println()

	// ///取得Boss前Ｎ條浮動賠率///
	// temp8 := &BossOdds
	// rowbossodds:= xlxs.GetRows("bossodds")

	// for i := 1; i < len(rowbossodds); i++ {
	// 	var slice []float64
	// 	for k := 2; k < len(rowbossodds[i]); k++ {
	// 		if rowbossodds[i][k] == "" {
	// 			continue
	// 		} else {

	// 			ele, _ := strconv.ParseFloat(rowbossodds[i][k], 64)

	// 			slice = append(slice, ele)
	// 		}
	// 	}
	// 	*temp8 = append(*temp8, slice)
	// }
	// //fmt.Println("Boss前Ｎ條命浮動賠率：", BossOdds)
	// fmt.Println()

	// ///取得浮動賠率權重///
	// temp9 := &BossOddsWeight
	// rowbossweight:= xlxs.GetRows("bossweight")

	// for i := 1; i < len(rowbossweight); i++ {
	// 	var slice []int
	// 	for k := 2; k < len(rowbossweight[i]); k++ {
	// 		if rowbossweight[i][k] == "" {
	// 			continue
	// 		} else {
	// 			ele, _ := strconv.Atoi(rowbossweight[i][k])
	// 			slice = append(slice, ele)
	// 		}
	// 	}
	// 	*temp9 = append(*temp9, slice)
	// }
	// //fmt.Println("Boss浮動賠率權重：", BossOddsWeight)
	// fmt.Println()

	///取得怪物各關出現權重///
	temp4 := &MonsterWeightJson
	rowmonsterweight := xlxs.GetRows("怪物各關出現權重")

	for i := 1; i < len(rowmonsterweight); i++ {
		var slice []int
		for k := 1; k < len(rowmonsterweight[i]); k++ {
			if rowmonsterweight[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowmonsterweight[i][k])
				slice = append(slice, ele)
			}
		}
		*temp4 = append(*temp4, slice)
	}

	fmt.Println()
	fmt.Println("Json獲取資料成功.....")
	fmt.Println()

}

func GetExcel() {
	xlxs, err := excelize.OpenFile("/Users/terry_hsiesh/go/game/Other/towerdefense玩家模擬/tower/gameset.xlsx")
	if err != nil {
		fmt.Println("excel open failed", err)
	}
	///取得盤面///
	temp1 := &GamePanel
	rowpanel := xlxs.GetRows("panel")
	fmt.Println(len(rowpanel))

	for i := 1; i < len(rowpanel); i++ {
		var slice []int
		for k := 1; k < len(rowpanel[i]); k++ {
			if rowpanel[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowpanel[i][k])
				slice = append(slice, ele)
			}
		}
		*temp1 = append(*temp1, slice)
	}
	fmt.Println("盤面：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", GamePanel[3*(i-1)+k])
		}
	}

	var countnumarr []int
	var totalnum int
	for i := 0; i < len(GamePanel); i++ {
		var num int
		for k := 0; k < len(GamePanel[i]); k++ {
			if GamePanel[i][k] != 0 {
				num++
			}
		}
		totalnum += num
		countnumarr = append(countnumarr, num)
	}
	fmt.Println("各關怪物總數：", countnumarr, "怪物總數：", totalnum)

	fmt.Println()

	///取得怪物擊殺率///
	temp0 := &Monsterkillrate
	rowkill := xlxs.GetRows("killprob")
	*temp0 = append(*temp0, []float64{})

	for i := 0; i < info.GameLevel-1; i++ {
		var slice = []float64{0}
		for k := 10*i + 2; k < 10*i+12; k++ {
			if rowkill[k][2] == "" {
				continue
			} else {
				ele, _ := strconv.ParseFloat(rowkill[k][2], 64)
				slice = append(slice, ele)

			}

		}

		*temp0 = append(*temp0, slice)
	}
	fmt.Println("怪物擊殺率：")
	for i := 1; i < info.GameLevel; i++ {

		fmt.Println("關卡", i, ":", Monsterkillrate[i])

	}
	fmt.Println()

	///取得浮動賠率///
	temp2 := &MonsterOdds
	rowodds := xlxs.GetRows("odds")

	for i := 1; i < len(rowodds); i++ {
		var slice []float64
		for k := 2; k < len(rowodds[i]); k++ {
			if rowodds[i][k] == "" {
				continue
			} else {

				ele, _ := strconv.ParseFloat(rowodds[i][k], 64)

				slice = append(slice, ele)
			}
		}
		*temp2 = append(*temp2, slice)
	}
	fmt.Println("浮動賠率：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 11; k++ {
			fmt.Println("關卡", i, "", "怪物", k, ":", MonsterOdds[10*(i-1)+k])
		}
		fmt.Println()
	}

	fmt.Println()

	///取得浮動賠率權重///
	temp3 := &OddsWeight
	rowweight := xlxs.GetRows("weight")

	for i := 1; i < len(rowweight); i++ {
		var slice []int
		for k := 2; k < len(rowweight[i]); k++ {
			if rowweight[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowweight[i][k])
				slice = append(slice, ele)
			}
		}
		*temp3 = append(*temp3, slice)
	}
	fmt.Println("浮動賠率權重：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 11; k++ {
			fmt.Println("關卡", i, "", "怪物", k, ":", OddsWeightJson[10*(i-1)+k])
		}
		fmt.Println()
	}

	fmt.Println()

	///取得Boss血量擊殺率///
	temp10 := &BossLifekillrate
	rowbosskill := xlxs.GetRows("bosskillprob")
	*temp10 = append(*temp10, []float64{})
	for i := 0; i < info.GameLevel-1; i++ {
		var slice = []float64{0}
		for k := 10*i + 2; k < 10*i+12; k++ {
			if rowbosskill[k][2] == "" {
				continue
			} else {
				ele, _ := strconv.ParseFloat(rowbosskill[k][2], 64)
				slice = append(slice, ele)

			}

		}
		*temp10 = append(*temp10, slice)
	}
	fmt.Println("Boss血量擊殺率：", BossLifekillrate)
	fmt.Println()

	///取得Boss前Ｎ條浮動賠率///
	temp8 := &BossOdds
	rowbossodds := xlxs.GetRows("bossodds")

	for i := 1; i < len(rowbossodds); i++ {
		var slice []float64
		for k := 2; k < len(rowbossodds[i]); k++ {
			if rowbossodds[i][k] == "" {
				continue
			} else {

				ele, _ := strconv.ParseFloat(rowbossodds[i][k], 64)

				slice = append(slice, ele)
			}
		}
		*temp8 = append(*temp8, slice)
	}
	fmt.Println("Boss前Ｎ條命浮動賠率：", BossOdds)
	fmt.Println()

	///取得浮動賠率權重///
	temp9 := &BossOddsWeight
	rowbossweight := xlxs.GetRows("bossweight")

	for i := 1; i < len(rowbossweight); i++ {
		var slice []int
		for k := 2; k < len(rowbossweight[i]); k++ {
			if rowbossweight[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowbossweight[i][k])
				slice = append(slice, ele)
			}
		}
		*temp9 = append(*temp9, slice)
	}
	fmt.Println("Boss浮動賠率權重：", BossOddsWeight)
	fmt.Println()

	///取得怪物各關出現權重///
	temp4 := &MonsterWeight
	rowmonsterweight := xlxs.GetRows("monsterweight")

	for i := 1; i < len(rowmonsterweight); i++ {
		var slice []int
		for k := 1; k < len(rowmonsterweight[i]); k++ {
			if rowmonsterweight[i][k] == "" {
				continue
			} else {
				ele, _ := strconv.Atoi(rowmonsterweight[i][k])
				slice = append(slice, ele)
			}
		}
		*temp4 = append(*temp4, slice)
	}
	fmt.Println("怪物各關出現權重：", MonsterWeight)

	fmt.Println()

	///取得擊殺目標數///
	temp5 := &KillTarget
	rowkilltarget := xlxs.GetRows("killtarget")

	for k := 1; k < len(rowkilltarget[1]); k++ {
		if rowkilltarget[1][k] == "" {
			continue
		} else {
			ele, _ := strconv.Atoi(rowkilltarget[1][k])
			*temp5 = append(*temp5, ele)
		}
	}

	fmt.Println("各等級目標擊殺數：", KillTarget)

	fmt.Println()

	temp6 := &RewardWeight

	for k := 1; k < len(rowkilltarget[1]); k++ {
		var slice []int

		for i := 2; i < len(rowkilltarget); i++ {

			if rowkilltarget[i][k] == "" {

				continue
			} else {
				ele, _ := strconv.Atoi(rowkilltarget[i][k])
				slice = append(slice, ele)
			}

		}
		*temp6 = append(*temp6, slice)
	}

	fmt.Println("各等級獎勵權重：", RewardWeight)

	fmt.Println()

	fmt.Println()
	fmt.Println("獲取資料成功.....")
	fmt.Println()

}
