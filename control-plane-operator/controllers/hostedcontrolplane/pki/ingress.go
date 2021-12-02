package pki

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"github.com/openshift/hypershift/support/config"
)

func ReconcileIngressCert(secret, ca *corev1.Secret, ownerRef config.OwnerRef, externalOAuthAddress, ingressSubdomain string) error {
	var ingressNumericIPs, ingressHostNames []string
	if isNumericIP(externalOAuthAddress) {
		ingressNumericIPs = append(ingressNumericIPs, externalOAuthAddress)
	} else {
		ingressHostNames = append(ingressHostNames, externalOAuthAddress)
	}
	ingressHostNames = append(ingressHostNames, fmt.Sprintf("*.%s", ingressSubdomain))
	return reconcileSignedCertWithAddresses(secret, ca, ownerRef, "openshift-ingress", []string{"openshift"}, X509DefaultUsage, X509UsageClientServerAuth, ingressHostNames, ingressNumericIPs)
}

func ReconcileOAuthServerCert(secret, sourceSecret, ca *corev1.Secret, ownerRef config.OwnerRef) error {
	secret.Data = sourceSecret.Data
	AnnotateWithCA(secret, ca)
	return nil
}
