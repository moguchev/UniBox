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

package models

import (
	"fmt"
	"time"
)

// ErrorType -
type ErrorType uint

// CtxKey - ctx key
type CtxKey string

const (
	// NoType - Unknown error
	NoType = ErrorType(iota)
	// BadRequest -
	BadRequest
	// NotFound -
	NotFound
	// AlreadyExists -
	AlreadyExists
	// Invalid -
	Invalid
	// Internal -
	Internal
)

// ReqIDKey -
const ReqIDKey = "rID"

// ErrorMessage -
type ErrorMessage struct {
	Target      string            `json:"target"`
	Message     string            `json:"message"`
	ContextInfo map[string]string `json:"internal"`
}

// Error - own realisation
type Error struct {
	Type        ErrorType
	Original    error
	Target      string
	Message     string
	ContextInfo map[string]string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Target, e.Message)
}

// User - пользователь
type User struct {
	Email            string    `json:"email"             db:"email"`
	Nickname         string    `json:"nickname"          db:"nickname"`
	Firstname        string    `json:"firstname"         db:"firstname"`
	Lastname         string    `json:"lastname"          db:"lastname"`
	RegistrationTime time.Time `json:"registration_time" db:"registration_time"`
}

// NewUser - новый пользователь
type NewUser struct {
	Email     string `json:"email"     db:"email"`
	Nickname  string `json:"nickname"  db:"nickname"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname"  db:"lastname"`
	Password  string `json:"password"  db:"password_digest"`
}
