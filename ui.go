/*
   Copyright 2017 OÜ Nevermore

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
package meelel

import (
	"bytes"
	"html/template"
)

func (p *Post) HTML() string {
	var b bytes.Buffer
	b.WriteString("<div>")

	b.WriteString("<h2>")
	b.WriteString(template.HTMLEscapeString(p.Title))
	b.WriteString("</h2>")

	b.WriteString(template.HTMLEscapeString(p.Content))

	b.WriteString("</div>")

	return b.String()
}
