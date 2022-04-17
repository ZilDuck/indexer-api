# NFTs

## Get NFTs

```shell
curl "https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d" \
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

fetch("https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
[
  {
    "contract": "0x8a79bac7a6383211ae902f34e86c6b729906346d",
    "name": "The Soulless Citadel",
    "symbol": "TSC",
    "tokenId": 1,
    "tokenUri": "https://soullesscitadel.com/metadata/chapter-1/1",
    "owner": "0x0d7cad239adeb9fa74205fb86318db21695b5655",
    "type": "ZRC6",
    "metadata": {
      "attributes": [
        {
          "trait_type": "Background",
          "value": "The Citadel"
        },
        ...
      ]
    }
  }
]
```

<p class="first">Fetch a paginated list all NFTs for a contract.</p>

### HTTP Request

`GET https://api.zildexr.com/nft/:contractAddr`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The NFT contract address to match


### URL Parameters

Parameter  | Description                                      | Default
---------- | ------------------------------------------------ | -------
size       | Number of results per pag (max 100)              | 10
page       | The page number to return                        | 1

### Pagination response headers

Header        | Value
------------- | -----
X-Pagination  | `{"size": 10, "page": 1, "total_pages": 556, "total_elements": 5555}`


## Get NFT

```shell
curl "https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1" \
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

fetch("https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
{
  "contract": "0x8a79bac7a6383211ae902f34e86c6b729906346d",
  "name": "The Soulless Citadel",
  "symbol": "TSC",
  "tokenId": 1,
  "tokenUri": "https://soullesscitadel.com/metadata/chapter-1/1",
  "owner": "0x0d7cad239adeb9fa74205fb86318db21695b5655",
  "type": "ZRC6",
  "metadata": {
    "attributes": [
      {
        "trait_type": "Background",
        "value": "The Citadel"
      },
      ...
    ]
  }
}
```

<p class="first">This endpoint will return the details for the NFT that matches the given contract address and token id.</p>

### HTTP Request

`GET https://api.zildexr.com/nft/:contractAddr/:tokenId`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The NFT contract address to match
tokenId      | The NFT token id to match


## Get NFT metadata

```shell
curl "https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1/metadata" \
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

fetch("https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1/metadata", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
{
  "attributes": [
    {
      "trait_type": "Background",
      "value": "The Citadel"
    },
    {
      "trait_type": "Weapons",
      "value": "Dragon Daggers Black"
    },
    {
      "trait_type": "Body",
      "value": "Eternal Soul"
    },
    ...
  ],
  "category": "Legend",
  "chapter": 1,
  "chapter_mint_count": 5555,
  "description": "5555 Light Walker Clan cursed to purgatory in The Soulless Citadel. Who can save their souls? We shall see...",
  "external_url": "https://soullesscitadel.com/gallery/chapter-1/1",
  "name": "Soulless #1",
  "rank": 1,
  "resources": [
    {
      "mime_type": "image/png",
      "uri": "https://soullesscitadel.com/images/chapter-1/1.png"
    }
  ]
}
```

<p class="first">This endpoint will return the metadata for the NFT that matches the given contract address and token id.</p>

### HTTP Request

`GET https://api.zildexr.com/nft/:contractAddr/:tokenId/metadata`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The NFT contract address to match
tokenId      | The NFT token id to match





## Get NFT actions

```shell
curl "https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1/actions" \
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

fetch("https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1/actions", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
[
  {
    "txId": "a50d92333e83ae7948551eb0d4df3c91ff6fe57f2520336624c6c0b58a8e47d3",
    "blockNum": 1851541,
    "action": "mint",
    "from": "",
    "to": "0xcde68a3a27554244269c2c85ee014fc1cfd89901"
  },
  {
    "txId": "b2b98f1de7fa9477e8a030efe75f5f40aa7237d4e33fac3312d21dd2a550c3ac",
    "blockNum": 1852819,
    "action": "transfer",
    "from": "0xcde68a3a27554244269c2c85ee014fc1cfd89901",
    "to": "0x0d7cad239adeb9fa74205fb86318db21695b5655"
  },
  ...
]
```

<p class="first">This endpoint will return the actions for the NFT that matches the given contract address and token id.</p>

### HTTP Request

`GET https://api.zildexr.com/nft/:contractAddr/:tokenId/actions`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The NFT contract address to match
tokenId      | The NFT token id to match


### URL Parameters

Parameter  | Description                                      | Default
---------- | ------------------------------------------------ | -------
size       | Number of results per pag (max 100)              | 10
page       | The page number to return                        | 1

### Pagination response headers

Header        | Value
------------- | -----
X-Pagination  | `{"size": 10, "page": 1, "total_pages": 10, "total_elements": 100}`




## Refresh NFT metadata

```shell
curl "https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1/refresh" \
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

fetch("https://api.zildexr.com/nft/0x8a79bac7a6383211ae902f34e86c6b729906346d/1/refresh", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
{
  "contract": "0x8a79bac7a6383211ae902f34e86c6b729906346d",
  "tokenId": 1,
  "state": true
}
```

<p class="first">This endpoint will schedule a refresh of the metadata for the NFT that matches the given contract address and token id.</p>

<p>Metadata refresh requests are queued up for processing. On a busy day this can take a few minutes to complete.</p>

### HTTP Request

`GET https://api.zildexr.com/nft/:contractAddr/:tokenId/refresh`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The NFT contract address to match
tokenId      | The NFT token id to match        

