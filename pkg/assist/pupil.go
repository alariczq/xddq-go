package assist

import (
	"context"
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/utils"
)

type PupilManager struct {
	a *Assistant
}

func NewPupilManager(a *Assistant) *PupilManager {
	pm := &PupilManager{
		a: a,
	}

	a.addRunner(pm)

	a.Client().OnEnterPupilSystemReq(pm.onEnterPupilSystemResp)
	a.Client().OnPupilTrainReq(pm.onPupilTrainResp)
	a.Client().OnPupilGraduatedUnMarryListSync(pm.onPupilGraduatedUnMarryListSync)

	return pm
}

func (pm *PupilManager) Run(ctx context.Context) error {
	pm.check(ctx)

	pm.a.addJob("pupil", "@every 5m", pm.check)

	return nil
}

func (pm *PupilManager) check(ctx context.Context) error {
	pm.enter(ctx)

	return nil
}

func (pm *PupilManager) Stop() error {
	return nil
}

// EnterPupilSystemReq 进入门徒系统
func (pm *PupilManager) enter(ctx context.Context) {
	pm.a.Client().SendEnterPupilSystemReq(ctx, &pb.EnterPupilSystemReq{})
}

// onEnterPupilSystemResp
func (pm *PupilManager) onEnterPupilSystemResp(ctx client.Context, resp *pb.EnterPupilSystemResp) error {
	pm.a.L().Info("进入门徒系统")

	pm.checkState(ctx, resp.GetSiteList())

	for i := resp.GetGetTimes(); i < 2; i++ {
		pm.getAdReward(ctx)
	}

	return nil
}

// PupilTrainReq 门徒培养
func (pm *PupilManager) train(ctx context.Context) error {
	return pm.a.Client().SendPupilTrainReq(ctx, &pb.PupilTrainReq{
		IsOneKey: utils.Ptr(int32(1)),
	})
}

func (pm *PupilManager) checkState(ctx client.Context, siteList []*pb.PupilSiteDataMsg) error {
	needTrain := false

	for _, v := range siteList {
		if pm.canGraduate(v) {
			pm.graduate(ctx, v.GetIndex())
			pm.recruit(ctx)
		}

		if v.GetTrainTimeInfo().GetRemainTrainTimes() > 0 {
			needTrain = true
		}
	}

	if needTrain {
		pm.train(ctx)
	}

	return nil
}

// onPupilTrainResp
func (pm *PupilManager) onPupilTrainResp(ctx client.Context, resp *pb.PupilTrainResp) error {
	pm.checkState(ctx, resp.GetSiteList())
	return nil
}

// onPupilGraduatedUnMarryListSync 门徒毕业未结婚列表同步
func (pm *PupilManager) onPupilGraduatedUnMarryListSync(ctx client.Context, resp *pb.PupilGraduatedUnMarryListSync) error {
	return nil
}

// recruit 招收弟子
func (pm *PupilManager) recruit(ctx client.Context) error {
	return pm.a.Client().SendPupilRecruitReq(ctx, &pb.PupilRecruitReq{})
}

func (pm *PupilManager) onRecruitResp(ctx client.Context, resp *pb.PupilRecruitResp) error {
	return nil
}

func (pm *PupilManager) canGraduate(msg *pb.PupilSiteDataMsg) bool {
	pd := msg.GetPupilData()
	td := msg.GetTrainTimeInfo()
	return td.GetTrainTimes() >= pd.GetLevel()*20
}

// 出师
func (pm *PupilManager) graduate(ctx client.Context, idx int32) error {
	return pm.a.Client().SendPupilGraduateReq(ctx, &pb.PupilGraduateReq{SiteIndex: &idx})
}

// getAdReward 获取广告奖励
func (pm *PupilManager) getAdReward(ctx client.Context) error {
	return pm.a.Client().SendPupilGetAdRewardReq(ctx, &pb.PupilGetAdRewardReq{IsUseADTime: utils.Ptr(false)})
}

func (pm *PupilManager) onGetAdRewardResp(ctx client.Context, resp *pb.PupilGetAdRewardResp) error {
	return nil
}
