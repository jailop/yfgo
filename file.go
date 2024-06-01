package main

import (
    "os"
    "path"
)

func FilePath(filename string) (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        println("User home directory not detected")
        return "", err
    }
    configDir := path.Join(homeDir, ".config")
    _, err = os.Stat(configDir)
    if err != nil {
        if err = os.Mkdir(configDir, os.ModePerm); err != nil {
            println("Config directory couldn't be created")
            return "", err
        }
    }
    appDir := path.Join(configDir, "yfgo")
    _, err = os.Stat(appDir)
    if err != nil {
        if err = os.Mkdir(appDir, os.ModePerm); err != nil {
            println("App directory couldn't be created")
            return "", err
        }
    }
    return path.Join(appDir, filename), nil
}

func FileExists(filePath string) bool {
    _, err := os.Stat(filePath)
    if err != nil {
        println("File doesn't exist")
        return false
    }
    return true
}
