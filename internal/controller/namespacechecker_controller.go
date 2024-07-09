package controller

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1alpha1 "github.com/cloudsteak/component-operator.git/api/v1alpha1"
)

// NamespaceCheckerReconciler reconciles a NamespaceChecker object
type NamespaceCheckerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=api.component.cloudsteak.com,resources=namespacecheckers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=api.component.cloudsteak.com,resources=namespacecheckers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=api.component.cloudsteak.com,resources=namespacecheckers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NamespaceChecker object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.2/pkg/reconcile
func (r *NamespaceCheckerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the NamespaceChecker instance
	var namespaceChecker apiv1alpha1.NamespaceChecker
	if err := r.Get(ctx, req.NamespacedName, &namespaceChecker); err != nil {
		log.Error(err, "unable to fetch NamespaceChecker")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Check namespaces existence
	nsExist := make(map[string]bool)
	for _, ns := range namespaceChecker.Spec.Namespaces {
		namespace := &v1.Namespace{}
		err := r.Get(ctx, types.NamespacedName{Name: ns}, namespace)
		if err != nil && errors.IsNotFound(err) {
			nsExist[ns] = false
		} else if err == nil {
			nsExist[ns] = true
		} else {
			log.Error(err, "unable to check namespace")
			return ctrl.Result{}, err
		}
	}

	// Check ConfigMaps existence and read data
	configMaps := &v1.ConfigMap{}
	configMapsExist := make(map[string]bool)
	configMapsData := make(map[string]map[string]string)
	if nsExist[namespaceChecker.Spec.ConfigMapNamespace] {
		for _, cm := range namespaceChecker.Spec.ConfigMapNames {
			err := r.Get(ctx, types.NamespacedName{Name: cm, Namespace: namespaceChecker.Spec.ConfigMapNamespace}, configMaps)
			if err != nil && errors.IsNotFound(err) {
				configMapsExist[cm] = false
			} else if err == nil {
				configMapsExist[cm] = true
				configMapsData[cm] = configMaps.Data
			} else {
				log.Error(err, "unable to check configmap", "configmap", cm)
				return ctrl.Result{}, err
			}

		}
	}

	// Check Secrets existence and read data
	secrets := &v1.Secret{}
	secretsExist := make(map[string]bool)
	secretsData := make(map[string]map[string][]byte)
	if nsExist[namespaceChecker.Spec.SecretsNamespace] {
		for _, s := range namespaceChecker.Spec.SecretsNames {
			err := r.Get(ctx, types.NamespacedName{Name: s, Namespace: namespaceChecker.Spec.SecretsNamespace}, secrets)
			if err != nil && errors.IsNotFound(err) {
				secretsExist[s] = false
			} else if err == nil {
				secretsExist[s] = true
				secretsData[s] = secrets.Data
			} else {
				log.Error(err, "unable to check secret", "secret", s)
				return ctrl.Result{}, err
			}
		}
	}

	// Log the status
	log.Info("1. ####### Namespace check result", "namespaces", nsExist)
	log.Info("2. ####### ConfigMap check result", "configmaps", configMapsExist)
	log.Info("3. ####### Secret check result", "secrets", secretsExist)
	log.Info("4. ####### ConfigMap data", "configmaps", configMapsData)
	log.Info("5. ####### Secret data", "secrets", secretsData)

	// Update the status
	namespaceChecker.Status.NamespacesExist = nsExist
	namespaceChecker.Status.ConfigMapsExists = configMapsExist
	namespaceChecker.Status.ConfigMapsData = configMapsData
	namespaceChecker.Status.SecretsExists = secretsExist
	namespaceChecker.Status.SecretsData = secretsData

	if err := r.Status().Update(ctx, &namespaceChecker); err != nil {
		log.Error(err, "unable to update NamespaceChecker status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NamespaceCheckerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.NamespaceChecker{}).
		Complete(r)
}
