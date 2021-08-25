package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"time"
	info "towerjs/tower/gamesrc/info"
	table "towerjs/tower/gamesrc/table"
	rng "towerjs/tower/gamesrc/tool"
)

const (
	printcontrol            = info.Printcontrol
	bombprintcontrol        = info.BombPrintcontrol
	MoveSet                 = info.MoveSet
	gameround           int = 16
	freebulletset       int = info.Freebulletset
	stoptimesset        int = info.Stoptimesset
	ecahroundplayercost     = info.PlayerCost
	ecahroundplayerlife     = info.PlayerLife
)

var (
	monsterkillrate = table.Monsterkillrate
	GamePanel       = table.GamePanel
	monsterodds     = table.MonsterOdds

	oddsweight = table.OddsWeight

	GameEndLevel []int
)

type TotalResult struct {

	//各關得分統計（未含過關獎勵）
	TotalRoundScore [gameround]float64
	//各關花費統計
	TotalRoundCost [gameround]int
	//各關ＲＴＰ（未含過關獎勵）
	TotalRoundRTP [gameround]float64

	//各關累積ＲＴＰ(含過關獎勵)
	TotalRoundAccumulateRTP [gameround]float64

	TotalRoundMonster [gameround][11]int

	TotalRoundMainScore   [gameround]float64
	TotalRoundMainGameRTP [gameround]float64

	TotalRoundMonsterCost  [gameround][11]int
	TotalRoundMonsterScore [gameround][11]float64
	TotalRoundMonsterRTP   [gameround][11]float64

	TotalGameEndLevel       [gameround]int
	TotalGameEndLevelRate   [gameround]float64
	TotalGameEnterLevel     [gameround]int
	TotalGameEnterLevelRate [gameround]float64

	TotalRoundRewardScore  [3][gameround]float64
	TotalRoundRewardRTP    [3][gameround]float64
	TotalLevelRewardScore  [3][info.RewardLevel]float64
	TotalLevelRewardRTP    [3][info.RewardLevel]float64
	TotalRewardTimes       [3]int
	TotalLRoundRewardTimes [gameround][3]int
	TotalRewardHitRate     [3]float64

	TotalRewardScore float64
	TotalRewardRTP   float64

	TotalMainGameScore float64
	MainGameRTP        float64

	EachLevelRTP [11]float64

	EachRoundScorePercentage [gameround]float64
	EachRoundCostPercentage  [gameround]float64

	TotalCost  int
	TotalScore float64
	TotalRTP   float64
}

type PlayerEachRoundResult struct {
	EachGameEnterLevel [gameround]int
	EachRoundMonster   [gameround][11]int

	EachRoundMainScore [gameround]float64
	EachRoundScore     [gameround]float64
	EachRoundCost      [gameround]int

	EachRoundRewardScore [3][gameround]float64
	EachLevelRewardScore [3][info.RewardLevel]float64
	EachRoundRewardTimes [gameround][3]int

	EachGameEndLevel int
	PlayerCost       int
	PlayerScore      float64

	EachRoundMonsterCost  [gameround][11]int
	EachRoundMonsterScore [gameround][11]float64

	EachRewardTimes [3]int
}

