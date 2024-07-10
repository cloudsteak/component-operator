# component-operator
Kubebuilder - Operator

## Commands

1. Create project

```bash
kubebuilder init --domain component.cloudsteak.com --owner "CloudSteak" --repo github.com/cloudsteak/component-operator.git --license 'none'
```

2. Create API

```bash
kubebuilder create api --kind NamespaceChecker --version v1alpha1 --group api
```

3. Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.

```bash
make generate
```

4. Create manifests (CRD, RBAC and Controller)

```bash
make manifests
```

5. Develop your code

6. Install CRDs

```bash
make install
```

7. Configure sample CRD

```bash
nano config/samples/api_v1alpha1_namespacechecker.yaml
```

```yaml
apiVersion: api.component.cloudsteak.com/v1alpha1
kind: NamespaceChecker
metadata:
  name: namespacechecker-sample
spec:
  namespaces:
  - default
  configMapNames:
  - configmap-01
  configMapsNamespace: default
  secretNames:
  - secret-1
  - secret-2
  secretsNamespace: default
  deploymentNamespace: default

```

9. Create scaler

```bash
kubectl apply -f config/samples/api_v1alpha1_namespacechecker.yaml
```

10. Run the controller

```bash
make run
```