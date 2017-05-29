#!/bin/bash

# Update: This is no longer needed. CRDs work directly on HEAD.
# Temporary steps to set up kube-apiextensions-server

# Before running these, please run `hack/local-up-cluster.sh`
# to get the kube-apiserver (with aggregator) ready and available.

# this creates the image for the kube-apiextensions-server since it isn't published
make WHAT=vendor/k8s.io/kube-apiextensions-server/ && vendor/k8s.io/kube-apiextensions-server/hack/build-image.sh

# give yourself cluster-admin powers
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig

# creates a namespace for the kube-apiextensions-server pod
kubectl create ns kube-apiextensions

# creates the resources (permissions, RC, service, apiservice, etc) for the kube-apiextensions-server.
# This wont be necessary once we're embedded.
kubectl create -f vendor/k8s.io/kube-apiextensions-server/artifacts/example/

# create the APIService for mygroup.example.com. This wont be necessary once we're embedded.
kubectl create -f vendor/k8s.io/kube-apiextensions-server/artifacts/customresource-01/noxu-apiservice.yaml

# create a CustomResourceDefinition
kubectl create -f vendor/k8s.io/kube-apiextensions-server/artifacts/customresource-01/noxu-resource-definition.yaml

# create an instance.
kubectl create -f vendor/k8s.io/kube-apiextensions-server/artifacts/customresource-01/noxu.yaml
