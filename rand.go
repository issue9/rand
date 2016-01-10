// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// 生成各种随机字符串的包
package rand

import (
	"errors"
	mr "math/rand"
	"time"
)

const (
	Upper = iota // 大写字母
	Lower        // 小写字母
	Digit        // 数字
	Punct        // 标点符号
	size
)

// 随机种子
var random = mr.New(mr.NewSource(time.Now().Unix()))

// 随机字符串取值的表
var table = [][]byte{
	Upper: []byte("ABCDEFGHIJKLMNPQRSTUVWXYZ"),
	Lower: []byte("abcdefghijkmnpqrstuvwxyz"),
	Digit: []byte("0123456789"),
	Punct: []byte("~!@#$%^&*()_+-={}|[]\\:\";'<>,.?/"),
}

// 产生随机字符数组，其长度为[min, max)。
// cats 随机字符串的类型，指定非法值是触发panic
func Bytes(min, max int, cats ...int) []byte {
	if err := checkArgs(min, max, cats); err != nil {
		panic(err)
	}

	return bytes(min+random.Intn(max-min), cats)
}

// 产生一个随机字符串，其长度为[min, max)。
// cats 随机字符串的类型，指定非法值是触发panic
func String(min, max int, cats ...int) string {
	return string(Bytes(min, max, cats...))
}

// Rand 提供了对参数的简单包装，方便用户批量产生相同的类型的随机字符串。
type Rand struct {
	min, max  int
	cats      []int
	hasBuffer bool
	channel   chan []byte
}

// 声明一个Rand变量。
// bufferSize 缓存的随机字符串数量，若为0,表示不缓存。
func New(bufferSize int, min, max int, cats ...int) (*Rand, error) {
	if err := checkArgs(min, max, cats); err != nil {
		return nil, err
	}

	ret := &Rand{
		min:       min,
		max:       max,
		cats:      cats,
		hasBuffer: bufferSize > 0,
	}
	if ret.hasBuffer {
		ret.channel = make(chan []byte, bufferSize)
		go func() {
			for {
				ret.channel <- bytes(min+random.Intn(max-min), cats)
			}
		}()
	}
	return ret, nil
}

// 产生随机字符数组，功能与全局函数Bytes()相同，但参数通过NewRand()预先指定。
func (r *Rand) Bytes() []byte {
	if r.hasBuffer {
		return <-r.channel
	}

	return bytes(r.min+random.Intn(r.max-r.min), r.cats)
}

// 产生一个随机字符串，功能与全局函数String()相同，但参数通过NewRand()预先指定。
func (r *Rand) String() string {
	if r.hasBuffer {
		return string(<-r.channel)
	}

	return string(r.Bytes())
}

// 生成指定指定长度的随机字符数组
func bytes(l int, cats []int) []byte {
	bs := make([]byte, 0, l)
	i := 0
	for {
		for _, cat := range cats {
			s := table[cat]
			index := random.Intn(len(s))
			bs = append(bs, s[index])

			i++
			if i >= l {
				return bs
			}
		}
	}
}

// 检测各个参数是否合法
func checkArgs(min, max int, cats []int) error {
	if min <= 0 {
		return errors.New("rand.checkArgs:min值必须大于0")
	}
	if max <= min {
		return errors.New("rand.checkArgs:max必须大于min")
	}

	if len(cats) == 0 {
		return errors.New("rand.checkArgs:无效的cat参数")
	}

	for _, cat := range cats {
		if cat < 0 || cat >= size {
			return errors.New("rand.Bytes:无效的cat参数")
		}
	}

	return nil
}
