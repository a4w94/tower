package tower

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Monster struct {
	MaxRound int `json:"maxRound"`
	MaxLive  int `json:"maxLive"`

	LevelBouns  []int `json:"LevelBonus"`
	BossMaxLive int   `json:"BossMaxLive"`

	GameGoals      map[int][]int           `json:"GameGoals"`
	KillBouns      map[string]float32      `json:"killbouns"`
	SymbolBouns    map[int]map[int]float64 `json:"SymbolBouns"`
	BounsTypeRate  map[int]map[int][]int   `json:"BounsTypeRate"`
	TimeStopSecond int                     `json:"TimeStopSecond"`
	FreeBullet     int                     `json:"FreeBullet"`
	FreeDouble     int                     `json:"FreeDouble"`

	MonsterRate     map[int]map[int][]int `json:"monsterRate"`
	MonsterDeadRate map[int][]float64     `json:"monsterDeadRate"`

	BossHitRate          map[int][]float64         ` json:"bossHitRate"`
	BossPayoffRateWeight map[int]map[int][]int     `json:"bossPayoffRateWeight"`
	BossPayoffRate       map[int]map[int][]float64 `json:"bossPayoffRate"`

	PayoffRateWeight map[int]map[int][]int     `json:"payoffRateWeight"`
	PayoffRate       map[int]map[int][]float64 `json:"payoffRate"`
	MonsterScript    map[int]map[int][]int     `json:"monsterScript"`
	MaxBet           int                       `json:"MaxBet"`
	Bet              int                       `json:"Bet"`
	CurrencyToRate   map[string][]int          `json:"CurrencyToRate"`
	DefaultBetBase   map[string]int            `json:"DefaultBetBase"`
}

// var (
// 	Monsterkillrate [][]float64

// 	GamePanel [][]int

// 	MonsterOdds [][]float64

// 	OddsWeight [][]int

// 	BossLifekillrate [][]float64

// 	BossOdds [][]float64

// 	BossOddsWeight [][]int

// 	MonsterWeight [][]int

// 	LevelReward []int

// 	BounsHitRate []float64

// 	BounsWeight [2][]int
// )

var (
	MonsterWeightJson   [][]int
	MonsterkillrateJson [][]float64
	OddsWeightJson      [][]int
	MonsterOddsJson     [][]float64
	GamePanelJson       [][]int
	LevelRewardJson     []int

	BossOddsJson         [][]float64
	BossLifekillrateJson [][]float64
	BossOddsWeightJson   [][]int

	KillTargetJson   []int
	RewardWeightJson [][]int
	BounsHitRateJson []float64
	BounsWeightJson  [2][]int
)

