package biz

import (
	"context"
	v1 "dhb/app/app/api"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID        int64
	Address   string
	Undo      int64
	CreatedAt time.Time
}

type UserInfo struct {
	ID               int64
	UserId           int64
	Vip              int64
	HistoryRecommend int64
}

type UserRecommend struct {
	ID            int64
	UserId        int64
	RecommendCode string
	CreatedAt     time.Time
}

type UserRecommendArea struct {
	ID            int64
	RecommendCode string
	Num           int64
	CreatedAt     time.Time
}

type UserArea struct {
	ID         int64
	UserId     int64
	Amount     int64
	SelfAmount int64
	Level      int64
}

type UserCurrentMonthRecommend struct {
	ID              int64
	UserId          int64
	RecommendUserId int64
	Date            time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type UserBalance struct {
	ID          int64
	UserId      int64
	BalanceUsdt int64
	BalanceDhb  int64
}

type Withdraw struct {
	ID              int64
	UserId          int64
	Amount          int64
	RelAmount       int64
	BalanceRecordId int64
	Status          string
	Type            string
	CreatedAt       time.Time
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

type UserUseCase struct {
	repo                          UserRepo
	urRepo                        UserRecommendRepo
	configRepo                    ConfigRepo
	uiRepo                        UserInfoRepo
	ubRepo                        UserBalanceRepo
	locationRepo                  LocationRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type LocationNew struct {
	ID                int64
	UserId            int64
	Status            string
	Current           int64
	CurrentMax        int64
	StopLocationAgain int64
	OutRate           int64
	StopCoin          int64
	StopDate          time.Time
	CreatedAt         time.Time
}

type BalanceReward struct {
	ID        int64
	UserId    int64
	Status    int64
	Amount    int64
	SetDate   time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Reward struct {
	ID               int64
	UserId           int64
	Amount           int64
	BalanceRecordId  int64
	Type             string
	TypeRecordId     int64
	Reason           string
	ReasonLocationId int64
	LocationType     string
	CreatedAt        time.Time
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type ConfigRepo interface {
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetConfigs(ctx context.Context) ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
}

type UserBalanceRepo interface {
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error
	SystemReward(ctx context.Context, amount int64, locationId int64) error
	SystemFee(ctx context.Context, amount int64, locationId int64) error
	GetSystemYesterdayDailyReward(ctx context.Context) (*Reward, error)
	UserFee(ctx context.Context, userId int64, amount int64) (int64, error)
	RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	Deposit(ctx context.Context, userId int64, amount int64) (int64, error)
	DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error)
	DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetUserRewardByUserId(ctx context.Context, userId int64) ([]*Reward, error)
	GetUserRewardByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserSortRecommendReward, error)
	GetUserRewards(ctx context.Context, b *Pagination, userId int64) ([]*Reward, error, int64)
	GetUserRewardsLastMonthFee(ctx context.Context) ([]*Reward, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceUsdtTotal(ctx context.Context) (int64, error)
	GreateWithdraw(ctx context.Context, userId int64, amount int64, coinType string) (*Withdraw, error)
	WithdrawUsdt(ctx context.Context, userId int64, amount int64) error
	WithdrawDhb(ctx context.Context, userId int64, amount int64) error
	GetWithdrawByUserId(ctx context.Context, userId int64) ([]*Withdraw, error)
	GetWithdraws(ctx context.Context, b *Pagination, userId int64) ([]*Withdraw, error, int64)
	GetWithdrawPassOrRewarded(ctx context.Context) ([]*Withdraw, error)
	UpdateWithdraw(ctx context.Context, id int64, status string) (*Withdraw, error)
	GetWithdrawById(ctx context.Context, id int64) (*Withdraw, error)
	GetWithdrawNotDeal(ctx context.Context) ([]*Withdraw, error)
	GetUserBalanceRecordUserUsdtTotal(ctx context.Context, userId int64) (int64, error)
	GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error)
	GetUserRewardUsdtTotal(ctx context.Context) (int64, error)
	GetSystemRewardUsdtTotal(ctx context.Context) (int64, error)
	UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*Withdraw, error)
	GetUserRewardRecommendSort(ctx context.Context) ([]*UserSortRecommendReward, error)
	GetUserRewardTodayTotalByUserId(ctx context.Context, userId int64) (*UserSortRecommendReward, error)

	SetBalanceReward(ctx context.Context, userId int64, amount int64) error
	UpdateBalanceReward(ctx context.Context, userId int64, id int64, amount int64, status int64) error
	GetBalanceRewardByUserId(ctx context.Context, userId int64) ([]*BalanceReward, error)
}

type UserRecommendRepo interface {
	GetUserRecommendByUserId(ctx context.Context, userId int64) (*UserRecommend, error)
	CreateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (*UserRecommend, error)
	UpdateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	CreateUserRecommendArea(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	DeleteOrOriginUserRecommendArea(ctx context.Context, code string, originCode string) (bool, error)
	GetUserRecommendLowArea(ctx context.Context, code string) ([]*UserRecommendArea, error)
	GetUserAreas(ctx context.Context, userIds []int64) ([]*UserArea, error)
	CreateUserArea(ctx context.Context, u *User) (bool, error)
	GetUserArea(ctx context.Context, userId int64) (*UserArea, error)
}

type UserCurrentMonthRecommendRepo interface {
	GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *Pagination, userId int64) ([]*UserCurrentMonthRecommend, error, int64)
	CreateUserCurrentMonthRecommend(ctx context.Context, u *UserCurrentMonthRecommend) (*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error)
	GetUserLastMonthRecommend(ctx context.Context) ([]int64, error)
}

type UserInfoRepo interface {
	CreateUserInfo(ctx context.Context, u *User) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, u *UserInfo) (*UserInfo, error)
	GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserInfo, error)
}

