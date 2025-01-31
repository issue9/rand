// SPDX-FileCopyrightText: 2016-2024 caixw
//
// SPDX-License-Identifier: MIT

package rands

import (
	"context"
	"math/rand/v2"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/issue9/assert/v4"
)

func TestCheckArgs(t *testing.T) {
	a := assert.New(t, false)

	// min < 0
	a.Panic(func() {
		checkArgs(-1, 1, []byte("12"))
	})

	// max <= min
	a.Panic(func() {
		checkArgs(5, 5, []byte("12"))
	})

	// cats为空
	a.Panic(func() {
		checkArgs(5, 6, []byte(""))
	})

	a.NotPanic(func() { checkArgs(5, 6, []byte("123")) })
}

func TestGen(t *testing.T) {
	a := assert.New(t, false)

	r1 := rand.IntN
	r2 := rand.Uint64
	a.NotEqual(gen(r1, r2, 10, 11, []byte("1234123lks;df")), gen(r1, r2, 10, 11, []byte("1234123lks;df")))
	a.NotEqual(gen(r1, r2, 10, 11, []byte("1234123lks;df")), gen(r1, r2, 10, 11, []byte("1234123lks;df")))
	a.NotEqual(gen(r1, r2, 10, 11, []byte("1234123lks;df")), gen(r1, r2, 10, 11, []byte("1234123lks;df")))
}

func TestAppend(t *testing.T) {
	a := assert.New(t, false)
	bs := []byte{}
	a.Equal(len(Append(bs, 8, 9, []byte("1ks;dfp123;4j;ladj;fpoqwe"))), 8)
}

func TestBytes(t *testing.T) {
	a := assert.New(t, false)

	// 测试固定长度
	a.Equal(len(Bytes(8, 9, []byte("1ks;dfp123;4j;ladj;fpoqwe"))), 8)
	a.Equal(utf8.RuneCount(Bytes(8, 9, []rune("中文内容也可以正常显示"))), 8)

	// 非固定长度
	l := len(Bytes(8, 10, []byte("adf;wieqpwekwjerpq")))
	a.True(l >= 8 && l <= 10)

}

func TestString(t *testing.T) {
	a := assert.New(t, false)

	// 查看是否输出完整的字符
	val := String(10, 11, []rune("中文内容也d可e以正常显示abc"))
	t.Log("这将显示一段随机的中英文：", val)
	for _, r := range val {
		a.True(utf8.ValidRune(r))
	}
}

func TestRands(t *testing.T) {
	a := assert.New(t, false)

	a.PanicString(func() {
		New(rand.New(rand.NewPCG(0, 0)), 0, 5, 7, []byte(";adkfjpqwei12124nbnb"))
	}, "bufferSize 必须大于零")

	r := New(nil, 2, 5, 7, []byte(";adkfjpqwei12124nbnb"))
	a.NotNil(r)
	ctx, cancel := context.WithCancel(context.Background())
	go r.Serve(ctx)
	time.Sleep(time.Microsecond * 500) // 等待 go 运行完成
	a.Equal(cap(r.channel), 2)

	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.String(), r.String())
	a.NotEqual(r.Bytes(), r.Bytes())

	cancel()
	time.Sleep(time.Microsecond * 500) // 等待 cancel 运行完成
	a.NotEqual(r.String(), r.String()) // 读取 channel 中剩余的数据
	a.Equal(r.String(), r.String())    // 没有数据了，都是空值
}