func main() {

	table.GetExcelJson()
	table.GetExcel()
	Point()

	var totalresult = TotalResult{}

	var session = info.Session
	s := time.Now()
	rand.Seed(int64(time.Now().UnixNano()))

	for i := 0; i < session; i++ {

		eachresult := PlayProcess()

		// fmt.Println(eachresult.EachRoundMonsterCost)
		// fmt.Println(eachresult.EachRoundMonsterScore)
		// //fmt.Println("result:", eachresult)

		//統計玩家總花費
		totalresult.TotalCost += eachresult.PlayerCost

		//統計玩家總得分
		totalresult.TotalScore += eachresult.PlayerScore

		//統計玩家結束關卡次數
		totalresult.TotalGameEndLevel[eachresult.EachGameEndLevel]++

		for i := 0; i < len(eachresult.EachRoundCost); i++ {
			//統計各關花費
			totalresult.TotalRoundCost[i] += eachresult.EachRoundCost[i]

			//統計各關得分
			totalresult.TotalRoundScore[i] += eachresult.EachRoundScore[i]

		}

		///計算各關怪物出現
		for i := 0; i < len(eachresult.EachRoundMonster); i++ {
			for k := 0; k < len(eachresult.EachRoundMonster[i]); k++ {
				totalresult.TotalRoundMonster[i][k] += eachresult.EachRoundMonster[i][k]
			}
		}

		///計算各關怪物花費與得分
		for i := 0; i < len(totalresult.TotalRoundMonsterCost); i++ {
			for k := 0; k < len(totalresult.TotalRoundMonsterCost[i]); k++ {

				totalresult.TotalRoundMonsterCost[i][k] += eachresult.EachRoundMonsterCost[i][k]
				totalresult.TotalRoundMonsterScore[i][k] += eachresult.EachRoundMonsterScore[i][k]
			}
		}

		for i := 0; i < gameround; i++ {
			//統計玩家進入關卡次數
			totalresult.TotalGameEnterLevel[i] += eachresult.EachGameEnterLevel[i]

			//計算各關一般子彈得分//
			totalresult.TotalRoundMainScore[i] += eachresult.EachRoundMainScore[i]

			//統計各關獎勵總得分//
			for k := 0; k < 3; k++ {
				totalresult.TotalRoundRewardScore[k][i] += eachresult.EachRoundRewardScore[k][i]
			}

		}
		for i := 0; i < info.RewardLevel; i++ {
			for k := 0; k < 3; k++ {
				totalresult.TotalLevelRewardScore[k][i] += eachresult.EachLevelRewardScore[k][i]

				totalresult.TotalRewardScore += eachresult.EachLevelRewardScore[k][i]
			}
		}

		for i := 0; i < 3; i++ {
			totalresult.TotalRewardTimes[i] += eachresult.EachRewardTimes[i]
		}
		for i := 0; i < gameround; i++ {
			for k := 0; k < 3; k++ {
				totalresult.TotalLRoundRewardTimes[i][k] += eachresult.EachRoundRewardTimes[i][k]
			}
		}
	}

	fmt.Println("各關一般子彈得分", totalresult.TotalRoundMainScore)

	for i := 0; i < gameround; i++ {
		//計算各關ＲＴＰ（不含過關獎勵）
		totalresult.TotalRoundRTP[i] = totalresult.TotalRoundScore[i] / float64(totalresult.TotalRoundCost[i])

		//計算各關一般子彈ＲＴＰ（不含過關獎勵）
		totalresult.TotalRoundMainGameRTP[i] = totalresult.TotalRoundMainScore[i] / float64(totalresult.TotalRoundCost[i])

		//計算一般子彈總得分
		totalresult.TotalMainGameScore += totalresult.TotalRoundMainScore[i]

		//統計玩家結束關卡比率
		totalresult.TotalGameEndLevelRate[i] = float64(totalresult.TotalGameEndLevel[i]) / float64(session)

		//統計玩家進入關卡比率
		totalresult.TotalGameEnterLevelRate[i] = float64(totalresult.TotalGameEnterLevel[i]) / float64(session)

		//計算各關得分佔比（不含過關獎勵）
		totalresult.EachRoundScorePercentage[i] = float64(totalresult.TotalRoundMainScore[i]) / float64(totalresult.TotalScore)

		//計算各關花費佔比
		totalresult.EachRoundCostPercentage[i] = float64(totalresult.TotalRoundCost[i]) / float64(totalresult.TotalCost)

		//計算各關獎勵ＲＴＰ（分母為各關得分)//
		for k := 0; k < 3; k++ {
			totalresult.TotalRoundRewardRTP[k][i] = totalresult.TotalRoundRewardScore[k][i] / float64(totalresult.TotalRoundCost[i])
		}

	}

	for i := 0; i < info.RewardLevel; i++ {
		for k := 0; k < 3; k++ {
			totalresult.TotalLevelRewardRTP[k][i] = totalresult.TotalLevelRewardScore[k][i] / float64(totalresult.TotalScore)
		}
	}

	totalresult.TotalRewardRTP = totalresult.TotalRewardScore / float64(totalresult.TotalCost)

	//統計各關累積ＲＴＰ
	for i := 0; i < gameround; i++ {
		var score float64
		var cost int

		for k := 0; k <= i; k++ {
			score += totalresult.TotalRoundScore[k]
			cost += totalresult.TotalRoundCost[k]
		}

		totalresult.TotalRoundAccumulateRTP[i] = score / float64(cost)

	}
	fmt.Println()

	//統計各等級
	for i := 1; i < info.GameLevel; i++ {
		var levelcost int
		var levelscore float64

		for k := 1; k < 4; k++ {
			levelscore += totalresult.TotalRoundScore[3*(i-1)+k]
			levelcost += totalresult.TotalRoundCost[3*(i-1)+k]

		}
		totalresult.EachLevelRTP[i] = (levelscore) / float64(levelcost)
	}

	fmt.Println("各等級ＲＴＰ(一整關):")

	for i := 1; i < 11; i++ {

		fmt.Println("等級:", i, ":", "ＲＴＰ：", totalresult.EachLevelRTP[i])

	}

	fmt.Println()

	totalresult.MainGameRTP = totalresult.TotalMainGameScore / float64(totalresult.TotalCost)

	fmt.Println("玩家各關得分：", totalresult.TotalRoundScore)
	//fmt.Println("玩家各關花費：", totalresult.TotalRoundCost)

	fmt.Println("各關累積ＲＴＰ(含觸發擊殺獎勵)：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", totalresult.TotalRoundAccumulateRTP[3*(i-1)+k])
		}
	}
	fmt.Println()

	fmt.Println("各關獨立ＲＴＰ(含觸發擊殺獎勵)：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", totalresult.TotalRoundRTP[3*(i-1)+k])
		}
	}
	fmt.Println()

	for i := 0; i < len(totalresult.TotalRoundMonsterRTP); i++ {
		for k := 0; k < len(totalresult.TotalRoundMonsterRTP[i]); k++ {
			totalresult.TotalRoundMonsterRTP[i][k] = totalresult.TotalRoundMonsterScore[i][k] / float64(totalresult.TotalRoundMonsterCost[i][k])
		}
	}

	fmt.Println()
	//fmt.Println("各關一般子彈得分：", totalresult.TotalRoundMainScore)
	fmt.Println("各關一般子彈RTP：", totalresult.TotalRoundMainGameRTP)

	// fmt.Println("各關怪物分佈：", totalresult.TotalRoundMonster)
	// fmt.Println("各關怪獸得分：", totalresult.TotalRoundMonsterScore)
	// fmt.Println("各關怪獸花費：", totalresult.TotalRoundMonsterCost)
	fmt.Println("各關怪獸RTP：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", totalresult.TotalRoundMonsterRTP[3*(i-1)+k])
		}
	}

	fmt.Println()

	fmt.Println("各關得分佔比(一般子彈不包含觸發擊殺)：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", totalresult.EachRoundScorePercentage[3*(i-1)+k])
		}
	}
	fmt.Println()

	var totalrewardtimes int
	for i := 0; i < 3; i++ {
		totalrewardtimes += totalresult.TotalRewardTimes[i]
	}
	for i := 0; i < 3; i++ {
		totalresult.TotalRewardHitRate[i] = float64(totalresult.TotalRewardTimes[i]) / float64(totalrewardtimes)
	}
	fmt.Println("觸發擊殺獎勵頻率：")
	fmt.Println(totalresult.TotalRewardHitRate)
	fmt.Println()
	fmt.Println("各關觸發擊殺獎勵次數：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", totalresult.TotalLRoundRewardTimes[3*(i-1)+k])
		}
	}

	fmt.Println()

	fmt.Println("各關花費佔比：")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", totalresult.EachRoundCostPercentage[3*(i-1)+k])
		}
	}
	fmt.Println()
	fmt.Println("各關觸發擊殺獎勵得分：1.血量補分 2.得分加乘 3.自爆")
	for i := 1; i < info.GameLevel; i++ {
		for k := 1; k < 4; k++ {
			fmt.Println("關卡", i, "-", k, ":", totalresult.TotalRoundRewardScore[0][3*(i-1)+k], totalresult.TotalRoundRewardScore[1][3*(i-1)+k], totalresult.TotalRoundRewardScore[2][3*(i-1)+k])
		}
	}
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(totalresult.TotalRoundRewardScore[i])
	// }
	fmt.Println()

	fmt.Println("各關卡觸發擊殺獎勵RTP：")
	for i := 0; i < 3; i++ {
		fmt.Println(totalresult.TotalRoundRewardRTP)
	}
	fmt.Println()

	fmt.Println("各等級觸發擊殺獎勵得分：")
	for i := 0; i < info.RewardLevel; i++ {
		fmt.Println("等級", i+1, ":", totalresult.TotalLevelRewardScore[0][i], totalresult.TotalLevelRewardScore[1][i], totalresult.TotalLevelRewardScore[2][i])
	}
	fmt.Println()

	fmt.Println("各等級觸發擊殺獎勵RTP(佔總得分)：1.血量補分 2.得分加乘 3.自爆")
	for i := 0; i < info.RewardLevel; i++ {
		fmt.Println("等級", i+1, ":", totalresult.TotalLevelRewardRTP[0][i], totalresult.TotalLevelRewardRTP[1][i], totalresult.TotalLevelRewardRTP[2][i])
	}

	fmt.Println()

	fmt.Println("玩家關卡分佈：", totalresult.TotalGameEndLevel)
	fmt.Println("玩家關卡分佈比率（結束）：", totalresult.TotalGameEndLevelRate)
	fmt.Println("進入關卡分佈：", totalresult.TotalGameEnterLevel)
	fmt.Println("進入關卡分佈比率：", totalresult.TotalGameEnterLevelRate)
	fmt.Println()

	fmt.Println("一般子彈總得分：", totalresult.TotalMainGameScore)
	fmt.Println("一般子彈ＲＴＰ：", totalresult.MainGameRTP)
	fmt.Println()
	fmt.Println()
	fmt.Println("觸發擊殺獎勵得分：", totalresult.TotalRewardScore)
	fmt.Println("觸發擊殺獎勵RTP：", totalresult.TotalRewardRTP)
	fmt.Println()

	fmt.Println("score:", totalresult.TotalScore)
	fmt.Println("cost:", totalresult.TotalCost)
	rtp := float32(totalresult.TotalScore) / float32(totalresult.TotalCost)
	fmt.Println("RTP:", rtp)
	fmt.Println(time.Since(s))

	WriteCsv(totalresult)
}

