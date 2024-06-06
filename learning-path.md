# Learning path

The learning path includes several steps outlined below:

1. - [ ] Understand the Basics of Kubernetes
2. - [ ] Learn the Concept of Operators
3. - [ ] Explore the Core Concepts Behind Operators
4. - [ ] Set Up Your Development Environment
5. - [ ] Start with simple, exemplary operator
6. - [ ] Build a Real Operator

> NOTE: This document is not a tutorial in itself. It merely guides you through external sources.
>
> NOTE: The aspect of a great GO programming language knowledge required is omitted in this doc

## 1.  Understand the Basics of Kubernetes

Before diving into Operators, ensure you have a solid understanding of Kubernetes fundamentals.

A solid understanding implies the ability to read and comprehend the official Kubernetes documentation with ease and naturally. This is a crucial requirement, as the subsequent learning steps involve extensive engagement with such materials.

At this stage, I recommend engaging with extensive courses on platforms like Udemy, Coursera, etc. It's also beneficial to get some hands-on practice during this phase. The course I attended and found helpful is the [Certified Kubernetes Administrator (CKA) with Practice Tests](https://www.udemy.com/course/certified-kubernetes-administrator-with-practice-tests/).

**Core Kubernetes Concepts**

1. - [x] **Pods**: Understand what a pod is, how it works, and its lifecycle. 
2. - [x] **Deployments and ReplicaSets**: Know how to manage a set of replicas of a pod for availability and scalability
3. - [x] **Services**: Understand how services provide a stable interface to pods.
4. - [x] **Namespaces**: Know how namespaces are used to organize resources within a cluster. 
5. - [x] **ConfigMaps and Secrets**: Understand how to manage application configuration and sensitive data. 
6. - [x] **StatefulSets**: Basic understanding, especially if you plan to manage stateful applications with Operators.:

**Intermediate Concepts**

1. - [x] **Ingress**: Understand how Ingress controllers provide external access to services. 
2. - [x] **Volumes and Persistent Storage**: Basic understanding of how Kubernetes handles storage, especially if your Operator will manage stateful applications. 
3. - [x] **Resource Limits and Requests**: Know how to allocate resources to containers. 

**Advanced Concepts (Helpful but not required initially)**

1. - [x] **Security**: Basic understanding of Kubernetes security practices, such as Role-Based Access Control (RBAC), Network Policies, and Pod Security Policies. 
2. - [x] **Networking**: Understand how pod networking works, including concepts like CNI (Container Network Interface). 
3. - [x] **Monitoring and Logging**: Familiarity with monitoring and logging practices in Kubernetes. 

## 2. Learn the Concept of Operator

Once you have a good understanding of Kubernetes, take your first look at what an Operator can be. Explore how it fits into the entire Kubernetes ecosystem and understand its purpose. The knowledge gained in this step will help you better organize and comprehend the information acquired in subsequent steps.

Here are some recommended YouTube materials:

- [Kubernetes Operator simply explained in 10 mins - TechWorld with Nana](https://youtu.be/ha3LjlD6g7g?si=BqIzTYkesRiLmunB)
- [Kuberntes Operators explained - IBM Technolog](https://youtu.be/UmIomb8aMkA?si=uWni5Kqndcl0STnA)

## 3. Explore the Core Concepts Behind Operators

First, familiarize yourself with [Objects in Kubernetes](https://kubernetes.io/docs/concepts/overview/working-with-objects/) and understand that:

- You can interact with them (Create, Read, Update, Delete) using [The Kubernetes API](https://kubernetes.io/docs/concepts/overview/kubernetes-api/).
- Each type of object has its own [Controller](https://kubernetes.io/docs/concepts/architecture/controller/).

Next, learn about defining your own custom kinds of Objects using the [Custom Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) mechanism. This step requires more time and possibly some hands-on practice. Try defining some Custom Resource Definitions (CRDs) and experiment with them. Here are some valuable resources:

- [Kubernetes crds what they are and why they are useful](https://thenewstack.io/kubernetes-crds-what-they-are-and-why-they-are-useful/)
- [Tutorial in this repo](crd-playground-lab.md)

Finally, it's a good time to delve into the [Operator Pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

The ultimate and most advanced material is [CNCF Operator White Paper](https://github.com/cncf/tag-app-delivery/blob/main/operator-wg/whitepaper/Operator-WhitePaper_v1-0.md). However, a comprehensive analysis of it is not necessary at this stage.

## 4.  Set Up Your Development Environment

All you need is a  Kubernetes Cluster, GO programming language and some text editor. 

You can follow my [My tutorial](https://github.com/0x41gawor/pdmgr/blob/master/minikube-dev-env.md) or get inspired by it.

## 5. Start with simple exemplary operator

The best way is to follow the guide from the official Kubebuilder webpage: [The cronjob tutorial](https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial). 

During this stage it is required to analyze the code and get familiar with:

- [Controller Runtime Project](https://pkg.go.dev/sigs.k8s.io/controller-runtime)
- [Controller Runtime Package](https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg)

## 6.  Build a Real Operator

- **Choose a Real-World Use Case**: Pick a more complex application or task you want to automate with an Operator.
- **Design Your Operator**: Plan what resources it will manage and how it will react to changes in those resources.