type UserRepo interface {
	GetUserById(ctx context.Context, Id int64) (*User, error)
	GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetUsers(ctx context.Context, b *Pagination, address string) ([]*User, error, int64)
	GetUserCount(ctx context.Context) (int64, error)
	GetUserCountToday(ctx context.Context) (int64, error)
}

func NewUserUseCase(repo UserRepo, tx Transaction, configRepo ConfigRepo, uiRepo UserInfoRepo, urRepo UserRecommendRepo, locationRepo LocationRepo, userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                          repo,
		tx:                            tx,
		configRepo:                    configRepo,
		locationRepo:                  locationRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		uiRepo:                        uiRepo,
		urRepo:                        urRepo,
		ubRepo:                        ubRepo,
		log:                           log.NewHelper(logger),
	}
}

func (uuc *UserUseCase) GetUserByAddress(ctx context.Context, Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(ctx, Addresses...)
}

func (uuc *UserUseCase) GetDhbConfig(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "level1Dhb", "level2Dhb", "level3Dhb")
}

func (uuc *UserUseCase) GetExistUserByAddressOrCreate(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user          *User
		recommendUser *UserRecommend
		userRecommend *UserRecommend
		userInfo      *UserInfo
		userBalance   *UserBalance
		err           error
		userId        int64
		decodeBytes   []byte
	)

	user, err = uuc.repo.GetUserByAddress(ctx, u.Address) // 查询用户
	if nil == user || nil != err {
		code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
		if "abf00dd52c08a9213f225827bc3fb100" != code {
			decodeBytes, err = base64.StdEncoding.DecodeString(code)
			code = string(decodeBytes)
			if 1 >= len(code) {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}
			if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}

			// 查询推荐人的相关信息
			recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
			if err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = uuc.repo.CreateUser(ctx, u) // 用户创建
			if err != nil {
				return err
			}

			userInfo, err = uuc.uiRepo.CreateUserInfo(ctx, user) // 创建用户信息
			if err != nil {
				return err
			}

			userRecommend, err = uuc.urRepo.CreateUserRecommend(ctx, user, recommendUser) // 创建用户推荐信息
			if err != nil {
				return err
			}

			_, err = uuc.urRepo.CreateUserArea(ctx, user)
			if err != nil {
				return err
			}

			userBalance, err = uuc.ubRepo.CreateUserBalance(ctx, user) // 创建余额信息
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (uuc *UserUseCase) UpdateUserRecommend(ctx context.Context, u *User, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	var (
		err                   error
		userId                int64
		recommendUser         *UserRecommend
		userRecommend         *UserRecommend
		locations             []*LocationNew
		myRecommendUser       *User
		myUserRecommendUserId int64
		decodeBytes           []byte
	)

	code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
	if "abf00dd52c08a9213f225827bc3fb100" != code {
		decodeBytes, err = base64.StdEncoding.DecodeString(code)
		code = string(decodeBytes)
		if 1 >= len(code) {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}
		if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}

		// 现有推荐人信息，判断推荐人是否改变
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, u.ID)
		if nil == userRecommend {
			return nil, err
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
			if nil != err {
				return nil, err
			}
		}
		if myRecommendUser.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, err
		}

		// 我的占位信息
		locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, u.ID)
		if nil != locations && 0 < len(locations) {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 查询推荐人的相关信息
		recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}

		// 推荐人信息
		myRecommendUser, err = uuc.repo.GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		// 更新
		_, err = uuc.urRepo.UpdateUserRecommend(ctx, u, recommendUser)
		if err != nil {
			return nil, err
		}
	}

	return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, err
}

