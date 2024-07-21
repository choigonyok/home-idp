package cmd

import (
	"github.com/choigonyok/home-idp/idpctl/pkg/cli"
	"github.com/spf13/cobra"
)

const (
	// Location to read istioctl defaults from
	defaultIstioctlConfig = "$HOME/.idpctl/config.yaml"
)

const (
	FlagCharts = "charts"
)

func Test() int {
	return 1
}

// // ConfigAndEnvProcessing uses spf13/viper for overriding CLI parameters
// func ConfigAndEnvProcessing() error {
// 	configPath := filepath.Dir(root.IstioConfig)
// 	baseName := filepath.Base(root.IstioConfig)
// 	configType := filepath.Ext(root.IstioConfig)
// 	configName := baseName[0 : len(baseName)-len(configType)]
// 	if configType != "" {
// 		configType = configType[1:]
// 	}

// 	// Allow users to override some variables through $HOME/.istioctl/config.yaml
// 	// and environment variables.
// 	viper.SetEnvPrefix("ISTIOCTL")
// 	viper.AutomaticEnv()
// 	viper.AllowEmptyEnv(true) // So we can say ISTIOCTL_CERT_DIR="" to suppress certs
// 	viper.SetConfigName(configName)
// 	viper.SetConfigType(configType)
// 	viper.AddConfigPath(configPath)
// 	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
// 	err := viper.ReadInConfig()
// 	// Ignore errors reading the configuration unless the file is explicitly customized
// 	if root.IstioConfig != defaultIstioctlConfig {
// 		return err
// 	}

// 	return nil
// }

// func init() {
// 	viper.SetDefault("istioNamespace", constants.IstioSystemNamespace)
// 	viper.SetDefault("xds-port", 15012)
// }

// GetRootCmd returns the root of the cobra command-tree.
func GetRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "idpctl",
		Short:             "Home-idp Command Line Interface",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		// PersistentPreRunE: ConfigureLogging,
		Long: `Command line interface to deploy home-idp application in kubernetes cluster`,
	}

	rootCmd.SetArgs(args)
	rootFlags := rootCmd.PersistentFlags()
	cli.AddRootFlags(rootFlags)
	rootCmd.Flags().AddFlagSet(rootFlags)

	// ctx := cli.NewCLIContext(rootOptions)

	installCmd := GetInstallCmd()
	// hideInheritedFlags(installCmd, cli.:, cli.FlagIstioNamespace, FlagCharts)
	rootCmd.AddCommand(installCmd)

	statusCmd := GetStatusCmd()
	// hideInheritedFlags(installCmd, cli.:, cli.FlagIstioNamespace, FlagCharts)
	rootCmd.AddCommand(statusCmd)

	return rootCmd
}

// func test(ctx cli.Context) *cobra.Command {
func GetInstallCmd() *cobra.Command {
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Install home-idp app",
		Long:  "Install home-idp application in kubernetes cluster with CLI",
	}
	return installCmd
}

func GetStatusCmd() *cobra.Command {
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "show status home-idp app",
		Long:  "Show status of home-idp application deployed in kubernetes cluster with CLI",
	}
	return statusCmd
}

func hideInheritedFlags(orig *cobra.Command, hidden ...string) {
	orig.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		for _, hidden := range hidden {
			_ = cmd.Flags().MarkHidden(hidden) // nolint: errcheck
		}

		orig.SetHelpFunc(nil)
		orig.HelpFunc()(cmd, args)
	})
}

// func ManifestGenerateCmd(ctx cli.Context, rootArgs *RootArgs, mgArgs *ManifestGenerateArgs) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "generate",
// 		Short: "Generates an Istio install manifest",
// 		Long:  "The generate subcommand generates an Istio install manifest and outputs to the console by default.",
// 		// nolint: lll
// 		Example: `  # Generate a default Istio installation
//   istioctl manifest generate

//   # Enable Tracing
//   istioctl manifest generate --set meshConfig.enableTracing=true

//   # Generate the demo profile
//   istioctl manifest generate --set profile=demo

