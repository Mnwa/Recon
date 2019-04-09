package database

import "github.com/prologic/bitcask"

var Client, _ = bitcask.Open("/tmp/db")
