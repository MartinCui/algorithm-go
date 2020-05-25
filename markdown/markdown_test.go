package markdown

import "testing"

func TestMarkdown(t *testing.T) {
	t.Log(Translate("abc"))
	t.Log(Translate(`abc
def`))
	t.Log(Translate(`abc

def`))
	t.Log(Translate(`abc

def

> ghi
`))
	t.Log(Translate(`abc

def

> ghi
> 
> 1234
> 76777

abc
`))
	t.Log(Translate(`abc

simple list:
- l1
- l2
- l3

> ghi
> 
>> this is embedded quote:
>> - list inside of embedded quote: 1
>> - list inside of embedded quote: 2
> dddd

abc
`))
	t.Log(Translate(`abc
# h1
## h2

### h3

simple **list** and **second strong** and *single em*:
- l1
- l2
- l3

> ghi
> 
>> this is embedded quote:
>> - list inside of embedded quote: 1
>> - list inside of embedded quote: 2
> dddd

abc
`))
}
