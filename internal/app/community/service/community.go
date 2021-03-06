package service

import (
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/pkg/constant/e"
)

// GetCommunityList 获取所有文章标签
func (s *communityService) GetCommunityList() (resList []CommunityDto, err error) {
	var communityList []model.Community
	communityList, err = s.dao.GetCommunityList()
	for _, c := range communityList {
		var dto CommunityDto
		if err = gconv.Struct(c, &dto); err != nil {
			zap.L().Error("communityService.GetCommunityList->gconv.Struct failed", zap.Error(err))
			return nil, err
		}
		resList = append(resList, dto)
	}
	return resList, nil
}

// GetCommunityDetailByID 根据社区id获取社区详情
func (s *communityService) GetCommunityDetailByID(ID int64) (commDto *CommunityDto, err error) {
	var commuModel *model.Community

	if commuModel, err = s.dao.GetCommunityDetailByID(ID); err != nil {
		return nil, err
	}
	if err = gconv.Struct(commuModel, &commDto); err != nil {
		zap.L().Error(e.CodeConvDataErr.Msg(), zap.Error(err))
		return nil, err
	}
	return commDto, nil
}
