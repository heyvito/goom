#!/usr/bin/env roundup
export GOOMFILE=test/examples/data.json
goom="./dist/goom"

describe "cli"

it_shows_help() {
  $goom help | grep "goom is a simple kv storage based on Boom, by Zach Holman"
}

it_shows_a_version() {
  $goom --version | grep "goom version"
}
