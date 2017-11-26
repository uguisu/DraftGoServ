package Config

// Server configaration
type ServerConfig struct {
    port      string
}

var _s ServerConfig

func GetPort() string {
    return ":" + port
}

func LoadConfig() <-chan bool {

    // Declare a channel for loading config
    configFinished = make(chan bool)

    go func() {
        _s.port = "8080"
        configFinished <- true
        close(configFinished)
    }()

    return configFinished
}
