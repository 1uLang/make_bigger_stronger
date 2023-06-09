/*
 * Author: lihy lihy@zhiannet.com
 * Date: 2022-11-12 16:32:42
 * LastEditors: lihy lihy@zhiannet.com
 * Note: slice tool funtions
 */
package tools

import (
	"errors"
	"reflect"
)

/*
 * Author: lihongyin
 * Date: 2022-11-13 00:01:06
 * Note: Remove items with value of t in slice a.
 * param {[]interface{}} a		origin slices
 * param {interface{}} t		the pattern to be removed
 * param {bool} flag			if true, will remove all items of value t, while false only remove the first.
 */
func TrimSlice(t interface{}, flag bool, a ...interface{}) (b []interface{}, e error) {
	if reflect.TypeOf(a[0]) != reflect.TypeOf(t) {
		return nil, &reflect.ValueError{}
	}
	b = []interface{}{}
	for i := 0; i <= len(a); i++ {
		if reflect.ValueOf(a[i]) == reflect.ValueOf(t) {
			if !flag {
				b = append(b, a[i+1:]...)
				break
			}
			continue
		}
		b = append(b, a[i])
	}
	return
}

func TrimInt64Slice(a []int64, t int64, flag bool) (b []int64) {
	b = []int64{}
	for i := 0; i < len(a); i++ {
		if a[i] == t {
			if !flag {
				if len(b) != i+1 {
					b = append(b, a[i:]...)
				}
				break
			}
			continue
		}
		b = append(b, a[i])
	}
	return
}

func FilterInPlace(sl interface{}, f interface{}) {
	v := reflect.ValueOf(sl).Elem()
	j := 0
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		if reflect.ValueOf(f).Call([]reflect.Value{e.Addr()})[0].Bool() {
			v.Index(j).Set(e)
			j++
		}
	}
	v.SetLen(j)
}

func Contains(obj, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not contains")
}

func DistinctInt64(target []int64) []int64 {
	if target == nil {
		return nil
	}
	m := make(map[int64]bool)
	for _, v := range target {
		m[v] = true
	}
	res := make([]int64, 0)
	for k := range m {
		res = append(res, k)
	}
	return res
}

func TrimStringArr(a []string, b string) []string {
	var c []string
	for _, v := range a {
		if v == b {
			continue
		}
		c = append(c, v)
	}
	return c
}
