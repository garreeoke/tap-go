custom_build('garreeoke/tap-go', 
  'tanzu apps workload apply -f k8s/workload.yaml --local-path=. --yes && \
    while [[ $(kubectl get ksvc tap-go -o \'jsonpath={..status.conditions[?(@.type=="Ready")].status}\') != "True" ]]; do echo "waiting for ksvc" && sleep 10; done',
  ['pom.xml', './target/classes'],
  live_update = [
    sync('./target/classes', '/workspace/BOOT-INF/classes')
  ],
  skips_local_docker=True
)

k8s_yaml('./config/workload.yaml')
k8s_kind('Workload', image_json_path='{.metadata.run-image}')
k8s_resource(workload='sample-app-java', extra_pod_selectors=[{'serving.knative.dev/service':'sample-app-java'}])
