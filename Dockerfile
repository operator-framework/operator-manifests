FROM quay.io/openshift/origin-operator-registry:latest
COPY community-operators manifests
RUN initializer
USER root
CMD ["registry-server"]