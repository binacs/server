package gateway

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/binacs/server/config"
	"github.com/binacsgo/log"
)

// LogCleanerService log cleanup service interface
type LogCleanerService interface {
	Start() error
	Stop() error
	CleanOnce() error
	GetStats() map[string]interface{}
}

// LogCleanerServiceImpl log cleanup service implementation
type LogCleanerServiceImpl struct {
	Config    *config.Config `inject-name:"Config"`
	Logger    log.Logger     `inject-name:"Logger"`
	logConfig config.LogConfig
	stopChan  chan struct{}
	isRunning bool
}

func (lcs *LogCleanerServiceImpl) AfterInject() error {
	lcs.logConfig = lcs.Config.LogConfig
	lcs.stopChan = make(chan struct{})
	lcs.isRunning = false
	return nil
}

// Start start log cleanup service
func (lcs *LogCleanerServiceImpl) Start() error {
	if lcs.isRunning {
		return nil
	}

	lcs.isRunning = true
	lcs.Logger.Info("Log cleanup service started", "logPath", lcs.logConfig.File, "maxAge", lcs.logConfig.MaxAge)

	go lcs.run()
	return nil
}

// Stop stop log cleanup service
func (lcs *LogCleanerServiceImpl) Stop() error {
	if !lcs.isRunning {
		return nil
	}

	close(lcs.stopChan)
	lcs.isRunning = false
	lcs.Logger.Info("Log cleanup service stopped")
	return nil
}

// run run cleanup goroutine
func (lcs *LogCleanerServiceImpl) run() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	// Execute cleanup immediately on startup
	if err := lcs.CleanOnce(); err != nil {
		lcs.Logger.Error("Initial cleanup failed", "error", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := lcs.CleanOnce(); err != nil {
				lcs.Logger.Error("Periodic cleanup failed", "error", err)
			}
		case <-lcs.stopChan:
			return
		}
	}
}

// CleanOnce execute cleanup once
func (lcs *LogCleanerServiceImpl) CleanOnce() error {
	if lcs.logConfig.MaxAge <= 0 {
		return nil
	}

	lcs.Logger.Debug("Start cleaning expired log files", "logPath", lcs.logConfig.File, "maxAge", lcs.logConfig.MaxAge)

	logDir := path.Dir(lcs.logConfig.File)
	baseName := path.Base(lcs.logConfig.File)
	baseNameWithoutExt := strings.TrimSuffix(baseName, ".log")

	// Read log directory
	files, err := os.ReadDir(logDir)
	if err != nil {
		lcs.Logger.Error("Failed to read log directory", "error", err, "logDir", logDir)
		return err
	}

	cutoffTime := time.Now().AddDate(0, 0, -lcs.logConfig.MaxAge)
	deletedCount := 0
	totalSize := int64(0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		// Check if it's a log file (main log file or backup)
		if !strings.HasPrefix(fileName, baseNameWithoutExt) {
			continue
		}

		filePath := path.Join(logDir, fileName)
		fileInfo, err := file.Info()
		if err != nil {
			lcs.Logger.Warn("Failed to get file info", "error", err, "file", fileName)
			continue
		}

		// If file exceeds max age, delete it
		if fileInfo.ModTime().Before(cutoffTime) {
			if err := os.Remove(filePath); err != nil {
				lcs.Logger.Error("Failed to delete expired log file", "error", err, "file", fileName)
				continue
			}

			deletedCount++
			totalSize += fileInfo.Size()
			lcs.Logger.Info("Deleted expired log file", "file", fileName, "size", fileInfo.Size(), "modTime", fileInfo.ModTime())
		}
	}

	if deletedCount > 0 {
		lcs.Logger.Info("Log cleanup completed", "deletedCount", deletedCount, "totalSize", totalSize)
	} else {
		lcs.Logger.Debug("No log files need cleanup")
	}

	return nil
}

// GetStats get cleanup statistics
func (lcs *LogCleanerServiceImpl) GetStats() map[string]interface{} {
	logDir := path.Dir(lcs.logConfig.File)
	baseName := path.Base(lcs.logConfig.File)
	baseNameWithoutExt := strings.TrimSuffix(baseName, ".log")

	files, err := os.ReadDir(logDir)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	totalFiles := 0
	totalSize := int64(0)
	oldestFile := time.Now()
	newestFile := time.Time{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if !strings.HasPrefix(fileName, baseNameWithoutExt) {
			continue
		}

		fileInfo, err := file.Info()
		if err != nil {
			continue
		}

		totalFiles++
		totalSize += fileInfo.Size()

		if fileInfo.ModTime().Before(oldestFile) {
			oldestFile = fileInfo.ModTime()
		}
		if fileInfo.ModTime().After(newestFile) {
			newestFile = fileInfo.ModTime()
		}
	}

	return map[string]interface{}{
		"totalFiles": totalFiles,
		"totalSize":  totalSize,
		"oldestFile": oldestFile,
		"newestFile": newestFile,
		"maxAge":     lcs.logConfig.MaxAge,
		"isRunning":  lcs.isRunning,
	}
}