func WriteCsv(result TotalResult) {
	// 不存在則建立;存在則清空;讀寫模式;
	file, err := os.Create("towerdata.csv")
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延遲關閉
	defer file.Close()

	// 寫入UTF-8 BOM，防止中文亂碼
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)

	title := []string{"關卡", "關卡1--1", "關卡1--2", "關卡1--3", "關卡2--1", "關卡2--2", "關卡2--3", "關卡3--1", "關卡3--2", "關卡3--3", "關卡4--1", "關卡4--2", "關卡4--3", "關卡5--1", "關卡5--2", "關卡5--3", "關卡6--1", "關卡6--2", "關卡6--3", "關卡7--1", "關卡7--2", "關卡7--3", "關卡8--1", "關卡8--2", "關卡8--3", "關卡9--1", "關卡9--2", "關卡9--3", "關卡10--1", "關卡10--2", "關卡10--3"}
	w.Write(title)
	w.Flush()

	str10 := []string{"各關累積ＲＴＰ(含過關獎勵）"}
	for i := 1; i < gameround; i++ {
		str10 = append(str10, strconv.FormatFloat(float64(result.TotalRoundAccumulateRTP[i]), 'E', 6, 32))
	}
	w.Write(str10)
	w.Flush()

	str := []string{"各關ＲＴＰ(不含過關獎勵）"}
	for i := 1; i < gameround; i++ {
		str = append(str, strconv.FormatFloat(float64(result.TotalRoundRTP[i]), 'E', 6, 32))
	}
	w.Write(str)
	w.Flush()

	str5 := []string{"各關一般子彈RTP："}
	for i := 1; i < gameround; i++ {
		str5 = append(str5, strconv.FormatFloat(float64(result.TotalRoundMainGameRTP[i]), 'E', 6, 32))
	}
	w.Write(str5)
	w.Flush()

	str1 := []string{"玩家關卡分佈比率（結束）："}
	for i := 1; i < gameround; i++ {
		str1 = append(str1, strconv.FormatFloat(float64(result.TotalGameEndLevelRate[i]), 'E', 6, 32))
	}
	w.Write(str1)
	w.Flush()

	str2 := []string{"進入關卡分佈比率："}
	for i := 1; i < gameround; i++ {
		str2 = append(str2, strconv.FormatFloat(float64(result.TotalGameEnterLevelRate[i]), 'E', 6, 32))
	}
	w.Write(str2)
	w.Flush()

	var spacestring []string
	w.Write(spacestring)
	w.Flush()

	// strreward := []string{"關卡", "1", "2", "3", "4", "5,", "6", "7", "8", "9", "11"}
	// w.Write(strreward)
	// w.Flush()

}

