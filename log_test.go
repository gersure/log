// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	. "github.com/pingcap/check"
)

var _ = Suite(&testLogSuite{})

type testLogSuite struct{}

func (t *testLogSuite) TestExport(c *C) {
	conf := &Config{Level: "debug", File: FileLogConfig{}, DisableTimestamp: true}
	lg := newZapTestLogger(conf, c)
	ReplaceGlobals(lg.Logger, nil)
	Info("Testing")
	Debug("Testing")
	Warn("Testing")
	Error("Testing")
	lg.AssertContains("log_test.go:")
}

//func (t *testLogSuite) TestLogPrint(c *C) {
//	config := &Config{
//		Level:               "info",
//		Format:              "text",
//		DisableTimestamp:    false,
//		File:                FileLogConfig{},
//		Development:         false,
//		DisableCaller:       false,
//		DisableStacktrace:   true,
//		DisableErrorVerbose: false,
//		Sampling:            nil,
//	}
//	gl, props, err := InitLogger(config)
//	if err != nil {
//		c.Error(err)
//	}
//	gl = gl.With(zap.Int32("conn", 123))
//	ReplaceGlobals(gl, props)
//	Info("Info Testing")
//	Debug("Debug Testing")
//	Warn("Warn Testing")
//	Error("Error Testing")
//}
