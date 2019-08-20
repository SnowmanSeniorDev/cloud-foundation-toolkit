package launchpad
// WARNING: Generated file, do not modify directly!

var statics = map[string]string {
	"static/tmpl/tf/_output.tf.tmpl": `output "{{ .Id }}" {
  value = "{{ .Val }}"
}
`,
	"static/tmpl/tf/_variable.tf.tmpl": `variable "{{ .Id }}" {
  description = "{{ .Description }}"
  default     = "{{ .Default }}"
}
`,
	"static/tmpl/tf/google_folder.tf.tmpl": `resource "google_folder" "{{ .Id }}" {
  display_name = "{{ .DisplayName }}"
  parent       = "{{ .Parent }}"
}
`,
	"static/tmpl/tf/google_provider.tf.tmpl": `provider "google" {
  credentials = "{{ .Credentials }}"
  version     = "{{ .Version }}"
}
`,
	"static/tmpl/tf/license.tf.tmpl": `/**
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

`,
}
