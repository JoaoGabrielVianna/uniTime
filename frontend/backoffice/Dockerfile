# Usa a imagem oficial do Node.js
FROM node:18-slim

# Define o diretório de trabalho
WORKDIR /app

# Copia apenas os arquivos essenciais
COPY package*.json ./

# Instala as dependências
RUN npm install

# Copia o restante do código
COPY . .

# Expondo a porta do frontend (alterar para 4000 no backoffice)
EXPOSE 4000

# Define a porta no Vite
ENV PORT=4000

# Inicia o projeto
CMD ["npm", "run", "dev"]
