/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package binary

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsString(t *testing.T) {
	tmpfile, err := os.CreateTemp(t.TempDir(), "")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	// Decode a fragment of kubectl into a temporary file:
	binData, err := base64.StdEncoding.DecodeString(kubectlFragment)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(tmpfile.Name(), binData, os.FileMode(0o644)))
	bin := Binary{
		options: &Options{
			Path: tmpfile.Name(),
		},
	}

	// The kubectl binary fragment is clipped where the tag "v1.20.2"
	// is located. The function should find it:
	cont, err := bin.ContainsStrings("v1.20.2")
	require.True(t, cont)
	require.NoError(t, err)

	// It should not, however, find a substring of the tag:
	cont, err = bin.ContainsStrings("1.20.2")
	require.False(t, cont)
	require.NoError(t, err)
}

var kubectlFragment = `nxsirlx0QAAAAAAA0HZAFANwVyHQekA7vuLSGA57QHEaitUNKXtAY+ef53SofUDqSbATP1Z+QGgo
7CEZK4RA97PI/X55hUACFbBWgMiFQO85+v5CLoZABGeTp8C4i0D///////+PQBhRnRjrAphA5jvf
zhnyo0BqJIxot/+oQB7FLgvj9rJAaUuYyn5qtECfyHUuMhK1QAAAAAAAiMNAER3/Jb8Vx0Dhka4+
lrfNQIWi5Wbutc9AZlLW1XK31ED7KgGmgjnXQHausXxFDtxAvvRkqnwx40AAAAAAAADwQAAAAAAA
avhAAAAAAICELkEAAAAAQHdLQQAAAADQEmNBAAAA/P//k0EAAAAAAACwQQAAAAAAAMBBAAAAAGXN
zUEAAACwjvArQgAAAMWFMYpCAAA0JvVrDEMAbCwrXJQpQwAAAAAAADBD////////P0MAAAAAAABA
QwAAAAAAAFBDAAAAAAAAkEMAAAAAAADgQwAAAAAAAPBDUO/i1uQaS0QAAADg///vRwAAAAAAAABI
AAAAAAAA0FJ9w5QlrUmyVP///////99/////////738AAAAAAADwfwAAAAAAAACAAQAAAAAAAIAf
l4qkysdQvMizP2k95MW9djx5Ne856r2V1iboCy4RvgAAAAAAADC+zj/Wc+fVM764BZFWAKx4vi3D
CW63/Yq+8WvSxUG9u76tXNJzp/7PviBhokJDnNC+aJYWusbF8L6sFgAS1ur4vrfbqp4ZzhS/6cYa
tqXmKL8P9URI5VVfvz8Hlgo4v2G/OHX3vrhZY7+Tvb4WbMFmv+RoZiORone/Sm/oORI0hL81ZA1g
EjSEv5f6bLHoTIy/TxnX21Eqnb/0EBERERGhv2KadO7y766/AAAAAAAAsL8f0MuZ1uSyv+woPj2Y
Y7y/QorDvLkZv78AAAAAAADAv4q8PBRmGMm/AAAAAAAA0L9niAEzw77SvzMzMzMzM9O/E7kcaX3N
1L/xw7j7QNLXvwAAAAAAAOC/g8jJbTBf5L/vOfr+Qi7mv2BzuuQWNOa/NMgyJd6R5r/ehcJwupPp
vwAAAGDBCuu/lCX3oX8A7L8AAAAAAADwvwAAAAAAAADA6CttR3yKC8BcwprG76AjwCbnsEEEHiXA
eoBrW1QoMMBamV9VCcIxwGItcULicDbAeQdU6D1OOsCNo8vkCjBPwI7sKP1pNlDAcwKINozAUsDy
0uRXZVJUwCW6BS2/uF7AmO3FQ10UZMBmIiiEsUxkwLsD0Py5tWXAsqvM61wTZ8A/ONybTjh+wCjy
dROI7IPAUTAt1RBJh8CSWS5qYQSQwAAAAAAAyJDA+y3aNmycpsAAAAAAACCswC3SBYaL8/XAT7tk
ut4/I8EAADQm9WsMw////////z/D////////7/8AAAAAAADw/wAAAQAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAABYAAAAAAAAAJgAAAAAAAAAAAAAAAAAAAP8AAAAAAAAA////////
//8CAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAAAAAAgAAAAAAAAACAAAAAAAAAAQAAAAAAAAA
CAAAAAAAAAAQAAAAAAAAAAEAAAAAAAAAAgAAAAAAAAAEAAAAAAAAAAgAAAAAAAAAEAAAAAAAAAAE
AAAAAAAAAAgAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJwAA
AAAAABAnAAAAAAAAAAAAAAAAAAAQJwAAAAAAABAnAAAAAAAAdjEuMjAuMgB2MS4yMC4yAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AD8AAAAAAAAAPwAAAAAAAAD9AAAAAAAAAD8AAAAAAAAAHAAAAAAAAAAPAAAAAAAAAAkAAAAAAAAA
ECcAAAAAAAAQJwAAAAAAABkAAAAAAAAANgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAW
AAAAAAAAAGEAAAAAAAAAYQAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYA
AAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAA
AAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAA
AAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAA
AAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAA
AAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAA
ABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAA
FgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAW
AAAAAAAAABYAAAAAAAAAFQAAAAAAAAARAAAAAAAAABYAAAAAAAAAJAAAAAAAAAAUAAAAAAAAABYA
AAAAAAAA/////wAAAAABAAAA/////0+8AO6qyNHMFAAAAAAAAAAkYc1jAQAAAAEAAAAAAAAAAQAA
AAAAAAABAAAAAAAAAAEAAAAAAAAAQJpDAAAAAAACAAAAAAAAAAEAAAAAAAAAFgAAAAAAAAAWAAAA
AAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAAGEAAAAAAAAAYQAAAAAAAAAWAAAAAAAAABYAAAAA
AAAAXwAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAABYAAAAAAAAAFgAAAAAA
AABhAAAAAAAAABYAAAAAAAAAFgAAAAAAAAAWAAAAAAAAAGEAAAAAAAAAIAAAAAAAAAAWAAAAAAAA
ABYAAAAAAAAAAgAAAAAAAAAgAAAAAAAAACAAAAAAAAAAAEBoOwAAAADa//////8PAP7//////w8A
AQAAAAgAAAAOAAAAAAAAAAEAAAAIAAAADgAAAAAAAAABAAAACAAAAA4AAAAAAAAAAQAAAAgAAAAD
AAAAAAAAAAEAAAAIAAAABwAAAAAAAAABAAAACAAAAAEAAAAAAAAAAQAAAAgAAAADAAAAAAAAAAEA
AAAIAAAAAwAAAAAAAAABAAAACAAAAAMAAAAAAAAAAQAAAAgAAAADAAAAAAAAAAEAAAAIAAAAAwAA
AAAAAAABAAAACAAAAAMAAAAAAAAAAQAAAAgAAAAHAAAAAAAAAAEAAAAIAAAABwAAAAAAAAABAAAA
CAAAAAcAAAAAAAAAAQAAAAgAAAAHAAAAAAAAAAEAAAAIAAAABwAAAAAAAAABAAAACAAAAAcAAAAA
AAAAAQAAAAgAAAAHAAAAAAAAAAEAAAAIAAAABwAAAAAAAAABAAAACAAAAAcAAAAAAAAAAQAAAAgA
AAAHAAAAAAAAAAEAAAAIAAAABwAAAAAAAAABAAAACAAAAAcAAAAAAAAAAQAAAAgAAAAHAAAAAAAA
AAEAAAAIAAAABwAAAAAAAAABAAAACAAAAAcAAAAAAAAAAQAAAAgAAAAHAAAAAAAAAAEAAAAIAAAA
BwAAAAAAAAABAAAACAAAAAcAAAAAAAAAAQAAAAgAAAAHAAAAAAAAAAEAAAAIAAAABwAAAAAAAAAB
AAAACAAAAAcAAAAAAAAAAQAAAAgAAAAHAAAAAAAAAAEAAAAIAAAABwAAAAAAAAABAAAACAAAAAcA
AAAAAAAAAQAAAAgAAAAHAAAAAAAAAAEAAAAIAAAABwAAAAAAAAABAAAACAAAAAcAAAAAAAAAAQAA
AAYAAAAAAAAAAAAAAAEAAAAEAAAAAQAAAAAAAAABAAAABAAAAAEAAAAAAAAAAQAAAAQAAAABAAAA
AAAAAAEAAAACAAAAAQAAAAAAAAABAAAAAgAAAAEAAAAAAAAAAQAAAAYAAAAHAAAAAAAAAAEAAAAE
AAAAAwAAAAAAAAABAAAACAAAAAMAAAAAAAAAAQAAAAoAAAATAAAAAAAAAAEAAAAK
`
