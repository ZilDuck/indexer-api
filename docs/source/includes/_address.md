# Address

## Get NFTs owned by address

```shell
curl "https://api.zildexr.com/address/0x0d7cad239adeb9fa74205fb86318db21695b5655/nft" \
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

fetch("https://api.zildexr.com/address/0x0d7cad239adeb9fa74205fb86318db21695b5655/nft", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns the following JSON structure:

```json
[
  {
    "contract": "0xd793f378a925b9f0d3c4b6ee544d31c707899386",
    "zrc6": false,
    "tokenIds": [39]
  },
  {
    "contract": "0x8a79bac7a6383211ae902f34e86c6b729906346d",
    "zrc6": true,
    "tokenIds": [1, 550, 574, 578, 2426, 2427, 2428, 2429]
  }
]
```

<p class="first">Get a list of NFTs owned by a given address</p>

### HTTP Request

`GET https://api.zildexr.com/address/:ownerAddress/nft`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
ownerAddr    | The owner address to match



## Get Contracts owned by address

```shell
curl "https://api.zildexr.com/address/0x3a9398c1abfd7c9845149d68ad30bb672ba66295/contract" \
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

fetch("https://api.zildexr.com/address/0x3a9398c1abfd7c9845149d68ad30bb672ba66295/contract", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns the following JSON structure:

```json
[
  "0x1378d0e2e800302fd055d30c758b255172410195",
  "0x8a79bac7a6383211ae902f34e86c6b729906346d",
  "0xa0c7fa38c274cc14d9a63d5ac25d1180de6e9a59"
]
```

> and when details=true:

```json
[
  {
    "name": "SoullessMarketplace",
    "address": "0x1378d0e2e800302fd055d30c758b255172410195",
    "blockNum": 1851610,
    ...
  },
  {
    "name": "NonfungibleToken",
    "address": "0x8a79bac7a6383211ae902f34e86c6b729906346d",
    "blockNum": 1851193,
    ...
  },
  {
    "name": "SoullessWhitelist",
    "address": "0xa0c7fa38c274cc14d9a63d5ac25d1180de6e9a59",
    "blockNum": 1851036,
    ...
  },
  ...
]
```

<p class="first">Get a list of contracts that are owned by a given address</p>

### HTTP Request

`GET https://api.zildexr.com/address/:ownerAddress/contract?details`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
ownerAddr    | The owner address to match

### Query parameters

Parameter    | Description                                            | Default
------------ | ------------------------------------------------------ | -------
details      | Should the response provide full contract information? | false

