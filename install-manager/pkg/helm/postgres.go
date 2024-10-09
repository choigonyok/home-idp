package helm

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/helm"
)

type Postgres struct {
	Namespace   string `json:"namespace"`
	ReleaseName string `json:"release_name"`
}

func NewPostgresClient(namespace, releaseName string) *Postgres {
	return &Postgres{
		Namespace:   namespace,
		ReleaseName: releaseName,
	}
}

func (c *Postgres) Install(h helm.HelmClient) error {
	fmt.Println("RELEASE: ", c.ReleaseName)
	fmt.Println("NAMESPACE: ", c.Namespace)
	h.Install("bitnami/postgresql:16.0.0", c.Namespace, c.ReleaseName, postgresOverrideValues())
	return nil
}

func (c *Postgres) Uninstall(h helm.HelmClient) error {
	h.Uninstall(c.ReleaseName, c.Namespace)
	return nil
}

// primary:
//   ## @param primary.name Name of the primary database (eg primary, master, leader, ...)
//   ##
//   name: primary
//   ## @param primary.configuration PostgreSQL Primary main configuration to be injected as ConfigMap
//   ## ref: https://www.postgresql.org/docs/current/static/runtime-config.html
//   ##
//   configuration: ""
//   ## @param primary.pgHbaConfiguration PostgreSQL Primary client authentication configuration
//   ## ref: https://www.postgresql.org/docs/current/static/auth-pg-hba-conf.html
//   ## e.g:#
//   ## pgHbaConfiguration: |-
//   ##   local all all trust
//   ##   host all all localhost trust
//   ##   host mydatabase mysuser 192.168.0.0/24 md5
//   ##
//   pgHbaConfiguration: ""
//   ## @param primary.existingConfigmap Name of an existing ConfigMap with PostgreSQL Primary configuration
//   ## NOTE: `primary.configuration` and `primary.pgHbaConfiguration` will be ignored
//   ##
//   existingConfigmap: ""
//   ## @param primary.extendedConfiguration Extended PostgreSQL Primary configuration (appended to main or default configuration)
//   ## ref: https://github.com/bitnami/containers/tree/main/bitnami/postgresql#allow-settings-to-be-loaded-from-files-other-than-the-default-postgresqlconf
//   ##
//   extendedConfiguration: ""
//   ## @param primary.existingExtendedConfigmap Name of an existing ConfigMap with PostgreSQL Primary extended configuration
//   ## NOTE: `primary.extendedConfiguration` will be ignored
//   ##
//   existingExtendedConfigmap: ""
//   ## Initdb configuration
//   ## ref: https://github.com/bitnami/containers/tree/main/bitnami/postgresql#specifying-initdb-arguments
//   ##
//   initdb:
//     ## @param primary.initdb.args PostgreSQL initdb extra arguments
//     ##
//     args: ""
//     ## @param primary.initdb.postgresqlWalDir Specify a custom location for the PostgreSQL transaction log
//     ##
//     postgresqlWalDir: ""
//     ## @param primary.initdb.scripts Dictionary of initdb scripts
//     ## Specify dictionary of scripts to be run at first boot
//     ## e.g:
//     ## scripts:

//     scripts:
// my_init_script.sh: |

//     ## @param primary.initdb.scriptsConfigMap ConfigMap with scripts to be run at first boot
//     ## NOTE: This will override `primary.initdb.scripts`
//     ##
//     scriptsConfigMap: ""
//     ## @param primary.initdb.scriptsSecret Secret with scripts to be run at first boot (in case it contains sensitive information)
//     ## NOTE: This can work along `primary.initdb.scripts` or `primary.initdb.scriptsConfigMap`
//     ##
//     scriptsSecret: ""
//     ## @param primary.initdb.user Specify the PostgreSQL username to execute the initdb scripts
//     ##
//     user: tester
//     ## @param primary.initdb.password Specify the PostgreSQL password to execute the initdb scripts
//     ##
//     password: tester1234

func postgresOverrideValues() map[string]interface{} {
	return map[string]interface{}{
		"auth": map[string]interface{}{
			"enablePostgresUser": false,
			"username":           env.Get("HOME_IDP_STORAGE_USERNAME"),
			"password":           env.Get("HOME_IDP_STORAGE_PASSWORD"),
			"database":           env.Get("HOME_IDP_STORAGE_DB"),
		},
		"primary": map[string]interface{}{
			"initdb": map[string]interface{}{
				"scriptsConfigMap": "home-idp-postgres-initdb",
				"user":             env.Get("HOME_IDP_STORAGE_USERNAME"),
				"password":         env.Get("HOME_IDP_STORAGE_PASSWORD"),
			},
		},
	}
}
