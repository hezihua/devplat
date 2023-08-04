package provider

import "hzh/devcloud/mcenter/apps/token"

var (
	store = map[token.GRANT_TYPE]Issuer{}
)

func Registry(key token.GRANT_TYPE, value Issuer) {
	store[key] = value
}

func Get(key token.GRANT_TYPE) Issuer {
	return store[key]
}

func Init() error {
	for _, v := range store {
		err := v.Config()
		if err != nil {
			return err
		}
	}
	return nil
}
