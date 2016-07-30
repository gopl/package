#!/bin/bash
export PS4='+[$(date +"[%F %T]") $LINENO:${FUNCNAME[0]}:$( basename "${BASH_SOURCE[0]}")] '
set -x

exit $1

