package v1beta1

import (
	"github.com/rancher/opni-monitoring/pkg/config/meta"
)

type GatewayConfig struct {
	meta.TypeMeta `json:",inline"`

	Spec GatewayConfigSpec `json:"spec,omitempty"`
}

type GatewayConfigSpec struct {
	ListenAddress  string         `json:"listenAddress,omitempty"`
	Management     ManagementSpec `json:"management,omitempty"`
	EnableMonitor  bool           `json:"enableMonitor,omitempty"`
	TrustedProxies []string       `json:"trustedProxies,omitempty"`
	Cortex         CortexSpec     `json:"cortex,omitempty"`
	AuthProvider   string         `json:"authProvider,omitempty"`
	Storage        StorageSpec    `json:"storage,omitempty"`
	Certs          CertsSpec      `json:"certs,omitempty"`
	Plugins        PluginsSpec    `json:"plugins,omitempty"`
}

type ManagementSpec struct {
	GRPCListenAddress string `json:"grpcListenAddress,omitempty"`
	HTTPListenAddress string `json:"httpListenAddress,omitempty"`
	WebListenAddress  string `json:"webListenAddress,omitempty"`
}

type CortexSpec struct {
	Distributor   DistributorSpec   `json:"distributor,omitempty"`
	Ingester      IngesterSpec      `json:"ingester,omitempty"`
	Alertmanager  AlertmanagerSpec  `json:"alertmanager,omitempty"`
	Ruler         RulerSpec         `json:"ruler,omitempty"`
	QueryFrontend QueryFrontendSpec `json:"queryFrontend,omitempty"`
	Certs         MTLSSpec          `json:"certs,omitempty"`
}

type DistributorSpec struct {
	Address string `json:"address,omitempty"`
}

type IngesterSpec struct {
	Address string `json:"address,omitempty"`
}

type AlertmanagerSpec struct {
	Address string `json:"address,omitempty"`
}

type RulerSpec struct {
	Address string `json:"address,omitempty"`
}

type QueryFrontendSpec struct {
	Address string `json:"address,omitempty"`
}

type MTLSSpec struct {
	ServerCA   string `json:"serverCA,omitempty"`
	ClientCA   string `json:"clientCA,omitempty"`
	ClientCert string `json:"clientCert,omitempty"`
	ClientKey  string `json:"clientKey,omitempty"`
}

type CertsSpec struct {
	// Path to a PEM encoded CA certificate file. Mutually exclusive with CACertData
	CACert *string `json:"caCert,omitempty"`
	// String containing PEM encoded CA certificate data. Mutually exclusive with CACert
	CACertData *string `json:"caCertData,omitempty"`
	// Path to a PEM encoded server certificate file. Mutually exclusive with ServingCertData
	ServingCert *string `json:"servingCert,omitempty"`
	// String containing PEM encoded server certificate data. Mutually exclusive with ServingCert
	ServingCertData *string `json:"servingCertData,omitempty"`
	// Path to a PEM encoded server key file. Mutually exclusive with ServingKeyData
	ServingKey *string `json:"servingKey,omitempty"`
	// String containing PEM encoded server key data. Mutually exclusive with ServingKey
	ServingKeyData *string `json:"servingKeyData,omitempty"`
}

type PluginsSpec struct {
	// Directories to look for plugins in
	Dirs []string `json:"dirs,omitempty"`
}

func (s *GatewayConfigSpec) SetDefaults() {
	if s == nil {
		return
	}
	if s.ListenAddress == "" {
		s.ListenAddress = ":8080"
	}
	if s.Cortex.Distributor.Address == "" {
		s.Cortex.Distributor.Address = "cortex-distributor:8080"
	}
	if s.Cortex.Ingester.Address == "" {
		s.Cortex.Ingester.Address = "cortex-ingester:8080"
	}
	if s.Cortex.Alertmanager.Address == "" {
		s.Cortex.Alertmanager.Address = "cortex-alertmanager:8080"
	}
	if s.Cortex.Ruler.Address == "" {
		s.Cortex.Ruler.Address = "cortex-ruler:8080"
	}
	if s.Cortex.QueryFrontend.Address == "" {
		s.Cortex.QueryFrontend.Address = "cortex-query-frontend:8080"
	}
}

type StorageType string

const (
	StorageTypeEtcd   StorageType = "etcd"
	StorageTypeSecret StorageType = "secret"
)

type LoggingStorageType string

const (
	StorageTypeKubernetes LoggingStorageType = "kubernetes"
)

type StorageSpec struct {
	Type        StorageType             `json:"type,omitempty"`
	LoggingType LoggingStorageType      `json:"loggingType,omitempty"`
	Etcd        *EtcdStorageSpec        `json:"etcd,omitempty"`
	Kubernetes  *KuberneteesStorageSpec `json:"kubernetes,omitempty"`
}

type EtcdStorageSpec struct {
	Endpoints []string  `json:"endpoints,omitempty"`
	Certs     *MTLSSpec `json:"certs,omitempty"`
}

type KuberneteesStorageSpec struct {
	SystemNamespace string `json:"systemNamespace,omitempty"`
}
