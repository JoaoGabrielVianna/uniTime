# Usa a imagem oficial do Node.js
FROM node:18-alpine

# Define o diretório de trabalho
WORKDIR /app

# Copia apenas os arquivos essenciais
COPY package.json package-lock.json ./

# Instala as dependências
RUN npm install

# Copia o restante do código
COPY . .

# Expondo a porta do frontend (alterar para 4000 no backoffice)
EXPOSE 3000

# Define a porta no Vite
ENV PORT=3000

# Inicia o projeto
CMD ["npm", "run", "dev"]
