package qall

import (
	"github.com/topxeq/qlang/lib/bufio"
	"github.com/topxeq/qlang/lib/bytes"
	"github.com/topxeq/qlang/lib/crypto/md5"
	"github.com/topxeq/qlang/lib/encoding/hex"
	"github.com/topxeq/qlang/lib/encoding/json"
	"github.com/topxeq/qlang/lib/eqlang"
	"github.com/topxeq/qlang/lib/errors"
	"github.com/topxeq/qlang/lib/io"
	"github.com/topxeq/qlang/lib/io/ioutil"
	"github.com/topxeq/qlang/lib/math"
	"github.com/topxeq/qlang/lib/meta"
	"github.com/topxeq/qlang/lib/net/http"
	"github.com/topxeq/qlang/lib/os"
	"github.com/topxeq/qlang/lib/path"
	"github.com/topxeq/qlang/lib/reflect"
	"github.com/topxeq/qlang/lib/runtime"
	"github.com/topxeq/qlang/lib/strconv"
	"github.com/topxeq/qlang/lib/strings"
	"github.com/topxeq/qlang/lib/sync"
	"github.com/topxeq/qlang/lib/terminal"
	"github.com/topxeq/qlang/lib/tpl/extractor"
	"github.com/topxeq/qlang/lib/version"
	qlang "github.com/topxeq/qlang/spec"

	// qlang builtin modules
	_ "github.com/topxeq/qlang/lib/builtin"
	_ "github.com/topxeq/qlang/lib/chan"
)

// -----------------------------------------------------------------------------

// Copyright prints qlang copyright information.
//
func Copyright() {
	version.Copyright()
}

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("", math.Exports) // import math as builtin package
	qlang.Import("", meta.Exports) // import meta package
	qlang.Import("bufio", bufio.Exports)
	qlang.Import("bytes", bytes.Exports)
	qlang.Import("md5", md5.Exports)
	qlang.Import("io", io.Exports)
	qlang.Import("ioutil", ioutil.Exports)
	qlang.Import("hex", hex.Exports)
	qlang.Import("json", json.Exports)
	qlang.Import("errors", errors.Exports)
	qlang.Import("eqlang", eqlang.Exports)
	qlang.Import("math", math.Exports)
	qlang.Import("os", os.Exports)
	qlang.Import("", os.InlineExports)
	qlang.Import("path", path.Exports)
	qlang.Import("http", http.Exports)
	qlang.Import("reflect", reflect.Exports)
	qlang.Import("runtime", runtime.Exports)
	qlang.Import("strconv", strconv.Exports)
	qlang.Import("strings", strings.Exports)
	qlang.Import("sync", sync.Exports)
	qlang.Import("terminal", terminal.Exports)
	qlang.Import("extractor", extractor.Exports)
}

// -----------------------------------------------------------------------------
