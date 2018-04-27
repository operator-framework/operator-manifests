This project is a component of the [Operator Framework](https://github.com/operator-framework), an open source toolkit to manage Kubernetes native applications, called Operators, in an effective, automated, and scalable way. Read more in the [introduction blog post](https://coreos.com/blog/introducing-operator-framework).

# Operator Manifests

Contains a cumulative list of Operator manifests that can be used to manage the lifecycle of an application via the Operator Lifecycle Manager.

## Contribution

Everyone is invited to add their Operator manifests into this list. Please send a pull request and we will add your manifest to the list.

Make sure that you create a new folder in `manifests` respective to your Operator and place all necessary files such as CRDs and CSVs into that. There is no pattern for all those files, but we suggest to use the following:
* {{name}}.crd.yaml -> replace the name with the name of your CRD
* {{name}}.{{version}}.clusterserviceversion.yaml -> replace name and version with the respective info about your operator

## Usage

It's very simple:

1. Install the Operator Lifecycle Manager
2. Run `kubectl apply -f <link to the CSV file>`
