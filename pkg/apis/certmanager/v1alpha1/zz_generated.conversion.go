// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	certmanager "github.com/jetstack/cert-manager/pkg/apis/certmanager"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ACME_To_certmanager_ACME,
		Convert_certmanager_ACME_To_v1alpha1_ACME,
		Convert_v1alpha1_ACMEDNSConfig_To_certmanager_ACMEDNSConfig,
		Convert_certmanager_ACMEDNSConfig_To_v1alpha1_ACMEDNSConfig,
		Convert_v1alpha1_ACMEDNSConfigCloudDNS_To_certmanager_ACMEDNSConfigCloudDNS,
		Convert_certmanager_ACMEDNSConfigCloudDNS_To_v1alpha1_ACMEDNSConfigCloudDNS,
		Convert_v1alpha1_Certificate_To_certmanager_Certificate,
		Convert_certmanager_Certificate_To_v1alpha1_Certificate,
		Convert_v1alpha1_CertificateList_To_certmanager_CertificateList,
		Convert_certmanager_CertificateList_To_v1alpha1_CertificateList,
		Convert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec,
		Convert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec,
		Convert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus,
		Convert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus,
		Convert_v1alpha1_ProviderConfig_To_certmanager_ProviderConfig,
		Convert_certmanager_ProviderConfig_To_v1alpha1_ProviderConfig,
	)
}

func autoConvert_v1alpha1_ACME_To_certmanager_ACME(in *ACME, out *certmanager.ACME, s conversion.Scope) error {
	out.Challenge = certmanager.ACMEChallengeType(in.Challenge)
	out.URL = in.URL
	out.Email = in.Email
	out.DNS = (*certmanager.ACMEDNSConfig)(unsafe.Pointer(in.DNS))
	return nil
}

func Convert_v1alpha1_ACME_To_certmanager_ACME(in *ACME, out *certmanager.ACME, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACME_To_certmanager_ACME(in, out, s)
}

func autoConvert_certmanager_ACME_To_v1alpha1_ACME(in *certmanager.ACME, out *ACME, s conversion.Scope) error {
	out.Challenge = ACMEChallengeType(in.Challenge)
	out.URL = in.URL
	out.Email = in.Email
	out.DNS = (*ACMEDNSConfig)(unsafe.Pointer(in.DNS))
	return nil
}

func Convert_certmanager_ACME_To_v1alpha1_ACME(in *certmanager.ACME, out *ACME, s conversion.Scope) error {
	return autoConvert_certmanager_ACME_To_v1alpha1_ACME(in, out, s)
}

func autoConvert_v1alpha1_ACMEDNSConfig_To_certmanager_ACMEDNSConfig(in *ACMEDNSConfig, out *certmanager.ACMEDNSConfig, s conversion.Scope) error {
	out.CloudDNS = (*certmanager.ACMEDNSConfigCloudDNS)(unsafe.Pointer(in.CloudDNS))
	return nil
}

func Convert_v1alpha1_ACMEDNSConfig_To_certmanager_ACMEDNSConfig(in *ACMEDNSConfig, out *certmanager.ACMEDNSConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEDNSConfig_To_certmanager_ACMEDNSConfig(in, out, s)
}

func autoConvert_certmanager_ACMEDNSConfig_To_v1alpha1_ACMEDNSConfig(in *certmanager.ACMEDNSConfig, out *ACMEDNSConfig, s conversion.Scope) error {
	out.CloudDNS = (*ACMEDNSConfigCloudDNS)(unsafe.Pointer(in.CloudDNS))
	return nil
}

func Convert_certmanager_ACMEDNSConfig_To_v1alpha1_ACMEDNSConfig(in *certmanager.ACMEDNSConfig, out *ACMEDNSConfig, s conversion.Scope) error {
	return autoConvert_certmanager_ACMEDNSConfig_To_v1alpha1_ACMEDNSConfig(in, out, s)
}

func autoConvert_v1alpha1_ACMEDNSConfigCloudDNS_To_certmanager_ACMEDNSConfigCloudDNS(in *ACMEDNSConfigCloudDNS, out *certmanager.ACMEDNSConfigCloudDNS, s conversion.Scope) error {
	return nil
}

func Convert_v1alpha1_ACMEDNSConfigCloudDNS_To_certmanager_ACMEDNSConfigCloudDNS(in *ACMEDNSConfigCloudDNS, out *certmanager.ACMEDNSConfigCloudDNS, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACMEDNSConfigCloudDNS_To_certmanager_ACMEDNSConfigCloudDNS(in, out, s)
}

func autoConvert_certmanager_ACMEDNSConfigCloudDNS_To_v1alpha1_ACMEDNSConfigCloudDNS(in *certmanager.ACMEDNSConfigCloudDNS, out *ACMEDNSConfigCloudDNS, s conversion.Scope) error {
	return nil
}

