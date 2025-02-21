# Step 1:
 -  Create an EC2 Instance(Make sure it's in public subnet so that we can download kind)
 -  Connect to the ec2 instance
 -  (option)switch to sudo using `sudo -i `

# Step 2: Install Docker

```
sudo yum update -y
sudo yum install -y docker
sudo usermod -a -G docker ec2-user
sudo systemctl status docker
sudo systemctl start docker
```

# Step 3: Install Kind

```

[ $(uname -m) = x86_64 ] && curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.27.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind

```

# Step 4 : Create Kind Cluster

```

kind create cluster --name mycluster

```

# Step 5: Install kubectl

```

curl -O https://s3.us-west-2.amazonaws.com/amazon-eks/1.32.0/2024-12-20/bin/linux/amd64/kubectl

curl -O https://s3.us-west-2.amazonaws.com/amazon-eks/1.32.0/2024-12-20/bin/linux/amd64/kubectl.sha256

sha256sum -c kubectl.sha256

chmod +x ./kubectl

mkdir -p $HOME/bin && cp ./kubectl $HOME/bin/kubectl && export PATH=$HOME/bin:$PATH

echo 'export PATH=$HOME/bin:$PATH' >> ~/.bashrc
```

# Step 6: Connect to the cluster

```
kind get clusters

# connect to the specific cluster:

kubectl cluster-info --context kind-<cluster_name>

```

# Create a cluster with multiple nodes
1. Create a config.yaml

```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker

```

2. kind create cluster --config config.yaml
   
