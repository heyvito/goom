#!/usr/bin/env roundup
export GOOMFILE=test/examples/data.json
goom="./dist/goom"

describe "lists"

it_shows_all_lists_by_default() {
  $goom | grep "urls"
  $goom | grep "jokes"
}

it_shows_a_list() {
  $goom urls | grep 'vito.io'
}

it_creates_a_list() {
  $goom enemies James Moriarty | grep 'Moriarty'
}

it_deletes_a_list() {
  $goom | grep "enemies"
  $goom rm-group enemies | grep "enemies is no more!"
  ! $goom | grep "enemies"
}

it_handles_delete_on_nonexistent_list() {
  ! $goom | grep "enemies"
  $goom rm-group "enemies" | grep "enemies is no more!"
}

it_handles_empty_goomfile() {
  cp /dev/null test/examples/empty.json
  export BOOMFILE=test/examples/empty.json
  $goom test item test | grep "test"
  export BOOMFILE=test/examples/data.json
}
