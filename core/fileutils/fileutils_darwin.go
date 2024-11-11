/*
 Copyright 2016 Red Hat, Inc.

 Mercator is free software: you can redistribute it and/or modify
 it under the terms of the GNU Lesser General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 Mercator is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Lesser General Public License for more details.
 You should have received a copy of the GNU Lesser General Public License
 along with Mercator. If not, see <http://www.gnu.org/licenses/>.
*/
package fileutils

import "strings"

func IsHidden(path string) (bool, error) {
	return strings.Contains(path, "/."), nil
}