func Convert_certmanager_ACMEDNSConfigCloudDNS_To_v1alpha1_ACMEDNSConfigCloudDNS(in *certmanager.ACMEDNSConfigCloudDNS, out *ACMEDNSConfigCloudDNS, s conversion.Scope) error {
	return autoConvert_certmanager_ACMEDNSConfigCloudDNS_To_v1alpha1_ACMEDNSConfigCloudDNS(in, out, s)
}

func autoConvert_v1alpha1_Certificate_To_certmanager_Certificate(in *Certificate, out *certmanager.Certificate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_Certificate_To_certmanager_Certificate(in *Certificate, out *certmanager.Certificate, s conversion.Scope) error {
	return autoConvert_v1alpha1_Certificate_To_certmanager_Certificate(in, out, s)
}

func autoConvert_certmanager_Certificate_To_v1alpha1_Certificate(in *certmanager.Certificate, out *Certificate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_certmanager_Certificate_To_v1alpha1_Certificate(in *certmanager.Certificate, out *Certificate, s conversion.Scope) error {
	return autoConvert_certmanager_Certificate_To_v1alpha1_Certificate(in, out, s)
}

func autoConvert_v1alpha1_CertificateList_To_certmanager_CertificateList(in *CertificateList, out *certmanager.CertificateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]certmanager.Certificate)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_CertificateList_To_certmanager_CertificateList(in *CertificateList, out *certmanager.CertificateList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateList_To_certmanager_CertificateList(in, out, s)
}

func autoConvert_certmanager_CertificateList_To_v1alpha1_CertificateList(in *certmanager.CertificateList, out *CertificateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]Certificate, 0)
	} else {
		out.Items = *(*[]Certificate)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_certmanager_CertificateList_To_v1alpha1_CertificateList(in *certmanager.CertificateList, out *CertificateList, s conversion.Scope) error {
	return autoConvert_certmanager_CertificateList_To_v1alpha1_CertificateList(in, out, s)
}

func autoConvert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(in *CertificateSpec, out *certmanager.CertificateSpec, s conversion.Scope) error {
	out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	if err := Convert_v1alpha1_ProviderConfig_To_certmanager_ProviderConfig(&in.ProviderConfig, &out.ProviderConfig, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(in *CertificateSpec, out *certmanager.CertificateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateSpec_To_certmanager_CertificateSpec(in, out, s)
}

func autoConvert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(in *certmanager.CertificateSpec, out *CertificateSpec, s conversion.Scope) error {
	if in.Domains == nil {
		out.Domains = make([]string, 0)
	} else {
		out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	}
	if err := Convert_certmanager_ProviderConfig_To_v1alpha1_ProviderConfig(&in.ProviderConfig, &out.ProviderConfig, s); err != nil {
		return err
	}
	return nil
}

func Convert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(in *certmanager.CertificateSpec, out *CertificateSpec, s conversion.Scope) error {
	return autoConvert_certmanager_CertificateSpec_To_v1alpha1_CertificateSpec(in, out, s)
}

func autoConvert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(in *CertificateStatus, out *certmanager.CertificateStatus, s conversion.Scope) error {
	return nil
}

func Convert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(in *CertificateStatus, out *certmanager.CertificateStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertificateStatus_To_certmanager_CertificateStatus(in, out, s)
}

func autoConvert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(in *certmanager.CertificateStatus, out *CertificateStatus, s conversion.Scope) error {
	return nil
}

func Convert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(in *certmanager.CertificateStatus, out *CertificateStatus, s conversion.Scope) error {
	return autoConvert_certmanager_CertificateStatus_To_v1alpha1_CertificateStatus(in, out, s)
}

func autoConvert_v1alpha1_ProviderConfig_To_certmanager_ProviderConfig(in *ProviderConfig, out *certmanager.ProviderConfig, s conversion.Scope) error {
	out.ACME = (*certmanager.ACME)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_v1alpha1_ProviderConfig_To_certmanager_ProviderConfig(in *ProviderConfig, out *certmanager.ProviderConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderConfig_To_certmanager_ProviderConfig(in, out, s)
}

func autoConvert_certmanager_ProviderConfig_To_v1alpha1_ProviderConfig(in *certmanager.ProviderConfig, out *ProviderConfig, s conversion.Scope) error {
	out.ACME = (*ACME)(unsafe.Pointer(in.ACME))
	return nil
}

func Convert_certmanager_ProviderConfig_To_v1alpha1_ProviderConfig(in *certmanager.ProviderConfig, out *ProviderConfig, s conversion.Scope) error {
	return autoConvert_certmanager_ProviderConfig_To_v1alpha1_ProviderConfig(in, out, s)
}