func Point() {
	normalpanel0 := &monsterkillrate
	normalpanel1 := &GamePanel

	normalpanel4 := &GameEndLevel

	*normalpanel0, *normalpanel1 = table.Monsterkillrate, table.GamePanel
	*normalpanel4 = make([]int, len(GamePanel))

}

func PlayProcess() PlayerEachRoundResult {
	var playermoney float64
	playermoney = info.PlayerInitialMoney

	var (
		result     PlayerEachRoundResult
		playercost = ecahroundplayercost //玩家下注金額

		playerscore  float64               //玩家得分
		eachroundpay int                   //玩家花費累積金額
		playerlife   = ecahroundplayerlife //玩家血量

		level int //關卡號

		playerAccumulateKill int //玩家累積擊殺怪物

		playerTargetKillLevel int //玩家擊殺怪物目標等級

		playerTargetKill int //玩家擊殺怪物目標數量

		//playerEnterAccumulateKill int //紀錄玩家進入該關卡累積擊殺怪物
	)

	//獎勵狀態///
	var (
		rewardstatus string

		multiplenum int
	)

	//擊殺目標數//
	playerTargetKill = table.KillTarget[playerTargetKillLevel]

	type PlayerEnterInfo struct {
		AccumulateKill  int
		TargetKillLevel int
		TargetKill      int
		round           int
	}
	var enterinfo PlayerEnterInfo

	gamecontinue := true //每次發射判斷是否繼續遊戲
	nextround := true    //判斷是否進下一局盤面
	//var gamestatus string = "Stop"

	/// 產生盤面 ///
	slice := rng.ProductPanel()

	lifeslice := rng.ProductPanelLife(slice)

	//紀錄原始盤面

	initslice := make([][]int, len(slice))
	initlifeslice := make([][]int, len(lifeslice))

	// initslice = append(slice)
	// initlifeslice = append(lifeslice)

	// fmt.Println(slice)
	// fmt.Println(initslice)
	for i := 0; i < len(slice); i++ {
		for k := 0; k < len(slice[i]); k++ {
			initslice[i] = append(initslice[i], slice[i][k])
			initlifeslice[i] = append(initlifeslice[i], lifeslice[i][k])
		}
	}
	// fmt.Println("開始前盤面", slice)
	// fmt.Println("儲存陣列", initslice)

	if printcontrol == true {
		fmt.Println("遊戲盤面", slice)
		fmt.Println("怪物血量盤面：", lifeslice)
	}

	for i := 0; i < len(slice); i++ {
		for k := 0; k < len(slice[i]); k++ {
			result.EachRoundMonster[i][slice[i][k]]++
		}
	}

	// for i := 0; i < gameround; i++ {
	// 	for k := 0; k < len(bonusslice[i]); k++ {
	// 		if bonusslice[i][k] == 1 {
	// 			result.FreeTimes[i]++
	// 		}
	// 	}
	// }

	var bullet int
	var revivaltimes int

	//遊戲流程//
	for round := 1; round < info.Gameround; round++ {
		level = ((round - 1) / 3) + 1

		if info.GameContinueControl == true {
			if (round-1)%3+1 == 1 {

				enterinfo.AccumulateKill = playerAccumulateKill

				if info.GameContinueResetAccumulateKill == true {
					enterinfo.AccumulateKill = 0
				}

				enterinfo.TargetKillLevel = playerTargetKillLevel
				enterinfo.TargetKill = playerTargetKill
				enterinfo.round = round

				if printcontrol == true {
					fmt.Println("玩家進入資訊：", "關卡：", ((enterinfo.round-1)/3)+1, "-", (enterinfo.round-1)%3+1, "累積擊殺：", enterinfo.AccumulateKill, "擊殺獎勵等級：", enterinfo.TargetKillLevel, "擊殺獎勵目標數目：", enterinfo.TargetKill)
				}
			}
		}

		if nextround == true {

			result.EachGameEnterLevel[round]++
			gamecontinue = true

			startposition := 0
			//fmt.Println(slice[round])

			normalpanel := slice[round][startposition : startposition+info.MonsterAmountInPanel]
			//fmt.Println(normalpanel)
			lifepanel := lifeslice[round][startposition : startposition+info.MonsterAmountInPanel]

			var moveset = 1

			for gamecontinue == true {

				if round == 1 {
					//	fmt.Println(normalpanel)
				}

				// fmt.Println("初始盤面", normalpanel)
				// time.Sleep(5 * time.Second)

				if printcontrol == true {
					fmt.Println("第", level, "關")
					fmt.Println("round:", (round-1)%3+1)

					fmt.Println("初始盤面", normalpanel)
					fmt.Println(slice[round])
					fmt.Println(initslice[round])
					fmt.Println("初始血量盤面：", lifepanel)

				}

				for position, monster := range normalpanel {

					///抓取場上怪的位置///
					if monster != 0 {
						///每次擊發子彈加上玩家花費

						eachroundpay += playercost
						result.EachRoundCost[round] += playercost

						playermoney -= float64(playercost)

						//發數+1
						bullet++
						if multiplenum >= 0 {
							multiplenum--
							if printcontrol == true {
								fmt.Println("玩家剩餘加成子彈：", multiplenum)
							}
						}

						result.EachRoundMonsterCost[round][monster] += playercost

						if printcontrol == true {
							fmt.Println("第", bullet, "發")
							fmt.Println("位置", position)
							fmt.Println("怪獸", monster)
						}

						levelofmonsterarray := 10*(level-1) + monster

						if round%3 == 0 && lifepanel[position] > 1 {
							monsterkillrate = table.BossLifekillrate
							monsterodds = table.BossOdds
							oddsweight = table.BossOddsWeight
						} else {
							monsterkillrate = table.Monsterkillrate
							monsterodds = table.MonsterOdds
							oddsweight = table.OddsWeight
						}

						///各關卡怪物擊殺率
						seed := rand.Float64() ///擊殺率///

						// if multiplenum != 0 {
						// 	seed = 0
						// }

						if seed < monsterkillrate[level][monster] {

							//怪物血量減一//
							lifepanel[position]--

							//怪物位置消除//
							if lifepanel[position] == 0 {
								normalpanel[position] = 0
							}

							if printcontrol == true {
								fmt.Println("kill")
								fmt.Println("擊殺盤面", normalpanel)
								fmt.Println("擊殺血量盤面：", lifepanel)

								fmt.Println("怪物", monster)
								fmt.Println("擊殺率：", monsterkillrate)
								fmt.Println("賠率權重", oddsweight[levelofmonsterarray])
								fmt.Println("賠率", monsterodds[levelofmonsterarray])
								fmt.Println("怪物陣列排序:", levelofmonsterarray)

								//fmt.Println("累積擊殺怪物：", playerAccumulateKill)
							}

							monsteroddsseed := rand.Intn(oddsweight[levelofmonsterarray][len(oddsweight[levelofmonsterarray])-1])
							//fmt.Println("賠率種子", monsteroddsseed)

							for i := 0; i < len(oddsweight[levelofmonsterarray])-1; i++ {
								if monsteroddsseed >= oddsweight[levelofmonsterarray][i] && monsteroddsseed < oddsweight[levelofmonsterarray][i+1] {

									pay := monsterodds[levelofmonsterarray][i+1] * float64(playercost)

									//得分加乘效果//

									//得分加成子彈數
									var multiplepay float64
									if multiplenum >= 0 {
										pay *= 1 + info.MutiplePercent
										multiplepay = pay * info.MutiplePercent

										//統計各等級加成得分
										result.EachLevelRewardScore[1][playerTargetKillLevel] += multiplepay
										//統計各關加成得分
										result.EachRoundRewardScore[1][round] += multiplepay

										if printcontrol == true {
											fmt.Println("加成後分數：", pay, "剩餘得分加成子彈：", multiplenum)
										}
									}

									result.EachRoundMonsterScore[round][monster] += pay - multiplepay
									result.EachRoundMainScore[round] += pay - multiplepay

									result.EachRoundScore[round] += pay

									if printcontrol == true {
										fmt.Println("得分", pay)
									}
									playermoney += pay
									playerscore += pay

									if pay >= 1000 {
										//fmt.Println("rtp", pay, float32(playerscore)/float32(eachroundpay))
									}

								}
							}

							//累積擊殺怪物
							playerAccumulateKill++

							if printcontrol == true {
								fmt.Println("累積擊殺怪物數：", playerAccumulateKill, "/", playerTargetKill)
							}
							if playerAccumulateKill == playerTargetKill {

								playerTargetKillLevel++
								playerTargetKill = table.KillTarget[playerTargetKillLevel]
								playerAccumulateKill = 0
								rewardstatus = rng.ProductReward(playerTargetKillLevel - 1)

								if multiplenum >= 0 {
									//fmt.Println(table.KillTarget)
									//fmt.Println("還有加成子彈：", multiplenum, playerTargetKillLevel+1, rewardstatus, "復活次數：", revivaltimes, "關卡：", round)
								}
								if printcontrol == true {
									fmt.Println("目標等級", playerTargetKillLevel, "完成")
									fmt.Println("觸發獎勵：", rewardstatus)
									fmt.Println("下一等級目標：", playerTargetKill)
								}

								// if playerTargetKill == 90 {
								// 	fmt.Println(playerTargetKill)
								// 	fmt.Println(result.EachLevelRewardScore)
								// }
								//rewardstatus = "Bomb"
								//獎勵判斷 Life Multiple Bomb//
								switch rewardstatus {
								case "Life":
									result.EachRoundRewardTimes[round][0]++
									result.EachRewardTimes[0]++
									if playerlife == 6 {
										if printcontrol == true {
											fmt.Println("滿血補分")
										}
										lifepay := float64(info.FullLifeScoreMultiple * playercost)
										playermoney += lifepay
										playerscore += lifepay

										//統計各關得分
										result.EachRoundScore[round] += lifepay
										//統計各等級血量得分
										result.EachLevelRewardScore[0][playerTargetKillLevel] += lifepay
										//統計各關血量得分
										result.EachRoundRewardScore[0][round] += lifepay
									} else {
										playerlife++
									}
								case "Multiple":
									result.EachRoundRewardTimes[round][1]++
									result.EachRewardTimes[1]++

									multiplenum = info.MultipleBulletNum
								case "Bomb":
									result.EachRoundRewardTimes[round][2]++
									result.EachRewardTimes[2]++

									var InputInfo BombInput
									InputInfo.normalpanel = normalpanel
									InputInfo.lifepanel = lifepanel
									InputInfo.playercost = playercost
									InputInfo.levelofmonsterarray = levelofmonsterarray
									InputInfo.level = level
									InputInfo.monsterkillrate = monsterkillrate
									InputInfo.monsterodds = monsterodds
									InputInfo.oddsweight = oddsweight

									//自爆盤面
									BombResult := BombReward(InputInfo)

									normalpanel = BombResult.normalpanel
									lifepanel = BombResult.lifepanel
									playerAccumulateKill = BombResult.playerAccumulateKill

									//統計各關得分
									result.EachRoundScore[round] += BombResult.pay
									//統計各等級炸彈得分
									result.EachLevelRewardScore[2][playerTargetKillLevel] += BombResult.pay
									//統計各關炸彈得分
									result.EachRoundRewardScore[2][round] += BombResult.pay

									playermoney += BombResult.pay
									playerscore += BombResult.pay
								}

							}

						} else {
							if printcontrol == true {
								fmt.Println("no kill")
								fmt.Println("擊殺率：", monsterkillrate)
							}

						}

						break ///只找第一隻怪
					}

				}

				if playermoney == 0 {
					if info.PlyerMoneyControl == true {
						playerlife = 0
					}
				}

				///若盤面第0隻還有怪則玩家血量減一///
				if moveset%MoveSet == 0 {
					if normalpanel[0] != 0 {
						playerlife--

						if printcontrol == true {
							fmt.Println("剩餘血量：", playerlife)
						}
					}
				}

				if playerlife == 0 {
					if printcontrol == true {
						fmt.Println("玩家死亡")
						fmt.Println("關卡", result.EachGameEndLevel)
					}
					if info.GameContinueControl != true {
						gamecontinue = false
						nextround = false
						result.EachGameEndLevel = round
						break
					} else {
						revivaltimes++

						if revivaltimes < 0 {
							nextround = false
							break
						}
						if revivaltimes%10000 == 0 {
							fmt.Println("復活", revivaltimes, "次", "level:", level)
						}

						eachroundpay += playercost * info.ContinuePayMultiple
						result.EachRoundCost[round] += playercost * info.ContinuePayMultiple

						playerlife = 6
						moveset = 1
						startposition = 0
						round = enterinfo.round

						slice = make([][]int, len(initslice))
						lifeslice = make([][]int, len(initlifeslice))
						for i := 0; i < len(initslice); i++ {
							for k := 0; k < len(initslice[i]); k++ {
								slice[i] = append(slice[i], initslice[i][k])
								lifeslice[i] = append(lifeslice[i], initlifeslice[i][k])
							}
						}
						//normalpanel = slice[round][startposition : startposition+info.MonsterAmountInPanel]
						//lifepanel = lifeslice[round][startposition : startposition+info.MonsterAmountInPanel]

						if printcontrol == true {
							fmt.Println("盤面比較：")
							fmt.Println(slice[round])
							fmt.Println(initslice[round])
						}

						playerAccumulateKill = enterinfo.AccumulateKill

						playerTargetKillLevel = enterinfo.TargetKillLevel
						playerTargetKill = table.KillTarget[playerTargetKillLevel]

						if printcontrol == true {
							fmt.Println("玩家復活")
							fmt.Println("從第", round, "round開始")
							fmt.Println("盤面：", slice)
							fmt.Println("玩家進入資訊：", "關卡：", ((enterinfo.round-1)/3)+1, "-", (enterinfo.round-1)%3+1, "累積擊殺：", playerAccumulateKill, "擊殺獎勵等級：", playerTargetKillLevel, "擊殺獎勵目標數目：", playerTargetKill)

							fmt.Println("nextround", nextround)
							fmt.Println("玩家剩餘金額：", playermoney)
							fmt.Println("玩家累積得分", playerscore)
							fmt.Println("玩家累積花費金額", eachroundpay)
							fmt.Println("剩餘血量：", playerlife)
							fmt.Println()

							fmt.Println()
						}

						//break
					}

				} else {
					result.EachGameEndLevel = round
				}

				///重整盤面///
				//fmt.Println("盤面：", slice)
				if moveset%MoveSet == 0 {
					startposition++

					if startposition+info.MonsterAmountInPanel < len(slice[round]) {
						normalpanel = slice[round][startposition : startposition+info.MonsterAmountInPanel]
						lifepanel = lifeslice[round][startposition : startposition+info.MonsterAmountInPanel]
					} else {
						normalpanel = slice[round][startposition:len(slice[round])]
						lifepanel = lifeslice[round][startposition:len(slice[round])]

					}
				}

				///盤面上若無怪則進入下一局
				var monsteramout int
				for _, k := range normalpanel {
					if k != 0 {
						monsteramout++
					}
				}
				if startposition+info.MonsterAmountInPanel <= len(slice[round]) {
					for k := startposition + info.MonsterAmountInPanel; k < len(slice[round]); k++ {

						if slice[round][k] != 0 {
							monsteramout++
						}
					}

				}
				if monsteramout == 0 {

					///過關獎勵//
					if round%3 == 0 {

						if printcontrol == true {
							fmt.Println("玩家過關：", "Level:", round/3)

							if round == 15 {
								fmt.Println("全數通關")
								fmt.Println(playerAccumulateKill)
								fmt.Println("下一等級目標：", playerTargetKill)
								fmt.Println()
							}
						}
					}

					break
				} else {

					gamecontinue = true
				}
				if printcontrol == true {
					fmt.Println("nextround", nextround)
					fmt.Println("玩家剩餘金額：", playermoney)
					fmt.Println("玩家累積得分", playerscore)
					fmt.Println("玩家累積花費金額", eachroundpay)
					fmt.Println("剩餘血量：", playerlife)
					fmt.Println()
				}

				moveset++

			}

		}

		// if nextround == false {
		// 	break
		// }
	}
	result.PlayerCost = eachroundpay
	result.PlayerScore = playerscore

	if printcontrol == true {
		fmt.Println("玩家最後剩餘金額：", playermoney)
		fmt.Println("玩家累積得分", playerscore)
		fmt.Println("玩家最後花費金額", eachroundpay)
		fmt.Println("總復活次數：", revivaltimes)
	}
	return result
}

