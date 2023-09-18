package cmd

import (
	logger2 "github.com/mohammadne/phone-book/pkg/logger"
	"github.com/peyman1987/phonebook/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

type Server struct {
}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run phonebook server",
		Run:   run,
	}
}

func (cmd *Server) run(cfg *config.Config, trap chan os.Signal) {
	logger := logger2.NewZap(cfg.Logger)

	logger.Error("error message")

	field := zap.String("signal trap", (<-trap).String())
	logger.Info("exiting by recieving unix signal", field)

}
