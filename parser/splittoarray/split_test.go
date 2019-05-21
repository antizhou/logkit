package parser

import (
	"log"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	lines := []string{
		`2019-05-16 00:00:00,120 INFO  ACCESS - 2019-05-16T00:00:00.108	100.116.189.158	8000	POST	/opentreaty-service/hashSign/evidence/attachment	404	11973	143	44bb057c-facc-45c5-a0b8-73953d9d0ada	286	"toapi.tsign.cn"	"47.96.79.204"	"-"	"-"	"-"	"okhttp/3.6.0"	"3438757846"	"cerberus-open/1.0.0"	"-"	"application/json;charset=UTF-8"`,
		`2019-05-16 00:00:32,186 INFO  ACCESS - 2019-05-16T00:00:32.181	100.116.189.180	8000	POST	/opentreaty-service/hashSign/evidence/attachment	404	4776	144	6aa3e65b-1f8e-4b74-ad87-378c14c45a59	287	"toapi.tsign.cn"	"47.96.79.204"	"-"	"-"	"-"	"okhttp/3.6.0"	"3438757846"	"cerberus-open/1.0.0"	"-"	"application/json;charset=UTF-8"`,
	}
	p := Parser{
		name:          "splittoarray",
		lineSeparator: `\s+`,
		fieldOrderMappings: map[string]string{
			"timestamp": "5",
			"remoteip":  "6",
			"uri":       "9",
			"status":    "10",
			"duration":  "11",
			"appid":     "21",
		},
		characterReplace: map[string]string{
			"\"": "",
		},
	}

	log.Println(p.Parse(lines))
}
