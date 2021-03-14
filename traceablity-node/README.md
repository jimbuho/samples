# Xooa Traceability Smart Contract

This page provides an overview to Xooa Traceability Smart Contract functionalities written.

This smart contract runs on `Hyperledger Fabric` and is written in `JavaScript`.

Originally created by IBM and forked from [here](https://github.com/hyperledger/fabric-samples).

## Overview

This smart contract provides 3 functions:
  
  * Invoke
  * Query
  * Delete


#### Invoke

Invoke method is used to post a trace information
This method expects fire arguments `[ "<product code>", "<step name>", "<content>", "<location>", "<datetime>" ]`.


#### Query

Set method is used to query a post by his code formed of productCode + step
This method expects an argument `[ "<key>" ]`.


#### Delete

This method is used to delete the post from world state.
This method expects an argument `[ "<key>" ]`.


## Deploy the Xooa-Traceability smart contract 
 
1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-Traceability-Node** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from Swagger.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab in the App. 
   > 2. Click **Actions** for the identity for which token is to be regenerated.
   > 3. Select **Regenerate API Token**, and then click on **Generate**.



## Explore the Xooa-Traceability-Node smart Contract end-points

1. Go to the **API** tab, and  under API Explorer paste API Token and click **Authorize API** .

2. Scroll below and select any endpoint and click **Try it out**.

3. In the **body** field, replace placeholder with your own data to store in the blockchain.

4. Click **Execute**. 

> **Congrats!** You have saved data in a blockchain using **Xooa**.

5. To view your transaction as part of the blockchain, go to [https://xooa.com/blockchain](https://xooa.com/blockchain/ledger?utm_source=samplesRepo) or navigate to **Ledger** from your Xooa console.

6. Navigate to **Transactions** tab.

7. You can expand the data field to see your transactions.


## What is swagger.json and it's purpose

Swagger is used to simplify API development and is used by Xooa for defining REST APIs. 
Using swagger Xooa provides you an interface to make API requests.
Place **swagger.json** file along with your Smart Contract and Xooa will pick it up while deploying and will integrate it with your app.

### Steps to access [swagger](https://docs.xooa.com/API.html)

1. Go to [https://xooa.com/blockchain](https://xooa.com/blockchain/custom-smart-contracts?utm_source=samplesRepo) or navigate to **Custom Smart Contracts** from your Xooa console.

2. Click **Deploy New** or select an existing app to upgrade.

3. Once the app is deployed or upgraded successfully, Navigate to **API** tab and select **Sandbox** option and there you will find API Explorer.

4. Paste API token in the required input box and click on **Authorize API** to access swagger REST APIs.