func (uuc *UserUseCase) UserInfo(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	var (
		myUser                   *User
		userInfo                 *UserInfo
		locations                []*LocationNew
		userBalance              *UserBalance
		userRecommend            *UserRecommend
		userRecommends           []*UserRecommend
		userRewards              []*Reward
		userRewardTotal          int64
		encodeString             string
		myUserRecommendUserId    int64
		myRecommendUser          *User
		recommendTeamNum         int64
		recommendTotal           int64
		recommendTeamTotal       int64
		locationDailyRewardTotal int64
		recommendAreaTotal       int64
		myCode                   string
		inviteUserAddress        string
		amount                   = "0"
		userCount                string
		status                   = "no"
		configs                  []*Config
		myLastStopLocations      []*LocationNew
		myLastLocationCurrent    int64
		myWithdraws              []*Withdraw
		totalDepoist             int64
		withdrawAmount           int64
		locationCount            int64
		userTodayReward          int64
		recommendTop             int64
		fybPrice                 int64
		areaAmount               int64
		maxAreaAmount            int64
		recommendAreaOne         int64
		recommendAreaTwo         int64
		recommendAreaThree       int64
		recommendAreaFour        int64
		recommendAreaOneName     string
		recommendAreaTwoName     string
		recommendAreaThreeName   string
		recommendAreaFourName    string
		areaName                 string
		timeAgain                int64
		stopCoin                 int64
		err                      error
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx, "user_count", "coin_price", "time_again", "recommend_area_one", "recommend_area_two", "recommend_area_three", "recommend_area_four")
	if nil != configs {
		for _, vConfig := range configs {
			if "user_count" == vConfig.KeyName {
				userCount = vConfig.Value
			}
			if "coin_price" == vConfig.KeyName {
				fybPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "time_again" == vConfig.KeyName {
				timeAgain, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "recommend_area_one" == vConfig.KeyName {
				recommendAreaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
				recommendAreaOneName = vConfig.Name
			}
			if "recommend_area_two" == vConfig.KeyName {
				recommendAreaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
				recommendAreaTwoName = vConfig.Name
			}
			if "recommend_area_three" == vConfig.KeyName {
				recommendAreaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
				recommendAreaThreeName = vConfig.Name
			}
			if "recommend_area_four" == vConfig.KeyName {
				recommendAreaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
				recommendAreaFourName = vConfig.Name
			}
		}
	}

	myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}

	locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, myUser.ID)
	if nil != locations && 0 < len(locations) {
		status = "stop"
		tmpCurrent := int64(0)
		for _, v := range locations {
			if "running" == v.Status {
				status = "yes"
				tmpCurrent += v.Current

			}
		}

		if tmpCurrent > 0 {
			status = "running"
			amount = fmt.Sprintf("%.2f", float64(tmpCurrent)/float64(10000000000))
		}
	}

	locationCount = int64(len(locations))

	// 提现记录
	myWithdraws, err = uuc.ubRepo.GetWithdrawByUserId(ctx, myUser.ID)
	for _, vMyWithdraw := range myWithdraws {
		withdrawAmount += vMyWithdraw.RelAmount
	}

	// 充值记录
	totalDepoist, err = uuc.ubRepo.GetUserBalanceRecordUserUsdtTotal(ctx, myUser.ID)

	// 冻结
	myLastStopLocations, err = uuc.locationRepo.GetMyStopLocationsLast(ctx, myUser.ID)
	now := time.Now().UTC().Add(8 * time.Hour)
	if nil != myLastStopLocations {
		for _, vMyLastStopLocations := range myLastStopLocations {
			if now.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
				myLastLocationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax // 补上
				stopCoin += vMyLastStopLocations.StopCoin
			}
		}
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}

	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	if nil == userRecommend {
		return nil, err
	}
	myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	codeByte := []byte(myCode)
	encodeString = base64.StdEncoding.EncodeToString(codeByte)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}
		myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
		if nil != err {
			return nil, err
		}
		inviteUserAddress = myRecommendUser.Address
		myCode = userRecommend.RecommendCode + myCode
	}

	// 团队
	userRecommends, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode)
	if nil != userRecommends {
		recommendTeamNum = int64(len(userRecommends))
	}

	// 累计奖励
	userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	if nil != userRewards {
		for _, vUserReward := range userRewards {
			userRewardTotal += vUserReward.Amount
			if "recommend_team" == vUserReward.Reason {
				recommendTeamTotal += vUserReward.Amount
			} else if "daily_recommend_area" == vUserReward.Reason {
				recommendAreaTotal += vUserReward.Amount
			} else if "location_daily_reward" == vUserReward.Reason {
				locationDailyRewardTotal += vUserReward.Amount
			} else if "recommend" == vUserReward.Reason {
				recommendTotal += vUserReward.Amount
			}
		}
	}

	// 小区信息
	if "" != myCode {
		var (
			myRecommendUsers   []*UserRecommend
			userAreas          []*UserArea
			myRecommendUserIds []int64
		)
		myRecommendUsers, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
		if nil == err {
			// 找直推
			for _, vMyRecommendUsers := range myRecommendUsers {
				myRecommendUserIds = append(myRecommendUserIds, vMyRecommendUsers.UserId)
			}
		}

		if 0 < len(myRecommendUserIds) {
			userAreas, err = uuc.urRepo.GetUserAreas(ctx, myRecommendUserIds)
			if nil == err {
				var (
					tmpTotalAreaAmount int64
				)
				for _, vUserAreas := range userAreas {
					tmpAreaAmount := vUserAreas.Amount + vUserAreas.SelfAmount
					tmpTotalAreaAmount += tmpAreaAmount
					if tmpAreaAmount > maxAreaAmount {
						maxAreaAmount = tmpAreaAmount
					}
				}

				areaAmount = tmpTotalAreaAmount - maxAreaAmount
			}

			// 比较级别
			if areaAmount >= recommendAreaOne {
				areaName = recommendAreaOneName
			}

			if areaAmount >= recommendAreaTwo {
				areaName = recommendAreaTwoName
			}

			if areaAmount >= recommendAreaThree {
				areaName = recommendAreaThreeName
			}

			if areaAmount >= recommendAreaFour {
				areaName = recommendAreaFourName
			}
		}

	}
	// 优先展示设定的
	var myUserArea *UserArea
	myUserArea, err = uuc.urRepo.GetUserArea(ctx, user.ID)
	if nil != myUserArea && 0 < myUserArea.Level {
		if myUserArea.Level >= 1 {
			areaName = recommendAreaOneName
		}
		if myUserArea.Level >= 2 {
			areaName = recommendAreaTwoName
		}
		if myUserArea.Level >= 3 {
			areaName = recommendAreaThreeName
		}
		if myUserArea.Level >= 4 {
			areaName = recommendAreaFourName
		}
	}

	var (
		balanceRewards           []*BalanceReward
		totalBalanceRewardAmount int64
	)
	balanceRewards, err = uuc.ubRepo.GetBalanceRewardByUserId(ctx, user.ID)
	if nil != balanceRewards {
		for _, vBalanceReward := range balanceRewards {
			totalBalanceRewardAmount += vBalanceReward.Amount
		}
	}

	return &v1.UserInfoReply{
		Address:             myUser.Address,
		Status:              status,
		Amount:              amount,
		BalanceUsdt:         fmt.Sprintf("%.4f", float64(userBalance.BalanceUsdt)/float64(10000000000)),
		BalanceDhb:          fmt.Sprintf("%.4f", float64(userBalance.BalanceDhb)/float64(10000000000)),
		InviteUrl:           encodeString,
		InviteUserAddress:   inviteUserAddress,
		RecommendNum:        userInfo.HistoryRecommend,
		RecommendTeamNum:    recommendTeamNum,
		Total:               fmt.Sprintf("%.4f", float64(userRewardTotal)/float64(10000000000)),
		WithdrawAmount:      fmt.Sprintf("%.3f", float64(withdrawAmount)/float64(10000000000)),
		RecommendTotal:      fmt.Sprintf("%.4f", float64(recommendTotal)/float64(10000000000)),
		Usdt:                "0x55d398326f99059fF775485246999027B3197955",
		Account:             "0x6b2c086C9bDb2e09A85f84CD1b8eed1d9C9B7eae",
		AmountB:             fmt.Sprintf("%.4f", float64(myLastLocationCurrent)/float64(10000000000)),
		AmountC:             fmt.Sprintf("%.4f", float64(stopCoin)/float64(10000000000)),
		UserCount:           userCount,
		TotalDeposit:        fmt.Sprintf("%.4f", float64(totalDepoist)/float64(10000000000)),
		LocationCount:       locationCount,
		TodayReward:         fmt.Sprintf("%.4f", float64(userTodayReward)/float64(10000000000)),
		RecommendTop:        fmt.Sprintf("%.4f", float64(recommendTop)/float64(10000000000)),
		FybPrice:            fmt.Sprintf("%.4f", float64(fybPrice)/float64(1000)),
		Undo:                myUser.Undo,
		AreaName:            areaName,
		AreaAmount:          fmt.Sprintf("%.4f", float64(areaAmount)/float64(100000)),
		AreaMaxAmount:       fmt.Sprintf("%.4f", float64(maxAreaAmount)/float64(100000)),
		RecommendAreaTotal:  fmt.Sprintf("%.4f", float64(recommendAreaTotal)/float64(10000000000)),
		AmountBalanceReward: fmt.Sprintf("%.4f", float64(totalBalanceRewardAmount)/float64(10000000000)),
	}, nil
}

