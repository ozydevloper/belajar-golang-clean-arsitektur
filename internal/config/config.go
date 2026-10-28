package config

import "time"

type ConfigApp struct {
	Name         string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type ConfigDB struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DBName       string
	SSL          string
	OpenConn     int
	IdleConn     int
	Lifetime     time.Duration
	IdleLifetime time.Duration
}

type ConfigLogger struct {
	Level int
}

type Config struct {
	Log ConfigLogger
	DB  ConfigDB
	App ConfigApp
}
