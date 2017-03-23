package tomlConfig

import (
    "testing"
)

type Config struct{
	Pizza string
}

func TestLoad(t *testing.T){
    filename := "settings.default.toml"
    // Test Creating a file
    err := Create(filename, Config{})
    if err != nil{
        t.Error(err)
    }

    // Test Loading this file
    var conf Config
    err = Load([]string{filename,}, &conf)
    if err != nil{
        t.Error(err)
    }
}

func TestLoadMissingFile(t *testing.T){
    // This file should not exist... Hopefully
    filename := "faszxcvuhinwerfrwezcvnmop.toml"
    var conf Config
    err := Load([]string{filename,}, &conf)
    if err == nil{
        t.Errorf("A file loaded should not exist.")
    }
}
