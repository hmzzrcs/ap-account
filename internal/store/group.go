package store

import (
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
	"reflect"
)

type IUserGroupStore interface {
	store.IBaseStore[UserGroup]
}

type IUserGroupMemberStore interface {
	store.IBaseStore[UserGroupMember]
}

type imlUserGroupStore struct {
	store.BaseStore[UserGroup]
}
type imlUserGroupMemberStore struct {
	store.BaseStore[UserGroupMember]
}

func init() {
	autowire.Auto[IUserGroupStore](func() reflect.Value {
		return reflect.ValueOf(new(imlUserGroupStore))
	})
	autowire.Auto[IUserGroupMemberStore](func() reflect.Value {
		return reflect.ValueOf(new(imlUserGroupMemberStore))
	})
}