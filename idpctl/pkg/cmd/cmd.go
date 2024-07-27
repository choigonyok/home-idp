package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/choigonyok/home-idp/pkg/file"
	"github.com/choigonyok/home-idp/pkg/git"
	"github.com/choigonyok/home-idp/pkg/kube"
	"github.com/choigonyok/home-idp/pkg/object"
	ptr "github.com/choigonyok/home-idp/pkg/pointer"
	"github.com/choigonyok/home-idp/pkg/secret"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/yaml"
)

const (
	// Location to read istioctl defaults from
	defaultIstioctlConfig = "$HOME/.idpctl/config.yaml"
)

const (
	FlagCharts = "charts"
)

func NewRootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:               "idpctl",
		Short:             "Home-idp Command Line Interface",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `Command line interface to deploy home-idp application in kubernetes cluster`,
	}
	c.SetArgs(os.Args[1:])
	addRootFlags(c)
	addSubCmds(c)
	return c
}

func addRootFlags(c *cobra.Command) {
	c.PersistentFlags().StringP("namespace", "n", "", "Kubernetes namespace")
}

func addSubCmds(c *cobra.Command) {
	c.AddCommand(getStatusCmd())

	c.AddCommand(getInstallCmd())
}

func getInstallCmd() *cobra.Command {
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Install home-idp app",
		Long:  "Install home-idp application in kubernetes cluster with CLI",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("INSTALL COMMAND EXECUTED")
			key := util.Hash("andromeda0085")
			fmt.Println(util.EqualWithKey("andromeda0085", key))
			return nil
		},
	}

	installCmd.AddCommand(getInstallSecretManagerCmd())

	return installCmd
}

func getInstallSecretManagerCmd() *cobra.Command {
	installCmd := &cobra.Command{
		Use:   "secret-manager",
		Short: "Install home-idp secret manager",
		Long:  "Install home-idp secret manager in kubernetes cluster with CLI",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("SECRET-MANAGER COMMAND EXECUTED")

			plaintext := "andromeda0085*"

			// 암호화
			encrypted, key, _ := util.Encrypt(plaintext)
			fmt.Println("EN:", encrypted)
			kk, _ := base64.RawStdEncoding.DecodeString(key)
			// 복호화
			decrypted, _ := util.Decrypt("andromeda0085", kk)
			fmt.Println("DE:", decrypted)

			return nil
		},
	}
	return installCmd
}

type statusFlags struct {
	config *string
}

func getStatusCmd() *cobra.Command {
	f := &statusFlags{
		config: ptr.Of[string](""),
	}

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "show status home-idp app",
		Long:  "Show status of home-idp application deployed in kubernetes cluster with CLI",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			restConfig, _ := kube.GetKubeConfig()
			client, _ := kube.NewClient(restConfig)
			dc, _ := kube.GetDynamicClient(restConfig)
			serviceList, _ := kube.ListServices("log-system", client)
			for _, podName := range serviceList.Items {
				fmt.Println(podName.Name)
			}
			fmt.Println()
			fmt.Println()
			fmt.Println()

			configFilePath, nil := cmd.PersistentFlags().GetString("config")
			fileContent, nil := file.ReadFile(configFilePath)
			// fileContent, nil := file.ReadYamlFile(*f.config, os.Stdin)
			// if yamlSyntaxErr := file.ValidateYamlFileFormat(fileContent); yamlSyntaxErr != nil {
			// 	return yamlSyntaxErr
			// }

			gvk, obj := object.ParseObjectsFromManifest(string(fileContent))

			mapIOP := make(map[string]any)
			yaml.Unmarshal([]byte(fileContent), &mapIOP)
			// yaml.NewYAMLToJSONDecoder(bytes.NewReader([]byte(mapIOP)))

			kube.ApplyManifest("pods", "default", dc, obj, gvk)

			fmt.Println(mapIOP)

			gc := git.NewClient("choigonyok", "")
			// git.ValidateClient(c)
			// fmt.Println(git.CreateRepository("test3", true, gc))
			git.ConnectRepository(gc, "test2")
			// fmt.Println(git.CreateWebhook("https://argocd.slexn.com/api/webhook", gc))

			// fmt.Println(git.DeleteRepository("test1", gc))
			// fmt.Println(git.DeleteRepository("test2", gc))
			// fmt.Println(git.DeleteRepository("test3", gc))
			// fmt.Println(git.DeleteRepository("test4", gc))
			// fmt.Println(git.DeleteRepository("test5", gc))
			// fmt.Println(git.DeleteRepository("test6", gc))
			// fmt.Println(git.DeleteRepository("test7", gc))
			rc := secret.NewClient()
			rc.LogInWithRoot()

			rc.AddPolicy("github", "/github/*", []string{"create", "read", "update", "delete", "list"})
			secretID, err := rc.ApplyAppRoleWithPolicies("github", []string{"default", "github"})
			fmt.Println("ERROR1", err)
			err = rc.LogInWithAppRole("github", secretID)
			fmt.Println("ERROR2", err)
			err = rc.PutSecretV2("github", "github", "test/config", map[string]any{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			})
			fmt.Println("ERROR3", err)
			fmt.Println("TEST3")
			return nil
		},
	}

	statusCmd.PersistentFlags().StringVarP(f.config, "config", "f", "", "Configuration File")

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
