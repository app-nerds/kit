/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package passwords

/*
HashedPasswordString is a type of string that represents a hashed
password. This is often useful when dealing with passwords that must
be stored in a database.
*/
type HashedPasswordString string

func (hp HashedPasswordString) Set(password string) {
	result, _ := HashPassword(password)
	hp = HashedPasswordString(result)
}

func (hp HashedPasswordString) IsSameAsPlaintextPassword(plaintextPassword string) bool {
	return IsPasswordValid(string(hp), plaintextPassword)
}
