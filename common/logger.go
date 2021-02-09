package common

import "go.uber.org/zap"

//TestLog xxx
func TestLog() {
	log := zap.NewExample()
	log.Debug("This is a DEBUG message")
	log.Info("This is an INFO message")
	log.Info("This is an INFO message with fields", zap.Strings("region", []string{"us_west"}), zap.Int("id", 2))
	log.Warn("This is a WARN message")
	log.Error("This is an ERROR message")
	log.DPanic("This is a DPANIC message")
}
