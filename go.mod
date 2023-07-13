module github.com/topxeq/qlang

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1007
	github.com/beevik/etree v1.1.0
	github.com/bhhbazinga/socks5-go v0.0.0-20191216170803-e6b0000a596a
	github.com/cavaliercoder/grab v2.0.0+incompatible
	github.com/dgraph-io/badger v1.6.1
	github.com/domodwyer/mailyak v3.1.1+incompatible
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807
	github.com/fogleman/gg v1.3.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jchv/go-webview2 v0.0.0-20221223143126-dc24628cff85
	// github.com/jchv/go-webview2 v0.0.0-20221223143126-dc24628cff85
	github.com/kbinani/screenshot v0.0.0-20210326165202-b96eb3309bb0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/nsf/termbox-go v1.1.1
	github.com/peterh/liner v1.2.1
	github.com/pterm/pterm v0.12.12
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/objx v0.5.0
	github.com/tinylib/msgp v1.1.5 // indirect
	github.com/topxeq/afero v0.0.0-20200914073911-38c8390e9ef4
	github.com/topxeq/awsapi v0.0.0-20191115074250-1192cb0fdb97
	github.com/topxeq/blockix v0.0.0-20210301072129-59538096fd95
	github.com/topxeq/charlang v0.0.0-20220224014453-a505ad57cf4e
	github.com/topxeq/dialog v0.0.0-20211124003827-315c3296b533
	github.com/topxeq/dlgs v0.0.0-20211122010615-d49596e82836
	github.com/topxeq/doc2vec v0.0.0-20200623133505-b167170c691e
	github.com/topxeq/docxrepl v0.0.0-20230223003559-0c82df798f9d
	github.com/topxeq/gods v0.0.0-20220125023913-b5718bb0c704
	github.com/topxeq/goph v0.0.0-20230116054750-120b087d86fb
	github.com/topxeq/imagetk v0.0.0-20230306082727-9b06565a8b58
	github.com/topxeq/regexpx v0.0.0-20200814082553-4bffc7d07029
	github.com/topxeq/socks v0.0.0-20200812112322-24acb126b5f3
	github.com/topxeq/sqltk v0.0.0-20230713014122-1cfca57c0814
	github.com/topxeq/text v0.0.0-20210710090213-e5fbc3508928
	github.com/topxeq/tk v1.0.1
	github.com/topxeq/xmlx v0.2.0
	github.com/visualfc/pkgwalk v1.0.0
	github.com/xuri/excelize/v2 v2.6.1
	gonum.org/v1/plot v0.12.0
)

// replace github.com/topxeq/text v0.0.0 => ../text

replace github.com/topxeq/tk v1.0.1 => ../tk

// replace github.com/topxeq/sqltk v0.0.0 => ../sqltk

// replace github.com/topxeq/charlang v0.0.0 => ../charlang

// replace github.com/topxeq/goph v0.0.0 => ../goph

// replace github.com/topxeq/go-sciter v0.0.0 => ../go-sciter

// replace github.com/topxeq/gods v0.0.0 => ../gods
