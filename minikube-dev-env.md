# Minikube Development Environment

Document describes how to setup environment for k8s operators development in a very simple way. 

To develop k8s operators the first thing you need is a k8s cluster. (Actually not directly to develop them, but to test them Custom Resources running in the cluster). Fortunately, you don't need a real k8s cluster. Your operator will communicate with the cluster object by a kube-api-server. Minikube emulates the whole cluster in a single Docker container and exposes kube-api-server port for communication. From operator point of view this solution is completely transparent and sufficient. 

K8s operator can be developed with Ansible, Helm or Go. Only Go gives the access to the full power, thus this is the choice in this solution.

## Architecture

Architecture choices are influenced by my personal set of favorite tools. As a Host OS for Minikube I've chosen the Ubuntu Server. As a host of Ubuntu Server the VritualBox Machine is applied. I am using Visual Studio Code and MobaXterm SSH clients to connect to the Ubuntu-Server Virtual Machine. 

![](img/1.png)

The Host OS in my case is Windows 10, but it is agnostic as long as you can run VirtualBox on it. 

I am using two SSH clients (Visual Studio Code and MobaXterm), because VSC is more convenient for coding and MobaXterm for other types of activities. 

# Setup steps

## Intro

The setup steps are listed below:

1. Install Oracle VM VirtualBox 

2. Create Ubuntu Server VM

3. Configure Ubuntu-Server VM

4. Install Minikube dependencies

5. Install Minikube itself

6. Install Go

Brief descriptions of each step with important stuff to note can be found below. Do not execute any actions at this point. This will only scaffold our way of doing. Step by step guide will be provide in next sections. 

**1. Install VirtualBox machine**

Purposely blank.

**2. Create Ubuntu Server VM**

Create Virtual Machine choosing the Linux type. Do not provide it with ISO image during creation. We will boot it first time with ISO image inserted as a storage device. Also remember to enable second network adapter (as Host Only) before the first boot (it is important and can cause some issues).

**3. Configure Ubuntu Server VM**

It is a good practice to setup a hostname of a machine and to set a static IP address within our host-only network. It will ease the process of SSH connecting. 

**4. Install Minikube dependecies**

As it was said before Minkube needs to Docker to be hosted as a container. Of course there are some other option as Podman, Hyperkit etc. but Docker was the chosen. 

We will also install Kubectl as a program separate to Minikube. Minikube comes with built-in Kubectl but usage of it is inconvenient. 

Also we will use a trick that will make our solution even more amicable by applying alias of `k` to `kubectl`. 

**5. Install Minikube itself**

No magic here. Installation of Minikube is very easy.

**6. Install Go**

We will use Go and Kubebuilder to develop k8s operators. Nonetheless this tutorial does not cover isntallation of kubebuilder it is better to do it right before the process of development.

**Small remark**

One small remark before going further. This document will be kept as simple and concise as it can get. Due to this reason most of the setup steps are refering and delegating to external pages. This has twofold purpose. First, to omit redundancy, second to keep in mind that specifics of some instructions can change in time. It is better to delegate the reader to the official pages and this document just guides him through the choice and configuration of tools. 

## 1 Install Oracle VM VirtualBox

Follow this page: https://www.virtualbox.org/wiki/Downloads

The version I am using is 7.0.8

## 2 Create Ubuntu Server VM

### 2.1 Download Ubuntu Server ISO image.

https://ubuntu.com/download/server 

Select Option 1 - Manual Server Installation.

### 2.2 Create new "empty" VM in VirtualBox

![](img/2.png)

![](img/3.png)

Adjust Hardware and Hard Disk to you preferences. 

The settings I used was:

- 4 CPU
- 8192 MB of RAM
- 50 GB of hard disk 

But I believe it is way more than needed. I was just biased by the fact that my previous installation with the choice of 10GB of hard disk failed to the insufficient memory for Docker. 

Do not boot up the machine yet!

### 2.3 Insert iso image

Go to the created VM settings and select the storage section.

![](img/4.png)

Select the option to Add optical Driver and choose the previously downloaded `.iso` file from your host OS filesystem. 

Do not boot up the machine yet!

### 2.4 Add second network adapter

The first network adapter is NAT (used for connecting the VM to the internet).

We need to make the machine visible by our Host OS. The easiest way to do this is to plug the machine into our Host-Only network. 

To do so, apply the settings presented below to your VM:

![](img/5.png)

Do not boot up the machine yet!

### 2.5  Boot up the machine and perform installation

First boot up of the machine will emulate the installation of Linux system from Compact Disc (CD). 

Remember to (when you will be given opportunity to at some point)

- Install OpenSSH server
- Uncheck the LVM group

## 3 Configure Ubuntu Server VM

### 3.1 Setup the hostname

```sh
sudo vi /etc/hostname
```

Change the contents of the file to your preferred name. 

> My personal convention is to postfix VM names with `-vm` and to give it a name of the last installed part of software or combination of software names that are important for the cause of which I will use the machine.
>
> In this case I would name the VM "minukube-go-vm".

If you don't like `vi` you can use `nano` instead.

```sh
sudo vi /etc/hosts
```

Also change the name in the first line.

### 3.2 Setup the static IP address

```sh
cd /etc/netplan
sudo vi 00-installer-config.yaml
```

> The name of the file can vary. It should be one file present under this directory anyway.

Switch off the dns under interface belonging to the host network and assign it some IP address. Exemplary file below:

```sh
# This is the network config written by 'subiquity'
network:
  ethernets:
    enp0s3:
      dhcp4: true
    enp0s8:
      dhcp4: no
      addresses: [192.168.56.109/24]
  version: 2
```

In this example, the VM will get IP address of `192.168.56.109/24`

After saving the file perform:

```sh
sudo netplan try
sudo netplan apply
# to check if netplan applied
ifconfig # or `ip a` if you don't want to install net-tools
```

After these changes reboot the system 

```sh
sudo shutdown -r now
```

## 4 Install Minikube dependencies

### 4.1 Docker

Follow: https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository

Then (to avoid using sudo for every Docker command) execute this command:

```sh
sudo usermod -aG docker $USER
```

These changes will be in force once you log back in.

### 4.2 kubectl

Follow: https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/#install-using-native-package-management

To create the alias:

```sh
echo 'alias k="kubectl"' >> ~/.bashrc
source ~/.bashrc
```

## 5 Install Minikube itself

https://minikube.sigs.k8s.io/docs/start/

## 6 Install GO

Follow: https://go.dev/doc/install

But to preserve between reboots the PATH variable set in the tutorial above add it to `~/,bashrc` so it can be loaded each time a system starts.

```sh
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

# The end

Now you are ready to go!

Have fun :happy:



