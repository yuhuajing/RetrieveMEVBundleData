# getMEVData

https://blocks.flashbots.net/

https://flashbots-explorer.marto.lol/

Get flashbots MEV data by API "https://blocks.flashbots.net/v1/blocks?block_number={blockNumber}"

responsed data
```json
{"transaction_hash":"0xa8ed3012283f5c8cd6b7509c86b377a4301bc0095c428a02144229c500039ff3",
"bundle_type":"flashbots",
"tx_index":1,
"bundle_index":8,
"block_number":17291332,
"eoa_address":"0xbEA8693896a92c4671692293D6aC660BBfecf26D",
"to_address":"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
"gas_used":184402,
"gas_price":"1056000000",
"eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"1947285120000",
"gas_fee_to_Miner":"0.00019472851200"}
```

# Get flashbot blocks' bundle info by bundle index
```json
./main --blockNumber=17291332 --bundleindex=0
```
response two transactions in bundle 0.
```json
[{"transaction_hash":"0x2d11ad37f16fa98c154ef3798019ff7587d5ed923e2c14b4fb39ed7fe23a8e91",
"bundle_type":"flashbots","tx_index":0,"bundle_index":0,"block_number":17291332,
"eoa_address":"0x94D66F5693a4032b3eB74d202E3AC2d4eB9EE747",
"to_address":"0xc80E5f7B6D94561c32C59e68ee67CdD2f613851C","gas_used":139340,
"gas_price":"164310162501","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"22894978042889340","gas_fee_to_Miner":0.02289497804} 
{"transaction_hash":"0xd2e6914bf9bad1ca66ae31cc07a540851dd49687eb4cd6772c398bb24aa5b7d0",
"bundle_type":"flashbots","tx_index":1,"bundle_index":0,"block_number":17291332,
"eoa_address":"0x9aE889Ed6a67db7d441741e90fC860f65BCd7ECe",
"to_address":"0x1111111254EEB25477B68fb85Ed929f73A960582","gas_used":117729,
"gas_price":"105384280790","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"12406785993125910","gas_fee_to_Miner":0.01240678599}]
```

# Get flashbot blocks' bundle info by sender address

```json
./main --blockNumber=17291332 --address="0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13"
```

```json
[{"transaction_hash":"0xd3445a14a4e5109a0089ddefebb97489ef59d214a1a962c29c1f4921fdea31dd",
"bundle_type":"flashbots","tx_index":0,"bundle_index":8,"block_number":17291332,
"eoa_address":"0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13",
"to_address":"0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80","gas_used":113221,"gas_price":"0",
"eth_sent_to_fee_recipient":"0","fee_recipient_eth_diff":"0","gas_fee_to_Miner":0} 
{"transaction_hash":"0xa8ed3012283f5c8cd6b7509c86b377a4301bc0095c428a02144229c500039ff3",
"bundle_type":"flashbots","tx_index":1,"bundle_index":8,"block_number":17291332,
"eoa_address":"0xbEA8693896a92c4671692293D6aC660BBfecf26D",
"to_address":"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D","gas_used":184402,
"gas_price":"1056000000","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"194728512000000","gas_fee_to_Miner":0.000194728512} 
{"transaction_hash":"0x98317ec3c7e3b6c7e741d9addc17372f36b061791ed43d5c0fa540c21cf67a78",
"bundle_type":"flashbots","tx_index":2,"bundle_index":8,"block_number":17291332,
"eoa_address":"0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13",
"to_address":"0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80","gas_used":107339,
"gas_price":"92563384919","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"9935661173820541","gas_fee_to_Miner":0.009935661174}]
```

# Get flashbot blocks' bundle info by sender address and bundle index

```json
./main --blockNumber=17291332 --address="0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13" --bundleindex=0
```

# Get flashbot blocks' bundle info by sender address from serial blocks.

```json
./main --lowblockNumber=17291327 --highblockNumber=17291332 --address="0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13"
```
responses:
```json
[{"transaction_hash":"0x8936cf617aa9b2333e4b296855b5a8590f0c3d24079d4581bd6ace335b2021d0",
"bundle_type":"flashbots","tx_index":0,"bundle_index":2,"block_number":17291327,
"eoa_address":"0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13",
"to_address":"0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80","gas_used":107592,"gas_price":"0",
"eth_sent_to_fee_recipient":"0","fee_recipient_eth_diff":"0","gas_fee_to_Miner":0} 
{"transaction_hash":"0xc7e19fb2476e6c53153b23316405cb4e641b7e2e985370dd265a2b210b1f6770",
"bundle_type":"flashbots","tx_index":1,"bundle_index":2,"block_number":17291327,
"eoa_address":"0xd3F73413dfAa7C0d09B0Bb8bDa2BeaA99a789c78",
"to_address":"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D","gas_used":125598,
"gas_price":"100000000","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"12559800000000","gas_fee_to_Miner":1.25598e-05} 
{"transaction_hash":"0x3a47728627ccdf6479488d44601800de6db8cfa83bbd02996a5114b0be702c31",
"bundle_type":"flashbots","tx_index":2,"bundle_index":2,"block_number":17291327,
"eoa_address":"0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13",
"to_address":"0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80","gas_used":94922,
"gas_price":"94203450402","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"8941979919058644","gas_fee_to_Miner":0.008941979919}]

[{"transaction_hash":"0xd3445a14a4e5109a0089ddefebb97489ef59d214a1a962c29c1f4921fdea31dd",
"bundle_type":"flashbots","tx_index":0,"bundle_index":8,"block_number":17291332,
"eoa_address":"0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13",
"to_address":"0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80","gas_used":113221,"gas_price":"0",
"eth_sent_to_fee_recipient":"0","fee_recipient_eth_diff":"0","gas_fee_to_Miner":0} 
{"transaction_hash":"0xa8ed3012283f5c8cd6b7509c86b377a4301bc0095c428a02144229c500039ff3",
"bundle_type":"flashbots","tx_index":1,"bundle_index":8,"block_number":17291332,
"eoa_address":"0xbEA8693896a92c4671692293D6aC660BBfecf26D",
"to_address":"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D","gas_used":184402,
"gas_price":"1056000000","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"194728512000000","gas_fee_to_Miner":0.000194728512} 
{"transaction_hash":"0x98317ec3c7e3b6c7e741d9addc17372f36b061791ed43d5c0fa540c21cf67a78",
"bundle_type":"flashbots","tx_index":2,"bundle_index":8,"block_number":17291332,
"eoa_address":"0xae2Fc483527B8EF99EB5D9B44875F005ba1FaE13",
"to_address":"0x6b75d8AF000000e20B7a7DDf000Ba900b4009A80","gas_used":107339,
"gas_price":"92563384919","eth_sent_to_fee_recipient":"0",
"fee_recipient_eth_diff":"9935661173820541","gas_fee_to_Miner":0.009935661174}]
```

