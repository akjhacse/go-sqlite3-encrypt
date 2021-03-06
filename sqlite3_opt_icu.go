// Copyright (C) 2014 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// +build sqlite_icu icu

package sqlite3

/*
#cgo LDFLAGS: -licui18n -licuuc -licudata -lm -lpthread -lstdc++
#cgo CFLAGS: -DSQLITE_ENABLE_ICU -DU_STATIC_IMPLEMENTATION
#cgo darwin CFLAGS: -I/usr/local/opt/icu4c/include
#cgo darwin LDFLAGS: -L/usr/local/opt/icu4c/lib
#cgo openbsd LDFLAGS: -lsqlite3
*/
import "C"