type BombInput struct {
	normalpanel         []int
	lifepanel           []int
	levelofmonsterarray int
	level               int
	monsterkillrate     [][]float64
	monsterodds         [][]float64
	oddsweight          [][]int

	playercost int
}
type BombOutput struct {
	normalpanel          []int
	lifepanel            []int
	playerAccumulateKill int
	pay                  float64
}

func BombReward(Input BombInput) BombOutput {
	var Output BombOutput

	//fmt.Println(Input.normalpanel)
	for bombtimes := 0; bombtimes < info.BombTimes; bombtimes++ {
		if info.BombPrintcontrol == true {
			fmt.Println("轟炸第", bombtimes, "次")
			fmt.Println("盤面：", Input.normalpanel)
		}
		for position, monster := range Input.normalpanel {

			///抓取場上怪的位置///
			if monster != 0 {
				if bombprintcontrol == true {

					fmt.Println("位置", position)
					fmt.Println("怪獸", monster)
				}

				///各關卡怪物擊殺率
				seed := rand.Float64() ///擊殺率///

				if seed < Input.monsterkillrate[Input.level][monster] {

					//怪物血量減一//
					Input.lifepanel[position]--

					//怪物位置消除//
					if Input.lifepanel[position] == 0 {
						Input.normalpanel[position] = 0
					}

					if bombprintcontrol == true {
						fmt.Println("kill")
						fmt.Println("擊殺盤面", Input.normalpanel)
						fmt.Println("擊殺血量盤面：", Input.lifepanel)

						fmt.Println("怪物", monster)
						fmt.Println("擊殺率：", Input.monsterkillrate)
						fmt.Println("賠率權重", Input.oddsweight[Input.levelofmonsterarray])
						fmt.Println("賠率", Input.monsterodds[Input.levelofmonsterarray])
						fmt.Println("怪物陣列排序:", Input.levelofmonsterarray)

						fmt.Println("累積擊殺怪物：", Output.playerAccumulateKill)
					}

					monsteroddsseed := rand.Intn(Input.oddsweight[Input.levelofmonsterarray][len(Input.oddsweight[Input.levelofmonsterarray])-1])
					//fmt.Println("賠率種子", monsteroddsseed)
					for i := 0; i < len(Input.oddsweight[Input.levelofmonsterarray])-1; i++ {
						if monsteroddsseed >= Input.oddsweight[Input.levelofmonsterarray][i] && monsteroddsseed < Input.oddsweight[Input.levelofmonsterarray][i+1] {
							Output.pay += monsterodds[Input.levelofmonsterarray][i+1] * float64(Input.playercost)

							if bombprintcontrol == true {
								fmt.Println("得分", Output.pay)
							}

						}
					}

					//累積擊殺怪物
					Output.playerAccumulateKill++

				} else {
					if bombprintcontrol == true {
						fmt.Println("no kill")
						fmt.Println("擊殺率：", monsterkillrate)
					}
				}

			}

		}
	}
	//fmt.Println("加上自爆初始得分前：", Output.pay)
	Output.pay += float64(Input.playercost * info.BombScoreMultiple)
	//fmt.Println("加上自爆初始得分後：", Output.pay)

	Output.lifepanel = Input.lifepanel
	Output.normalpanel = Input.normalpanel

	return Output
}

func ChangeTable(RTP int) {
	switch RTP {
	case 95:
		monsterkillrate = table.Monsterkillrate95
		oddsweight = table.OddsWeight95
	case 99:
		monsterkillrate = table.Monsterkillrate99
		oddsweight = table.OddsWeight99
	case 965:
		monsterkillrate = table.Monsterkillrate
		oddsweight = table.OddsWeight
	}

}
