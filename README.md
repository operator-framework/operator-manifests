[![Docker Repository on Quay](https://quay.io/repository/operatorframework/operator-manifests/status?token=ac6092db-69a1-4355-a982-555194dd2a88 "Docker Repository on Quay")](https://quay.io/repository/operatorframework/operator-manifests)

This project is a component of the [Operator Framework](https://github.com/operator-framework), an open source toolkit to manage Kubernetes native applications, called Operators, in an effective, automated, and scalable way. Read more in the [introduction blog post](https://coreos.com/blog/introducing-operator-framework).

# Operator Manifests

Contains a cumulative list of Operator manifests that can be used to manage the lifecycle of an application via the Operator Lifecycle Manager.

## Contribution

Everyone is invited to add their Operator manifests into this list. Please send a pull request and we will add your manifest to the list.

Make sure that you create a new folder in `manifests` respective to your Operator and place all necessary files such as CRDs and CSVs into that. There is no pattern for all those files, but we suggest to use the following:
* {{name}}.crd.yaml -> replace the name with the name of your CRD
* {{name}}.{{version}}.clusterserviceversion.yaml -> replace name and version with the respective info about your operator

## Usage

### Individual Operator

It's very simple:

1. Install the Operator Lifecycle Manager
2. Run `kubectl create -f <link to CRD file>` for each necessary Operator CRD
2. Run `kubectl apply -f <link to the CSV file>`

### Catalog Source

Run the following to create a new catalog source in your cluster containing all of the Operators in `/manifests`:

1. Install the Operator Lifecycle Manager
2. Run `docker run -it quay.io/operatorframework/operator-manifests > ./out`
3. Run `kubectl create -f ./out`

You can now create install plans and subscriptions for the Operator packages, and OLM will take care of creating CRDs and updates. See the [OLM catalog Operator docs](https://github.com/operator-framework/operator-lifecycle-manager/blob/master/Documentation/design/architecture.md#catalog-operator) for more information.
