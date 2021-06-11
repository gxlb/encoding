package encoding

import (
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	type testCase struct {
		json, yml, xml string
	}
	cases := []*testCase{
		&testCase{
			`
{
	"a": "aaa",	
	"b": 1234,
	"b2": 1234.56,
	"c": true,
	"d": ["xx", "yy", "zz"],
	"e": [11, 22, 33]		
}
`,
			``,
			`
<data>
  <ret>0</ret>
  <errcode>0</errcode>
  <msg>ok</msg>
  <data>
     <timestamp>128679200</timestamp>
     <hasnext>0</hasnext>
     <info>
          <name></name>
          <openid>B624064BA065E01CB73F835017FE96FA</openid>
          <nick>aaaa</nick>
          <head>http://app.qlogo.cn/mbloghead/563ad8b6be488a07a694</head>
          <sex>1</sex>
          <location>广东 深圳</location>
          <country_code>1</country_code>
          <province_code>44</province_code>		   
          <city_code>3</city_code>
          <tweet>
               <text></text>
               <from>来自网页</from>
               <id>7987543214334</id>
               <timestamp>1285813236</timestamp>		         
          </tweet>
          <fansnum>15</fansnum>
          <idolnum>20</idolnum>
          <isfans>1</isfans>
          <isvip>0</isvip>
          <tag>
               <id>1</id>
               <name></name>
          </tag>
     </info>
  </data>
</data>
			`,
		},
	}

	for i, v := range cases {
		{
			d, _ := FromJson(v.json)
			fmt.Printf("%d %s %#v\n", i, v, d)
			yml, _ := ToYaml(d, true)
			fmt.Printf("yml:\n%s\n", yml)
			xml, _ := ToXml(d, true)
			fmt.Printf("xml:\n%s\n", xml)
		}
		if v.xml != "" {
			d, _ := FromXml(v.xml)
			fmt.Printf("%d %s %#v\n", i, v, d)
		}

	}

	var xx = []interface{}{
		123,
		"string",
		true,
		12.34,
		[]string{"a", "b", "c"},
		[]int{1, 2, 3, 4},
	}
	xml, _ := ToXml(xx, true)
	fmt.Printf("xml:\n%s\n", xml)
	xmls, _ := FromXml(xml)
	fmt.Printf("xml:\n%#v\n", xmls)

	jsn, _ := ToJson(xx, true)
	fmt.Printf("json:\n%s\n", jsn)
}
