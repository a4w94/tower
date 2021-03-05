package tower

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Luxurioust/excelize"
)

type Monster struct {
	MaxRound         int                       `json:"maxRound"`
	MaxLive          int                       `json:"maxLive"`
	MonsterRate      map[int]map[int][]int     `json:"monsterRate"`
	MonsterDeadRate  map[int][]float64         `json:"monsterDeadRate"`
	PayoffRateWeight map[int]map[int][]int     `json:"payoffRateWeight"`
	PayoffRate       map[int]map[int][]float64 `json:"payoffRate"`
	MonsterScript    map[int]map[int][]int     `json:"monsterScript"`
	MaxBet           int                       `json:"MaxBet"`
	Bet              int                       `json:"Bet"`
	CurrencyToRate   map[string][]int          `json:"CurrencyToRate"`
	DefaultBetBase   map[string]int            `json:"DefaultBetBase"`
}

var (
	Monsterkillrate [][]float64

	GamePanel [][]int

	MonsterOdds [][]float64

	OddsWeight [][]int

	BossLifekillrate [][]float64

	BossOdds [][]float64

	BossOddsWeight [][]int

	MonsterWeight [][]int

	LevelReward []int

	BounsHitRate []float64

	BounsWeight [2][]int
)

var (
	MonsterWeightJson   [][]int
	MonsterkillrateJson [][]float64
	OddsWeightJson      [][]int
	MonsterOddsJson     [][]float64
	GamePanelJson       [][]int
)

func TowerJson(excelroute string) []byte {
	GetExcelJson(excelroute)

	var result Monster
	result.MaxRound = 10
	result.MaxLive = 6
	result.MaxBet = 10
	result.Bet = 10
	result.CurrencyToRate = map[string][]int{
		"HKD": []int{7},
		"IDR": []int{14},
		"JPY": []int{10},
		"KRW": []int{13},
		"MYR": []int{7},
		"RMB": []int{7},
		"SGD": []int{7},
		"THB": []int{9},
		"TWD": []int{9},
		"USD": []int{7},
		"EUR": []int{7},
		"GBP": []int{7},
		"VND": []int{14},
		"INR": []int{10},
	}

	result.DefaultBetBase = map[string]int{
		"HKD": 7,
		"IDR": 14,
		"JPY": 10,
		"KRW": 13,
		"MYR": 7,
		"RMB": 7,
		"SGD": 7,
		"THB": 9,
		"TWD": 9,
		"USD": 7,
		"EUR": 7,
		"GBP": 7,
		"VND": 14,
		"INR": 10,
	}

	//怪物出現權重
	monsterrate := map[int]map[int][]int{}
	for i := 0; i < 10; i++ {
		temp := map[int][]int{}
		for k := 1; k < 4; k++ {
			temp[k] = MonsterWeightJson[3*i+k]

		}

		monsterrate[i+1] = temp

	}
	result.MonsterRate = monsterrate

	//各關怪物死亡率

	mosterdeadrate := map[int][]float64{}
	for i := 1; i < 11; i++ {

		mosterdeadrate[i] = MonsterkillrateJson[i]

	}
	result.MonsterDeadRate = mosterdeadrate

	//各關怪物賠率權重
	monsteroddweight := map[int]map[int][]int{}
	for i := 0; i < 10; i++ {
		temp := map[int][]int{}
		for k := 1; k < 11; k++ {
			temp[k] = OddsWeightJson[10*i+k][1:]

		}

		monsteroddweight[i+1] = temp

	}
	result.PayoffRateWeight = monsteroddweight

	//各關怪物賠率
	monsterodd := map[int]map[int][]float64{}
	for i := 0; i < 10; i++ {
		temp := map[int][]float64{}
		for k := 1; k < 11; k++ {
			temp[k] = MonsterOddsJson[10*i+k][1:]

		}

		monsterodd[i+1] = temp

	}
	result.PayoffRate = monsterodd

	//各關怪物出現盤面
	monsterpanel := map[int]map[int][]int{}
	for i := 0; i < 10; i++ {
		temp := map[int][]int{}
		for k := 1; k < 4; k++ {
			temp[k] = GamePanelJson[3*i+k]

		}

		monsterpanel[i+1] = temp

	}
	result.MonsterScript = monsterpanel

	resultjson, _ := json.Marshal(result)
	//fmt.Println(string(resultjson))

	return resultjson

}