//   # To override a setting that includes dots, escape them with a backslash (\).  Your shell may require enclosing quotes.
//   istioctl manifest generate --set "values.sidecarInjectorWebhook.injectedAnnotations.container\.apparmor\.security\.beta\.kubernetes\.io/istio-proxy=runtime/default"
// `,
// 		Args: func(cmd *cobra.Command, args []string) error {
// 			if len(args) != 0 {
// 				return fmt.Errorf("generate accepts no positional arguments, got %#v", args)
// 			}
// 			return nil
// 		},
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			if kubeClientFunc == nil {
// 				kubeClientFunc = ctx.CLIClient
// 			}
// 			var kubeClient kube.CLIClient
// 			if mgArgs.EnableClusterSpecific {
// 				kc, err := kubeClientFunc()
// 				if err != nil {
// 					return err
// 				}
// 				kubeClient = kc
// 			}
// 			l := clog.NewConsoleLogger(cmd.OutOrStdout(), cmd.ErrOrStderr(), installerScope)
// 			return ManifestGenerate(kubeClient, rootArgs, mgArgs, l)
// 		},
// 	}
// }

// _ = rootCmd.RegisterFlagCompletionFunc(cli.FlagNamespace, func(
// 	cmd *cobra.Command, args []string, toComplete string,
// ) ([]string, cobra.ShellCompDirective) {
// 	return completion.ValidNamespaceArgs(cmd, ctx, args, toComplete)
// })

// 	// Attach the Istio logging options to the command.
// 	root.LoggingOptions.AttachCobraFlags(rootCmd)
// 	hiddenFlags := []string{
// 		"log_as_json", "log_rotate", "log_rotate_max_age", "log_rotate_max_backups",
// 		"log_rotate_max_size", "log_stacktrace_level", "log_target", "log_caller", "log_output_level",
// 	}
// 	for _, opt := range hiddenFlags {
// 		_ = rootCmd.PersistentFlags().MarkHidden(opt)
// 	}

// 	cmd.AddFlags(rootCmd)

// 	kubeInjectCmd := kubeinject.InjectCommand(ctx)
// 	hideInheritedFlags(kubeInjectCmd, cli.FlagNamespace)
// 	rootCmd.AddCommand(kubeInjectCmd)

// 	experimentalCmd := &cobra.Command{
// 		Use:     "experimental",
// 		Aliases: []string{"x", "exp"},
// 		Short:   "Experimental commands that may be modified or deprecated",
// 	}

// 	xdsBasedTroubleshooting := []*cobra.Command{
// 		// TODO(hanxiaop): I think experimental version still has issues, so we keep the old version for now.
// 		version.XdsVersionCommand(ctx),
// 		// TODO(hanxiaop): this is kept for some releases in case someone is using it.
// 		proxystatus.XdsStatusCommand(ctx),
// 	}
// 	troubleshootingCommands := []*cobra.Command{
// 		version.NewVersionCommand(ctx),
// 		proxystatus.StableXdsStatusCommand(ctx),
// 	}
// 	var debugCmdAttachmentPoint *cobra.Command
// 	if viper.GetBool("PREFER-EXPERIMENTAL") {
// 		legacyCmd := &cobra.Command{
// 			Use:   "legacy",
// 			Short: "Legacy command variants",
// 		}
// 		rootCmd.AddCommand(legacyCmd)
// 		for _, c := range xdsBasedTroubleshooting {
// 			rootCmd.AddCommand(c)
// 		}
// 		debugCmdAttachmentPoint = legacyCmd
// 	} else {
// 		debugCmdAttachmentPoint = rootCmd
// 	}
// 	for _, c := range xdsBasedTroubleshooting {
// 		experimentalCmd.AddCommand(c)
// 	}
// 	for _, c := range troubleshootingCommands {
// 		debugCmdAttachmentPoint.AddCommand(c)
// 	}

// 	rootCmd.AddCommand(experimentalCmd)
// 	rootCmd.AddCommand(proxyconfig.ProxyConfig(ctx))
// 	rootCmd.AddCommand(admin.Cmd(ctx))
// 	experimentalCmd.AddCommand(injector.Cmd(ctx))

// 	rootCmd.AddCommand(mesh.NewVerifyCommand(ctx))
// 	rootCmd.AddCommand(mesh.UninstallCmd(ctx))

// 	experimentalCmd.AddCommand(authz.AuthZ(ctx))

// 	experimentalCmd.AddCommand(metrics.Cmd(ctx))
// 	experimentalCmd.AddCommand(describe.Cmd(ctx))
// 	experimentalCmd.AddCommand(config.Cmd())
// 	experimentalCmd.AddCommand(workload.Cmd(ctx))
// 	experimentalCmd.AddCommand(internaldebug.DebugCommand(ctx))
// 	experimentalCmd.AddCommand(precheck.Cmd(ctx))
// 	experimentalCmd.AddCommand(proxyconfig.StatsConfigCmd(ctx))
// 	experimentalCmd.AddCommand(checkinject.Cmd(ctx))
// 	rootCmd.AddCommand(waypoint.Cmd(ctx))
// 	rootCmd.AddCommand(ztunnelconfig.ZtunnelConfig(ctx))

// 	analyzeCmd := analyze.Analyze(ctx)
// 	hideInheritedFlags(analyzeCmd, cli.FlagIstioNamespace)
// 	rootCmd.AddCommand(analyzeCmd)

// 	dashboardCmd := dashboard.Dashboard(ctx)
// 	hideInheritedFlags(dashboardCmd, cli.FlagNamespace, cli.FlagIstioNamespace)
// 	rootCmd.AddCommand(dashboardCmd)

// 	manifestCmd := mesh.ManifestCmd(ctx)
// 	hideInheritedFlags(manifestCmd, cli.FlagNamespace, cli.FlagIstioNamespace, FlagCharts)
// 	rootCmd.AddCommand(manifestCmd)

// 	operatorCmd := mesh.OperatorCmd(ctx)
// 	hideInheritedFlags(operatorCmd, cli.FlagNamespace, cli.FlagIstioNamespace, FlagCharts)
// 	rootCmd.AddCommand(operatorCmd)

// 	installCmd := mesh.InstallCmd(ctx)
// 	hideInheritedFlags(installCmd, cli.FlagNamespace, cli.FlagIstioNamespace, FlagCharts)
// 	rootCmd.AddCommand(installCmd)

// 	profileCmd := mesh.ProfileCmd(ctx)
// 	hideInheritedFlags(profileCmd, cli.FlagNamespace, cli.FlagIstioNamespace, FlagCharts)
// 	rootCmd.AddCommand(profileCmd)

// 	upgradeCmd := mesh.UpgradeCmd(ctx)
// 	hideInheritedFlags(upgradeCmd, cli.FlagNamespace, cli.FlagIstioNamespace, FlagCharts)
// 	rootCmd.AddCommand(upgradeCmd)

// 	bugReportCmd := bugreport.Cmd(ctx, root.LoggingOptions)
// 	hideInheritedFlags(bugReportCmd, cli.FlagNamespace, cli.FlagIstioNamespace)
// 	rootCmd.AddCommand(bugReportCmd)

// 	tagCmd := tag.TagCommand(ctx)
// 	hideInheritedFlags(tag.TagCommand(ctx), cli.FlagNamespace, cli.FlagIstioNamespace, FlagCharts)
// 	rootCmd.AddCommand(tagCmd)

// 	// leave the multicluster commands in x for backwards compat
// 	rootCmd.AddCommand(multicluster.NewCreateRemoteSecretCommand(ctx))
// 	rootCmd.AddCommand(proxyconfig.ClustersCommand(ctx))

// 	rootCmd.AddCommand(collateral.CobraCommand(rootCmd, collateral.Metadata{
// 		Title:   "Istio Control",
// 		Section: "istioctl CLI",
// 		Manual:  "Istio Control",
// 	}))

// 	validateCmd := validate.NewValidateCommand(ctx)
// 	hideInheritedFlags(validateCmd, "kubeconfig")
// 	rootCmd.AddCommand(validateCmd)

// 	rootCmd.AddCommand(optionsCommand(rootCmd))

// 	// BFS applies the flag error function to all subcommands
// 	seenCommands := make(map[*cobra.Command]bool)
// 	var commandStack []*cobra.Command

// 	commandStack = append(commandStack, rootCmd)

// 	for len(commandStack) > 0 {
// 		n := len(commandStack) - 1
// 		curCmd := commandStack[n]
// 		commandStack = commandStack[:n]
// 		seenCommands[curCmd] = true
// 		for _, command := range curCmd.Commands() {
// 			if !seenCommands[command] {
// 				commandStack = append(commandStack, command)
// 			}
// 		}
// 		curCmd.SetFlagErrorFunc(func(_ *cobra.Command, e error) error {
// 			return util.CommandParseError{Err: e}
// 		})
// 	}

// 	return rootCmd
// }

// func hideInheritedFlags(orig *cobra.Command, hidden ...string) {
// 	orig.SetHelpFunc(func(cmd *cobra.Command, args []string) {
// 		for _, hidden := range hidden {
// 			_ = cmd.Flags().MarkHidden(hidden) // nolint: errcheck
// 		}

// 		orig.SetHelpFunc(nil)
// 		orig.HelpFunc()(cmd, args)
// 	})
// }

// func ConfigureLogging(_ *cobra.Command, _ []string) error {
// 	return log.Configure(root.LoggingOptions)
// }

// // seeExperimentalCmd is used for commands that have been around for a release but not graduated from
// // Other alternative
// // for graduatedCmd see https://github.com/istio/istio/pull/26408
// // for softGraduatedCmd see https://github.com/istio/istio/pull/26563
// func seeExperimentalCmd(name string) *cobra.Command {
// 	msg := fmt.Sprintf("(%s is experimental. Use `istioctl experimental %s`)", name, name)
// 	return &cobra.Command{
// 		Use:   name,
// 		Short: msg,
// 		RunE: func(_ *cobra.Command, _ []string) error {
// 			return errors.New(msg)
// 		},
// 	}
// }
