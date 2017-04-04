#!/usr/bin/env roundup
export GOOMFILE=test/examples/data.json
goom="./dist/goom"

describe "items"

it_adds_an_item() {
  $goom urls google 'http://google.com'
  $goom urls | grep google.com
}

it_deletes_an_item() {
  yes | $goom rm-item urls google | grep 'Removed google from urls'
  $goom urls google | grep "isn't present"
}

it_echos_an_item() {
  $goom echo site | grep 'vito.io'
}
