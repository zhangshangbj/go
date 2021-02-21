package test1

import "testing"

func TestMd5(t *testing.T) {
	hashs := map[string]string{
		"a": "0cc175b9c0f1b6a831c399e269772661",
		"b": "92eb5ffee6ae2fec3ad71c777531578f",
	}
	for k, v := range hashs {
		if Md5(k) != v {
			t.Errorf("%s md5 not is %s", k, v)
		}
	}
}
