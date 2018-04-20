# Openshift monitoring

This is a set of templates and containers to perform
smoke tests and montoring to Openshift components.

To start a monitor process the template:

oc process -f router-monitor-template.yaml | oc create -f -




