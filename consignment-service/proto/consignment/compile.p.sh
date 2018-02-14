#!/bin/bash
protoc -I.  --go_out=plugins=micro:./ \
 consignment.proto
