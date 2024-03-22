package lib

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlQuery(t *testing.T) {
	vars := url.Values{
		"x": []string{"1&xx"},
		"y": []string{"2%yy"},
	}
	unescaped, err := url.QueryUnescape(vars.Encode())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(unescaped)
}

const GrafanaAutoLoginUrlTmpl = `https://shield.cdb.woa.com/auto_login?dest_url=%s`

func TestUrlEncode(t *testing.T) {
	var target = "https://shield.cdb.woa.com/d/CCBVbWcIz/man-ri-zhi-cha-xun?orgId=1&var-region=test&var-datasource=log_cluster_test&var-insttype=master&var-insttype=slave&var-insttype=ro&var-instance_id=70ebc680-8370-450d-964e-32151c332792&var-node_id=dbn-fjdvgpfk"
	autoUrl := fmt.Sprintf(GrafanaAutoLoginUrlTmpl, url.QueryEscape(target))
	t.Log(autoUrl)

	vars := url.Values{
		"var-instid":     []string{"1111"},
		"var-datasource": []string{"log_" + "dataSource"},
		"var-insttype":   []string{"master", "slave", "ro", "tke"},
		"var-keyword":    []string{"%error%"},
	}
	unescaped, _ := url.QueryUnescape(vars.Encode())
	u := &url.URL{
		Scheme:   "https",
		Host:     "woa.com",
		Path:     "/d/CCBVbWcIz/man-ri-zhi-cha-xun",
		RawQuery: unescaped,
	}
	t.Log(u.String())
}
