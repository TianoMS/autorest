/*
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for
 * license information.
 * 
 * Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
 * Changes may cause incorrect behavior and will be lost if the code is
 * regenerated.
 */

'use strict';

var models = require('./index');

/**
 * @class
 * Initializes a new instance of the ClassWrapper class.
 * @constructor
 * @member {object} value
 * 
 * @member {number} [value.id]
 * 
 * @member {string} [value.name]
 * 
 */
function ClassWrapper(parameters) {
  if (parameters !== null && parameters !== undefined) {
    if (parameters.value !== null && parameters.value !== undefined) {
      this.value = new models['Product'](parameters.value);
    }
  }    
}


/**
 * Validate the payload against the ClassWrapper schema
 *
 * @param {JSON} payload
 *
 */
ClassWrapper.prototype.serialize = function () {
  var payload = {};
  if (this['value']) {
    payload['value'] = this['value'].serialize();
  }
   else {  throw new Error('this[\'value\'] cannot be null or undefined.');
  }
};

/**
 * Deserialize the instance to ClassWrapper schema
 *
 * @param {JSON} instance
 *
 */
ClassWrapper.prototype.deserialize = function (instance) {
  if (instance) {
    if (instance['value'] !== null && instance['value'] !== undefined) {
      this['value'] = new models['Product']().deserialize(instance['value']);
    }
  }
};

module.exports = ClassWrapper;
