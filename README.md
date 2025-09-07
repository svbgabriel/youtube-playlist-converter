# YouTube Playlist Converter

Um conversor de playlists do YouTube para Spotify, desenvolvido em Go. Esta ferramenta permite copiar playlists públicas do YouTube diretamente para sua conta do Spotify de forma interativa.

## 🎵 Funcionalidades

- **Conversão de playlists**: Converte playlists públicas do YouTube para o Spotify
- **Busca inteligente**: Encontra automaticamente as músicas correspondentes no Spotify
- **Seleção interativa**: Quando múltiplas opções são encontradas, permite escolher a música correta
- **Personalização**: Possibilidade de customizar nome e descrição da playlist no Spotify
- **Autenticação OAuth**: Integração segura com as APIs do YouTube e Spotify

## 📋 Pré-requisitos

- Go 1.18 ou superior
- Chave da API do YouTube Data API v3
- Credenciais de aplicativo do Spotify (Client ID e Client Secret)

## 🚀 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/youtube-playlist-converter.git
cd youtube-playlist-converter
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as credenciais (veja secção de configuração abaixo)

## ⚙️ Configuração

1. Copie o arquivo de exemplo:
```bash
cp config.yml.example config.yml
```

2. Edite o arquivo `config.yml` com as suas credenciais:
```yaml
youtube:
  key: "sua-chave-da-api-do-youtube"

spotify:
  clientId: "seu-client-id-do-spotify"
  clientSecret: "seu-client-secret-do-spotify"
```

### Obtendo as credenciais:

**YouTube API:**
1. Acesse o [Google Cloud Console](https://console.cloud.google.com/)
2. Crie um projeto ou selecione um existente
3. Ative a YouTube Data API v3
4. Gere uma chave de API

**Spotify API:**
1. Acesse o [Spotify for Developers](https://developer.spotify.com/dashboard/)
2. Crie um novo aplicativo
3. Configure o Redirect URI como: `http://localhost:8888/callback`
4. Copie o Client ID e Client Secret

## 📱 Uso

Execute o programa fornecendo o ID da playlist do YouTube como parâmetro:

```bash
go run . PLAYLIST_ID
```

### Como encontrar o ID da playlist:
- Na URL da playlist do YouTube: `https://www.youtube.com/playlist?list=PLrAWmaBOcF_ID`
- O ID é a parte após `list=`

### Fluxo de uso:
1. O programa irá buscar informações da playlist no YouTube
2. Será solicitada autenticação com a sua conta Spotify (via navegador)
3. Você poderá personalizar o nome e descrição da nova playlist
4. Para cada música encontrada:
    - Se houver apenas uma correspondência, será adicionada automaticamente
    - Se houver múltiplas opções, você escolherá qual adicionar
    - Se não encontrar correspondências, a música será pulada

## 🏗️ Estrutura do Projeto

```
youtube-playlist-converter/
├── config/
│   └── config.go          # Configurações da aplicação
├── config.yml             # Arquivo de configuração (não versionado)
├── config.yml.example     # Exemplo de configuração
├── main.go                # Arquivo principal
├── model.go               # Estruturas de dados para API do YouTube
├── spotify.go             # Funções de integração com Spotify
├── youtube.go             # Funções de integração com YouTube
├── go.mod                 # Dependências do projeto
└── README.md              # Documentação
```

## 📦 Dependências

- **viper**: Gerenciamento de configurações
- **spotify/v2**: Client oficial do Spotify para Go
- **oauth2**: Autenticação OAuth2

## 🔒 Segurança

- As credenciais são armazenadas localmente no arquivo `config.yml`
- Nunca faça o commit do arquivo `config.yml`
- A autenticação OAuth é realizada de forma segura através do navegador

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para a sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit as suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## ⚠️ Limitações

- Funciona apenas com playlists públicas do YouTube
- A qualidade da correspondência depende da similaridade dos títulos das músicas
- Algumas músicas podem não ser encontradas no Spotify devido a diferenças de disponibilidade regional
- A API do YouTube tem limites de quota diários

## 🐛 Problemas Conhecidos

- Músicas com caracteres especiais podem ter dificuldade na busca
- Playlists muito grandes podem demorar para processar
- É necessário conexão com internet durante todo o processo

---

**Nota**: Esta ferramenta é para uso pessoal e educacional. Respeite os termos de uso das APIs do YouTube e Spotify.
