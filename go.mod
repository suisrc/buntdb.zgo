module github.com/suisrc/buntdb.zgo

go 1.16

require (
	github.com/stretchr/testify v1.7.0
	github.com/suisrc/res.zgo v0.0.0
	github.com/tidwall/buntdb v1.2.0
)

replace github.com/suisrc/res.zgo v0.0.0 => ../res
