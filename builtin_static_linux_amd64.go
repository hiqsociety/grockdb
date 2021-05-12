// +build builtin_static !static

package grocksdb

// #cgo CFLAGS: -I${SRCDIR}/dist/linux_amd64/include
// #cgo CXXFLAGS: -I${SRCDIR}/dist/linux_amd64/include
// #cgo LDFLAGS: -L${SRCDIR}/dist/linux_amd64/lib -L${SRCDIR}/dist/linux_amd64/lib64 -lterarkdb -pthread -lstdc++ -ldl -lzstd -llz4 -lz -lsnappy -lm

import "C"
