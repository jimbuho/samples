/**
 *  Xooa balance transfer JavaScript smart contract
 *
 *  Copyright 2019 Xooa
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at:
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License
 *  for the specific language governing permissions and limitations under the License.
 */
/*
 * Original source via IBM Corp:
 *  https://github.com/hyperledger/fabric-samples
 *
 * Modifications from Xooa:
 *  https://github.com/xooa/samples
 */
 

const shim = require('fabric-shim');
const util = require('util');

var Chaincode = class {

  // Initialize the chaincode
  async Init(stub) {
    console.info('========= Init =========');    
  }

  async Invoke(stub) {
    let ret = stub.getFunctionAndParameters();
    console.info(ret);
    let method = this[ret.fcn];
    if (!method) {
      console.log('no method of name:' + ret.fcn + ' found');
      return shim.error('no method of name:' + ret.fcn + ' found');
    }
    try {
      let payload = await method(stub, ret.params);
      return shim.success(payload);
    } catch (err) {
      console.log(err);
      return shim.error(err);
    }
  }

  async invoke(stub, args) {

    if (args.length != 5) {
      throw new Error('Incorrect number of arguments. Expecting 5');
    }

    let productCode = args[0];
    let step = args[1];
    let content = args[2];
    let location = args[3];
    let datetime = args[4];

    var datetimeDt = new Date(datetime);

    if (datetimeDt.valueOf() !== 0) {
      return shim.error('Expecting datetime value');
    }

    let parts = location.split(',');

    if(parts.length != 2) {
      return shim.error('Invalid geolocation');
    }

    let lon = parseFloat(parts[0]);
    let lat = parseFloat(parts[0]);

    if(lon >= -180 && lon <=180 || lat >= -180 && lat <=180) {
      return shim.error('Invalid coordinates');
    }

    const post = {
      productCode,
      step,
      content,
      location,
      datetime
    };

    let key = productCode + "-" + step;

    try {
      await ctx.stub.putState(key, Buffer.from(JSON.stringify(post)));
    } catch(err) {
      return shim.error(err);
    }
  
  }

  // Deletes an entity from state
  async delete(stub, args) {
    if (args.length != 1) {
      throw new Error('Incorrect number of arguments. Expecting 1');
    }

    let A = args[0];

    // Delete the key from the state in ledger
    await stub.deleteState(A);
  }

  // query callback representing the query of a chaincode
  async query(stub, args) {
    if (args.length != 1) {
      throw new Error('Incorrect number of arguments. Expecting name of the person to query')
    }

    let jsonResp = {};
    let A = args[0];

    // Get the state from the ledger
    let Avalbytes = await stub.getState(A);
    console.info(jsonResp);
    return Avalbytes;
    if (!Avalbytes) {
      jsonResp.error = 'Failed to get state for ' + A;
      throw new Error(JSON.stringify(jsonResp));
    }

    jsonResp.name = A;
    jsonResp.amount = Avalbytes.toString();
    console.info('Query Response:'); 
  }
};

shim.start(new Chaincode());
