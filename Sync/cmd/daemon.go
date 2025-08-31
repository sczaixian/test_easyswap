package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ProjectsTask/Base/logger/xzap"
	"github.com/ProjectsTask/Sync/service"
	"github.com/ProjectsTask/Sync/service/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var DeamonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "sync easy swap order info",
	Long:  `sync easy swap order info`,
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		// rpc退出信号通知chan
		onSyncExit := make(chan error, 1)

		go func() {
			defer wg.Done()

			cfg, err := config.UnmarshalCmdConfig() // 读取和解析配置文件
			if err != nil {
				xzap.WithContext(ctx).Error("Failed to unmarshal config", zap.Error(err))
				onSyncExit <- err
				return
			}

			_, err = xzap.SetUp(*cfg.Log) // 初始化日志模块
			if err != nil {
				xzap.WithContext(ctx).Error("Failed to set up logger", zap.Error(err))
				onSyncExit <- err
				return
			}

			xzap.WithContext(ctx).Info("sync server start", zap.Any("config", cfg))

			s, err := service.New(ctx, cfg) // 初始化服务
			if err != nil {
				xzap.WithContext(ctx).Error("Failed to create sync server", zap.Error(err))
				onSyncExit <- err
				return
			}

			if err := s.Start(); err != nil { // 启动服务
				xzap.WithContext(ctx).Error("Failed to start sync server", zap.Error(err))
				onSyncExit <- err
				return
			}

			if cfg.Monitor.PprofEnable { // 开启pprof，用于性能监控
				http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.Monitor.PprofPort), nil)
			}
		}()

		onSignal := make(chan os.Signal)
		signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-onSignal:
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP:
				cancel()
				xzap.WithContext(ctx).Infof("Exit by signal", zap.String("signal", sig.String()))
			}
		case err := <-onSyncExit:
			cancel()
			xzap.WithContext(ctx).Infof("Exit by error", zap.Error(err))
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(DeamonCmd)
}
