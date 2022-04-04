---
title: API Reference

language_tabs: # must be one of https://git.io/vQNgJ
  - shell
  - javascript

toc_footers:
  - <a href="https://zildexr.com">Sign Up for a Developer Key</a>

includes:
  - errors

search: true

code_clipboard: true

meta:
  - name: description
    content: Zildexr - API Documentation
---

# Zildex.com API

Welcome to the Zildexr API! You can use our API to access Zilliqa information for contracts, NFTs and their owners.

We have language bindings in Shell, and JavaScript! You can view code examples in the dark area to the right, and you  
can switch the programming language of the examples with the tabs in the top right.

Visit <a href="https://zildexr.com">https://zildexr.com</a> to request an API key for access.

<aside class="success">
The path addresses (contractAddr and ownerAddr) can be either a Base16 or Base32 addresses
</aside>

# Authentication

## Authenticating with the API

> To authenticate, use this code:

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
limit      | Number of results per pag (max 100)              | 10
offset     | The record offset to start the response fro  m   | 1

### Pagination response headers

Header        | Value
------------- | -----
X-Pagination  | `{"limit": 10, "offset": 1, "total_pages": 556, "total_elements": 5555}`


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


# Contract

## Get Contract

```shell
curl "https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d" \
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

fetch("https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
{
  "name": "NonfungibleToken",
  "address": "0x8a79bac7a6383211ae902f34e86c6b729906346d",
  "blockNum": 1851193,
  "data": {
    "params": [
      {"vname": "name", "type": "String", "value": "The Soulless Citadel"},
      ...
    ]
  },
  "mutableParams": [
    {"vname": "token_name", "type": "String"},
    ...
  ],
  "immutableParams": [
    {"vname": "initial_contract_owner", "type": "ByStr20"},
    ...
  ],
  "transitions": [
    {
      "name": "Mint",
      "arguments": {"to": "ByStr20", "token_uri": "String"}
    },
    ...
  ],
  "shape": {
    "ZRC6": true
  }
}
```

<p class="first">This endpoint will return the details for the Contract that matches the given contract address.</p>

### HTTP Request

`GET https://api.zildexr.com/contract/:contractAddr`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The contract address to match    


## Get Contract Code

```shell
curl "https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d/code" \
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

fetch("https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d/code", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```text
(* SPDX-License-Identifier: MIT *)
scilla_version 0

(***************************************************)
(*               Associated library                *)
(***************************************************)
import BoolUtils ListUtils IntUtils
library NonfungibleToken

type Operation =
| Add
| Sub

(* Global variables *)
let zero_address = 0x0000000000000000000000000000000000000000
let false = False
let true = True
let zero = Uint256 0
let one = Uint256 1
let empty_string = ""

let add_operation = Add
let sub_operation = Sub
let min_fee_bps = Uint128 1
let max_fee_bps = Uint128 10000
...
```

<p class="first">This endpoint will return the code for the Contract that matches the given contract address.</p>

### HTTP Request

`GET https://api.zildexr.com/contract/:contractAddr/code`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The contract address to match


## Get Contract Attributes

```shell
curl "https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d/attributes" \
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

fetch("https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d/attributes", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
{
  "Background": {
    "Cream": 716,
    "Cream Lines": 516,
    "Enlightened": 112,
    ...
  },
  "Body": {
    "Ancient": 1,
    "Blue Body": 363,
    "Blue Body Dragon Tats": 384,
    ...
  },
  "Eyes": {
    "Ancient Tears": 1,
    "Angry Glowing Eyes": 142,
    "Blood Thirsty": 1,
    ...
  },
  ...
}
```

<p class="first">This endpoint will return all of the NFT attributes for the minted NFTs along with a usage count.</p>

### HTTP Request

`GET https://api.zildexr.com/contract/:contractAddr/attributes?tokenIds=`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The contract address to match

### Query parameters

Parameter    | Description
------------ | ---------------------------------------------
tokenIds     | Comma delimited list of token ids to consider (optional)


## Get Contract State

```shell
curl "https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d/state" \
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

fetch("https://api.zildexr.com/contract/0x8a79bac7a6383211ae902f34e86c6b729906346d/state", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
{
  "_balance": 0,
  "balances": {
    "0x002fe06015ad71f26152a870342804a6f9926710": "10",
    "0x009ab69adb1fff60610019ea60df9d3bd3a08c38": "3",
    "0x00a42601c656bb0b9ace5318970336fbb3f32991": "2",
    ...
  },
  "base_uri": "https://soullesscitadel.com/metadata/chapter-1/",
  "contract_owner": "0x3a9398c1abfd7c9845149d68ad30bb672ba66295",
  ...
}
```

<p class="first">This endpoint will return the state for the given contract address.</p>

### HTTP Request

`GET https://api.zildexr.com/contract/:contractAddr/state`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The contract address to match


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

> The above command returns JSON structured like this:

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

<p class="first">Get a list of NFTs that are owned by a given address</p>

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

fetch("https://api.zildexr.com/address/0x3a9398c1abfd7c9845149d68ad30bb672ba66295/nft", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
[
  "0x1378d0e2e800302fd055d30c758b255172410195",
  "0x8a79bac7a6383211ae902f34e86c6b729906346d",
  "0xa0c7fa38c274cc14d9a63d5ac25d1180de6e9a59"
]
```

> If the details query parameter is set to true

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

Parameter    | Description
------------ | ---------------------------------
details      | Should the repsonse provide full contract infomration?

