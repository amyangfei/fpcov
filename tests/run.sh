#!/bin/bash

TEST_DIR="/tmp/failpoint_test"
CUR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

mkdir -p $TEST_DIR

$CUR/../bin/goodbye.test -test.coverprofile="$TEST_DIR/goodbye.$(date +"%s").out" DEVEL >> $TEST_DIR/goodbye.log 2>&1
