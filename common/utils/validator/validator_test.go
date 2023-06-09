/*
 * Author: lihy lihy@zhiannet.com
 * Date: 2023-01-09 12:25:21
 * LastEditors: lihy lihy@zhiannet.com
 * Note: Need note condition
 */
package validator

import (
	"strings"
	"testing"
)

func TestCheckIPV4(t *testing.T) {
	ips := "12.15"
	ipsArr := strings.Split(ips, ";")
	for _, v := range ipsArr {
		t.Log(v, CheckIPV4("12.15"))
	}
}
