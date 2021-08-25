package tower

import (
	"math/rand"
	info "towerjs/tower/gamesrc/info"
	t "towerjs/tower/gamesrc/table"
)

const BossLife = info.BossLife

func ProductPanel() [][]int {

	panel := make([][]int, len(t.GamePanel))

	for i := 0; i < len(t.GamePanel); i++ {

		slice := make([]int, len(t.GamePanel[i]))
		for k := 0; k < len(t.GamePanel[i]); k++ {
			if t.GamePanel[i][k] != 0 {
				seed := rand.Intn(t.MonsterWeight[i][len(t.MonsterWeight[i])-1])

				for m := 0; m < len(t.MonsterWeight[i])-1; m++ {
					if seed >= t.MonsterWeight[i][m] && seed < t.MonsterWeight[i][m+1] {
						slice[k] = t.MonsterWeight[0][m+1]
						//fmt.Println(t.MonsterWeight[0][m+1])
					}
				}

			} else {
				slice[k] = t.GamePanel[i][k]
			}

		}

		panel[i] = slice
		//fmt.Println(slice)
	}

	//fmt.Println("產盤", panel)
	return panel
}

func ProductPanelLife(panel [][]int) [][]int {

	lifepanel := make([][]int, len(panel))
	for i := 1; i < len(panel); i++ {
		for k := 0; k < len(panel[i]); k++ {
			if panel[i][k] == 0 {
				lifepanel[i] = append(lifepanel[i], 0)
			} else {
				if i%3 == 0 {
					lifepanel[i] = append(lifepanel[i], BossLife)
				} else {
					lifepanel[i] = append(lifepanel[i], 1)
				}

			}
		}
	}

	return lifepanel
}

func ProductReward(level int) string {

	var reward int
	var rewardname string

	//fmt.Println(t.RewardWeight[level])
	seed := rand.Intn(t.RewardWeight[level][len(t.RewardWeight[level])-1])
	//fmt.Println(t.RewardWeight[level])

	for i := 0; i < len(t.RewardWeight[level])-1; i++ {
		//fmt.Println(t.RewardWeight[level][i-1], t.RewardWeight[level][i])

		if seed >= t.RewardWeight[level][i] && seed < t.RewardWeight[level][i+1] {
			reward = i

		}
	}

	switch reward {
	case 0:
		rewardname = "Life"
	case 1:
		rewardname = "Multiple"
	case 2:
		rewardname = "Bomb"
	}
	//fmt.Println(seed, reward, rewardname)
	return rewardname
}
