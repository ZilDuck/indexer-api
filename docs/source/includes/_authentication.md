# Authentication

## Authenticating with the API

> To authenticate with the API you'll need to provide your api key 

```shell
# With shell, you can just pass the correct header with each request
curl "https://api.zildexr.com" \
     -H "X_API_KEY: yourZildexrApiKey"
```

```javascript
var requestOptions = {
  method: 'GET',
  redirect: 'follow',
  headers: {
    'X-API-KEY': 'yourZildexrApiKey',
  }
};

fetch("https://api.zildexr.com/", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> Make sure to replace `yourZildexrApiKey` with your API key.

<p class="first">Zildexr uses API keys to allow access to the API. You can register a new zildexr API key at <a href="https://zildexr.com">zildexr.com</a>.</p>

Zildexr expects for the API key to be included in all API requests to the server in a header that looks like the following:

`X-API-KEY: yourZildexrApiKey`

<aside class="notice">
You must replace <code>yourZildexrApiKey</code> with your zildexr API key.
</aside>





