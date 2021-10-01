module github.com/Assetsadapter/standardchain-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.1
	github.com/blocktree/go-owcrypt v1.1.2
	github.com/blocktree/openwallet/v2 v2.0.7
	github.com/blocktree/quorum-adapter v1.4.1
	github.com/ethereum/go-ethereum v1.9.9
	github.com/influxdata/influxdb v1.2.3-0.20180221223340-01288bdb0883
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876
	gotest.tools v2.2.0+incompatible
)

//replace github.com/blocktree/quorum-adapter => ../quorum-adapter

//replace github.com/blocktree/openwallet/v2 => ../../openwallet
