// +build !ignore_autogenerated

//
// Copyright 2022 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	metav1 "github.com/ibm/ibm-cert-manager-operator/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolver) DeepCopyInto(out *ACMEChallengeSolver) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(CertificateDNSNameSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTP01 != nil {
		in, out := &in.HTTP01, &out.HTTP01
		*out = new(ACMEChallengeSolverHTTP01)
		(*in).DeepCopyInto(*out)
	}
	if in.DNS01 != nil {
		in, out := &in.DNS01, &out.DNS01
		*out = new(ACMEChallengeSolverDNS01)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolver.
func (in *ACMEChallengeSolver) DeepCopy() *ACMEChallengeSolver {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverDNS01) DeepCopyInto(out *ACMEChallengeSolverDNS01) {
	*out = *in
	if in.Akamai != nil {
		in, out := &in.Akamai, &out.Akamai
		*out = new(ACMEIssuerDNS01ProviderAkamai)
		**out = **in
	}
	if in.CloudDNS != nil {
		in, out := &in.CloudDNS, &out.CloudDNS
		*out = new(ACMEIssuerDNS01ProviderCloudDNS)
		(*in).DeepCopyInto(*out)
	}
	if in.Cloudflare != nil {
		in, out := &in.Cloudflare, &out.Cloudflare
		*out = new(ACMEIssuerDNS01ProviderCloudflare)
		(*in).DeepCopyInto(*out)
	}
	if in.Route53 != nil {
		in, out := &in.Route53, &out.Route53
		*out = new(ACMEIssuerDNS01ProviderRoute53)
		**out = **in
	}
	if in.AzureDNS != nil {
		in, out := &in.AzureDNS, &out.AzureDNS
		*out = new(ACMEIssuerDNS01ProviderAzureDNS)
		(*in).DeepCopyInto(*out)
	}
	if in.DigitalOcean != nil {
		in, out := &in.DigitalOcean, &out.DigitalOcean
		*out = new(ACMEIssuerDNS01ProviderDigitalOcean)
		**out = **in
	}
	if in.AcmeDNS != nil {
		in, out := &in.AcmeDNS, &out.AcmeDNS
		*out = new(ACMEIssuerDNS01ProviderAcmeDNS)
		**out = **in
	}
	if in.RFC2136 != nil {
		in, out := &in.RFC2136, &out.RFC2136
		*out = new(ACMEIssuerDNS01ProviderRFC2136)
		**out = **in
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(ACMEIssuerDNS01ProviderWebhook)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverDNS01.
func (in *ACMEChallengeSolverDNS01) DeepCopy() *ACMEChallengeSolverDNS01 {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverDNS01)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverHTTP01) DeepCopyInto(out *ACMEChallengeSolverHTTP01) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(ACMEChallengeSolverHTTP01Ingress)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverHTTP01.
func (in *ACMEChallengeSolverHTTP01) DeepCopy() *ACMEChallengeSolverHTTP01 {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverHTTP01)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverHTTP01Ingress) DeepCopyInto(out *ACMEChallengeSolverHTTP01Ingress) {
	*out = *in
	if in.Class != nil {
		in, out := &in.Class, &out.Class
		*out = new(string)
		**out = **in
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(ACMEChallengeSolverHTTP01IngressPodTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.IngressTemplate != nil {
		in, out := &in.IngressTemplate, &out.IngressTemplate
		*out = new(ACMEChallengeSolverHTTP01IngressTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverHTTP01Ingress.
func (in *ACMEChallengeSolverHTTP01Ingress) DeepCopy() *ACMEChallengeSolverHTTP01Ingress {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverHTTP01Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverHTTP01IngressObjectMeta) DeepCopyInto(out *ACMEChallengeSolverHTTP01IngressObjectMeta) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverHTTP01IngressObjectMeta.
func (in *ACMEChallengeSolverHTTP01IngressObjectMeta) DeepCopy() *ACMEChallengeSolverHTTP01IngressObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverHTTP01IngressObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverHTTP01IngressPodObjectMeta) DeepCopyInto(out *ACMEChallengeSolverHTTP01IngressPodObjectMeta) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverHTTP01IngressPodObjectMeta.
func (in *ACMEChallengeSolverHTTP01IngressPodObjectMeta) DeepCopy() *ACMEChallengeSolverHTTP01IngressPodObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverHTTP01IngressPodObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverHTTP01IngressPodSpec) DeepCopyInto(out *ACMEChallengeSolverHTTP01IngressPodSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverHTTP01IngressPodSpec.
func (in *ACMEChallengeSolverHTTP01IngressPodSpec) DeepCopy() *ACMEChallengeSolverHTTP01IngressPodSpec {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverHTTP01IngressPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverHTTP01IngressPodTemplate) DeepCopyInto(out *ACMEChallengeSolverHTTP01IngressPodTemplate) {
	*out = *in
	in.ACMEChallengeSolverHTTP01IngressPodObjectMeta.DeepCopyInto(&out.ACMEChallengeSolverHTTP01IngressPodObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverHTTP01IngressPodTemplate.
func (in *ACMEChallengeSolverHTTP01IngressPodTemplate) DeepCopy() *ACMEChallengeSolverHTTP01IngressPodTemplate {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverHTTP01IngressPodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEChallengeSolverHTTP01IngressTemplate) DeepCopyInto(out *ACMEChallengeSolverHTTP01IngressTemplate) {
	*out = *in
	in.ACMEChallengeSolverHTTP01IngressObjectMeta.DeepCopyInto(&out.ACMEChallengeSolverHTTP01IngressObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEChallengeSolverHTTP01IngressTemplate.
func (in *ACMEChallengeSolverHTTP01IngressTemplate) DeepCopy() *ACMEChallengeSolverHTTP01IngressTemplate {
	if in == nil {
		return nil
	}
	out := new(ACMEChallengeSolverHTTP01IngressTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEExternalAccountBinding) DeepCopyInto(out *ACMEExternalAccountBinding) {
	*out = *in
	out.Key = in.Key
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEExternalAccountBinding.
func (in *ACMEExternalAccountBinding) DeepCopy() *ACMEExternalAccountBinding {
	if in == nil {
		return nil
	}
	out := new(ACMEExternalAccountBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuer) DeepCopyInto(out *ACMEIssuer) {
	*out = *in
	if in.ExternalAccountBinding != nil {
		in, out := &in.ExternalAccountBinding, &out.ExternalAccountBinding
		*out = new(ACMEExternalAccountBinding)
		**out = **in
	}
	out.PrivateKey = in.PrivateKey
	if in.Solvers != nil {
		in, out := &in.Solvers, &out.Solvers
		*out = make([]ACMEChallengeSolver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuer.
func (in *ACMEIssuer) DeepCopy() *ACMEIssuer {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderAcmeDNS) DeepCopyInto(out *ACMEIssuerDNS01ProviderAcmeDNS) {
	*out = *in
	out.AccountSecret = in.AccountSecret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderAcmeDNS.
func (in *ACMEIssuerDNS01ProviderAcmeDNS) DeepCopy() *ACMEIssuerDNS01ProviderAcmeDNS {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderAcmeDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderAkamai) DeepCopyInto(out *ACMEIssuerDNS01ProviderAkamai) {
	*out = *in
	out.ClientToken = in.ClientToken
	out.ClientSecret = in.ClientSecret
	out.AccessToken = in.AccessToken
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderAkamai.
func (in *ACMEIssuerDNS01ProviderAkamai) DeepCopy() *ACMEIssuerDNS01ProviderAkamai {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderAkamai)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderAzureDNS) DeepCopyInto(out *ACMEIssuerDNS01ProviderAzureDNS) {
	*out = *in
	if in.ClientSecret != nil {
		in, out := &in.ClientSecret, &out.ClientSecret
		*out = new(metav1.SecretKeySelector)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderAzureDNS.
func (in *ACMEIssuerDNS01ProviderAzureDNS) DeepCopy() *ACMEIssuerDNS01ProviderAzureDNS {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderAzureDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderCloudDNS) DeepCopyInto(out *ACMEIssuerDNS01ProviderCloudDNS) {
	*out = *in
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(metav1.SecretKeySelector)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderCloudDNS.
func (in *ACMEIssuerDNS01ProviderCloudDNS) DeepCopy() *ACMEIssuerDNS01ProviderCloudDNS {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderCloudDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderCloudflare) DeepCopyInto(out *ACMEIssuerDNS01ProviderCloudflare) {
	*out = *in
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(metav1.SecretKeySelector)
		**out = **in
	}
	if in.APIToken != nil {
		in, out := &in.APIToken, &out.APIToken
		*out = new(metav1.SecretKeySelector)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderCloudflare.
func (in *ACMEIssuerDNS01ProviderCloudflare) DeepCopy() *ACMEIssuerDNS01ProviderCloudflare {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderCloudflare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderDigitalOcean) DeepCopyInto(out *ACMEIssuerDNS01ProviderDigitalOcean) {
	*out = *in
	out.Token = in.Token
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderDigitalOcean.
func (in *ACMEIssuerDNS01ProviderDigitalOcean) DeepCopy() *ACMEIssuerDNS01ProviderDigitalOcean {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderDigitalOcean)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderRFC2136) DeepCopyInto(out *ACMEIssuerDNS01ProviderRFC2136) {
	*out = *in
	out.TSIGSecret = in.TSIGSecret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderRFC2136.
func (in *ACMEIssuerDNS01ProviderRFC2136) DeepCopy() *ACMEIssuerDNS01ProviderRFC2136 {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderRFC2136)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderRoute53) DeepCopyInto(out *ACMEIssuerDNS01ProviderRoute53) {
	*out = *in
	out.SecretAccessKey = in.SecretAccessKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderRoute53.
func (in *ACMEIssuerDNS01ProviderRoute53) DeepCopy() *ACMEIssuerDNS01ProviderRoute53 {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderRoute53)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderWebhook) DeepCopyInto(out *ACMEIssuerDNS01ProviderWebhook) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderWebhook.
func (in *ACMEIssuerDNS01ProviderWebhook) DeepCopy() *ACMEIssuerDNS01ProviderWebhook {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerStatus) DeepCopyInto(out *ACMEIssuerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerStatus.
func (in *ACMEIssuerStatus) DeepCopy() *ACMEIssuerStatus {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateDNSNameSelector) DeepCopyInto(out *CertificateDNSNameSelector) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNSZones != nil {
		in, out := &in.DNSZones, &out.DNSZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateDNSNameSelector.
func (in *CertificateDNSNameSelector) DeepCopy() *CertificateDNSNameSelector {
	if in == nil {
		return nil
	}
	out := new(CertificateDNSNameSelector)
	in.DeepCopyInto(out)
	return out
}
