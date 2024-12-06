package config

func Init() {
    random_number := 20
    logger := GetLogger("config")
    logger.Infof("CONFIG INITIALIZED: %d", random_number)
}

//Get new log instace
func GetLogger(p string) *Logger {
    logger := NewLogger(p)
    return logger
}
