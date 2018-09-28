#! /usr/bin

declare_env_variables() {
  #  
  export KOPS_STATE_STORE=s3://cluster1.eventor.tk

}

setup_cluster() {
  kops create cluster --name=cluster1.eventor.tk --state=s3://cluster1.eventor.tk --node-count=2 --node-size=t2.micro --master-size=t2.micro --zones=us-west-2b --networking=kubenet --yes
  kubectl create -R -f k8s/
}

main() {
  declare_env_variables
  setup_cluster
}

main