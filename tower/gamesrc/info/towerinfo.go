package tower

const (
	Printcontrol = false

	//一秒四發
	MoveSet                  = 6
	Gameround            int = 16
	GameLevel            int = 6 //五關＋1
	Freebulletset        int = 0 //設定值為發數+1
	FreebulletMultiple   int = 2
	FreebulletDeadSet        = false //控制免費子彈是否直接擊殺
	Stoptimesset         int = 0
	BossLife                 = 1
	PlayerCost               = 10
	PlayerLife               = 6
	Session                  = 1000000
	MonsterAmountInPanel     = 14

	PlyerMoneyControl  = false //控制玩家是否花錢遊玩
	PlayerInitialMoney = 5000

	GameContinueControl             = false //控制是否開啟玩家復活
	GameContinueResetAccumulateKill = false //控制是否開啟玩家復活歸零累積擊殺數
	ContinuePayMultiple             = 5     //玩家復活花費倍數

	TargetReward              = false //控制是否開啟全部通關獎勵
	RewardLevel           int = 13
	MutiplePercent            = 0.2 //得分加乘趴數//
	MultipleBulletNum         = 60  // 得分加成發數//
	BombPrintcontrol          = false
	BombTimes                 = 3 //自爆射擊次數
	BombScoreMultiple         = 5 //自爆基本分數倍數
	FullLifeScoreMultiple     = 5 //滿血補分基本倍數

	AllPassRewardMultiple = 200

	RiskControl = true
)

var (
	Risk = [5][3]int{
		{500000, 482500, 0},
		{400000, 386000, 0},
		{400000, 386000, 0},
		{300000, 289500, 0},
		{300000, 289500, 0},
	}
)
