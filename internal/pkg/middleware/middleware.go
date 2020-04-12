/*
 * Copyright (C) 2020. Leonid Moguchev
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package middleware

// Middleware represent the data-struct for middleware
type Middleware struct {
	// another stuff , may be needed by middleware
}

// InitMiddleware intialize the middleware
func InitMiddleware() *Middleware {
	return &Middleware{}
}
