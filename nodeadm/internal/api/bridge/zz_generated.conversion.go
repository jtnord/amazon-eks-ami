//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package bridge

import (
	unsafe "unsafe"

	v1alpha1 "github.com/awslabs/amazon-eks-ami/nodeadm/api/v1alpha1"
	api "github.com/awslabs/amazon-eks-ami/nodeadm/internal/api"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ClusterDetails)(nil), (*api.ClusterDetails)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClusterDetails_To_api_ClusterDetails(a.(*v1alpha1.ClusterDetails), b.(*api.ClusterDetails), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.ClusterDetails)(nil), (*v1alpha1.ClusterDetails)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_ClusterDetails_To_v1alpha1_ClusterDetails(a.(*api.ClusterDetails), b.(*v1alpha1.ClusterDetails), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.KubeletOptions)(nil), (*api.KubeletOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeletOptions_To_api_KubeletOptions(a.(*v1alpha1.KubeletOptions), b.(*api.KubeletOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.KubeletOptions)(nil), (*v1alpha1.KubeletOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_KubeletOptions_To_v1alpha1_KubeletOptions(a.(*api.KubeletOptions), b.(*v1alpha1.KubeletOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.NodeConfig)(nil), (*api.NodeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeConfig_To_api_NodeConfig(a.(*v1alpha1.NodeConfig), b.(*api.NodeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.NodeConfig)(nil), (*v1alpha1.NodeConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_NodeConfig_To_v1alpha1_NodeConfig(a.(*api.NodeConfig), b.(*v1alpha1.NodeConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.NodeConfigList)(nil), (*api.NodeConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeConfigList_To_api_NodeConfigList(a.(*v1alpha1.NodeConfigList), b.(*api.NodeConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.NodeConfigList)(nil), (*v1alpha1.NodeConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_NodeConfigList_To_v1alpha1_NodeConfigList(a.(*api.NodeConfigList), b.(*v1alpha1.NodeConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.NodeConfigSpec)(nil), (*api.NodeConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeConfigSpec_To_api_NodeConfigSpec(a.(*v1alpha1.NodeConfigSpec), b.(*api.NodeConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.NodeConfigSpec)(nil), (*v1alpha1.NodeConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_NodeConfigSpec_To_v1alpha1_NodeConfigSpec(a.(*api.NodeConfigSpec), b.(*v1alpha1.NodeConfigSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ClusterDetails_To_api_ClusterDetails(in *v1alpha1.ClusterDetails, out *api.ClusterDetails, s conversion.Scope) error {
	out.Name = in.Name
	out.APIServerEndpoint = in.APIServerEndpoint
	out.CertificateAuthority = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthority))
	out.CIDR = in.CIDR
	out.EnableOutpost = (*bool)(unsafe.Pointer(in.EnableOutpost))
	out.ID = in.ID
	return nil
}

// Convert_v1alpha1_ClusterDetails_To_api_ClusterDetails is an autogenerated conversion function.
func Convert_v1alpha1_ClusterDetails_To_api_ClusterDetails(in *v1alpha1.ClusterDetails, out *api.ClusterDetails, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterDetails_To_api_ClusterDetails(in, out, s)
}

func autoConvert_api_ClusterDetails_To_v1alpha1_ClusterDetails(in *api.ClusterDetails, out *v1alpha1.ClusterDetails, s conversion.Scope) error {
	out.Name = in.Name
	out.APIServerEndpoint = in.APIServerEndpoint
	out.CertificateAuthority = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthority))
	out.CIDR = in.CIDR
	out.EnableOutpost = (*bool)(unsafe.Pointer(in.EnableOutpost))
	out.ID = in.ID
	return nil
}

// Convert_api_ClusterDetails_To_v1alpha1_ClusterDetails is an autogenerated conversion function.
func Convert_api_ClusterDetails_To_v1alpha1_ClusterDetails(in *api.ClusterDetails, out *v1alpha1.ClusterDetails, s conversion.Scope) error {
	return autoConvert_api_ClusterDetails_To_v1alpha1_ClusterDetails(in, out, s)
}

func autoConvert_v1alpha1_KubeletOptions_To_api_KubeletOptions(in *v1alpha1.KubeletOptions, out *api.KubeletOptions, s conversion.Scope) error {
	out.Config = *(*api.InlineDocument)(unsafe.Pointer(&in.Config))
	out.Flags = *(*[]string)(unsafe.Pointer(&in.Flags))
	return nil
}

// Convert_v1alpha1_KubeletOptions_To_api_KubeletOptions is an autogenerated conversion function.
func Convert_v1alpha1_KubeletOptions_To_api_KubeletOptions(in *v1alpha1.KubeletOptions, out *api.KubeletOptions, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletOptions_To_api_KubeletOptions(in, out, s)
}

func autoConvert_api_KubeletOptions_To_v1alpha1_KubeletOptions(in *api.KubeletOptions, out *v1alpha1.KubeletOptions, s conversion.Scope) error {
	out.Config = *(*map[string]runtime.RawExtension)(unsafe.Pointer(&in.Config))
	out.Flags = *(*[]string)(unsafe.Pointer(&in.Flags))
	return nil
}

// Convert_api_KubeletOptions_To_v1alpha1_KubeletOptions is an autogenerated conversion function.
func Convert_api_KubeletOptions_To_v1alpha1_KubeletOptions(in *api.KubeletOptions, out *v1alpha1.KubeletOptions, s conversion.Scope) error {
	return autoConvert_api_KubeletOptions_To_v1alpha1_KubeletOptions(in, out, s)
}

func autoConvert_v1alpha1_NodeConfig_To_api_NodeConfig(in *v1alpha1.NodeConfig, out *api.NodeConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_NodeConfigSpec_To_api_NodeConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_NodeConfig_To_api_NodeConfig is an autogenerated conversion function.
func Convert_v1alpha1_NodeConfig_To_api_NodeConfig(in *v1alpha1.NodeConfig, out *api.NodeConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeConfig_To_api_NodeConfig(in, out, s)
}

func autoConvert_api_NodeConfig_To_v1alpha1_NodeConfig(in *api.NodeConfig, out *v1alpha1.NodeConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_api_NodeConfigSpec_To_v1alpha1_NodeConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	// INFO: in.Status opted out of conversion generation
	return nil
}

// Convert_api_NodeConfig_To_v1alpha1_NodeConfig is an autogenerated conversion function.
func Convert_api_NodeConfig_To_v1alpha1_NodeConfig(in *api.NodeConfig, out *v1alpha1.NodeConfig, s conversion.Scope) error {
	return autoConvert_api_NodeConfig_To_v1alpha1_NodeConfig(in, out, s)
}

func autoConvert_v1alpha1_NodeConfigList_To_api_NodeConfigList(in *v1alpha1.NodeConfigList, out *api.NodeConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.NodeConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_NodeConfig_To_api_NodeConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_NodeConfigList_To_api_NodeConfigList is an autogenerated conversion function.
func Convert_v1alpha1_NodeConfigList_To_api_NodeConfigList(in *v1alpha1.NodeConfigList, out *api.NodeConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeConfigList_To_api_NodeConfigList(in, out, s)
}

func autoConvert_api_NodeConfigList_To_v1alpha1_NodeConfigList(in *api.NodeConfigList, out *v1alpha1.NodeConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha1.NodeConfig, len(*in))
		for i := range *in {
			if err := Convert_api_NodeConfig_To_v1alpha1_NodeConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_api_NodeConfigList_To_v1alpha1_NodeConfigList is an autogenerated conversion function.
func Convert_api_NodeConfigList_To_v1alpha1_NodeConfigList(in *api.NodeConfigList, out *v1alpha1.NodeConfigList, s conversion.Scope) error {
	return autoConvert_api_NodeConfigList_To_v1alpha1_NodeConfigList(in, out, s)
}

func autoConvert_v1alpha1_NodeConfigSpec_To_api_NodeConfigSpec(in *v1alpha1.NodeConfigSpec, out *api.NodeConfigSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_ClusterDetails_To_api_ClusterDetails(&in.Cluster, &out.Cluster, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeletOptions_To_api_KubeletOptions(&in.Kubelet, &out.Kubelet, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_NodeConfigSpec_To_api_NodeConfigSpec is an autogenerated conversion function.
func Convert_v1alpha1_NodeConfigSpec_To_api_NodeConfigSpec(in *v1alpha1.NodeConfigSpec, out *api.NodeConfigSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeConfigSpec_To_api_NodeConfigSpec(in, out, s)
}

func autoConvert_api_NodeConfigSpec_To_v1alpha1_NodeConfigSpec(in *api.NodeConfigSpec, out *v1alpha1.NodeConfigSpec, s conversion.Scope) error {
	if err := Convert_api_ClusterDetails_To_v1alpha1_ClusterDetails(&in.Cluster, &out.Cluster, s); err != nil {
		return err
	}
	if err := Convert_api_KubeletOptions_To_v1alpha1_KubeletOptions(&in.Kubelet, &out.Kubelet, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_NodeConfigSpec_To_v1alpha1_NodeConfigSpec is an autogenerated conversion function.
func Convert_api_NodeConfigSpec_To_v1alpha1_NodeConfigSpec(in *api.NodeConfigSpec, out *v1alpha1.NodeConfigSpec, s conversion.Scope) error {
	return autoConvert_api_NodeConfigSpec_To_v1alpha1_NodeConfigSpec(in, out, s)
}
