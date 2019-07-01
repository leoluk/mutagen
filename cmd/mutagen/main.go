package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/havoc-io/mutagen/cmd"
	"github.com/havoc-io/mutagen/cmd/mutagen/daemon"
	"github.com/havoc-io/mutagen/cmd/mutagen/sync"
	"github.com/havoc-io/mutagen/pkg/prompt"

	// Explicitly import packages that need to register protocol handlers.
	_ "github.com/havoc-io/mutagen/pkg/session/protocols/docker"
	_ "github.com/havoc-io/mutagen/pkg/session/protocols/local"
	_ "github.com/havoc-io/mutagen/pkg/session/protocols/ssh"
)

func rootMain(command *cobra.Command, arguments []string) error {
	// If no commands were given, then print help information and bail. We don't
	// have to worry about warning about arguments being present here (which
	// would be incorrect usage) because arguments can't even reach this point
	// (they will be mistaken for subcommands and a error will be displayed).
	command.Help()

	// Success.
	return nil
}

var rootCommand = &cobra.Command{
	Use:          "mutagen",
	Short:        "Mutagen provides simple, continuous, bi-directional file synchronization.",
	RunE:         rootMain,
	SilenceUsage: true,
}

var rootConfiguration struct {
	// help indicates whether or not help information should be shown for the
	// command.
	help bool
}

func init() {
	// Disable alphabetical sorting of commands in help output. This is a global
	// setting that affects all Cobra command instances.
	cobra.EnableCommandSorting = false

	// Grab a handle for the command line flags.
	flags := rootCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&rootConfiguration.help, "help", "h", false, "Show help information")

	// Disable Cobra's use of mousetrap. This breaks daemon registration on
	// Windows because it tries to enforce that the CLI only be launched from
	// a console, which it's not when running automatically.
	cobra.MousetrapHelpText = ""

	// Register commands.
	// HACK: Add the legacy sync commands for temporary backward compatibility.
	commands := []*cobra.Command{
		sync.RootCommand,
		daemon.RootCommand,
		versionCommand,
		legalCommand,
		generateCommand,
	}
	commands = append(commands, sync.LegacyCommands...)
	rootCommand.AddCommand(commands...)
}

func main() {
	// Check if a prompting environment is set. If so, treat this as a prompt
	// request. Prompting is sort of a special pseudo-command that's indicated
	// by the presence of an environment variable, and hence it has to be
	// handled in a bit of a special manner.
	if _, ok := os.LookupEnv(prompt.PrompterEnvironmentVariable); ok {
		if err := promptMain(os.Args[1:]); err != nil {
			cmd.Fatal(err)
		}
		return
	}

	// Handle terminal compatibility issues. If this call returns, it means that
	// we should proceed normally.
	cmd.HandleTerminalCompatibility()

	// HACK: Modify the root command's help function to hide legacy commands.
	defaultHelpFunction := rootCommand.HelpFunc()
	rootCommand.SetHelpFunc(func(command *cobra.Command, arguments []string) {
		if command == rootCommand {
			for _, command := range sync.LegacyCommands {
				command.Hidden = true
			}
		}
		defaultHelpFunction(command, arguments)
	})

	// Execute the root command.
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