func (uuc *UserUseCase) RewardList(ctx context.Context, req *v1.RewardListRequest, user *User) (*v1.RewardListReply, error) {

	res := &v1.RewardListReply{
		Rewards: make([]*v1.RewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) RecommendRewardList(ctx context.Context, user *User) (*v1.RecommendRewardListReply, error) {

	res := &v1.RecommendRewardListReply{
		Rewards: make([]*v1.RecommendRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) FeeRewardList(ctx context.Context, user *User) (*v1.FeeRewardListReply, error) {
	res := &v1.FeeRewardListReply{
		Rewards: make([]*v1.FeeRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) WithdrawList(ctx context.Context, user *User) (*v1.WithdrawListReply, error) {

	var (
		withdraws []*Withdraw
		err       error
	)

	res := &v1.WithdrawListReply{
		Withdraw: make([]*v1.WithdrawListReply_List, 0),
	}

	withdraws, err = uuc.ubRepo.GetWithdrawByUserId(ctx, user.ID)
	if nil != err {
		return res, err
	}

	for _, v := range withdraws {
		res.Withdraw = append(res.Withdraw, &v1.WithdrawListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.RelAmount)/float64(10000000000)),
			Status:    v.Status,
			Type:      v.Type,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) Withdraw(ctx context.Context, req *v1.WithdrawRequest, user *User) (*v1.WithdrawReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	if "dhb" != req.SendBody.Type && "usdt" != req.SendBody.Type {
		return &v1.WithdrawReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 10000000000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.WithdrawReply{
			Status: "fail",
		}, nil
	}

	if "dhb" == req.SendBody.Type && userBalance.BalanceDhb < amount {
		return &v1.WithdrawReply{
			Status: "fail",
		}, nil
	}

	if "usdt" == req.SendBody.Type && userBalance.BalanceUsdt < amount {
		return &v1.WithdrawReply{
			Status: "fail",
		}, nil
	}
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		if "usdt" == req.SendBody.Type {
			err = uuc.ubRepo.WithdrawUsdt(ctx, user.ID, amount) // 提现
			if nil != err {
				return err
			}
			_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount, req.SendBody.Type)
			if nil != err {
				return err
			}

		} else if "dhb" == req.SendBody.Type {
			err = uuc.ubRepo.WithdrawDhb(ctx, user.ID, amount) // 提现
			if nil != err {
				return err
			}
			_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount, req.SendBody.Type)
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) SetBalanceReward(ctx context.Context, req *v1.SetBalanceRewardRequest, user *User) (*v1.SetBalanceRewardReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 10000000000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.SetBalanceReward(ctx, user.ID, amount) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.SetBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) DeleteBalanceReward(ctx context.Context, req *v1.DeleteBalanceRewardRequest, user *User) (*v1.DeleteBalanceRewardReply, error) {
	var (
		err            error
		balanceRewards []*BalanceReward
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 10000000000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	balanceRewards, err = uuc.ubRepo.GetBalanceRewardByUserId(ctx, user.ID)
	if nil != err {
		return &v1.DeleteBalanceRewardReply{
			Status: "无余额宝数据",
		}, nil
	}

	var totalBalanceRewardAmount int64
	for _, vBalanceReward := range balanceRewards {
		totalBalanceRewardAmount += vBalanceReward.Amount
	}

	if totalBalanceRewardAmount < amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "余额宝金额不足",
		}, nil
	}

	for _, vBalanceReward := range balanceRewards {
		tmpAmount := int64(0)
		Status := int64(1)

		if amount-vBalanceReward.Amount < 0 {
			tmpAmount = amount
		} else {
			tmpAmount = vBalanceReward.Amount
			Status = 2
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.UpdateBalanceReward(ctx, user.ID, vBalanceReward.ID, tmpAmount, Status) // 提现
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
		amount -= tmpAmount

		if amount <= 0 {
			break
		}
	}

	return &v1.DeleteBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	res := &v1.AdminRewardListReply{
		Rewards: make([]*v1.AdminRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {

	res := &v1.AdminUserListReply{
		Users: make([]*v1.AdminUserListReply_UserList, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error) {
	return uuc.repo.GetUserByUserIds(ctx, userIds...)
}

func (uuc *UserUseCase) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}
	return res, nil

}

func (uuc *UserUseCase) AdminRecommendList(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	res := &v1.AdminUserRecommendReply{
		Users: make([]*v1.AdminUserRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {

	res := &v1.AdminMonthRecommendReply{
		Users: make([]*v1.AdminMonthRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	res := &v1.AdminConfigReply{
		Config: make([]*v1.AdminConfigReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	res := &v1.AdminConfigUpdateReply{}
	return res, nil
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedList(ctx context.Context) ([]*Withdraw, error) {
	return uuc.ubRepo.GetWithdrawPassOrRewarded(ctx)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	res := &v1.AdminWithdrawListReply{
		Withdraw: make([]*v1.AdminWithdrawListReply_List, 0),
	}

	return res, nil

}

func (uuc *UserUseCase) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return &v1.AdminFeeReply{}, nil
}

func (uuc *UserUseCase) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {

	return &v1.AdminAllReply{}, nil
}

func (uuc *UserUseCase) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return &v1.AdminWithdrawReply{}, nil
}
