// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// GroupStore is an autogenerated mock type for the GroupStore type
type GroupStore struct {
	mock.Mock
}

// AdminRoleGroupsForSyncableMember provides a mock function with given fields: userID, syncableID, syncableType
func (_m *GroupStore) AdminRoleGroupsForSyncableMember(userID string, syncableID string, syncableType model.GroupSyncableType) ([]string, error) {
	ret := _m.Called(userID, syncableID, syncableType)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) []string); ok {
		r0 = rf(userID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GroupSyncableType) error); ok {
		r1 = rf(userID, syncableID, syncableType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelMembersMinusGroupMembers provides a mock function with given fields: channelID, groupIDs, page, perPage
func (_m *GroupStore) ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, error) {
	ret := _m.Called(channelID, groupIDs, page, perPage)

	var r0 []*model.UserWithGroups
	if rf, ok := ret.Get(0).(func(string, []string, int, int) []*model.UserWithGroups); ok {
		r0 = rf(channelID, groupIDs, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserWithGroups)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, int, int) error); ok {
		r1 = rf(channelID, groupIDs, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelMembersToAdd provides a mock function with given fields: since, channelID, includeRemovedMembers
func (_m *GroupStore) ChannelMembersToAdd(since int64, channelID *string, includeRemovedMembers bool) ([]*model.UserChannelIDPair, error) {
	ret := _m.Called(since, channelID, includeRemovedMembers)

	var r0 []*model.UserChannelIDPair
	if rf, ok := ret.Get(0).(func(int64, *string, bool) []*model.UserChannelIDPair); ok {
		r0 = rf(since, channelID, includeRemovedMembers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserChannelIDPair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *string, bool) error); ok {
		r1 = rf(since, channelID, includeRemovedMembers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelMembersToRemove provides a mock function with given fields: channelID
func (_m *GroupStore) ChannelMembersToRemove(channelID *string) ([]*model.ChannelMember, error) {
	ret := _m.Called(channelID)

	var r0 []*model.ChannelMember
	if rf, ok := ret.Get(0).(func(*string) []*model.ChannelMember); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string) error); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountChannelMembersMinusGroupMembers provides a mock function with given fields: channelID, groupIDs
func (_m *GroupStore) CountChannelMembersMinusGroupMembers(channelID string, groupIDs []string) (int64, error) {
	ret := _m.Called(channelID, groupIDs)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, []string) int64); ok {
		r0 = rf(channelID, groupIDs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(channelID, groupIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountGroupsByChannel provides a mock function with given fields: channelID, opts
func (_m *GroupStore) CountGroupsByChannel(channelID string, opts model.GroupSearchOpts) (int64, error) {
	ret := _m.Called(channelID, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) int64); ok {
		r0 = rf(channelID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) error); ok {
		r1 = rf(channelID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountGroupsByTeam provides a mock function with given fields: teamID, opts
func (_m *GroupStore) CountGroupsByTeam(teamID string, opts model.GroupSearchOpts) (int64, error) {
	ret := _m.Called(teamID, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) int64); ok {
		r0 = rf(teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) error); ok {
		r1 = rf(teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountTeamMembersMinusGroupMembers provides a mock function with given fields: teamID, groupIDs
func (_m *GroupStore) CountTeamMembersMinusGroupMembers(teamID string, groupIDs []string) (int64, error) {
	ret := _m.Called(teamID, groupIDs)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, []string) int64); ok {
		r0 = rf(teamID, groupIDs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(teamID, groupIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: group
func (_m *GroupStore) Create(group *model.Group) (*model.Group, error) {
	ret := _m.Called(group)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(*model.Group) *model.Group); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Group) error); ok {
		r1 = rf(group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, error) {
	ret := _m.Called(groupSyncable)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) *model.GroupSyncable); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.GroupSyncable) error); ok {
		r1 = rf(groupSyncable)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: groupID
func (_m *GroupStore) Delete(groupID string) (*model.Group, error) {
	ret := _m.Called(groupID)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string) *model.Group); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) *model.GroupSyncable); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GroupSyncableType) error); ok {
		r1 = rf(groupID, syncableID, syncableType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) DeleteMember(groupID string, userID string) (*model.GroupMember, error) {
	ret := _m.Called(groupID, userID)

	var r0 *model.GroupMember
	if rf, ok := ret.Get(0).(func(string, string) *model.GroupMember); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(groupID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DistinctGroupMemberCount provides a mock function with given fields:
func (_m *GroupStore) DistinctGroupMemberCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: groupID
func (_m *GroupStore) Get(groupID string) (*model.Group, error) {
	ret := _m.Called(groupID)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string) *model.Group); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBySource provides a mock function with given fields: groupSource
func (_m *GroupStore) GetAllBySource(groupSource model.GroupSource) ([]*model.Group, error) {
	ret := _m.Called(groupSource)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func(model.GroupSource) []*model.Group); ok {
		r0 = rf(groupSource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GroupSource) error); ok {
		r1 = rf(groupSource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllGroupSyncablesByGroupId provides a mock function with given fields: groupID, syncableType
func (_m *GroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, error) {
	ret := _m.Called(groupID, syncableType)

	var r0 []*model.GroupSyncable
	if rf, ok := ret.Get(0).(func(string, model.GroupSyncableType) []*model.GroupSyncable); ok {
		r0 = rf(groupID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.GroupSyncable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSyncableType) error); ok {
		r1 = rf(groupID, syncableType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDs provides a mock function with given fields: groupIDs
func (_m *GroupStore) GetByIDs(groupIDs []string) ([]*model.Group, error) {
	ret := _m.Called(groupIDs)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func([]string) []*model.Group); ok {
		r0 = rf(groupIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(groupIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name, opts
func (_m *GroupStore) GetByName(name string, opts model.GroupSearchOpts) (*model.Group, error) {
	ret := _m.Called(name, opts)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) *model.Group); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRemoteID provides a mock function with given fields: remoteID, groupSource
func (_m *GroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, error) {
	ret := _m.Called(remoteID, groupSource)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string, model.GroupSource) *model.Group); ok {
		r0 = rf(remoteID, groupSource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSource) error); ok {
		r1 = rf(remoteID, groupSource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySource provides a mock function with given fields: groupSource, page, perPage
func (_m *GroupStore) GetBySource(groupSource model.GroupSource, page int, perPage int) ([]*model.Group, error) {
	ret := _m.Called(groupSource, page, perPage)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func(model.GroupSource, int, int) []*model.Group); ok {
		r0 = rf(groupSource, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GroupSource, int, int) error); ok {
		r1 = rf(groupSource, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUser provides a mock function with given fields: userID
func (_m *GroupStore) GetByUser(userID string) ([]*model.Group, error) {
	ret := _m.Called(userID)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func(string) []*model.Group); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) *model.GroupSyncable); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GroupSyncableType) error); ok {
		r1 = rf(groupID, syncableID, syncableType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroups provides a mock function with given fields: page, perPage, opts
func (_m *GroupStore) GetGroups(page int, perPage int, opts model.GroupSearchOpts) ([]*model.Group, error) {
	ret := _m.Called(page, perPage, opts)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func(int, int, model.GroupSearchOpts) []*model.Group); ok {
		r0 = rf(page, perPage, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, model.GroupSearchOpts) error); ok {
		r1 = rf(page, perPage, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupsAssociatedToChannelsByTeam provides a mock function with given fields: teamID, opts
func (_m *GroupStore) GetGroupsAssociatedToChannelsByTeam(teamID string, opts model.GroupSearchOpts) (map[string][]*model.GroupWithSchemeAdmin, error) {
	ret := _m.Called(teamID, opts)

	var r0 map[string][]*model.GroupWithSchemeAdmin
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) map[string][]*model.GroupWithSchemeAdmin); ok {
		r0 = rf(teamID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]*model.GroupWithSchemeAdmin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) error); ok {
		r1 = rf(teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupsByChannel provides a mock function with given fields: channelID, opts
func (_m *GroupStore) GetGroupsByChannel(channelID string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, error) {
	ret := _m.Called(channelID, opts)

	var r0 []*model.GroupWithSchemeAdmin
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) []*model.GroupWithSchemeAdmin); ok {
		r0 = rf(channelID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.GroupWithSchemeAdmin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) error); ok {
		r1 = rf(channelID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupsByTeam provides a mock function with given fields: teamID, opts
func (_m *GroupStore) GetGroupsByTeam(teamID string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, error) {
	ret := _m.Called(teamID, opts)

	var r0 []*model.GroupWithSchemeAdmin
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) []*model.GroupWithSchemeAdmin); ok {
		r0 = rf(teamID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.GroupWithSchemeAdmin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) error); ok {
		r1 = rf(teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMemberCount provides a mock function with given fields: groupID
func (_m *GroupStore) GetMemberCount(groupID string) (int64, error) {
	ret := _m.Called(groupID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(groupID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMemberUsers provides a mock function with given fields: groupID
func (_m *GroupStore) GetMemberUsers(groupID string) ([]*model.User, error) {
	ret := _m.Called(groupID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMemberUsersInTeam provides a mock function with given fields: groupID, teamID
func (_m *GroupStore) GetMemberUsersInTeam(groupID string, teamID string) ([]*model.User, error) {
	ret := _m.Called(groupID, teamID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string) []*model.User); ok {
		r0 = rf(groupID, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(groupID, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMemberUsersNotInChannel provides a mock function with given fields: groupID, channelID
func (_m *GroupStore) GetMemberUsersNotInChannel(groupID string, channelID string) ([]*model.User, error) {
	ret := _m.Called(groupID, channelID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string) []*model.User); ok {
		r0 = rf(groupID, channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(groupID, channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMemberUsersPage provides a mock function with given fields: groupID, page, perPage
func (_m *GroupStore) GetMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, error) {
	ret := _m.Called(groupID, page, perPage)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.User); ok {
		r0 = rf(groupID, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(groupID, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupChannelCount provides a mock function with given fields:
func (_m *GroupStore) GroupChannelCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupCount provides a mock function with given fields:
func (_m *GroupStore) GroupCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupCountWithAllowReference provides a mock function with given fields:
func (_m *GroupStore) GroupCountWithAllowReference() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupMemberCount provides a mock function with given fields:
func (_m *GroupStore) GroupMemberCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupTeamCount provides a mock function with given fields:
func (_m *GroupStore) GroupTeamCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteMembersByUser provides a mock function with given fields: userID
func (_m *GroupStore) PermanentDeleteMembersByUser(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PermittedSyncableAdmins provides a mock function with given fields: syncableID, syncableType
func (_m *GroupStore) PermittedSyncableAdmins(syncableID string, syncableType model.GroupSyncableType) ([]string, error) {
	ret := _m.Called(syncableID, syncableType)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, model.GroupSyncableType) []string); ok {
		r0 = rf(syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.GroupSyncableType) error); ok {
		r1 = rf(syncableID, syncableType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamMembersMinusGroupMembers provides a mock function with given fields: teamID, groupIDs, page, perPage
func (_m *GroupStore) TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, error) {
	ret := _m.Called(teamID, groupIDs, page, perPage)

	var r0 []*model.UserWithGroups
	if rf, ok := ret.Get(0).(func(string, []string, int, int) []*model.UserWithGroups); ok {
		r0 = rf(teamID, groupIDs, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserWithGroups)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, int, int) error); ok {
		r1 = rf(teamID, groupIDs, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamMembersToAdd provides a mock function with given fields: since, teamID, includeRemovedMembers
func (_m *GroupStore) TeamMembersToAdd(since int64, teamID *string, includeRemovedMembers bool) ([]*model.UserTeamIDPair, error) {
	ret := _m.Called(since, teamID, includeRemovedMembers)

	var r0 []*model.UserTeamIDPair
	if rf, ok := ret.Get(0).(func(int64, *string, bool) []*model.UserTeamIDPair); ok {
		r0 = rf(since, teamID, includeRemovedMembers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserTeamIDPair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *string, bool) error); ok {
		r1 = rf(since, teamID, includeRemovedMembers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamMembersToRemove provides a mock function with given fields: teamID
func (_m *GroupStore) TeamMembersToRemove(teamID *string) ([]*model.TeamMember, error) {
	ret := _m.Called(teamID)

	var r0 []*model.TeamMember
	if rf, ok := ret.Get(0).(func(*string) []*model.TeamMember); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: group
func (_m *GroupStore) Update(group *model.Group) (*model.Group, error) {
	ret := _m.Called(group)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(*model.Group) *model.Group); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Group) error); ok {
		r1 = rf(group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, error) {
	ret := _m.Called(groupSyncable)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) *model.GroupSyncable); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.GroupSyncable) error); ok {
		r1 = rf(groupSyncable)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) UpsertMember(groupID string, userID string) (*model.GroupMember, error) {
	ret := _m.Called(groupID, userID)

	var r0 *model.GroupMember
	if rf, ok := ret.Get(0).(func(string, string) *model.GroupMember); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(groupID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
