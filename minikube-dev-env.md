# Minikube Development Environment

Document describes how to setup environment for k8s operators development in a very simple way. 

To develop k8s operators the first thing you need is a k8s cluster. (Actually not directly to develop them, but to test them against Custom Resources). Fortunately, you don't need a real cluster. Your operator will communicate with the cluster object by a kube-api-server. Minikube emulates the whole cluster in a single Docker container and exposes kube-api-server port for communication. From operator point of view this solution is completely transparent and sufficient. 

## Architecture

Architecture choices are influenced by my personal set of favorite tools. As a HOST OS for Minikube I've chosen the Ubuntu Server. As a host of Ubuntu Server the VritualBox Machine is applied. I am using Visual Studio Code and MobaXterm SSH clients to connect to the Ubuntu-Server Virtual Machine. 

![](img/1.png)

The Host OS in my case is Windows 10, but it is agnostic as long as you can run VirtualBox on it. 

I am using two SSH clients (Visual Studio Code and MobaXterm), because one VSC is more convenient for coding and MobaXterm for other types of activites. 