node -v
npm -v

mkdir /app && cd "$_"
echo "Clonando o repositorio..."
git clone https://github.com/wesleywillians/argocd.git

echo "Realizando o build..."
cat argocd/main.go

echo "Publicando artefatos..."
