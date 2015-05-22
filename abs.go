// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package abs

func SignAbsi(n int) (int, int) {
	if n > 0 {
		return 1, n
	} else if n < 0 {
		return -1, -n
	} else {
		return 0, 0
	}
}
func SignAbsf(n float64) (float64, float64) {
	if n > 0 {
		return 1, n
	} else if n < 0 {
		return -1, -n
	} else {
		return 0, 0
	}
}
func Signi(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	} else {
		return 0
	}
}
func Signf(n float64) float64 {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	} else {
		return 0
	}
}
func Absi(n int) int {
	if n > 0 {
		return n
	} else if n < 0 {
		return -n
	} else {
		return 0
	}
}
func Absf(n float64) float64 {
	if n > 0 {
		return n
	} else if n < 0 {
		return -n
	} else {
		return 0
	}
}
