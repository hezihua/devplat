package token

import (
	"fmt"
	"time"
)

func (req *IssueTokenRequest) Validate() error {

	switch req.GrantType {
	case GRANT_TYPE_PASSWORD, GRANT_TYPE_LDAP:
		if req.Username == "" || req.Password == "" {
			return fmt.Errorf("用户名或密码缺失")
		}

	}
	return nil
	// validator.V().Struct(i)
}

func NewToken() *Token {
	return &Token{
		AccessExpiredAt:  600,
		RefreshExpiredAt: 600 * 4,
		IssueAt:          time.Now().Unix(),
		Status:           &Status{},
		Meta:             map[string]string{},
	}
}

func (t *Token) CheckAliable() error {
	if t.Status.IsBlock {
		return fmt.Errorf(t.Status.BlockReason)
	}

	dura := time.Since(t.ExpiredTime())
	if dura >= 0 {
		return fmt.Errorf("Token 已经过期")
	}
	return nil
}

func (t *Token) ExpiredTime() time.Time {
	dura := time.Duration(t.AccessExpiredAt * int64(time.Second))
	return time.Unix(t.IssueAt, 0).Add(dura)
}

func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}