func GetExcelJson(excelroute string) {
	xlxs, err := excelize.OpenFile(excelroute)
	if err != nil {
		fmt.Println("excel open failed", err)
	}

	///取得怪物擊殺率///
	temp0 := &MonsterkillrateJson
	rowkill, _ := xlxs.GetRows("killprob")
	*temp0 = append(*temp0, []float64{})

	for i := 0; i < 10; i++ {
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
	fmt.Println("怪物擊殺率：", MonsterkillrateJson)
	fmt.Println()

	///取得Boss血量擊殺率///
	temp10 := &BossLifekillrate
	rowbosskill, _ := xlxs.GetRows("bosskillprob")
	*temp10 = append(*temp10, []float64{})
	for i := 0; i < 10; i++ {
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

	///取得盤面///
	temp1 := &GamePanelJson
	rowpanel, _ := xlxs.GetRows("panel")

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
	fmt.Println("盤面：", GamePanelJson)
	fmt.Println()

	///取得浮動賠率///
	temp2 := &MonsterOddsJson
	rowodds, _ := xlxs.GetRows("odds")
	//fmt.Println(rowodds)

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
		fmt.Println(i, slice)
		*temp2 = append(*temp2, slice)
	}
	fmt.Println("浮動賠率：", MonsterOddsJson)
	fmt.Println()

	///取得浮動賠率權重///
	temp3 := &OddsWeightJson
	rowweight, _ := xlxs.GetRows("賠率權重")

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
	fmt.Println("浮動賠率權重：", OddsWeightJson)
	fmt.Println()

	///取得Boss前Ｎ條浮動賠率///
	temp8 := &BossOdds
	rowbossodds, _ := xlxs.GetRows("bossodds")

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
	rowbossweight, _ := xlxs.GetRows("bossweight")

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
	temp4 := &MonsterWeightJson
	rowmonsterweight, _ := xlxs.GetRows("怪物各關出現權重")

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
	fmt.Println("怪物各關出現權重：", MonsterWeightJson)

	fmt.Println()

	//取得各關獎勵倍數//
	temp5 := &LevelReward
	rowlevelreward, _ := xlxs.GetRows("levelreward")

	for k := 1; k < len(rowlevelreward[2]); k++ {
		if rowlevelreward[2][k] == "" {
			continue
		} else {
			ele, _ := strconv.Atoi(rowlevelreward[2][k])
			*temp5 = append(*temp5, ele)
		}
	}

	fmt.Println("各關獎勵倍數：", LevelReward)

	fmt.Println()

	//取得各關Bouns 權重與命中率//
	temp6 := &BounsHitRate
	rowbonus, _ := xlxs.GetRows("bonus")

	for k := 1; k < len(rowbonus[2]); k++ {
		if rowbonus[2][k] == "" {
			continue
		} else {
			ele, _ := strconv.ParseFloat(rowbonus[2][k], 64)
			*temp6 = append(*temp6, ele)
		}
	}

	fmt.Println("各關Bouns觸發率：", BounsHitRate)

	fmt.Println()

	temp7 := &BounsWeight
	for k := 1; k < len(rowbonus[3]); k++ {
		if rowbonus[3][k] == "" {
			continue
		} else {
			ele, _ := strconv.Atoi(rowbonus[3][k])
			ele1, _ := strconv.Atoi(rowbonus[4][k])

			temp7[0] = append(temp7[0], ele)
			temp7[1] = append(temp7[1], ele1)

		}
	}
	fmt.Println("各關免費子彈權重：", BounsWeight[0])
	fmt.Println("各關時間暫停權重：", BounsWeight[1])

	fmt.Println()
	fmt.Println("獲取資料成功.....")
	fmt.Println()

}
