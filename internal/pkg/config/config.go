package config

import "time"

// MainAppPort - порт, который слушает главный сервер
const MainAppPort = 3000

// MainAppWriteTimeout - лимит на запись для главного сервер
const MainAppWriteTimeout = 15 * time.Second

// MainAppReadTimeout - лимит на чтение для главного сервер
const MainAppReadTimeout = 5 * time.Second

// Debug - дебаг
const Debug = true

// ContextTimeout - лимит на запрос в DB
const ContextTimeout = 2 * time.Second
