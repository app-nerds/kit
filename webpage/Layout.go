/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package webpage

/*
NewGoLayoutFromFile creates a new Mustache-based layout from a file
*/
func NewGoLayoutFromFile(fileName string) (ILayout, error) {
	result := &GoLayout{}
	err := result.LoadLayoutFile(fileName)

	return result, err
}

/*
NewGoLayoutFromString creates a new Mustache-based layout from a byte array
*/
func NewGoLayoutFromString(layout []byte) (ILayout, error) {
	result := &GoLayout{}
	err := result.LoadLayoutString(layout)

	return result, err
}

/*
NewMustacheLayoutFromFile creates a new Mustache-based layout from a file
*/
func NewMustacheLayoutFromFile(fileName string) (ILayout, error) {
	result := &MustacheLayout{}
	err := result.LoadLayoutFile(fileName)

	return result, err
}

/*
NewMustacheLayoutFromString creates a new Mustache-based layout from a byte array
*/
func NewMustacheLayoutFromString(layout []byte) (ILayout, error) {
	result := &MustacheLayout{}
	err := result.LoadLayoutString(layout)

	return result, err
}
