package member

import (
	"context"
	"reflect"

	user_dto "github.com/eolinker/ap-account/module/user/dto"
	department_member "github.com/eolinker/ap-account/service/department-member"
	"github.com/eolinker/ap-account/service/member"
	"github.com/eolinker/ap-account/service/user"
	user_group "github.com/eolinker/ap-account/service/user-group"
	"github.com/eolinker/go-common/auto"
	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/utils"
)

var (
	_ IMemberModule = (*imlMemberModule)(nil)
)

type IMemberModule interface {
	UserGroupMember(ctx context.Context, keyword string, groupId ...string) ([]*user_dto.UserInfo, error)
}

type imlMemberModule struct {
	memberService           user_group.IUserGroupMemberService `autowired:""`
	departmentMemberService department_member.IMemberService   `autowired:""`
	userService             user.IUserService                  `autowired:""`
}

func (m *imlMemberModule) UserGroupMember(ctx context.Context, keyword string, groupId ...string) ([]*user_dto.UserInfo, error) {
	us, err := m.userService.Search(ctx, keyword, -1)
	if err != nil {
		return nil, err
	}
	userIds := utils.SliceToSlice(us, func(s *user.User) string {
		return s.UID
	})

	members, err := m.memberService.Members(ctx, groupId, userIds)
	if err != nil {
		return nil, err
	}

	userids := utils.SliceToSlice(members, member.UserID, func(m *member.Member) bool {
		return m.Come != ""
	})

	if len(userids) == 0 {
		return nil, nil
	}
	users, err := m.userService.Get(ctx, userids...)
	if err != nil {
		return nil, err
	}
	result := utils.SliceToSlice(users, user_dto.CreateUserInfoFromModel)

	groups, err := m.memberService.FilterMembersForUser(ctx, userids...)
	if err != nil {
		return nil, err
	}
	departmentMembers, err := m.departmentMemberService.Members(ctx, nil, userids)
	if err != nil {
		return nil, err
	}
	departmentMemberMap := utils.SliceToMapArrayO(utils.SliceToSlice(departmentMembers, func(s *member.Member) *member.Member {
		return s
	}, func(m *member.Member) bool {
		return m.Come != ""
	}), func(t *member.Member) (string, string) {
		return t.UID, t.Come
	})
	for _, r := range result {
		r.Department = auto.List(departmentMemberMap[r.Uid])
		r.UserGroups = auto.List(groups[r.Uid])
	}
	return result, nil
}

func init() {
	autowire.Auto[IMemberModule](func() reflect.Value {
		return reflect.ValueOf(new(imlMemberModule))
	})
}
