package config

import (
//    "net/http"
    "log"
//    Controller "github.com/uguisu/controller"
)

// Server configaration
type ServerConfig struct {
    port      string
}

var _s ServerConfig

func GetPort() string {
    return ":" + _s.port
}

/**
 * Load Url mappings
 */
func loadUrlMappings() <-chan bool {
    // Declare a channel for loading config
    loadUrlMappingConfigFinished := make(chan bool)

    go func() {

        log.Println("Execute loadUrlMappings() ...")

//        http.HandleFunc("/hello", Controller.HelloServer)
        loadUrlMappingConfigFinished <- true
        close(loadUrlMappingConfigFinished)
    }()

    return loadUrlMappingConfigFinished
}

/**
 * Load Server settings
 */
func loadServerConfig() <-chan bool {
    // Declare a channel for loading config
    loadServerConfigFinished := make(chan bool)

    go func() {

        log.Println("Execute loadServerConfig() ...")

        _s.port = "8080"
        loadServerConfigFinished <- true
        close(loadServerConfigFinished)
    }()

    return loadServerConfigFinished
}

/**
 * Load Config
 */
func LoadConfig() <-chan bool {

    // Declare a channel for loading config
    configFinished := make(chan bool)

    loadUrlMappingConfigFinished := loadUrlMappings()
    loadServerConfigFinished     := loadServerConfig()

    go func() {
        <-loadUrlMappingConfigFinished
        <-loadServerConfigFinished
        configFinished <- true
        close(configFinished)
    }()

    return configFinished
}
