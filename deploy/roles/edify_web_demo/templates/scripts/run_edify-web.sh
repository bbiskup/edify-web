#!/bin/bash

echo "****** Starting edify-web ******"
[[ -s "{{demo_home}}/.gvm/scripts/gvm" ]] && source "{{demo_home}}/.gvm/scripts/gvm"
env
gvm use {{go_version}}
exec edify-web
