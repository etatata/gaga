package main

import (
	"errors"
	"time"
)

//第三题
type cache struct {
	key     string
	val     string
	expire  int
	creatAt int
}

const defaultExpire = 100 //默认过期时间

type cacheManager struct {
	m map[string]cache
}

func (m *cacheManager) add(key string, val string, expire int) error {

	if key == "" {
		return errors.New("key is empty")
	}

	if expire <= 0 {
		expire = defaultExpire
	}

	c := cache{key, val, expire, time.Now().Second()}
	m.m[key] = c

	return nil
}

func (m *cacheManager) delete(key string) {

	delete(m.m, key)
}

func (m *cacheManager) get(key string) (string, error) {
	val := m.m[key]

	if _, ok := m.m[key]; !ok {
		return "", errors.New("not found")
	}

	now := time.Now().Second()
	if (now - val.creatAt) >= val.expire {
		m.delete(key)
		return "", errors.New("expire")
	}

	return val.val, nil
}

func (m *cacheManager) modify(key string, val string, expire int) (bool, error) {
	if key == "" {
		return false, errors.New("key empty")
	}

	if _, ok := m.m[key]; !ok {
		return false, errors.New("not found")
	}

	value := m.m[key]
	value.val = val
	if expire > 0 {
		value.expire = expire
	}

	m.m[key] = value

	return true, nil
}
