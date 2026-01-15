package lib

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func fakeString() string {
	return "xxx"
}

func fakeY() string {
	return "yyy"
}

func TestStringEqual(t *testing.T) {
	var (
		x = "xxx"
	)
	t.Log(x == fakeString())
	t.Log(x == fakeY())
}

// SQL> SELECT banner FROM v$version;
func TestRegex(t *testing.T) {
	reg := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+\.\d+)`)
	ss := []string{
		"Oracle Database 11g Enterprise Edition Release 11.2.0.4.0 - 64bit Production",
		"PL/SQL Release 11.2.0.4.0 - Production",
		"CORE    11.2.0.4.0      Production",
		"TNS for Linux: Version 11.2.0.4.0 - Production",
		"NLSRTL Version 11.2.0.4.0 - Production",
	}
	for _, s := range ss {
		version := reg.FindString(s)
		t.Log(version)
		split := strings.Split(version, ".")
		if len(split) == 5 {
			t.Log(split[0])
		} else {
			t.Fatalf("invalid version: %s", version)
		}
	}
}

func TestStringToInt(t *testing.T) {
	var ss = []string{
		"0.62",
		"1.3",
		"4",
	}
	for _, s := range ss {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(atoi)
	}
}

func TestStringIndex(t *testing.T) {
	var ss = []string{
		"oracle/12.2.2.2.2",
		"cls/xcss.s.c.s",
		"oracle/19.3.3.3",
	}
	for _, s := range ss {
		if strings.HasPrefix(s, "oracle/") {
			t.Log(s[7:])
		}
	}
}

func TestEmptyString(t *testing.T) {
	var s = ""
	t.Log(strings.Split(s, ";"))
}

func TestStrSplit(t *testing.T) {
	var strs = []string{
		"xxx.ccc_CM_ddd",
		"xxx.ccc_ddd",
		"xxx.ddd",
		" ",
		"",
	}
	for _, str := range strs {
		t.Log(strings.Split(str, "_CM_"))
	}
}

func TestStrConv(t *testing.T) {
	var ss = []string{
		"xxx",
		"34",
	}
	for _, s := range ss {
		atoi, err := strconv.Atoi(s)
		if err != nil {
			t.Log(err)
		}
		t.Log(">>>", atoi)
	}
}

func TestStr2Slice(t *testing.T) {
	var s = "[\"成都电信西区803DC11号楼M705\"，\"上海电信西区803DC11号楼M705\"]"
	var ss = gconv.Strings(s)

	t.Log(len(ss), ss)
}

type GlobalSearchReq struct {
	MessageType string `json:"MessageType" dc:"消息类型"`
	Duration    int64  `json:"Duration" dc:"查询耗时(ms)"`
	Succeed     bool   `json:"Succeed" dc:"是否查询成功"`
	Message     string `json:"Message" dc:"查询失败时返回原因"`
}

// MD5 计算一次请求的md5，用于过滤同一个链接的多次请求
// struct的json序列化是有序的，因此可以先进行json序列化然后求md5
func (g *GlobalSearchReq) MD5() string {
	j, err := json.Marshal(g)
	if err != nil {
		return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(g))))
	}

	return fmt.Sprintf("%x", md5.Sum(j))
}

func TestMD5(t *testing.T) {
	var (
		items = []GlobalSearchReq{
			{"1", int64(time.Minute), true, "test"},
			{"1", int64(time.Minute), true, "test"},
			{"2", int64(time.Minute), true, "test"},
			{"3", int64(time.Minute), true, "test"},
		}
	)
	for _, item := range items {
		t.Log(item.MD5())
	}

}

func TestSplitStr(t *testing.T) {
	var strs = []string{
		"育英技术中心",
		"|TEG技术工程事业群|数据库研发部|运营技术中心",
		"",
	}
	for _, c := range strs {
		items := gstr.SplitAndTrim(c, "|")
		if len(items) > 0 {
			t.Log(items[len(items)-1])
		}
	}

	t.Log("pass")
}

func TestFindSubString(t *testing.T) {
	var (
		AZServiceNameRegex     = regexp.MustCompile(`^txstore_(?P<RegionKey>[a-zA-Z][a-zA-Z0-9-]{1,49})_(?P<ZoneKey>[a-zA-Z][a-zA-Z0-9-]{1,49})_(?P<Suffix>[a-z-]{2,16})$`)
		RegionServiceNameRegex = regexp.MustCompile(`^txstore_(?P<RegionKey>[a-zA-Z][a-zA-Z0-9-]{1,49})_(?P<Suffix>[a-z-]{2,16})$`)
		ss                     = []string{
			"txstore_global_api-server",
			"txstore_cross-region-1_cross-region-1-az_dbres",
			"txstore_ncdbtest-80036_cdp-checker",
			"txstore_region-W8HJ1D_az-773482_dbmgr",
			"xxx-cxcsjdjs",
		}
	)

	for _, s := range ss {
		switch {
		case RegionServiceNameRegex.MatchString(s):
			t.Log("地域匹配", s)
			items := RegionServiceNameRegex.FindStringSubmatch(s)
			if len(items) > 0 {
				t.Log(items[RegionServiceNameRegex.SubexpIndex("RegionKey")])
			}
		case AZServiceNameRegex.MatchString(s):
			t.Log("az匹配", s)
			items := AZServiceNameRegex.FindStringSubmatch(s)
			if len(items) > 0 {
				t.Log(items[AZServiceNameRegex.SubexpIndex("ZoneKey")])
			}
		default:
			t.Log("不匹配", s)
		}
	}
}
