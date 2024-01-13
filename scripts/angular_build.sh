echo "Instalando o git..."
apt update
apt install git -y

mkdir /app && cd "$_"
echo "Clonando o repositorio..."
git clone https://github.com/wesleywillians/argocd.git

echo "Realizando o build..."
cat argocd/main.go

echo "Publicando artefatos..."
