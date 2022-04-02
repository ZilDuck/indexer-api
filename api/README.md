# Indexer API

Available endpoints:

GET `/contract/:contractAddr/attributes`

```
[
  {
    "traitType": string,
    "values":{
      trait1: count,
      trait2: count,
      trait3: count,
      ...
    }
  },
  ...
]
```

GET `/nft/:contractAddr`

```
[
  {
    "contract": string,
    "name":     string,
    "symbol":   string,
    "tokenId":  number, 
    "tokenUri": string, // tokenUri || baseUri + tokenId
    "type":     string  // zrc6 || zrc1
    "owner":    string,
    "burnedAt"  number, // optional 
  },
...
]
```

GET /nft/:contractAddr/:tokenId
```
{
  "contract": string,
  "name":     string,
  "symbol":   string,
  "tokenId":  number, 
  "tokenUri": string, // tokenUri || baseUri + tokenId
  "type":     string  // zrc6 || zrc1
  "owner":    string,
  "burnedAt"  number, // optional 
}
```

GET /nft/:contractAddr/:tokenId/metadata
```
{
  attributes: [
  {
    display_type: "string",
    trait_type: "Base",
    value: "Pochard"
  },
  ...
  ],
  quick_resource: string
  resource: "ipfs://string",
  resource_mimetype: "image/png",
  ...
}
```

GET /nft/:contractAddr/:tokenId/asset
```
Returns the NFT resource
```

GET /nft/:contractAddr/:tokenId/refresh

GET /wallets/:ownerAddr
```
[
  {
    "contract": string,
    "zrc6":     bool,
    "tokenIds": []number,
  },
  ...
]
```

