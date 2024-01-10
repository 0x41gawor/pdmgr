 # LAB

```sh
ejek@minikube-vm:~/minikube$ k get loops
error: the server doesn't have a resource type "loops"
ejek@minikube-vm:~/minikube$
```

Now we will register a new type of ApiObject by called **Custom Resource**. First we need to define it with **Custom Resource Definition**. I will use such file:

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: loops.stable.gawor.com
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: stable.gawor.com
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                name:
                  type: string
                company:
                  type: string
                year:
                  type: integer
  # either Namespaced or Cluster
  scope: Namespaced
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: loops
    # singular name to be used as an alias on the CLI and for display
    singular: loop
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: Loop
    # shortNames allow shorter string to match your resource on the CLI
    shortNames:
    - lp
```

After file creation I applied it to k8s:

```sh
ejek@minikube-vm:~/minikube$ mkdir custom-resource-definitions
ejek@minikube-vm:~/minikube$ cd custom-resource-definitions/
ejek@minikube-vm:~/minikube/custom-resource-definitions$ vi loops.yaml
ejek@minikube-vm:~/minikube/custom-resource-definitions$ k apply -f loops.yaml
customresourcedefinition.apiextensions.k8s.io/loops.stable.gawor.com created
```

Now server does have a resource `loops`, but the number of objects of this resource is zero:

```sh
ejek@minikube-vm:~/minikube/custom-resource-definitions$ k get loops
No resources found in default namespace.
ejek@minikube-vm:~/minikube/custom-resource-definitions$ k get loops -A
No resources found
ejek@minikube-vm:~/minikube/custom-resource-definitions$
```

Ok, so lets create some:

```yaml
apiVersion: "stable.gawor.com/v1"
kind: Loop
metadata:
  name: ooda
spec:
  name: ooda
  company: US Air Force
  year: 1976
```

```yaml
apiVersion: "stable.gawor.com/v1"
kind: Loop
metadata:
  name: mapek
spec:
  name: mapek
  company: IBM
  year: 2000
```

```yaml
apiVersion: "stable.gawor.com/v1"
kind: Loop
metadata:
  name: focale
spec:
  name: focale
  company: Google
  year: 2005
```

Now create with:

```
k apply -f <filename>
```

And see now that:

```sh
ejek@minikube-vm:~/minikube$ k get loops
NAME     AGE
focale   2m56s
mapek    2m52s
ooda     2m46s
ejek@minikube-vm:~/minikube$
```

```sh
ejek@minikube-vm:~/minikube$ k describe loop mapek
Name:         mapek
Namespace:    default
Labels:       <none>
Annotations:  <none>
API Version:  stable.gawor.com/v1
Kind:         Loop
Metadata:
  Creation Timestamp:  2023-12-27T16:23:44Z
  Generation:          1
  Resource Version:    7852
  UID:                 05c091e4-e791-4350-a911-5c607673da51
Spec:
  Company:  IBM
  Name:     mapek
  Year:     2000
Events:     <none>
ejek@minikube-vm:~/minikube$
```

When creating the next yaml definition file for loop I added some label

```sh
ejek@minikube-vm:~/minikube$ k get loops --show-labels
NAME     AGE   LABELS
focale   10m   <none>
gana     4s    environment=earth
mapek    10m   <none>
ooda     10m   <none>
ejek@minikube-vm:~/minikube$
```

I want `k get` to display company and year of loop to do so I need to update  CRD file with this lines under `spec.versions[0]`:

```yaml
  additionalPrinterColumns:
    - name: "Company"
      type: string
      jsonPath: .spec.company
    - name: "Year"
      type: integer
      jsonPath: .spec.year
```

```sh
ejek@minikube-vm:~/minikube/custom-resource-definitions$ k apply -f loops.yaml
customresourcedefinition.apiextensions.k8s.io/loops.stable.gawor.com configured

ejek@minikube-vm:~$ k get loops
NAME     COMPANY        YEAR
focale   Google         2005
gana     Microsoft      2001
mapek    IBM            2000
ooda     US Air Force   1976
ejek@minikube-vm:~$
```
