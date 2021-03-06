// Copyright (c) 2017 Sagar Gubbi. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const forgotpassSrc = `
{{ define "content" }}

<form action="/forgotpass" method="POST">
<input type="hidden" name="csrf" value={{ .Common.CSRF }}>
<table class="form">
	<tr>
		<th><label for="username">Username:</label></th>
		<td><input type="text" name="username" id="username" required></td>
	</tr>
{{ if .Common.Msg }}
	<tr>
		<th></th>
		<td><span class="alert">{{ .Common.Msg }}</span></td>
	</tr>
{{ end }}
	<tr>
		<th></th>
		<td><input type="submit" value="E-mail password reset link"></td>
	</tr>
</table>
</form>


{{ end }}`
