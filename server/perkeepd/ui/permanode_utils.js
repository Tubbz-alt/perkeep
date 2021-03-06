/*
Copyright 2014 The Perkeep Authors

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

goog.provide('cam.permanodeUtils');

goog.require('goog.array');

cam.permanodeUtils.getSingleAttr = function(permanode, name) {
	var val = permanode.attr[name];
	if (val) {
		return goog.isArray(val) ? val[0] : val;
	}
	return null;
};

cam.permanodeUtils.isContainer = function(permanode) {
	return goog.object.some(permanode.attr, function(v, k) { return k == 'camliMember' || goog.string.startsWith(k, 'camliPath:'); });
};

cam.permanodeUtils.getCamliNodeType = function(permanode) {
	return cam.permanodeUtils.getSingleAttr(permanode, 'camliNodeType');
};