func TowerJson(excelroute string) []byte {
	GetExcelJson(excelroute)

	var result Monster
	result.MaxRound = 5
	result.MaxLive = 6

	result.BossMaxLive = 1
	result.KillBouns = map[string]float32{
		"Healing":      1,
		"MaxLiveBonus": 5,
		"PayoffBonus":  1.2,
	}
	result.MaxBet = 10
	result.Bet = 10
	result.BossMaxLive = 1
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

	//各關過關獎勵
	result.LevelBouns = LevelRewardJson

	//各關Boss血量命中率
	bosslifehitrate := map[int][]float64{}
	for i := 1; i < 11; i++ {

		bosslifehitrate[i] = BossLifekillrateJson[i]

	}
	result.BossHitRate = bosslifehitrate

	//各關Boss血量賠率
	bosslifeodd := map[int]map[int][]float64{}
	for i := 0; i < 10; i++ {
		temp := map[int][]float64{}
		for k := 1; k < 11; k++ {
			temp[k] = BossOddsJson[10*i+k][1:]

		}

		bosslifeodd[i+1] = temp

	}
	result.BossPayoffRate = bosslifeodd

	//各關Boos血量賠率權重
	bosslifeoddweight := map[int]map[int][]int{}
	for i := 0; i < 10; i++ {
		temp := map[int][]int{}
		for k := 1; k < 11; k++ {
			temp[k] = BossOddsWeightJson[10*i+k][1:]

		}

		bosslifeoddweight[i+1] = temp

	}
	result.BossPayoffRateWeight = bosslifeoddweight

	//bouns獎勵
	bonushitrate := map[int]map[int]float64{}
	for i := 0; i < 10; i++ {
		temp := map[int]float64{}
		for k := 0; k < 3; k++ {

			temp[k+1] = BounsHitRateJson[3*i+k]

		}
		bonushitrate[i+1] = temp
	}
	result.SymbolBouns = bonushitrate

	//fmt.Println(BounsWeight)
	bonusweight := map[int]map[int][]int{}
	for i := 0; i < 10; i++ {
		temp := map[int][]int{}
		for k := 0; k < 3; k++ {

			temp[k+1] = []int{0, BounsWeightJson[0][3*i+k], BounsWeightJson[1][3*i+k]}
		}
		bonusweight[i+1] = temp
	}
	//fmt.Println(bonusweight)
	result.BounsTypeRate = bonusweight

	targetarr := map[int][]int{}
	for i := 0; i < len(KillTargetJson); i++ {
		temp := []int{}
		temp = append(temp, KillTargetJson[i], RewardWeightJson[i][0], RewardWeightJson[i][1], RewardWeightJson[i][2])
		targetarr[i+1] = temp
	}
	result.GameGoals = targetarr

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
	rowkill := xlxs.GetRows("killprob")
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
	//fmt.Println("怪物擊殺率：", MonsterkillrateJson)
	//fmt.Println()

	///取得Boss血量擊殺率///
	temp10 := &BossLifekillrateJson
	rowbosskill := xlxs.GetRows("bosskillprob")
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
	//fmt.Println("Boss血量擊殺率：", BossLifekillrateJson)
	//fmt.Println()

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
	//fmt.Println()

	///取得浮動賠率///
	temp2 := &MonsterOddsJson
	rowodds := xlxs.GetRows("odds")
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
		*temp2 = append(*temp2, slice)
	}
	//fmt.Println("浮動賠率：", MonsterOddsJson)
	//fmt.Println()

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
	//fmt.Println()

	///取得Boss前Ｎ條浮動賠率///
	temp8 := &BossOddsJson
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
	//fmt.Println("Boss前Ｎ條命浮動賠率：", BossOddsJson)
	fmt.Println()

	///取得浮動賠率權重///
	temp9 := &BossOddsWeightJson
	rowbossweight := xlxs.GetRows("BossLifeOddsWeight")

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
	//fmt.Println("Boss浮動賠率權重：", BossOddsWeightJson)
	//fmt.Println()

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
	//fmt.Println("怪物各關出現權重：", MonsterWeightJson)

	//fmt.Println()

	//取得各關獎勵倍數//
	temp5 := &LevelRewardJson
	rowlevelreward := xlxs.GetRows("levelreward")

	for k := 1; k < len(rowlevelreward[2]); k++ {
		if rowlevelreward[2][k] == "" {
			continue
		} else {
			ele, _ := strconv.Atoi(rowlevelreward[2][k])
			*temp5 = append(*temp5, ele)
		}
	}

	//fmt.Println("各關獎勵倍數：", LevelRewardJson)

	//fmt.Println()

	//取得各關Bouns 權重與命中率//
	temp6 := &BounsHitRateJson
	rowbonus := xlxs.GetRows("bonus")

	for k := 1; k < len(rowbonus[2]); k++ {
		if rowbonus[2][k] == "" {
			continue
		} else {
			ele, _ := strconv.ParseFloat(rowbonus[2][k], 64)
			*temp6 = append(*temp6, ele)
		}
	}

	//fmt.Println("各關Bouns觸發率：", BounsHitRateJson)

	//fmt.Println()

	temp7 := &BounsWeightJson
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
	//fmt.Println("各關免費子彈權重：", BounsWeightJson[0])
	//fmt.Println("各關時間暫停權重：", BounsWeightJson[1])

	///取得擊殺目標數///
	temp11 := &KillTargetJson
	rowkilltarget := xlxs.GetRows("killtarget")

	for k := 1; k < len(rowkilltarget[1]); k++ {
		if rowkilltarget[1][k] == "" {
			continue
		} else {
			ele, _ := strconv.Atoi(rowkilltarget[1][k])
			*temp11 = append(*temp11, ele)
		}
	}

	//fmt.Println("各等級目標擊殺數：", KillTargetJson)

	//fmt.Println()

	rowkilltargetweight := xlxs.GetRows("targetbonusweight")
	temp12 := &RewardWeightJson

	for k := 1; k < len(rowkilltargetweight[1]); k++ {
		var slice []int

		for i := 2; i < len(rowkilltargetweight); i++ {

			if rowkilltargetweight[i][k] == "" {

				continue
			} else {
				ele, _ := strconv.Atoi(rowkilltargetweight[i][k])
				slice = append(slice, ele)
			}

		}
		*temp12 = append(*temp12, slice)
	}

	//fmt.Println("各等級獎勵權重：", RewardWeightJson)
	//fmt.Println()
	fmt.Println("獲取資料成功.....")
	fmt.Println()

}
