package api

type ClusterParams struct {
	Namespace                           string                 `json:"namespace"`
	ExternalAPIAddress                  string                 `json:"externalAPIAddress"`
	ExternalAPIPort                     uint                   `json:"externalAPIPort"`
	ExternalAPIIPAddress                string                 `json:"externalAPIIPAddress"`
	ExternalOpenVPNAddress              string                 `json:"externalVPNAddress"`
	ExternalOpenVPNPort                 uint                   `json:"externalVPNPort"`
	ExternalOAuthAddress                string                 `json:"externalOAuthAddress"`
	ExternalOauthPort                   uint                   `json:"externalOauthPort"`
	IdentityProviders                   string                 `json:"identityProviders"`
	ServiceCIDR                         string                 `json:"serviceCIDR"`
	NamedCerts                          []NamedCert            `json:"namedCerts,omitempty"`
	PodCIDR                             string                 `json:"podCIDR"`
	ReleaseImage                        string                 `json:"releaseImage"`
	IngressSubdomain                    string                 `json:"ingressSubdomain"`
	OpenShiftAPIClusterIP               string                 `json:"openshiftAPIClusterIP"`
	ImageRegistryHTTPSecret             string                 `json:"imageRegistryHTTPSecret"`
	RouterNodePortHTTP                  string                 `json:"routerNodePortHTTP"`
	RouterNodePortHTTPS                 string                 `json:"routerNodePortHTTPS"`
	BaseDomain                          string                 `json:"baseDomain"`
	NetworkType                         string                 `json:"networkType"`
	Replicas                            string                 `json:"replicas"`
	EtcdClientName                      string                 `json:"etcdClientName"`
	OriginReleasePrefix                 string                 `json:"originReleasePrefix"`
	OpenshiftAPIServerCABundle          string                 `json:"openshiftAPIServerCABundle"`
	CloudProvider                       string                 `json:"cloudProvider"`
	CVOSetupImage                       string                 `json:"cvoSetupImage"`
	InternalAPIPort                     uint                   `json:"internalAPIPort"`
	RouterServiceType                   string                 `json:"routerServiceType"`
	KubeAPIServerResources              []ResourceRequirements `json:"kubeAPIServerResources"`
	OpenshiftControllerManagerResources []ResourceRequirements `json:"openshiftControllerManagerResources"`
	ClusterVersionOperatorResources     []ResourceRequirements `json:"clusterVersionOperatorResources"`
	KubeControllerManagerResources      []ResourceRequirements `json:"kubeControllerManagerResources"`
	OpenshiftAPIServerResources         []ResourceRequirements `json:"openshiftAPIServerResources"`
	KubeSchedulerResources              []ResourceRequirements `json:"kubeSchedulerResources"`
	ControlPlaneOperatorResources       []ResourceRequirements `json:"controlPlaneOperatorResources"`
	OAuthServerResources                []ResourceRequirements `json:"oAuthServerResources"`
	ClusterPolicyControllerResources    []ResourceRequirements `json:"clusterPolicyControllerResources"`
	AutoApproverResources               []ResourceRequirements `json:"autoApproverResources"`
	OpenVPNClientResources              []ResourceRequirements `json:"openVPNClientResources"`
	OpenVPNServerResources              []ResourceRequirements `json:"openVPNServerResources"`
	APIServerAuditEnabled               bool                   `json:"apiServerAuditEnabled"`
	RestartDate                         string                 `json:"restartDate"`
	ControlPlaneOperatorImage           string                 `json:"controlPlaneOperatorImage"`
	ControlPlaneOperatorControllers     []string               `json:"controlPlaneOperatorControllers"`
	ExtraFeatureGates                   []string               `json:"extraFeatureGates"`
	ControlPlaneOperatorSecurity        string                 `json:"controlPlaneOperatorSecurity"`
	ApiserverLivenessPath               string                 `json:"apiserverLivenessPath"`
	DefaultFeatureGates                 []string
	PlatformType                        string `json:"platformType"`
}

type NamedCert struct {
	NamedCertPrefix string `json:"namedCertPrefix"`
	NamedCertDomain string `json:"namedCertDomain"`
}

type ResourceRequirements struct {
	ResourceLimit   []ResourceLimit   `json:"resourceLimit"`
	ResourceRequest []ResourceRequest `json:"resourceRequest"`
}

type ResourceLimit struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

type ResourceRequest struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
