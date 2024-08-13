package models

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	tealog "github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"
)

const (
	host = "localhost"
	port = "6060"
)

type server struct {
	*ssh.Server
}

func CreateServer() (*server, error) {
	svr := new(server)

	sshServer, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.MiddlewareWithProgramHandler(svr.ProgramHandler, termenv.ANSI256),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)

	if err != nil {
		return svr, fmt.Errorf("could not start server, error %w", err)
	}

	svr.Server = sshServer
	return svr, nil
}

func (svr *server) Start() {
	var err error

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	tealog.Info("Starting SSH server", "host", host, "port", port)

	go func() {
		if err = svr.ListenAndServe(); err != nil {
			tealog.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done

	tealog.Info("Stopping SSH server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		tealog.Error("Could not stop server", "error", err)
	}
}

func (svr *server) ProgramHandler(session ssh.Session) *tea.Program {
	model := NewModel()
	model.server = svr
	model.id = session.User()
	hash := sha256.New()

	if _, err := hash.Write(session.PublicKey().Marshal()); err != nil {
		tealog.Fatal(err)
	}

	tealog.Infof("session public key %s\n", string(session.PublicKey().Marshal()))
	tealog.Infof("%x\n", hash.Sum(nil))

	return tea.NewProgram(model, bubbletea.MakeOptions(session)...)
}
