# Contract

## Get Contracts

```shell
curl "https://api.zildexr.com/contract" \
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

fetch("https://api.zildexr.com/contract", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

> The above command returns JSON structured like this:

```json
[
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
  },
  ...
]
```

<p class="first">This endpoint will return the details for a list of Contracts.</p>

### HTTP Request

`GET https://api.zildexr.com/contract`

### Path parameters

Parameter    | Description
------------ | ---------------------------------
contractAddr | The contract address to match

### Query parameters

Parameter    | Description                                      | Default
------------ | -------------------------------------------------|--------------
size         | Number of results to return per page             | 10
page         | The result set page to return                    | 1 
sort         | The field and direction to sort the results      | blockNum:desc
shape        | Comma delimited list of contract shapes to match | null
from         | The blockNum to start results from               | 0

sort options currently supported are:  
`sort=blockNum:asc`, `sort=blockNum:desc`

shape options supported are:  
`shape=ZRC1,ZRC2,ZRC3,ZRC4,ZRC6`

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
