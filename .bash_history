cd ..
ls
cd frontendchart/
ls
cd ..
ls
helm package backendchart/
helm package frontendchart/
helm lint backendchart
helm lint frontendchart
ls
cd ..
ls
cd terraform-gke-case/
ls
terraform init
terraform plan
terraform apply
cd ..
ls
cd gke-tf-helm/
ls
terraform init
terraform plan
terraform apply --auto-approve
kubectl get nodes
terraform destroy --auto-approve
terraform init
terraform plan
terraform apply --auto-approve
gcloud container clusters get-credentials gke-tf-cluster --zone=europe-west3-a
gcloud container clusters get-credentials gke-tf-cluster --zone=europe-west3
kubectl get nodes
kubectl get svc
curl 10.16.2.61:80
kubectl get deployments
terraform destroy --auto-aprove
terraform destroy --auto-approve
ssh -V
ssh-keygen -t rsa -b 2048 -C "mustafa.ozdemir.dev@gmail.com"
cat /home/mustafa_ozdemir_dev/.ssh/id_rsa.pub
ls
rename
ls
rmdir terraform-gke-pipeline/
ls
mv terraform-gke-case/ terraform-gke-pipeline
ls
cd terraform-gke-pipeline/
ls
mv acceptance-test-master/ ./
ls
cd ..
pwd
cd terraform-gke-pipeline/
ls
mv acceptance-test-master/ /home/mustafa_ozdemir_dev
ls
mv backend-app-master/ /home/mustafa_ozdemir_dev
mv frontend-app-master/ /home/mustafa_ozdemir_dev
ls
cd ..
ls
rm README-cloudshell.txt
rm terraform-gke-pipeline/
cd terraform-gke-pipeline/
ls
cd ..
rmdir terraform-gke-pipeline/
rmdir -r  terraform-gke-pipeline/
rmdir -f terraform-gke-pipeline/
rmdir --help
rm -r terraform-gke-pipeline/
ls
rm -r terraform-gke-pipeline/
ls
cd gke-tf-infra/
ls
terraform init
terraform destroy
touch .gitignore
nano .gitignore
ls
ls -al
cat .gitignore
cd ..
ls
cd gke-tf-helm
ls
nano .gitignore
cat .gitignore
ls
cat commands.txt 
ls
rm commands.txt
ls
ls -al
cat .terraform/
ls
cd ..
ls
cd gke-tf-infra/
ls
terraform init
terraform plan
terraform apply --auto-approve
ls
cd ..
ls
cd gke-tf-helm/
ls
terraform init
terraform plan
terraform apply --auto-approve
terraform init
terraform plan
terraform apply --auto-approve
cd ..
ls
cd gke-tf-infra/
ls
terraform init
terraform plan
terraform apply --auto-approve
terraform destroy --auto-approve
cd ..
ls
cd gke-tf-helm/
terraform destroy --auto-approve
ls
cd gke-tf-infra/
ls
terraform init
terraform plan
terraform apply --auto-approve
cd ..
ls
cd gke-tf-helm/
ls
terraform init
terraform plan
terraform apply --auto-approve
ls
terraform destroy --auto-approve
ls
cd ..
ls
cd gke-tf-helm/
ls
terraform init
terraform plan
terraform apply --auto-approve
terraform refresh
terraform init
terraform plan
terraform apply --auto-approve
helm init
ls
rm frontendchart-0.1.0.tgz
ls
rm backendchart-0.1.0.tgz
ls
helm lint backendchart/
helm lint frontendchart/
terraform init
terraform plan
terraform apply --auto-approve
terraform destroy --auto-approve
helm ls --all-namespaces
gcloud container clusters get-credentials gke-tf-cluster --zone=europe-west3
gcloud config set project robotic-jet-301718
gcloud container clusters get-credentials gke-tf-cluster --zone=europe-west3
helm ls --all-namespaces
helm ls -A
ls
cd gke-tf-helm/
ls
terraform init
terraform plan
terraform apply --auto-approve
ls
terraform destroy
helm install backendchart/
helm lint backendchart/
helm lint frontendchart/
ls
helm init
$ helm repo list
helm
helm history
helm show
helm status
helm show chart
helm show all
helm show all chartd
helm show all chart
ls
helm verify backendchart/
helm lint ./backendchart
helm lint ./frontendchart/
ls
ls -al
helm template
helm init
ls
cat helm.tf
terraform init
terraform plan
terraform refresh
terraform providers
terraform apply --auto-approve
ls
cd ..
mkdir gke-helm-version
ls
cd gke-helm-version/
ls
helm create backendchart
helm create frontendchart
ls
cd ..
ls
cd gke-helm-version/
ls
helm lint backendchart/
helm lint frontendchart/
ls
helm package backendchart/
helm package frontendchart/
ls
terraform init
terraform plan
terraform apply --auto-approve
terrafrom destroy
terraform destroy
helm uninstall backend
ls
terraform init
terraform plan
terraform apply --auto-approve
kubectl get pods
kubectl describe pod backend-backendchart-9bc86844b-5hjjc 
terraform init
terraform plan
terraform apply --auto-approve
terraform destroy --auto-approve
terraform -lock=false destroy --auto-approve
terraform destroy -lock=false --auto-approve
terraform init
terraform plan
terraform plan -lock=false
terraform apply -lock=false --auto-approve
ls
kubectl get pods
kubectl describe pod backend-backendchart-c67cd5dff-bh7v9
kubectl get pods
kubectl describe pod frontend-frontendchart-6f948d65c9-24x89 
kubectl logs frontend-frontendchart-6f948d65c9-24x89 
kubectl logs backend-backendchart-c67cd5dff-bh7v9
kubectl logs frontend-frontendchart-6f948d65c9-24x89 
kubectl logs backend-backendchart-c67cd5dff-bh7v9 
ls
gitlab 
curl -LJO "https://gitlab-runner-downloads.s3.amazonaws.com/latest/deb/gitlab-runner_<arch>.deb"
dpkg -i gitlab-runner_<arch>.deb
ls
rpm -i gitlab-runner_<arch>.rpm
sudo gitlab-runner register
dpkg -i 'gitlab-runner_<arch>.deb'
sudo dpkg -i 'gitlab-runner_<arch>.deb'
ls
rm -r 'gitlab-runner_<arch>.deb'
ls
docker volume create gitlab-runner-config
docker run -d --name gitlab-runner --restart always     -v /var/run/docker.sock:/var/run/docker.sock     -v gitlab-runner-config:/etc/gitlab-runner     gitlab/gitlab-runner:latest
docker images
docker rmi 570
docker stop f012
docker stop f01
docker rm f01
docker rmi 570
ls
docker volume create gitlab-runner-config
docker run -d --name gitlab-runner --restart always     -v /var/run/docker.sock:/var/run/docker.sock     -v gitlab-runner-config:/etc/gitlab-runner     gitlab/gitlab-runner:latest
docker run -d --name gitlab-runner --restart always   -v /var/run/docker.sock:/var/run/docker.sock   -v /srv/gitlab-runner/config:/etc/gitlab-runner   gitlab/gitlab-runner:latest
docker images
docker run --rm -it -v gitlab-runner-config:/etc/gitlab-runner gitlab/gitlab-runner:latest register
ls
cd acceptance-test-master/
ls
cd çç
cd ..
ls -al
ls
cd acceptance-test-master/
ls -al
ls
ls -al
rm .gitlab-ci.yaml
ls
ls -al
touch .gitlab-ci.yml
ls
ls -al
rm .gitlab-ci.yaml
ls
rm .gitlab-ci.yml
ls
cd gke-tf-infra/
ls
ls -al
cd ..
ls
cd backend-app-master/
ls
ls -al
clear
ls -al
cd ..
ls
pwd
ls
cd backend-app-master/
ls
ls -al
nano .gitlab-ci.yaml
ls
cd backend-app-master/
ls -al
rm .gitlab-ci.yaml
rm ..gitlab-ci.yaml.swp
ls
ls -al
clear
cd ..
ls
cd frontend-app-master/
ls
ls -al
rm .gitlab-ci.yaml
ls
ls -al
rm .gitignore
cd ..
ls
ls -al
rm .gitignore
ls
cat .gitignore
clear
git init
git remote add origin git@gitlab.com:terraform-project1/gke-terraform-helm.git
git add .
git commit -m "Initial commit"
git config --global user.email "mustafa.ozdemir.dev@gmail.com"
git config --global user.name "Mustafa"
git commit -m "Initial commit"
git push -u origin master
nano .gitignore
ls
git add .
git add *
git add -f
git add . 
git status
git commit -m ".gitignore updated"
git push -u origin master
ls
cd backend-app-master/
ls
ls -al
mv .gitlab-ci.yaml .gitlab-ci.yml
ls -al
cd ..
ls
cd frontend-app-master/
ls
mv .gitlab-ci.yaml .gitlab-ci.yml
ls -al
cd ..
ls
cd gke-tf-helm/
ls
ls -al
mv .gitlab-ci.yaml .gitlab-ci.yml
ls -al
ls
cd ..
ls
cd gke-tf-infra/
ls
ls -al
mv .gitlab-ci.yaml .gitlab-ci.yml
ls -al
cd ..
git add .
git status
git commit -m ".yml update"
git push -u origin master
ls
git init
ls -al
nano .gitignore
ls -al
git add .
git status
git commit -m ".gitignore update"
git push -u origin master
ssh-keygen -t rsa -b 2048 -C "mustafa.ozdemir.dev@gmail.com"
sudo cat /home/mustafa_ozdemir_dev/.ssh/id_rsa.pub
git clone git@gitlab.com:mustafa_ozdemir/gke-terraform-helm-pipeline.git
ls
git init
git remote add origin git@gitlab.com:mustafa_ozdemir/gke-terraform-helm-pipeline.git
git add .
git commit -m "Initial commit"
git push -u origin master
ls
ls -al
ls
cd gke-helm-version/
ls
ls -al
rm .gitlab-ci.yaml
cd ..
ls -al
git add .
git status
mv .gitlab-ci.yaml .gitlab-ci.yml
git add .
git status
git commit -m "gitlab-ci.yml update"
git push -u origin master
ls -al
rm .gitlab-ci.yml
ls
ls -al
git int
git init
git add *
git add .
git status
git commit -m "root ci yaml updated"
git push -u origin master
git add .
git status
git commit -m "root ci yml updated"
git push -u origin master
git add .
git status
git commit -m "root ci yml updated"
git push -u origin master
git fetch origin master
git push -u origin master
git add .
git status
git commit -m "root ci yml updated"
git push -u origin master
git fetch origin master
git push -u origin master
git merge origin master
git fetch origin master:tmp
git rebase tmp
git push origin HEAD:master
it push origin master --force
git push origin master --force
git remote add origin git@gitlab.com:terraform-project1/gke-terraform-helm.git
git add .
git status
git commit -m "root ci yml updated"
git push origin master --force
git pull origin master --allow-unrelated-histories
$ git push -u origin master 
git push -u origin master 
git push --help
git pull origin
git add .
git push -u origin master 
cat .gitlab-ci.yml
rm .gitlab-ci.yml
ls
ls -al
git add .
git status
mv .gitlab-ci.yaml gitlab-ci.yml
git add .
git status
git commit -m "root ci yml updated"
git push -u origin master 
mv .gitlab-ci.yaml .gitlab-ci.yml
mv gitlab-ci.yml .gitlab-ci.yml
git add .
git status
git commit -m "root ci yml updated"
git push -u origin master 
ls
git add .
git status
git commit -m "README.md added"
git push -u origin master
ls -al
rm .gitlab-ci.yml
ls -al
