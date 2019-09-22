# Facebook Crawler

Crawler para buscar comentários não respondidos em [IFrames do Facebook](https://developers.facebook.com/docs/plugins/comments/).

## Exemplo
```
url := "https://www.facebook.com/plugins/feedback.php?info=put_your_iframe_url_here?href=http://dynamicurl.com/"
user := "YOUR_OFFICIAL_USER_ID"
hasUnansweredComments, err := VerifyFacebookPage(url, user)

if err != nil {
  fmt.Println(err)
  os.Exit(1)
}

fmt.Println(hasUnansweredComments)
```

## Setup

Você precisa definir a url base do iframe e alterar apenas o parâmetro ```href```.
É necessário informar o id do usuário "oficial". Serão considerados comentários respondidos somente aqueles que tiveram pelo menos uma resposta do usuário "oficial".