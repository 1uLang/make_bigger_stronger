/*
 * Author: lihy lihy@zhiannet.com
 * Date: 2022-11-13 22:51:07
 * LastEditors: lihy lihy@zhiannet.com
 * Note: Need note condition
 */
package tools

import "testing"

func TestFilterInPlace(t *testing.T) {
	sl := interface{}(1)
	f := interface{}(2)
	FilterInPlace(sl, f)
}

func TestContains(t *testing.T) {
	a := []string{"account1", "test1234", "aaa"}
	b := "account1"
	t.Log(Contains(b, a))
}
